package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slog"
)

type Handler interface {
	AddUrl(url string) (hash string, err error)
}

type API struct {
	handler Handler
}

func Bind(r *httprouter.Router, h Handler) {
	a := &API{handler: h}

	r.POST("/api/v1/url", a.AddUrl)
}

type AddUrlRequest struct {
	Url  string `json:"url,omitempty"`
}
type AddUrlResponse struct {
	Url  string `json:"url,omitempty"`
	Hash string `json:"hash,omitempty"`
}

func (a *API) AddUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var v AddUrlRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		// TODO error
		return
	}

	hash, err := a.handler.AddUrl(v.Url)
	if err != nil {
		// TODO error
		return
	}

	err = json.NewEncoder(w).Encode(AddUrlResponse{Url: v.Url, Hash: hash})
	if err != nil {
		slog.Error(err.Error())
		return
	}
}