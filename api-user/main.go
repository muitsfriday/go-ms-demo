package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muitsfriday/go-ms-demo/api-user/consumers"
	"github.com/muitsfriday/go-ms-demo/api-user/consumers/repositories"
	"github.com/muitsfriday/go-ms-demo/api-user/consumers/services"
	"gopkg.in/mgo.v2"
)

// UserLogin form data
type UserLogin struct {
	Email    string `form:"email" json:"email" bson:"email" binding:"required,email"`
	Password string `form:"password" json:"-" bson:"password" binding:"required,alphanum,gte=6,lte=20"`
}

func main() {

	ss := strings.Split(os.Getenv("MONGODB_URI"), "/")
	dbname := ss[len(ss)-1]

	// init mongo collection.
	var session, err = mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic("end")
	}
	defer session.Close()

	// init dependentcy.
	userCollection := session.DB(dbname).C("user")
	incrCollection := session.DB(dbname).C("user_counter")

	userRepo := repositories.New(userCollection, incrCollection)
	articleRepo := repositories.NewRemoteArticleRepository()
	userService := services.New(userRepo, &articleRepo)

	// init router.
	router := gin.Default()

	// ------ register api -------
	router.POST("/register", func(c *gin.Context) {
		var formData consumers.User
		if err := c.ShouldBind(&formData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := userService.Register(formData.Email, formData.Password, formData.Alias)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	// ------ login api -------
	router.POST("/login", func(c *gin.Context) {
		var formData UserLogin
		if err := c.ShouldBind(&formData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.Login(formData.Email, formData.Password)
		if err != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status": true,
			"user":   user,
		})
		return
	})

	router.GET("/user/:id/articles", func(c *gin.Context) {

		var err error

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(200, gin.H{"status": false, "err": err.Error()})
			return
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			c.JSON(200, gin.H{"status": false, "err": err.Error()})
			return
		}

		if articles, err := userService.GetArticleByOwnerID(id, page); err != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    err.Error(),
			})
			return
		} else {
			fmt.Println("arararar xxxx222", articles)
			c.JSON(200, gin.H{
				"status":   true,
				"articles": articles,
			})
		}

	})

	router.GET("/user/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		if id, err := strconv.Atoi(idParam); err != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    err.Error(),
			})
			return
		} else {
			user := userRepo.GetUser(id)
			c.JSON(200, gin.H{
				"status": user.ID == id,
				"user":   user,
			})
		}
	})

	// for test only //

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	// for init mongo
	router.GET("/seed", func(c *gin.Context) {
		var createdID int
		var err error
		createdID, err = userRepo.CreateUser(consumers.User{ID: 1, Email: "muitsfriday@gmail.com", Password: "1234", Alias: "muitsfriday"})
		createdID, err = userRepo.CreateUser(consumers.User{ID: 2, Email: "annie12@gmail.com", Password: "1234", Alias: "annie12"})
		c.JSON(200, gin.H{
			"status": false,
			"id":     createdID,
			"err":    err,
		})
	})

	// router.GET("/users", func(c *gin.Context) {
	// 	users := userRepo.()
	// 	c.JSON(200, gin.H{
	// 		"users": users,
	// 	})
	// 	return
	// })

	router.Run()

}
