# PageLimitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Page limit | 
**Used** | **int32** | Pages used | 
**Remaining** | **int32** | Pages remaining | 

## Methods

### NewPageLimitStatus

`func NewPageLimitStatus(limit int32, used int32, remaining int32, ) *PageLimitStatus`

NewPageLimitStatus instantiates a new PageLimitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageLimitStatusWithDefaults

`func NewPageLimitStatusWithDefaults() *PageLimitStatus`

NewPageLimitStatusWithDefaults instantiates a new PageLimitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PageLimitStatus) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageLimitStatus) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageLimitStatus) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetUsed

`func (o *PageLimitStatus) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *PageLimitStatus) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *PageLimitStatus) SetUsed(v int32)`

SetUsed sets Used field to given value.


### GetRemaining

`func (o *PageLimitStatus) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *PageLimitStatus) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *PageLimitStatus) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


