package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/MatheusNP/fc-ports-adapters/adapters/db"
	"github.com/MatheusNP/fc-ports-adapters/application"
	"github.com/stretchr/testify/require"
)

var sqlDB *sql.DB

func setUp() {
	sqlDB, _ = sql.Open("sqlite3", ":memory:")
	createTable(sqlDB)
	createProduct(sqlDB)
}

func createTable(sqlDB *sql.DB) {
	table := `create table products(
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := sqlDB.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(sqlDB *sql.DB) {
	insert := `insert into products values(
		"abc",
		"Product Test",
		0,
		"disabled"
	);`

	stmt, err := sqlDB.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer sqlDB.Close()

	productDB := db.NewProductDB(sqlDB)

	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer sqlDB.Close()

	productDB := db.NewProductDB(sqlDB)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 10.0

	got, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, got.GetName())
	require.Equal(t, product.Price, got.GetPrice())
	require.Equal(t, product.Status, got.GetStatus())

	product.Enable()
	got, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, got.GetName())
	require.Equal(t, product.Price, got.GetPrice())
	require.Equal(t, "enabled", got.GetStatus())
}
