# \SubscriptionsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscriptionTrial**](SubscriptionsAPI.md#CancelSubscriptionTrial) | **Post** /subscriptions/{subscription_id}/cancel-subscription-trial | Cancel Subscription Trial
[**CreateSubscriptions**](SubscriptionsAPI.md#CreateSubscriptions) | **Post** /subscriptions/ | Create Subscriptions
[**DeleteSubscription**](SubscriptionsAPI.md#DeleteSubscription) | **Delete** /subscriptions/{subscription_id} | Delete Subscription
[**DeleteSubscriptionDiscount**](SubscriptionsAPI.md#DeleteSubscriptionDiscount) | **Delete** /subscriptions/{subscription_id}/discount | Delete Subscription Discount
[**GetSubscription**](SubscriptionsAPI.md#GetSubscription) | **Get** /subscriptions/{subscription_id} | Get Subscription
[**ListSubscriptions**](SubscriptionsAPI.md#ListSubscriptions) | **Post** /subscriptions/list | List Subscriptions
[**PauseSubscription**](SubscriptionsAPI.md#PauseSubscription) | **Put** /subscriptions/{subscription_id}/pause | Pause Subscription
[**RefreshSubscriptionStatus**](SubscriptionsAPI.md#RefreshSubscriptionStatus) | **Post** /subscriptions/{subscription_id}/refresh-status | Refresh Subscription Status
[**ResumeSubscription**](SubscriptionsAPI.md#ResumeSubscription) | **Put** /subscriptions/{subscription_id}/resume | Resume Subscription
[**SearchSubscriptions**](SubscriptionsAPI.md#SearchSubscriptions) | **Post** /subscriptions/search | Search Subscriptions
[**UpdateSubscription**](SubscriptionsAPI.md#UpdateSubscription) | **Put** /subscriptions/{subscription_id} | Update Subscription



## CancelSubscriptionTrial

> SubscriptionExternal CancelSubscriptionTrial(ctx, subscriptionId).Execute()

Cancel Subscription Trial



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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.CancelSubscriptionTrial(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CancelSubscriptionTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelSubscriptionTrial`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CancelSubscriptionTrial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscriptions

> CreateSubscriptionResponse CreateSubscriptions(ctx).CreateSubscriptionRequest(createSubscriptionRequest).Execute()

Create Subscriptions

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
    createSubscriptionRequest := *openapiclient.NewCreateSubscriptionRequest("cus_dev_abc123", []openapiclient.SelectedPriceQuantity{*openapiclient.NewSelectedPriceQuantity("PriceId_example", int32(123))}, int32(9000)) // CreateSubscriptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.CreateSubscriptions(context.Background()).CreateSubscriptionRequest(createSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CreateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscriptions`: CreateSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CreateSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionRequest** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md) |  | 

### Return type

[**CreateSubscriptionResponse**](CreateSubscriptionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> SubscriptionExternal DeleteSubscription(ctx, subscriptionId).DeleteSubscriptionRequest(deleteSubscriptionRequest).Execute()

Delete Subscription

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.
    deleteSubscriptionRequest := *openapiclient.NewDeleteSubscriptionRequest() // DeleteSubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.DeleteSubscription(context.Background(), subscriptionId).DeleteSubscriptionRequest(deleteSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscription`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.DeleteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSubscriptionRequest** | [**DeleteSubscriptionRequest**](DeleteSubscriptionRequest.md) |  | 

### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionDiscount

> SubscriptionExternal DeleteSubscriptionDiscount(ctx, subscriptionId).Execute()

Delete Subscription Discount

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.DeleteSubscriptionDiscount(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.DeleteSubscriptionDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriptionDiscount`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.DeleteSubscriptionDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionDiscountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> SubscriptionExternal GetSubscription(ctx, subscriptionId).Execute()

Get Subscription

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.GetSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscription`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> ListResponseSubscriptionExternal ListSubscriptions(ctx).SubscriptionQueryParams(subscriptionQueryParams).Execute()

List Subscriptions

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
    subscriptionQueryParams := *openapiclient.NewSubscriptionQueryParams() // SubscriptionQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.ListSubscriptions(context.Background()).SubscriptionQueryParams(subscriptionQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptions`: ListResponseSubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionQueryParams** | [**SubscriptionQueryParams**](SubscriptionQueryParams.md) |  | 

### Return type

[**ListResponseSubscriptionExternal**](ListResponseSubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseSubscription

> SubscriptionExternal PauseSubscription(ctx, subscriptionId).SubscriptionPauseRequest(subscriptionPauseRequest).Execute()

Pause Subscription

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.
    subscriptionPauseRequest := *openapiclient.NewSubscriptionPauseRequest(int32(123)) // SubscriptionPauseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.PauseSubscription(context.Background(), subscriptionId).SubscriptionPauseRequest(subscriptionPauseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PauseSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseSubscription`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PauseSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPauseRequest** | [**SubscriptionPauseRequest**](SubscriptionPauseRequest.md) |  | 

### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshSubscriptionStatus

> SubscriptionExternal RefreshSubscriptionStatus(ctx, subscriptionId).Execute()

Refresh Subscription Status

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.RefreshSubscriptionStatus(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.RefreshSubscriptionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshSubscriptionStatus`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.RefreshSubscriptionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshSubscriptionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeSubscription

> SubscriptionExternal ResumeSubscription(ctx, subscriptionId).SubscriptionResumeRequest(subscriptionResumeRequest).Execute()

Resume Subscription

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.
    subscriptionResumeRequest := *openapiclient.NewSubscriptionResumeRequest() // SubscriptionResumeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.ResumeSubscription(context.Background(), subscriptionId).SubscriptionResumeRequest(subscriptionResumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ResumeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeSubscription`: SubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ResumeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionResumeRequest** | [**SubscriptionResumeRequest**](SubscriptionResumeRequest.md) |  | 

### Return type

[**SubscriptionExternal**](SubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSubscriptions

> ListResponseSubscriptionExternal SearchSubscriptions(ctx).SearchSubscriptionRequest(searchSubscriptionRequest).Execute()

Search Subscriptions

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
    searchSubscriptionRequest := *openapiclient.NewSearchSubscriptionRequest("Query_example") // SearchSubscriptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.SearchSubscriptions(context.Background()).SearchSubscriptionRequest(searchSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SearchSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSubscriptions`: ListResponseSubscriptionExternal
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SearchSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSubscriptionRequest** | [**SearchSubscriptionRequest**](SearchSubscriptionRequest.md) |  | 

### Return type

[**ListResponseSubscriptionExternal**](ListResponseSubscriptionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> UpdateSubscriptionResponse UpdateSubscription(ctx, subscriptionId).UpdateSubscriptionRequest(updateSubscriptionRequest).Execute()

Update Subscription

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
    subscriptionId := "subscription_dev_abc123" // string | Unique identifier of the subscription.
    updateSubscriptionRequest := *openapiclient.NewUpdateSubscriptionRequest() // UpdateSubscriptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsAPI.UpdateSubscription(context.Background(), subscriptionId).UpdateSubscriptionRequest(updateSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: UpdateSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier of the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriptionRequest** | [**UpdateSubscriptionRequest**](UpdateSubscriptionRequest.md) |  | 

### Return type

[**UpdateSubscriptionResponse**](UpdateSubscriptionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

