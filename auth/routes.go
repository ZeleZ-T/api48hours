package auth

import (
	"api48hours/models"
	"api48hours/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"time"
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

	data.Password, _ = hashPassword(data.Password)

	if err := repository.MySqlRepo.CreateUser(*data.User); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, nil)
		return
	}

	w.Write([]byte("registered" + data.Email))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	hash, _ := repository.MySqlRepo.FindUserByEmail(data.Email)

	if !verifyPassword(data.Password, hash.Password) {
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, nil)
		return
	}

	jwt := generateJWT(data.Email, time.Now())

	w.Write([]byte("logged in: " + jwt))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	if emailJWT, err := ValidateJWT(r.Header.Get("Authorization")); emailJWT != data.Email || err != nil {
		if err.Error() == "token expired" {
			w.Write([]byte("token expired"))
		}
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, nil)
		return
	}

	if !validPassword(data.Password) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, nil)
		return
	}

	data.Password, _ = hashPassword(data.Password)

	if err := repository.MySqlRepo.ChangePassword(data.Email, data.Password); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, nil)
		return
	}

	w.Write([]byte("password changed"))
	render.Status(r, http.StatusOK)
	render.Render(w, r, nil)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	if emailJWT, err := ValidateJWT(r.Header.Get("Authorization")); emailJWT != data.Email || err != nil {
		if err.Error() == "token expired" {
			w.Write([]byte("token expired"))
		}
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, nil)
		return
	}

	if err := repository.MySqlRepo.DeleteAccount(data.Email); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, nil)
		return
	}

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
	*models.User
}

func (a *request) Bind(r *http.Request) error {
	return nil
}
