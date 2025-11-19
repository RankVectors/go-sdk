# GetPageChanges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**[]PageChange**](PageChange.md) |  | [optional] 
**Summary** | Pointer to [**GetPageChanges200ResponseSummary**](GetPageChanges200ResponseSummary.md) |  | [optional] 

## Methods

### NewGetPageChanges200Response

`func NewGetPageChanges200Response() *GetPageChanges200Response`

NewGetPageChanges200Response instantiates a new GetPageChanges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPageChanges200ResponseWithDefaults

`func NewGetPageChanges200ResponseWithDefaults() *GetPageChanges200Response`

NewGetPageChanges200ResponseWithDefaults instantiates a new GetPageChanges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *GetPageChanges200Response) GetChanges() []PageChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *GetPageChanges200Response) GetChangesOk() (*[]PageChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *GetPageChanges200Response) SetChanges(v []PageChange)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *GetPageChanges200Response) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetSummary

`func (o *GetPageChanges200Response) GetSummary() GetPageChanges200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetPageChanges200Response) GetSummaryOk() (*GetPageChanges200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetPageChanges200Response) SetSummary(v GetPageChanges200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetPageChanges200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


