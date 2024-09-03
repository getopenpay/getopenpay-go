# \CreditNotesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreditNote**](CreditNotesAPI.md#CreateCreditNote) | **Post** /credit-notes/ | Create Credit Note
[**GetCreditNote**](CreditNotesAPI.md#GetCreditNote) | **Get** /credit-notes/{credit_note_id} | Get Credit Note
[**ListCreditNotes**](CreditNotesAPI.md#ListCreditNotes) | **Post** /credit-notes/list | List Credit Notes



## CreateCreditNote

> CreditNoteExternal CreateCreditNote(ctx).CreateCreditNoteRequest(createCreditNoteRequest).Execute()

Create Credit Note



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
    createCreditNoteRequest := *openapiclient.NewCreateCreditNoteRequest("InvoiceId_example", int32(123), []openapiclient.CreateCreditNoteLine{*openapiclient.NewCreateCreditNoteLine(int32(123), openapiclient.CurrencyEnum("usd"), openapiclient.CreditNoteLineType("invoice_line_item"))}) // CreateCreditNoteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.CreateCreditNote(context.Background()).CreateCreditNoteRequest(createCreditNoteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreateCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreditNote`: CreditNoteExternal
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreateCreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreditNoteRequest** | [**CreateCreditNoteRequest**](CreateCreditNoteRequest.md) |  | 

### Return type

[**CreditNoteExternal**](CreditNoteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNote

> CreditNoteExternal GetCreditNote(ctx, creditNoteId).Execute()

Get Credit Note

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
    creditNoteId := "credit_note_dev_abc123" // string | Unique Identifier of the credit_note.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.GetCreditNote(context.Background(), creditNoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.GetCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreditNote`: CreditNoteExternal
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.GetCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **string** | Unique Identifier of the credit_note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNoteExternal**](CreditNoteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCreditNotes

> ListResponseCreditNoteExternal ListCreditNotes(ctx).CreditNoteQueryParams(creditNoteQueryParams).Execute()

List Credit Notes

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
    creditNoteQueryParams := *openapiclient.NewCreditNoteQueryParams() // CreditNoteQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditNotesAPI.ListCreditNotes(context.Background()).CreditNoteQueryParams(creditNoteQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.ListCreditNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreditNotes`: ListResponseCreditNoteExternal
    fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.ListCreditNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreditNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteQueryParams** | [**CreditNoteQueryParams**](CreditNoteQueryParams.md) |  | 

### Return type

[**ListResponseCreditNoteExternal**](ListResponseCreditNoteExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

