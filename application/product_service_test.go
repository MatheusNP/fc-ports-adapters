package application_test

import (
	"testing"

	"github.com/MatheusNP/fc-ports-adapters/application"
	mock_application "github.com/MatheusNP/fc-ports-adapters/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.NewProductService(persistence)

	got, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, got)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.NewProductService(persistence)

	got, err := service.Create("product 1", 100)
	require.Nil(t, err)
	require.Equal(t, product, got)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.NewProductService(persistence)

	got, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, got)

	got, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, got)
}
