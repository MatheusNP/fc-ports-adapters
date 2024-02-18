package main

import (
	"database/sql"

	"github.com/MatheusNP/fc-ports-adapters/adapters/db"
	"github.com/MatheusNP/fc-ports-adapters/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlDB, _ := sql.Open("sqlite3", "sqlite.db")

	productDBAdapter := db.NewProductDB(sqlDB)

	productService := application.NewProductService(productDBAdapter)

	productService.Create("Product Example", 30)
}
