# ProjectAnalyticsSuggestionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Pending** | Pointer to **int32** |  | [optional] 
**Approved** | Pointer to **int32** |  | [optional] 
**Rejected** | Pointer to **int32** |  | [optional] 
**Implemented** | Pointer to **int32** |  | [optional] 
**AcceptanceRate** | Pointer to **float32** |  | [optional] 
**ImplementationRate** | Pointer to **float32** |  | [optional] 
**AvgRelevanceScore** | Pointer to **float32** |  | [optional] 
**RelevanceScoreDistribution** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProjectAnalyticsSuggestionStats

`func NewProjectAnalyticsSuggestionStats() *ProjectAnalyticsSuggestionStats`

NewProjectAnalyticsSuggestionStats instantiates a new ProjectAnalyticsSuggestionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAnalyticsSuggestionStatsWithDefaults

`func NewProjectAnalyticsSuggestionStatsWithDefaults() *ProjectAnalyticsSuggestionStats`

NewProjectAnalyticsSuggestionStatsWithDefaults instantiates a new ProjectAnalyticsSuggestionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ProjectAnalyticsSuggestionStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProjectAnalyticsSuggestionStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProjectAnalyticsSuggestionStats) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProjectAnalyticsSuggestionStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPending

`func (o *ProjectAnalyticsSuggestionStats) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ProjectAnalyticsSuggestionStats) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ProjectAnalyticsSuggestionStats) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ProjectAnalyticsSuggestionStats) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetApproved

`func (o *ProjectAnalyticsSuggestionStats) GetApproved() int32`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ProjectAnalyticsSuggestionStats) GetApprovedOk() (*int32, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ProjectAnalyticsSuggestionStats) SetApproved(v int32)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ProjectAnalyticsSuggestionStats) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetRejected

`func (o *ProjectAnalyticsSuggestionStats) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *ProjectAnalyticsSuggestionStats) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *ProjectAnalyticsSuggestionStats) SetRejected(v int32)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *ProjectAnalyticsSuggestionStats) HasRejected() bool`

HasRejected returns a boolean if a field has been set.

### GetImplemented

`func (o *ProjectAnalyticsSuggestionStats) GetImplemented() int32`

GetImplemented returns the Implemented field if non-nil, zero value otherwise.

### GetImplementedOk

`func (o *ProjectAnalyticsSuggestionStats) GetImplementedOk() (*int32, bool)`

GetImplementedOk returns a tuple with the Implemented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplemented

`func (o *ProjectAnalyticsSuggestionStats) SetImplemented(v int32)`

SetImplemented sets Implemented field to given value.

### HasImplemented

`func (o *ProjectAnalyticsSuggestionStats) HasImplemented() bool`

HasImplemented returns a boolean if a field has been set.

### GetAcceptanceRate

`func (o *ProjectAnalyticsSuggestionStats) GetAcceptanceRate() float32`

GetAcceptanceRate returns the AcceptanceRate field if non-nil, zero value otherwise.

### GetAcceptanceRateOk

`func (o *ProjectAnalyticsSuggestionStats) GetAcceptanceRateOk() (*float32, bool)`

GetAcceptanceRateOk returns a tuple with the AcceptanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceRate

`func (o *ProjectAnalyticsSuggestionStats) SetAcceptanceRate(v float32)`

SetAcceptanceRate sets AcceptanceRate field to given value.

### HasAcceptanceRate

`func (o *ProjectAnalyticsSuggestionStats) HasAcceptanceRate() bool`

HasAcceptanceRate returns a boolean if a field has been set.

### GetImplementationRate

`func (o *ProjectAnalyticsSuggestionStats) GetImplementationRate() float32`

GetImplementationRate returns the ImplementationRate field if non-nil, zero value otherwise.

### GetImplementationRateOk

`func (o *ProjectAnalyticsSuggestionStats) GetImplementationRateOk() (*float32, bool)`

GetImplementationRateOk returns a tuple with the ImplementationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationRate

`func (o *ProjectAnalyticsSuggestionStats) SetImplementationRate(v float32)`

SetImplementationRate sets ImplementationRate field to given value.

### HasImplementationRate

`func (o *ProjectAnalyticsSuggestionStats) HasImplementationRate() bool`

HasImplementationRate returns a boolean if a field has been set.

### GetAvgRelevanceScore

`func (o *ProjectAnalyticsSuggestionStats) GetAvgRelevanceScore() float32`

GetAvgRelevanceScore returns the AvgRelevanceScore field if non-nil, zero value otherwise.

### GetAvgRelevanceScoreOk

`func (o *ProjectAnalyticsSuggestionStats) GetAvgRelevanceScoreOk() (*float32, bool)`

GetAvgRelevanceScoreOk returns a tuple with the AvgRelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRelevanceScore

`func (o *ProjectAnalyticsSuggestionStats) SetAvgRelevanceScore(v float32)`

SetAvgRelevanceScore sets AvgRelevanceScore field to given value.

### HasAvgRelevanceScore

`func (o *ProjectAnalyticsSuggestionStats) HasAvgRelevanceScore() bool`

HasAvgRelevanceScore returns a boolean if a field has been set.

### GetRelevanceScoreDistribution

`func (o *ProjectAnalyticsSuggestionStats) GetRelevanceScoreDistribution() map[string]interface{}`

GetRelevanceScoreDistribution returns the RelevanceScoreDistribution field if non-nil, zero value otherwise.

### GetRelevanceScoreDistributionOk

`func (o *ProjectAnalyticsSuggestionStats) GetRelevanceScoreDistributionOk() (*map[string]interface{}, bool)`

GetRelevanceScoreDistributionOk returns a tuple with the RelevanceScoreDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScoreDistribution

`func (o *ProjectAnalyticsSuggestionStats) SetRelevanceScoreDistribution(v map[string]interface{})`

SetRelevanceScoreDistribution sets RelevanceScoreDistribution field to given value.

### HasRelevanceScoreDistribution

`func (o *ProjectAnalyticsSuggestionStats) HasRelevanceScoreDistribution() bool`

HasRelevanceScoreDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


