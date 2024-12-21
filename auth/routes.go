package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func SetRoutes(r *chi.Mux) {
	r.Post("/auth/register", register)
	r.Get("/auth/login", login)
	r.Patch("/auth/change-password", changePassword)
	r.Delete("/auth/delete-account", deleteAccount)
}

func register(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if bindData(w, r, data) {
		return
	}

	if !validEmail(data.Email) || !validPassword(data.Password) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, nil)
		return
	}

	var err error
	if data.Password, err = hashPassword(data.Password); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, nil)
		return
	}

	print(data)

	w.Write([]byte("registered"))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	if !verifyPassword(data.Password, "hash") {
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, nil)
		return
	}

	w.Write([]byte("logged in"))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	if !validPassword(data.Password) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, nil)
		return
	}

	print(data)

	w.Write([]byte("password changed"))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	// auth

	// delete

	w.Write([]byte("account deleted"))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func bindData(w http.ResponseWriter, r *http.Request, data *request) bool {
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, nil)
		return false
	} else {
		return true
	}
}

type request struct {
	*User
}

func (a *request) Bind(r *http.Request) error {
	return nil
}
