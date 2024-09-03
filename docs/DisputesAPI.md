# \DisputesAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDispute**](DisputesAPI.md#CreateDispute) | **Post** /dispute/ | Create Dispute
[**GetDispute**](DisputesAPI.md#GetDispute) | **Get** /dispute/{dispute_id} | Get Dispute



## CreateDispute

> DisputeExternal CreateDispute(ctx).CreateDisputeRequest(createDisputeRequest).Execute()

Create Dispute

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
    createDisputeRequest := *openapiclient.NewCreateDisputeRequest() // CreateDisputeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisputesAPI.CreateDispute(context.Background()).CreateDisputeRequest(createDisputeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.CreateDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDispute`: DisputeExternal
    fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.CreateDispute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDisputeRequest** | [**CreateDisputeRequest**](CreateDisputeRequest.md) |  | 

### Return type

[**DisputeExternal**](DisputeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispute

> DisputeExternal GetDispute(ctx, disputeId).Execute()

Get Dispute

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
    disputeId := "disputeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisputesAPI.GetDispute(context.Background(), disputeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.GetDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDispute`: DisputeExternal
    fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.GetDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisputeExternal**](DisputeExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

