# \IntegrationAPI

All URIs are relative to *https://secure.octo.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackPost**](IntegrationAPI.md#CallbackPost) | **Post** /callback | Handle Octo callback
[**NotificationPost**](IntegrationAPI.md#NotificationPost) | **Post** /notification | Receive notification from Octo



## CallbackPost

> CallbackResponse CallbackPost(ctx).CallbackRequest(callbackRequest).Execute()

Handle Octo callback



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
	callbackRequest := *openapiclient.NewCallbackRequest("OctoSecret_example", "OctoPaymentUUID_example", "AcceptStatus_example") // CallbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.CallbackPost(context.Background()).CallbackRequest(callbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CallbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallbackPost`: CallbackResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CallbackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackRequest** | [**CallbackRequest**](CallbackRequest.md) |  | 

### Return type

[**CallbackResponse**](CallbackResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationPost

> NotificationPost(ctx).NotificationRequest(notificationRequest).Execute()

Receive notification from Octo



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
	notificationRequest := *openapiclient.NewNotificationRequest("ShopTransactionId_example", "OctoPaymentUUID_example", "Status_example", "Signature_example", "HashKey_example") // NotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.NotificationPost(context.Background()).NotificationRequest(notificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.NotificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

