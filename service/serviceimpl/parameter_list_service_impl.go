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
	ParameterService        service.ParameterService
	DB                      *sql.DB
	validate                *validator.Validate
}

func NewParameterListService(
	parameterListRepository repository.ParameterListRepository,
	parameterService service.ParameterService,
	db *sql.DB,
	validate *validator.Validate,
) service.ParameterListService {
	return &ParameterListServiceImpl{
		ParameterListRepository: parameterListRepository,
		ParameterService:        parameterService,
		DB:                      db,
		validate:                validate,
	}
}

func (parameterListService *ParameterListServiceImpl) Create(ctx context.Context, parameterListRequest request.ParameterListRequest) response.ParameterListResponse {
	parameterListService.validateRequest(parameterListRequest)

	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	parameterList := domain.ParameterList{}
	parameterListService.setParameterList(ctx, &parameterList, &parameterListRequest)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList = parameterListService.ParameterListRepository.Save(repoCtx, parameterList)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) Update(
	ctx context.Context,
	id int64,
	parameterListRequest request.ParameterListRequest,
) response.ParameterListResponse {
	parameterListService.validateRequest(parameterListRequest)

	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	parameterList := parameterListService.FindByIdDomain(ctx, id)
	parameterListService.setParameterList(ctx, &parameterList, &parameterListRequest)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList = parameterListService.ParameterListRepository.Update(repoCtx, parameterList)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) FindById(ctx context.Context, id int64) response.ParameterListResponse {
	parameterList := parameterListService.FindByIdDomain(ctx, id)
	return response.ToParameterListResponse(parameterList)
}

func (parameterListService *ParameterListServiceImpl) FindAll(ctx context.Context, parameterListSeach search.ParameterListSearch) []response.ParameterListResponse {
	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterLists := parameterListService.ParameterListRepository.FindAll(repoCtx, parameterListSeach)
	return response.ToParameterListResponses(parameterLists)
}

func (parameterListService *ParameterListServiceImpl) FindByIdDomain(ctx context.Context, id int64) domain.ParameterList {
	tx := service.BeginTransaction(parameterListService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameterList, err := parameterListService.ParameterListRepository.FindById(repoCtx, id)
	exception.PanicErrorBadRequest(err)
	return parameterList
}

func (parameterListService *ParameterListServiceImpl) validateRequest(parameterListRequest request.ParameterListRequest) {
	err := parameterListService.validate.Struct(parameterListRequest)
	helper.PanicIfError(err)
}

func (parameterListService *ParameterListServiceImpl) setParameterList(
	ctx context.Context,
	parameterList *domain.ParameterList,
	parameterListRequest *request.ParameterListRequest,
) {
	parameterList.Name = parameterListRequest.Name
	parameterList.Parameter = parameterListService.ParameterService.FindByIdDomain(ctx, parameterListRequest.ParameterId)
}
