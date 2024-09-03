# \SubscriptionItemsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptionItem**](SubscriptionItemsAPI.md#CreateSubscriptionItem) | **Post** /subscription-items/ | Create Subscription Item
[**DeleteSubscriptionItem**](SubscriptionItemsAPI.md#DeleteSubscriptionItem) | **Delete** /subscription-items/{subscription_item_id} | Delete Subscription Item
[**GetSubscriptionItem**](SubscriptionItemsAPI.md#GetSubscriptionItem) | **Get** /subscription-items/{subscription_item_id} | Get Subscription Item
[**ListSubscriptionItems**](SubscriptionItemsAPI.md#ListSubscriptionItems) | **Post** /subscription-items/list | List Subscription Items
[**UpdateSubscriptionItem**](SubscriptionItemsAPI.md#UpdateSubscriptionItem) | **Put** /subscription-items/{subscription_item_id} | Update Subscription Item



## CreateSubscriptionItem

> SubscriptionItemExternal CreateSubscriptionItem(ctx).CreateSubscriptionItemRequest(createSubscriptionItemRequest).Execute()

Create Subscription Item



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
    createSubscriptionItemRequest := *openapiclient.NewCreateSubscriptionItemRequest("subi_dev_abc123", "price_dev_abc123") // CreateSubscriptionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionItemsAPI.CreateSubscriptionItem(context.Background()).CreateSubscriptionItemRequest(createSubscriptionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionItemsAPI.CreateSubscriptionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscriptionItem`: SubscriptionItemExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionItemsAPI.CreateSubscriptionItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionItemRequest** | [**CreateSubscriptionItemRequest**](CreateSubscriptionItemRequest.md) |  | 

### Return type

[**SubscriptionItemExternal**](SubscriptionItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionItem

> DeleteSubscriptionItemResponse DeleteSubscriptionItem(ctx, subscriptionItemId).DeleteSubscriptionItemRequest(deleteSubscriptionItemRequest).Execute()

Delete Subscription Item



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
    subscriptionItemId := "subi_dev_abc123" // string | Unique identifier of the subscription_item.
    deleteSubscriptionItemRequest := *openapiclient.NewDeleteSubscriptionItemRequest() // DeleteSubscriptionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionItemsAPI.DeleteSubscriptionItem(context.Background(), subscriptionItemId).DeleteSubscriptionItemRequest(deleteSubscriptionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionItemsAPI.DeleteSubscriptionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriptionItem`: DeleteSubscriptionItemResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionItemsAPI.DeleteSubscriptionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionItemId** | **string** | Unique identifier of the subscription_item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSubscriptionItemRequest** | [**DeleteSubscriptionItemRequest**](DeleteSubscriptionItemRequest.md) |  | 

### Return type

[**DeleteSubscriptionItemResponse**](DeleteSubscriptionItemResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionItem

> SubscriptionItemExternal GetSubscriptionItem(ctx, subscriptionItemId).Execute()

Get Subscription Item

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
    subscriptionItemId := "subi_dev_abc123" // string | Unique identifier of the subscription_item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionItemsAPI.GetSubscriptionItem(context.Background(), subscriptionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionItemsAPI.GetSubscriptionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptionItem`: SubscriptionItemExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionItemsAPI.GetSubscriptionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionItemId** | **string** | Unique identifier of the subscription_item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionItemExternal**](SubscriptionItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionItems

> ListResponseSubscriptionItemExternal ListSubscriptionItems(ctx).SubscriptionItemQueryParams(subscriptionItemQueryParams).Execute()

List Subscription Items

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
    subscriptionItemQueryParams := *openapiclient.NewSubscriptionItemQueryParams("subi_dev_abc123") // SubscriptionItemQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionItemsAPI.ListSubscriptionItems(context.Background()).SubscriptionItemQueryParams(subscriptionItemQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionItemsAPI.ListSubscriptionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionItems`: ListResponseSubscriptionItemExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionItemsAPI.ListSubscriptionItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionItemQueryParams** | [**SubscriptionItemQueryParams**](SubscriptionItemQueryParams.md) |  | 

### Return type

[**ListResponseSubscriptionItemExternal**](ListResponseSubscriptionItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscriptionItem

> SubscriptionItemExternal UpdateSubscriptionItem(ctx, subscriptionItemId).UpdateSubscriptionItemRequest(updateSubscriptionItemRequest).Execute()

Update Subscription Item

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
    subscriptionItemId := "subi_dev_abc123" // string | Unique identifier of the subscription_item.
    updateSubscriptionItemRequest := *openapiclient.NewUpdateSubscriptionItemRequest() // UpdateSubscriptionItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionItemsAPI.UpdateSubscriptionItem(context.Background(), subscriptionItemId).UpdateSubscriptionItemRequest(updateSubscriptionItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionItemsAPI.UpdateSubscriptionItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscriptionItem`: SubscriptionItemExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionItemsAPI.UpdateSubscriptionItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionItemId** | **string** | Unique identifier of the subscription_item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriptionItemRequest** | [**UpdateSubscriptionItemRequest**](UpdateSubscriptionItemRequest.md) |  | 

### Return type

[**SubscriptionItemExternal**](SubscriptionItemExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

