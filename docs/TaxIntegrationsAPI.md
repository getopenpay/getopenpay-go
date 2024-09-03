# \TaxIntegrationsAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaxIntegration**](TaxIntegrationsAPI.md#CreateTaxIntegration) | **Post** /integrations/tax/ | Create Tax Integration



## CreateTaxIntegration

> TaxIntegrationExternal CreateTaxIntegration(ctx).CreateTaxIntegrationRequest(createTaxIntegrationRequest).Execute()

Create Tax Integration

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
    createTaxIntegrationRequest := *openapiclient.NewCreateTaxIntegrationRequest(openapiclient.TaxIntegrationApiName("anrok_v1"), map[string]string({api_key=00000000000/aaaaaaaaaaaaa/secret.ZZZZZZZZZZZZZZZZZZZZZZZZZZ}), "US/Pacific") // CreateTaxIntegrationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxIntegrationsAPI.CreateTaxIntegration(context.Background()).CreateTaxIntegrationRequest(createTaxIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxIntegrationsAPI.CreateTaxIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTaxIntegration`: TaxIntegrationExternal
    fmt.Fprintf(os.Stdout, "Response from `TaxIntegrationsAPI.CreateTaxIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTaxIntegrationRequest** | [**CreateTaxIntegrationRequest**](CreateTaxIntegrationRequest.md) |  | 

### Return type

[**TaxIntegrationExternal**](TaxIntegrationExternal.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

