package controllerimpl

import (
	"net/http"
	"siakad/constant"
	"siakad/controller"
	"siakad/helper"
	"siakad/model/web/request"
	"siakad/service"

	"github.com/julienschmidt/httprouter"
)

type ParameterControllerImpl struct {
	ParameterService service.ParameterService
}

func NewParameterController(parameterService service.ParameterService) controller.ParameterController {
	return &ParameterControllerImpl{
		ParameterService: parameterService,
	}
}

func (parameterController *ParameterControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterRequest := request.ParameterRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterRequest)

	parameterResponse := parameterController.ParameterService.Create(httpRequest.Context(), parameterRequest)
	helper.WriteToResponseBody(writer, parameterResponse)
}

func (parameterController *ParameterControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterRequest := request.ParameterRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterRequest)

	id := helper.GetIdFromPath(params, constant.PARAMETER_ID)
	parameterResponse := parameterController.ParameterService.Update(httpRequest.Context(), id, parameterRequest)
	helper.WriteToResponseBody(writer, parameterResponse)
}

func (parameterController *ParameterControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	id := helper.GetIdFromPath(params, constant.PARAMETER_ID)
	parameterResponse := parameterController.ParameterService.FindById(httpRequest.Context(), id)
	helper.WriteToResponseBody(writer, parameterResponse)
}
