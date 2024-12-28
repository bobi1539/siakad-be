package service

import (
	"context"
	"siakad/model/search"
	"siakad/model/web/request"
	"siakad/model/web/response"
)

type ParameterListService interface {
	Create(ctx context.Context, parameterListRequest request.ParameterListRequest) response.ParameterListResponse
	Update(ctx context.Context, id int64, parameterListRequest request.ParameterListRequest) response.ParameterListResponse
	FindById(ctx context.Context, id int64) response.ParameterListResponse
	FindAll(ctx context.Context, parameterListSeach search.ParameterListSearch) []response.ParameterListResponse
}
