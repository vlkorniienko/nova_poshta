package json

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type resp struct{}

type Response interface {
	RespondError(w http.ResponseWriter, statusCode int, msg string) error
}

func (j resp) RespondError(w http.ResponseWriter, statusCode int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	e := map[string]interface{}{"error": map[string]interface{}{"status": http.StatusText(statusCode), "message": msg}}
	err := json.NewEncoder(w).Encode(e)

	if err != nil {
		return errors.Wrap(err, "json encode failed")
	}

	return nil
}

func New() Response {
	return resp{}
}
