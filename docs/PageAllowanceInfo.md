# PageAllowanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagesTotal** | **int32** | Total pages allowed (tier + packs) | 
**PagesUsed** | **int32** | Pages currently indexed across all projects | 
**PagesRemaining** | **int32** | Remaining pages before upgrade or page pack | 

## Methods

### NewPageAllowanceInfo

`func NewPageAllowanceInfo(pagesTotal int32, pagesUsed int32, pagesRemaining int32, ) *PageAllowanceInfo`

NewPageAllowanceInfo instantiates a new PageAllowanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageAllowanceInfoWithDefaults

`func NewPageAllowanceInfoWithDefaults() *PageAllowanceInfo`

NewPageAllowanceInfoWithDefaults instantiates a new PageAllowanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagesTotal

`func (o *PageAllowanceInfo) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *PageAllowanceInfo) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *PageAllowanceInfo) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.


### GetPagesUsed

`func (o *PageAllowanceInfo) GetPagesUsed() int32`

GetPagesUsed returns the PagesUsed field if non-nil, zero value otherwise.

### GetPagesUsedOk

`func (o *PageAllowanceInfo) GetPagesUsedOk() (*int32, bool)`

GetPagesUsedOk returns a tuple with the PagesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesUsed

`func (o *PageAllowanceInfo) SetPagesUsed(v int32)`

SetPagesUsed sets PagesUsed field to given value.


### GetPagesRemaining

`func (o *PageAllowanceInfo) GetPagesRemaining() int32`

GetPagesRemaining returns the PagesRemaining field if non-nil, zero value otherwise.

### GetPagesRemainingOk

`func (o *PageAllowanceInfo) GetPagesRemainingOk() (*int32, bool)`

GetPagesRemainingOk returns a tuple with the PagesRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesRemaining

`func (o *PageAllowanceInfo) SetPagesRemaining(v int32)`

SetPagesRemaining sets PagesRemaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


