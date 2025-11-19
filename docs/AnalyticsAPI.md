# \AnalyticsAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectAnalytics**](AnalyticsAPI.md#GetProjectAnalytics) | **Get** /api/projects/{projectId}/analytics | Get project analytics



## GetProjectAnalytics

> ProjectAnalytics GetProjectAnalytics(ctx, projectId).StartDate(startDate).EndDate(endDate).Execute()

Get project analytics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/rankvectors/rankvectors-go-sdk"
)

func main() {
	projectId := "proj-123" // string | Unique identifier for the project
	startDate := time.Now() // time.Time | Start date for analytics (optional)
	endDate := time.Now() // time.Time | End date for analytics (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetProjectAnalytics(context.Background(), projectId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetProjectAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectAnalytics`: ProjectAnalytics
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetProjectAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Start date for analytics | 
 **endDate** | **time.Time** | End date for analytics | 

### Return type

[**ProjectAnalytics**](ProjectAnalytics.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

