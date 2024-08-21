package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status   int    `json:"status"`
	Messagge string `json:"messagge"`
}

func InternalServerError(msg string) Response {
	return errorR(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errorR(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return errorR(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return errorR(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return errorR(msg, http.StatusBadRequest)
}

func errorR(msg string, code int) Response {
	return &ErrorResponse{
		Messagge: msg,
		Status:   code,
	}
}

func (e ErrorResponse) Error() string {
	return e.Messagge
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (s ErrorResponse) GetData() interface{} {
	return nil
}
