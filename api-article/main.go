package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muitsfriday/go-ms-demo/api-article/consumers"
	"github.com/muitsfriday/go-ms-demo/api-article/consumers/repositories"
	"github.com/muitsfriday/go-ms-demo/api-article/consumers/services"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	fmt.Println("go is running")

	router := gin.Default()

	ss := strings.Split(os.Getenv("MONGODB_URI"), "/")
	dbname := ss[len(ss)-1]

	var session, err = mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		fmt.Print(os.Getenv("MONGODB_URI"))
		panic("end")
	}
	defer session.Close()

	articleCollection := session.DB(dbname).C("article")
	incrCollection := session.DB(dbname).C("article_counter")
	remoteUserRepo := repositories.NewUserRepository()
	articleRepo := repositories.New(articleCollection, incrCollection)
	articleService := services.NewArticleService(&articleRepo, &remoteUserRepo)

	router.GET("/article/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    err.Error(),
			})
			return
		}

		if a, errFinding := articleService.GetArticleByID(id); errFinding != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    errFinding.Error(),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"status":  false,
				"article": a,
			})
		}

	})

	router.GET("/user/:id/articles", func(c *gin.Context) {
		idParam := c.Param("id")
		pageParam := c.DefaultQuery("page", "1")

		id, errParseID := strconv.Atoi(idParam)
		page, errParsePage := strconv.Atoi(pageParam)

		if errParsePage != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    errParsePage.Error(),
			})
			return
		}

		if errParseID != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    errParseID.Error(),
			})
			return
		}

		var articles []consumers.Article
		if errFinding := articleService.GetArticleByOwnerID(id, page, &articles); errFinding != nil {
			c.JSON(200, gin.H{
				"status": false,
				"err":    errFinding.Error(),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"status":   true,
				"articles": articles,
			})
		}

	})

	router.POST("/article", func(c *gin.Context) {
		var formData consumers.Article
		if err := c.ShouldBind(&formData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := articleService.CreateArticle(
			formData.Title,
			formData.Description,
			formData.Content,
			formData.Tags,
			formData.OwnerID,
		)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	// router.GET("/user/:id", func(c *gin.Context) {

	// 	idParam := c.Param("id")
	// 	id, err := strconv.Atoi(idParam)
	// 	if err != nil {
	// 		c.JSON(200, gin.H{
	// 			"status": false,
	// 			"err":    err.Error(),
	// 		})
	// 		return
	// 	}

	// 	user, err := remoteUserRepo.GetUser(id)
	// 	if err != nil {
	// 		c.JSON(200, gin.H{
	// 			"status": false,
	// 			"err":    err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"status": user.ID == id,
	// 		"user":   user,
	// 	})

	// })

	router.Run()
}
