package helper

import (
	"encoding/json"
	"net/http"
	"siakad/constant"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add(constant.CONTENT_TYPE, constant.APPLICATION_JSON)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
