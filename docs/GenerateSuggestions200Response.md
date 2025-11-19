# GenerateSuggestions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Count** | Pointer to **int32** | Number of suggestions generated | [optional] 
**AutoApproved** | Pointer to **int32** | Number of suggestions auto-approved | [optional] 
**Opportunities** | Pointer to [**[]LinkOpportunity**](LinkOpportunity.md) | Top 10 opportunities preview | [optional] 

## Methods

### NewGenerateSuggestions200Response

`func NewGenerateSuggestions200Response() *GenerateSuggestions200Response`

NewGenerateSuggestions200Response instantiates a new GenerateSuggestions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSuggestions200ResponseWithDefaults

`func NewGenerateSuggestions200ResponseWithDefaults() *GenerateSuggestions200Response`

NewGenerateSuggestions200ResponseWithDefaults instantiates a new GenerateSuggestions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GenerateSuggestions200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GenerateSuggestions200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GenerateSuggestions200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GenerateSuggestions200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCount

`func (o *GenerateSuggestions200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GenerateSuggestions200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GenerateSuggestions200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GenerateSuggestions200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAutoApproved

`func (o *GenerateSuggestions200Response) GetAutoApproved() int32`

GetAutoApproved returns the AutoApproved field if non-nil, zero value otherwise.

### GetAutoApprovedOk

`func (o *GenerateSuggestions200Response) GetAutoApprovedOk() (*int32, bool)`

GetAutoApprovedOk returns a tuple with the AutoApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproved

`func (o *GenerateSuggestions200Response) SetAutoApproved(v int32)`

SetAutoApproved sets AutoApproved field to given value.

### HasAutoApproved

`func (o *GenerateSuggestions200Response) HasAutoApproved() bool`

HasAutoApproved returns a boolean if a field has been set.

### GetOpportunities

`func (o *GenerateSuggestions200Response) GetOpportunities() []LinkOpportunity`

GetOpportunities returns the Opportunities field if non-nil, zero value otherwise.

### GetOpportunitiesOk

`func (o *GenerateSuggestions200Response) GetOpportunitiesOk() (*[]LinkOpportunity, bool)`

GetOpportunitiesOk returns a tuple with the Opportunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunities

`func (o *GenerateSuggestions200Response) SetOpportunities(v []LinkOpportunity)`

SetOpportunities sets Opportunities field to given value.

### HasOpportunities

`func (o *GenerateSuggestions200Response) HasOpportunities() bool`

HasOpportunities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


