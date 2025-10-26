# ContentVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Safe** | **bool** | Whether it&#39;s safe to implement the link | 
**Reason** | **string** | Reason for the verification result | 
**ChangeResult** | Pointer to [**ChangeResult**](ChangeResult.md) |  | [optional] 

## Methods

### NewContentVerification

`func NewContentVerification(safe bool, reason string, ) *ContentVerification`

NewContentVerification instantiates a new ContentVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentVerificationWithDefaults

`func NewContentVerificationWithDefaults() *ContentVerification`

NewContentVerificationWithDefaults instantiates a new ContentVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafe

`func (o *ContentVerification) GetSafe() bool`

GetSafe returns the Safe field if non-nil, zero value otherwise.

### GetSafeOk

`func (o *ContentVerification) GetSafeOk() (*bool, bool)`

GetSafeOk returns a tuple with the Safe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafe

`func (o *ContentVerification) SetSafe(v bool)`

SetSafe sets Safe field to given value.


### GetReason

`func (o *ContentVerification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContentVerification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContentVerification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetChangeResult

`func (o *ContentVerification) GetChangeResult() ChangeResult`

GetChangeResult returns the ChangeResult field if non-nil, zero value otherwise.

### GetChangeResultOk

`func (o *ContentVerification) GetChangeResultOk() (*ChangeResult, bool)`

GetChangeResultOk returns a tuple with the ChangeResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeResult

`func (o *ContentVerification) SetChangeResult(v ChangeResult)`

SetChangeResult sets ChangeResult field to given value.

### HasChangeResult

`func (o *ContentVerification) HasChangeResult() bool`

HasChangeResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


