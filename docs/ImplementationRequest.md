# ImplementationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestionIds** | **[]string** | Array of suggestion IDs to implement | 
**Platform** | **string** | Platform type | 
**Credentials** | [**CustomCredentials**](CustomCredentials.md) |  | 
**SkipContentVerification** | Pointer to **bool** | Skip content change detection | [optional] [default to false]
**ImplementationMethod** | Pointer to **string** | Implementation method | [optional] [default to "api"]

## Methods

### NewImplementationRequest

`func NewImplementationRequest(suggestionIds []string, platform string, credentials CustomCredentials, ) *ImplementationRequest`

NewImplementationRequest instantiates a new ImplementationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplementationRequestWithDefaults

`func NewImplementationRequestWithDefaults() *ImplementationRequest`

NewImplementationRequestWithDefaults instantiates a new ImplementationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestionIds

`func (o *ImplementationRequest) GetSuggestionIds() []string`

GetSuggestionIds returns the SuggestionIds field if non-nil, zero value otherwise.

### GetSuggestionIdsOk

`func (o *ImplementationRequest) GetSuggestionIdsOk() (*[]string, bool)`

GetSuggestionIdsOk returns a tuple with the SuggestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionIds

`func (o *ImplementationRequest) SetSuggestionIds(v []string)`

SetSuggestionIds sets SuggestionIds field to given value.


### GetPlatform

`func (o *ImplementationRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ImplementationRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ImplementationRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCredentials

`func (o *ImplementationRequest) GetCredentials() CustomCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ImplementationRequest) GetCredentialsOk() (*CustomCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ImplementationRequest) SetCredentials(v CustomCredentials)`

SetCredentials sets Credentials field to given value.


### GetSkipContentVerification

`func (o *ImplementationRequest) GetSkipContentVerification() bool`

GetSkipContentVerification returns the SkipContentVerification field if non-nil, zero value otherwise.

### GetSkipContentVerificationOk

`func (o *ImplementationRequest) GetSkipContentVerificationOk() (*bool, bool)`

GetSkipContentVerificationOk returns a tuple with the SkipContentVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipContentVerification

`func (o *ImplementationRequest) SetSkipContentVerification(v bool)`

SetSkipContentVerification sets SkipContentVerification field to given value.

### HasSkipContentVerification

`func (o *ImplementationRequest) HasSkipContentVerification() bool`

HasSkipContentVerification returns a boolean if a field has been set.

### GetImplementationMethod

`func (o *ImplementationRequest) GetImplementationMethod() string`

GetImplementationMethod returns the ImplementationMethod field if non-nil, zero value otherwise.

### GetImplementationMethodOk

`func (o *ImplementationRequest) GetImplementationMethodOk() (*string, bool)`

GetImplementationMethodOk returns a tuple with the ImplementationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationMethod

`func (o *ImplementationRequest) SetImplementationMethod(v string)`

SetImplementationMethod sets ImplementationMethod field to given value.

### HasImplementationMethod

`func (o *ImplementationRequest) HasImplementationMethod() bool`

HasImplementationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


