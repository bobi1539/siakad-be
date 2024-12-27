package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ParameterController interface {
	Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params)
}
