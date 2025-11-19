# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Page identifier | 
**ProjectId** | **string** | Project identifier | 
**Url** | **string** | Page URL | 
**Title** | Pointer to **NullableString** | Page title | [optional] 
**WordCount** | Pointer to **NullableInt32** | Word count | [optional] 
**InternalLinks** | Pointer to **[]map[string]interface{}** | Internal links on the page | [optional] 
**ExternalLinks** | Pointer to **[]map[string]interface{}** | External links on the page | [optional] 
**SyncedAt** | Pointer to **NullableTime** | Last sync timestamp | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | Last update timestamp | [optional] 

## Methods

### NewPage

`func NewPage(id string, projectId string, url string, ) *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Page) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Page) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Page) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Page) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUrl

`func (o *Page) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Page) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Page) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *Page) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Page) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Page) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Page) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Page) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Page) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWordCount

`func (o *Page) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *Page) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *Page) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *Page) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.

### SetWordCountNil

`func (o *Page) SetWordCountNil(b bool)`

 SetWordCountNil sets the value for WordCount to be an explicit nil

### UnsetWordCount
`func (o *Page) UnsetWordCount()`

UnsetWordCount ensures that no value is present for WordCount, not even an explicit nil
### GetInternalLinks

`func (o *Page) GetInternalLinks() []map[string]interface{}`

GetInternalLinks returns the InternalLinks field if non-nil, zero value otherwise.

### GetInternalLinksOk

`func (o *Page) GetInternalLinksOk() (*[]map[string]interface{}, bool)`

GetInternalLinksOk returns a tuple with the InternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalLinks

`func (o *Page) SetInternalLinks(v []map[string]interface{})`

SetInternalLinks sets InternalLinks field to given value.

### HasInternalLinks

`func (o *Page) HasInternalLinks() bool`

HasInternalLinks returns a boolean if a field has been set.

### GetExternalLinks

`func (o *Page) GetExternalLinks() []map[string]interface{}`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *Page) GetExternalLinksOk() (*[]map[string]interface{}, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *Page) SetExternalLinks(v []map[string]interface{})`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *Page) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetSyncedAt

`func (o *Page) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *Page) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *Page) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *Page) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *Page) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *Page) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Page) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Page) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Page) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Page) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Page) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Page) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


