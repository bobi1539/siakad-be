package main

import (
	"net/http"
	"os"
	"siakad/app"
	"siakad/constant"
	_ "siakad/docs"
	"siakad/endpoint"
	"siakad/exception"
	"siakad/helper"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//	@title			SIAKAD API
//	@version		1.0
//	@description	This is a siakad api service.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/api
func main() {
	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	endpoint.SetSwaggerEndpoint(router)
	endpoint.SetParameterEndpoint(router, db, validate)
	endpoint.SetParameterListEndpoint(router, db, validate)
	router.PanicHandler = exception.ErrorHandler

	log := helper.GetLogger()
	log.Info("Start siakad application")

	runServer(router)
}

func runServer(router *httprouter.Router) {
	server := http.Server{
		Addr:    os.Getenv(constant.APP_HOST),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
