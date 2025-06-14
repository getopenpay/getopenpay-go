# \SignupAnswersAPI

All URIs are relative to *https://connto.getopenpay.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignupQuestionnaireStatus**](SignupAnswersAPI.md#SignupQuestionnaireStatus) | **Get** /signup-answers/ | Signup Questionnaire Status



## SignupQuestionnaireStatus

> SignupQuestionnaireStatus SignupQuestionnaireStatus(ctx).Execute()

Signup Questionnaire Status

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignupAnswersAPI.SignupQuestionnaireStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignupAnswersAPI.SignupQuestionnaireStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignupQuestionnaireStatus`: SignupQuestionnaireStatus
	fmt.Fprintf(os.Stdout, "Response from `SignupAnswersAPI.SignupQuestionnaireStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSignupQuestionnaireStatusRequest struct via the builder pattern


### Return type

[**SignupQuestionnaireStatus**](SignupQuestionnaireStatus.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

