# \ChargesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharge**](ChargesAPI.md#GetCharge) | **Get** /charges/{charge_id} | Get Charge
[**ListCharges**](ChargesAPI.md#ListCharges) | **Post** /charges/list | List Charges
[**SearchCharges**](ChargesAPI.md#SearchCharges) | **Post** /charges/search | Search Charges
[**UpdateCharge**](ChargesAPI.md#UpdateCharge) | **Put** /charges/{charge_id} | Update Charge



## GetCharge

> ChargeExternal GetCharge(ctx, chargeId).Execute()

Get Charge

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
    chargeId := "chargeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesAPI.GetCharge(context.Background(), chargeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.GetCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCharge`: ChargeExternal
    fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.GetCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChargeExternal**](ChargeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCharges

> ListResponseChargeExternal ListCharges(ctx).ChargeQueryParams(chargeQueryParams).Execute()

List Charges

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
    chargeQueryParams := *openapiclient.NewChargeQueryParams() // ChargeQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesAPI.ListCharges(context.Background()).ChargeQueryParams(chargeQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.ListCharges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCharges`: ListResponseChargeExternal
    fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.ListCharges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeQueryParams** | [**ChargeQueryParams**](ChargeQueryParams.md) |  | 

### Return type

[**ListResponseChargeExternal**](ListResponseChargeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCharges

> ListResponseChargeExternal SearchCharges(ctx).SearchChargeRequest(searchChargeRequest).Execute()

Search Charges

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
    searchChargeRequest := *openapiclient.NewSearchChargeRequest("Query_example") // SearchChargeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesAPI.SearchCharges(context.Background()).SearchChargeRequest(searchChargeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.SearchCharges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCharges`: ListResponseChargeExternal
    fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.SearchCharges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchChargeRequest** | [**SearchChargeRequest**](SearchChargeRequest.md) |  | 

### Return type

[**ListResponseChargeExternal**](ListResponseChargeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCharge

> ChargeExternal UpdateCharge(ctx, chargeId).UpdateChargeRequest(updateChargeRequest).Execute()

Update Charge

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
    chargeId := "chargeId_example" // string | 
    updateChargeRequest := *openapiclient.NewUpdateChargeRequest() // UpdateChargeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesAPI.UpdateCharge(context.Background(), chargeId).UpdateChargeRequest(updateChargeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.UpdateCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCharge`: ChargeExternal
    fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.UpdateCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChargeRequest** | [**UpdateChargeRequest**](UpdateChargeRequest.md) |  | 

### Return type

[**ChargeExternal**](ChargeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

