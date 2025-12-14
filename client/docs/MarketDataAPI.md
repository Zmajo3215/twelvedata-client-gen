# \MarketDataAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrencyConversion**](MarketDataAPI.md#GetCurrencyConversion) | **Get** /currency_conversion | Currency conversion
[**GetEod**](MarketDataAPI.md#GetEod) | **Get** /eod | End of day price
[**GetExchangeRate**](MarketDataAPI.md#GetExchangeRate) | **Get** /exchange_rate | Exchange rate
[**GetMarketMovers**](MarketDataAPI.md#GetMarketMovers) | **Get** /market_movers/{market} | Market movers
[**GetPrice**](MarketDataAPI.md#GetPrice) | **Get** /price | Latest price
[**GetQuote**](MarketDataAPI.md#GetQuote) | **Get** /quote | Quote



## GetCurrencyConversion

> GetCurrencyConversion200Response GetCurrencyConversion(ctx).Symbol(symbol).Amount(amount).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()

Currency conversion



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
	symbol := "EUR/USD" // string | The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct
	amount := float64(100) // float64 | Amount of base currency to be converted into quote currency. Supports values in the range from `0` and above
	date := "2006-01-02T15:04:05" // string | If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone (optional)
	format := "format_example" // string | Value can be `JSON` or `CSV`. Default `JSON` (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the `CSV` file. Default semicolon `;` (optional) (default to ";")
	dp := int64(789) // int64 | The number of decimal places for the data (optional) (default to 5)
	timezone := "UTC" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetCurrencyConversion(context.Background()).Symbol(symbol).Amount(amount).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetCurrencyConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrencyConversion`: GetCurrencyConversion200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetCurrencyConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct | 
 **amount** | **float64** | Amount of base currency to be converted into quote currency. Supports values in the range from &#x60;0&#x60; and above | 
 **date** | **string** | If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone | 
 **format** | **string** | Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60; | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60; | [default to &quot;;&quot;]
 **dp** | **int64** | The number of decimal places for the data | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | 

### Return type

[**GetCurrencyConversion200Response**](GetCurrencyConversion200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEod

> GetEod200Response GetEod(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Date(date).Prepost(prepost).Dp(dp).Execute()

End of day price



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	type_ := "ETF" // string | The asset class to which the instrument belongs (optional)
	date := "2006-01-02" // string | If not null, then return data from a specific date (optional)
	prepost := true // bool | Parameter is optional. Only for `Pro` and above plans. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume (optional) (default to false)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetEod(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Date(date).Prepost(prepost).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetEod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEod`: GetEod200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetEod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | **string** | The asset class to which the instrument belongs | 
 **date** | **string** | If not null, then return data from a specific date | 
 **prepost** | **bool** | Parameter is optional. Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume | [default to false]
 **dp** | **int64** | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive | [default to 5]

### Return type

[**GetEod200Response**](GetEod200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeRate

> GetExchangeRate200Response GetExchangeRate(ctx).Symbol(symbol).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()

Exchange rate



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
	symbol := "EUR/USD" // string | The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct
	date := "2006-01-02T15:04:05" // string | If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone (optional)
	format := "format_example" // string | Value can be `JSON` or `CSV`. Default `JSON` (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the `CSV` file. Default semicolon `;` (optional) (default to ";")
	dp := int64(789) // int64 | The number of decimal places for the data (optional) (default to 5)
	timezone := "UTC" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetExchangeRate(context.Background()).Symbol(symbol).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetExchangeRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeRate`: GetExchangeRate200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetExchangeRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct | 
 **date** | **string** | If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone | 
 **format** | **string** | Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60; | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60; | [default to &quot;;&quot;]
 **dp** | **int64** | The number of decimal places for the data | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | 

### Return type

[**GetExchangeRate200Response**](GetExchangeRate200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketMovers

> MarketMoversResponseBody GetMarketMovers(ctx, market).Direction(direction).Outputsize(outputsize).Country(country).PriceGreaterThan(priceGreaterThan).Dp(dp).Execute()

Market movers



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
	market := "stocks" // string | Maket type
	direction := "direction_example" // string | Specifies direction of the snapshot gainers or losers (optional) (default to "gainers")
	outputsize := int64(789) // int64 | Specifies the size of the snapshot. Can be in a range from `1` to `50` (optional) (default to 30)
	country := "country_example" // string | Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code (optional) (default to "USA")
	priceGreaterThan := "175.5" // string | Takes values with price grater than specified value (optional)
	dp := "dp_example" // string | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to "5")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketMovers(context.Background(), market).Direction(direction).Outputsize(outputsize).Country(country).PriceGreaterThan(priceGreaterThan).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketMovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketMovers`: MarketMoversResponseBody
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketMovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**market** | **string** | Maket type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketMoversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** | Specifies direction of the snapshot gainers or losers | [default to &quot;gainers&quot;]
 **outputsize** | **int64** | Specifies the size of the snapshot. Can be in a range from &#x60;1&#x60; to &#x60;50&#x60; | [default to 30]
 **country** | **string** | Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code | [default to &quot;USA&quot;]
 **priceGreaterThan** | **string** | Takes values with price grater than specified value | 
 **dp** | **string** | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive | [default to &quot;5&quot;]

### Return type

[**MarketMoversResponseBody**](MarketMoversResponseBody.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrice

> GetPrice200Response GetPrice(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Dp(dp).Execute()

Latest price



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	type_ := "ETF" // string | The asset class to which the instrument belongs (optional)
	format := "format_example" // string | Value can be JSON or CSV (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	prepost := true // bool | Parameter is optional. Only for Pro and above plans. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume. (optional) (default to false)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetPrice(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrice`: GetPrice200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | **string** | The asset class to which the instrument belongs | 
 **format** | **string** | Value can be JSON or CSV | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **prepost** | **bool** | Parameter is optional. Only for Pro and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume. | [default to false]
 **dp** | **int64** | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive | [default to 5]

### Return type

[**GetPrice200Response**](GetPrice200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> GetQuote200Response GetQuote(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Interval(interval).Exchange(exchange).MicCode(micCode).Country(country).VolumeTimePeriod(volumeTimePeriod).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Eod(eod).RollingPeriod(rollingPeriod).Dp(dp).Timezone(timezone).Execute()

Quote



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI) (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN) (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Add-ons</a> section (optional)
	interval := "interval_example" // string | Interval of the quote (optional) (default to "1day")
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	volumeTimePeriod := int64(789) // int64 | Number of periods for Average Volume (optional) (default to 9)
	type_ := "ETF" // string | The asset class to which the instrument belongs (optional)
	format := "format_example" // string | Value can be JSON or CSV Default JSON (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	prepost := true // bool | Parameter is optional. Only for `Pro` and above plans. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume. (optional) (default to false)
	eod := true // bool | If true, then return data for closed day (optional) (default to false)
	rollingPeriod := int64(789) // int64 | Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168]. (optional) (default to 24)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive (optional) (default to 5)
	timezone := "timezone_example" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional) (default to "Exchange")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetQuote(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Interval(interval).Exchange(exchange).MicCode(micCode).Country(country).VolumeTimePeriod(volumeTimePeriod).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Eod(eod).RollingPeriod(rollingPeriod).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: GetQuote200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI) | 
 **isin** | **string** | Filter by international securities identification number (ISIN) | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section | 
 **interval** | **string** | Interval of the quote | [default to &quot;1day&quot;]
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **volumeTimePeriod** | **int64** | Number of periods for Average Volume | [default to 9]
 **type_** | **string** | The asset class to which the instrument belongs | 
 **format** | **string** | Value can be JSON or CSV Default JSON | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **prepost** | **bool** | Parameter is optional. Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume. | [default to false]
 **eod** | **bool** | If true, then return data for closed day | [default to false]
 **rollingPeriod** | **int64** | Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168]. | [default to 24]
 **dp** | **int64** | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | [default to &quot;Exchange&quot;]

### Return type

[**GetQuote200Response**](GetQuote200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader), [queryParameter](../README.md#queryParameter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

