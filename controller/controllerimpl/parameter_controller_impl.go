package controllerimpl

import (
	"net/http"
	"siakad/constant"
	"siakad/controller"
	"siakad/helper"
	"siakad/model/search"
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

// @Tags		Parameter
// @Accept	json
// @Produce	json
// @Param		request	body		request.ParameterRequest	true	"Request body"
// @Success	200		{object}	response.WebResponse{data=response.ParameterResponse}
// @Failure	400		{object}	response.WebResponse
// @Failure	500		{object}	response.WebResponse
// @Router	/parameters [post]
func (parameterController *ParameterControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterRequest := request.ParameterRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterRequest)

	parameterResponse := parameterController.ParameterService.Create(httpRequest.Context(), parameterRequest)
	helper.WriteSuccessResponse(writer, parameterResponse)
}

// @Tags		Parameter
// @Accept	json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Param		request	body		request.ParameterRequest	true	"Request body"
// @Success	200		{object}	response.WebResponse{data=response.ParameterResponse}
// @Failure	400		{object}	response.WebResponse
// @Failure	500		{object}	response.WebResponse
// @Router	/parameters/parameter/{id} [put]
func (parameterController *ParameterControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterRequest := request.ParameterRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterRequest)

	id := helper.GetIdFromPath(params, constant.PARAMETER_ID)
	parameterResponse := parameterController.ParameterService.Update(httpRequest.Context(), id, parameterRequest)
	helper.WriteSuccessResponse(writer, parameterResponse)
}

// @Tags		Parameter
// @Accept	json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Success	200	{object}	response.WebResponse{data=response.ParameterResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/parameters/parameter/{id} [get]
func (parameterController *ParameterControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	id := helper.GetIdFromPath(params, constant.PARAMETER_ID)
	parameterResponse := parameterController.ParameterService.FindById(httpRequest.Context(), id)
	helper.WriteSuccessResponse(writer, parameterResponse)
}

// @Tags		Parameter
// @Accept	json
// @Produce	json
// @Param   search  query   string  false  "Search"
// @Success	200	{object}	response.WebResponse{data=[]response.ParameterResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/parameters/all [get]
func (parameterController *ParameterControllerImpl) FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	generalSearch := search.GeneralSearch{
		Search:   helper.GetQueryParam(httpRequest, constant.SEARCH),
		PageSize: search.BuildPageSize(0, 0),
	}

	parameterResponses := parameterController.ParameterService.FindAll(httpRequest.Context(), generalSearch)
	helper.WriteSuccessResponse(writer, parameterResponses)
}
