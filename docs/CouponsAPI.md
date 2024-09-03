# \CouponsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCoupon**](CouponsAPI.md#CreateCoupon) | **Post** /coupons/ | Create Coupon
[**DeleteCoupon**](CouponsAPI.md#DeleteCoupon) | **Delete** /coupons/{coupon_id} | Delete Coupon
[**GetCoupon**](CouponsAPI.md#GetCoupon) | **Get** /coupons/{coupon_id} | Get Coupon
[**ListCoupons**](CouponsAPI.md#ListCoupons) | **Post** /coupons/list | List Coupons
[**UpdateCoupon**](CouponsAPI.md#UpdateCoupon) | **Put** /coupons/{coupon_id} | Update Coupon



## CreateCoupon

> CouponExternal CreateCoupon(ctx).CreateCouponRequest(createCouponRequest).Execute()

Create Coupon

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
    createCouponRequest := *openapiclient.NewCreateCouponRequest("Name_example") // CreateCouponRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.CreateCoupon(context.Background()).CreateCouponRequest(createCouponRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CreateCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCoupon`: CouponExternal
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CreateCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCouponRequest** | [**CreateCouponRequest**](CreateCouponRequest.md) |  | 

### Return type

[**CouponExternal**](CouponExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupon

> DeleteCoupon(ctx, couponId).Execute()

Delete Coupon

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
    couponId := "coupon_dev_abc123" // string | Unique identifier of the Coupon.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CouponsAPI.DeleteCoupon(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.DeleteCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **string** | Unique identifier of the Coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoupon

> CouponExternal GetCoupon(ctx, couponId).Execute()

Get Coupon

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
    couponId := "coupon_dev_abc123" // string | Unique identifier of the Coupon.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.GetCoupon(context.Background(), couponId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.GetCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCoupon`: CouponExternal
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.GetCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **string** | Unique identifier of the Coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CouponExternal**](CouponExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCoupons

> ListResponseCouponExternal ListCoupons(ctx).CouponQueryParams(couponQueryParams).Execute()

List Coupons

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
    couponQueryParams := *openapiclient.NewCouponQueryParams() // CouponQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.ListCoupons(context.Background()).CouponQueryParams(couponQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.ListCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCoupons`: ListResponseCouponExternal
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.ListCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponQueryParams** | [**CouponQueryParams**](CouponQueryParams.md) |  | 

### Return type

[**ListResponseCouponExternal**](ListResponseCouponExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> CouponExternal UpdateCoupon(ctx, couponId).UpdateCouponRequest(updateCouponRequest).Execute()

Update Coupon

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
    couponId := "coupon_dev_abc123" // string | Unique identifier of the Coupon.
    updateCouponRequest := *openapiclient.NewUpdateCouponRequest() // UpdateCouponRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponsAPI.UpdateCoupon(context.Background(), couponId).UpdateCouponRequest(updateCouponRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.UpdateCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCoupon`: CouponExternal
    fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.UpdateCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **string** | Unique identifier of the Coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCouponRequest** | [**UpdateCouponRequest**](UpdateCouponRequest.md) |  | 

### Return type

[**CouponExternal**](CouponExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

