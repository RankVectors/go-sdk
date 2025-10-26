# Crawl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique crawl identifier | 
**ProjectId** | **string** | Project identifier | 
**Status** | **string** | Crawl status | 
**StartedAt** | **time.Time** | Crawl start timestamp | 
**CompletedAt** | Pointer to **time.Time** | Crawl completion timestamp | [optional] 
**PagesCrawled** | Pointer to **int32** | Number of pages crawled | [optional] 
**ErrorMessage** | Pointer to **string** | Error message if crawl failed | [optional] 

## Methods

### NewCrawl

`func NewCrawl(id string, projectId string, status string, startedAt time.Time, ) *Crawl`

NewCrawl instantiates a new Crawl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlWithDefaults

`func NewCrawlWithDefaults() *Crawl`

NewCrawlWithDefaults instantiates a new Crawl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Crawl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Crawl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Crawl) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Crawl) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Crawl) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Crawl) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStatus

`func (o *Crawl) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Crawl) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Crawl) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *Crawl) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Crawl) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Crawl) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *Crawl) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Crawl) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Crawl) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Crawl) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetPagesCrawled

`func (o *Crawl) GetPagesCrawled() int32`

GetPagesCrawled returns the PagesCrawled field if non-nil, zero value otherwise.

### GetPagesCrawledOk

`func (o *Crawl) GetPagesCrawledOk() (*int32, bool)`

GetPagesCrawledOk returns a tuple with the PagesCrawled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCrawled

`func (o *Crawl) SetPagesCrawled(v int32)`

SetPagesCrawled sets PagesCrawled field to given value.

### HasPagesCrawled

`func (o *Crawl) HasPagesCrawled() bool`

HasPagesCrawled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Crawl) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Crawl) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Crawl) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Crawl) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


