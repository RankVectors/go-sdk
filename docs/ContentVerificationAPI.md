# com.rankvectors\ContentVerificationAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyContent**](ContentVerificationAPI.md#VerifyContent) | **Post** /api/projects/{projectId}/verify-content | Verify page content



## VerifyContent

> VerifyContent200Response VerifyContent(ctx, projectId).VerifyContentRequest(verifyContentRequest).Execute()

Verify page content



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
	projectId := "proj-123" // string | Unique identifier for the project
	verifyContentRequest := *openapiclient.NewVerifyContentRequest("https://example.com/page", "sugg-123") // VerifyContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentVerificationAPI.VerifyContent(context.Background(), projectId).VerifyContentRequest(verifyContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentVerificationAPI.VerifyContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyContent`: VerifyContent200Response
	fmt.Fprintf(os.Stdout, "Response from `ContentVerificationAPI.VerifyContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyContentRequest** | [**VerifyContentRequest**](VerifyContentRequest.md) |  | 

### Return type

[**VerifyContent200Response**](VerifyContent200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

