# \TransactionManagementAPI

All URIs are relative to *https://secure.octo.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundPost**](TransactionManagementAPI.md#RefundPost) | **Post** /refund | Refund a payment
[**SetAcceptPost**](TransactionManagementAPI.md#SetAcceptPost) | **Post** /set_accept | Confirm or cancel a transaction



## RefundPost

> RefundResponse RefundPost(ctx).RefundRequest(refundRequest).Execute()

Refund a payment



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
	refundRequest := *openapiclient.NewRefundRequest(int32(123), "ShopRefundId_example", "OctoSecret_example", "OctoPaymentUUID_example", float64(123)) // RefundRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionManagementAPI.RefundPost(context.Background()).RefundRequest(refundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionManagementAPI.RefundPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundPost`: RefundResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionManagementAPI.RefundPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundRequest** | [**RefundRequest**](RefundRequest.md) |  | 

### Return type

[**RefundResponse**](RefundResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAcceptPost

> SetAcceptResponse SetAcceptPost(ctx).SetAcceptRequest(setAcceptRequest).Execute()

Confirm or cancel a transaction



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
	setAcceptRequest := *openapiclient.NewSetAcceptRequest(int32(123), "OctoSecret_example", "OctoPaymentUUID_example", "AcceptStatus_example", float64(123)) // SetAcceptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionManagementAPI.SetAcceptPost(context.Background()).SetAcceptRequest(setAcceptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionManagementAPI.SetAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAcceptPost`: SetAcceptResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionManagementAPI.SetAcceptPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setAcceptRequest** | [**SetAcceptRequest**](SetAcceptRequest.md) |  | 

### Return type

[**SetAcceptResponse**](SetAcceptResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

