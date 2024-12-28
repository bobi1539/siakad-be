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

type ParameterListControllerImpl struct {
	ParameterListService service.ParameterListService
}

func NewParameterListController(parameterListService service.ParameterListService) controller.ParameterListController {
	return &ParameterListControllerImpl{
		ParameterListService: parameterListService,
	}
}

// @Tags		Parameter List
// @Accept	json
// @Produce	json
// @Param		request	body		request.ParameterListRequest	true	"Request body"
// @Success	200		{object}	response.WebResponse{data=response.ParameterListResponse}
// @Failure	400		{object}	response.WebResponse
// @Failure	500		{object}	response.WebResponse
// @Router	/parameter-lists [post]
func (parameterListController *ParameterListControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterListRequest := request.ParameterListRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterListRequest)

	parameterListResponse := parameterListController.ParameterListService.Create(httpRequest.Context(), parameterListRequest)
	helper.WriteSuccessResponse(writer, parameterListResponse)
}

// @Tags		Parameter List
// @Accept	json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Param		request	body		request.ParameterListRequest	true	"Request body"
// @Success	200		{object}	response.WebResponse{data=response.ParameterListResponse}
// @Failure	400		{object}	response.WebResponse
// @Failure	500		{object}	response.WebResponse
// @Router	/parameter-lists/parameter-list/{id} [put]
func (parameterListController *ParameterListControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterListRequest := request.ParameterListRequest{}
	helper.ReadFromRequestBody(httpRequest, &parameterListRequest)

	id := helper.GetIdFromPath(params, constant.PARAMETER_LIST_ID)
	parameterListResponse := parameterListController.ParameterListService.Update(httpRequest.Context(), id, parameterListRequest)
	helper.WriteSuccessResponse(writer, parameterListResponse)
}

// @Tags		Parameter List
// @Accept	json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Success	200	{object}	response.WebResponse{data=response.ParameterListResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/parameter-lists/parameter-list/{id} [get]
func (parameterListController *ParameterListControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	id := helper.GetIdFromPath(params, constant.PARAMETER_LIST_ID)
	parameterListResponse := parameterListController.ParameterListService.FindById(httpRequest.Context(), id)
	helper.WriteSuccessResponse(writer, parameterListResponse)
}

// @Tags		Parameter List
// @Accept	json
// @Produce	json
// @Param   search  			query   string  false  	"Search"
// @Param   parameterId  	query   int  		true  	"Parameter ID"
// @Success	200	{object}	response.WebResponse{data=[]response.ParameterListResponse}
// @Failure	400	{object}	response.WebResponse
// @Failure	500	{object}	response.WebResponse
// @Router	/parameter-lists/all [get]
func (parameterListController *ParameterListControllerImpl) FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	parameterId := helper.GetQueryParam(httpRequest, constant.PARAMETER_ID)
	generalSearch := search.ParameterListSearch{
		ParameterId: helper.StringToInt64(parameterId),
		GeneralSearch: search.GeneralSearch{
			Search:   helper.GetQueryParam(httpRequest, constant.SEARCH),
			PageSize: search.BuildPageSize(0, 0),
		},
	}

	parameterListResponses := parameterListController.ParameterListService.FindAll(httpRequest.Context(), generalSearch)
	helper.WriteSuccessResponse(writer, parameterListResponses)
}
