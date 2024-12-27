package main

import (
	"net/http"
	"os"
	"siakad/app"
	"siakad/constant"
	"siakad/endpoint"
	"siakad/exception"
	"siakad/helper"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	endpoint.SetParameterEndpoint(router, db, validate)
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
