# SpendingLimitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether spending limit is enabled | 
**Limit** | **float32** | Spending limit | 
**Used** | **float32** | Amount used | 
**Remaining** | **float32** | Amount remaining | 

## Methods

### NewSpendingLimitStatus

`func NewSpendingLimitStatus(enabled bool, limit float32, used float32, remaining float32, ) *SpendingLimitStatus`

NewSpendingLimitStatus instantiates a new SpendingLimitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingLimitStatusWithDefaults

`func NewSpendingLimitStatusWithDefaults() *SpendingLimitStatus`

NewSpendingLimitStatusWithDefaults instantiates a new SpendingLimitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SpendingLimitStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpendingLimitStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpendingLimitStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLimit

`func (o *SpendingLimitStatus) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SpendingLimitStatus) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SpendingLimitStatus) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetUsed

`func (o *SpendingLimitStatus) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *SpendingLimitStatus) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *SpendingLimitStatus) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetRemaining

`func (o *SpendingLimitStatus) GetRemaining() float32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *SpendingLimitStatus) GetRemainingOk() (*float32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *SpendingLimitStatus) SetRemaining(v float32)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


