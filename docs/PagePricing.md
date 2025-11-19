# PagePricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tiers** | Pointer to [**[]PagePricingTiersInner**](PagePricingTiersInner.md) |  | [optional] 
**PagePacks** | Pointer to [**[]PagePricingPagePacksInner**](PagePricingPagePacksInner.md) |  | [optional] 

## Methods

### NewPagePricing

`func NewPagePricing() *PagePricing`

NewPagePricing instantiates a new PagePricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagePricingWithDefaults

`func NewPagePricingWithDefaults() *PagePricing`

NewPagePricingWithDefaults instantiates a new PagePricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiers

`func (o *PagePricing) GetTiers() []PagePricingTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PagePricing) GetTiersOk() (*[]PagePricingTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PagePricing) SetTiers(v []PagePricingTiersInner)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PagePricing) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetPagePacks

`func (o *PagePricing) GetPagePacks() []PagePricingPagePacksInner`

GetPagePacks returns the PagePacks field if non-nil, zero value otherwise.

### GetPagePacksOk

`func (o *PagePricing) GetPagePacksOk() (*[]PagePricingPagePacksInner, bool)`

GetPagePacksOk returns a tuple with the PagePacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagePacks

`func (o *PagePricing) SetPagePacks(v []PagePricingPagePacksInner)`

SetPagePacks sets PagePacks field to given value.

### HasPagePacks

`func (o *PagePricing) HasPagePacks() bool`

HasPagePacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


