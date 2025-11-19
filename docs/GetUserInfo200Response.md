# GetUserInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**GetUserInfo200ResponseUser**](GetUserInfo200ResponseUser.md) |  | [optional] 
**Links** | Pointer to [**GetUserInfo200ResponseLinks**](GetUserInfo200ResponseLinks.md) |  | [optional] 
**ProjectsCount** | Pointer to **int32** |  | [optional] 
**RecentUsage** | Pointer to **float32** | Recent usage (pages indexed in last 30 days) | [optional] 

## Methods

### NewGetUserInfo200Response

`func NewGetUserInfo200Response() *GetUserInfo200Response`

NewGetUserInfo200Response instantiates a new GetUserInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInfo200ResponseWithDefaults

`func NewGetUserInfo200ResponseWithDefaults() *GetUserInfo200Response`

NewGetUserInfo200ResponseWithDefaults instantiates a new GetUserInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetUserInfo200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetUserInfo200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetUserInfo200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetUserInfo200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUser

`func (o *GetUserInfo200Response) GetUser() GetUserInfo200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetUserInfo200Response) GetUserOk() (*GetUserInfo200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetUserInfo200Response) SetUser(v GetUserInfo200ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetUserInfo200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserInfo200Response) GetLinks() GetUserInfo200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserInfo200Response) GetLinksOk() (*GetUserInfo200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserInfo200Response) SetLinks(v GetUserInfo200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserInfo200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProjectsCount

`func (o *GetUserInfo200Response) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *GetUserInfo200Response) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *GetUserInfo200Response) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.

### HasProjectsCount

`func (o *GetUserInfo200Response) HasProjectsCount() bool`

HasProjectsCount returns a boolean if a field has been set.

### GetRecentUsage

`func (o *GetUserInfo200Response) GetRecentUsage() float32`

GetRecentUsage returns the RecentUsage field if non-nil, zero value otherwise.

### GetRecentUsageOk

`func (o *GetUserInfo200Response) GetRecentUsageOk() (*float32, bool)`

GetRecentUsageOk returns a tuple with the RecentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentUsage

`func (o *GetUserInfo200Response) SetRecentUsage(v float32)`

SetRecentUsage sets RecentUsage field to given value.

### HasRecentUsage

`func (o *GetUserInfo200Response) HasRecentUsage() bool`

HasRecentUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


