package consumers

// ArticleService Interface
type ArticleService interface {
	GetArticleByID(id int) (Article, error)
	GetArticleByOwnerID(id int, page int, articles *[]Article) error
	CreateArticle(title string, description string, content string, tags string, ownerID int) (int, error)
}
