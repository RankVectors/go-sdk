# ProjectAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImplementationStats** | [**ProjectAnalyticsImplementationStats**](ProjectAnalyticsImplementationStats.md) |  | 
**SuggestionStats** | [**ProjectAnalyticsSuggestionStats**](ProjectAnalyticsSuggestionStats.md) |  | 
**SyncStats** | [**ProjectAnalyticsSyncStats**](ProjectAnalyticsSyncStats.md) |  | 
**ContentStats** | [**ProjectAnalyticsContentStats**](ProjectAnalyticsContentStats.md) |  | 
**TopSourcePages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TopTargetPages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OrphanedPages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HighPerformingPages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MostSuggestedPages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MostLinkedToPages** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RecentImplementations** | Pointer to [**[]Implementation**](Implementation.md) |  | [optional] 
**RecentSuggestions** | Pointer to [**[]Suggestion**](Suggestion.md) |  | [optional] 
**RecentSyncs** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewProjectAnalytics

`func NewProjectAnalytics(implementationStats ProjectAnalyticsImplementationStats, suggestionStats ProjectAnalyticsSuggestionStats, syncStats ProjectAnalyticsSyncStats, contentStats ProjectAnalyticsContentStats, ) *ProjectAnalytics`

NewProjectAnalytics instantiates a new ProjectAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAnalyticsWithDefaults

`func NewProjectAnalyticsWithDefaults() *ProjectAnalytics`

NewProjectAnalyticsWithDefaults instantiates a new ProjectAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImplementationStats

`func (o *ProjectAnalytics) GetImplementationStats() ProjectAnalyticsImplementationStats`

GetImplementationStats returns the ImplementationStats field if non-nil, zero value otherwise.

### GetImplementationStatsOk

`func (o *ProjectAnalytics) GetImplementationStatsOk() (*ProjectAnalyticsImplementationStats, bool)`

GetImplementationStatsOk returns a tuple with the ImplementationStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationStats

`func (o *ProjectAnalytics) SetImplementationStats(v ProjectAnalyticsImplementationStats)`

SetImplementationStats sets ImplementationStats field to given value.


### GetSuggestionStats

`func (o *ProjectAnalytics) GetSuggestionStats() ProjectAnalyticsSuggestionStats`

GetSuggestionStats returns the SuggestionStats field if non-nil, zero value otherwise.

### GetSuggestionStatsOk

`func (o *ProjectAnalytics) GetSuggestionStatsOk() (*ProjectAnalyticsSuggestionStats, bool)`

GetSuggestionStatsOk returns a tuple with the SuggestionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionStats

`func (o *ProjectAnalytics) SetSuggestionStats(v ProjectAnalyticsSuggestionStats)`

SetSuggestionStats sets SuggestionStats field to given value.


### GetSyncStats

`func (o *ProjectAnalytics) GetSyncStats() ProjectAnalyticsSyncStats`

GetSyncStats returns the SyncStats field if non-nil, zero value otherwise.

### GetSyncStatsOk

`func (o *ProjectAnalytics) GetSyncStatsOk() (*ProjectAnalyticsSyncStats, bool)`

GetSyncStatsOk returns a tuple with the SyncStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStats

`func (o *ProjectAnalytics) SetSyncStats(v ProjectAnalyticsSyncStats)`

SetSyncStats sets SyncStats field to given value.


### GetContentStats

`func (o *ProjectAnalytics) GetContentStats() ProjectAnalyticsContentStats`

GetContentStats returns the ContentStats field if non-nil, zero value otherwise.

### GetContentStatsOk

`func (o *ProjectAnalytics) GetContentStatsOk() (*ProjectAnalyticsContentStats, bool)`

GetContentStatsOk returns a tuple with the ContentStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentStats

`func (o *ProjectAnalytics) SetContentStats(v ProjectAnalyticsContentStats)`

SetContentStats sets ContentStats field to given value.


### GetTopSourcePages

`func (o *ProjectAnalytics) GetTopSourcePages() []map[string]interface{}`

GetTopSourcePages returns the TopSourcePages field if non-nil, zero value otherwise.

### GetTopSourcePagesOk

`func (o *ProjectAnalytics) GetTopSourcePagesOk() (*[]map[string]interface{}, bool)`

GetTopSourcePagesOk returns a tuple with the TopSourcePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSourcePages

`func (o *ProjectAnalytics) SetTopSourcePages(v []map[string]interface{})`

SetTopSourcePages sets TopSourcePages field to given value.

### HasTopSourcePages

`func (o *ProjectAnalytics) HasTopSourcePages() bool`

HasTopSourcePages returns a boolean if a field has been set.

### GetTopTargetPages

`func (o *ProjectAnalytics) GetTopTargetPages() []map[string]interface{}`

GetTopTargetPages returns the TopTargetPages field if non-nil, zero value otherwise.

### GetTopTargetPagesOk

`func (o *ProjectAnalytics) GetTopTargetPagesOk() (*[]map[string]interface{}, bool)`

GetTopTargetPagesOk returns a tuple with the TopTargetPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopTargetPages

`func (o *ProjectAnalytics) SetTopTargetPages(v []map[string]interface{})`

SetTopTargetPages sets TopTargetPages field to given value.

### HasTopTargetPages

`func (o *ProjectAnalytics) HasTopTargetPages() bool`

HasTopTargetPages returns a boolean if a field has been set.

### GetOrphanedPages

`func (o *ProjectAnalytics) GetOrphanedPages() []map[string]interface{}`

GetOrphanedPages returns the OrphanedPages field if non-nil, zero value otherwise.

### GetOrphanedPagesOk

`func (o *ProjectAnalytics) GetOrphanedPagesOk() (*[]map[string]interface{}, bool)`

GetOrphanedPagesOk returns a tuple with the OrphanedPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedPages

`func (o *ProjectAnalytics) SetOrphanedPages(v []map[string]interface{})`

SetOrphanedPages sets OrphanedPages field to given value.

### HasOrphanedPages

`func (o *ProjectAnalytics) HasOrphanedPages() bool`

HasOrphanedPages returns a boolean if a field has been set.

### GetHighPerformingPages

`func (o *ProjectAnalytics) GetHighPerformingPages() []map[string]interface{}`

GetHighPerformingPages returns the HighPerformingPages field if non-nil, zero value otherwise.

### GetHighPerformingPagesOk

`func (o *ProjectAnalytics) GetHighPerformingPagesOk() (*[]map[string]interface{}, bool)`

GetHighPerformingPagesOk returns a tuple with the HighPerformingPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPerformingPages

`func (o *ProjectAnalytics) SetHighPerformingPages(v []map[string]interface{})`

SetHighPerformingPages sets HighPerformingPages field to given value.

### HasHighPerformingPages

`func (o *ProjectAnalytics) HasHighPerformingPages() bool`

HasHighPerformingPages returns a boolean if a field has been set.

### GetMostSuggestedPages

`func (o *ProjectAnalytics) GetMostSuggestedPages() []map[string]interface{}`

GetMostSuggestedPages returns the MostSuggestedPages field if non-nil, zero value otherwise.

### GetMostSuggestedPagesOk

`func (o *ProjectAnalytics) GetMostSuggestedPagesOk() (*[]map[string]interface{}, bool)`

GetMostSuggestedPagesOk returns a tuple with the MostSuggestedPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostSuggestedPages

`func (o *ProjectAnalytics) SetMostSuggestedPages(v []map[string]interface{})`

SetMostSuggestedPages sets MostSuggestedPages field to given value.

### HasMostSuggestedPages

`func (o *ProjectAnalytics) HasMostSuggestedPages() bool`

HasMostSuggestedPages returns a boolean if a field has been set.

### GetMostLinkedToPages

`func (o *ProjectAnalytics) GetMostLinkedToPages() []map[string]interface{}`

GetMostLinkedToPages returns the MostLinkedToPages field if non-nil, zero value otherwise.

### GetMostLinkedToPagesOk

`func (o *ProjectAnalytics) GetMostLinkedToPagesOk() (*[]map[string]interface{}, bool)`

GetMostLinkedToPagesOk returns a tuple with the MostLinkedToPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostLinkedToPages

`func (o *ProjectAnalytics) SetMostLinkedToPages(v []map[string]interface{})`

SetMostLinkedToPages sets MostLinkedToPages field to given value.

### HasMostLinkedToPages

`func (o *ProjectAnalytics) HasMostLinkedToPages() bool`

HasMostLinkedToPages returns a boolean if a field has been set.

### GetRecentImplementations

`func (o *ProjectAnalytics) GetRecentImplementations() []Implementation`

GetRecentImplementations returns the RecentImplementations field if non-nil, zero value otherwise.

### GetRecentImplementationsOk

`func (o *ProjectAnalytics) GetRecentImplementationsOk() (*[]Implementation, bool)`

GetRecentImplementationsOk returns a tuple with the RecentImplementations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentImplementations

`func (o *ProjectAnalytics) SetRecentImplementations(v []Implementation)`

SetRecentImplementations sets RecentImplementations field to given value.

### HasRecentImplementations

`func (o *ProjectAnalytics) HasRecentImplementations() bool`

HasRecentImplementations returns a boolean if a field has been set.

### GetRecentSuggestions

`func (o *ProjectAnalytics) GetRecentSuggestions() []Suggestion`

GetRecentSuggestions returns the RecentSuggestions field if non-nil, zero value otherwise.

### GetRecentSuggestionsOk

`func (o *ProjectAnalytics) GetRecentSuggestionsOk() (*[]Suggestion, bool)`

GetRecentSuggestionsOk returns a tuple with the RecentSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSuggestions

`func (o *ProjectAnalytics) SetRecentSuggestions(v []Suggestion)`

SetRecentSuggestions sets RecentSuggestions field to given value.

### HasRecentSuggestions

`func (o *ProjectAnalytics) HasRecentSuggestions() bool`

HasRecentSuggestions returns a boolean if a field has been set.

### GetRecentSyncs

`func (o *ProjectAnalytics) GetRecentSyncs() []map[string]interface{}`

GetRecentSyncs returns the RecentSyncs field if non-nil, zero value otherwise.

### GetRecentSyncsOk

`func (o *ProjectAnalytics) GetRecentSyncsOk() (*[]map[string]interface{}, bool)`

GetRecentSyncsOk returns a tuple with the RecentSyncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSyncs

`func (o *ProjectAnalytics) SetRecentSyncs(v []map[string]interface{})`

SetRecentSyncs sets RecentSyncs field to given value.

### HasRecentSyncs

`func (o *ProjectAnalytics) HasRecentSyncs() bool`

HasRecentSyncs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


