/*
OpenPay API

Testing CustomersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package getopenpay

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_getopenpay_CustomersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CustomersAPIService CreateCustomer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService CreateCustomerBalanceTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.CreateCustomerBalanceTransaction(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService DeleteCustomerDiscount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.DeleteCustomerDiscount(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomer(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomerBalanceTransaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var transactionId string

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomerBalanceTransaction(context.Background(), customerId, transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService GetCustomerBalanceTransactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.GetCustomerBalanceTransactions(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService ListCustomerPaymentMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.ListCustomerPaymentMethods(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService ListCustomers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.ListCustomers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService ListValidSubscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string

		resp, httpRes, err := apiClient.CustomersAPI.ListValidSubscriptions(context.Background(), customerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService SearchCustomers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CustomersAPI.SearchCustomers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CustomersAPIService UpdateCustomer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerExternalId string

		resp, httpRes, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), customerExternalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}