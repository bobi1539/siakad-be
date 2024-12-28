package repositoryimpl

import (
	"database/sql"
	"errors"
	"fmt"
	"siakad/constant"
	"siakad/helper"
	"siakad/model/domain"
	"siakad/model/dto"
	"siakad/model/search"
	"siakad/repository"
)

type ParameterListRepositoryImpl struct {
}

func NewParameterListRepository() repository.ParameterListRepository {
	return &ParameterListRepositoryImpl{}
}

func (parameterListRepository *ParameterListRepositoryImpl) Save(repoCtx dto.RepositoryContext, parameterList domain.ParameterList) domain.ParameterList {
	result, err := repoCtx.Tx.ExecContext(repoCtx.Ctx, sqlSaveParameterList(), parameterList.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	parameterList.Id = id
	return parameterList
}

func (parameterListRepository *ParameterListRepositoryImpl) Update(repoCtx dto.RepositoryContext, parameterList domain.ParameterList) domain.ParameterList {
	_, err := repoCtx.Tx.ExecContext(repoCtx.Ctx, sqlUpdateParameterList(), parameterList.Name, parameterList.Id)
	helper.PanicIfError(err)

	return parameterList
}

func (parameterListRepository *ParameterListRepositoryImpl) FindById(repoCtx dto.RepositoryContext, id int64) (domain.ParameterList, error) {
	sqlQuery := fmt.Sprintf("%s AND id = ?", sqlSelectParameterList())
	rows := repository.FetchRows(repoCtx, sqlQuery, id)

	parameterList := domain.ParameterList{}

	if rows.Next() {
		scanParameterList(rows, &parameterList)
		return parameterList, nil
	}

	return parameterList, errors.New(constant.DATA_NOT_FOUND)
}

func (parameterListRepository *ParameterListRepositoryImpl) FindAll(repoCtx dto.RepositoryContext, parameterListSearch search.ParameterListSearch) []domain.ParameterList {
	sqlSearch, args := sqlSearchParameterList(parameterListSearch)
	sqlQuery := sqlSelectParameterList() + sqlSearch

	rows := repository.FetchRows(repoCtx, sqlQuery, args...)
	defer rows.Close()

	return getParameterLists(rows)
}

func sqlSaveParameterList() string {
	return "INSERT INTO m_parameter_list(" +
		"name " +
		") " +
		"VALUES (?,?)"
}

func sqlUpdateParameterList() string {
	return "UPDATE m_parameter_list SET " +
		"name = ? " +
		"WHERE id = ?"
}

func sqlSelectParameterList() string {
	return "SELECT " +
		"id, " +
		"name " +
		"FROM m_parameter_list " +
		"WHERE 1 = 1 "
}

func sqlSearchParameterList(parameterListSearch search.ParameterListSearch) (string, []any) {
	var args []any
	sqlQuery := "AND parameter_id = ? "
	args = append(args, parameterListSearch.ParameterId)

	if len(parameterListSearch.Search) != 0 {
		sqlQuery += "AND LOWER(name) LIKE ? "
		searchLike := helper.StringQueryLike(parameterListSearch.Search)
		args = append(args, searchLike, searchLike)
	}

	sqlQuery += "ORDER BY id ASC"
	return sqlQuery, args
}

func scanParameterList(rows *sql.Rows, parameterList *domain.ParameterList) {
	err := rows.Scan(
		&parameterList.Id,
		&parameterList.Name,
	)
	helper.PanicIfError(err)
}

func getParameterLists(rows *sql.Rows) []domain.ParameterList {
	var parameterLists []domain.ParameterList
	for rows.Next() {
		parameterList := domain.ParameterList{}
		scanParameterList(rows, &parameterList)
		parameterLists = append(parameterLists, parameterList)
	}
	return parameterLists
}
