# com.rankvectors\WebhooksAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImplementationInstructions**](WebhooksAPI.md#GetImplementationInstructions) | **Get** /api/webhook/implement-link | Get implementation instructions
[**ReportImplementationStatus**](WebhooksAPI.md#ReportImplementationStatus) | **Post** /api/webhook/implement-link | Report implementation status



## GetImplementationInstructions

> ImplementationInstructions GetImplementationInstructions(ctx).SuggestionId(suggestionId).Execute()

Get implementation instructions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rankvectors/rankvectors-go-sdk"
)

func main() {
	suggestionId := "sugg-123" // string | ID of the suggestion to implement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetImplementationInstructions(context.Background()).SuggestionId(suggestionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetImplementationInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImplementationInstructions`: ImplementationInstructions
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetImplementationInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImplementationInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suggestionId** | **string** | ID of the suggestion to implement | 

### Return type

[**ImplementationInstructions**](ImplementationInstructions.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportImplementationStatus

> ReportImplementationStatus200Response ReportImplementationStatus(ctx).ReportImplementationStatusRequest(reportImplementationStatusRequest).Execute()

Report implementation status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rankvectors/rankvectors-go-sdk"
)

func main() {
	reportImplementationStatusRequest := *openapiclient.NewReportImplementationStatusRequest("sugg-123", "rv_1234567890abcdef", "success") // ReportImplementationStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.ReportImplementationStatus(context.Background()).ReportImplementationStatusRequest(reportImplementationStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ReportImplementationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportImplementationStatus`: ReportImplementationStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ReportImplementationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportImplementationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportImplementationStatusRequest** | [**ReportImplementationStatusRequest**](ReportImplementationStatusRequest.md) |  | 

### Return type

[**ReportImplementationStatus200Response**](ReportImplementationStatus200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

