package helper

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetIdFromPath(params httprouter.Params, name string) int64 {
	paramName := params.ByName(name)
	id, err := strconv.Atoi(paramName)
	PanicIfError(err)
	return int64(id)
}
