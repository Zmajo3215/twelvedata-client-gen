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


// FundamentalsAPIService FundamentalsAPI service
type FundamentalsAPIService service

type ApiGetBalanceSheetRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetBalanceSheetRequest) Symbol(symbol string) ApiGetBalanceSheetRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetBalanceSheetRequest) Figi(figi string) ApiGetBalanceSheetRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetBalanceSheetRequest) Isin(isin string) ApiGetBalanceSheetRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetBalanceSheetRequest) Cusip(cusip string) ApiGetBalanceSheetRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetBalanceSheetRequest) Exchange(exchange string) ApiGetBalanceSheetRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetBalanceSheetRequest) MicCode(micCode string) ApiGetBalanceSheetRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetBalanceSheetRequest) Country(country string) ApiGetBalanceSheetRequest {
	r.country = &country
	return r
}

// The reporting period for the balane sheet data
func (r ApiGetBalanceSheetRequest) Period(period string) ApiGetBalanceSheetRequest {
	r.period = &period
	return r
}

// Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetBalanceSheetRequest) StartDate(startDate string) ApiGetBalanceSheetRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetBalanceSheetRequest) EndDate(endDate string) ApiGetBalanceSheetRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetBalanceSheetRequest) Outputsize(outputsize int64) ApiGetBalanceSheetRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetBalanceSheetRequest) Execute() (*GetBalanceSheet200Response, *http.Response, error) {
	return r.ApiService.GetBalanceSheetExecute(r)
}

/*
GetBalanceSheet Balance sheet

The balance sheet endpoint provides a detailed financial statement for a company, outlining its assets, liabilities, and shareholders' equity. This endpoint returns structured data that includes current and non-current assets, total liabilities, and equity figures, enabling users to assess a company's financial health and stability.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBalanceSheetRequest
*/
func (a *FundamentalsAPIService) GetBalanceSheet(ctx context.Context) ApiGetBalanceSheetRequest {
	return ApiGetBalanceSheetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetBalanceSheet200Response
func (a *FundamentalsAPIService) GetBalanceSheetExecute(r ApiGetBalanceSheetRequest) (*GetBalanceSheet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBalanceSheet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetBalanceSheet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/balance_sheet"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	} else {
        var defaultValue string = "annual"
        parameterAddToHeaderOrQuery(localVarQueryParams, "period", defaultValue, "form", "")
        r.period = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetBalanceSheetConsolidatedRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetBalanceSheetConsolidatedRequest) Symbol(symbol string) ApiGetBalanceSheetConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetBalanceSheetConsolidatedRequest) Figi(figi string) ApiGetBalanceSheetConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetBalanceSheetConsolidatedRequest) Isin(isin string) ApiGetBalanceSheetConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetBalanceSheetConsolidatedRequest) Cusip(cusip string) ApiGetBalanceSheetConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetBalanceSheetConsolidatedRequest) Exchange(exchange string) ApiGetBalanceSheetConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetBalanceSheetConsolidatedRequest) MicCode(micCode string) ApiGetBalanceSheetConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetBalanceSheetConsolidatedRequest) Country(country string) ApiGetBalanceSheetConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the balance sheet data.
func (r ApiGetBalanceSheetConsolidatedRequest) Period(period string) ApiGetBalanceSheetConsolidatedRequest {
	r.period = &period
	return r
}

// Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetBalanceSheetConsolidatedRequest) StartDate(startDate string) ApiGetBalanceSheetConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetBalanceSheetConsolidatedRequest) EndDate(endDate string) ApiGetBalanceSheetConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetBalanceSheetConsolidatedRequest) Outputsize(outputsize int64) ApiGetBalanceSheetConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetBalanceSheetConsolidatedRequest) Execute() (*GetBalanceSheetConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetBalanceSheetConsolidatedExecute(r)
}

/*
GetBalanceSheetConsolidated Balance sheet consolidated

The balance sheet consolidated endpoint provides a detailed overview of a company's raw balance sheet, including a comprehensive summary of its assets, liabilities, and shareholders' equity. This endpoint is useful for retrieving financial data that reflects the overall financial position of a company, allowing users to access critical information about its financial health and structure.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBalanceSheetConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetBalanceSheetConsolidated(ctx context.Context) ApiGetBalanceSheetConsolidatedRequest {
	return ApiGetBalanceSheetConsolidatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetBalanceSheetConsolidated200Response
func (a *FundamentalsAPIService) GetBalanceSheetConsolidatedExecute(r ApiGetBalanceSheetConsolidatedRequest) (*GetBalanceSheetConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBalanceSheetConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetBalanceSheetConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/balance_sheet/consolidated"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	} else {
        var defaultValue string = "annual"
        parameterAddToHeaderOrQuery(localVarQueryParams, "period", defaultValue, "form", "")
        r.period = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetCashFlowRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetCashFlowRequest) Symbol(symbol string) ApiGetCashFlowRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetCashFlowRequest) Figi(figi string) ApiGetCashFlowRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetCashFlowRequest) Isin(isin string) ApiGetCashFlowRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetCashFlowRequest) Cusip(cusip string) ApiGetCashFlowRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetCashFlowRequest) Exchange(exchange string) ApiGetCashFlowRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetCashFlowRequest) MicCode(micCode string) ApiGetCashFlowRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetCashFlowRequest) Country(country string) ApiGetCashFlowRequest {
	r.country = &country
	return r
}

// The reporting period for the cash flow statements
func (r ApiGetCashFlowRequest) Period(period string) ApiGetCashFlowRequest {
	r.period = &period
	return r
}

// Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60;
func (r ApiGetCashFlowRequest) StartDate(startDate string) ApiGetCashFlowRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60;
func (r ApiGetCashFlowRequest) EndDate(endDate string) ApiGetCashFlowRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetCashFlowRequest) Outputsize(outputsize int64) ApiGetCashFlowRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetCashFlowRequest) Execute() (*GetCashFlow200Response, *http.Response, error) {
	return r.ApiService.GetCashFlowExecute(r)
}

/*
GetCashFlow Cash flow

The cash flow endpoint provides detailed information on a company's cash flow activities, including the net cash and cash equivalents moving in and out of the business. This data includes operating, investing, and financing cash flows, offering a comprehensive view of the company's liquidity and financial health.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCashFlowRequest
*/
func (a *FundamentalsAPIService) GetCashFlow(ctx context.Context) ApiGetCashFlowRequest {
	return ApiGetCashFlowRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCashFlow200Response
func (a *FundamentalsAPIService) GetCashFlowExecute(r ApiGetCashFlowRequest) (*GetCashFlow200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCashFlow200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetCashFlow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cash_flow"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	} else {
        var defaultValue string = "annual"
        parameterAddToHeaderOrQuery(localVarQueryParams, "period", defaultValue, "form", "")
        r.period = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetCashFlowConsolidatedRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetCashFlowConsolidatedRequest) Symbol(symbol string) ApiGetCashFlowConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetCashFlowConsolidatedRequest) Figi(figi string) ApiGetCashFlowConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetCashFlowConsolidatedRequest) Isin(isin string) ApiGetCashFlowConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetCashFlowConsolidatedRequest) Cusip(cusip string) ApiGetCashFlowConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetCashFlowConsolidatedRequest) Exchange(exchange string) ApiGetCashFlowConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetCashFlowConsolidatedRequest) MicCode(micCode string) ApiGetCashFlowConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetCashFlowConsolidatedRequest) Country(country string) ApiGetCashFlowConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the cash flow statements
func (r ApiGetCashFlowConsolidatedRequest) Period(period string) ApiGetCashFlowConsolidatedRequest {
	r.period = &period
	return r
}

// Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60;
func (r ApiGetCashFlowConsolidatedRequest) StartDate(startDate string) ApiGetCashFlowConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60;
func (r ApiGetCashFlowConsolidatedRequest) EndDate(endDate string) ApiGetCashFlowConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetCashFlowConsolidatedRequest) Outputsize(outputsize int64) ApiGetCashFlowConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetCashFlowConsolidatedRequest) Execute() (*GetCashFlowConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetCashFlowConsolidatedExecute(r)
}

/*
GetCashFlowConsolidated Cash flow consolidated

The cash flow consolidated endpoint provides raw data on a company's consolidated cash flow, including the net cash and cash equivalents moving in and out of the business. It returns information on operating, investing, and financing activities, helping users track liquidity and financial health over a specified period.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCashFlowConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetCashFlowConsolidated(ctx context.Context) ApiGetCashFlowConsolidatedRequest {
	return ApiGetCashFlowConsolidatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCashFlowConsolidated200Response
func (a *FundamentalsAPIService) GetCashFlowConsolidatedExecute(r ApiGetCashFlowConsolidatedRequest) (*GetCashFlowConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCashFlowConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetCashFlowConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cash_flow/consolidated"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	} else {
        var defaultValue string = "annual"
        parameterAddToHeaderOrQuery(localVarQueryParams, "period", defaultValue, "form", "")
        r.period = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetDividendsRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	range_ *string
	startDate *string
	endDate *string
	adjust *bool
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetDividendsRequest) Symbol(symbol string) ApiGetDividendsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetDividendsRequest) Figi(figi string) ApiGetDividendsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetDividendsRequest) Isin(isin string) ApiGetDividendsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetDividendsRequest) Cusip(cusip string) ApiGetDividendsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetDividendsRequest) Exchange(exchange string) ApiGetDividendsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetDividendsRequest) MicCode(micCode string) ApiGetDividendsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetDividendsRequest) Country(country string) ApiGetDividendsRequest {
	r.country = &country
	return r
}

// Specifies the time range for which to retrieve dividend data. Accepts values such as &#x60;last&#x60; (most recent dividend), &#x60;next&#x60; (upcoming dividend), &#x60;1m&#x60; - &#x60;5y&#x60; for respective periods, or &#x60;full&#x60; for all available data. If provided together with &#x60;start_date&#x60; and/or &#x60;end_date&#x60;, this parameter takes precedence.
func (r ApiGetDividendsRequest) Range_(range_ string) ApiGetDividendsRequest {
	r.range_ = &range_
	return r
}

// Start date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence.
func (r ApiGetDividendsRequest) StartDate(startDate string) ApiGetDividendsRequest {
	r.startDate = &startDate
	return r
}

// End date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence.
func (r ApiGetDividendsRequest) EndDate(endDate string) ApiGetDividendsRequest {
	r.endDate = &endDate
	return r
}

// Specifies if there should be an adjustment
func (r ApiGetDividendsRequest) Adjust(adjust bool) ApiGetDividendsRequest {
	r.adjust = &adjust
	return r
}

func (r ApiGetDividendsRequest) Execute() (*GetDividends200Response, *http.Response, error) {
	return r.ApiService.GetDividendsExecute(r)
}

/*
GetDividends Dividends

The dividends endpoint provides historical dividend data for a specified stock, in many cases covering over a decade. It returns information on dividend payouts, including the amount, payment date, and frequency. This endpoint is ideal for users tracking dividend histories or evaluating the income potential of stocks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDividendsRequest
*/
func (a *FundamentalsAPIService) GetDividends(ctx context.Context) ApiGetDividendsRequest {
	return ApiGetDividendsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetDividends200Response
func (a *FundamentalsAPIService) GetDividendsExecute(r ApiGetDividendsRequest) (*GetDividends200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDividends200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetDividends")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dividends"

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
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range", r.range_, "form", "")
	} else {
        var defaultValue string = "last"
        parameterAddToHeaderOrQuery(localVarQueryParams, "range", defaultValue, "form", "")
        r.range_ = &defaultValue
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

type ApiGetDividendsCalendarRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	startDate *string
	endDate *string
	outputsize *int64
	page *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetDividendsCalendarRequest) Symbol(symbol string) ApiGetDividendsCalendarRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetDividendsCalendarRequest) Figi(figi string) ApiGetDividendsCalendarRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetDividendsCalendarRequest) Isin(isin string) ApiGetDividendsCalendarRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetDividendsCalendarRequest) Cusip(cusip string) ApiGetDividendsCalendarRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetDividendsCalendarRequest) Exchange(exchange string) ApiGetDividendsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetDividendsCalendarRequest) MicCode(micCode string) ApiGetDividendsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetDividendsCalendarRequest) Country(country string) ApiGetDividendsCalendarRequest {
	r.country = &country
	return r
}

// Start date for the dividends calendar query. Only dividends with ex-dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;
func (r ApiGetDividendsCalendarRequest) StartDate(startDate string) ApiGetDividendsCalendarRequest {
	r.startDate = &startDate
	return r
}

// End date for the dividends calendar query. Only dividends with ex-dates on or before this date will be returned. Format &#x60;2006-01-02&#x60;
func (r ApiGetDividendsCalendarRequest) EndDate(endDate string) ApiGetDividendsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum
func (r ApiGetDividendsCalendarRequest) Outputsize(outputsize int64) ApiGetDividendsCalendarRequest {
	r.outputsize = &outputsize
	return r
}

// Page number
func (r ApiGetDividendsCalendarRequest) Page(page int64) ApiGetDividendsCalendarRequest {
	r.page = &page
	return r
}

func (r ApiGetDividendsCalendarRequest) Execute() ([]DividendsCalendarItem, *http.Response, error) {
	return r.ApiService.GetDividendsCalendarExecute(r)
}

/*
GetDividendsCalendar Dividends calendar

The dividends calendar endpoint provides a detailed schedule of upcoming and past dividend events for specified date ranges. By using the `start_date` and `end_date` parameters, users can retrieve a list of companies issuing dividends, including the ex-dividend date, payment date, and dividend amount. This endpoint is ideal for tracking dividend payouts and planning investment strategies based on dividend schedules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDividendsCalendarRequest
*/
func (a *FundamentalsAPIService) GetDividendsCalendar(ctx context.Context) ApiGetDividendsCalendarRequest {
	return ApiGetDividendsCalendarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DividendsCalendarItem
func (a *FundamentalsAPIService) GetDividendsCalendarExecute(r ApiGetDividendsCalendarRequest) ([]DividendsCalendarItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DividendsCalendarItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetDividendsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dividends_calendar"

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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 100
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
        var defaultValue int64 = 1
        parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
        r.page = &defaultValue
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

type ApiGetEarningsRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	type_ *string
	period *string
	outputsize *int64
	format *string
	delimiter *string
	startDate *string
	endDate *string
	dp *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetEarningsRequest) Symbol(symbol string) ApiGetEarningsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetEarningsRequest) Figi(figi string) ApiGetEarningsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetEarningsRequest) Isin(isin string) ApiGetEarningsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetEarningsRequest) Cusip(cusip string) ApiGetEarningsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetEarningsRequest) Exchange(exchange string) ApiGetEarningsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetEarningsRequest) MicCode(micCode string) ApiGetEarningsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetEarningsRequest) Country(country string) ApiGetEarningsRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetEarningsRequest) Type_(type_ string) ApiGetEarningsRequest {
	r.type_ = &type_
	return r
}

// Type of earning, returns only 1 record. When is not empty, dates and outputsize parameters are ignored
func (r ApiGetEarningsRequest) Period(period string) ApiGetEarningsRequest {
	r.period = &period
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;1000&#x60;. Default &#x60;10&#x60; when no date parameters are set, otherwise set to maximum
func (r ApiGetEarningsRequest) Outputsize(outputsize int64) ApiGetEarningsRequest {
	r.outputsize = &outputsize
	return r
}

// The format of the response data
func (r ApiGetEarningsRequest) Format(format string) ApiGetEarningsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetEarningsRequest) Delimiter(delimiter string) ApiGetEarningsRequest {
	r.delimiter = &delimiter
	return r
}

// The date from which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;.
func (r ApiGetEarningsRequest) StartDate(startDate string) ApiGetEarningsRequest {
	r.startDate = &startDate
	return r
}

// The date to which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;.
func (r ApiGetEarningsRequest) EndDate(endDate string) ApiGetEarningsRequest {
	r.endDate = &endDate
	return r
}

// The number of decimal places in the response data. Should be in range [0,11] inclusive
func (r ApiGetEarningsRequest) Dp(dp int64) ApiGetEarningsRequest {
	r.dp = &dp
	return r
}

func (r ApiGetEarningsRequest) Execute() (*GetEarnings200Response, *http.Response, error) {
	return r.ApiService.GetEarningsExecute(r)
}

/*
GetEarnings Earnings

The earnings endpoint provides comprehensive earnings data for a specified company, including both the estimated and actual Earnings Per Share (EPS) figures. This endpoint delivers historical earnings information, allowing users to track a company's financial performance over time.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEarningsRequest
*/
func (a *FundamentalsAPIService) GetEarnings(ctx context.Context) ApiGetEarningsRequest {
	return ApiGetEarningsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetEarnings200Response
func (a *FundamentalsAPIService) GetEarningsExecute(r ApiGetEarningsRequest) (*GetEarnings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEarnings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetEarnings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earnings"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 10
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 2
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

type ApiGetEarningsCalendarRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	exchange *string
	micCode *string
	country *string
	format *string
	delimiter *string
	startDate *string
	endDate *string
	dp *int64
}

// Exchange where instrument is traded
func (r ApiGetEarningsCalendarRequest) Exchange(exchange string) ApiGetEarningsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetEarningsCalendarRequest) MicCode(micCode string) ApiGetEarningsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetEarningsCalendarRequest) Country(country string) ApiGetEarningsCalendarRequest {
	r.country = &country
	return r
}

// Value can be JSON or CSV
func (r ApiGetEarningsCalendarRequest) Format(format string) ApiGetEarningsCalendarRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r ApiGetEarningsCalendarRequest) Delimiter(delimiter string) ApiGetEarningsCalendarRequest {
	r.delimiter = &delimiter
	return r
}

// Can be used separately and together with end_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;
func (r ApiGetEarningsCalendarRequest) StartDate(startDate string) ApiGetEarningsCalendarRequest {
	r.startDate = &startDate
	return r
}

// Can be used separately and together with start_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;
func (r ApiGetEarningsCalendarRequest) EndDate(endDate string) ApiGetEarningsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive
func (r ApiGetEarningsCalendarRequest) Dp(dp int64) ApiGetEarningsCalendarRequest {
	r.dp = &dp
	return r
}

func (r ApiGetEarningsCalendarRequest) Execute() (*GetEarningsCalendar200Response, *http.Response, error) {
	return r.ApiService.GetEarningsCalendarExecute(r)
}

/*
GetEarningsCalendar Earnings calendar

The earnings calendar endpoint provides a schedule of company earnings announcements for a specified date range. By default, it returns earnings data for the current day. Users can customize the date range using the `start_date` and `end_date` parameters to retrieve earnings information for specific periods. This endpoint is useful for tracking upcoming earnings reports and planning around key financial announcements.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEarningsCalendarRequest
*/
func (a *FundamentalsAPIService) GetEarningsCalendar(ctx context.Context) ApiGetEarningsCalendarRequest {
	return ApiGetEarningsCalendarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetEarningsCalendar200Response
func (a *FundamentalsAPIService) GetEarningsCalendarExecute(r ApiGetEarningsCalendarRequest) (*GetEarningsCalendar200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEarningsCalendar200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetEarningsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earnings_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.dp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dp", r.dp, "form", "")
	} else {
        var defaultValue int64 = 2
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

type ApiGetIncomeStatementRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetIncomeStatementRequest) Symbol(symbol string) ApiGetIncomeStatementRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetIncomeStatementRequest) Figi(figi string) ApiGetIncomeStatementRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetIncomeStatementRequest) Isin(isin string) ApiGetIncomeStatementRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetIncomeStatementRequest) Cusip(cusip string) ApiGetIncomeStatementRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetIncomeStatementRequest) Exchange(exchange string) ApiGetIncomeStatementRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetIncomeStatementRequest) MicCode(micCode string) ApiGetIncomeStatementRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetIncomeStatementRequest) Country(country string) ApiGetIncomeStatementRequest {
	r.country = &country
	return r
}

// The reporting period for the income statement data
func (r ApiGetIncomeStatementRequest) Period(period string) ApiGetIncomeStatementRequest {
	r.period = &period
	return r
}

// Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetIncomeStatementRequest) StartDate(startDate string) ApiGetIncomeStatementRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetIncomeStatementRequest) EndDate(endDate string) ApiGetIncomeStatementRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetIncomeStatementRequest) Outputsize(outputsize int64) ApiGetIncomeStatementRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetIncomeStatementRequest) Execute() (*GetIncomeStatement200Response, *http.Response, error) {
	return r.ApiService.GetIncomeStatementExecute(r)
}

/*
GetIncomeStatement Income statement

The income statement endpoint provides detailed financial data on a company's income statement, including revenues, expenses, and net income for specified periods, either annually or quarterly. This endpoint is essential for retrieving comprehensive financial performance metrics of a company, allowing users to access historical and current financial results.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIncomeStatementRequest
*/
func (a *FundamentalsAPIService) GetIncomeStatement(ctx context.Context) ApiGetIncomeStatementRequest {
	return ApiGetIncomeStatementRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetIncomeStatement200Response
func (a *FundamentalsAPIService) GetIncomeStatementExecute(r ApiGetIncomeStatementRequest) (*GetIncomeStatement200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetIncomeStatement200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIncomeStatement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income_statement"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetIncomeStatementConsolidatedRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	period *string
	startDate *string
	endDate *string
	outputsize *int64
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetIncomeStatementConsolidatedRequest) Symbol(symbol string) ApiGetIncomeStatementConsolidatedRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetIncomeStatementConsolidatedRequest) Figi(figi string) ApiGetIncomeStatementConsolidatedRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetIncomeStatementConsolidatedRequest) Isin(isin string) ApiGetIncomeStatementConsolidatedRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetIncomeStatementConsolidatedRequest) Cusip(cusip string) ApiGetIncomeStatementConsolidatedRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetIncomeStatementConsolidatedRequest) Exchange(exchange string) ApiGetIncomeStatementConsolidatedRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetIncomeStatementConsolidatedRequest) MicCode(micCode string) ApiGetIncomeStatementConsolidatedRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetIncomeStatementConsolidatedRequest) Country(country string) ApiGetIncomeStatementConsolidatedRequest {
	r.country = &country
	return r
}

// The reporting period for the income statement data
func (r ApiGetIncomeStatementConsolidatedRequest) Period(period string) ApiGetIncomeStatementConsolidatedRequest {
	r.period = &period
	return r
}

// Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetIncomeStatementConsolidatedRequest) StartDate(startDate string) ApiGetIncomeStatementConsolidatedRequest {
	r.startDate = &startDate
	return r
}

// End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetIncomeStatementConsolidatedRequest) EndDate(endDate string) ApiGetIncomeStatementConsolidatedRequest {
	r.endDate = &endDate
	return r
}

// Number of records in response
func (r ApiGetIncomeStatementConsolidatedRequest) Outputsize(outputsize int64) ApiGetIncomeStatementConsolidatedRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetIncomeStatementConsolidatedRequest) Execute() (*GetIncomeStatementConsolidated200Response, *http.Response, error) {
	return r.ApiService.GetIncomeStatementConsolidatedExecute(r)
}

/*
GetIncomeStatementConsolidated Income statement consolidated

The income statement consolidated endpoint provides a company's raw income statement, detailing revenue, expenses, and net income for specified periods, either annually or quarterly. This data is essential for evaluating a company's financial performance over time, allowing users to access comprehensive financial results in a structured format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIncomeStatementConsolidatedRequest
*/
func (a *FundamentalsAPIService) GetIncomeStatementConsolidated(ctx context.Context) ApiGetIncomeStatementConsolidatedRequest {
	return ApiGetIncomeStatementConsolidatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetIncomeStatementConsolidated200Response
func (a *FundamentalsAPIService) GetIncomeStatementConsolidatedExecute(r ApiGetIncomeStatementConsolidatedRequest) (*GetIncomeStatementConsolidated200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetIncomeStatementConsolidated200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIncomeStatementConsolidated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/income_statement/consolidated"

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
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 6
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetIpoCalendarRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	exchange *string
	micCode *string
	country *string
	startDate *string
	endDate *string
}

// Exchange where instrument is traded
func (r ApiGetIpoCalendarRequest) Exchange(exchange string) ApiGetIpoCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetIpoCalendarRequest) MicCode(micCode string) ApiGetIpoCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetIpoCalendarRequest) Country(country string) ApiGetIpoCalendarRequest {
	r.country = &country
	return r
}

// The earliest IPO date to include in the results. Format: &#x60;2006-01-02&#x60;
func (r ApiGetIpoCalendarRequest) StartDate(startDate string) ApiGetIpoCalendarRequest {
	r.startDate = &startDate
	return r
}

// The latest IPO date to include in the results. Format: &#x60;2006-01-02&#x60;
func (r ApiGetIpoCalendarRequest) EndDate(endDate string) ApiGetIpoCalendarRequest {
	r.endDate = &endDate
	return r
}

func (r ApiGetIpoCalendarRequest) Execute() (*map[string][]GetIpoCalendar200ResponseValueInner, *http.Response, error) {
	return r.ApiService.GetIpoCalendarExecute(r)
}

/*
GetIpoCalendar IPO calendar

The IPO Calendar endpoint provides detailed information on initial public offerings (IPOs), including those that have occurred in the past, are happening today, or are scheduled for the future. Users can access data such as company names, IPO dates, and offering details, allowing them to track and monitor IPO activity efficiently.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIpoCalendarRequest
*/
func (a *FundamentalsAPIService) GetIpoCalendar(ctx context.Context) ApiGetIpoCalendarRequest {
	return ApiGetIpoCalendarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string][]GetIpoCalendar200ResponseValueInner
func (a *FundamentalsAPIService) GetIpoCalendarExecute(r ApiGetIpoCalendarRequest) (*map[string][]GetIpoCalendar200ResponseValueInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]GetIpoCalendar200ResponseValueInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetIpoCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ipo_calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
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

type ApiGetKeyExecutivesRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetKeyExecutivesRequest) Symbol(symbol string) ApiGetKeyExecutivesRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetKeyExecutivesRequest) Figi(figi string) ApiGetKeyExecutivesRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetKeyExecutivesRequest) Isin(isin string) ApiGetKeyExecutivesRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetKeyExecutivesRequest) Cusip(cusip string) ApiGetKeyExecutivesRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetKeyExecutivesRequest) Exchange(exchange string) ApiGetKeyExecutivesRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetKeyExecutivesRequest) MicCode(micCode string) ApiGetKeyExecutivesRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetKeyExecutivesRequest) Country(country string) ApiGetKeyExecutivesRequest {
	r.country = &country
	return r
}

func (r ApiGetKeyExecutivesRequest) Execute() (*GetKeyExecutives200Response, *http.Response, error) {
	return r.ApiService.GetKeyExecutivesExecute(r)
}

/*
GetKeyExecutives Key executives

The key executives endpoint provides detailed information about a company's key executives identified by a specific stock symbol. It returns data such as names, titles, and roles of the executives, which can be useful for understanding the leadership structure of the company.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKeyExecutivesRequest
*/
func (a *FundamentalsAPIService) GetKeyExecutives(ctx context.Context) ApiGetKeyExecutivesRequest {
	return ApiGetKeyExecutivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetKeyExecutives200Response
func (a *FundamentalsAPIService) GetKeyExecutivesExecute(r ApiGetKeyExecutivesRequest) (*GetKeyExecutives200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetKeyExecutives200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetKeyExecutives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key_executives"

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

type ApiGetLastChangesRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	endpoint string
	startDate *string
	symbol *string
	exchange *string
	micCode *string
	country *string
	page *int64
	outputsize *int64
}

// The starting date and time for data selection, in &#x60;2006-01-02T15:04:05&#x60; format
func (r ApiGetLastChangesRequest) StartDate(startDate string) ApiGetLastChangesRequest {
	r.startDate = &startDate
	return r
}

// Filter by symbol
func (r ApiGetLastChangesRequest) Symbol(symbol string) ApiGetLastChangesRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name
func (r ApiGetLastChangesRequest) Exchange(exchange string) ApiGetLastChangesRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetLastChangesRequest) MicCode(micCode string) ApiGetLastChangesRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetLastChangesRequest) Country(country string) ApiGetLastChangesRequest {
	r.country = &country
	return r
}

// Page number
func (r ApiGetLastChangesRequest) Page(page int64) ApiGetLastChangesRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r ApiGetLastChangesRequest) Outputsize(outputsize int64) ApiGetLastChangesRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetLastChangesRequest) Execute() (*GetLastChanges200Response, *http.Response, error) {
	return r.ApiService.GetLastChangesExecute(r)
}

/*
GetLastChanges Last changes

The last change endpoint provides the most recent updates to fundamental data for a specified dataset. It returns a timestamp indicating when the data was last modified, allowing users to efficiently manage API requests by only fetching new data when changes occur. This helps optimize data retrieval and reduce unnecessary API credit usage.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param endpoint Endpoint name
 @return ApiGetLastChangesRequest
*/
func (a *FundamentalsAPIService) GetLastChanges(ctx context.Context, endpoint string) ApiGetLastChangesRequest {
	return ApiGetLastChangesRequest{
		ApiService: a,
		ctx: ctx,
		endpoint: endpoint,
	}
}

// Execute executes the request
//  @return GetLastChanges200Response
func (a *FundamentalsAPIService) GetLastChangesExecute(r ApiGetLastChangesRequest) (*GetLastChanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetLastChanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetLastChanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/last_change/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", url.PathEscape(parameterValueToString(r.endpoint, "endpoint")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
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
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
        var defaultValue int64 = 1
        parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
        r.page = &defaultValue
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 30
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetLogoRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	exchange *string
	micCode *string
	country *string
}

// The ticker symbol of an instrument for which data is requested, e.g., &#x60;AAPL&#x60;, &#x60;BTC/USD&#x60;, &#x60;EUR/USD&#x60;.
func (r ApiGetLogoRequest) Symbol(symbol string) ApiGetLogoRequest {
	r.symbol = &symbol
	return r
}

// The exchange name where the instrument is traded, e.g., &#x60;NASDAQ&#x60;, &#x60;NSE&#x60;
func (r ApiGetLogoRequest) Exchange(exchange string) ApiGetLogoRequest {
	r.exchange = &exchange
	return r
}

// The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., &#x60;XNAS&#x60;, &#x60;XLON&#x60;
func (r ApiGetLogoRequest) MicCode(micCode string) ApiGetLogoRequest {
	r.micCode = &micCode
	return r
}

// The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetLogoRequest) Country(country string) ApiGetLogoRequest {
	r.country = &country
	return r
}

func (r ApiGetLogoRequest) Execute() (*GetLogo200Response, *http.Response, error) {
	return r.ApiService.GetLogoExecute(r)
}

/*
GetLogo Logo

The logo endpoint provides the official logo image for a specified company, cryptocurrency, or forex pair. This endpoint is useful for integrating visual branding elements into financial applications, websites, or reports, ensuring that users can easily identify and associate the correct logo with the respective financial asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLogoRequest
*/
func (a *FundamentalsAPIService) GetLogo(ctx context.Context) ApiGetLogoRequest {
	return ApiGetLogoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetLogo200Response
func (a *FundamentalsAPIService) GetLogoExecute(r ApiGetLogoRequest) (*GetLogo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetLogo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetLogo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMarketCapRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	startDate *string
	endDate *string
	page *int64
	outputsize *int64
}

// Filter by symbol
func (r ApiGetMarketCapRequest) Symbol(symbol string) ApiGetMarketCapRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMarketCapRequest) Figi(figi string) ApiGetMarketCapRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMarketCapRequest) Isin(isin string) ApiGetMarketCapRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMarketCapRequest) Cusip(cusip string) ApiGetMarketCapRequest {
	r.cusip = &cusip
	return r
}

// Filter by exchange name
func (r ApiGetMarketCapRequest) Exchange(exchange string) ApiGetMarketCapRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetMarketCapRequest) MicCode(micCode string) ApiGetMarketCapRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMarketCapRequest) Country(country string) ApiGetMarketCapRequest {
	r.country = &country
	return r
}

// Start date for market capitalization data retrieval. Data will be returned from this date onwards. Format &#x60;2006-01-02&#x60;
func (r ApiGetMarketCapRequest) StartDate(startDate string) ApiGetMarketCapRequest {
	r.startDate = &startDate
	return r
}

// End date for market capitalization data retrieval. Data will be returned up to and including this date. Format &#x60;2006-01-02&#x60;
func (r ApiGetMarketCapRequest) EndDate(endDate string) ApiGetMarketCapRequest {
	r.endDate = &endDate
	return r
}

// Page number
func (r ApiGetMarketCapRequest) Page(page int64) ApiGetMarketCapRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r ApiGetMarketCapRequest) Outputsize(outputsize int64) ApiGetMarketCapRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetMarketCapRequest) Execute() (*GetMarketCap200Response, *http.Response, error) {
	return r.ApiService.GetMarketCapExecute(r)
}

/*
GetMarketCap Market capitalization

The Market Capitalization History endpoint provides historical data on a company's market capitalization over a specified time period. It returns a time series of market cap values, allowing users to track changes in a company's market value.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMarketCapRequest
*/
func (a *FundamentalsAPIService) GetMarketCap(ctx context.Context) ApiGetMarketCapRequest {
	return ApiGetMarketCapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMarketCap200Response
func (a *FundamentalsAPIService) GetMarketCapExecute(r ApiGetMarketCapRequest) (*GetMarketCap200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMarketCap200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetMarketCap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_cap"

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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
        var defaultValue int64 = 1
        parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
        r.page = &defaultValue
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 10
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
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

type ApiGetProfileRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetProfileRequest) Symbol(symbol string) ApiGetProfileRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetProfileRequest) Figi(figi string) ApiGetProfileRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetProfileRequest) Isin(isin string) ApiGetProfileRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetProfileRequest) Cusip(cusip string) ApiGetProfileRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetProfileRequest) Exchange(exchange string) ApiGetProfileRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetProfileRequest) MicCode(micCode string) ApiGetProfileRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetProfileRequest) Country(country string) ApiGetProfileRequest {
	r.country = &country
	return r
}

func (r ApiGetProfileRequest) Execute() (*GetProfile200Response, *http.Response, error) {
	return r.ApiService.GetProfileExecute(r)
}

/*
GetProfile Profile

The profile endpoint provides detailed company information, including the company's name, industry, sector, CEO, headquarters location, and market capitalization. This data is useful for obtaining a comprehensive overview of a company's business and financial standing.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProfileRequest
*/
func (a *FundamentalsAPIService) GetProfile(ctx context.Context) ApiGetProfileRequest {
	return ApiGetProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetProfile200Response
func (a *FundamentalsAPIService) GetProfileExecute(r ApiGetProfileRequest) (*GetProfile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetProfile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profile"

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

type ApiGetSplitsRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	range_ *string
	startDate *string
	endDate *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetSplitsRequest) Symbol(symbol string) ApiGetSplitsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetSplitsRequest) Figi(figi string) ApiGetSplitsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetSplitsRequest) Isin(isin string) ApiGetSplitsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetSplitsRequest) Cusip(cusip string) ApiGetSplitsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetSplitsRequest) Exchange(exchange string) ApiGetSplitsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetSplitsRequest) MicCode(micCode string) ApiGetSplitsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetSplitsRequest) Country(country string) ApiGetSplitsRequest {
	r.country = &country
	return r
}

// Range of data to be returned
func (r ApiGetSplitsRequest) Range_(range_ string) ApiGetSplitsRequest {
	r.range_ = &range_
	return r
}

// The starting date for data selection. Format &#x60;2006-01-02&#x60;
func (r ApiGetSplitsRequest) StartDate(startDate string) ApiGetSplitsRequest {
	r.startDate = &startDate
	return r
}

// The ending date for data selection. Format &#x60;2006-01-02&#x60;
func (r ApiGetSplitsRequest) EndDate(endDate string) ApiGetSplitsRequest {
	r.endDate = &endDate
	return r
}

func (r ApiGetSplitsRequest) Execute() (*GetSplits200Response, *http.Response, error) {
	return r.ApiService.GetSplitsExecute(r)
}

/*
GetSplits Splits

The splits endpoint provides historical data on stock split events for a specified company. It returns details including the date of each split and the corresponding split factor.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSplitsRequest
*/
func (a *FundamentalsAPIService) GetSplits(ctx context.Context) ApiGetSplitsRequest {
	return ApiGetSplitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSplits200Response
func (a *FundamentalsAPIService) GetSplitsExecute(r ApiGetSplitsRequest) (*GetSplits200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSplits200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetSplits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/splits"

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
	if r.range_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range", r.range_, "form", "")
	} else {
        var defaultValue string = "last"
        parameterAddToHeaderOrQuery(localVarQueryParams, "range", defaultValue, "form", "")
        r.range_ = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
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

type ApiGetSplitsCalendarRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
	startDate *string
	endDate *string
	outputsize *int64
	page *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetSplitsCalendarRequest) Symbol(symbol string) ApiGetSplitsCalendarRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetSplitsCalendarRequest) Figi(figi string) ApiGetSplitsCalendarRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetSplitsCalendarRequest) Isin(isin string) ApiGetSplitsCalendarRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetSplitsCalendarRequest) Cusip(cusip string) ApiGetSplitsCalendarRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetSplitsCalendarRequest) Exchange(exchange string) ApiGetSplitsCalendarRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetSplitsCalendarRequest) MicCode(micCode string) ApiGetSplitsCalendarRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetSplitsCalendarRequest) Country(country string) ApiGetSplitsCalendarRequest {
	r.country = &country
	return r
}

// The starting date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60;
func (r ApiGetSplitsCalendarRequest) StartDate(startDate string) ApiGetSplitsCalendarRequest {
	r.startDate = &startDate
	return r
}

// The ending date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60;
func (r ApiGetSplitsCalendarRequest) EndDate(endDate string) ApiGetSplitsCalendarRequest {
	r.endDate = &endDate
	return r
}

// Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum
func (r ApiGetSplitsCalendarRequest) Outputsize(outputsize int64) ApiGetSplitsCalendarRequest {
	r.outputsize = &outputsize
	return r
}

// Page number
func (r ApiGetSplitsCalendarRequest) Page(page string) ApiGetSplitsCalendarRequest {
	r.page = &page
	return r
}

func (r ApiGetSplitsCalendarRequest) Execute() ([]SplitsCalendarResponseItem, *http.Response, error) {
	return r.ApiService.GetSplitsCalendarExecute(r)
}

/*
GetSplitsCalendar Splits calendar

The splits calendar endpoint provides a detailed calendar of stock split events within a specified date range. By setting the `start_date` and `end_date` parameters, users can retrieve a list of upcoming or past stock splits, including the company name, split ratio, and effective date. This endpoint is useful for tracking changes in stock structure and planning investment strategies around these events.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSplitsCalendarRequest
*/
func (a *FundamentalsAPIService) GetSplitsCalendar(ctx context.Context) ApiGetSplitsCalendarRequest {
	return ApiGetSplitsCalendarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SplitsCalendarResponseItem
func (a *FundamentalsAPIService) GetSplitsCalendarExecute(r ApiGetSplitsCalendarRequest) ([]SplitsCalendarResponseItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SplitsCalendarResponseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetSplitsCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/splits_calendar"

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
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 100
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
        var defaultValue string = "1"
        parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
        r.page = &defaultValue
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

type ApiGetStatisticsRequest struct {
	ctx context.Context
	ApiService *FundamentalsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	country *string
}

// Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct
func (r ApiGetStatisticsRequest) Symbol(symbol string) ApiGetStatisticsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetStatisticsRequest) Figi(figi string) ApiGetStatisticsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetStatisticsRequest) Isin(isin string) ApiGetStatisticsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetStatisticsRequest) Cusip(cusip string) ApiGetStatisticsRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded
func (r ApiGetStatisticsRequest) Exchange(exchange string) ApiGetStatisticsRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard
func (r ApiGetStatisticsRequest) MicCode(micCode string) ApiGetStatisticsRequest {
	r.micCode = &micCode
	return r
}

// Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetStatisticsRequest) Country(country string) ApiGetStatisticsRequest {
	r.country = &country
	return r
}

func (r ApiGetStatisticsRequest) Execute() (*GetStatistics200Response, *http.Response, error) {
	return r.ApiService.GetStatisticsExecute(r)
}

/*
GetStatistics Statistics

The statistics endpoint provides a comprehensive snapshot of a company's key financial statistics, including valuation metrics, revenue figures, profit margins, and other essential financial data. This endpoint is ideal for users seeking detailed insights into a company's financial health and performance metrics.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStatisticsRequest
*/
func (a *FundamentalsAPIService) GetStatistics(ctx context.Context) ApiGetStatisticsRequest {
	return ApiGetStatisticsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetStatistics200Response
func (a *FundamentalsAPIService) GetStatisticsExecute(r ApiGetStatisticsRequest) (*GetStatistics200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStatistics200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FundamentalsAPIService.GetStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/statistics"

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
