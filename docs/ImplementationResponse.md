# ImplementationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether any implementations succeeded | 
**Results** | [**[]ImplementationResult**](ImplementationResult.md) | Individual implementation results | 
**Summary** | [**ImplementationSummary**](ImplementationSummary.md) |  | 

## Methods

### NewImplementationResponse

`func NewImplementationResponse(success bool, results []ImplementationResult, summary ImplementationSummary, ) *ImplementationResponse`

NewImplementationResponse instantiates a new ImplementationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplementationResponseWithDefaults

`func NewImplementationResponseWithDefaults() *ImplementationResponse`

NewImplementationResponseWithDefaults instantiates a new ImplementationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ImplementationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ImplementationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ImplementationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetResults

`func (o *ImplementationResponse) GetResults() []ImplementationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ImplementationResponse) GetResultsOk() (*[]ImplementationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ImplementationResponse) SetResults(v []ImplementationResult)`

SetResults sets Results field to given value.


### GetSummary

`func (o *ImplementationResponse) GetSummary() ImplementationSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ImplementationResponse) GetSummaryOk() (*ImplementationSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ImplementationResponse) SetSummary(v ImplementationSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


