package services

import (
	"errors"

	"github.com/muitsfriday/go-ms-demo/api-user/consumers"
)

// UserServiceV1 implementation
type UserServiceV1 struct {
	userRepository    consumers.UserRepository
	articleRepository consumers.ArticleRepository
}

// New create new userservice instance
func New(up consumers.UserRepository, ap consumers.ArticleRepository) UserServiceV1 {
	var u UserServiceV1
	u.userRepository = up
	u.articleRepository = ap
	return u
}

// Login authentication user
func (u *UserServiceV1) Login(email string, password string) (consumers.User, error) {
	user, _ := u.userRepository.GetUserByEmail(email)

	if user.Email != email {
		return user, errors.New("user not found")
	}

	if user.Password != password {
		return user, errors.New("password incorrect")
	}

	return user, nil
}

// Register user
func (u *UserServiceV1) Register(email string, password string, alias string) (int, error) {
	user, _ := u.userRepository.GetUserByEmail(email)
	if user.Email == email {
		return 0, errors.New("email in use")
	}

	return u.userRepository.CreateUser(consumers.User{Email: email, Password: password, Alias: alias})
}

// GetArticleByOwnerID get article on remote host
func (u *UserServiceV1) GetArticleByOwnerID(id int, page int) ([]consumers.Article, error) {
	return u.articleRepository.GetArticleByOwnerID(id, page)
}
