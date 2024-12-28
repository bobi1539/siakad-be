package repository

import (
	"siakad/model/domain"
	"siakad/model/dto"
	"siakad/model/search"
)

type ParameterListRepository interface {
	Save(repoCtx dto.RepositoryContext, parameterList domain.ParameterList) domain.ParameterList
	Update(repoCtx dto.RepositoryContext, parameterList domain.ParameterList) domain.ParameterList
	FindById(repoCtx dto.RepositoryContext, id int64) (domain.ParameterList, error)
	FindAll(repoCtx dto.RepositoryContext, parameterListSearch search.ParameterListSearch) []domain.ParameterList
}
