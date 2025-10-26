# AutoRechargeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether auto-recharge is enabled | 
**Threshold** | **float32** | Credit threshold for auto-recharge | 
**Amount** | **float32** | Amount to recharge | 

## Methods

### NewAutoRechargeSettings

`func NewAutoRechargeSettings(enabled bool, threshold float32, amount float32, ) *AutoRechargeSettings`

NewAutoRechargeSettings instantiates a new AutoRechargeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRechargeSettingsWithDefaults

`func NewAutoRechargeSettingsWithDefaults() *AutoRechargeSettings`

NewAutoRechargeSettingsWithDefaults instantiates a new AutoRechargeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoRechargeSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRechargeSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRechargeSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetThreshold

`func (o *AutoRechargeSettings) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AutoRechargeSettings) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AutoRechargeSettings) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.


### GetAmount

`func (o *AutoRechargeSettings) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AutoRechargeSettings) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AutoRechargeSettings) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


