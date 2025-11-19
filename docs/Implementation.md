# Implementation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique implementation identifier | 
**Status** | **string** | Implementation status | 
**Platform** | **string** | Platform used | 
**ImplementationMethod** | Pointer to **string** | Implementation method | [optional] 
**CreditsUsed** | Pointer to **float32** | Credits consumed | [optional] 
**CreatedAt** | **time.Time** | Implementation start timestamp | 
**CompletedAt** | Pointer to **time.Time** | Implementation completion timestamp | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Platform-specific metadata | [optional] 
**Suggestion** | [**Suggestion**](Suggestion.md) |  | 
**Rollbacks** | Pointer to [**[]Rollback**](Rollback.md) | Rollback history | [optional] 

## Methods

### NewImplementation

`func NewImplementation(id string, status string, platform string, createdAt time.Time, suggestion Suggestion, ) *Implementation`

NewImplementation instantiates a new Implementation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplementationWithDefaults

`func NewImplementationWithDefaults() *Implementation`

NewImplementationWithDefaults instantiates a new Implementation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Implementation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Implementation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Implementation) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Implementation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Implementation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Implementation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *Implementation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Implementation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Implementation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetImplementationMethod

`func (o *Implementation) GetImplementationMethod() string`

GetImplementationMethod returns the ImplementationMethod field if non-nil, zero value otherwise.

### GetImplementationMethodOk

`func (o *Implementation) GetImplementationMethodOk() (*string, bool)`

GetImplementationMethodOk returns a tuple with the ImplementationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationMethod

`func (o *Implementation) SetImplementationMethod(v string)`

SetImplementationMethod sets ImplementationMethod field to given value.

### HasImplementationMethod

`func (o *Implementation) HasImplementationMethod() bool`

HasImplementationMethod returns a boolean if a field has been set.

### GetCreditsUsed

`func (o *Implementation) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *Implementation) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *Implementation) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.

### HasCreditsUsed

`func (o *Implementation) HasCreditsUsed() bool`

HasCreditsUsed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Implementation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Implementation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Implementation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *Implementation) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Implementation) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Implementation) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Implementation) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *Implementation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Implementation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Implementation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Implementation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSuggestion

`func (o *Implementation) GetSuggestion() Suggestion`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Implementation) GetSuggestionOk() (*Suggestion, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Implementation) SetSuggestion(v Suggestion)`

SetSuggestion sets Suggestion field to given value.


### GetRollbacks

`func (o *Implementation) GetRollbacks() []Rollback`

GetRollbacks returns the Rollbacks field if non-nil, zero value otherwise.

### GetRollbacksOk

`func (o *Implementation) GetRollbacksOk() (*[]Rollback, bool)`

GetRollbacksOk returns a tuple with the Rollbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbacks

`func (o *Implementation) SetRollbacks(v []Rollback)`

SetRollbacks sets Rollbacks field to given value.

### HasRollbacks

`func (o *Implementation) HasRollbacks() bool`

HasRollbacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


