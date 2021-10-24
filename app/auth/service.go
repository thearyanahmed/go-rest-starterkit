package auth

import (
	"context"
	"github.com/thearyanahmed/kloudlabllc/app/user"
	"gopkg.in/mgo.v2/bson"
)

type ServiceInterface interface {
	Create(context.Context, *user.User) error
	Login(context.Context, *user.Credential) (*user.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}

type Service struct {
	repository user.Repository
}

func NewService(userRepo user.Repository) ServiceInterface {
	return &Service{repository: userRepo}
}

func (service *Service) Create(ctx context.Context, user *user.User) error {
	return service.repository.Create(ctx, user)
}

func (service *Service) Login(ctx context.Context, credential *user.Credential) (*user.User, error) {
	query := bson.M{"email": credential.Email}
	foundUser, err := service.repository.FindOne(ctx, query)

	if err != nil {
		return nil, err
	}

	if err = foundUser.ComparePassword(credential.Password); err != nil {
		return nil, err
	}
	return foundUser, nil

}

func (service *Service) IsUserAlreadyExists(ctx context.Context, email string) bool {
	return service.repository.IsUserAlreadyExists(ctx, email)
}
