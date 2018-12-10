package consumers

// Article struct
type Article struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `form:"title" json:"title" bson:"title" binding:"required,lte=255"`
	Description string `form:"description" json:"description" bson:"description" binding:"required,lte=255"`
	Content     string `form:"content" json:"content" bson:"content" binding:"required"`
	Tags        string `form:"tags" json:"tags" bson:"tags" binding:"-"`
	OwnerID     int    `form:"ownerId" json:"ownerId" bson:"ownerId" binding:"required,gt=0"`
	OwnerAlias  string `form:"ownerAlias" json:"ownerAlias" bson:"ownerAlias" binding:"-"`
	CreatedAt   int64  `form:"createdAt" json:"createdAt" bson:"createdAt" binding:"-"`
	UpdatedAt   int64  `form:"updatedAt" json:"updatedAt" bson:"updatedAt" binding:"-"`
}

// User struct
type User struct {
	ID    int    `json:"id" bson:"id"`
	Alias string `json:"alias" bson:"alias"`
}

// ArticleRepository interface for user repository
type ArticleRepository interface {
	GetArticle(id int) Article
	GetArticleByOwnerID(id int, page int, articles *[]Article) error
	CreateArticle(article Article) (int, error)
	//UpdateArticle(id int, article Article) bool
	//DeleteArticle(id int) bool
}

// UserRepository interface
type UserRepository interface {
	GetUser(id int) (User, error)
}
