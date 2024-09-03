# \PromotionCodesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePromoCode**](PromotionCodesAPI.md#CreatePromoCode) | **Post** /promotion-codes/ | Create Promo Code
[**GetPromoCode**](PromotionCodesAPI.md#GetPromoCode) | **Get** /promotion-codes/{promo_code_id} | Get Promo Code
[**GetPromoCodeByCode**](PromotionCodesAPI.md#GetPromoCodeByCode) | **Get** /promotion-codes/code/{promo_code} | Get Promo Code By Code
[**ListPromoCodes**](PromotionCodesAPI.md#ListPromoCodes) | **Post** /promotion-codes/list | List Promo Codes
[**UpdatePromoCode**](PromotionCodesAPI.md#UpdatePromoCode) | **Put** /promotion-codes/{promo_code_id} | Update Promo Code



## CreatePromoCode

> PromotionCodeExternal CreatePromoCode(ctx).CreatePromoCodeRequest(createPromoCodeRequest).Execute()

Create Promo Code

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
    createPromoCodeRequest := *openapiclient.NewCreatePromoCodeRequest("CouponId_example", "Code_example") // CreatePromoCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionCodesAPI.CreatePromoCode(context.Background()).CreatePromoCodeRequest(createPromoCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionCodesAPI.CreatePromoCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePromoCode`: PromotionCodeExternal
    fmt.Fprintf(os.Stdout, "Response from `PromotionCodesAPI.CreatePromoCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePromoCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPromoCodeRequest** | [**CreatePromoCodeRequest**](CreatePromoCodeRequest.md) |  | 

### Return type

[**PromotionCodeExternal**](PromotionCodeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromoCode

> PromotionCodeExternal GetPromoCode(ctx, promoCodeId).Execute()

Get Promo Code

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
    promoCodeId := "promo_dev_abc123" // string | Unique identifier of the PromotionCode.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionCodesAPI.GetPromoCode(context.Background(), promoCodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionCodesAPI.GetPromoCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPromoCode`: PromotionCodeExternal
    fmt.Fprintf(os.Stdout, "Response from `PromotionCodesAPI.GetPromoCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promoCodeId** | **string** | Unique identifier of the PromotionCode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromoCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromotionCodeExternal**](PromotionCodeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromoCodeByCode

> PromotionCodeExternal GetPromoCodeByCode(ctx, promoCode).Execute()

Get Promo Code By Code

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
    promoCode := "promoCode_example" // string | Unique PromotionCode code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionCodesAPI.GetPromoCodeByCode(context.Background(), promoCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionCodesAPI.GetPromoCodeByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPromoCodeByCode`: PromotionCodeExternal
    fmt.Fprintf(os.Stdout, "Response from `PromotionCodesAPI.GetPromoCodeByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promoCode** | **string** | Unique PromotionCode code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromoCodeByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromotionCodeExternal**](PromotionCodeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPromoCodes

> ListResponsePromotionCodeExternal ListPromoCodes(ctx).PromoCodeQueryParams(promoCodeQueryParams).Execute()

List Promo Codes

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
    promoCodeQueryParams := *openapiclient.NewPromoCodeQueryParams() // PromoCodeQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionCodesAPI.ListPromoCodes(context.Background()).PromoCodeQueryParams(promoCodeQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionCodesAPI.ListPromoCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPromoCodes`: ListResponsePromotionCodeExternal
    fmt.Fprintf(os.Stdout, "Response from `PromotionCodesAPI.ListPromoCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPromoCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promoCodeQueryParams** | [**PromoCodeQueryParams**](PromoCodeQueryParams.md) |  | 

### Return type

[**ListResponsePromotionCodeExternal**](ListResponsePromotionCodeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePromoCode

> PromotionCodeExternal UpdatePromoCode(ctx, promoCodeId).UpdatePromoCodeRequest(updatePromoCodeRequest).Execute()

Update Promo Code

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
    promoCodeId := "promo_dev_abc123" // string | Unique identifier of the PromotionCode.
    updatePromoCodeRequest := *openapiclient.NewUpdatePromoCodeRequest(false) // UpdatePromoCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionCodesAPI.UpdatePromoCode(context.Background(), promoCodeId).UpdatePromoCodeRequest(updatePromoCodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionCodesAPI.UpdatePromoCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePromoCode`: PromotionCodeExternal
    fmt.Fprintf(os.Stdout, "Response from `PromotionCodesAPI.UpdatePromoCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promoCodeId** | **string** | Unique identifier of the PromotionCode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePromoCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePromoCodeRequest** | [**UpdatePromoCodeRequest**](UpdatePromoCodeRequest.md) |  | 

### Return type

[**PromotionCodeExternal**](PromotionCodeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

