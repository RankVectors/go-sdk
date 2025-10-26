# CreditBalanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Project identifier | 
**CreditsTotal** | **float32** | Total credits purchased | 
**CreditsUsed** | **float32** | Credits consumed | 
**CreditsRemaining** | **float32** | Credits remaining | 
**LastResetAt** | Pointer to **time.Time** | Last credit reset timestamp | [optional] 

## Methods

### NewCreditBalanceInfo

`func NewCreditBalanceInfo(projectId string, creditsTotal float32, creditsUsed float32, creditsRemaining float32, ) *CreditBalanceInfo`

NewCreditBalanceInfo instantiates a new CreditBalanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceInfoWithDefaults

`func NewCreditBalanceInfoWithDefaults() *CreditBalanceInfo`

NewCreditBalanceInfoWithDefaults instantiates a new CreditBalanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreditBalanceInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreditBalanceInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreditBalanceInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreditsTotal

`func (o *CreditBalanceInfo) GetCreditsTotal() float32`

GetCreditsTotal returns the CreditsTotal field if non-nil, zero value otherwise.

### GetCreditsTotalOk

`func (o *CreditBalanceInfo) GetCreditsTotalOk() (*float32, bool)`

GetCreditsTotalOk returns a tuple with the CreditsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsTotal

`func (o *CreditBalanceInfo) SetCreditsTotal(v float32)`

SetCreditsTotal sets CreditsTotal field to given value.


### GetCreditsUsed

`func (o *CreditBalanceInfo) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *CreditBalanceInfo) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *CreditBalanceInfo) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.


### GetCreditsRemaining

`func (o *CreditBalanceInfo) GetCreditsRemaining() float32`

GetCreditsRemaining returns the CreditsRemaining field if non-nil, zero value otherwise.

### GetCreditsRemainingOk

`func (o *CreditBalanceInfo) GetCreditsRemainingOk() (*float32, bool)`

GetCreditsRemainingOk returns a tuple with the CreditsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsRemaining

`func (o *CreditBalanceInfo) SetCreditsRemaining(v float32)`

SetCreditsRemaining sets CreditsRemaining field to given value.


### GetLastResetAt

`func (o *CreditBalanceInfo) GetLastResetAt() time.Time`

GetLastResetAt returns the LastResetAt field if non-nil, zero value otherwise.

### GetLastResetAtOk

`func (o *CreditBalanceInfo) GetLastResetAtOk() (*time.Time, bool)`

GetLastResetAtOk returns a tuple with the LastResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResetAt

`func (o *CreditBalanceInfo) SetLastResetAt(v time.Time)`

SetLastResetAt sets LastResetAt field to given value.

### HasLastResetAt

`func (o *CreditBalanceInfo) HasLastResetAt() bool`

HasLastResetAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


