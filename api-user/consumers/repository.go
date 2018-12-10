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
	ID       int    `json:"id" bson:"id"`
	Email    string `form:"email" json:"email" bson:"email" binding:"required,email"`
	Password string `form:"password" json:"-" bson:"password" binding:"required,alphanum,gte=6,lte=20"`
	Alias    string `form:"alias" json:"alias" bson:"alias" binding:"required,gte=6,lte=32"`
}

// UserRepository interface for user repository
type UserRepository interface {
	GetUser(id int) User
	GetUserByEmail(email string) (User, error)
	CreateUser(user User) (int, error)
	UpdateUser(id int, user User) bool
	DeleteUser(id int) bool
}

// ArticleRepository interface for article repository
type ArticleRepository interface {
	GetArticleByOwnerID(id int, page int) ([]Article, error)
}
