package repo

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"time"
	"user_api/pkg/model"
)

const store = `users.json`

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
}

type UserList map[string]User

type UserStore struct {
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

type Repository interface {
	CreateUser(model.MessageCreatUser) (model.MessageCreatedUser, error)
}

type UserRepository struct {
	UserStore
}

func (u *UserRepository) CreateUser(msg model.MessageCreatUser) (model.MessageCreatedUser, error) {
	fmt.Println(msg.Email, msg.DisplayName)

	data, err := ioutil.ReadFile(store)
	if err != nil {
		return model.MessageCreatedUser{}, err
	}

	err = json.Unmarshal(data, &u.UserStore)
	if err != nil {
		return model.MessageCreatedUser{}, err
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
		return model.MessageCreatedUser{}, err
	}

	err = ioutil.WriteFile(store, encode, fs.ModePerm)
	if err != nil {
		return model.MessageCreatedUser{}, err
	}
	message := model.MessageCreatedUser{
		CreatedAt: raw.CreatedAt,
		UserId:    id,
	}
	return message, nil
}
