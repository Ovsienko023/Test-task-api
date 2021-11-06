package model

import (
	"net/http"
	"time"
)

type MessageGetUser struct {
	UserId string `json:"user_id"`
}

type MessageSearchUsers struct{}

type MessageUser struct {
	UserId      string    `json:"id"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}

type MessageUsers struct {
	Count int           `json:"count"`
	Users []MessageUser `json:"users"`
}

type MessageCreatUser struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *MessageCreatUser) Bind(r *http.Request) error { return nil }

type MessageCreatedUser struct {
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageUpdateUser struct {
	UserId      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *MessageUpdateUser) Bind(r *http.Request) error { return nil }

type MessageDeleteUser struct {
	UserId string `json:"user_id"`
}
