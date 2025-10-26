# AddCredits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Balance** | Pointer to [**CreditBalanceInfo**](CreditBalanceInfo.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAddCredits200Response

`func NewAddCredits200Response() *AddCredits200Response`

NewAddCredits200Response instantiates a new AddCredits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredits200ResponseWithDefaults

`func NewAddCredits200ResponseWithDefaults() *AddCredits200Response`

NewAddCredits200ResponseWithDefaults instantiates a new AddCredits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddCredits200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCredits200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCredits200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCredits200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBalance

`func (o *AddCredits200Response) GetBalance() CreditBalanceInfo`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddCredits200Response) GetBalanceOk() (*CreditBalanceInfo, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddCredits200Response) SetBalance(v CreditBalanceInfo)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AddCredits200Response) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetMessage

`func (o *AddCredits200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AddCredits200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AddCredits200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AddCredits200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


