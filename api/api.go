package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//tämä consumer määrittää miten kutsutaan ei siis implementoija
//top down implementointi, kutsuja määrittää interfacen
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

func (a *API) AddUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
