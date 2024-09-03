# \ProductsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductsAPI.md#CreateProduct) | **Post** /products/ | Create Product
[**DeleteProduct**](ProductsAPI.md#DeleteProduct) | **Delete** /products/{product_id} | Delete Product
[**GetProduct**](ProductsAPI.md#GetProduct) | **Get** /products/{product_id} | Get Product
[**ListProducts**](ProductsAPI.md#ListProducts) | **Post** /products/list | List Products
[**SearchProducts**](ProductsAPI.md#SearchProducts) | **Post** /products/search | Search Products
[**UpdateProduct**](ProductsAPI.md#UpdateProduct) | **Put** /products/{product_id} | Update Product



## CreateProduct

> ProductExternal CreateProduct(ctx).CreateProductRequest(createProductRequest).Execute()

Create Product

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
    createProductRequest := *openapiclient.NewCreateProductRequest("product1", "This is description for product1") // CreateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.CreateProduct(context.Background()).CreateProductRequest(createProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: ProductExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) |  | 

### Return type

[**ProductExternal**](ProductExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProductResponse DeleteProduct(ctx, productId).Execute()

Delete Product

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
    productId := "product_dev_abc123" // string | Unique identifier of the product.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.DeleteProduct(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.DeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProduct`: DeleteProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Unique identifier of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteProductResponse**](DeleteProductResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> ProductExternal GetProduct(ctx, productId).Execute()

Get Product

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
    productId := "product_dev_abc123" // string | Unique identifier of the product.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.GetProduct(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: ProductExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Unique identifier of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductExternal**](ProductExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ListResponseProductExternal ListProducts(ctx).ProductQueryParams(productQueryParams).Execute()

List Products

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
    productQueryParams := *openapiclient.NewProductQueryParams() // ProductQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.ListProducts(context.Background()).ProductQueryParams(productQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ListResponseProductExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productQueryParams** | [**ProductQueryParams**](ProductQueryParams.md) |  | 

### Return type

[**ListResponseProductExternal**](ListResponseProductExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProducts

> ListResponseProductExternal SearchProducts(ctx).SearchProductRequest(searchProductRequest).Execute()

Search Products

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
    searchProductRequest := *openapiclient.NewSearchProductRequest("Query_example") // SearchProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.SearchProducts(context.Background()).SearchProductRequest(searchProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.SearchProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProducts`: ListResponseProductExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.SearchProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchProductRequest** | [**SearchProductRequest**](SearchProductRequest.md) |  | 

### Return type

[**ListResponseProductExternal**](ListResponseProductExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> ProductExternal UpdateProduct(ctx, productId).UpdateProductRequest(updateProductRequest).Execute()

Update Product

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
    productId := "product_dev_abc123" // string | Unique identifier of the product.
    updateProductRequest := *openapiclient.NewUpdateProductRequest() // UpdateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsAPI.UpdateProduct(context.Background(), productId).UpdateProductRequest(updateProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.UpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProduct`: ProductExternal
    fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Unique identifier of the product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProductRequest** | [**UpdateProductRequest**](UpdateProductRequest.md) |  | 

### Return type

[**ProductExternal**](ProductExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

