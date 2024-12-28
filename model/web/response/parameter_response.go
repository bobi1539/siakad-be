package response

import "siakad/model/domain"

type ParameterResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToParameterResponse(parameter domain.Parameter) ParameterResponse {
	return ParameterResponse{
		Id:          parameter.Id,
		Name:        parameter.Name,
		Description: parameter.Description,
	}
}

func ToParameterResponses(parameters []domain.Parameter) []ParameterResponse {
	if parameters == nil {
		return make([]ParameterResponse, 0)
	}

	var responses []ParameterResponse
	for _, parameter := range parameters {
		responses = append(responses, ToParameterResponse(parameter))
	}
	return responses
}
