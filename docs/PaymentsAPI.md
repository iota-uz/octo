# \PaymentsAPI

All URIs are relative to *https://secure.octo.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreparePaymentPost**](PaymentsAPI.md#PreparePaymentPost) | **Post** /prepare_payment | Initiate or check payment



## PreparePaymentPost

> PreparePaymentResponse PreparePaymentPost(ctx).PreparePaymentRequest(preparePaymentRequest).Execute()

Initiate or check payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/octo/octoapi"
)

func main() {
	preparePaymentRequest := *openapiclient.NewPreparePaymentRequest(int32(123), "OctoSecret_example", "ShopTransactionId_example", false, false, "InitTime_example", float64(123), "Currency_example", "Description_example", "ReturnUrl_example", "NotifyUrl_example", "Language_example") // PreparePaymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PreparePaymentPost(context.Background()).PreparePaymentRequest(preparePaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PreparePaymentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreparePaymentPost`: PreparePaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PreparePaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreparePaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preparePaymentRequest** | [**PreparePaymentRequest**](PreparePaymentRequest.md) |  | 

### Return type

[**PreparePaymentResponse**](PreparePaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

