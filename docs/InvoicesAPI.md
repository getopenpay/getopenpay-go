# \InvoicesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddComment**](InvoicesAPI.md#AddComment) | **Post** /invoices/{invoice_external_id}/comment | Add Comment
[**CreateInvoice**](InvoicesAPI.md#CreateInvoice) | **Post** /invoices/ | Create Invoice
[**FinalizeInvoice**](InvoicesAPI.md#FinalizeInvoice) | **Post** /invoices/{invoice_external_id}/finalize | Finalize Invoice
[**GetInvoice**](InvoicesAPI.md#GetInvoice) | **Get** /invoices/{invoice_external_id} | Get Invoice
[**GetInvoicePublic**](InvoicesAPI.md#GetInvoicePublic) | **Get** /invoices/public/{public_permanent_token} | Get Invoice Public
[**ListInvoices**](InvoicesAPI.md#ListInvoices) | **Post** /invoices/list | List Invoices
[**MarkInvoiceAsUncollectible**](InvoicesAPI.md#MarkInvoiceAsUncollectible) | **Post** /invoices/{invoice_external_id}/mark_uncollectible | Mark Invoice As Uncollectible
[**MarkInvoiceAsVoid**](InvoicesAPI.md#MarkInvoiceAsVoid) | **Post** /invoices/{invoice_external_id}/void | Mark Invoice As Void
[**PayInvoice**](InvoicesAPI.md#PayInvoice) | **Post** /invoices/{invoice_external_id}/pay | Pay Invoice
[**PreviewNextInvoice**](InvoicesAPI.md#PreviewNextInvoice) | **Get** /invoices/preview_next_invoice/{subscription_id} | Preview Next Invoice
[**SearchInvoices**](InvoicesAPI.md#SearchInvoices) | **Post** /invoices/search | Search Invoices
[**UpdateInvoice**](InvoicesAPI.md#UpdateInvoice) | **Put** /invoices/{invoice_id} | Update Invoice



## AddComment

> InvoiceExternal AddComment(ctx, invoiceExternalId).AddCommentRequest(addCommentRequest).Execute()

Add Comment

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    addCommentRequest := *openapiclient.NewAddCommentRequest("Marked as paid.") // AddCommentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.AddComment(context.Background(), invoiceExternalId).AddCommentRequest(addCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AddComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddComment`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AddComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCommentRequest** | [**AddCommentRequest**](AddCommentRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> InvoiceExternal CreateInvoice(ctx).CreateInvoiceRequest(createInvoiceRequest).Execute()

Create Invoice



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
    createInvoiceRequest := *openapiclient.NewCreateInvoiceRequest("CustomerId_example") // CreateInvoiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.CreateInvoice(context.Background()).CreateInvoiceRequest(createInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvoiceRequest** | [**CreateInvoiceRequest**](CreateInvoiceRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeInvoice

> InvoiceExternal FinalizeInvoice(ctx, invoiceExternalId).Execute()

Finalize Invoice

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.FinalizeInvoice(context.Background(), invoiceExternalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.FinalizeInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.FinalizeInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> InvoiceExternal GetInvoice(ctx, invoiceExternalId).Expand(expand).Execute()

Get Invoice

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    expand := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.GetInvoice(context.Background(), invoiceExternalId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicePublic

> InvoicePublic GetInvoicePublic(ctx, publicPermanentToken).Execute()

Get Invoice Public

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
    publicPermanentToken := "00000000-aaaa-bbbb-cccc-dddddddddddd" // string | Public token for the invoice.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.GetInvoicePublic(context.Background(), publicPermanentToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoicePublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoicePublic`: InvoicePublic
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoicePublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicPermanentToken** | **string** | Public token for the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoicePublic**](InvoicePublic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> ListResponseInvoiceExternal ListInvoices(ctx).InvoiceQueryParams(invoiceQueryParams).Execute()

List Invoices

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
    invoiceQueryParams := *openapiclient.NewInvoiceQueryParams() // InvoiceQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.ListInvoices(context.Background()).InvoiceQueryParams(invoiceQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListResponseInvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceQueryParams** | [**InvoiceQueryParams**](InvoiceQueryParams.md) |  | 

### Return type

[**ListResponseInvoiceExternal**](ListResponseInvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkInvoiceAsUncollectible

> InvoiceExternal MarkInvoiceAsUncollectible(ctx, invoiceExternalId).MarkUncollectibleRequest(markUncollectibleRequest).Execute()

Mark Invoice As Uncollectible

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    markUncollectibleRequest := *openapiclient.NewMarkUncollectibleRequest() // MarkUncollectibleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.MarkInvoiceAsUncollectible(context.Background(), invoiceExternalId).MarkUncollectibleRequest(markUncollectibleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.MarkInvoiceAsUncollectible``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkInvoiceAsUncollectible`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.MarkInvoiceAsUncollectible`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkInvoiceAsUncollectibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markUncollectibleRequest** | [**MarkUncollectibleRequest**](MarkUncollectibleRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkInvoiceAsVoid

> InvoiceExternal MarkInvoiceAsVoid(ctx, invoiceExternalId).MarkVoidRequest(markVoidRequest).Execute()

Mark Invoice As Void

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    markVoidRequest := *openapiclient.NewMarkVoidRequest() // MarkVoidRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.MarkInvoiceAsVoid(context.Background(), invoiceExternalId).MarkVoidRequest(markVoidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.MarkInvoiceAsVoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkInvoiceAsVoid`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.MarkInvoiceAsVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkInvoiceAsVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markVoidRequest** | [**MarkVoidRequest**](MarkVoidRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayInvoice

> InvoiceExternal PayInvoice(ctx, invoiceExternalId).PayInvoiceRequest(payInvoiceRequest).Execute()

Pay Invoice

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
    invoiceExternalId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    payInvoiceRequest := *openapiclient.NewPayInvoiceRequest() // PayInvoiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.PayInvoice(context.Background(), invoiceExternalId).PayInvoiceRequest(payInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.PayInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.PayInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceExternalId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payInvoiceRequest** | [**PayInvoiceRequest**](PayInvoiceRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewNextInvoice

> InvoiceExternal PreviewNextInvoice(ctx, subscriptionId).Execute()

Preview Next Invoice

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
    subscriptionId := "subscriptionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.PreviewNextInvoice(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.PreviewNextInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewNextInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.PreviewNextInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewNextInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInvoices

> ListResponseInvoiceExternal SearchInvoices(ctx).SearchInvoiceRequest(searchInvoiceRequest).Execute()

Search Invoices

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
    searchInvoiceRequest := *openapiclient.NewSearchInvoiceRequest("Query_example") // SearchInvoiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.SearchInvoices(context.Background()).SearchInvoiceRequest(searchInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.SearchInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchInvoices`: ListResponseInvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.SearchInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchInvoiceRequest** | [**SearchInvoiceRequest**](SearchInvoiceRequest.md) |  | 

### Return type

[**ListResponseInvoiceExternal**](ListResponseInvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> InvoiceExternal UpdateInvoice(ctx, invoiceId).UpdateInvoiceRequest(updateInvoiceRequest).Execute()

Update Invoice

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
    invoiceId := "invoice_dev_abc123" // string | Unique identifier of the invoice.
    updateInvoiceRequest := *openapiclient.NewUpdateInvoiceRequest() // UpdateInvoiceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesAPI.UpdateInvoice(context.Background(), invoiceId).UpdateInvoiceRequest(updateInvoiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInvoice`: InvoiceExternal
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | Unique identifier of the invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInvoiceRequest** | [**UpdateInvoiceRequest**](UpdateInvoiceRequest.md) |  | 

### Return type

[**InvoiceExternal**](InvoiceExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

