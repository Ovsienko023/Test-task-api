package service

import (
	"fmt"
	"time"
	"user_api/pkg/model"
	repo "user_api/pkg/repository"
)

type Service interface {
	GetUser(msg model.MessageGetUser) (model.MessageUser, error)
	SearchUsers(msg model.MessageSearchUsers) (model.MessageUsers, error)
	CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error)
	DeleteUser(msg model.MessageDeleteUser) error
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

func (u *UserService) GetUser(msg model.MessageGetUser) (model.MessageUser, error) {
	var user repo.Repository = &repo.UserRepository{}

	data, err := user.GetUser(msg)
	if err != nil {
		return model.MessageUser{}, err
	}
	message := model.MessageUser{
		UserId:      msg.UserId,
		DisplayName: data.DisplayName,
		Email:       data.Email,
		CreatedAt:   data.CreatedAt,
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

func (u *UserService) DeleteUser(msg model.MessageDeleteUser) error {
	var user repo.Repository = &repo.UserRepository{}
	err := user.DeleteUser(msg)
	if err != nil {
		return err
	}

	fmt.Println("User ", msg.UserId, " deleted!")

	return nil
}
