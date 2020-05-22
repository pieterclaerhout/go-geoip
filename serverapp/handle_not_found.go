package serverapp

import (
	"net/http"

	"github.com/pieterclaerhout/go-webserver/v2/respond"
)

func (a *serverApp) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond.NotFound("nothing here").Write(w, r)
	}
}
