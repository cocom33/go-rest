package main

import (
	"go-restapi/app"
	"go-restapi/controller"
	"go-restapi/helpers"
	"go-restapi/middleware"
	"go-restapi/repository"
	"go-restapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryController()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryConroller := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryConroller)
	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}