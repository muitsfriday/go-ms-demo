package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/muitsfriday/go-ms-demo/api-article/consumers"
)

// UserResponse struct of user response from remote host
type UserResponse struct {
	Status bool `json:"status"`
	User   struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Alias string `json:"alias"`
	} `json:"user"`
}

// RemoteUserRepository remote request to get user
type RemoteUserRepository struct {
}

// NewUserRepository new user erepository instance
func NewUserRepository() RemoteUserRepository {
	var r RemoteUserRepository
	return r
}

// GetUser from remote host
func (r *RemoteUserRepository) GetUser(id int) (consumers.User, error) {
	var ur UserResponse
	var u consumers.User

	response, err := http.Get(os.Getenv("SERVICE_USER_URI") + "/user/" + strconv.Itoa(id))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return u, errors.New("HTTP error")
	}
	data, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(data, &ur); err != nil {
		panic(err)
	}

	if ur.Status == false {
		return u, errors.New("user not found")
	}

	u.ID = ur.User.ID
	u.Alias = ur.User.Alias

	return u, nil
}
