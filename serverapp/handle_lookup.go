package serverapp

import (
	"errors"
	"net/http"

	"github.com/pieterclaerhout/go-webserver/v2/binder"
	"github.com/pieterclaerhout/go-webserver/v2/respond"
)

func (a *serverApp) handleLookup() http.HandlerFunc {

	type request struct {
		IPAddress string `json:"ip" form:"ip" query:"ip"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := binder.Bind(r, &req); err != nil {
			respond.Error(err).Write(w, r)
			return
		}

		if req.IPAddress == "" {
			respond.Error(errors.New("No IP address specified")).Write(w, r)
			return
		}

		result, err := a.db.Lookup(req.IPAddress)
		if err != nil {
			respond.Error(err).Write(w, r)
			return
		}

		respond.OK(result).Write(w, r)

	}
}
