# VerifyContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageUrl** | **string** | URL of the page to verify | 
**SuggestionId** | **string** | ID of the suggestion to verify | 

## Methods

### NewVerifyContentRequest

`func NewVerifyContentRequest(pageUrl string, suggestionId string, ) *VerifyContentRequest`

NewVerifyContentRequest instantiates a new VerifyContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyContentRequestWithDefaults

`func NewVerifyContentRequestWithDefaults() *VerifyContentRequest`

NewVerifyContentRequestWithDefaults instantiates a new VerifyContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageUrl

`func (o *VerifyContentRequest) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *VerifyContentRequest) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *VerifyContentRequest) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetSuggestionId

`func (o *VerifyContentRequest) GetSuggestionId() string`

GetSuggestionId returns the SuggestionId field if non-nil, zero value otherwise.

### GetSuggestionIdOk

`func (o *VerifyContentRequest) GetSuggestionIdOk() (*string, bool)`

GetSuggestionIdOk returns a tuple with the SuggestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionId

`func (o *VerifyContentRequest) SetSuggestionId(v string)`

SetSuggestionId sets SuggestionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


