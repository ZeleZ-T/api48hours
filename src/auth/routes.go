package auth

import (
	httpRender "api48hours/httpRenderer"
	"api48hours/models"
	"api48hours/repository"
	"errors"
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

	emailExists := repository.MySqlRepo.EmailExists(data.Email)
	if !validEmail(data.Email) || emailExists {
		message := "invalid email"
		if emailExists {
			message = "email already registered"
		}
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			errors.New(message),
			message),
		)
		return
	}
	if !validPassword(data.Password) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			errors.New("invalid password"),
			"invalid password"),
		)
		return
	}

	var err error
	if data.Password, err = hashPassword(data.Password); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not hash password"),
		)
		return
	}

	if err = repository.MySqlRepo.CreateUser(*data.User); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not create user"),
		)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(
		http.StatusCreated,
		"Registered: "+data.Email),
	)
}

func login(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	hash, _ := repository.MySqlRepo.FindUserByEmail(data.Email)

	if !verifyPassword(data.Password, hash.Password) {
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			errors.New("invalid email or password"),
			"invalid email or password"),
		)
		return
	}

	jwt := generateJWT(data.Email, time.Now())

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(http.StatusOK, jwt))
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	data := &request{}

	if !bindData(w, r, data) {
		return
	}

	if emailJWT, err := ValidateJWT(r.Header.Get("Authorization")); emailJWT != data.Email || err != nil {
		message := "invalid token"
		if err.Error() == "token expired" {
			message = "token expired"
		}
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			err, message))
		return
	}

	if !validPassword(data.Password) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			errors.New("invalid password"),
			"invalid password"),
		)
		return
	}

	var err error
	if data.Password, err = hashPassword(data.Password); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not hash password"),
		)
		return
	}

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

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(
		http.StatusOK, "account deleted"))
}

func bindData(w http.ResponseWriter, r *http.Request, data *request) bool {
	if err := render.Bind(r, data); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			err, "invalid request"))
		return true
	} else {
		return false
	}
}

type request struct {
	*models.User
}

func (a *request) Bind(r *http.Request) error {
	return nil
}