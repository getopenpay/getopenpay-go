# \PaymentRoutesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentRoute**](PaymentRoutesAPI.md#CreatePaymentRoute) | **Post** /payment-routes/ | Create Payment Route
[**DeletePaymentRoute**](PaymentRoutesAPI.md#DeletePaymentRoute) | **Delete** /payment-routes/{payment_route_id} | Delete Payment Route
[**GetPaymentRoute**](PaymentRoutesAPI.md#GetPaymentRoute) | **Get** /payment-routes/{payment_route_id} | Get Payment Route
[**ListPaymentRoutes**](PaymentRoutesAPI.md#ListPaymentRoutes) | **Post** /payment-routes/list | List Payment Routes
[**UpdatePaymentRoute**](PaymentRoutesAPI.md#UpdatePaymentRoute) | **Put** /payment-routes/{payment_route_id} | Update Payment Route



## CreatePaymentRoute

> PaymentRouteExternal CreatePaymentRoute(ctx).CreatePaymentRouteRequest(createPaymentRouteRequest).Execute()

Create Payment Route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	createPaymentRouteRequest := *openapiclient.NewCreatePaymentRouteRequest("My Payment Route", map[string]interface{}(123)) // CreatePaymentRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRoutesAPI.CreatePaymentRoute(context.Background()).CreatePaymentRouteRequest(createPaymentRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRoutesAPI.CreatePaymentRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentRoute`: PaymentRouteExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentRoutesAPI.CreatePaymentRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentRouteRequest** | [**CreatePaymentRouteRequest**](CreatePaymentRouteRequest.md) |  | 

### Return type

[**PaymentRouteExternal**](PaymentRouteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentRoute

> DeletePaymentRouteResponse DeletePaymentRoute(ctx, paymentRouteId).Execute()

Delete Payment Route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	paymentRouteId := "paymentRouteId_example" // string | The payment route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRoutesAPI.DeletePaymentRoute(context.Background(), paymentRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRoutesAPI.DeletePaymentRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePaymentRoute`: DeletePaymentRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentRoutesAPI.DeletePaymentRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRouteId** | **string** | The payment route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePaymentRouteResponse**](DeletePaymentRouteResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRoute

> PaymentRouteExternal GetPaymentRoute(ctx, paymentRouteId).Execute()

Get Payment Route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	paymentRouteId := "paymentRouteId_example" // string | The payment route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRoutesAPI.GetPaymentRoute(context.Background(), paymentRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRoutesAPI.GetPaymentRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentRoute`: PaymentRouteExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentRoutesAPI.GetPaymentRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRouteId** | **string** | The payment route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentRouteExternal**](PaymentRouteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentRoutes

> ListResponsePaymentRouteExternal ListPaymentRoutes(ctx).Execute()

List Payment Routes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRoutesAPI.ListPaymentRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRoutesAPI.ListPaymentRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentRoutes`: ListResponsePaymentRouteExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentRoutesAPI.ListPaymentRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentRoutesRequest struct via the builder pattern


### Return type

[**ListResponsePaymentRouteExternal**](ListResponsePaymentRouteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentRoute

> PaymentRouteExternal UpdatePaymentRoute(ctx, paymentRouteId).UpdatePaymentRouteRequest(updatePaymentRouteRequest).Execute()

Update Payment Route

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/getopenpay/getopenpay-go"
)

func main() {
	paymentRouteId := "paymentRouteId_example" // string | The payment route ID
	updatePaymentRouteRequest := *openapiclient.NewUpdatePaymentRouteRequest("My Payment Route", map[string]interface{}(123)) // UpdatePaymentRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRoutesAPI.UpdatePaymentRoute(context.Background(), paymentRouteId).UpdatePaymentRouteRequest(updatePaymentRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRoutesAPI.UpdatePaymentRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePaymentRoute`: PaymentRouteExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentRoutesAPI.UpdatePaymentRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRouteId** | **string** | The payment route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaymentRouteRequest** | [**UpdatePaymentRouteRequest**](UpdatePaymentRouteRequest.md) |  | 

### Return type

[**PaymentRouteExternal**](PaymentRouteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

