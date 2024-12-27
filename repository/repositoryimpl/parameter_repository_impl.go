package repositoryimpl

import (
	"database/sql"
	"errors"
	"fmt"
	"siakad/constant"
	"siakad/helper"
	"siakad/model/domain"
	"siakad/model/dto"
	"siakad/repository"
)

type ParameterRepositoryImpl struct {
}

func NewParameterRepository() repository.ParameterRepository {
	return &ParameterRepositoryImpl{}
}

func (parameterRepository *ParameterRepositoryImpl) Save(repoCtx dto.RepositoryContext, parameter domain.Parameter) domain.Parameter {
	result, err := repoCtx.Tx.ExecContext(repoCtx.Ctx, sqlSave(), parameter.Name, parameter.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	parameter.Id = id
	return parameter
}

func (parameterRepository *ParameterRepositoryImpl) Update(repoCtx dto.RepositoryContext, parameter domain.Parameter) domain.Parameter {
	_, err := repoCtx.Tx.ExecContext(repoCtx.Ctx, sqlUpdate(), parameter.Name, parameter.Description, parameter.Id)
	helper.PanicIfError(err)

	return parameter
}

func (parameterRepository *ParameterRepositoryImpl) FindById(repoCtx dto.RepositoryContext, id int64) (domain.Parameter, error) {
	sqlQuery := fmt.Sprintf("%s AND id = ?", sqlSelect())
	rows := repository.FetchRows(repoCtx, sqlQuery, id)
	defer rows.Close()

	parameter := domain.Parameter{}

	if rows.Next() {
		scanParameter(rows, &parameter)
		return parameter, nil
	}

	return parameter, errors.New(constant.DATA_NOT_FOUND)
}

func sqlSave() string {
	return "INSERT INTO m_parameter(" +
		"name, " +
		"description " +
		") " +
		"VALUES (?,?)"
}

func sqlUpdate() string {
	return "UPDATE m_parameter SET " +
		"name = ?, " +
		"description = ? " +
		"WHERE id = ?"
}

func sqlSelect() string {
	return "SELECT " +
		"id, " +
		"name, " +
		"description " +
		"FROM m_parameter " +
		"WHERE 1 = 1 "
}

func scanParameter(rows *sql.Rows, parameter *domain.Parameter) {
	err := rows.Scan(
		&parameter.Id,
		&parameter.Name,
		&parameter.Description,
	)
	helper.PanicIfError(err)
}
