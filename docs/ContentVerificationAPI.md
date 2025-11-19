# \ContentVerificationAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPageChanges**](ContentVerificationAPI.md#GetPageChanges) | **Get** /api/projects/{projectId}/changes | Get page changes
[**VerifyContent**](ContentVerificationAPI.md#VerifyContent) | **Post** /api/projects/{projectId}/verify-content | Verify page content



## GetPageChanges

> GetPageChanges200Response GetPageChanges(ctx, projectId).Status(status).PageId(pageId).Limit(limit).Execute()

Get page changes



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
	status := "status_example" // string | Filter by change status (optional)
	pageId := "pageId_example" // string | Filter by page ID (optional)
	limit := int32(56) // int32 | Results per page (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentVerificationAPI.GetPageChanges(context.Background(), projectId).Status(status).PageId(pageId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentVerificationAPI.GetPageChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageChanges`: GetPageChanges200Response
	fmt.Fprintf(os.Stdout, "Response from `ContentVerificationAPI.GetPageChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by change status | 
 **pageId** | **string** | Filter by page ID | 
 **limit** | **int32** | Results per page | [default to 50]

### Return type

[**GetPageChanges200Response**](GetPageChanges200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

