# \SuggestionsAPI

All URIs are relative to *https://api.rankvectors.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSuggestion**](SuggestionsAPI.md#DeleteSuggestion) | **Delete** /api/projects/{projectId}/suggestions/{suggestionId} | Delete suggestion
[**GenerateSuggestions**](SuggestionsAPI.md#GenerateSuggestions) | **Post** /api/projects/{projectId}/suggestions | Generate link suggestions
[**GetSuggestions**](SuggestionsAPI.md#GetSuggestions) | **Get** /api/projects/{projectId}/suggestions | Get link suggestions
[**UpdateSuggestion**](SuggestionsAPI.md#UpdateSuggestion) | **Patch** /api/projects/{projectId}/suggestions/{suggestionId} | Update suggestion status



## DeleteSuggestion

> DeleteSuggestion200Response DeleteSuggestion(ctx, projectId, suggestionId).Execute()

Delete suggestion



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
	suggestionId := "sugg-123" // string | Unique identifier for the suggestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAPI.DeleteSuggestion(context.Background(), projectId, suggestionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAPI.DeleteSuggestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSuggestion`: DeleteSuggestion200Response
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAPI.DeleteSuggestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 
**suggestionId** | **string** | Unique identifier for the suggestion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteSuggestion200Response**](DeleteSuggestion200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSuggestions

> GenerateSuggestions200Response GenerateSuggestions(ctx, projectId).GenerateSuggestionsRequest(generateSuggestionsRequest).Execute()

Generate link suggestions



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
	generateSuggestionsRequest := *openapiclient.NewGenerateSuggestionsRequest() // GenerateSuggestionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAPI.GenerateSuggestions(context.Background(), projectId).GenerateSuggestionsRequest(generateSuggestionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAPI.GenerateSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSuggestions`: GenerateSuggestions200Response
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAPI.GenerateSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateSuggestionsRequest** | [**GenerateSuggestionsRequest**](GenerateSuggestionsRequest.md) |  | 

### Return type

[**GenerateSuggestions200Response**](GenerateSuggestions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestions

> []Suggestion GetSuggestions(ctx, projectId).Status(status).Execute()

Get link suggestions



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
	status := "pending" // string | Filter suggestions by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAPI.GetSuggestions(context.Background(), projectId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAPI.GetSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestions`: []Suggestion
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAPI.GetSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter suggestions by status | 

### Return type

[**[]Suggestion**](Suggestion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuggestion

> Suggestion UpdateSuggestion(ctx, projectId, suggestionId).UpdateSuggestionRequest(updateSuggestionRequest).Execute()

Update suggestion status



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
	suggestionId := "sugg-123" // string | Unique identifier for the suggestion
	updateSuggestionRequest := *openapiclient.NewUpdateSuggestionRequest("approved") // UpdateSuggestionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAPI.UpdateSuggestion(context.Background(), projectId, suggestionId).UpdateSuggestionRequest(updateSuggestionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAPI.UpdateSuggestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSuggestion`: Suggestion
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAPI.UpdateSuggestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique identifier for the project | 
**suggestionId** | **string** | Unique identifier for the suggestion | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSuggestionRequest** | [**UpdateSuggestionRequest**](UpdateSuggestionRequest.md) |  | 

### Return type

[**Suggestion**](Suggestion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

