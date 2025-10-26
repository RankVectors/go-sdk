# ImplementationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether implementation succeeded | 
**ImplementationId** | Pointer to **string** | Implementation identifier | [optional] 
**CreditsUsed** | Pointer to **float32** | Credits consumed | [optional] 
**Error** | Pointer to **string** | Error message if failed | [optional] 

## Methods

### NewImplementationResult

`func NewImplementationResult(success bool, ) *ImplementationResult`

NewImplementationResult instantiates a new ImplementationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplementationResultWithDefaults

`func NewImplementationResultWithDefaults() *ImplementationResult`

NewImplementationResultWithDefaults instantiates a new ImplementationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ImplementationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ImplementationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ImplementationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetImplementationId

`func (o *ImplementationResult) GetImplementationId() string`

GetImplementationId returns the ImplementationId field if non-nil, zero value otherwise.

### GetImplementationIdOk

`func (o *ImplementationResult) GetImplementationIdOk() (*string, bool)`

GetImplementationIdOk returns a tuple with the ImplementationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationId

`func (o *ImplementationResult) SetImplementationId(v string)`

SetImplementationId sets ImplementationId field to given value.

### HasImplementationId

`func (o *ImplementationResult) HasImplementationId() bool`

HasImplementationId returns a boolean if a field has been set.

### GetCreditsUsed

`func (o *ImplementationResult) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *ImplementationResult) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *ImplementationResult) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.

### HasCreditsUsed

`func (o *ImplementationResult) HasCreditsUsed() bool`

HasCreditsUsed returns a boolean if a field has been set.

### GetError

`func (o *ImplementationResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImplementationResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImplementationResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImplementationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


