# SyncProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain to sync | [optional] 
**ProjectName** | Pointer to **string** | Project name | [optional] 

## Methods

### NewSyncProjectRequest

`func NewSyncProjectRequest() *SyncProjectRequest`

NewSyncProjectRequest instantiates a new SyncProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncProjectRequestWithDefaults

`func NewSyncProjectRequestWithDefaults() *SyncProjectRequest`

NewSyncProjectRequestWithDefaults instantiates a new SyncProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SyncProjectRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SyncProjectRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SyncProjectRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SyncProjectRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetProjectName

`func (o *SyncProjectRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *SyncProjectRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *SyncProjectRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *SyncProjectRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


