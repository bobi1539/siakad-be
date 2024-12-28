package serviceimpl

import (
	"context"
	"database/sql"
	"siakad/exception"
	"siakad/helper"
	"siakad/model/domain"
	"siakad/model/dto"
	"siakad/model/search"
	"siakad/model/web/request"
	"siakad/model/web/response"
	"siakad/repository"
	"siakad/service"

	"github.com/go-playground/validator/v10"
)

type ParameterListServiceImpl struct {
	ParameterListRepository repository.ParameterListRepository
	DB                      *sql.DB
	validate                *validator.Validate
}

func NewParameterListService(parameterListRepository repository.ParameterListRepository, db *sql.DB, validate *validator.Validate) service.ParameterListService {
	return &ParameterListServiceImpl{
		ParameterListRepository: parameterListRepository,
		DB:                      db,
		validate:                validate,
	}
}

func (parameterListService *ParameterListServiceImpl) Create(ctx context.Context, parameterListRequest request.ParameterListRequest) response.ParameterListResponse {
	parameterListService.validateRequest(parameterListRequest)

	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList := domain.ParameterList{
		Name: parameterListRequest.Name,
	}
	parameterList = parameterListService.ParameterListRepository.Save(repoCtx, parameterList)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) Update(ctx context.Context, id int64, parameterListRequest request.ParameterListRequest) response.ParameterListResponse {
	parameterListService.validateRequest(parameterListRequest)

	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList := parameterListService.findParameterListById(repoCtx, id)
	parameterList.Name = parameterListRequest.Name

	parameterList = parameterListService.ParameterListRepository.Update(repoCtx, parameterList)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) FindById(ctx context.Context, id int64) response.ParameterListResponse {
	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList := parameterListService.findParameterListById(repoCtx, id)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) FindAll(ctx context.Context, parameterListSeach search.ParameterListSearch) []response.ParameterListResponse {
	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterLists := parameterListService.ParameterListRepository.FindAll(repoCtx, parameterListSeach)
	return response.ToParameterListResponses(parameterLists)
}

func (parameterListService *ParameterListServiceImpl) validateRequest(parameterListRequest request.ParameterListRequest) {
	err := parameterListService.validate.Struct(parameterListRequest)
	helper.PanicIfError(err)
}

func (parameterListService *ParameterListServiceImpl) findParameterListById(repoCtx dto.RepositoryContext, id int64) domain.ParameterList {
	parameterList, err := parameterListService.ParameterListRepository.FindById(repoCtx, id)
	exception.PanicErrorBadRequest(err)
	return parameterList
}
