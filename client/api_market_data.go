/*
Twelve Data API

## Overview  Welcome to Twelve Data developer docs — your gateway to comprehensive financial market data through a powerful and easy-to-use API. Twelve Data provides access to financial markets across over 50 global countries, covering more than 1 million public instruments, including stocks, forex, ETFs, mutual funds, commodities, and cryptocurrencies.  ## Quickstart  To get started, you'll need to sign up for an API key. Once you have your API key, you can start making requests to the API.  ### Step 1: Create Twelve Data account  Sign up on the Twelve Data website to create your account [here](https://twelvedata.com/register). This gives you access to the API dashboard and your API key.  ### Step 2: Get your API key  After signing in, navigate to your [dashboard](https://twelvedata.com/account/api-keys) to find your unique API key. This key is required to authenticate all API and WebSocket requests.  ### Step 3: Make your first request  Try a simple API call with cURL to fetch the latest price for Apple (AAPL):  ``` curl \"https://api.twelvedata.com/price?symbol=AAPL&apikey=your_api_key\" ```  ### Step 4: Make a request from Python or Javascript  Use our client libraries or standard HTTP clients to make API calls programmatically. Here’s an example in [Python](https://github.com/twelvedata/twelvedata-python) and JavaScript:  #### Python (using official Twelve Data SDK):  ```python from twelvedata import TDClient  # Initialize client with your API key td = TDClient(apikey=\"your_api_key\")  # Get latest price for Apple price = td.price(symbol=\"AAPL\").as_json()  print(price) ```  #### JavaScript (Node.js):  ```javascript const fetch = require('node-fetch');  fetch('https://api.twelvedata.com/price?symbol=AAPL&apikey=your_api_key') &nbsp;&nbsp;.then(response => response.json()) &nbsp;&nbsp;.then(data => console.log(data)); ```  ### Step 5: Perform correlation analysis between Tesla and Microsoft prices  Fetch historical price data for Tesla (TSLA) and Microsoft (MSFT) and calculate the correlation of their closing prices:  ```python from twelvedata import TDClient import pandas as pd  # Initialize client with your API key td = TDClient(apikey=\"your_api_key\")  # Fetch historical price data for Tesla tsla_ts = td.time_series( &nbsp;&nbsp;&nbsp;&nbsp;symbol=\"TSLA\", &nbsp;&nbsp;&nbsp;&nbsp;interval=\"1day\", &nbsp;&nbsp;&nbsp;&nbsp;outputsize=100 ).as_pandas()  # Fetch historical price data for Microsoft msft_ts = td.time_series( &nbsp;&nbsp;&nbsp;&nbsp;symbol=\"MSFT\", &nbsp;&nbsp;&nbsp;&nbsp;interval=\"1day\", &nbsp;&nbsp;&nbsp;&nbsp;outputsize=100 ).as_pandas()  # Align data on datetime index combined = pd.concat( &nbsp;&nbsp;&nbsp;&nbsp;[tsla_ts['close'].astype(float), msft_ts['close'].astype(float)], &nbsp;&nbsp;&nbsp;&nbsp;axis=1, &nbsp;&nbsp;&nbsp;&nbsp;keys=[\"TSLA\", \"MSFT\"] ).dropna()  # Calculate correlation correlation = combined[\"TSLA\"].corr(combined[\"MSFT\"]) print(f\"Correlation of closing prices between TSLA and MSFT: {correlation:.2f}\") ```  ### Authentication  Authenticate your requests using one of these methods:  #### Query parameter method ``` GET https://api.twelvedata.com/endpoint?symbol=AAPL&apikey=your_api_key ```  #### HTTP header method (recommended) ``` Authorization: apikey your_api_key ```  ##### API key useful information <ul> <li> Demo API key (<code>apikey=demo</code>) available for demo requests</li> <li> Personal API key required for full access</li> <li> Premium endpoints and data require higher-tier plans (testable with <a href=\"https://twelvedata.com/exchanges\">trial symbols</a>)</li> </ul>  ### API endpoints   Service | Base URL | ---------|----------|  REST API | `https://api.twelvedata.com` |  WebSocket | `wss://ws.twelvedata.com` |  ### Parameter guidelines <ul> <li><b>Separator:</b> Use <code>&</code> to separate multiple parameters</li> <li><b>Case sensitivity:</b> Parameter names are case-insensitive</li>  <ul><li><code>symbol=AAPL</code> = <code>symbol=aapl</code></li></ul>  <li><b>Multiple values:</b> Separate with commas where supported</li> </ul>  ### Response handling  #### Default format All responses return JSON format by default unless otherwise specified.  #### Null values <b>Important:</b> Some response fields may contain `null` values when data is unavailable for specific metrics. This is expected behavior, not an error.  ##### Best Practices: <ul> <li>Always implement <code>null</code> value handling in your application</li> <li>Use defensive programming techniques for data processing</li> <li>Consider fallback values or error handling for critical metrics</li> </ul>  #### Error handling Structure your code to gracefully handle: <ul> <li>Network timeouts</li> <li>Rate limiting responses</li> <li>Invalid parameter errors</li> <li>Data unavailability periods</li> </ul>  ##### Best practices <ul> <li><b>Rate limits:</b> Adhere to your plan’s rate limits to avoid throttling. Check your dashboard for details.</li> <li><b>Error handling:</b> Implement retry logic for transient errors (e.g., <code>429 Too Many Requests</code>).</li> <li><b>Caching:</b> Cache responses for frequently accessed data to reduce API calls and improve performance.</li> <li><b>Secure storage:</b> Store your API key securely and never expose it in client-side code or public repositories.</li> </ul>  ## Errors  Twelve Data API employs a standardized error response format, delivering a JSON object with `code`, `message`, and `status` keys for clear and consistent error communication.  ### Codes  Below is a table of possible error codes, their HTTP status, meanings, and resolution steps:   Code | status | Meaning | Resolution |  --- | --- | --- | --- |  **400** | Bad Request | Invalid or incorrect parameter(s) provided. | Check the `message` in the response for details. Refer to the API Documenta­tion to correct the input. |  **401** | Unauthor­ized | Invalid or incorrect API key. | Verify your API key is correct. Sign up for a key <a href=\"https://twelvedata.com/account/api-keys\">here</a>. |  **403** | Forbidden | API key lacks permissions for the requested resource (upgrade required). | Upgrade your plan <a href=\"https://twelvedata.com/pricing\">here</a>. |  **404** | Not Found | Requested data could not be found. | Adjust parameters to be less strict as they may be too restrictive. |  **414** | Parameter Too Long | Input parameter array exceeds the allowed length. | Follow the `message` guidance to adjust the parameter length. |  **429** | Too Many Requests | API request limit reached for your key. | Wait briefly or upgrade your plan <a href=\"https://twelvedata.com/pricing\">here</a>. |  **500** | Internal Server Error | Server-side issue occurred; retry later. | Contact support <a href=\"https://twelvedata.com/contact\">here</a> for assistance. |  ### Example error response  Consider the following invalid request:  ``` https://api.twelvedata.com/time_series?symbol=AAPL&interval=0.99min&apikey=your_api_key ```  Due to the incorrect `interval` value, the API returns:  ```json { &nbsp;&nbsp;\"code\": 400, &nbsp;&nbsp;\"message\": \"Invalid **interval** provided: 0.99min. Supported intervals: 1min, 5min, 15min, 30min, 45min, 1h, 2h, 4h, 8h, 1day, 1week, 1month\", &nbsp;&nbsp;\"status\": \"error\" } ```  Refer to the API Documentation for valid parameter values to resolve such errors.  ## Libraries  Twelve Data provides a growing ecosystem of libraries and integrations to help you build faster and smarter in your preferred environment. Official libraries are actively maintained by the Twelve Data team, while selected community-built libraries offer additional flexibility.  A full list is available on our [GitHub profile](https://github.com/search?q=twelvedata).  ### Official SDKs <ul> <li><b>Python:</b> <a href=\"https://github.com/twelvedata/twelvedata-python\">twelvedata-python</a></li> <li><b>R:</b> <a href=\"https://github.com/twelvedata/twelvedata-r-sdk\">twelvedata-r-sdk</a></li> </ul>  ### AI integrations <ul> <li><b>Twelve Data MCP Server:</b> <a href=\"https://github.com/twelvedata/mcp\">Repository</a> — Model Context Protocol (MCP) server that provides seamless integration with AI assistants and language models, enabling direct access to Twelve Data's financial market data within conversational interfaces and AI workflows.</li> </ul>  ### Spreadsheet add-ons <ul> <li><b>Excel:</b> <a href=\"https://twelvedata.com/excel\">Excel Add-in</a></li> <li><b>Google Sheets:</b> <a href=\"https://twelvedata.com/google-sheets\">Google Sheets Add-on</a></li> </ul>  ### Community libraries  The community has developed libraries in several popular languages. You can explore more community libraries on [GitHub](https://github.com/search?q=twelvedata). <ul> <li><b>C#:</b> <a href=\"https://github.com/pseudomarkets/TwelveDataSharp\">TwelveDataSharp</a></li> <li><b>JavaScript:</b> <a href=\"https://github.com/evzaboun/twelvedata\">twelvedata</a></li> <li><b>PHP:</b> <a href=\"https://github.com/ingelby/twelvedata\">twelvedata</a></li> <li><b>Go:</b> <a href=\"https://github.com/soulgarden/twelvedata\">twelvedata</a></li> <li><b>TypeScript:</b> <a href=\"https://github.com/Clyde-Goodall/twelve-data-wrapper\">twelve-data-wrapper</a></li> </ul>  ### Other Twelve Data repositories <ul> <li><b>searchindex</b> <i>(Go)</i>: <a href=\"https://github.com/twelvedata/searchindex\">Repository</a> — In-memory search index by strings</li> <li><b>ws-tools</b> <i>(Python)</i>: <a href=\"https://github.com/twelvedata/ws-tools\">Repository</a> — Utility tools for WebSocket stream handling</li> </ul>  ### API specification <ul> <li><b>OpenAPI / Swagger:</b> Access the <a href=\"https://api.twelvedata.com/doc/swagger/openapi.json\">complete API specification</a> in OpenAPI format. You can use this file to automatically generate client libraries in your preferred programming language, explore the API interactively via Swagger tools, or integrate Twelve Data seamlessly into your AI and LLM workflows.</li> </ul>

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MarketDataAPIService MarketDataAPI service
type MarketDataAPIService service

type ApiGetCurrencyConversionRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	symbol *string
	amount *float64
	date *string
	format *string
	delimiter *string
	dp *int64
	timezone *string
}

// The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct
func (r ApiGetCurrencyConversionRequest) Symbol(symbol string) ApiGetCurrencyConversionRequest {
	r.symbol = &symbol
	return r
}

// Amount of base currency to be converted into quote currency. Supports values in the range from &#x60;0&#x60; and above
func (r ApiGetCurrencyConversionRequest) Amount(amount float64) ApiGetCurrencyConversionRequest {
	r.amount = &amount
	return r
}

// If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone
func (r ApiGetCurrencyConversionRequest) Date(date string) ApiGetCurrencyConversionRequest {
	r.date = &date
	return r
}

// Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60;
func (r ApiGetCurrencyConversionRequest) Format(format string) ApiGetCurrencyConversionRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60;
func (r ApiGetCurrencyConversionRequest) Delimiter(delimiter string) ApiGetCurrencyConversionRequest {
	r.delimiter = &delimiter
	return r
}

// The number of decimal places for the data
func (r ApiGetCurrencyConversionRequest) Dp(dp int64) ApiGetCurrencyConversionRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetCurrencyConversionRequest) Timezone(timezone string) ApiGetCurrencyConversionRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetCurrencyConversionRequest) Execute() (*GetCurrencyConversion200Response, *http.Response, error) {
	return r.ApiService.GetCurrencyConversionExecute(r)
}

/*
GetCurrencyConversion Currency conversion

The currency conversion endpoint provides real-time exchange rates and calculates the converted amount for specified currency pairs, including both forex and cryptocurrencies. This endpoint is useful for obtaining up-to-date conversion values between two currencies, facilitating tasks such as financial reporting, e-commerce transactions, and travel budgeting.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCurrencyConversionRequest
*/
func (a *MarketDataAPIService) GetCurrencyConversion(ctx context.Context) ApiGetCurrencyConversionRequest {
	return ApiGetCurrencyConversionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCurrencyConversion200Response
func (a *MarketDataAPIService) GetCurrencyConversionExecute(r ApiGetCurrencyConversionRequest) (*GetCurrencyConversion200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCurrencyConversion200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetCurrencyConversion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/currency_conversion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
        var defaultValue string = "JSON"
        parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
        r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
        var defaultValue string = ";"
        parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
        r.delimiter = &defaultValue
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 5
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEodRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	type_ *string
	date *string
	prepost *bool
	dp *int64
}

// Symbol ticker of the instrument
func (r ApiGetEodRequest) Symbol(symbol string) ApiGetEodRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetEodRequest) Figi(figi string) ApiGetEodRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetEodRequest) Isin(isin string) ApiGetEodRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetEodRequest) Cusip(cusip string) ApiGetEodRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetEodRequest) Exchange(exchange string) ApiGetEodRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetEodRequest) MicCode(micCode string) ApiGetEodRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetEodRequest) Country(country string) ApiGetEodRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetEodRequest) Type_(type_ string) ApiGetEodRequest {
	r.type_ = &type_
	return r
}

// If not null, then return data from a specific date
func (r ApiGetEodRequest) Date(date string) ApiGetEodRequest {
	r.date = &date
	return r
}

// Parameter is optional. Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume
func (r ApiGetEodRequest) Prepost(prepost bool) ApiGetEodRequest {
	r.prepost = &prepost
	return r
}

// Specifies the number of decimal places for floating values Should be in range [0,11] inclusive
func (r ApiGetEodRequest) Dp(dp int64) ApiGetEodRequest {
	r.dp = &dp
	return r
}

func (r ApiGetEodRequest) Execute() (*GetEod200Response, *http.Response, error) {
	return r.ApiService.GetEodExecute(r)
}

/*
GetEod End of day price

The End of Day (EOD) Prices endpoint provides the closing price and other relevant metadata for a financial instrument at the end of a trading day. This endpoint is useful for retrieving daily historical data for stocks, ETFs, or other securities, allowing users to track performance over time and compare daily market movements.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEodRequest
*/
func (a *MarketDataAPIService) GetEod(ctx context.Context) ApiGetEodRequest {
	return ApiGetEodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetEod200Response
func (a *MarketDataAPIService) GetEodExecute(r ApiGetEodRequest) (*GetEod200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEod200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetEod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eod"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", defaultValue, "form", "")
        r.prepost = &defaultValue
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 5
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetExchangeRateRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	symbol *string
	date *string
	format *string
	delimiter *string
	dp *int64
	timezone *string
}

// The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct
func (r ApiGetExchangeRateRequest) Symbol(symbol string) ApiGetExchangeRateRequest {
	r.symbol = &symbol
	return r
}

// If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone
func (r ApiGetExchangeRateRequest) Date(date string) ApiGetExchangeRateRequest {
	r.date = &date
	return r
}

// Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60;
func (r ApiGetExchangeRateRequest) Format(format string) ApiGetExchangeRateRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60;
func (r ApiGetExchangeRateRequest) Delimiter(delimiter string) ApiGetExchangeRateRequest {
	r.delimiter = &delimiter
	return r
}

// The number of decimal places for the data
func (r ApiGetExchangeRateRequest) Dp(dp int64) ApiGetExchangeRateRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetExchangeRateRequest) Timezone(timezone string) ApiGetExchangeRateRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetExchangeRateRequest) Execute() (*GetExchangeRate200Response, *http.Response, error) {
	return r.ApiService.GetExchangeRateExecute(r)
}

/*
GetExchangeRate Exchange rate

The exchange rate endpoint provides real-time exchange rates for specified currency pairs, including both forex and cryptocurrency. It returns the current exchange rate value between two currencies, allowing users to quickly access up-to-date conversion rates for financial transactions or market analysis.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetExchangeRateRequest
*/
func (a *MarketDataAPIService) GetExchangeRate(ctx context.Context) ApiGetExchangeRateRequest {
	return ApiGetExchangeRateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetExchangeRate200Response
func (a *MarketDataAPIService) GetExchangeRateExecute(r ApiGetExchangeRateRequest) (*GetExchangeRate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchangeRate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetExchangeRate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange_rate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
        var defaultValue string = "JSON"
        parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
        r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
        var defaultValue string = ";"
        parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
        r.delimiter = &defaultValue
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 5
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMarketMoversRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	market string
	direction *string
	outputsize *int64
	country *string
	priceGreaterThan *string
	dp *string
}

// Specifies direction of the snapshot gainers or losers
func (r ApiGetMarketMoversRequest) Direction(direction string) ApiGetMarketMoversRequest {
	r.direction = &direction
	return r
}

// Specifies the size of the snapshot. Can be in a range from &#x60;1&#x60; to &#x60;50&#x60;
func (r ApiGetMarketMoversRequest) Outputsize(outputsize int64) ApiGetMarketMoversRequest {
	r.outputsize = &outputsize
	return r
}

// Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code
func (r ApiGetMarketMoversRequest) Country(country string) ApiGetMarketMoversRequest {
	r.country = &country
	return r
}

// Takes values with price grater than specified value
func (r ApiGetMarketMoversRequest) PriceGreaterThan(priceGreaterThan string) ApiGetMarketMoversRequest {
	r.priceGreaterThan = &priceGreaterThan
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r ApiGetMarketMoversRequest) Dp(dp string) ApiGetMarketMoversRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMarketMoversRequest) Execute() (*MarketMoversResponseBody, *http.Response, error) {
	return r.ApiService.GetMarketMoversExecute(r)
}

/*
GetMarketMovers Market movers

The market movers endpoint provides a ranked list of the top-gaining and losing assets for the current trading day. It returns detailed data on the highest percentage price increases and decreases since the previous day's close. This endpoint supports international equities, forex, and cryptocurrencies, enabling users to quickly identify significant market movements across various asset classes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param market Maket type
 @return ApiGetMarketMoversRequest
*/
func (a *MarketDataAPIService) GetMarketMovers(ctx context.Context, market string) ApiGetMarketMoversRequest {
	return ApiGetMarketMoversRequest{
		ApiService: a,
		ctx: ctx,
		market: market,
	}
}

// Execute executes the request
//  @return MarketMoversResponseBody
func (a *MarketDataAPIService) GetMarketMoversExecute(r ApiGetMarketMoversRequest) (*MarketMoversResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketMoversResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetMarketMovers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_movers/{market}"
	localVarPath = strings.Replace(localVarPath, "{"+"market"+"}", url.PathEscape(parameterValueToString(r.market, "market")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
	} else {
        var defaultValue string = "gainers"
        parameterAddToHeaderOrQuery(localVarQueryParams, "direction", defaultValue, "form", "")
        r.direction = &defaultValue
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 30
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	} else {
        var defaultValue string = "USA"
        parameterAddToHeaderOrQuery(localVarQueryParams, "country", defaultValue, "form", "")
        r.country = &defaultValue
	}
	if r.priceGreaterThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "price_greater_than", r.priceGreaterThan, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue string = "5"
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPriceRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	type_ *string
	format *string
	delimiter *string
	prepost *bool
	dp *int64
}

// Symbol ticker of the instrument
func (r ApiGetPriceRequest) Symbol(symbol string) ApiGetPriceRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetPriceRequest) Figi(figi string) ApiGetPriceRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetPriceRequest) Isin(isin string) ApiGetPriceRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetPriceRequest) Cusip(cusip string) ApiGetPriceRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetPriceRequest) Exchange(exchange string) ApiGetPriceRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetPriceRequest) MicCode(micCode string) ApiGetPriceRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetPriceRequest) Country(country string) ApiGetPriceRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetPriceRequest) Type_(type_ string) ApiGetPriceRequest {
	r.type_ = &type_
	return r
}

// Value can be JSON or CSV
func (r ApiGetPriceRequest) Format(format string) ApiGetPriceRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r ApiGetPriceRequest) Delimiter(delimiter string) ApiGetPriceRequest {
	r.delimiter = &delimiter
	return r
}

// Parameter is optional. Only for Pro and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r ApiGetPriceRequest) Prepost(prepost bool) ApiGetPriceRequest {
	r.prepost = &prepost
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r ApiGetPriceRequest) Dp(dp int64) ApiGetPriceRequest {
	r.dp = &dp
	return r
}

func (r ApiGetPriceRequest) Execute() (*GetPrice200Response, *http.Response, error) {
	return r.ApiService.GetPriceExecute(r)
}

/*
GetPrice Latest price

The latest price endpoint provides the latest market price for a specified financial instrument. It returns a single data point representing the current (or the most recently available) trading price.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPriceRequest
*/
func (a *MarketDataAPIService) GetPrice(ctx context.Context) ApiGetPriceRequest {
	return ApiGetPriceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPrice200Response
func (a *MarketDataAPIService) GetPriceExecute(r ApiGetPriceRequest) (*GetPrice200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPrice200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetPrice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
        var defaultValue string = "JSON"
        parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
        r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
        var defaultValue string = ";"
        parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
        r.delimiter = &defaultValue
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", defaultValue, "form", "")
        r.prepost = &defaultValue
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 5
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetQuoteRequest struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	interval *string
	exchange *string
	micCode *string
	country *string
	volumeTimePeriod *int64
	type_ *string
	format *string
	delimiter *string
	prepost *bool
	eod *bool
	rollingPeriod *int64
	dp *int64
	timezone *string
}

// Symbol ticker of the instrument
func (r ApiGetQuoteRequest) Symbol(symbol string) ApiGetQuoteRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetQuoteRequest) Figi(figi string) ApiGetQuoteRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetQuoteRequest) Isin(isin string) ApiGetQuoteRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetQuoteRequest) Cusip(cusip string) ApiGetQuoteRequest {
	r.cusip = &cusip
	return r
}

// Interval of the quote
func (r ApiGetQuoteRequest) Interval(interval string) ApiGetQuoteRequest {
	r.interval = &interval
	return r
}

// Exchange where instrument is traded
func (r ApiGetQuoteRequest) Exchange(exchange string) ApiGetQuoteRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetQuoteRequest) MicCode(micCode string) ApiGetQuoteRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetQuoteRequest) Country(country string) ApiGetQuoteRequest {
	r.country = &country
	return r
}

// Number of periods for Average Volume
func (r ApiGetQuoteRequest) VolumeTimePeriod(volumeTimePeriod int64) ApiGetQuoteRequest {
	r.volumeTimePeriod = &volumeTimePeriod
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetQuoteRequest) Type_(type_ string) ApiGetQuoteRequest {
	r.type_ = &type_
	return r
}

// Value can be JSON or CSV Default JSON
func (r ApiGetQuoteRequest) Format(format string) ApiGetQuoteRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r ApiGetQuoteRequest) Delimiter(delimiter string) ApiGetQuoteRequest {
	r.delimiter = &delimiter
	return r
}

// Parameter is optional. Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r ApiGetQuoteRequest) Prepost(prepost bool) ApiGetQuoteRequest {
	r.prepost = &prepost
	return r
}

// If true, then return data for closed day
func (r ApiGetQuoteRequest) Eod(eod bool) ApiGetQuoteRequest {
	r.eod = &eod
	return r
}

// Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168].
func (r ApiGetQuoteRequest) RollingPeriod(rollingPeriod int64) ApiGetQuoteRequest {
	r.rollingPeriod = &rollingPeriod
	return r
}

// Specifies the number of decimal places for floating values Should be in range [0,11] inclusive
func (r ApiGetQuoteRequest) Dp(dp int64) ApiGetQuoteRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetQuoteRequest) Timezone(timezone string) ApiGetQuoteRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetQuoteRequest) Execute() (*GetQuote200Response, *http.Response, error) {
	return r.ApiService.GetQuoteExecute(r)
}

/*
GetQuote Quote

The quote endpoint provides real-time data for a selected financial instrument, returning essential information such as the latest price, open, high, low, close, volume, and price change. This endpoint is ideal for users needing up-to-date market data to track price movements and trading activity for specific stocks, ETFs, or other securities.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuoteRequest
*/
func (a *MarketDataAPIService) GetQuote(ctx context.Context) ApiGetQuoteRequest {
	return ApiGetQuoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetQuote200Response
func (a *MarketDataAPIService) GetQuoteExecute(r ApiGetQuoteRequest) (*GetQuote200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetQuote200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.GetQuote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/quote"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	} else {
        var defaultValue string = "1day"
        parameterAddToHeaderOrQuery(localVarQueryParams, "interval", defaultValue, "form", "")
        r.interval = &defaultValue
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.volumeTimePeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "volume_time_period", r.volumeTimePeriod, "form", "")
	} else {
        var defaultValue int64 = 9
        parameterAddToHeaderOrQuery(localVarQueryParams, "volume_time_period", defaultValue, "form", "")
        r.volumeTimePeriod = &defaultValue
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
        var defaultValue string = "JSON"
        parameterAddToHeaderOrQuery(localVarQueryParams, "format", defaultValue, "form", "")
        r.format = &defaultValue
	}
	if r.delimiter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", r.delimiter, "form", "")
	} else {
        var defaultValue string = ";"
        parameterAddToHeaderOrQuery(localVarQueryParams, "delimiter", defaultValue, "form", "")
        r.delimiter = &defaultValue
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", defaultValue, "form", "")
        r.prepost = &defaultValue
	}
	if r.eod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eod", r.eod, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "eod", defaultValue, "form", "")
        r.eod = &defaultValue
	}
	if r.rollingPeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rolling_period", r.rollingPeriod, "form", "")
	} else {
        var defaultValue int64 = 24
        parameterAddToHeaderOrQuery(localVarQueryParams, "rolling_period", defaultValue, "form", "")
        r.rollingPeriod = &defaultValue
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 5
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	} else {
        var defaultValue string = "Exchange"
        parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", defaultValue, "form", "")
        r.timezone = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorizationHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["queryParameter"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
