# \PaymentIntentsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentIntent**](PaymentIntentsAPI.md#GetPaymentIntent) | **Get** /payment-intents/{payment_intent_id} | Get Payment Intent
[**ListPaymentIntents**](PaymentIntentsAPI.md#ListPaymentIntents) | **Post** /payment-intents/list | List Payment Intents
[**SearchPaymentIntents**](PaymentIntentsAPI.md#SearchPaymentIntents) | **Post** /payment-intents/search | Search Payment Intents
[**UpdatePaymentIntent**](PaymentIntentsAPI.md#UpdatePaymentIntent) | **Put** /payment-intents/{payment_intent_id} | Update Payment Intent



## GetPaymentIntent

> PaymentIntentExternal GetPaymentIntent(ctx, paymentIntentId).Expand(expand).Execute()

Get Payment Intent

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
    paymentIntentId := "paymentIntentId_example" // string | 
    expand := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentIntentsAPI.GetPaymentIntent(context.Background(), paymentIntentId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.GetPaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentIntent`: PaymentIntentExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.GetPaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentIntentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** |  | 

### Return type

[**PaymentIntentExternal**](PaymentIntentExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentIntents

> ListResponsePaymentIntentExternal ListPaymentIntents(ctx).PaymentIntentQueryParams(paymentIntentQueryParams).Execute()

List Payment Intents

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
    paymentIntentQueryParams := *openapiclient.NewPaymentIntentQueryParams() // PaymentIntentQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentIntentsAPI.ListPaymentIntents(context.Background()).PaymentIntentQueryParams(paymentIntentQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.ListPaymentIntents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentIntents`: ListResponsePaymentIntentExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.ListPaymentIntents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentIntentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentIntentQueryParams** | [**PaymentIntentQueryParams**](PaymentIntentQueryParams.md) |  | 

### Return type

[**ListResponsePaymentIntentExternal**](ListResponsePaymentIntentExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPaymentIntents

> ListResponsePaymentIntentExternal SearchPaymentIntents(ctx).SearchPaymentIntentRequest(searchPaymentIntentRequest).Execute()

Search Payment Intents

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
    searchPaymentIntentRequest := *openapiclient.NewSearchPaymentIntentRequest("Query_example") // SearchPaymentIntentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentIntentsAPI.SearchPaymentIntents(context.Background()).SearchPaymentIntentRequest(searchPaymentIntentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.SearchPaymentIntents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPaymentIntents`: ListResponsePaymentIntentExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.SearchPaymentIntents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPaymentIntentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPaymentIntentRequest** | [**SearchPaymentIntentRequest**](SearchPaymentIntentRequest.md) |  | 

### Return type

[**ListResponsePaymentIntentExternal**](ListResponsePaymentIntentExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentIntent

> PaymentIntentExternal UpdatePaymentIntent(ctx, paymentIntentId).UpdatePaymentIntent(updatePaymentIntent).Execute()

Update Payment Intent

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
    paymentIntentId := "paymentIntentId_example" // string | 
    updatePaymentIntent := *openapiclient.NewUpdatePaymentIntent() // UpdatePaymentIntent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentIntentsAPI.UpdatePaymentIntent(context.Background(), paymentIntentId).UpdatePaymentIntent(updatePaymentIntent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.UpdatePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentIntent`: PaymentIntentExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.UpdatePaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentIntentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaymentIntent** | [**UpdatePaymentIntent**](UpdatePaymentIntent.md) |  | 

### Return type

[**PaymentIntentExternal**](PaymentIntentExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

