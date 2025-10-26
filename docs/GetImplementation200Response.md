# GetImplementation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Implementation** | Pointer to [**Implementation**](Implementation.md) |  | [optional] 

## Methods

### NewGetImplementation200Response

`func NewGetImplementation200Response() *GetImplementation200Response`

NewGetImplementation200Response instantiates a new GetImplementation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImplementation200ResponseWithDefaults

`func NewGetImplementation200ResponseWithDefaults() *GetImplementation200Response`

NewGetImplementation200ResponseWithDefaults instantiates a new GetImplementation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetImplementation200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetImplementation200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetImplementation200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetImplementation200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetImplementation

`func (o *GetImplementation200Response) GetImplementation() Implementation`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *GetImplementation200Response) GetImplementationOk() (*Implementation, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *GetImplementation200Response) SetImplementation(v Implementation)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *GetImplementation200Response) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


