package model

import (
	"net/http"
	"time"
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }

type MessageCreatUser struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type MessageCreatedUser struct {
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
