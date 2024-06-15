package service

import (
	"context"
	"database/sql"
	"go-restapi/exception"
	"go-restapi/helpers"
	"go-restapi/model/domain"
	"go-restapi/model/web"
	"go-restapi/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		validate: validate,
	}
}

func (service *CategoryServiceImpl) Create (ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommirOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update (ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.validate.Struct(request)
	helpers.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommirOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	category.Name = request.Name
	
	category = service.CategoryRepository.Update(ctx, tx, category)
	
	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete (ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommirOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById (ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommirOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll (ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommirOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helpers.ToCategoryResponse(category))
	}
	
	return categoryResponses
}
