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
)


// TimeSeriesAPIService TimeSeriesAPI service
type TimeSeriesAPIService service

type ApiGetTimeSeriesRequest struct {
	ctx context.Context
	ApiService *TimeSeriesAPIService
	interval *string
	symbol *string
	isin *string
	figi *string
	cusip *string
	outputsize *int64
	exchange *string
	micCode *string
	country *string
	type_ *string
	timezone *string
	startDate *string
	endDate *string
	date *string
	order *string
	prepost *bool
	format *string
	delimiter *string
	dp *int64
	previousClose *bool
	adjust *string
}

// Interval between two consecutive points in time series
func (r ApiGetTimeSeriesRequest) Interval(interval string) ApiGetTimeSeriesRequest {
	r.interval = &interval
	return r
}

// Symbol ticker of the instrument. E.g. &#x60;AAPL&#x60;, &#x60;EUR/USD&#x60;, &#x60;ETH/BTC&#x60;, ...
func (r ApiGetTimeSeriesRequest) Symbol(symbol string) ApiGetTimeSeriesRequest {
	r.symbol = &symbol
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetTimeSeriesRequest) Isin(isin string) ApiGetTimeSeriesRequest {
	r.isin = &isin
	return r
}

// The FIGI of an instrument for which data is requested
func (r ApiGetTimeSeriesRequest) Figi(figi string) ApiGetTimeSeriesRequest {
	r.figi = &figi
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetTimeSeriesRequest) Cusip(cusip string) ApiGetTimeSeriesRequest {
	r.cusip = &cusip
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum
func (r ApiGetTimeSeriesRequest) Outputsize(outputsize int64) ApiGetTimeSeriesRequest {
	r.outputsize = &outputsize
	return r
}

// Exchange where instrument is traded
func (r ApiGetTimeSeriesRequest) Exchange(exchange string) ApiGetTimeSeriesRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetTimeSeriesRequest) MicCode(micCode string) ApiGetTimeSeriesRequest {
	r.micCode = &micCode
	return r
}

// The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetTimeSeriesRequest) Country(country string) ApiGetTimeSeriesRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetTimeSeriesRequest) Type_(type_ string) ApiGetTimeSeriesRequest {
	r.type_ = &type_
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetTimeSeriesRequest) Timezone(timezone string) ApiGetTimeSeriesRequest {
	r.timezone = &timezone
	return r
}

// Can be used separately and together with &#x60;end_date&#x60;. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;  Default location: &lt;ul&gt; &lt;li&gt;Forex and Cryptocurrencies - &lt;code&gt;UTC&lt;/code&gt;&lt;/li&gt; &lt;li&gt;Stocks - where exchange is located (e.g. for AAPL it will be &lt;code&gt;America/New_York&lt;/code&gt;)&lt;/li&gt; &lt;/ul&gt; Both parameters take into account if &lt;code&gt;timezone&lt;/code&gt; parameter is provided.&lt;br/&gt; If &lt;code&gt;timezone&lt;/code&gt; is given then, &lt;code&gt;start_date&lt;/code&gt; and &lt;code&gt;end_date&lt;/code&gt; will be used in the specified location  Examples: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;&amp;symbol&#x3D;AAPL&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 New York time up to current date&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;&amp;symbol&#x3D;EUR/USD&amp;timezone&#x3D;Asia/Singapore&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date&lt;/li&gt; &lt;li&gt;3. &lt;code&gt;&amp;symbol&#x3D;ETH/BTC&amp;timezone&#x3D;Europe/Zurich&amp;start_date&#x3D;2019-08-09T15:50:00&amp;end_date&#x3D;2019-08-09T15:55:00&amp;...&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00&lt;/li&gt; &lt;/ul&gt;
func (r ApiGetTimeSeriesRequest) StartDate(startDate string) ApiGetTimeSeriesRequest {
	r.startDate = &startDate
	return r
}

// The ending date and time for data selection, see &#x60;start_date&#x60; description for details.
func (r ApiGetTimeSeriesRequest) EndDate(endDate string) ApiGetTimeSeriesRequest {
	r.endDate = &endDate
	return r
}

// Specifies the exact date to get the data for. Could be the exact date, e.g. &#x60;2021-10-27&#x60;, or in human language &#x60;today&#x60; or &#x60;yesterday&#x60;
func (r ApiGetTimeSeriesRequest) Date(date string) ApiGetTimeSeriesRequest {
	r.date = &date
	return r
}

// Sorting order of the output
func (r ApiGetTimeSeriesRequest) Order(order string) ApiGetTimeSeriesRequest {
	r.order = &order
	return r
}

// Returns quotes that include pre-market and post-market data. Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume
func (r ApiGetTimeSeriesRequest) Prepost(prepost bool) ApiGetTimeSeriesRequest {
	r.prepost = &prepost
	return r
}

// The format of the response data
func (r ApiGetTimeSeriesRequest) Format(format string) ApiGetTimeSeriesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetTimeSeriesRequest) Delimiter(delimiter string) ApiGetTimeSeriesRequest {
	r.delimiter = &delimiter
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided
func (r ApiGetTimeSeriesRequest) Dp(dp int64) ApiGetTimeSeriesRequest {
	r.dp = &dp
	return r
}

// A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object
func (r ApiGetTimeSeriesRequest) PreviousClose(previousClose bool) ApiGetTimeSeriesRequest {
	r.previousClose = &previousClose
	return r
}

// Adjusting mode for prices
func (r ApiGetTimeSeriesRequest) Adjust(adjust string) ApiGetTimeSeriesRequest {
	r.adjust = &adjust
	return r
}

func (r ApiGetTimeSeriesRequest) Execute() (*GetTimeSeries200Response, *http.Response, error) {
	return r.ApiService.GetTimeSeriesExecute(r)
}

/*
GetTimeSeries Time series

The time series endpoint provides detailed historical data for a specified financial instrument. It returns two main components: metadata, which includes essential information about the instrument, and a time series dataset. The time series consists of chronological entries with Open, High, Low, and Close prices, and for applicable instruments, it also includes trading volume. This endpoint is ideal for retrieving comprehensive historical price data for analysis or visualization purposes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTimeSeriesRequest
*/
func (a *TimeSeriesAPIService) GetTimeSeries(ctx context.Context) ApiGetTimeSeriesRequest {
	return ApiGetTimeSeriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTimeSeries200Response
func (a *TimeSeriesAPIService) GetTimeSeriesExecute(r ApiGetTimeSeriesRequest) (*GetTimeSeries200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTimeSeries200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeSeriesAPIService.GetTimeSeries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time_series"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.isin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isin", r.isin, "form", "")
	}
	if r.figi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "figi", r.figi, "form", "")
	}
	if r.cusip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cusip", r.cusip, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 30
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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
	if r.timezone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", r.timezone, "form", "")
	} else {
        var defaultValue string = "Exchange"
        parameterAddToHeaderOrQuery(localVarQueryParams, "timezone", defaultValue, "form", "")
        r.timezone = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	} else {
        var defaultValue string = "desc"
        parameterAddToHeaderOrQuery(localVarQueryParams, "order", defaultValue, "form", "")
        r.order = &defaultValue
	}
	if r.prepost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", r.prepost, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "prepost", defaultValue, "form", "")
        r.prepost = &defaultValue
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
        var defaultValue int64 = -1
        parameterAddToHeaderOrQuery(localVarQueryParams, "dp", defaultValue, "form", "")
        r.dp = &defaultValue
	}
	if r.previousClose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "previous_close", r.previousClose, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "previous_close", defaultValue, "form", "")
        r.previousClose = &defaultValue
	}
	if r.adjust != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adjust", r.adjust, "form", "")
	} else {
        var defaultValue string = "splits"
        parameterAddToHeaderOrQuery(localVarQueryParams, "adjust", defaultValue, "form", "")
        r.adjust = &defaultValue
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

type ApiGetTimeSeriesCrossRequest struct {
	ctx context.Context
	ApiService *TimeSeriesAPIService
	base *string
	quote *string
	interval *string
	baseType *string
	baseExchange *string
	baseMicCode *string
	quoteType *string
	quoteExchange *string
	quoteMicCode *string
	outputsize *int64
	format *string
	delimiter *string
	prepost *bool
	startDate *string
	endDate *string
	adjust *bool
	dp *int64
	timezone *string
}

// Base currency symbol
func (r ApiGetTimeSeriesCrossRequest) Base(base string) ApiGetTimeSeriesCrossRequest {
	r.base = &base
	return r
}

// Quote currency symbol
func (r ApiGetTimeSeriesCrossRequest) Quote(quote string) ApiGetTimeSeriesCrossRequest {
	r.quote = &quote
	return r
}

// Interval between two consecutive points in time series
func (r ApiGetTimeSeriesCrossRequest) Interval(interval string) ApiGetTimeSeriesCrossRequest {
	r.interval = &interval
	return r
}

// Base instrument type according to the &#x60;/instrument_type&#x60; endpoint
func (r ApiGetTimeSeriesCrossRequest) BaseType(baseType string) ApiGetTimeSeriesCrossRequest {
	r.baseType = &baseType
	return r
}

// Base exchange
func (r ApiGetTimeSeriesCrossRequest) BaseExchange(baseExchange string) ApiGetTimeSeriesCrossRequest {
	r.baseExchange = &baseExchange
	return r
}

// Base MIC code
func (r ApiGetTimeSeriesCrossRequest) BaseMicCode(baseMicCode string) ApiGetTimeSeriesCrossRequest {
	r.baseMicCode = &baseMicCode
	return r
}

// Quote instrument type according to the &#x60;/instrument_type&#x60; endpoint
func (r ApiGetTimeSeriesCrossRequest) QuoteType(quoteType string) ApiGetTimeSeriesCrossRequest {
	r.quoteType = &quoteType
	return r
}

// Quote exchange
func (r ApiGetTimeSeriesCrossRequest) QuoteExchange(quoteExchange string) ApiGetTimeSeriesCrossRequest {
	r.quoteExchange = &quoteExchange
	return r
}

// Quote MIC code
func (r ApiGetTimeSeriesCrossRequest) QuoteMicCode(quoteMicCode string) ApiGetTimeSeriesCrossRequest {
	r.quoteMicCode = &quoteMicCode
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum
func (r ApiGetTimeSeriesCrossRequest) Outputsize(outputsize int64) ApiGetTimeSeriesCrossRequest {
	r.outputsize = &outputsize
	return r
}

// Format of the response data
func (r ApiGetTimeSeriesCrossRequest) Format(format string) ApiGetTimeSeriesCrossRequest {
	r.format = &format
	return r
}

// Delimiter used in CSV file
func (r ApiGetTimeSeriesCrossRequest) Delimiter(delimiter string) ApiGetTimeSeriesCrossRequest {
	r.delimiter = &delimiter
	return r
}

// Only for &#x60;Pro&#x60; and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume.
func (r ApiGetTimeSeriesCrossRequest) Prepost(prepost bool) ApiGetTimeSeriesCrossRequest {
	r.prepost = &prepost
	return r
}

// Start date for the time series data
func (r ApiGetTimeSeriesCrossRequest) StartDate(startDate string) ApiGetTimeSeriesCrossRequest {
	r.startDate = &startDate
	return r
}

// End date for the time series data
func (r ApiGetTimeSeriesCrossRequest) EndDate(endDate string) ApiGetTimeSeriesCrossRequest {
	r.endDate = &endDate
	return r
}

// Specifies if there should be an adjustment
func (r ApiGetTimeSeriesCrossRequest) Adjust(adjust bool) ApiGetTimeSeriesCrossRequest {
	r.adjust = &adjust
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive.
func (r ApiGetTimeSeriesCrossRequest) Dp(dp int64) ApiGetTimeSeriesCrossRequest {
	r.dp = &dp
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetTimeSeriesCrossRequest) Timezone(timezone string) ApiGetTimeSeriesCrossRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetTimeSeriesCrossRequest) Execute() (*GetTimeSeriesCross200Response, *http.Response, error) {
	return r.ApiService.GetTimeSeriesCrossExecute(r)
}

/*
GetTimeSeriesCross Time series cross

The Time Series Cross endpoint calculates and returns historical cross-rate data for exotic forex pairs, cryptocurrencies, or stocks (e.g., Apple Inc. price in Indian Rupees) on the fly. It provides metadata about the requested symbol and a time series array with Open, High, Low, and Close prices, sorted descending by time, enabling analysis of price history and market trends.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTimeSeriesCrossRequest
*/
func (a *TimeSeriesAPIService) GetTimeSeriesCross(ctx context.Context) ApiGetTimeSeriesCrossRequest {
	return ApiGetTimeSeriesCrossRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTimeSeriesCross200Response
func (a *TimeSeriesAPIService) GetTimeSeriesCrossExecute(r ApiGetTimeSeriesCrossRequest) (*GetTimeSeriesCross200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTimeSeriesCross200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeSeriesAPIService.GetTimeSeriesCross")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time_series/cross"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.base == nil {
		return localVarReturnValue, nil, reportError("base is required and must be specified")
	}
	if r.quote == nil {
		return localVarReturnValue, nil, reportError("quote is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "base", r.base, "form", "")
	if r.baseType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_type", r.baseType, "form", "")
	}
	if r.baseExchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_exchange", r.baseExchange, "form", "")
	}
	if r.baseMicCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_mic_code", r.baseMicCode, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "quote", r.quote, "form", "")
	if r.quoteType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_type", r.quoteType, "form", "")
	}
	if r.quoteExchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_exchange", r.quoteExchange, "form", "")
	}
	if r.quoteMicCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quote_mic_code", r.quoteMicCode, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.adjust != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adjust", r.adjust, "form", "")
	} else {
        var defaultValue bool = true
        parameterAddToHeaderOrQuery(localVarQueryParams, "adjust", defaultValue, "form", "")
        r.adjust = &defaultValue
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
