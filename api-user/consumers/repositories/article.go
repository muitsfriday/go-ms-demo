package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/muitsfriday/go-ms-demo/api-user/consumers"
)

// ArticlesResponse for array response for http
type ArticlesResponse struct {
	Status   bool                `json:"status"`
	Articles []consumers.Article `json:"articles"`
}

// RemoteArticleRepository remote article service
type RemoteArticleRepository struct {
}

// NewRemoteArticleRepository create new service
func NewRemoteArticleRepository() RemoteArticleRepository {
	var r RemoteArticleRepository
	return r
}

// GetArticleByOwnerID list articles for userid
func (ras *RemoteArticleRepository) GetArticleByOwnerID(id int, page int) ([]consumers.Article, error) {
	var ar ArticlesResponse
	var articles []consumers.Article

	response, err := http.Get(os.Getenv("SERVICE_ARTICLE_URI") + "/user/" + strconv.Itoa(id) + "/articles")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return articles, errors.New("HTTP error")
	}

	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("aaaaaa", string(data))
	if err := json.Unmarshal(data, &ar); err != nil {
		return articles, err
	}

	if ar.Status == false {
		return articles, errors.New("user not found")
	}

	//fmt.Println("arararar xxxx", ar.Articles)

	return ar.Articles, nil
}
