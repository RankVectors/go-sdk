# PageAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowance** | Pointer to [**PageAllowanceInfo**](PageAllowanceInfo.md) |  | [optional] 
**UsageHistory** | Pointer to [**[]PageAllowanceUsageHistoryInner**](PageAllowanceUsageHistoryInner.md) |  | [optional] 

## Methods

### NewPageAllowance

`func NewPageAllowance() *PageAllowance`

NewPageAllowance instantiates a new PageAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageAllowanceWithDefaults

`func NewPageAllowanceWithDefaults() *PageAllowance`

NewPageAllowanceWithDefaults instantiates a new PageAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowance

`func (o *PageAllowance) GetAllowance() PageAllowanceInfo`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *PageAllowance) GetAllowanceOk() (*PageAllowanceInfo, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *PageAllowance) SetAllowance(v PageAllowanceInfo)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *PageAllowance) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetUsageHistory

`func (o *PageAllowance) GetUsageHistory() []PageAllowanceUsageHistoryInner`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *PageAllowance) GetUsageHistoryOk() (*[]PageAllowanceUsageHistoryInner, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *PageAllowance) SetUsageHistory(v []PageAllowanceUsageHistoryInner)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *PageAllowance) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


