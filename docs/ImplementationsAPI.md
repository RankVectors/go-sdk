# \ImplementationsAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImplementation**](ImplementationsAPI.md#GetImplementation) | **Get** /api/projects/{projectId}/implementations/{implementationId} | Get implementation details
[**ImplementLinks**](ImplementationsAPI.md#ImplementLinks) | **Post** /api/projects/{projectId}/implementations | Implement link suggestions
[**ListImplementations**](ImplementationsAPI.md#ListImplementations) | **Get** /api/projects/{projectId}/implementations | List implementations
[**RollbackImplementation**](ImplementationsAPI.md#RollbackImplementation) | **Post** /api/projects/{projectId}/implementations/{implementationId}/rollback | Rollback implementation



## GetImplementation

> GetImplementation200Response GetImplementation(ctx, projectId, implementationId).Execute()

Get implementation details



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
	implementationId := "impl-123" // string | Unique identifier for the implementation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImplementationsAPI.GetImplementation(context.Background(), projectId, implementationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImplementationsAPI.GetImplementation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImplementation`: GetImplementation200Response
	fmt.Fprintf(os.Stdout, "Response from `ImplementationsAPI.GetImplementation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 
**implementationId** | **string** | Unique identifier for the implementation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImplementationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetImplementation200Response**](GetImplementation200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImplementLinks

> ImplementationResponse ImplementLinks(ctx, projectId).ImplementationRequest(implementationRequest).Execute()

Implement link suggestions



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
	implementationRequest := *openapiclient.NewImplementationRequest([]string{"SuggestionIds_example"}, "custom", *openapiclient.NewCustomCredentials("https://yourapi.com/rankvectors-webhook", "your-webhook-secret-key")) // ImplementationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImplementationsAPI.ImplementLinks(context.Background(), projectId).ImplementationRequest(implementationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImplementationsAPI.ImplementLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImplementLinks`: ImplementationResponse
	fmt.Fprintf(os.Stdout, "Response from `ImplementationsAPI.ImplementLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiImplementLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **implementationRequest** | [**ImplementationRequest**](ImplementationRequest.md) |  | 

### Return type

[**ImplementationResponse**](ImplementationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImplementations

> ListImplementations200Response ListImplementations(ctx, projectId).Status(status).Platform(platform).Limit(limit).Offset(offset).Execute()

List implementations



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
	status := "status_example" // string | Filter by implementation status (optional)
	platform := "platform_example" // string | Filter by platform (optional)
	limit := int32(56) // int32 | Results per page (optional) (default to 50)
	offset := int32(56) // int32 | Pagination offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImplementationsAPI.ListImplementations(context.Background(), projectId).Status(status).Platform(platform).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImplementationsAPI.ListImplementations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImplementations`: ListImplementations200Response
	fmt.Fprintf(os.Stdout, "Response from `ImplementationsAPI.ListImplementations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImplementationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by implementation status | 
 **platform** | **string** | Filter by platform | 
 **limit** | **int32** | Results per page | [default to 50]
 **offset** | **int32** | Pagination offset | [default to 0]

### Return type

[**ListImplementations200Response**](ListImplementations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackImplementation

> RollbackImplementation200Response RollbackImplementation(ctx, projectId, implementationId).RollbackImplementationRequest(rollbackImplementationRequest).Execute()

Rollback implementation



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
	implementationId := "impl-123" // string | Unique identifier for the implementation
	rollbackImplementationRequest := *openapiclient.NewRollbackImplementationRequest() // RollbackImplementationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImplementationsAPI.RollbackImplementation(context.Background(), projectId, implementationId).RollbackImplementationRequest(rollbackImplementationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImplementationsAPI.RollbackImplementation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackImplementation`: RollbackImplementation200Response
	fmt.Fprintf(os.Stdout, "Response from `ImplementationsAPI.RollbackImplementation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 
**implementationId** | **string** | Unique identifier for the implementation | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackImplementationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollbackImplementationRequest** | [**RollbackImplementationRequest**](RollbackImplementationRequest.md) |  | 

### Return type

[**RollbackImplementation200Response**](RollbackImplementation200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

