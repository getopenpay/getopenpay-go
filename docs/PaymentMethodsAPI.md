# \PaymentMethodsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachPaymentMethod**](PaymentMethodsAPI.md#AttachPaymentMethod) | **Post** /payment-methods/{payment_method_id}/attach | Attach Payment Method
[**AuthorizePaymentMethod**](PaymentMethodsAPI.md#AuthorizePaymentMethod) | **Post** /payment-methods/authorize | Authorize Payment Method
[**DetachPaymentMethod**](PaymentMethodsAPI.md#DetachPaymentMethod) | **Post** /payment-methods/{payment_method_id}/detach | Detach Payment Method
[**GetPaymentMethod**](PaymentMethodsAPI.md#GetPaymentMethod) | **Get** /payment-methods/{payment_method_id} | Get Payment Method



## AttachPaymentMethod

> PaymentMethodExternal AttachPaymentMethod(ctx, paymentMethodId).AttachPaymentMethodRequest(attachPaymentMethodRequest).Execute()

Attach Payment Method



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
    paymentMethodId := "pm_dev_abc123" // string | Unique Identifier of the payment_method.
    attachPaymentMethodRequest := *openapiclient.NewAttachPaymentMethodRequest("cus_dev_abc123") // AttachPaymentMethodRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsAPI.AttachPaymentMethod(context.Background(), paymentMethodId).AttachPaymentMethodRequest(attachPaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.AttachPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachPaymentMethod`: PaymentMethodExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.AttachPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | Unique Identifier of the payment_method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachPaymentMethodRequest** | [**AttachPaymentMethodRequest**](AttachPaymentMethodRequest.md) |  | 

### Return type

[**PaymentMethodExternal**](PaymentMethodExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePaymentMethod

> PaymentMethodExternal AuthorizePaymentMethod(ctx).AuthorizePaymentMethodRequest(authorizePaymentMethodRequest).Execute()

Authorize Payment Method

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
    authorizePaymentMethodRequest := *openapiclient.NewAuthorizePaymentMethodRequest("pm_dev_abc123") // AuthorizePaymentMethodRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsAPI.AuthorizePaymentMethod(context.Background()).AuthorizePaymentMethodRequest(authorizePaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.AuthorizePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizePaymentMethod`: PaymentMethodExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.AuthorizePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizePaymentMethodRequest** | [**AuthorizePaymentMethodRequest**](AuthorizePaymentMethodRequest.md) |  | 

### Return type

[**PaymentMethodExternal**](PaymentMethodExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachPaymentMethod

> PaymentMethodExternal DetachPaymentMethod(ctx, paymentMethodId).Execute()

Detach Payment Method



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
    paymentMethodId := "pm_dev_abc123" // string | Unique Identifier of the payment_method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsAPI.DetachPaymentMethod(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.DetachPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachPaymentMethod`: PaymentMethodExternal
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.DetachPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | Unique Identifier of the payment_method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentMethodExternal**](PaymentMethodExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet GetPaymentMethod(ctx, paymentMethodId).Execute()

Get Payment Method

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
    paymentMethodId := "pm_dev_abc123" // string | Unique Identifier of the payment_method.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentMethodsAPI.GetPaymentMethod(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.GetPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethod`: ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.GetPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | Unique Identifier of the payment_method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet**](ResponseGetPaymentMethodPaymentMethodsPaymentMethodIdGet.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

