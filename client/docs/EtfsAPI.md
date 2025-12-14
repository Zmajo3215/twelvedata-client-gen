# \EtfsAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetETFsWorld**](EtfsAPI.md#GetETFsWorld) | **Get** /etfs/world | ETF full data
[**GetETFsWorldComposition**](EtfsAPI.md#GetETFsWorldComposition) | **Get** /etfs/world/composition | Composition
[**GetETFsWorldPerformance**](EtfsAPI.md#GetETFsWorldPerformance) | **Get** /etfs/world/performance | Performance
[**GetETFsWorldRisk**](EtfsAPI.md#GetETFsWorldRisk) | **Get** /etfs/world/risk | Risk
[**GetETFsWorldSummary**](EtfsAPI.md#GetETFsWorldSummary) | **Get** /etfs/world/summary | Summary



## GetETFsWorld

> GetETFsWorld200Response GetETFsWorld(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

ETF full data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorld(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorld`: GetETFsWorld200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorld200Response**](GetETFsWorld200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldComposition

> GetETFsWorldComposition200Response GetETFsWorldComposition(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Composition



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldComposition(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldComposition`: GetETFsWorldComposition200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldComposition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldComposition200Response**](GetETFsWorldComposition200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldPerformance

> GetETFsWorldPerformance200Response GetETFsWorldPerformance(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Performance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldPerformance(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldPerformance`: GetETFsWorldPerformance200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldPerformance200Response**](GetETFsWorldPerformance200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldRisk

> GetETFsWorldRisk200Response GetETFsWorldRisk(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Risk



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldRisk(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldRisk`: GetETFsWorldRisk200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldRisk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldRisk200Response**](GetETFsWorldRisk200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldSummary

> GetETFsWorldSummary200Response GetETFsWorldSummary(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldSummary(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldSummary`: GetETFsWorldSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldSummary200Response**](GetETFsWorldSummary200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

