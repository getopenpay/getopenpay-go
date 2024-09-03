/*
OpenPay API

Testing ChargesAPIService

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

func Test_getopenpay_ChargesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChargesAPIService GetCharge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var chargeId string

		resp, httpRes, err := apiClient.ChargesAPI.GetCharge(context.Background(), chargeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargesAPIService ListCharges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChargesAPI.ListCharges(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargesAPIService SearchCharges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChargesAPI.SearchCharges(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargesAPIService UpdateCharge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var chargeId string

		resp, httpRes, err := apiClient.ChargesAPI.UpdateCharge(context.Background(), chargeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
