package serviceimpl

import (
	"context"
	"database/sql"
	"siakad/exception"
	"siakad/helper"
	"siakad/model/domain"
	"siakad/model/dto"
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

	parameter := domain.Parameter{
		Name:        parameterRequest.Name,
		Description: parameterRequest.Description,
	}

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter = parameterService.ParameterRepository.Save(repoCtx, parameter)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) Update(ctx context.Context, id int64, parameterRequest request.ParameterRequest) response.ParameterResponse {
	parameterService.validateRequest(parameterRequest)

	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter := parameterService.findParameterById(repoCtx, id)
	parameter.Name = parameterRequest.Name
	parameter.Description = parameterRequest.Description

	parameter = parameterService.ParameterRepository.Update(repoCtx, parameter)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) FindById(ctx context.Context, id int64) response.ParameterResponse {
	tx := service.BeginTransaction(parameterService.DB)
	defer helper.CommitOrRollback(tx)

	repoCtx := dto.BuildRepoCtx(ctx, tx)
	parameter := parameterService.findParameterById(repoCtx, id)
	return response.ToParameterResponse(parameter)
}

func (parameterService *ParameterServiceImpl) validateRequest(parameterRequest request.ParameterRequest) {
	err := parameterService.validate.Struct(parameterRequest)
	helper.PanicIfError(err)
}

func (parameterService *ParameterServiceImpl) findParameterById(repoCtx dto.RepositoryContext, id int64) domain.Parameter {
	parameter, err := parameterService.ParameterRepository.FindById(repoCtx, id)
	exception.PanicErrorBadRequest(err)
	return parameter
}