package helper

import (
	"net/http"
	"siakad/constant"
	"siakad/model/web/response"
)

func WriteSuccessResponse(writer http.ResponseWriter, data any) {
	webResponse := BuildSuccessResponse(data)
	WriteToResponseBody(writer, webResponse)
}

func WriteErrorResponse(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	webResponse := BuildErrorResponse(code, message)
	WriteToResponseBody(writer, webResponse)
}

func BuildSuccessResponse(data any) response.WebResponse {
	return response.WebResponse{
		Code:    200,
		Message: constant.SUCCESS,
		Data:    data,
	}
}

func BuildErrorResponse(code int, message string) response.WebResponse {
	return response.WebResponse{
		Code:    code,
		Message: message,
	}
}
