package repo

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"
	http_error "user_api/pkg/errors"
	"user_api/pkg/model"
)

const store = `users.json`

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

type UserList map[string]User

type UserCreate struct {
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

type Repository interface {
	GetUser(msg model.MessageGetUser) (User, error)
	SearchUsers(msg model.MessageSearchUsers) (UserStore, error)
	CreateUser(msg model.MessageCreatUser) (UserCreate, error)
	UpdateUser(msg model.MessageUpdateUser) error
	DeleteUser(msg model.MessageDeleteUser) error
}

type UserRepository struct {
	UserStore
}

func (u *UserRepository) SearchUsers(msg model.MessageSearchUsers) (UserStore, error) {
	data, err := ioutil.ReadFile(store)
	if err != nil {
		return UserStore{}, err
	}

	err = json.Unmarshal(data, &u.UserStore)
	if err != nil {
		return UserStore{}, err
	}

	return u.UserStore, nil
}

func (u *UserRepository) GetUser(msg model.MessageGetUser) (User, error) {
	data, err := ioutil.ReadFile(store)
	if err != nil {
		return User{}, nil
	}
	userStore := UserStore{}

	err = json.Unmarshal(data, &userStore)
	if err != nil {
		return User{}, nil
	}

	if _, ok := userStore.List[msg.UserId]; !ok {
		return User{}, http_error.ErrUserNotFound
	}

	user := userStore.List[msg.UserId]
	message := User{
		CreatedAt:   user.CreatedAt,
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}

	return message, nil
}

func (u *UserRepository) CreateUser(msg model.MessageCreatUser) (UserCreate, error) {
	data, err := ioutil.ReadFile(store)
	if err != nil {
		return UserCreate{}, err
	}

	err = json.Unmarshal(data, &u.UserStore)
	if err != nil {
		return UserCreate{}, err
	}

	u.UserStore.Increment++
	raw := User{
		CreatedAt:   time.Now(),
		DisplayName: msg.DisplayName,
		Email:       msg.Email,
	}

	id := strconv.Itoa(u.UserStore.Increment)
	u.UserStore.List[id] = raw

	encode, err := json.Marshal(&u.UserStore)
	if err != nil {
		return UserCreate{}, err
	}

	err = ioutil.WriteFile(store, encode, fs.ModePerm)
	if err != nil {
		return UserCreate{}, err
	}
	message := UserCreate{
		CreatedAt: raw.CreatedAt,
		UserId:    id,
	}
	return message, nil
}

func (u *UserRepository) UpdateUser(msg model.MessageUpdateUser) error {
	data, err := ioutil.ReadFile(store)
	if err != nil {
		return err
	}
	userStore := UserStore{}

	err = json.Unmarshal(data, &userStore)
	if err != nil {
		return err
	}

	if _, ok := userStore.List[msg.UserId]; !ok {
		return http_error.ErrUserNotFound
	}

	user := userStore.List[msg.UserId]
	if msg.DisplayName != "" {
		user.DisplayName = msg.DisplayName
	}
	if msg.Email != "" {
		user.Email = msg.Email
	}
	userStore.List[msg.UserId] = user

	encode, err := json.Marshal(&userStore)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(store, encode, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(msg model.MessageDeleteUser) error {
	data, err := ioutil.ReadFile(store)
	if err != nil {
		return err
	}
	userStore := UserStore{}

	err = json.Unmarshal(data, &userStore)
	if err != nil {
		return err
	}

	if _, ok := userStore.List[msg.UserId]; !ok {
		return http_error.ErrUserNotFound
	}

	delete(userStore.List, msg.UserId)

	encode, err := json.Marshal(&userStore)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(store, encode, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
