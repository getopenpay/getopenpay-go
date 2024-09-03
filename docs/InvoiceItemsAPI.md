# \InvoiceItemsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInvoiceItemById**](InvoiceItemsAPI.md#DeleteInvoiceItemById) | **Delete** /invoice-items/{invoice_item_id} | Delete Invoice Item By Id
[**GetInvoiceItem**](InvoiceItemsAPI.md#GetInvoiceItem) | **Get** /invoice-items/{invoice_item_id} | Get Invoice Item
[**ListInvoiceItems**](InvoiceItemsAPI.md#ListInvoiceItems) | **Post** /invoice-items/list | List Invoice Items



## DeleteInvoiceItemById

> DeleteInvoiceItemResponse DeleteInvoiceItemById(ctx, invoiceItemId).Execute()

Delete Invoice Item By Id

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
    invoiceItemId := "invoiceitem_dev_abc123" // string | Unique identifier of the invoice_item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceItemsAPI.DeleteInvoiceItemById(context.Background(), invoiceItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceItemsAPI.DeleteInvoiceItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInvoiceItemById`: DeleteInvoiceItemResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoiceItemsAPI.DeleteInvoiceItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceItemId** | **string** | Unique identifier of the invoice_item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInvoiceItemResponse**](DeleteInvoiceItemResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceItem

> InvoiceItemExternal GetInvoiceItem(ctx, invoiceItemId).Execute()

Get Invoice Item

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
    invoiceItemId := "invoiceitem_dev_abc123" // string | Unique identifier of the invoice_item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceItemsAPI.GetInvoiceItem(context.Background(), invoiceItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceItemsAPI.GetInvoiceItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceItem`: InvoiceItemExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoiceItemsAPI.GetInvoiceItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceItemId** | **string** | Unique identifier of the invoice_item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceItemExternal**](InvoiceItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoiceItems

> ListResponseInvoiceItemExternal ListInvoiceItems(ctx).InvoiceItemsQueryParams(invoiceItemsQueryParams).Execute()

List Invoice Items

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
    invoiceItemsQueryParams := *openapiclient.NewInvoiceItemsQueryParams() // InvoiceItemsQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoiceItemsAPI.ListInvoiceItems(context.Background()).InvoiceItemsQueryParams(invoiceItemsQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceItemsAPI.ListInvoiceItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoiceItems`: ListResponseInvoiceItemExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoiceItemsAPI.ListInvoiceItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoiceItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemsQueryParams** | [**InvoiceItemsQueryParams**](InvoiceItemsQueryParams.md) |  | 

### Return type

[**ListResponseInvoiceItemExternal**](ListResponseInvoiceItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

