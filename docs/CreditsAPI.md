# com.rankvectors\CreditsAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredits**](CreditsAPI.md#AddCredits) | **Post** /api/projects/{projectId}/credits | Add credits
[**GetCredits**](CreditsAPI.md#GetCredits) | **Get** /api/projects/{projectId}/credits | Get credit balance



## AddCredits

> AddCredits200Response AddCredits(ctx, projectId).AddCreditsRequest(addCreditsRequest).Execute()

Add credits



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
	addCreditsRequest := *openapiclient.NewAddCreditsRequest(float32(100)) // AddCreditsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditsAPI.AddCredits(context.Background(), projectId).AddCreditsRequest(addCreditsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditsAPI.AddCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCredits`: AddCredits200Response
	fmt.Fprintf(os.Stdout, "Response from `CreditsAPI.AddCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCreditsRequest** | [**AddCreditsRequest**](AddCreditsRequest.md) |  | 

### Return type

[**AddCredits200Response**](AddCredits200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredits

> CreditBalance GetCredits(ctx, projectId).IncludeHistory(includeHistory).StartDate(startDate).EndDate(endDate).Execute()

Get credit balance



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
	includeHistory := true // bool | Include usage history (optional) (default to false)
	startDate := time.Now() // time.Time | History start date (ISO 8601) (optional)
	endDate := time.Now() // time.Time | History end date (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditsAPI.GetCredits(context.Background(), projectId).IncludeHistory(includeHistory).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditsAPI.GetCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCredits`: CreditBalance
	fmt.Fprintf(os.Stdout, "Response from `CreditsAPI.GetCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHistory** | **bool** | Include usage history | [default to false]
 **startDate** | **time.Time** | History start date (ISO 8601) | 
 **endDate** | **time.Time** | History end date (ISO 8601) | 

### Return type

[**CreditBalance**](CreditBalance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

