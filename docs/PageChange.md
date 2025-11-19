# PageChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Change identifier | 
**PageId** | **string** | Page identifier | 
**ProjectId** | **string** | Project identifier | 
**ChangeStatus** | **string** | Change status | 
**Visibility** | **string** | Page visibility | 
**PreviousScrapeAt** | Pointer to **NullableTime** | Previous scrape timestamp | [optional] 
**CurrentScrapeAt** | **time.Time** | Current scrape timestamp | 
**Page** | Pointer to [**PageChangePage**](PageChangePage.md) |  | [optional] 

## Methods

### NewPageChange

`func NewPageChange(id string, pageId string, projectId string, changeStatus string, visibility string, currentScrapeAt time.Time, ) *PageChange`

NewPageChange instantiates a new PageChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageChangeWithDefaults

`func NewPageChangeWithDefaults() *PageChange`

NewPageChangeWithDefaults instantiates a new PageChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageChange) SetId(v string)`

SetId sets Id field to given value.


### GetPageId

`func (o *PageChange) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageChange) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageChange) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetProjectId

`func (o *PageChange) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PageChange) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PageChange) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetChangeStatus

`func (o *PageChange) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *PageChange) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *PageChange) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.


### GetVisibility

`func (o *PageChange) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PageChange) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PageChange) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetPreviousScrapeAt

`func (o *PageChange) GetPreviousScrapeAt() time.Time`

GetPreviousScrapeAt returns the PreviousScrapeAt field if non-nil, zero value otherwise.

### GetPreviousScrapeAtOk

`func (o *PageChange) GetPreviousScrapeAtOk() (*time.Time, bool)`

GetPreviousScrapeAtOk returns a tuple with the PreviousScrapeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousScrapeAt

`func (o *PageChange) SetPreviousScrapeAt(v time.Time)`

SetPreviousScrapeAt sets PreviousScrapeAt field to given value.

### HasPreviousScrapeAt

`func (o *PageChange) HasPreviousScrapeAt() bool`

HasPreviousScrapeAt returns a boolean if a field has been set.

### SetPreviousScrapeAtNil

`func (o *PageChange) SetPreviousScrapeAtNil(b bool)`

 SetPreviousScrapeAtNil sets the value for PreviousScrapeAt to be an explicit nil

### UnsetPreviousScrapeAt
`func (o *PageChange) UnsetPreviousScrapeAt()`

UnsetPreviousScrapeAt ensures that no value is present for PreviousScrapeAt, not even an explicit nil
### GetCurrentScrapeAt

`func (o *PageChange) GetCurrentScrapeAt() time.Time`

GetCurrentScrapeAt returns the CurrentScrapeAt field if non-nil, zero value otherwise.

### GetCurrentScrapeAtOk

`func (o *PageChange) GetCurrentScrapeAtOk() (*time.Time, bool)`

GetCurrentScrapeAtOk returns a tuple with the CurrentScrapeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScrapeAt

`func (o *PageChange) SetCurrentScrapeAt(v time.Time)`

SetCurrentScrapeAt sets CurrentScrapeAt field to given value.


### GetPage

`func (o *PageChange) GetPage() PageChangePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageChange) GetPageOk() (*PageChangePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageChange) SetPage(v PageChangePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageChange) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


