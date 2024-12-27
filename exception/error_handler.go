package exception

import (
	"net/http"
	"runtime/debug"
	"siakad/constant"
	"siakad/helper"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = helper.GetLogger()

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	log.Error(err, string(debug.Stack()))

	if badRequestError(writer, err) {
		return
	}

	if validationError(writer, err) {
		return
	}

	internalServerError(writer)
}

func badRequestError(writer http.ResponseWriter, err any) bool {
	exception, ok := err.(ErrorBadRequest)
	if ok {
		helper.WriteErrorResponse(writer, http.StatusBadRequest, exception.Error)
		return true
	}
	return false
}

func validationError(writer http.ResponseWriter, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		helper.WriteErrorResponse(writer, http.StatusBadRequest, exception.Error())
		return true
	}
	return false
}

func internalServerError(writer http.ResponseWriter) {
	helper.WriteErrorResponse(writer, http.StatusInternalServerError, constant.INTERNAL_SERVER_ERROR)
}

func PanicErrorBadRequest(err error) {
	if err != nil {
		panic(NewErrorBadRequest(err.Error()))
	}
}
