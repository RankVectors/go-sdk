# ProjectAnalyticsSyncStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncDate** | Pointer to **NullableTime** |  | [optional] 
**TotalPagesSynced** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjectAnalyticsSyncStats

`func NewProjectAnalyticsSyncStats() *ProjectAnalyticsSyncStats`

NewProjectAnalyticsSyncStats instantiates a new ProjectAnalyticsSyncStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAnalyticsSyncStatsWithDefaults

`func NewProjectAnalyticsSyncStatsWithDefaults() *ProjectAnalyticsSyncStats`

NewProjectAnalyticsSyncStatsWithDefaults instantiates a new ProjectAnalyticsSyncStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSyncDate

`func (o *ProjectAnalyticsSyncStats) GetLastSyncDate() time.Time`

GetLastSyncDate returns the LastSyncDate field if non-nil, zero value otherwise.

### GetLastSyncDateOk

`func (o *ProjectAnalyticsSyncStats) GetLastSyncDateOk() (*time.Time, bool)`

GetLastSyncDateOk returns a tuple with the LastSyncDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDate

`func (o *ProjectAnalyticsSyncStats) SetLastSyncDate(v time.Time)`

SetLastSyncDate sets LastSyncDate field to given value.

### HasLastSyncDate

`func (o *ProjectAnalyticsSyncStats) HasLastSyncDate() bool`

HasLastSyncDate returns a boolean if a field has been set.

### SetLastSyncDateNil

`func (o *ProjectAnalyticsSyncStats) SetLastSyncDateNil(b bool)`

 SetLastSyncDateNil sets the value for LastSyncDate to be an explicit nil

### UnsetLastSyncDate
`func (o *ProjectAnalyticsSyncStats) UnsetLastSyncDate()`

UnsetLastSyncDate ensures that no value is present for LastSyncDate, not even an explicit nil
### GetTotalPagesSynced

`func (o *ProjectAnalyticsSyncStats) GetTotalPagesSynced() int32`

GetTotalPagesSynced returns the TotalPagesSynced field if non-nil, zero value otherwise.

### GetTotalPagesSyncedOk

`func (o *ProjectAnalyticsSyncStats) GetTotalPagesSyncedOk() (*int32, bool)`

GetTotalPagesSyncedOk returns a tuple with the TotalPagesSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagesSynced

`func (o *ProjectAnalyticsSyncStats) SetTotalPagesSynced(v int32)`

SetTotalPagesSynced sets TotalPagesSynced field to given value.

### HasTotalPagesSynced

`func (o *ProjectAnalyticsSyncStats) HasTotalPagesSynced() bool`

HasTotalPagesSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


