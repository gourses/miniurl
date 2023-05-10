package api

import (
	"encoding/json"
	"fmt"
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

type AddUrlReq struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}
type AddUrlResp struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type ErrorResp struct {
	Msg string `json:"msg"`
}

func (a *API) AddUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var v AddUrlReq
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		responsJson(w, v, "Error response")
	}
	hash, err := a.handler.AddUrl(v.Url)
	if err != nil {
		return
	}
	responsJson(w, v, hash)

}

func responsJson(w http.ResponseWriter, v AddUrlReq, hash string) {
	err := json.NewEncoder(w).Encode(AddUrlResp{Url: v.Url, Hash: hash})
	if err != nil {
		//TODO: correct slog
		fmt.Println(err)
		//sys.Error.err.Error
	}
}
