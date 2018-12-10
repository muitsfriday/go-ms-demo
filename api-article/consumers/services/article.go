package services

import (
	"errors"

	"github.com/muitsfriday/go-ms-demo/api-article/consumers"
)

// ArticleServiceV1 version 1
type ArticleServiceV1 struct {
	articleRepo consumers.ArticleRepository
	userRepo    consumers.UserRepository
}

// NewArticleService create new instance
func NewArticleService(ar consumers.ArticleRepository, ur consumers.UserRepository) ArticleServiceV1 {
	var a ArticleServiceV1
	a.articleRepo = ar
	a.userRepo = ur
	return a
}

// GetArticleByID get article given id
func (as *ArticleServiceV1) GetArticleByID(id int) (consumers.Article, error) {
	article := as.articleRepo.GetArticle(id)
	if article.ID != id {
		return article, errors.New("article not found")
	}
	return article, nil
}

// GetArticleByOwnerID get article from id owner , with page
func (as *ArticleServiceV1) GetArticleByOwnerID(id int, page int, articles *[]consumers.Article) error {
	return as.articleRepo.GetArticleByOwnerID(id, page, articles)
}

// CreateArticle create new article
func (as *ArticleServiceV1) CreateArticle(title string, description string, content string, tags string, ownerID int) (int, error) {

	user, err := as.userRepo.GetUser(ownerID)
	if err != nil {
		return 0, err
	}

	if user.ID != ownerID {
		return 0, errors.New("owner id miss match")
	}

	id, errCreate := as.articleRepo.CreateArticle(consumers.Article{
		Title:       title,
		Description: description,
		Content:     content,
		Tags:        tags,
		OwnerID:     ownerID,
		OwnerAlias:  user.Alias,
	})

	if errCreate != nil {
		return id, errCreate
	}

	return id, nil
}
