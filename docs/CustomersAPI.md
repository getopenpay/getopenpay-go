# \CustomersAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersAPI.md#CreateCustomer) | **Post** /customers/ | Create Customer
[**CreateCustomerBalanceTransaction**](CustomersAPI.md#CreateCustomerBalanceTransaction) | **Post** /customers/{customer_id}/balance-transactions | Create Customer Balance Transaction
[**DeleteCustomerDiscount**](CustomersAPI.md#DeleteCustomerDiscount) | **Delete** /customers/{customer_id}/discount | Delete Customer Discount
[**GetCustomer**](CustomersAPI.md#GetCustomer) | **Get** /customers/{customer_id} | Get Customer
[**GetCustomerBalanceTransaction**](CustomersAPI.md#GetCustomerBalanceTransaction) | **Get** /customers/{customer_id}/balance-transactions/{transaction_id} | Get Customer Balance Transaction
[**GetCustomerBalanceTransactions**](CustomersAPI.md#GetCustomerBalanceTransactions) | **Get** /customers/{customer_id}/balance-transactions | Get Customer Balance Transactions
[**ListCustomerPaymentMethods**](CustomersAPI.md#ListCustomerPaymentMethods) | **Post** /customers/{customer_id}/payment-methods | List Customer Payment Methods
[**ListCustomers**](CustomersAPI.md#ListCustomers) | **Post** /customers/list | List Customers
[**ListValidSubscriptions**](CustomersAPI.md#ListValidSubscriptions) | **Post** /customers{customer_id}/list_valid_subscriptions | List Valid Subscriptions
[**SearchCustomers**](CustomersAPI.md#SearchCustomers) | **Post** /customers/search | Search Customers
[**UpdateCustomer**](CustomersAPI.md#UpdateCustomer) | **Put** /customers/{customer_external_id} | Update Customer



## CreateCustomer

> CustomerExternal CreateCustomer(ctx).CreateCustomerRequest(createCustomerRequest).Execute()

Create Customer

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
    createCustomerRequest := *openapiclient.NewCreateCustomerRequest("johndoes@xyz.com") // CreateCustomerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).CreateCustomerRequest(createCustomerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: CustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerRequest** | [**CreateCustomerRequest**](CreateCustomerRequest.md) |  | 

### Return type

[**CustomerExternal**](CustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerBalanceTransaction

> CustomerBalanceTransactionExternal CreateCustomerBalanceTransaction(ctx, customerId).CreateCustomerBalanceTransactionRequest(createCustomerBalanceTransactionRequest).Execute()

Create Customer Balance Transaction



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
    customerId := "customerId_example" // string | 
    createCustomerBalanceTransactionRequest := *openapiclient.NewCreateCustomerBalanceTransactionRequest(int32(123)) // CreateCustomerBalanceTransactionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.CreateCustomerBalanceTransaction(context.Background(), customerId).CreateCustomerBalanceTransactionRequest(createCustomerBalanceTransactionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomerBalanceTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerBalanceTransaction`: CustomerBalanceTransactionExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomerBalanceTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerBalanceTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomerBalanceTransactionRequest** | [**CreateCustomerBalanceTransactionRequest**](CreateCustomerBalanceTransactionRequest.md) |  | 

### Return type

[**CustomerBalanceTransactionExternal**](CustomerBalanceTransactionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerDiscount

> CustomerExternal DeleteCustomerDiscount(ctx, customerId).Execute()

Delete Customer Discount

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
    customerId := "customerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.DeleteCustomerDiscount(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.DeleteCustomerDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomerDiscount`: CustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.DeleteCustomerDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerDiscountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerExternal**](CustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> CustomerExternal GetCustomer(ctx, customerId).Expand(expand).Execute()

Get Customer

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
    customerId := "customerId_example" // string | 
    expand := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.GetCustomer(context.Background(), customerId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: CustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** |  | 

### Return type

[**CustomerExternal**](CustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerBalanceTransaction

> CustomerBalanceTransactionExternal GetCustomerBalanceTransaction(ctx, customerId, transactionId).Execute()

Get Customer Balance Transaction

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
    customerId := "customerId_example" // string | 
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.GetCustomerBalanceTransaction(context.Background(), customerId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerBalanceTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerBalanceTransaction`: CustomerBalanceTransactionExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerBalanceTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerBalanceTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomerBalanceTransactionExternal**](CustomerBalanceTransactionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerBalanceTransactions

> []CustomerBalanceTransactionExternal GetCustomerBalanceTransactions(ctx, customerId).Execute()

Get Customer Balance Transactions

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
    customerId := "customerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.GetCustomerBalanceTransactions(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerBalanceTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerBalanceTransactions`: []CustomerBalanceTransactionExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerBalanceTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerBalanceTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomerBalanceTransactionExternal**](CustomerBalanceTransactionExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerPaymentMethods

> ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal ListCustomerPaymentMethods(ctx, customerId).CustomerPaymentMethodQueryParams(customerPaymentMethodQueryParams).Execute()

List Customer Payment Methods



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
    customerId := "customerId_example" // string | Unique Identifier of the customer.
    customerPaymentMethodQueryParams := *openapiclient.NewCustomerPaymentMethodQueryParams() // CustomerPaymentMethodQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.ListCustomerPaymentMethods(context.Background(), customerId).CustomerPaymentMethodQueryParams(customerPaymentMethodQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomerPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomerPaymentMethods`: ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Unique Identifier of the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerPaymentMethodQueryParams** | [**CustomerPaymentMethodQueryParams**](CustomerPaymentMethodQueryParams.md) |  | 

### Return type

[**ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal**](ListResponseUnionTokenizedCreditCardExternalPaymentMethodExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> ListResponseCustomerExternal ListCustomers(ctx).CustomerQueryParams(customerQueryParams).Execute()

List Customers

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
    customerQueryParams := *openapiclient.NewCustomerQueryParams() // CustomerQueryParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.ListCustomers(context.Background()).CustomerQueryParams(customerQueryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomers`: ListResponseCustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerQueryParams** | [**CustomerQueryParams**](CustomerQueryParams.md) |  | 

### Return type

[**ListResponseCustomerExternal**](ListResponseCustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListValidSubscriptions

> ActiveSubResponse ListValidSubscriptions(ctx, customerId).ListActiveSubParams(listActiveSubParams).Execute()

List Valid Subscriptions



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
    customerId := "customerId_example" // string | Unique Identifier of the customer.
    listActiveSubParams := *openapiclient.NewListActiveSubParams() // ListActiveSubParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.ListValidSubscriptions(context.Background(), customerId).ListActiveSubParams(listActiveSubParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListValidSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListValidSubscriptions`: ActiveSubResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListValidSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Unique Identifier of the customer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListValidSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listActiveSubParams** | [**ListActiveSubParams**](ListActiveSubParams.md) |  | 

### Return type

[**ActiveSubResponse**](ActiveSubResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCustomers

> ListResponseCustomerExternal SearchCustomers(ctx).SearchCustomerRequest(searchCustomerRequest).Execute()

Search Customers

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
    searchCustomerRequest := *openapiclient.NewSearchCustomerRequest("Query_example") // SearchCustomerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.SearchCustomers(context.Background()).SearchCustomerRequest(searchCustomerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.SearchCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCustomers`: ListResponseCustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.SearchCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCustomerRequest** | [**SearchCustomerRequest**](SearchCustomerRequest.md) |  | 

### Return type

[**ListResponseCustomerExternal**](ListResponseCustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> CustomerExternal UpdateCustomer(ctx, customerExternalId).UpdateCustomerRequest(updateCustomerRequest).Execute()

Update Customer

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
    customerExternalId := "customerExternalId_example" // string | 
    updateCustomerRequest := *openapiclient.NewUpdateCustomerRequest() // UpdateCustomerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), customerExternalId).UpdateCustomerRequest(updateCustomerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: CustomerExternal
    fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerExternalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerRequest** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md) |  | 

### Return type

[**CustomerExternal**](CustomerExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

