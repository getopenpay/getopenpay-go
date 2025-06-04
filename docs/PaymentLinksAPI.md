# \PaymentLinksAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentLink**](PaymentLinksAPI.md#CreatePaymentLink) | **Post** /payment-link/ | Create Payment Link
[**GetPaymentLink**](PaymentLinksAPI.md#GetPaymentLink) | **Get** /payment-link/{plink_id} | Get Payment Link
[**ListPaymentLinks**](PaymentLinksAPI.md#ListPaymentLinks) | **Post** /payment-link/list | List Payment Links
[**OpenPaymentLinkPagePublic**](PaymentLinksAPI.md#OpenPaymentLinkPagePublic) | **Get** /payment-link/public/{secure_token} | Open Payment Link Page Public



## CreatePaymentLink

> PaymentLinkExternal CreatePaymentLink(ctx).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()

Create Payment Link

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
	createPaymentLinkRequest := *openapiclient.NewCreatePaymentLinkRequest(openapiclient.CheckoutMode("payment"), "SuccessUrl_example") // CreatePaymentLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.CreatePaymentLink(context.Background()).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.CreatePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentLink`: PaymentLinkExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.CreatePaymentLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentLinkRequest** | [**CreatePaymentLinkRequest**](CreatePaymentLinkRequest.md) |  | 

### Return type

[**PaymentLinkExternal**](PaymentLinkExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentLink

> PaymentLinkExternal GetPaymentLink(ctx, plinkId).Execute()

Get Payment Link

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
	plinkId := "plink_dev_abc123" // string | Unique identifier of the payment link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.GetPaymentLink(context.Background(), plinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.GetPaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentLink`: PaymentLinkExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.GetPaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plinkId** | **string** | Unique identifier of the payment link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinkExternal**](PaymentLinkExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentLinks

> ListResponsePaymentLinkExternal ListPaymentLinks(ctx).PaymentLinkQueryParams(paymentLinkQueryParams).Execute()

List Payment Links

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
	paymentLinkQueryParams := *openapiclient.NewPaymentLinkQueryParams() // PaymentLinkQueryParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.ListPaymentLinks(context.Background()).PaymentLinkQueryParams(paymentLinkQueryParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.ListPaymentLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentLinks`: ListResponsePaymentLinkExternal
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.ListPaymentLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentLinkQueryParams** | [**PaymentLinkQueryParams**](PaymentLinkQueryParams.md) |  | 

### Return type

[**ListResponsePaymentLinkExternal**](ListResponsePaymentLinkExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenPaymentLinkPagePublic

> interface{} OpenPaymentLinkPagePublic(ctx, secureToken).RefId(refId).CustomerId(customerId).Execute()

Open Payment Link Page Public

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
	secureToken := "00000000-aaaa-bbbb-cccc-dddddddddddd" // string | Secure token of the payment link.
	refId := "ORDER-12345" // string | Optional reference ID for the payment link (optional)
	customerId := "cus_dev_abc123" // string | Optional ID for the customer using the payment link (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.OpenPaymentLinkPagePublic(context.Background(), secureToken).RefId(refId).CustomerId(customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.OpenPaymentLinkPagePublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenPaymentLinkPagePublic`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.OpenPaymentLinkPagePublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureToken** | **string** | Secure token of the payment link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenPaymentLinkPagePublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refId** | **string** | Optional reference ID for the payment link | 
 **customerId** | **string** | Optional ID for the customer using the payment link | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

