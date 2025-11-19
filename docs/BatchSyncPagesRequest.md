# BatchSyncPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | [**[]PageData**](PageData.md) |  | 

## Methods

### NewBatchSyncPagesRequest

`func NewBatchSyncPagesRequest(pages []PageData, ) *BatchSyncPagesRequest`

NewBatchSyncPagesRequest instantiates a new BatchSyncPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSyncPagesRequestWithDefaults

`func NewBatchSyncPagesRequestWithDefaults() *BatchSyncPagesRequest`

NewBatchSyncPagesRequestWithDefaults instantiates a new BatchSyncPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *BatchSyncPagesRequest) GetPages() []PageData`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *BatchSyncPagesRequest) GetPagesOk() (*[]PageData, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *BatchSyncPagesRequest) SetPages(v []PageData)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


