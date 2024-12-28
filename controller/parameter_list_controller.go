package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ParameterListController interface {
	Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
