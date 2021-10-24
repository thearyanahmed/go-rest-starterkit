package user

import (
	"context"
	"github.com/thearyanahmed/kloudlabllc/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {

	// Create , will perform db operation to save user
	// Returns modified user and error if occurs
	Create(context.Context, *User) error

	// FindAll , returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]*User, error)

	// FindOneById , find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, string) (*User, error)

	// Update , will update user data by id
	// return error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete , will remove user entry from DB
	// Return error if any
	Delete(context.Context, *User) error

	// FindOne , will find one entry of user matched by the query.
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*User, error)

	IsUserAlreadyExists(context.Context, string) bool
}

type UserRepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func NewRepository(db *mgo.Session, c *config.Configuration) Repository {
	return &UserRepositoryImp{db: db, config: c}
}

func (service *UserRepositoryImp) Create(ctx context.Context, user *User) error {
	return service.collection().Insert(user)
}

func (service *UserRepositoryImp) FindAll(ctx context.Context) ([]*User, error) {
	return nil, nil
}

func (service *UserRepositoryImp) Update(ctx context.Context, query, change interface{}) error {

	return service.collection().Update(query, change)
}

func (service *UserRepositoryImp) FindOneById(ctx context.Context, id string) (*User, error) {
	var user User
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	e := service.collection().Find(query).Select(bson.M{"password": 0, "salt": 0}).One(&user)
	return &user, e
}

func (service *UserRepositoryImp) Delete(ctx context.Context, user *User) error {
	return nil
}

func (service *UserRepositoryImp) FindOne(ctx context.Context, query interface{}) (*User, error) {
	var user User
	e := service.collection().Find(query).One(&user)
	return &user, e
}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *UserRepositoryImp) IsUserAlreadyExists(ctx context.Context, email string) bool {
	query := bson.M{"email": email}
	_, e := service.FindOne(ctx, query)
	if e != nil {
		return false
	}
	return true
}

func (service *UserRepositoryImp) collection() *mgo.Collection {
	return service.db.DB(service.config.DataBaseName).C("users")
}
