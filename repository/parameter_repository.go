package repository

import (
	"siakad/model/domain"
	"siakad/model/dto"
)

type ParameterRepository interface {
	Save(repoCtx dto.RepositoryContext, parameter domain.Parameter) domain.Parameter
	Update(repoCtx dto.RepositoryContext, parameter domain.Parameter) domain.Parameter
	FindById(repoCtx dto.RepositoryContext, id int64) (domain.Parameter, error)
}
