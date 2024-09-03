# \EventsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /events/{event_id} | Get Event
[**ListEvents**](EventsAPI.md#ListEvents) | **Post** /events/list | List Events
[**SearchEvents**](EventsAPI.md#SearchEvents) | **Post** /events/search | Search Events



## GetEvent

> EventExternal GetEvent(ctx, eventId).Execute()

Get Event

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
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEvent(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: EventExternal
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventExternal**](EventExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListResponseEventExternal ListEvents(ctx).EventsQueryParams(eventsQueryParams).Execute()

List Events

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
    eventsQueryParams := *openapiclient.NewEventsQueryParams() // EventsQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.ListEvents(context.Background()).EventsQueryParams(eventsQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: ListResponseEventExternal
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventsQueryParams** | [**EventsQueryParams**](EventsQueryParams.md) |  | 

### Return type

[**ListResponseEventExternal**](ListResponseEventExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvents

> ListResponseEventExternal SearchEvents(ctx).EventSearchParams(eventSearchParams).Execute()

Search Events

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
    eventSearchParams := *openapiclient.NewEventSearchParams("ObjectId_example") // EventSearchParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.SearchEvents(context.Background()).EventSearchParams(eventSearchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.SearchEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEvents`: ListResponseEventExternal
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.SearchEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSearchParams** | [**EventSearchParams**](EventSearchParams.md) |  | 

### Return type

[**ListResponseEventExternal**](ListResponseEventExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

