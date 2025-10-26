# ChangeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageUrl** | **string** | URL of the page | 
**ChangeStatus** | **string** | Change status | 
**Visibility** | **string** | Page visibility | 
**PreviousScrapeAt** | Pointer to **time.Time** | Previous scrape timestamp | [optional] 

## Methods

### NewChangeResult

`func NewChangeResult(pageUrl string, changeStatus string, visibility string, ) *ChangeResult`

NewChangeResult instantiates a new ChangeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeResultWithDefaults

`func NewChangeResultWithDefaults() *ChangeResult`

NewChangeResultWithDefaults instantiates a new ChangeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageUrl

`func (o *ChangeResult) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *ChangeResult) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *ChangeResult) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetChangeStatus

`func (o *ChangeResult) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *ChangeResult) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *ChangeResult) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.


### GetVisibility

`func (o *ChangeResult) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ChangeResult) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ChangeResult) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetPreviousScrapeAt

`func (o *ChangeResult) GetPreviousScrapeAt() time.Time`

GetPreviousScrapeAt returns the PreviousScrapeAt field if non-nil, zero value otherwise.

### GetPreviousScrapeAtOk

`func (o *ChangeResult) GetPreviousScrapeAtOk() (*time.Time, bool)`

GetPreviousScrapeAtOk returns a tuple with the PreviousScrapeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousScrapeAt

`func (o *ChangeResult) SetPreviousScrapeAt(v time.Time)`

SetPreviousScrapeAt sets PreviousScrapeAt field to given value.

### HasPreviousScrapeAt

`func (o *ChangeResult) HasPreviousScrapeAt() bool`

HasPreviousScrapeAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


