package repositories

import (
	"github.com/muitsfriday/go-ms-demo/api-user/consumers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// autoIncrement is user id increment counter
type autoIncrement struct {
	ID int `bson:"counter"`
}

// UserRepositoryMongoDB implement UserRepository
type UserRepositoryMongoDB struct {
	userCollection *mgo.Collection
	incrCollection *mgo.Collection
}

// New factory function
func New(userCollection *mgo.Collection, incrCollection *mgo.Collection) consumers.UserRepository {

	return &UserRepositoryMongoDB{userCollection, incrCollection}
}

// GetUser given id
func (r *UserRepositoryMongoDB) GetUser(id int) consumers.User {

	var user consumers.User
	r.userCollection.Find(bson.M{"id": id}).One(&user)

	return user
}

// CreateUser create new user given user scheme
func (r *UserRepositoryMongoDB) CreateUser(u consumers.User) (int, error) {

	if u.ID == 0 {
		if id, err := r.getAutoIncrementID(); err != nil {
			return 0, err
		} else {
			u.ID = id
		}
	}

	if err := r.userCollection.Insert(&u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// UpdateUser update user info given user'id
func (r *UserRepositoryMongoDB) UpdateUser(id int, u consumers.User) bool {
	return false
}

// DeleteUser delete user by id
func (r *UserRepositoryMongoDB) DeleteUser(id int) bool {
	return false
}

// ListUser list all users
func (r *UserRepositoryMongoDB) ListUser() []consumers.User {
	var list []consumers.User
	err := r.userCollection.Find(nil).All(&list)
	if err != nil {
		return list
	}
	return list
}

// GetUserByEmail get user given email
func (r *UserRepositoryMongoDB) GetUserByEmail(email string) (consumers.User, error) {
	var u consumers.User
	err := r.userCollection.Find(bson.M{"email": email}).One(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

// getAutoIncrementID get next id for user record.
func (r *UserRepositoryMongoDB) getAutoIncrementID() (int, error) {

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
