package serverapp

import (
	"net/http"

	"github.com/pieterclaerhout/go-webserver/v2/respond"
)

func (a *serverApp) handleLookup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ipAddress := r.URL.Query().Get("ip")

		result, err := a.db.Lookup(ipAddress)
		if err != nil {
			respond.Error(err).Write(w, r)
			return
		}

		respond.OK(result).Write(w, r)

	}
}
