# \PagesAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchSyncPages**](PagesAPI.md#BatchSyncPages) | **Post** /api/projects/{projectId}/pages/batch | Batch sync pages
[**ListPages**](PagesAPI.md#ListPages) | **Get** /api/projects/{projectId}/pages | List pages



## BatchSyncPages

> BatchSyncPages200Response BatchSyncPages(ctx, projectId).BatchSyncPagesRequest(batchSyncPagesRequest).Execute()

Batch sync pages



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
	batchSyncPagesRequest := *openapiclient.NewBatchSyncPagesRequest([]openapiclient.PageData{*openapiclient.NewPageData("https://example.com/page")}) // BatchSyncPagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PagesAPI.BatchSyncPages(context.Background(), projectId).BatchSyncPagesRequest(batchSyncPagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagesAPI.BatchSyncPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchSyncPages`: BatchSyncPages200Response
	fmt.Fprintf(os.Stdout, "Response from `PagesAPI.BatchSyncPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchSyncPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchSyncPagesRequest** | [**BatchSyncPagesRequest**](BatchSyncPagesRequest.md) |  | 

### Return type

[**BatchSyncPages200Response**](BatchSyncPages200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPages

> ListPages200Response ListPages(ctx, projectId).Limit(limit).Offset(offset).Execute()

List pages



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
	limit := int32(56) // int32 | Results per page (optional) (default to 50)
	offset := int32(56) // int32 | Pagination offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PagesAPI.ListPages(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PagesAPI.ListPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPages`: ListPages200Response
	fmt.Fprintf(os.Stdout, "Response from `PagesAPI.ListPages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Results per page | [default to 50]
 **offset** | **int32** | Pagination offset | [default to 0]

### Return type

[**ListPages200Response**](ListPages200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

