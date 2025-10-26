# ReportImplementationStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestionId** | **string** | ID of the suggestion | 
**ApiKey** | **string** | Your RankVectors API key | 
**Status** | **string** | Implementation status | 

## Methods

### NewReportImplementationStatusRequest

`func NewReportImplementationStatusRequest(suggestionId string, apiKey string, status string, ) *ReportImplementationStatusRequest`

NewReportImplementationStatusRequest instantiates a new ReportImplementationStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportImplementationStatusRequestWithDefaults

`func NewReportImplementationStatusRequestWithDefaults() *ReportImplementationStatusRequest`

NewReportImplementationStatusRequestWithDefaults instantiates a new ReportImplementationStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestionId

`func (o *ReportImplementationStatusRequest) GetSuggestionId() string`

GetSuggestionId returns the SuggestionId field if non-nil, zero value otherwise.

### GetSuggestionIdOk

`func (o *ReportImplementationStatusRequest) GetSuggestionIdOk() (*string, bool)`

GetSuggestionIdOk returns a tuple with the SuggestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionId

`func (o *ReportImplementationStatusRequest) SetSuggestionId(v string)`

SetSuggestionId sets SuggestionId field to given value.


### GetApiKey

`func (o *ReportImplementationStatusRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ReportImplementationStatusRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ReportImplementationStatusRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetStatus

`func (o *ReportImplementationStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportImplementationStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportImplementationStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


