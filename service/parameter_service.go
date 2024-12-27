package service

import (
	"context"
	"siakad/model/web/request"
	"siakad/model/web/response"
)

type ParameterService interface {
	Create(ctx context.Context, parameterRequest request.ParameterRequest) response.ParameterResponse
	Update(ctx context.Context, id int64, parameterRequest request.ParameterRequest) response.ParameterResponse
	FindById(ctx context.Context, id int64) response.ParameterResponse
}
