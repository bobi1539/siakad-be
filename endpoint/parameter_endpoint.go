package endpoint

import (
	"database/sql"
	"siakad/constant"
	"siakad/controller"
	"siakad/controller/controllerimpl"
	"siakad/repository"
	"siakad/repository/repositoryimpl"
	"siakad/service"
	"siakad/service/serviceimpl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const PARAMETERS = constant.PREFIX_API + "/parameters"
const PARAMETERS_PARAMETER = PARAMETERS + "/parameter/:" + constant.PARAMETER_ID
const PARAMETERS_ALL = PARAMETERS + "/all"

func SetParameterEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	parameterController := getParameterController(db, validate)
	router.POST(PARAMETERS, parameterController.Create)
	router.PUT(PARAMETERS_PARAMETER, parameterController.Update)
	router.GET(PARAMETERS_PARAMETER, parameterController.FindById)
	router.GET(PARAMETERS_ALL, parameterController.FindAll)
}

func getParameterController(db *sql.DB, validate *validator.Validate) controller.ParameterController {
	return controllerimpl.NewParameterController(getParameterService(db, validate))
}

func getParameterService(db *sql.DB, validate *validator.Validate) service.ParameterService {
	return serviceimpl.NewParameterService(getParameterRepository(), db, validate)
}

func getParameterRepository() repository.ParameterRepository {
	return repositoryimpl.NewParameterRepository()
}
