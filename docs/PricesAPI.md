# \PricesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriceForProduct**](PricesAPI.md#CreatePriceForProduct) | **Post** /prices/ | Create Price For Product
[**DeletePriceForProduct**](PricesAPI.md#DeletePriceForProduct) | **Delete** /prices/{price_id} | Delete Price For Product
[**GetPriceAddOns**](PricesAPI.md#GetPriceAddOns) | **Get** /prices/{price_id}/add-ons | Get Price Add Ons
[**GetPriceForProduct**](PricesAPI.md#GetPriceForProduct) | **Get** /prices/{price_id} | Get Price For Product
[**ListPrices**](PricesAPI.md#ListPrices) | **Post** /prices/list | List Prices
[**SearchPrices**](PricesAPI.md#SearchPrices) | **Post** /prices/search | Search Prices
[**UpdatePriceForProduct**](PricesAPI.md#UpdatePriceForProduct) | **Put** /prices/{price_id} | Update Price For Product



## CreatePriceForProduct

> PriceExternal CreatePriceForProduct(ctx).CreatePriceRequest(createPriceRequest).Execute()

Create Price For Product

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
    createPriceRequest := *openapiclient.NewCreatePriceRequest(true, openapiclient.PricingModel("standard"), NullableInt32(123), openapiclient.PriceTypeEnum("one_time"), "product_dev_abc123") // CreatePriceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.CreatePriceForProduct(context.Background()).CreatePriceRequest(createPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.CreatePriceForProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePriceForProduct`: PriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.CreatePriceForProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriceForProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPriceRequest** | [**CreatePriceRequest**](CreatePriceRequest.md) |  | 

### Return type

[**PriceExternal**](PriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriceForProduct

> DeletePriceResponse DeletePriceForProduct(ctx, priceId).Execute()

Delete Price For Product

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
    priceId := "price_dev_abc123" // string | Unique identifier of the price.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.DeletePriceForProduct(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.DeletePriceForProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePriceForProduct`: DeletePriceResponse
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.DeletePriceForProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | **string** | Unique identifier of the price. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceForProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePriceResponse**](DeletePriceResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceAddOns

> []PriceExternal GetPriceAddOns(ctx, priceId).Execute()

Get Price Add Ons



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
    priceId := "price_dev_abc123" // string | Unique identifier of the price.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.GetPriceAddOns(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetPriceAddOns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceAddOns`: []PriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetPriceAddOns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | **string** | Unique identifier of the price. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceAddOnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PriceExternal**](PriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceForProduct

> PriceExternal GetPriceForProduct(ctx, priceId).Expand(expand).Execute()

Get Price For Product

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
    priceId := "price_dev_abc123" // string | Unique identifier of the price.
    expand := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.GetPriceForProduct(context.Background(), priceId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetPriceForProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceForProduct`: PriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetPriceForProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | **string** | Unique identifier of the price. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceForProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** |  | 

### Return type

[**PriceExternal**](PriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrices

> ListResponsePriceExternal ListPrices(ctx).PriceQueryParams(priceQueryParams).Execute()

List Prices

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
    priceQueryParams := *openapiclient.NewPriceQueryParams() // PriceQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.ListPrices(context.Background()).PriceQueryParams(priceQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.ListPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrices`: ListResponsePriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.ListPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceQueryParams** | [**PriceQueryParams**](PriceQueryParams.md) |  | 

### Return type

[**ListResponsePriceExternal**](ListResponsePriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPrices

> ListResponsePriceExternal SearchPrices(ctx).SearchPriceRequest(searchPriceRequest).Execute()

Search Prices

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
    searchPriceRequest := *openapiclient.NewSearchPriceRequest("Query_example") // SearchPriceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.SearchPrices(context.Background()).SearchPriceRequest(searchPriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.SearchPrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPrices`: ListResponsePriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.SearchPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPriceRequest** | [**SearchPriceRequest**](SearchPriceRequest.md) |  | 

### Return type

[**ListResponsePriceExternal**](ListResponsePriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriceForProduct

> PriceExternal UpdatePriceForProduct(ctx, priceId).UpdatePriceRequest(updatePriceRequest).Execute()

Update Price For Product

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
    priceId := "price_dev_abc123" // string | Unique identifier of the price.
    updatePriceRequest := *openapiclient.NewUpdatePriceRequest() // UpdatePriceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PricesAPI.UpdatePriceForProduct(context.Background(), priceId).UpdatePriceRequest(updatePriceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.UpdatePriceForProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePriceForProduct`: PriceExternal
    fmt.Fprintf(os.Stdout, "Response from `PricesAPI.UpdatePriceForProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | **string** | Unique identifier of the price. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceForProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePriceRequest** | [**UpdatePriceRequest**](UpdatePriceRequest.md) |  | 

### Return type

[**PriceExternal**](PriceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

