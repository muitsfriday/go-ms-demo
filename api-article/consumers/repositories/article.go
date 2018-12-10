package repositories

import (
	"time"

	"github.com/muitsfriday/go-ms-demo/api-article/consumers"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoArticleRepository implements ArticleRepository
type MongoArticleRepository struct {
	articleCollection *mgo.Collection
	incrCollection    *mgo.Collection
}

// New create new MongoArticleRepository instance
func New(c *mgo.Collection, i *mgo.Collection) MongoArticleRepository {
	var r MongoArticleRepository
	r.articleCollection = c
	r.incrCollection = i
	return r
}

// GetArticle get article given id
func (r *MongoArticleRepository) GetArticle(id int) consumers.Article {
	var a consumers.Article
	r.articleCollection.Find(bson.M{"id": id}).One(&a)
	return a
}

// GetArticleByOwnerID get article by owner id
func (r *MongoArticleRepository) GetArticleByOwnerID(id int, page int, articles *[]consumers.Article) error {
	limit := 60
	err := r.articleCollection.Find(bson.M{"ownerId": id}).Skip((page - 1) * limit).Limit(limit).All(articles)
	return err
}

// CreateArticle create article
func (r *MongoArticleRepository) CreateArticle(a consumers.Article) (int, error) {
	if a.ID == 0 {
		id, err := r.getAutoIncrementID()
		if err != nil {
			panic(err)
		}
		a.ID = id
	}

	a.CreatedAt = time.Now().Unix()
	a.UpdatedAt = time.Now().Unix()

	if err := r.articleCollection.Insert(&a); err != nil {
		return 0, err
	}

	return a.ID, nil
}

// UpdateArticle update article
func (r *MongoArticleRepository) UpdateArticle(id int, a consumers.Article) bool {
	return false
}

// DeleteArticle delete article
func (r *MongoArticleRepository) DeleteArticle(id int) bool {
	return false
}

// increments is user id increment counter
type autoIncrement struct {
	ID int `bson:"counter"`
}

// getAutoIncrementID get next id for user record.
func (r *MongoArticleRepository) getAutoIncrementID() (int, error) {

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
		Upsert:    true,
	}
	var i autoIncrement

	_, err := r.incrCollection.Find(nil).Apply(change, &i)
	if err != nil {
		return 0, err
	}

	return i.ID, nil

}
