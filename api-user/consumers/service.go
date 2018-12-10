package consumers

// ID ja
type ID int

// UserService interface for user repository
type UserService interface {
	Login(email string, password string) (string, error)
	Register(email string, password string, alias string) (ID, error)
	GetArticleByOwnerID(id int, page int) ([]Article, error)
}
