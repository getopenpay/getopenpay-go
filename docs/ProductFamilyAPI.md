# \ProductFamilyAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductFamily**](ProductFamilyAPI.md#CreateProductFamily) | **Post** /product-family/ | Create Product Family
[**DeleteProductFamily**](ProductFamilyAPI.md#DeleteProductFamily) | **Delete** /product-family/{id} | Delete Product Family
[**GetProductFamily**](ProductFamilyAPI.md#GetProductFamily) | **Get** /product-family/{id} | Get Product Family
[**ListProductFamilies**](ProductFamilyAPI.md#ListProductFamilies) | **Post** /product-family/list | List Product Families
[**UpdateProductFamily**](ProductFamilyAPI.md#UpdateProductFamily) | **Put** /product-family/{id} | Update Product Family



## CreateProductFamily

> ProductFamilyExternal CreateProductFamily(ctx).CreateProductFamilyRequest(createProductFamilyRequest).Execute()

Create Product Family

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
    createProductFamilyRequest := *openapiclient.NewCreateProductFamilyRequest("CRM products", []string{"Products_example"}) // CreateProductFamilyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductFamilyAPI.CreateProductFamily(context.Background()).CreateProductFamilyRequest(createProductFamilyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFamilyAPI.CreateProductFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductFamily`: ProductFamilyExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductFamilyAPI.CreateProductFamily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductFamilyRequest** | [**CreateProductFamilyRequest**](CreateProductFamilyRequest.md) |  | 

### Return type

[**ProductFamilyExternal**](ProductFamilyExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProductFamily

> DeleteProductFamily(ctx, id).Execute()

Delete Product Family

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProductFamilyAPI.DeleteProductFamily(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFamilyAPI.DeleteProductFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductFamily

> ProductFamilyExternal GetProductFamily(ctx, id).Execute()

Get Product Family

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductFamilyAPI.GetProductFamily(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFamilyAPI.GetProductFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductFamily`: ProductFamilyExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductFamilyAPI.GetProductFamily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductFamilyExternal**](ProductFamilyExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductFamilies

> ListResponseProductFamilyExternal ListProductFamilies(ctx).ProductFamilyQueryParams(productFamilyQueryParams).Execute()

List Product Families

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
    productFamilyQueryParams := *openapiclient.NewProductFamilyQueryParams() // ProductFamilyQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductFamilyAPI.ListProductFamilies(context.Background()).ProductFamilyQueryParams(productFamilyQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFamilyAPI.ListProductFamilies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductFamilies`: ListResponseProductFamilyExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductFamilyAPI.ListProductFamilies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductFamiliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productFamilyQueryParams** | [**ProductFamilyQueryParams**](ProductFamilyQueryParams.md) |  | 

### Return type

[**ListResponseProductFamilyExternal**](ListResponseProductFamilyExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductFamily

> ProductFamilyExternal UpdateProductFamily(ctx, id).UpdateProductFamilyRequest(updateProductFamilyRequest).Execute()

Update Product Family

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
    id := "id_example" // string | 
    updateProductFamilyRequest := *openapiclient.NewUpdateProductFamilyRequest("Name_example", []string{"Products_example"}) // UpdateProductFamilyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductFamilyAPI.UpdateProductFamily(context.Background(), id).UpdateProductFamilyRequest(updateProductFamilyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFamilyAPI.UpdateProductFamily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProductFamily`: ProductFamilyExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductFamilyAPI.UpdateProductFamily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProductFamilyRequest** | [**UpdateProductFamilyRequest**](UpdateProductFamilyRequest.md) |  | 

### Return type

[**ProductFamilyExternal**](ProductFamilyExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

