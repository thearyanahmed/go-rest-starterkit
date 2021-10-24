package user

import (
	"context"
	"github.com/thearyanahmed/kloudlabllc/utility"

	"github.com/thearyanahmed/kloudlabllc/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ServiceInterface interface {
	Update(context.Context, string, *UserUpdate) error
	Get(context.Context, string) (*User, error)
}

// UserService , implements UserService
// and perform user related business logics
type UserService struct {
	db         *mgo.Session
	repository Repository
	config     *config.Configuration
}
//
//// New function will initialize UserService
func NewService(userRepo Repository) ServiceInterface {
	return &UserService{repository: userRepo}
}

// Update function will update the user info
// return error if any
func (service *UserService) Update(ctx context.Context, id string, user *UserUpdate) error {
	query := bson.M{"_id": bson.ObjectIdHex(id), "isActive": true}
	CustomBson := &utility.CustomBson{}
	change, err := CustomBson.Set(user)
	if err != nil {
		return err
	}
	return service.repository.Update(ctx, query, change)
}

// Get function will find user by id
// return user and error if any
func (service *UserService) Get(ctx context.Context, id string) (*User, error) {
	return service.repository.FindOneById(ctx, id)
}
