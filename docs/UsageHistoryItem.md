# UsageHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Usage item identifier | 
**Action** | **string** | Action performed | 
**CreditsUsed** | **float32** | Credits used | 
**CreatedAt** | **time.Time** | Usage timestamp | 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata | [optional] 

## Methods

### NewUsageHistoryItem

`func NewUsageHistoryItem(id string, action string, creditsUsed float32, createdAt time.Time, ) *UsageHistoryItem`

NewUsageHistoryItem instantiates a new UsageHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageHistoryItemWithDefaults

`func NewUsageHistoryItemWithDefaults() *UsageHistoryItem`

NewUsageHistoryItemWithDefaults instantiates a new UsageHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsageHistoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageHistoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageHistoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetAction

`func (o *UsageHistoryItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UsageHistoryItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UsageHistoryItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetCreditsUsed

`func (o *UsageHistoryItem) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *UsageHistoryItem) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *UsageHistoryItem) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.


### GetCreatedAt

`func (o *UsageHistoryItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsageHistoryItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsageHistoryItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetadata

`func (o *UsageHistoryItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageHistoryItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageHistoryItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageHistoryItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


