# CustomCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | **string** | Webhook URL for implementation | 
**ApiKey** | **string** | Webhook secret key | 

## Methods

### NewCustomCredentials

`func NewCustomCredentials(webhookUrl string, apiKey string, ) *CustomCredentials`

NewCustomCredentials instantiates a new CustomCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCredentialsWithDefaults

`func NewCustomCredentialsWithDefaults() *CustomCredentials`

NewCustomCredentialsWithDefaults instantiates a new CustomCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *CustomCredentials) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CustomCredentials) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CustomCredentials) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetApiKey

`func (o *CustomCredentials) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CustomCredentials) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CustomCredentials) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


