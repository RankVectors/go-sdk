# Suggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique suggestion identifier | 
**ProjectId** | **string** | Project identifier | 
**Status** | **string** | Suggestion status | 
**RelevanceScore** | **float32** | AI-calculated relevance score (0-1) | 
**AnchorText** | **string** | Suggested anchor text | 
**Context** | Pointer to **string** | Context where the link should be placed | [optional] 
**SourcePage** | [**PageInfo**](PageInfo.md) |  | 
**TargetPage** | [**PageInfo**](PageInfo.md) |  | 
**Reasoning** | Pointer to **string** | AI reasoning for the suggestion | [optional] 
**CreatedAt** | **time.Time** | Suggestion creation timestamp | 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 

## Methods

### NewSuggestion

`func NewSuggestion(id string, projectId string, status string, relevanceScore float32, anchorText string, sourcePage PageInfo, targetPage PageInfo, createdAt time.Time, ) *Suggestion`

NewSuggestion instantiates a new Suggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestionWithDefaults

`func NewSuggestionWithDefaults() *Suggestion`

NewSuggestionWithDefaults instantiates a new Suggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Suggestion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Suggestion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Suggestion) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Suggestion) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Suggestion) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Suggestion) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStatus

`func (o *Suggestion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Suggestion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Suggestion) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRelevanceScore

`func (o *Suggestion) GetRelevanceScore() float32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *Suggestion) GetRelevanceScoreOk() (*float32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *Suggestion) SetRelevanceScore(v float32)`

SetRelevanceScore sets RelevanceScore field to given value.


### GetAnchorText

`func (o *Suggestion) GetAnchorText() string`

GetAnchorText returns the AnchorText field if non-nil, zero value otherwise.

### GetAnchorTextOk

`func (o *Suggestion) GetAnchorTextOk() (*string, bool)`

GetAnchorTextOk returns a tuple with the AnchorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorText

`func (o *Suggestion) SetAnchorText(v string)`

SetAnchorText sets AnchorText field to given value.


### GetContext

`func (o *Suggestion) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Suggestion) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Suggestion) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Suggestion) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSourcePage

`func (o *Suggestion) GetSourcePage() PageInfo`

GetSourcePage returns the SourcePage field if non-nil, zero value otherwise.

### GetSourcePageOk

`func (o *Suggestion) GetSourcePageOk() (*PageInfo, bool)`

GetSourcePageOk returns a tuple with the SourcePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePage

`func (o *Suggestion) SetSourcePage(v PageInfo)`

SetSourcePage sets SourcePage field to given value.


### GetTargetPage

`func (o *Suggestion) GetTargetPage() PageInfo`

GetTargetPage returns the TargetPage field if non-nil, zero value otherwise.

### GetTargetPageOk

`func (o *Suggestion) GetTargetPageOk() (*PageInfo, bool)`

GetTargetPageOk returns a tuple with the TargetPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPage

`func (o *Suggestion) SetTargetPage(v PageInfo)`

SetTargetPage sets TargetPage field to given value.


### GetReasoning

`func (o *Suggestion) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *Suggestion) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *Suggestion) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *Suggestion) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Suggestion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Suggestion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Suggestion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Suggestion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Suggestion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Suggestion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Suggestion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


