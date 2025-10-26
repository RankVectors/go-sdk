# VerifyContent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Verification** | Pointer to [**ContentVerification**](ContentVerification.md) |  | [optional] 

## Methods

### NewVerifyContent200Response

`func NewVerifyContent200Response() *VerifyContent200Response`

NewVerifyContent200Response instantiates a new VerifyContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyContent200ResponseWithDefaults

`func NewVerifyContent200ResponseWithDefaults() *VerifyContent200Response`

NewVerifyContent200ResponseWithDefaults instantiates a new VerifyContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *VerifyContent200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VerifyContent200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VerifyContent200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *VerifyContent200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetVerification

`func (o *VerifyContent200Response) GetVerification() ContentVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *VerifyContent200Response) GetVerificationOk() (*ContentVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *VerifyContent200Response) SetVerification(v ContentVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *VerifyContent200Response) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


