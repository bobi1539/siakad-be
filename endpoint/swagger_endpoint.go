package endpoint

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetSwaggerEndpoint(router *httprouter.Router) {
	router.GET("/swagger/*filepath", swaggerHandler)
}

func swaggerHandler(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	httpSwagger.Handler(
		httpSwagger.URL(getUrlSwaggerDoc()),
		httpSwagger.DocExpansion("none"),
	).ServeHTTP(writer, httpRequest)
}

func getUrlSwaggerDoc() string {
	return "/swagger/doc.json"
}
