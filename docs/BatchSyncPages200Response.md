# BatchSyncPages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Synced** | Pointer to **int32** | Number of pages synced | [optional] 
**Skipped** | Pointer to **int32** | Number of pages skipped | [optional] 

## Methods

### NewBatchSyncPages200Response

`func NewBatchSyncPages200Response() *BatchSyncPages200Response`

NewBatchSyncPages200Response instantiates a new BatchSyncPages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSyncPages200ResponseWithDefaults

`func NewBatchSyncPages200ResponseWithDefaults() *BatchSyncPages200Response`

NewBatchSyncPages200ResponseWithDefaults instantiates a new BatchSyncPages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *BatchSyncPages200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BatchSyncPages200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BatchSyncPages200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BatchSyncPages200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSynced

`func (o *BatchSyncPages200Response) GetSynced() int32`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *BatchSyncPages200Response) GetSyncedOk() (*int32, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *BatchSyncPages200Response) SetSynced(v int32)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *BatchSyncPages200Response) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetSkipped

`func (o *BatchSyncPages200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BatchSyncPages200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BatchSyncPages200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *BatchSyncPages200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


