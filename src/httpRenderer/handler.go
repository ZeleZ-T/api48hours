package httpRender

import (
	"api48hours/models"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

type HTTPErr struct {
	Err  error
	Code int
}

func (e HTTPErr) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

func New(err error, code int) *HTTPErr {
	return &HTTPErr{
		Err:  err,
		Code: code,
	}
}

func Error(err interface{}) *HTTPErr {
	switch err.(type) {
	case *HTTPErr:
		return err.(*HTTPErr)
	case error:
		return &HTTPErr{
			Err:  err.(error),
			Code: http.StatusInternalServerError,
		}
	default:
		return &HTTPErr{
			Err:  errors.New("Unknown error"),
			Code: http.StatusInternalServerError,
		}
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (rd *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Response struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data, omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int                 `json:"-"`
	Status         models.ResponseMeta `json:"status"`
	AppCode        int64               `json:"code,omitempty"`
}

func WrapHandlerFunc(route string, name string, handler http.HandlerFunc) (string, http.HandlerFunc) {
	return route, handler
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		Status: &models.ResponseMeta{
			AppStatusCode: status,
			Message:       "SUCCESS",
		},
		Data: data,
	}
}

func ErrInvalidRequest(err error, message string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusOK,
		Status: models.ResponseMeta{
			AppStatusCode: http.StatusBadRequest,
			Message:       "ERROR",
			ErrorMessage:  "Invalid Request",
			ErrorDetail:   message,
			DevMessage:    err.Error(),
		},
	}
}

func ErrServerInternal(err error, message string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Status: models.ResponseMeta{
			AppStatusCode: http.StatusInternalServerError,
			Message:       "ERROR",
			ErrorMessage:  message,
			DevMessage:    err.Error(),
		},
	}
}

var ErrNotFound = &ErrResponse{
	HTTPStatusCode: http.StatusOK,
	Status: models.ResponseMeta{
		AppStatusCode: http.StatusNotFound,
		Message:       "ERROR",
		ErrorDetail:   "Resource not found",
		ErrorMessage:  "The endpoint you were seeking burned down",
	},
}
