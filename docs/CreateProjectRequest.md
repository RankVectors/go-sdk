# CreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Project name | 
**Domain** | **string** | Website domain URL | 
**Prompt** | Pointer to **string** | Natural language prompt for crawling | [optional] 
**SearchQuery** | Pointer to **string** | Search query for targeted crawling | [optional] 
**SitemapMode** | Pointer to **string** | How to handle sitemaps | [optional] [default to "include"]
**IncludeSubdomains** | Pointer to **bool** | Whether to include subdomains | [optional] [default to true]
**IgnoreQueryParams** | Pointer to **bool** | Whether to ignore URL query parameters | [optional] [default to true]
**MaxDiscoveryDepth** | Pointer to **int32** | Maximum crawl depth | [optional] 
**ExcludePaths** | Pointer to **[]string** | Paths to exclude from crawling | [optional] 
**IncludePaths** | Pointer to **[]string** | Specific paths to include | [optional] 
**CrawlEntireDomain** | Pointer to **bool** | Whether to crawl the entire domain | [optional] [default to false]
**AllowExternalLinks** | Pointer to **bool** | Whether to allow external links | [optional] [default to false]
**MaxPages** | Pointer to **int32** | Maximum number of pages to crawl | [optional] [default to 100]
**CrawlDelay** | Pointer to **int32** | Delay between crawl requests (ms) | [optional] 
**CrawlMaxConcurrency** | Pointer to **int32** | Maximum concurrent crawl requests | [optional] 
**OnlyMainContent** | Pointer to **bool** | Whether to extract only main content | [optional] [default to true]
**CustomHeaders** | Pointer to **map[string]string** | Custom headers for crawling | [optional] 
**WaitFor** | Pointer to **int32** | Wait time for page load (ms) | [optional] [default to 0]
**BlockAds** | Pointer to **bool** | Whether to block ads | [optional] [default to true]
**ProxyMode** | Pointer to **string** | Proxy mode for crawling | [optional] [default to "auto"]
**UseReranking** | Pointer to **bool** | Whether to use AI reranking | [optional] [default to true]
**EnableChangeTracking** | Pointer to **bool** | Whether to enable change tracking | [optional] [default to false]

## Methods

### NewCreateProjectRequest

`func NewCreateProjectRequest(name string, domain string, ) *CreateProjectRequest`

NewCreateProjectRequest instantiates a new CreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectRequestWithDefaults

`func NewCreateProjectRequestWithDefaults() *CreateProjectRequest`

NewCreateProjectRequestWithDefaults instantiates a new CreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *CreateProjectRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateProjectRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateProjectRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPrompt

`func (o *CreateProjectRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CreateProjectRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CreateProjectRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CreateProjectRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSearchQuery

`func (o *CreateProjectRequest) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *CreateProjectRequest) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *CreateProjectRequest) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *CreateProjectRequest) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetSitemapMode

`func (o *CreateProjectRequest) GetSitemapMode() string`

GetSitemapMode returns the SitemapMode field if non-nil, zero value otherwise.

### GetSitemapModeOk

`func (o *CreateProjectRequest) GetSitemapModeOk() (*string, bool)`

GetSitemapModeOk returns a tuple with the SitemapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitemapMode

`func (o *CreateProjectRequest) SetSitemapMode(v string)`

SetSitemapMode sets SitemapMode field to given value.

### HasSitemapMode

`func (o *CreateProjectRequest) HasSitemapMode() bool`

HasSitemapMode returns a boolean if a field has been set.

### GetIncludeSubdomains

`func (o *CreateProjectRequest) GetIncludeSubdomains() bool`

GetIncludeSubdomains returns the IncludeSubdomains field if non-nil, zero value otherwise.

### GetIncludeSubdomainsOk

`func (o *CreateProjectRequest) GetIncludeSubdomainsOk() (*bool, bool)`

GetIncludeSubdomainsOk returns a tuple with the IncludeSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubdomains

`func (o *CreateProjectRequest) SetIncludeSubdomains(v bool)`

SetIncludeSubdomains sets IncludeSubdomains field to given value.

### HasIncludeSubdomains

`func (o *CreateProjectRequest) HasIncludeSubdomains() bool`

HasIncludeSubdomains returns a boolean if a field has been set.

### GetIgnoreQueryParams

`func (o *CreateProjectRequest) GetIgnoreQueryParams() bool`

GetIgnoreQueryParams returns the IgnoreQueryParams field if non-nil, zero value otherwise.

### GetIgnoreQueryParamsOk

`func (o *CreateProjectRequest) GetIgnoreQueryParamsOk() (*bool, bool)`

GetIgnoreQueryParamsOk returns a tuple with the IgnoreQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreQueryParams

`func (o *CreateProjectRequest) SetIgnoreQueryParams(v bool)`

SetIgnoreQueryParams sets IgnoreQueryParams field to given value.

### HasIgnoreQueryParams

`func (o *CreateProjectRequest) HasIgnoreQueryParams() bool`

HasIgnoreQueryParams returns a boolean if a field has been set.

### GetMaxDiscoveryDepth

`func (o *CreateProjectRequest) GetMaxDiscoveryDepth() int32`

GetMaxDiscoveryDepth returns the MaxDiscoveryDepth field if non-nil, zero value otherwise.

### GetMaxDiscoveryDepthOk

`func (o *CreateProjectRequest) GetMaxDiscoveryDepthOk() (*int32, bool)`

GetMaxDiscoveryDepthOk returns a tuple with the MaxDiscoveryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveryDepth

`func (o *CreateProjectRequest) SetMaxDiscoveryDepth(v int32)`

SetMaxDiscoveryDepth sets MaxDiscoveryDepth field to given value.

### HasMaxDiscoveryDepth

`func (o *CreateProjectRequest) HasMaxDiscoveryDepth() bool`

HasMaxDiscoveryDepth returns a boolean if a field has been set.

### GetExcludePaths

`func (o *CreateProjectRequest) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *CreateProjectRequest) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *CreateProjectRequest) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *CreateProjectRequest) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### GetIncludePaths

`func (o *CreateProjectRequest) GetIncludePaths() []string`

GetIncludePaths returns the IncludePaths field if non-nil, zero value otherwise.

### GetIncludePathsOk

`func (o *CreateProjectRequest) GetIncludePathsOk() (*[]string, bool)`

GetIncludePathsOk returns a tuple with the IncludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaths

`func (o *CreateProjectRequest) SetIncludePaths(v []string)`

SetIncludePaths sets IncludePaths field to given value.

### HasIncludePaths

`func (o *CreateProjectRequest) HasIncludePaths() bool`

HasIncludePaths returns a boolean if a field has been set.

### GetCrawlEntireDomain

`func (o *CreateProjectRequest) GetCrawlEntireDomain() bool`

GetCrawlEntireDomain returns the CrawlEntireDomain field if non-nil, zero value otherwise.

### GetCrawlEntireDomainOk

`func (o *CreateProjectRequest) GetCrawlEntireDomainOk() (*bool, bool)`

GetCrawlEntireDomainOk returns a tuple with the CrawlEntireDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlEntireDomain

`func (o *CreateProjectRequest) SetCrawlEntireDomain(v bool)`

SetCrawlEntireDomain sets CrawlEntireDomain field to given value.

### HasCrawlEntireDomain

`func (o *CreateProjectRequest) HasCrawlEntireDomain() bool`

HasCrawlEntireDomain returns a boolean if a field has been set.

### GetAllowExternalLinks

`func (o *CreateProjectRequest) GetAllowExternalLinks() bool`

GetAllowExternalLinks returns the AllowExternalLinks field if non-nil, zero value otherwise.

### GetAllowExternalLinksOk

`func (o *CreateProjectRequest) GetAllowExternalLinksOk() (*bool, bool)`

GetAllowExternalLinksOk returns a tuple with the AllowExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExternalLinks

`func (o *CreateProjectRequest) SetAllowExternalLinks(v bool)`

SetAllowExternalLinks sets AllowExternalLinks field to given value.

### HasAllowExternalLinks

`func (o *CreateProjectRequest) HasAllowExternalLinks() bool`

HasAllowExternalLinks returns a boolean if a field has been set.

### GetMaxPages

`func (o *CreateProjectRequest) GetMaxPages() int32`

GetMaxPages returns the MaxPages field if non-nil, zero value otherwise.

### GetMaxPagesOk

`func (o *CreateProjectRequest) GetMaxPagesOk() (*int32, bool)`

GetMaxPagesOk returns a tuple with the MaxPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPages

`func (o *CreateProjectRequest) SetMaxPages(v int32)`

SetMaxPages sets MaxPages field to given value.

### HasMaxPages

`func (o *CreateProjectRequest) HasMaxPages() bool`

HasMaxPages returns a boolean if a field has been set.

### GetCrawlDelay

`func (o *CreateProjectRequest) GetCrawlDelay() int32`

GetCrawlDelay returns the CrawlDelay field if non-nil, zero value otherwise.

### GetCrawlDelayOk

`func (o *CreateProjectRequest) GetCrawlDelayOk() (*int32, bool)`

GetCrawlDelayOk returns a tuple with the CrawlDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlDelay

`func (o *CreateProjectRequest) SetCrawlDelay(v int32)`

SetCrawlDelay sets CrawlDelay field to given value.

### HasCrawlDelay

`func (o *CreateProjectRequest) HasCrawlDelay() bool`

HasCrawlDelay returns a boolean if a field has been set.

### GetCrawlMaxConcurrency

`func (o *CreateProjectRequest) GetCrawlMaxConcurrency() int32`

GetCrawlMaxConcurrency returns the CrawlMaxConcurrency field if non-nil, zero value otherwise.

### GetCrawlMaxConcurrencyOk

`func (o *CreateProjectRequest) GetCrawlMaxConcurrencyOk() (*int32, bool)`

GetCrawlMaxConcurrencyOk returns a tuple with the CrawlMaxConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlMaxConcurrency

`func (o *CreateProjectRequest) SetCrawlMaxConcurrency(v int32)`

SetCrawlMaxConcurrency sets CrawlMaxConcurrency field to given value.

### HasCrawlMaxConcurrency

`func (o *CreateProjectRequest) HasCrawlMaxConcurrency() bool`

HasCrawlMaxConcurrency returns a boolean if a field has been set.

### GetOnlyMainContent

`func (o *CreateProjectRequest) GetOnlyMainContent() bool`

GetOnlyMainContent returns the OnlyMainContent field if non-nil, zero value otherwise.

### GetOnlyMainContentOk

`func (o *CreateProjectRequest) GetOnlyMainContentOk() (*bool, bool)`

GetOnlyMainContentOk returns a tuple with the OnlyMainContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyMainContent

`func (o *CreateProjectRequest) SetOnlyMainContent(v bool)`

SetOnlyMainContent sets OnlyMainContent field to given value.

### HasOnlyMainContent

`func (o *CreateProjectRequest) HasOnlyMainContent() bool`

HasOnlyMainContent returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CreateProjectRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateProjectRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateProjectRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CreateProjectRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetWaitFor

`func (o *CreateProjectRequest) GetWaitFor() int32`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *CreateProjectRequest) GetWaitForOk() (*int32, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *CreateProjectRequest) SetWaitFor(v int32)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *CreateProjectRequest) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### GetBlockAds

`func (o *CreateProjectRequest) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *CreateProjectRequest) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *CreateProjectRequest) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *CreateProjectRequest) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetProxyMode

`func (o *CreateProjectRequest) GetProxyMode() string`

GetProxyMode returns the ProxyMode field if non-nil, zero value otherwise.

### GetProxyModeOk

`func (o *CreateProjectRequest) GetProxyModeOk() (*string, bool)`

GetProxyModeOk returns a tuple with the ProxyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyMode

`func (o *CreateProjectRequest) SetProxyMode(v string)`

SetProxyMode sets ProxyMode field to given value.

### HasProxyMode

`func (o *CreateProjectRequest) HasProxyMode() bool`

HasProxyMode returns a boolean if a field has been set.

### GetUseReranking

`func (o *CreateProjectRequest) GetUseReranking() bool`

GetUseReranking returns the UseReranking field if non-nil, zero value otherwise.

### GetUseRerankingOk

`func (o *CreateProjectRequest) GetUseRerankingOk() (*bool, bool)`

GetUseRerankingOk returns a tuple with the UseReranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReranking

`func (o *CreateProjectRequest) SetUseReranking(v bool)`

SetUseReranking sets UseReranking field to given value.

### HasUseReranking

`func (o *CreateProjectRequest) HasUseReranking() bool`

HasUseReranking returns a boolean if a field has been set.

### GetEnableChangeTracking

`func (o *CreateProjectRequest) GetEnableChangeTracking() bool`

GetEnableChangeTracking returns the EnableChangeTracking field if non-nil, zero value otherwise.

### GetEnableChangeTrackingOk

`func (o *CreateProjectRequest) GetEnableChangeTrackingOk() (*bool, bool)`

GetEnableChangeTrackingOk returns a tuple with the EnableChangeTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChangeTracking

`func (o *CreateProjectRequest) SetEnableChangeTracking(v bool)`

SetEnableChangeTracking sets EnableChangeTracking field to given value.

### HasEnableChangeTracking

`func (o *CreateProjectRequest) HasEnableChangeTracking() bool`

HasEnableChangeTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


