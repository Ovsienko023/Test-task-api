package handler

import (
	"net/http"

	http_error "user_api/pkg/errors"
	"user_api/pkg/model"
	"user_api/pkg/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	message := model.MessageSearchUsers{}
	var user service.Service = &service.UserService{}

	result, err := user.SearchUsers(message)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	message := model.MessageCreatUser{}

	if err := render.Bind(r, &message); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	var user service.Service = &service.UserService{}
	result, err := user.CreateUser(message)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	message := model.MessageGetUser{
		UserId: chi.URLParam(r, "user_id"),
	}

	var user service.Service = &service.UserService{}
	result, err := user.GetUser(message)

	if err == http_error.ErrUserNotFound {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &http.ErrAbortHandler)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	message := model.MessageUpdateUser{}

	if err := render.Bind(r, &message); err != nil {
		_ = render.Render(w, r, http_error.ErrInvalidRequest(err))
		return
	}
	message.UserId = chi.URLParam(r, "user_id")

	var user service.Service = &service.UserService{}
	err := user.UpdateUser(message)

	if err == http_error.ErrUserNotFound {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	message := model.MessageDeleteUser{
		UserId: chi.URLParam(r, "user_id"),
	}

	var user service.Service = &service.UserService{}
	err := user.DeleteUser(message)

	if err == http_error.ErrUserNotFound {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	render.Status(r, http.StatusNoContent)
}
