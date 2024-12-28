package helper

import (
	"net/http"
	"siakad/constant"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetIdFromPath(params httprouter.Params, name string) int64 {
	paramName := params.ByName(name)
	id, err := strconv.Atoi(paramName)
	PanicIfError(err)
	return int64(id)
}

func GetPageOrSize(httpRequest *http.Request, name string) int {
	value := GetQueryParam(httpRequest, name)
	if len(value) == 0 {
		switch name {
		case constant.PAGE:
			return 0
		case constant.SIZE:
			return 10
		default:
			return 0
		}
	}
	return StringToInt(value)
}

func GetQueryParam(httpRequest *http.Request, name string) string {
	queryParam := httpRequest.URL.Query()
	return queryParam.Get(name)
}
