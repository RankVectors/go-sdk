# CreditBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Balance** | [**CreditBalanceInfo**](CreditBalanceInfo.md) |  | 
**PageLimit** | [**PageLimitStatus**](PageLimitStatus.md) |  | 
**SpendingLimit** | [**SpendingLimitStatus**](SpendingLimitStatus.md) |  | 
**AutoRecharge** | [**AutoRechargeSettings**](AutoRechargeSettings.md) |  | 
**CurrentPeriod** | [**PeriodCharges**](PeriodCharges.md) |  | 
**UsageHistory** | Pointer to [**[]UsageHistoryItem**](UsageHistoryItem.md) | Usage history (if requested) | [optional] 

## Methods

### NewCreditBalance

`func NewCreditBalance(success bool, balance CreditBalanceInfo, pageLimit PageLimitStatus, spendingLimit SpendingLimitStatus, autoRecharge AutoRechargeSettings, currentPeriod PeriodCharges, ) *CreditBalance`

NewCreditBalance instantiates a new CreditBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceWithDefaults

`func NewCreditBalanceWithDefaults() *CreditBalance`

NewCreditBalanceWithDefaults instantiates a new CreditBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreditBalance) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreditBalance) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreditBalance) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetBalance

`func (o *CreditBalance) GetBalance() CreditBalanceInfo`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreditBalance) GetBalanceOk() (*CreditBalanceInfo, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreditBalance) SetBalance(v CreditBalanceInfo)`

SetBalance sets Balance field to given value.


### GetPageLimit

`func (o *CreditBalance) GetPageLimit() PageLimitStatus`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *CreditBalance) GetPageLimitOk() (*PageLimitStatus, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *CreditBalance) SetPageLimit(v PageLimitStatus)`

SetPageLimit sets PageLimit field to given value.


### GetSpendingLimit

`func (o *CreditBalance) GetSpendingLimit() SpendingLimitStatus`

GetSpendingLimit returns the SpendingLimit field if non-nil, zero value otherwise.

### GetSpendingLimitOk

`func (o *CreditBalance) GetSpendingLimitOk() (*SpendingLimitStatus, bool)`

GetSpendingLimitOk returns a tuple with the SpendingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimit

`func (o *CreditBalance) SetSpendingLimit(v SpendingLimitStatus)`

SetSpendingLimit sets SpendingLimit field to given value.


### GetAutoRecharge

`func (o *CreditBalance) GetAutoRecharge() AutoRechargeSettings`

GetAutoRecharge returns the AutoRecharge field if non-nil, zero value otherwise.

### GetAutoRechargeOk

`func (o *CreditBalance) GetAutoRechargeOk() (*AutoRechargeSettings, bool)`

GetAutoRechargeOk returns a tuple with the AutoRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecharge

`func (o *CreditBalance) SetAutoRecharge(v AutoRechargeSettings)`

SetAutoRecharge sets AutoRecharge field to given value.


### GetCurrentPeriod

`func (o *CreditBalance) GetCurrentPeriod() PeriodCharges`

GetCurrentPeriod returns the CurrentPeriod field if non-nil, zero value otherwise.

### GetCurrentPeriodOk

`func (o *CreditBalance) GetCurrentPeriodOk() (*PeriodCharges, bool)`

GetCurrentPeriodOk returns a tuple with the CurrentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriod

`func (o *CreditBalance) SetCurrentPeriod(v PeriodCharges)`

SetCurrentPeriod sets CurrentPeriod field to given value.


### GetUsageHistory

`func (o *CreditBalance) GetUsageHistory() []UsageHistoryItem`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *CreditBalance) GetUsageHistoryOk() (*[]UsageHistoryItem, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *CreditBalance) SetUsageHistory(v []UsageHistoryItem)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *CreditBalance) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


