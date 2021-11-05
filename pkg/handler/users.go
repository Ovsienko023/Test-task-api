package handler

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"net/http"

	http_error "user_api/pkg/errors"
	model "user_api/pkg/model"
	repo "user_api/pkg/repository"
	"user_api/pkg/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const store = `users.json`

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
	request := model.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &http.ErrAbortHandler)
		return
	}

	var message model.MessageCreatUser
	message.DisplayName = request.DisplayName
	message.Email = request.Email

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
	f, _ := ioutil.ReadFile(store)
	s := repo.UserStore{}
	_ = json.Unmarshal(f, &s)

	id := chi.URLParam(r, "id")

	render.JSON(w, r, s.List[id])
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := repo.UserStore{}
	_ = json.Unmarshal(f, &s)

	request := model.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, http_error.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, http_error.ErrInvalidRequest(http_error.UserNotFound))
		return
	}

	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := repo.UserStore{}
	_ = json.Unmarshal(f, &s)

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, http_error.ErrInvalidRequest(http_error.UserNotFound))
		render.Status(r, http.StatusNotFound)
		return
	}

	delete(s.List, id)

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}
