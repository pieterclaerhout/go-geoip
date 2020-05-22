package serverapp

import (
	"net/http"

	"github.com/pieterclaerhout/go-webserver/v2/respond"
)

func (a *serverApp) handleMethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respond.MethodNotAllowed(r.Method+" is not supported").Write(w, r)
	}
}
