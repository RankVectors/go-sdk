# RollbackImplementationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | Reason for rollback | 
**Credentials** | Pointer to [**CustomCredentials**](CustomCredentials.md) |  | [optional] 

## Methods

### NewRollbackImplementationRequest

`func NewRollbackImplementationRequest(reason string, ) *RollbackImplementationRequest`

NewRollbackImplementationRequest instantiates a new RollbackImplementationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackImplementationRequestWithDefaults

`func NewRollbackImplementationRequestWithDefaults() *RollbackImplementationRequest`

NewRollbackImplementationRequestWithDefaults instantiates a new RollbackImplementationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *RollbackImplementationRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RollbackImplementationRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RollbackImplementationRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCredentials

`func (o *RollbackImplementationRequest) GetCredentials() CustomCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RollbackImplementationRequest) GetCredentialsOk() (*CustomCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RollbackImplementationRequest) SetCredentials(v CustomCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RollbackImplementationRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


