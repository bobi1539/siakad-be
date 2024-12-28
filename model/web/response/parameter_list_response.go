package response

import "siakad/model/domain"

type ParameterListResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func ToParameterListResponse(parameterList domain.ParameterList) ParameterListResponse {
	return ParameterListResponse{
		Id:   parameterList.Id,
		Name: parameterList.Name,
	}
}

func ToParameterListResponses(parameterLists []domain.ParameterList) []ParameterListResponse {
	if parameterLists == nil {
		return make([]ParameterListResponse, 0)
	}

	var responses []ParameterListResponse
	for _, parameterList := range parameterLists {
		responses = append(responses, ToParameterListResponse(parameterList))
	}
	return responses
}
