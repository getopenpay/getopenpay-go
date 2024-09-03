# \BillingPortalAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortalSession**](BillingPortalAPI.md#CreatePortalSession) | **Post** /billing_portal/sessions | Create Portal Session



## CreatePortalSession

> PortalSessionExternal CreatePortalSession(ctx).CreatePortalSessionRequest(createPortalSessionRequest).Execute()

Create Portal Session

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
    createPortalSessionRequest := *openapiclient.NewCreatePortalSessionRequest("cust_dev_abc123", "http://example.com/account") // CreatePortalSessionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPortalAPI.CreatePortalSession(context.Background()).CreatePortalSessionRequest(createPortalSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPortalAPI.CreatePortalSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePortalSession`: PortalSessionExternal
    fmt.Fprintf(os.Stdout, "Response from `BillingPortalAPI.CreatePortalSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortalSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortalSessionRequest** | [**CreatePortalSessionRequest**](CreatePortalSessionRequest.md) |  | 

### Return type

[**PortalSessionExternal**](PortalSessionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

