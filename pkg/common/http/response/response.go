package response

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/vlkorniienko/novaposhta/pkg/common/http/response/json"
)

func RespondError(w http.ResponseWriter, statusCode int, logErr error, msg string) {
	r := json.New()

	log.Debug().Str("error message", logErr.Error())

	err := r.RespondError(w, statusCode, msg)
	if err != nil {
		err2 := r.RespondError(w, http.StatusInternalServerError, "error on building error response")
		if err2 != nil {
			panic(errors.Wrap(err, "server can't respond with error"))
		}
	}
}
