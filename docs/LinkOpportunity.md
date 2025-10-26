# LinkOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelevanceScore** | **float32** | AI-calculated relevance score (0-1) | 
**AnchorText** | **string** | Suggested anchor text | 
**Context** | Pointer to **string** | Context where the link should be placed | [optional] 
**SourcePage** | [**PageInfo**](PageInfo.md) |  | 
**TargetPage** | [**PageInfo**](PageInfo.md) |  | 
**Reasoning** | Pointer to **string** | AI reasoning for the suggestion | [optional] 

## Methods

### NewLinkOpportunity

`func NewLinkOpportunity(relevanceScore float32, anchorText string, sourcePage PageInfo, targetPage PageInfo, ) *LinkOpportunity`

NewLinkOpportunity instantiates a new LinkOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkOpportunityWithDefaults

`func NewLinkOpportunityWithDefaults() *LinkOpportunity`

NewLinkOpportunityWithDefaults instantiates a new LinkOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelevanceScore

`func (o *LinkOpportunity) GetRelevanceScore() float32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *LinkOpportunity) GetRelevanceScoreOk() (*float32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *LinkOpportunity) SetRelevanceScore(v float32)`

SetRelevanceScore sets RelevanceScore field to given value.


### GetAnchorText

`func (o *LinkOpportunity) GetAnchorText() string`

GetAnchorText returns the AnchorText field if non-nil, zero value otherwise.

### GetAnchorTextOk

`func (o *LinkOpportunity) GetAnchorTextOk() (*string, bool)`

GetAnchorTextOk returns a tuple with the AnchorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorText

`func (o *LinkOpportunity) SetAnchorText(v string)`

SetAnchorText sets AnchorText field to given value.


### GetContext

`func (o *LinkOpportunity) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LinkOpportunity) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LinkOpportunity) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *LinkOpportunity) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSourcePage

`func (o *LinkOpportunity) GetSourcePage() PageInfo`

GetSourcePage returns the SourcePage field if non-nil, zero value otherwise.

### GetSourcePageOk

`func (o *LinkOpportunity) GetSourcePageOk() (*PageInfo, bool)`

GetSourcePageOk returns a tuple with the SourcePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePage

`func (o *LinkOpportunity) SetSourcePage(v PageInfo)`

SetSourcePage sets SourcePage field to given value.


### GetTargetPage

`func (o *LinkOpportunity) GetTargetPage() PageInfo`

GetTargetPage returns the TargetPage field if non-nil, zero value otherwise.

### GetTargetPageOk

`func (o *LinkOpportunity) GetTargetPageOk() (*PageInfo, bool)`

GetTargetPageOk returns a tuple with the TargetPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPage

`func (o *LinkOpportunity) SetTargetPage(v PageInfo)`

SetTargetPage sets TargetPage field to given value.


### GetReasoning

`func (o *LinkOpportunity) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *LinkOpportunity) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *LinkOpportunity) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *LinkOpportunity) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


