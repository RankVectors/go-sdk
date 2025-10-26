# ImplementationInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestionId** | **string** | Suggestion identifier | 
**PageUrl** | **string** | URL of the page to modify | 
**AnchorText** | **string** | Text to turn into a link | 
**TargetUrl** | **string** | URL to link to | 
**Context** | **string** | Context where the link should be placed | 
**Instructions** | [**StepByStepInstructions**](StepByStepInstructions.md) |  | 

## Methods

### NewImplementationInstructions

`func NewImplementationInstructions(suggestionId string, pageUrl string, anchorText string, targetUrl string, context string, instructions StepByStepInstructions, ) *ImplementationInstructions`

NewImplementationInstructions instantiates a new ImplementationInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplementationInstructionsWithDefaults

`func NewImplementationInstructionsWithDefaults() *ImplementationInstructions`

NewImplementationInstructionsWithDefaults instantiates a new ImplementationInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestionId

`func (o *ImplementationInstructions) GetSuggestionId() string`

GetSuggestionId returns the SuggestionId field if non-nil, zero value otherwise.

### GetSuggestionIdOk

`func (o *ImplementationInstructions) GetSuggestionIdOk() (*string, bool)`

GetSuggestionIdOk returns a tuple with the SuggestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionId

`func (o *ImplementationInstructions) SetSuggestionId(v string)`

SetSuggestionId sets SuggestionId field to given value.


### GetPageUrl

`func (o *ImplementationInstructions) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *ImplementationInstructions) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *ImplementationInstructions) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetAnchorText

`func (o *ImplementationInstructions) GetAnchorText() string`

GetAnchorText returns the AnchorText field if non-nil, zero value otherwise.

### GetAnchorTextOk

`func (o *ImplementationInstructions) GetAnchorTextOk() (*string, bool)`

GetAnchorTextOk returns a tuple with the AnchorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorText

`func (o *ImplementationInstructions) SetAnchorText(v string)`

SetAnchorText sets AnchorText field to given value.


### GetTargetUrl

`func (o *ImplementationInstructions) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *ImplementationInstructions) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *ImplementationInstructions) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetContext

`func (o *ImplementationInstructions) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ImplementationInstructions) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ImplementationInstructions) SetContext(v string)`

SetContext sets Context field to given value.


### GetInstructions

`func (o *ImplementationInstructions) GetInstructions() StepByStepInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ImplementationInstructions) GetInstructionsOk() (*StepByStepInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ImplementationInstructions) SetInstructions(v StepByStepInstructions)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


