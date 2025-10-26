# ListImplementations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Implementations** | Pointer to [**[]Implementation**](Implementation.md) |  | [optional] 
**Count** | Pointer to **int32** | Number of implementations returned | [optional] 
**Filters** | Pointer to **map[string]interface{}** | Applied filters | [optional] 

## Methods

### NewListImplementations200Response

`func NewListImplementations200Response() *ListImplementations200Response`

NewListImplementations200Response instantiates a new ListImplementations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImplementations200ResponseWithDefaults

`func NewListImplementations200ResponseWithDefaults() *ListImplementations200Response`

NewListImplementations200ResponseWithDefaults instantiates a new ListImplementations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListImplementations200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListImplementations200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListImplementations200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListImplementations200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetImplementations

`func (o *ListImplementations200Response) GetImplementations() []Implementation`

GetImplementations returns the Implementations field if non-nil, zero value otherwise.

### GetImplementationsOk

`func (o *ListImplementations200Response) GetImplementationsOk() (*[]Implementation, bool)`

GetImplementationsOk returns a tuple with the Implementations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementations

`func (o *ListImplementations200Response) SetImplementations(v []Implementation)`

SetImplementations sets Implementations field to given value.

### HasImplementations

`func (o *ListImplementations200Response) HasImplementations() bool`

HasImplementations returns a boolean if a field has been set.

### GetCount

`func (o *ListImplementations200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListImplementations200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListImplementations200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListImplementations200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFilters

`func (o *ListImplementations200Response) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListImplementations200Response) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListImplementations200Response) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListImplementations200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


