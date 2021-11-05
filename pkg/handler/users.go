package handler

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	http_error "refactoring/pkg/errors"
	model "refactoring/pkg/model"
	repo "refactoring/pkg/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const store = `users.json`

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := repo.UserStore{}
	_ = json.Unmarshal(f, &s)

	render.JSON(w, r, s.List)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := repo.UserStore{}
	_ = json.Unmarshal(f, &s)

	request := model.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, http_error.ErrInvalidRequest(err))
		return
	}

	s.Increment++
	u := repo.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(store, b, fs.ModePerm)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
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
