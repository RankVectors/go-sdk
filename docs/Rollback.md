# Rollback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique rollback identifier | 
**Reason** | **string** | Reason for rollback | 
**CreatedAt** | **time.Time** | Rollback timestamp | 
**CreditsRefunded** | Pointer to **float32** | Credits refunded | [optional] 

## Methods

### NewRollback

`func NewRollback(id string, reason string, createdAt time.Time, ) *Rollback`

NewRollback instantiates a new Rollback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackWithDefaults

`func NewRollbackWithDefaults() *Rollback`

NewRollbackWithDefaults instantiates a new Rollback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rollback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rollback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rollback) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *Rollback) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Rollback) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Rollback) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatedAt

`func (o *Rollback) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Rollback) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Rollback) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreditsRefunded

`func (o *Rollback) GetCreditsRefunded() float32`

GetCreditsRefunded returns the CreditsRefunded field if non-nil, zero value otherwise.

### GetCreditsRefundedOk

`func (o *Rollback) GetCreditsRefundedOk() (*float32, bool)`

GetCreditsRefundedOk returns a tuple with the CreditsRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsRefunded

`func (o *Rollback) SetCreditsRefunded(v float32)`

SetCreditsRefunded sets CreditsRefunded field to given value.

### HasCreditsRefunded

`func (o *Rollback) HasCreditsRefunded() bool`

HasCreditsRefunded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


