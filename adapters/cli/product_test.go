package cli_test

import (
	"fmt"
	"testing"

	"github.com/MatheusNP/fc-ports-adapters/adapters/cli"
	mock_application "github.com/MatheusNP/fc-ports-adapters/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productID := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)

	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf(
		"Product ID %s with name %s has been created with price %f and status %s",
		productID,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(
		"Product %s has been enabled",
		productName,
	)

	result, err = cli.Run(service, "enable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(
		"Product %s has been disabled",
		productName,
	)

	result, err = cli.Run(service, "disable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productID,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(service, "get", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
