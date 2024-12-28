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

const PARAMETER_LISTS = constant.PREFIX_API + "/parameter-lists"
const PARAMETER_LISTS_PARAMETER_LIST = PARAMETER_LISTS + "/parameter-list/:" + constant.PARAMETER_LIST_ID
const PARAMETER_LISTS_ALL = PARAMETER_LISTS + "/all"

func SetParameterListEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	parameterListController := getParameterListController(db, validate)
	router.POST(PARAMETER_LISTS, parameterListController.Create)
	router.PUT(PARAMETER_LISTS_PARAMETER_LIST, parameterListController.Update)
	router.GET(PARAMETER_LISTS_PARAMETER_LIST, parameterListController.FindById)
	router.GET(PARAMETER_LISTS_ALL, parameterListController.FindAll)
}

func getParameterListController(db *sql.DB, validate *validator.Validate) controller.ParameterListController {
	parameterListService := serviceimpl.NewParameterListService(getParameterListRepository(), db, validate)
	return controllerimpl.NewParameterListController(parameterListService)
}

func getParameterListRepository() repository.ParameterListRepository {
	return repositoryimpl.NewParameterListRepository()
}
