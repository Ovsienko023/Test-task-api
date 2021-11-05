package service

import (
	"fmt"
	"time"
	"user_api/pkg/model"
	repo "user_api/pkg/repository"
)

type Service interface {
	CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error)
}

type UserService struct {
}

func (u *UserService) CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error) {
	var user repo.Repository = &repo.UserRepository{}

	data, err := user.CreateUser(msg)
	if err != nil {
		return model.MessageCreatedUser{}, err
	}

	data.CreatedAt = time.Now()
	fmt.Println("User ", data.UserId, " created!")

	// return message, errors.New("barnacles")
	return data, nil
}
