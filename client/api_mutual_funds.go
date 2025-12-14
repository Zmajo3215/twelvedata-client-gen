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


// MutualFundsAPIService MutualFundsAPI service
type MutualFundsAPIService service

type ApiGetMutualFundsWorldRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldRequest) Symbol(symbol string) ApiGetMutualFundsWorldRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldRequest) Figi(figi string) ApiGetMutualFundsWorldRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldRequest) Isin(isin string) ApiGetMutualFundsWorldRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldRequest) Cusip(cusip string) ApiGetMutualFundsWorldRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldRequest) Country(country string) ApiGetMutualFundsWorldRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldRequest) Dp(dp int64) ApiGetMutualFundsWorldRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldRequest) Execute() (*GetMutualFundsWorld200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldExecute(r)
}

/*
GetMutualFundsWorld MF full data

The mutual full data endpoint provides detailed information about global mutual funds. It returns a comprehensive dataset that includes a summary of the fund, its performance metrics, risk assessment, ratings, asset composition, purchase details, and sustainability factors. This endpoint is essential for users seeking in-depth insights into mutual funds on a global scale, allowing them to evaluate various aspects such as investment performance, risk levels, and environmental impact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorld(ctx context.Context) ApiGetMutualFundsWorldRequest {
	return ApiGetMutualFundsWorldRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorld200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldExecute(r ApiGetMutualFundsWorldRequest) (*GetMutualFundsWorld200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorld200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorld")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldCompositionRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldCompositionRequest) Symbol(symbol string) ApiGetMutualFundsWorldCompositionRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldCompositionRequest) Figi(figi string) ApiGetMutualFundsWorldCompositionRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldCompositionRequest) Isin(isin string) ApiGetMutualFundsWorldCompositionRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldCompositionRequest) Cusip(cusip string) ApiGetMutualFundsWorldCompositionRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldCompositionRequest) Country(country string) ApiGetMutualFundsWorldCompositionRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldCompositionRequest) Dp(dp int64) ApiGetMutualFundsWorldCompositionRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldCompositionRequest) Execute() (*GetMutualFundsWorldComposition200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldCompositionExecute(r)
}

/*
GetMutualFundsWorldComposition Composition

The mutual funds compositions endpoint provides detailed information about the portfolio composition of a specified mutual fund. It returns data on sector allocations, individual holdings, and their respective weighted exposures. This endpoint is useful for users seeking to understand the investment distribution and risk profile of a mutual fund.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldCompositionRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldComposition(ctx context.Context) ApiGetMutualFundsWorldCompositionRequest {
	return ApiGetMutualFundsWorldCompositionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldComposition200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldCompositionExecute(r ApiGetMutualFundsWorldCompositionRequest) (*GetMutualFundsWorldComposition200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldComposition200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldComposition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/composition"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldPerformanceRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldPerformanceRequest) Symbol(symbol string) ApiGetMutualFundsWorldPerformanceRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldPerformanceRequest) Figi(figi string) ApiGetMutualFundsWorldPerformanceRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldPerformanceRequest) Isin(isin string) ApiGetMutualFundsWorldPerformanceRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldPerformanceRequest) Cusip(cusip string) ApiGetMutualFundsWorldPerformanceRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldPerformanceRequest) Country(country string) ApiGetMutualFundsWorldPerformanceRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldPerformanceRequest) Dp(dp int64) ApiGetMutualFundsWorldPerformanceRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldPerformanceRequest) Execute() (*GetMutualFundsWorldPerformance200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldPerformanceExecute(r)
}

/*
GetMutualFundsWorldPerformance Performance

The mutual funds performances endpoint provides comprehensive performance data for mutual funds globally. It returns metrics such as trailing returns, annual returns, quarterly returns, and load-adjusted returns.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldPerformanceRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldPerformance(ctx context.Context) ApiGetMutualFundsWorldPerformanceRequest {
	return ApiGetMutualFundsWorldPerformanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldPerformance200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldPerformanceExecute(r ApiGetMutualFundsWorldPerformanceRequest) (*GetMutualFundsWorldPerformance200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldPerformance200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldPerformance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/performance"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldPurchaseInfoRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Symbol(symbol string) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Figi(figi string) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Isin(isin string) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Cusip(cusip string) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Country(country string) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Dp(dp int64) ApiGetMutualFundsWorldPurchaseInfoRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldPurchaseInfoRequest) Execute() (*GetMutualFundsWorldPurchaseInfo200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldPurchaseInfoExecute(r)
}

/*
GetMutualFundsWorldPurchaseInfo Purchase info

The mutual funds purchase information endpoint provides detailed purchasing details for global mutual funds. It returns data on minimum investment requirements, current pricing, and a list of brokerages where the mutual fund can be purchased. This endpoint is useful for users looking to understand the entry requirements and options available for investing in specific mutual funds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldPurchaseInfoRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldPurchaseInfo(ctx context.Context) ApiGetMutualFundsWorldPurchaseInfoRequest {
	return ApiGetMutualFundsWorldPurchaseInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldPurchaseInfo200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldPurchaseInfoExecute(r ApiGetMutualFundsWorldPurchaseInfoRequest) (*GetMutualFundsWorldPurchaseInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldPurchaseInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldPurchaseInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/purchase_info"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldRatingsRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldRatingsRequest) Symbol(symbol string) ApiGetMutualFundsWorldRatingsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldRatingsRequest) Figi(figi string) ApiGetMutualFundsWorldRatingsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldRatingsRequest) Isin(isin string) ApiGetMutualFundsWorldRatingsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldRatingsRequest) Cusip(cusip string) ApiGetMutualFundsWorldRatingsRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldRatingsRequest) Country(country string) ApiGetMutualFundsWorldRatingsRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldRatingsRequest) Dp(dp int64) ApiGetMutualFundsWorldRatingsRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldRatingsRequest) Execute() (*GetMutualFundsWorldRatings200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldRatingsExecute(r)
}

/*
GetMutualFundsWorldRatings Ratings

The mutual funds ratings endpoint provides detailed ratings for mutual funds across global markets. It returns data on the performance and quality of mutual funds, including ratings calculated in-house by Twelve Data and from various financial institutions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldRatingsRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldRatings(ctx context.Context) ApiGetMutualFundsWorldRatingsRequest {
	return ApiGetMutualFundsWorldRatingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldRatings200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldRatingsExecute(r ApiGetMutualFundsWorldRatingsRequest) (*GetMutualFundsWorldRatings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldRatings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldRatings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/ratings"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldRiskRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldRiskRequest) Symbol(symbol string) ApiGetMutualFundsWorldRiskRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldRiskRequest) Figi(figi string) ApiGetMutualFundsWorldRiskRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldRiskRequest) Isin(isin string) ApiGetMutualFundsWorldRiskRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldRiskRequest) Cusip(cusip string) ApiGetMutualFundsWorldRiskRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldRiskRequest) Country(country string) ApiGetMutualFundsWorldRiskRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldRiskRequest) Dp(dp int64) ApiGetMutualFundsWorldRiskRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldRiskRequest) Execute() (*GetMutualFundsWorldRisk200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldRiskExecute(r)
}

/*
GetMutualFundsWorldRisk Risk

The mutual funds risk endpoint provides detailed risk metrics for global mutual funds. It returns data such as standard deviation, beta, and Sharpe ratio, which help assess the volatility and risk profile of mutual funds across different markets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldRiskRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldRisk(ctx context.Context) ApiGetMutualFundsWorldRiskRequest {
	return ApiGetMutualFundsWorldRiskRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldRisk200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldRiskExecute(r ApiGetMutualFundsWorldRiskRequest) (*GetMutualFundsWorldRisk200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldRisk200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldRisk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/risk"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldSummaryRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldSummaryRequest) Symbol(symbol string) ApiGetMutualFundsWorldSummaryRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldSummaryRequest) Figi(figi string) ApiGetMutualFundsWorldSummaryRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldSummaryRequest) Isin(isin string) ApiGetMutualFundsWorldSummaryRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldSummaryRequest) Cusip(cusip string) ApiGetMutualFundsWorldSummaryRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldSummaryRequest) Country(country string) ApiGetMutualFundsWorldSummaryRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldSummaryRequest) Dp(dp int64) ApiGetMutualFundsWorldSummaryRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldSummaryRequest) Execute() (*GetMutualFundsWorldSummary200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldSummaryExecute(r)
}

/*
GetMutualFundsWorldSummary Summary

The mutual funds summary endpoint provides a concise overview of global mutual funds, including key details such as fund name, symbol, asset class, and region. This endpoint is useful for quickly obtaining essential information about various mutual funds worldwide, aiding in the comparison and selection of funds for investment portfolios.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldSummaryRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldSummary(ctx context.Context) ApiGetMutualFundsWorldSummaryRequest {
	return ApiGetMutualFundsWorldSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldSummary200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldSummaryExecute(r ApiGetMutualFundsWorldSummaryRequest) (*GetMutualFundsWorldSummary200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldSummary200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/summary"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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

type ApiGetMutualFundsWorldSustainabilityRequest struct {
	ctx context.Context
	ApiService *MutualFundsAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	country *string
	dp *int64
}

// Symbol ticker of mutual fund
func (r ApiGetMutualFundsWorldSustainabilityRequest) Symbol(symbol string) ApiGetMutualFundsWorldSustainabilityRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsWorldSustainabilityRequest) Figi(figi string) ApiGetMutualFundsWorldSustainabilityRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsWorldSustainabilityRequest) Isin(isin string) ApiGetMutualFundsWorldSustainabilityRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsWorldSustainabilityRequest) Cusip(cusip string) ApiGetMutualFundsWorldSustainabilityRequest {
	r.cusip = &cusip
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsWorldSustainabilityRequest) Country(country string) ApiGetMutualFundsWorldSustainabilityRequest {
	r.country = &country
	return r
}

// Number of decimal places for floating values. Accepts value in range [0,11]
func (r ApiGetMutualFundsWorldSustainabilityRequest) Dp(dp int64) ApiGetMutualFundsWorldSustainabilityRequest {
	r.dp = &dp
	return r
}

func (r ApiGetMutualFundsWorldSustainabilityRequest) Execute() (*GetMutualFundsWorldSustainability200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsWorldSustainabilityExecute(r)
}

/*
GetMutualFundsWorldSustainability Sustainability

The mutual funds sustainability endpoint provides detailed information on the sustainability and Environmental, Social, and Governance (ESG) ratings of global mutual funds. It returns data such as ESG scores, sustainability metrics, and fund identifiers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsWorldSustainabilityRequest
*/
func (a *MutualFundsAPIService) GetMutualFundsWorldSustainability(ctx context.Context) ApiGetMutualFundsWorldSustainabilityRequest {
	return ApiGetMutualFundsWorldSustainabilityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsWorldSustainability200Response
func (a *MutualFundsAPIService) GetMutualFundsWorldSustainabilityExecute(r ApiGetMutualFundsWorldSustainabilityRequest) (*GetMutualFundsWorldSustainability200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsWorldSustainability200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MutualFundsAPIService.GetMutualFundsWorldSustainability")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/world/sustainability"

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
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
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
