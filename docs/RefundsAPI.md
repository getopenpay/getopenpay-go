# \RefundsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRefund**](RefundsAPI.md#CreateRefund) | **Post** /refunds/ | Create Refund
[**ListRefunds**](RefundsAPI.md#ListRefunds) | **Post** /refunds/list | List Refunds



## CreateRefund

> RefundExternal CreateRefund(ctx).CreateRefundRequest(createRefundRequest).Execute()

Create Refund

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
    createRefundRequest := *openapiclient.NewCreateRefundRequest() // CreateRefundRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsAPI.CreateRefund(context.Background()).CreateRefundRequest(createRefundRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.CreateRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRefund`: RefundExternal
    fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.CreateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRefundRequest** | [**CreateRefundRequest**](CreateRefundRequest.md) |  | 

### Return type

[**RefundExternal**](RefundExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefunds

> ListResponseRefundExternal ListRefunds(ctx).RefundQueryParams(refundQueryParams).Execute()

List Refunds

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
    refundQueryParams := *openapiclient.NewRefundQueryParams() // RefundQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsAPI.ListRefunds(context.Background()).RefundQueryParams(refundQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.ListRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRefunds`: ListResponseRefundExternal
    fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.ListRefunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundQueryParams** | [**RefundQueryParams**](RefundQueryParams.md) |  | 

### Return type

[**ListResponseRefundExternal**](ListResponseRefundExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

