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

type ParameterServiceImpl struct {
	ParameterRepository repository.ParameterRepository
	DB                  *sql.DB
	validate            *validator.Validate
}

func NewParameterService(parameterRepository repository.ParameterRepository, db *sql.DB, validate *validator.Validate) service.ParameterService {
	return &ParameterServiceImpl{
		ParameterRepository: parameterRepository,
		DB:                  db,
		validate:            validate,
	}
}

func (parameterService *ParameterServiceImpl) Create(ctx context.Context, parameterRequest request.ParameterRequest) response.ParameterResponse {
	parameterService.validateRequest(parameterRequest)

	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	parameter := domain.Parameter{}
	parameterService.setParameter(&parameter, &parameterRequest)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter = parameterService.ParameterRepository.Save(repoCtx, parameter)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) Update(ctx context.Context, id int64, parameterRequest request.ParameterRequest) response.ParameterResponse {
	parameterService.validateRequest(parameterRequest)

	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	parameter := parameterService.FindByIdDomain(ctx, id)
	parameterService.setParameter(&parameter, &parameterRequest)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter = parameterService.ParameterRepository.Update(repoCtx, parameter)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) FindById(ctx context.Context, id int64) response.ParameterResponse {
	parameter := parameterService.FindByIdDomain(ctx, id)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) FindAll(ctx context.Context, generalSearch search.GeneralSearch) []response.ParameterResponse {
	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameters := parameterService.ParameterRepository.FindAll(repoCtx, generalSearch)
	return response.ToParameterResponses(parameters)
}

func (parameterService *ParameterServiceImpl) FindByIdDomain(ctx context.Context, id int64) domain.Parameter {
	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter, err := parameterService.ParameterRepository.FindById(repoCtx, id)
	exception.PanicErrorBadRequest(err)
	return parameter
}

func (parameterService *ParameterServiceImpl) validateRequest(parameterRequest request.ParameterRequest) {
	err := parameterService.validate.Struct(parameterRequest)
	helper.PanicIfError(err)
}

func (parameterService *ParameterServiceImpl) setParameter(parameter *domain.Parameter, parameterRequest *request.ParameterRequest) {
	parameter.Name = parameterRequest.Name
	parameter.Description = parameterRequest.Description
}
