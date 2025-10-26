# AddCreditsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Number of credits to add | 
**Source** | Pointer to **string** | Source of credits | [optional] [default to "manual_addition"]

## Methods

### NewAddCreditsRequest

`func NewAddCreditsRequest(amount float32, ) *AddCreditsRequest`

NewAddCreditsRequest instantiates a new AddCreditsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCreditsRequestWithDefaults

`func NewAddCreditsRequestWithDefaults() *AddCreditsRequest`

NewAddCreditsRequestWithDefaults instantiates a new AddCreditsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddCreditsRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddCreditsRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddCreditsRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetSource

`func (o *AddCreditsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddCreditsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddCreditsRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddCreditsRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


