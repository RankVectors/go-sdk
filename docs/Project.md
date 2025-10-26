# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique project identifier | 
**Name** | **string** | Project name | 
**Domain** | **string** | Website domain URL | 
**UserId** | **string** | User who owns the project | 
**Prompt** | Pointer to **string** | Natural language prompt for crawling | [optional] 
**SearchQuery** | Pointer to **string** | Search query for targeted crawling | [optional] 
**SitemapMode** | Pointer to **string** | How to handle sitemaps | [optional] 
**IncludeSubdomains** | Pointer to **bool** | Whether to include subdomains | [optional] 
**IgnoreQueryParams** | Pointer to **bool** | Whether to ignore URL query parameters | [optional] 
**MaxDiscoveryDepth** | Pointer to **int32** | Maximum crawl depth | [optional] 
**ExcludePaths** | Pointer to **[]string** | Paths to exclude from crawling | [optional] 
**IncludePaths** | Pointer to **[]string** | Specific paths to include | [optional] 
**CrawlEntireDomain** | Pointer to **bool** | Whether to crawl the entire domain | [optional] 
**AllowExternalLinks** | Pointer to **bool** | Whether to allow external links | [optional] 
**MaxPages** | Pointer to **int32** | Maximum number of pages to crawl | [optional] 
**CrawlDelay** | Pointer to **int32** | Delay between crawl requests (ms) | [optional] 
**CrawlMaxConcurrency** | Pointer to **int32** | Maximum concurrent crawl requests | [optional] 
**OnlyMainContent** | Pointer to **bool** | Whether to extract only main content | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Custom headers for crawling | [optional] 
**WaitFor** | Pointer to **int32** | Wait time for page load (ms) | [optional] 
**BlockAds** | Pointer to **bool** | Whether to block ads | [optional] 
**ProxyMode** | Pointer to **string** | Proxy mode for crawling | [optional] 
**UseReranking** | Pointer to **bool** | Whether to use AI reranking | [optional] 
**EnableChangeTracking** | Pointer to **bool** | Whether to enable change tracking | [optional] 
**CreatedAt** | **time.Time** | Project creation timestamp | 
**UpdatedAt** | **time.Time** | Last update timestamp | 
**Count** | Pointer to [**ProjectCount**](ProjectCount.md) |  | [optional] 

## Methods

### NewProject

`func NewProject(id string, name string, domain string, userId string, createdAt time.Time, updatedAt time.Time, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *Project) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Project) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Project) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUserId

`func (o *Project) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Project) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Project) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPrompt

`func (o *Project) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *Project) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *Project) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *Project) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Project) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Project) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Project) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Project) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetSitemapMode

`func (o *Project) GetSitemapMode() string`

GetSitemapMode returns the SitemapMode field if non-nil, zero value otherwise.

### GetSitemapModeOk

`func (o *Project) GetSitemapModeOk() (*string, bool)`

GetSitemapModeOk returns a tuple with the SitemapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitemapMode

`func (o *Project) SetSitemapMode(v string)`

SetSitemapMode sets SitemapMode field to given value.

### HasSitemapMode

`func (o *Project) HasSitemapMode() bool`

HasSitemapMode returns a boolean if a field has been set.

### GetIncludeSubdomains

`func (o *Project) GetIncludeSubdomains() bool`

GetIncludeSubdomains returns the IncludeSubdomains field if non-nil, zero value otherwise.

### GetIncludeSubdomainsOk

`func (o *Project) GetIncludeSubdomainsOk() (*bool, bool)`

GetIncludeSubdomainsOk returns a tuple with the IncludeSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubdomains

`func (o *Project) SetIncludeSubdomains(v bool)`

SetIncludeSubdomains sets IncludeSubdomains field to given value.

### HasIncludeSubdomains

`func (o *Project) HasIncludeSubdomains() bool`

HasIncludeSubdomains returns a boolean if a field has been set.

### GetIgnoreQueryParams

`func (o *Project) GetIgnoreQueryParams() bool`

GetIgnoreQueryParams returns the IgnoreQueryParams field if non-nil, zero value otherwise.

### GetIgnoreQueryParamsOk

`func (o *Project) GetIgnoreQueryParamsOk() (*bool, bool)`

GetIgnoreQueryParamsOk returns a tuple with the IgnoreQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreQueryParams

`func (o *Project) SetIgnoreQueryParams(v bool)`

SetIgnoreQueryParams sets IgnoreQueryParams field to given value.

### HasIgnoreQueryParams

`func (o *Project) HasIgnoreQueryParams() bool`

HasIgnoreQueryParams returns a boolean if a field has been set.

### GetMaxDiscoveryDepth

`func (o *Project) GetMaxDiscoveryDepth() int32`

GetMaxDiscoveryDepth returns the MaxDiscoveryDepth field if non-nil, zero value otherwise.

### GetMaxDiscoveryDepthOk

`func (o *Project) GetMaxDiscoveryDepthOk() (*int32, bool)`

GetMaxDiscoveryDepthOk returns a tuple with the MaxDiscoveryDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveryDepth

`func (o *Project) SetMaxDiscoveryDepth(v int32)`

SetMaxDiscoveryDepth sets MaxDiscoveryDepth field to given value.

### HasMaxDiscoveryDepth

`func (o *Project) HasMaxDiscoveryDepth() bool`

HasMaxDiscoveryDepth returns a boolean if a field has been set.

### GetExcludePaths

`func (o *Project) GetExcludePaths() []string`

GetExcludePaths returns the ExcludePaths field if non-nil, zero value otherwise.

### GetExcludePathsOk

`func (o *Project) GetExcludePathsOk() (*[]string, bool)`

GetExcludePathsOk returns a tuple with the ExcludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePaths

`func (o *Project) SetExcludePaths(v []string)`

SetExcludePaths sets ExcludePaths field to given value.

### HasExcludePaths

`func (o *Project) HasExcludePaths() bool`

HasExcludePaths returns a boolean if a field has been set.

### GetIncludePaths

`func (o *Project) GetIncludePaths() []string`

GetIncludePaths returns the IncludePaths field if non-nil, zero value otherwise.

### GetIncludePathsOk

`func (o *Project) GetIncludePathsOk() (*[]string, bool)`

GetIncludePathsOk returns a tuple with the IncludePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaths

`func (o *Project) SetIncludePaths(v []string)`

SetIncludePaths sets IncludePaths field to given value.

### HasIncludePaths

`func (o *Project) HasIncludePaths() bool`

HasIncludePaths returns a boolean if a field has been set.

### GetCrawlEntireDomain

`func (o *Project) GetCrawlEntireDomain() bool`

GetCrawlEntireDomain returns the CrawlEntireDomain field if non-nil, zero value otherwise.

### GetCrawlEntireDomainOk

`func (o *Project) GetCrawlEntireDomainOk() (*bool, bool)`

GetCrawlEntireDomainOk returns a tuple with the CrawlEntireDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlEntireDomain

`func (o *Project) SetCrawlEntireDomain(v bool)`

SetCrawlEntireDomain sets CrawlEntireDomain field to given value.

### HasCrawlEntireDomain

`func (o *Project) HasCrawlEntireDomain() bool`

HasCrawlEntireDomain returns a boolean if a field has been set.

### GetAllowExternalLinks

`func (o *Project) GetAllowExternalLinks() bool`

GetAllowExternalLinks returns the AllowExternalLinks field if non-nil, zero value otherwise.

### GetAllowExternalLinksOk

`func (o *Project) GetAllowExternalLinksOk() (*bool, bool)`

GetAllowExternalLinksOk returns a tuple with the AllowExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExternalLinks

`func (o *Project) SetAllowExternalLinks(v bool)`

SetAllowExternalLinks sets AllowExternalLinks field to given value.

### HasAllowExternalLinks

`func (o *Project) HasAllowExternalLinks() bool`

HasAllowExternalLinks returns a boolean if a field has been set.

### GetMaxPages

`func (o *Project) GetMaxPages() int32`

GetMaxPages returns the MaxPages field if non-nil, zero value otherwise.

### GetMaxPagesOk

`func (o *Project) GetMaxPagesOk() (*int32, bool)`

GetMaxPagesOk returns a tuple with the MaxPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPages

`func (o *Project) SetMaxPages(v int32)`

SetMaxPages sets MaxPages field to given value.

### HasMaxPages

`func (o *Project) HasMaxPages() bool`

HasMaxPages returns a boolean if a field has been set.

### GetCrawlDelay

`func (o *Project) GetCrawlDelay() int32`

GetCrawlDelay returns the CrawlDelay field if non-nil, zero value otherwise.

### GetCrawlDelayOk

`func (o *Project) GetCrawlDelayOk() (*int32, bool)`

GetCrawlDelayOk returns a tuple with the CrawlDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlDelay

`func (o *Project) SetCrawlDelay(v int32)`

SetCrawlDelay sets CrawlDelay field to given value.

### HasCrawlDelay

`func (o *Project) HasCrawlDelay() bool`

HasCrawlDelay returns a boolean if a field has been set.

### GetCrawlMaxConcurrency

`func (o *Project) GetCrawlMaxConcurrency() int32`

GetCrawlMaxConcurrency returns the CrawlMaxConcurrency field if non-nil, zero value otherwise.

### GetCrawlMaxConcurrencyOk

`func (o *Project) GetCrawlMaxConcurrencyOk() (*int32, bool)`

GetCrawlMaxConcurrencyOk returns a tuple with the CrawlMaxConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlMaxConcurrency

`func (o *Project) SetCrawlMaxConcurrency(v int32)`

SetCrawlMaxConcurrency sets CrawlMaxConcurrency field to given value.

### HasCrawlMaxConcurrency

`func (o *Project) HasCrawlMaxConcurrency() bool`

HasCrawlMaxConcurrency returns a boolean if a field has been set.

### GetOnlyMainContent

`func (o *Project) GetOnlyMainContent() bool`

GetOnlyMainContent returns the OnlyMainContent field if non-nil, zero value otherwise.

### GetOnlyMainContentOk

`func (o *Project) GetOnlyMainContentOk() (*bool, bool)`

GetOnlyMainContentOk returns a tuple with the OnlyMainContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyMainContent

`func (o *Project) SetOnlyMainContent(v bool)`

SetOnlyMainContent sets OnlyMainContent field to given value.

### HasOnlyMainContent

`func (o *Project) HasOnlyMainContent() bool`

HasOnlyMainContent returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *Project) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *Project) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *Project) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *Project) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetWaitFor

`func (o *Project) GetWaitFor() int32`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *Project) GetWaitForOk() (*int32, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *Project) SetWaitFor(v int32)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *Project) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### GetBlockAds

`func (o *Project) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *Project) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *Project) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *Project) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetProxyMode

`func (o *Project) GetProxyMode() string`

GetProxyMode returns the ProxyMode field if non-nil, zero value otherwise.

### GetProxyModeOk

`func (o *Project) GetProxyModeOk() (*string, bool)`

GetProxyModeOk returns a tuple with the ProxyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyMode

`func (o *Project) SetProxyMode(v string)`

SetProxyMode sets ProxyMode field to given value.

### HasProxyMode

`func (o *Project) HasProxyMode() bool`

HasProxyMode returns a boolean if a field has been set.

### GetUseReranking

`func (o *Project) GetUseReranking() bool`

GetUseReranking returns the UseReranking field if non-nil, zero value otherwise.

### GetUseRerankingOk

`func (o *Project) GetUseRerankingOk() (*bool, bool)`

GetUseRerankingOk returns a tuple with the UseReranking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReranking

`func (o *Project) SetUseReranking(v bool)`

SetUseReranking sets UseReranking field to given value.

### HasUseReranking

`func (o *Project) HasUseReranking() bool`

HasUseReranking returns a boolean if a field has been set.

### GetEnableChangeTracking

`func (o *Project) GetEnableChangeTracking() bool`

GetEnableChangeTracking returns the EnableChangeTracking field if non-nil, zero value otherwise.

### GetEnableChangeTrackingOk

`func (o *Project) GetEnableChangeTrackingOk() (*bool, bool)`

GetEnableChangeTrackingOk returns a tuple with the EnableChangeTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChangeTracking

`func (o *Project) SetEnableChangeTracking(v bool)`

SetEnableChangeTracking sets EnableChangeTracking field to given value.

### HasEnableChangeTracking

`func (o *Project) HasEnableChangeTracking() bool`

HasEnableChangeTracking returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Project) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Project) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCount

`func (o *Project) GetCount() ProjectCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Project) GetCountOk() (*ProjectCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Project) SetCount(v ProjectCount)`

SetCount sets Count field to given value.

### HasCount

`func (o *Project) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


