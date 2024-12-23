package noiseMap

import (
	"api48hours/auth"
	httpRender "api48hours/httpRenderer"
	"api48hours/models"
	"api48hours/repository"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"image"
	"net/http"
	"strings"
)

func SetRoutes(r *chi.Mux) {
	r.Post("/world-map", createMap)
	r.Get("/world-map", getMap)
	r.Patch("/world-map", updateMapName)
	r.Delete("/world-map", deleteMap)
}

func createMap(w http.ResponseWriter, r *http.Request) {
	data := &request{}
	email := authUser(w, r)

	if bindData(w, r, data) || email == "unauthorized" {
		return
	}

	if repository.MySqlRepo.MapExistsToUser(email, data.Name) {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, httpRender.ErrInvalidRequest(
			errors.New(""), "map with that name already exists"),
		)
		return
	}

	mapData, err := MapCreation(data.MapCreationParams)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not create map"),
		)
		return
	}

	mapSave := models.MapSaveParams{UserEmail: email, Name: email, CreationParams: data.MapCreationParams}
	if err := repository.MySqlRepo.SaveMap(email, mapSave); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not save map"),
		)
		return
	}

	message := "map created"
	mapImage, err := MapPlainImage(&mapData)
	if err != nil {
		mapImage = image.NewRGBA(image.Rect(0, 0, 1, 1))
		message = "map created but image could not be generated"
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, httpRender.NewSuccessResponse(
		http.StatusCreated,
		response{
			message,
			mapData,
			mapImage,
		}))
}

func getMap(w http.ResponseWriter, r *http.Request) {
	data := &request{}
	email := authUser(w, r)

	if bindData(w, r, data) || email == "unauthorized" {
		return
	}

	mapData, err := repository.MySqlRepo.FindMap(email, data.Name)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not get map"),
		)
		return
	}

	message := "map found"
	mapImage, err := MapPlainImage(&mapData)
	if err != nil {
		mapImage = image.NewRGBA(image.Rect(0, 0, 1, 1))
		message = "map found but image could not be generated"
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(
		http.StatusOK,
		response{
			message,
			mapData,
			mapImage,
		}))
}

func updateMapName(w http.ResponseWriter, r *http.Request) {
	data := &request{}
	email := authUser(w, r)

	if bindData(w, r, data) || email == "unauthorized" {
		return
	}

	if err := repository.MySqlRepo.ChangeMapName(email, data.Name, data.NewName); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not update map name"),
		)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(http.StatusOK, "map name updated"))
}

func deleteMap(w http.ResponseWriter, r *http.Request) {
	data := &request{}
	email := authUser(w, r)

	if bindData(w, r, data) || email == "unauthorized" {
		return
	}

	if err := repository.MySqlRepo.DeleteMap(email, data.Name); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, httpRender.ErrServerInternal(
			err, "could not delete map"),
		)
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httpRender.NewSuccessResponse(http.StatusOK, "map deleted"))
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

func authUser(w http.ResponseWriter, r *http.Request) string {
	token := strings.Split(r.Header.Values("Authorization")[0], " ")[1]
	if emailJWT, err := auth.ValidateJWT(token); err != nil || !repository.MySqlRepo.EmailExists(emailJWT) {
		render.Status(r, http.StatusUnauthorized)
		render.Render(w, r, httpRender.ErrInvalidRequest(errors.New(""), "unauthorized"))
		return "unauthorized"
	} else {
		return emailJWT
	}
}

type request struct {
	models.MapCreationParams
	models.MapSaveParams
	models.MapUpdateParams
}

type response struct {
	string
	models.WorldMap
	image.Image
}

func (a *request) Bind(r *http.Request) error {
	return nil
}
