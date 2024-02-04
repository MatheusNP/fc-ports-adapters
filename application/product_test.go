package application_test

import (
	"testing"

	"github.com/MatheusNP/fc-ports-adapters/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewString(),
		Name:   "hello",
		Status: application.DISABLED,
		Price:  10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "invalid status for product", err.Error())

	product.Status = application.ENABLED
	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	newUUID := uuid.NewString()
	product := application.Product{
		ID: newUUID,
	}

	got := product.GetID()
	require.Equal(t, newUUID, got)
}

func TestProduct_GetName(t *testing.T) {
	newName := "hello"
	product := application.Product{
		Name: newName,
	}

	got := product.GetName()
	require.Equal(t, newName, got)
}

func TestProduct_GetStatus(t *testing.T) {
	newStatus := application.DISABLED
	product := application.Product{
		Status: newStatus,
	}

	got := product.GetStatus()
	require.Equal(t, newStatus, got)
}

func TestProduct_GetPrice(t *testing.T) {
	newPrice := float64(10)
	product := application.Product{
		Price: newPrice,
	}

	got := product.GetPrice()
	require.Equal(t, newPrice, got)
}
