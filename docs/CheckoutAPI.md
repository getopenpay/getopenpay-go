# \CheckoutAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckoutSession**](CheckoutAPI.md#CreateCheckoutSession) | **Post** /checkout/sessions | Create Checkout Session
[**GetCheckoutSession**](CheckoutAPI.md#GetCheckoutSession) | **Get** /checkout/sessions/{session_id} | Get Checkout Session
[**ListCheckoutSessions**](CheckoutAPI.md#ListCheckoutSessions) | **Post** /checkout/list | List Checkout Sessions



## CreateCheckoutSession

> CheckoutSessionExternal CreateCheckoutSession(ctx).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()

Create Checkout Session

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
    createCheckoutSessionRequest := *openapiclient.NewCreateCheckoutSessionRequest(openapiclient.CheckoutMode("payment")) // CreateCheckoutSessionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutAPI.CreateCheckoutSession(context.Background()).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CreateCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCheckoutSession`: CheckoutSessionExternal
    fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CreateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckoutSessionRequest** | [**CreateCheckoutSessionRequest**](CreateCheckoutSessionRequest.md) |  | 

### Return type

[**CheckoutSessionExternal**](CheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckoutSession

> CheckoutSessionExternal GetCheckoutSession(ctx, sessionId).Execute()

Get Checkout Session

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
    sessionId := "cs_dev_abc123" // string | Unique identifier of the checkout session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutAPI.GetCheckoutSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.GetCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutSession`: CheckoutSessionExternal
    fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique identifier of the checkout session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSessionExternal**](CheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckoutSessions

> ListResponseCheckoutSessionExternal ListCheckoutSessions(ctx).CheckoutSessionQueryParams(checkoutSessionQueryParams).Execute()

List Checkout Sessions

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
    checkoutSessionQueryParams := *openapiclient.NewCheckoutSessionQueryParams() // CheckoutSessionQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutAPI.ListCheckoutSessions(context.Background()).CheckoutSessionQueryParams(checkoutSessionQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.ListCheckoutSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckoutSessions`: ListResponseCheckoutSessionExternal
    fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.ListCheckoutSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCheckoutSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutSessionQueryParams** | [**CheckoutSessionQueryParams**](CheckoutSessionQueryParams.md) |  | 

### Return type

[**ListResponseCheckoutSessionExternal**](ListResponseCheckoutSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

