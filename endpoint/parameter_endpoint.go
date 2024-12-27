package endpoint

import (
	"database/sql"
	"siakad/constant"
	"siakad/controller"
	"siakad/controller/controllerimpl"
	"siakad/repository"
	"siakad/repository/repositoryimpl"
	"siakad/service/serviceimpl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const PARAMETERS = constant.PREFIX_API + "/parameters"
const PARAMETERS_PARAMETER = PARAMETERS + "/parameter/:" + constant.PARAMETER_ID

func SetParameterEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	parameterController := getParameterController(db, validate)
	router.POST(PARAMETERS, parameterController.Create)
	router.PUT(PARAMETERS_PARAMETER, parameterController.Update)
	router.GET(PARAMETERS_PARAMETER, parameterController.FindById)
}

func getParameterController(db *sql.DB, validate *validator.Validate) controller.ParameterController {
	parameterService := serviceimpl.NewParameterService(getParameterRepository(), db, validate)
	return controllerimpl.NewParameterController(parameterService)
}

func getParameterRepository() repository.ParameterRepository {
	return repositoryimpl.NewParameterRepository()
}
