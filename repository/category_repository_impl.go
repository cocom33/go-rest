package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restapi/helpers"
	"go-restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryController() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category (name) values (?)"
	res, err := tx.ExecContext(ctx, sql, category.Name)
	helpers.PanicIfError(err)
	
	id, err := res.LastInsertId()
	helpers.PanicIfError(err)
	
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helpers.PanicIfError(err)
	
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helpers.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helpers.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id, name from category"
	rows, err := tx.QueryContext(ctx, sql)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories = []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		
		categories = append(categories, category)
	}
	
	return categories
}