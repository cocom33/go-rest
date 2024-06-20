package app

import (
	"database/sql"
	"go-restapi/helpers"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_database_migration")
	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_rest_api")
	helpers.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	
	return db
	
	// db migration
	// migrate create -ext sql -dir {directory} {file name}
	// migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path {directory} up // n if want execute n table
	// migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path {directory} down // n if want execute n table

	// error / dirty
	// delete incorrect table
	// change version in scheme_migration 	// migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path {directory} force {version} 
	// 
}