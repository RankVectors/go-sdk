# ProjectAnalyticsImplementationStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Success** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**Pending** | Pointer to **int32** |  | [optional] 
**InProgress** | Pointer to **int32** |  | [optional] 
**SuccessRate** | Pointer to **float32** |  | [optional] 
**AvgCreditsUsed** | Pointer to **float32** |  | [optional] 
**PlatformBreakdown** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProjectAnalyticsImplementationStats

`func NewProjectAnalyticsImplementationStats() *ProjectAnalyticsImplementationStats`

NewProjectAnalyticsImplementationStats instantiates a new ProjectAnalyticsImplementationStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAnalyticsImplementationStatsWithDefaults

`func NewProjectAnalyticsImplementationStatsWithDefaults() *ProjectAnalyticsImplementationStats`

NewProjectAnalyticsImplementationStatsWithDefaults instantiates a new ProjectAnalyticsImplementationStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ProjectAnalyticsImplementationStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProjectAnalyticsImplementationStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProjectAnalyticsImplementationStats) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProjectAnalyticsImplementationStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *ProjectAnalyticsImplementationStats) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProjectAnalyticsImplementationStats) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProjectAnalyticsImplementationStats) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ProjectAnalyticsImplementationStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *ProjectAnalyticsImplementationStats) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProjectAnalyticsImplementationStats) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProjectAnalyticsImplementationStats) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProjectAnalyticsImplementationStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *ProjectAnalyticsImplementationStats) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ProjectAnalyticsImplementationStats) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ProjectAnalyticsImplementationStats) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ProjectAnalyticsImplementationStats) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetInProgress

`func (o *ProjectAnalyticsImplementationStats) GetInProgress() int32`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *ProjectAnalyticsImplementationStats) GetInProgressOk() (*int32, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *ProjectAnalyticsImplementationStats) SetInProgress(v int32)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *ProjectAnalyticsImplementationStats) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetSuccessRate

`func (o *ProjectAnalyticsImplementationStats) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *ProjectAnalyticsImplementationStats) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *ProjectAnalyticsImplementationStats) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *ProjectAnalyticsImplementationStats) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetAvgCreditsUsed

`func (o *ProjectAnalyticsImplementationStats) GetAvgCreditsUsed() float32`

GetAvgCreditsUsed returns the AvgCreditsUsed field if non-nil, zero value otherwise.

### GetAvgCreditsUsedOk

`func (o *ProjectAnalyticsImplementationStats) GetAvgCreditsUsedOk() (*float32, bool)`

GetAvgCreditsUsedOk returns a tuple with the AvgCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgCreditsUsed

`func (o *ProjectAnalyticsImplementationStats) SetAvgCreditsUsed(v float32)`

SetAvgCreditsUsed sets AvgCreditsUsed field to given value.

### HasAvgCreditsUsed

`func (o *ProjectAnalyticsImplementationStats) HasAvgCreditsUsed() bool`

HasAvgCreditsUsed returns a boolean if a field has been set.

### GetPlatformBreakdown

`func (o *ProjectAnalyticsImplementationStats) GetPlatformBreakdown() map[string]interface{}`

GetPlatformBreakdown returns the PlatformBreakdown field if non-nil, zero value otherwise.

### GetPlatformBreakdownOk

`func (o *ProjectAnalyticsImplementationStats) GetPlatformBreakdownOk() (*map[string]interface{}, bool)`

GetPlatformBreakdownOk returns a tuple with the PlatformBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBreakdown

`func (o *ProjectAnalyticsImplementationStats) SetPlatformBreakdown(v map[string]interface{})`

SetPlatformBreakdown sets PlatformBreakdown field to given value.

### HasPlatformBreakdown

`func (o *ProjectAnalyticsImplementationStats) HasPlatformBreakdown() bool`

HasPlatformBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


