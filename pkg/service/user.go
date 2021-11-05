package service

import (
	"fmt"
	"time"
	"user_api/pkg/model"
	repo "user_api/pkg/repository"
)

type Service interface {
	SearchUsers(msg model.MessageSearchUsers) (model.MessageUsers, error)
	CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error)
}

type UserService struct {
}

func (u *UserService) SearchUsers(msg model.MessageSearchUsers) (model.MessageUsers, error) {
	var user repo.Repository = &repo.UserRepository{}

	data, err := user.SearchUsers(msg)
	if err != nil {
		return model.MessageUsers{}, err
	}
	message := model.MessageUsers{Count: data.Increment}

	for id, user := range data.List {
		item := model.MessageUser{
			UserId:      id,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
		}
		message.Users = append(message.Users, item)
	}
	return message, nil
}

func (u *UserService) CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error) {
	var user repo.Repository = &repo.UserRepository{}

	data, err := user.CreateUser(msg)
	if err != nil {
		return model.MessageCreatedUser{}, err
	}

	data.CreatedAt = time.Now()
	fmt.Println("User ", data.UserId, " created!")

	message := model.MessageCreatedUser{
		UserId:    data.UserId,
		CreatedAt: data.CreatedAt,
	}

	return message, nil
}
