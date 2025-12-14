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


// ReferenceDataAPIService ReferenceDataAPI service
type ReferenceDataAPIService service

type ApiGetBondsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	exchange *string
	country *string
	format *string
	delimiter *string
	showPlan *bool
	page *int64
	outputsize *int64
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetBondsRequest) Symbol(symbol string) ApiGetBondsRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name
func (r ApiGetBondsRequest) Exchange(exchange string) ApiGetBondsRequest {
	r.exchange = &exchange
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetBondsRequest) Country(country string) ApiGetBondsRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ApiGetBondsRequest) Format(format string) ApiGetBondsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetBondsRequest) Delimiter(delimiter string) ApiGetBondsRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ApiGetBondsRequest) ShowPlan(showPlan bool) ApiGetBondsRequest {
	r.showPlan = &showPlan
	return r
}

// Page number of the results to fetch
func (r ApiGetBondsRequest) Page(page int64) ApiGetBondsRequest {
	r.page = &page
	return r
}

// Determines the number of data points returned in the output
func (r ApiGetBondsRequest) Outputsize(outputsize int64) ApiGetBondsRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetBondsRequest) Execute() (*GetBonds200Response, *http.Response, error) {
	return r.ApiService.GetBondsExecute(r)
}

/*
GetBonds Fixed income

The fixed income endpoint provides a daily updated list of available bonds. It returns an array containing detailed information about each bond, including identifiers, names, and other relevant attributes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBondsRequest
*/
func (a *ReferenceDataAPIService) GetBonds(ctx context.Context) ApiGetBondsRequest {
	return ApiGetBondsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetBonds200Response
func (a *ReferenceDataAPIService) GetBondsExecute(r ApiGetBondsRequest) (*GetBonds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBonds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetBonds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bonds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
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
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
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
        var defaultValue int64 = 5000
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

type ApiGetCommoditiesRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	category *string
	format *string
	delimiter *string
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetCommoditiesRequest) Symbol(symbol string) ApiGetCommoditiesRequest {
	r.symbol = &symbol
	return r
}

// Filter by category of commodity
func (r ApiGetCommoditiesRequest) Category(category string) ApiGetCommoditiesRequest {
	r.category = &category
	return r
}

// The format of the response data
func (r ApiGetCommoditiesRequest) Format(format string) ApiGetCommoditiesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetCommoditiesRequest) Delimiter(delimiter string) ApiGetCommoditiesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ApiGetCommoditiesRequest) Execute() (*GetCommodities200Response, *http.Response, error) {
	return r.ApiService.GetCommoditiesExecute(r)
}

/*
GetCommodities Commodities

The commodities endpoint provides a daily updated list of available commodity pairs, across precious metals, livestock, softs, grains, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCommoditiesRequest
*/
func (a *ReferenceDataAPIService) GetCommodities(ctx context.Context) ApiGetCommoditiesRequest {
	return ApiGetCommoditiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCommodities200Response
func (a *ReferenceDataAPIService) GetCommoditiesExecute(r ApiGetCommoditiesRequest) (*GetCommodities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCommodities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCommodities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commodities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "form", "")
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

type ApiGetCountriesRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
}

func (r ApiGetCountriesRequest) Execute() (*GetCountries200Response, *http.Response, error) {
	return r.ApiService.GetCountriesExecute(r)
}

/*
GetCountries Countries

The countries endpoint provides a comprehensive list of countries, including their ISO codes, official names, capitals, and currencies. This data is essential for applications requiring accurate country information for tasks such as localization, currency conversion, or geographic analysis.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCountriesRequest
*/
func (a *ReferenceDataAPIService) GetCountries(ctx context.Context) ApiGetCountriesRequest {
	return ApiGetCountriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCountries200Response
func (a *ReferenceDataAPIService) GetCountriesExecute(r ApiGetCountriesRequest) (*GetCountries200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCountries200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCountries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/countries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetCrossListingsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	exchange *string
	micCode *string
	country *string
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetCrossListingsRequest) Symbol(symbol string) ApiGetCrossListingsRequest {
	r.symbol = &symbol
	return r
}

// Exchange where instrument is traded
func (r ApiGetCrossListingsRequest) Exchange(exchange string) ApiGetCrossListingsRequest {
	r.exchange = &exchange
	return r
}

// Market identifier code (MIC) under ISO 10383 standard
func (r ApiGetCrossListingsRequest) MicCode(micCode string) ApiGetCrossListingsRequest {
	r.micCode = &micCode
	return r
}

// Country to which stock exchange belongs to
func (r ApiGetCrossListingsRequest) Country(country string) ApiGetCrossListingsRequest {
	r.country = &country
	return r
}

func (r ApiGetCrossListingsRequest) Execute() (*GetCrossListings200Response, *http.Response, error) {
	return r.ApiService.GetCrossListingsExecute(r)
}

/*
GetCrossListings Cross listings

The cross_listings endpoint provides a daily updated list of cross-listed symbols for a specified financial instrument. Cross-listed symbols represent the same security available on multiple exchanges. This endpoint is useful for identifying all the exchanges where a particular security is traded, allowing users to access comprehensive trading information across different markets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCrossListingsRequest
*/
func (a *ReferenceDataAPIService) GetCrossListings(ctx context.Context) ApiGetCrossListingsRequest {
	return ApiGetCrossListingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCrossListings200Response
func (a *ReferenceDataAPIService) GetCrossListingsExecute(r ApiGetCrossListingsRequest) (*GetCrossListings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCrossListings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCrossListings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cross_listings"

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

type ApiGetCryptocurrenciesRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	exchange *string
	currencyBase *string
	currencyQuote *string
	format *string
	delimiter *string
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetCryptocurrenciesRequest) Symbol(symbol string) ApiGetCryptocurrenciesRequest {
	r.symbol = &symbol
	return r
}

// Filter by exchange name. E.g. &#x60;Binance&#x60;, &#x60;Coinbase&#x60;, etc.
func (r ApiGetCryptocurrenciesRequest) Exchange(exchange string) ApiGetCryptocurrenciesRequest {
	r.exchange = &exchange
	return r
}

// Filter by currency base
func (r ApiGetCryptocurrenciesRequest) CurrencyBase(currencyBase string) ApiGetCryptocurrenciesRequest {
	r.currencyBase = &currencyBase
	return r
}

// Filter by currency quote
func (r ApiGetCryptocurrenciesRequest) CurrencyQuote(currencyQuote string) ApiGetCryptocurrenciesRequest {
	r.currencyQuote = &currencyQuote
	return r
}

// The format of the response data
func (r ApiGetCryptocurrenciesRequest) Format(format string) ApiGetCryptocurrenciesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetCryptocurrenciesRequest) Delimiter(delimiter string) ApiGetCryptocurrenciesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ApiGetCryptocurrenciesRequest) Execute() (*GetCryptocurrencies200Response, *http.Response, error) {
	return r.ApiService.GetCryptocurrenciesExecute(r)
}

/*
GetCryptocurrencies Cryptocurrency pairs

The cryptocurrencies endpoint provides a daily updated list of all available cryptos. It returns an array containing detailed information about each cryptocurrency, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of cryptocurrencies for applications that require up-to-date market listings or need to display available crypto assets to users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCryptocurrenciesRequest
*/
func (a *ReferenceDataAPIService) GetCryptocurrencies(ctx context.Context) ApiGetCryptocurrenciesRequest {
	return ApiGetCryptocurrenciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCryptocurrencies200Response
func (a *ReferenceDataAPIService) GetCryptocurrenciesExecute(r ApiGetCryptocurrenciesRequest) (*GetCryptocurrencies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCryptocurrencies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCryptocurrencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cryptocurrencies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.currencyBase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_base", r.currencyBase, "form", "")
	}
	if r.currencyQuote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_quote", r.currencyQuote, "form", "")
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

type ApiGetCryptocurrencyExchangesRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	format *string
	delimiter *string
}

// The format of the response data
func (r ApiGetCryptocurrencyExchangesRequest) Format(format string) ApiGetCryptocurrencyExchangesRequest {
	r.format = &format
	return r
}

// Specify the delimiter used when downloading the CSV file
func (r ApiGetCryptocurrencyExchangesRequest) Delimiter(delimiter string) ApiGetCryptocurrencyExchangesRequest {
	r.delimiter = &delimiter
	return r
}

func (r ApiGetCryptocurrencyExchangesRequest) Execute() (*GetCryptocurrencyExchanges200Response, *http.Response, error) {
	return r.ApiService.GetCryptocurrencyExchangesExecute(r)
}

/*
GetCryptocurrencyExchanges Cryptocurrency exchanges

The cryptocurrency exchanges endpoint provides a daily updated list of available cryptocurrency exchanges. It returns an array containing details about each exchange, such as exchange names and identifiers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCryptocurrencyExchangesRequest
*/
func (a *ReferenceDataAPIService) GetCryptocurrencyExchanges(ctx context.Context) ApiGetCryptocurrencyExchangesRequest {
	return ApiGetCryptocurrencyExchangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetCryptocurrencyExchanges200Response
func (a *ReferenceDataAPIService) GetCryptocurrencyExchangesExecute(r ApiGetCryptocurrencyExchangesRequest) (*GetCryptocurrencyExchanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCryptocurrencyExchanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetCryptocurrencyExchanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cryptocurrency_exchanges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetETFsFamilyRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	country *string
	fundFamily *string
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetETFsFamilyRequest) Country(country string) ApiGetETFsFamilyRequest {
	r.country = &country
	return r
}

// Filter by investment company that manages the fund
func (r ApiGetETFsFamilyRequest) FundFamily(fundFamily string) ApiGetETFsFamilyRequest {
	r.fundFamily = &fundFamily
	return r
}

func (r ApiGetETFsFamilyRequest) Execute() (*GetETFsFamily200Response, *http.Response, error) {
	return r.ApiService.GetETFsFamilyExecute(r)
}

/*
GetETFsFamily ETFs families

Retrieve a comprehensive list of exchange-traded fund (ETF) families, providing users with detailed information on various ETF groups available in the market. This endpoint is ideal for users looking to explore different ETF categories, compare offerings, or integrate ETF family data into their financial applications.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetETFsFamilyRequest
*/
func (a *ReferenceDataAPIService) GetETFsFamily(ctx context.Context) ApiGetETFsFamilyRequest {
	return ApiGetETFsFamilyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetETFsFamily200Response
func (a *ReferenceDataAPIService) GetETFsFamilyExecute(r ApiGetETFsFamilyRequest) (*GetETFsFamily200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetETFsFamily200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetETFsFamily")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/etfs/family"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.fundFamily != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_family", r.fundFamily, "form", "")
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

type ApiGetETFsListRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	cik *string
	country *string
	fundFamily *string
	fundType *string
	page *int64
	outputsize *int64
}

// Filter by symbol
func (r ApiGetETFsListRequest) Symbol(symbol string) ApiGetETFsListRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetETFsListRequest) Figi(figi string) ApiGetETFsListRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetETFsListRequest) Isin(isin string) ApiGetETFsListRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetETFsListRequest) Cusip(cusip string) ApiGetETFsListRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ApiGetETFsListRequest) Cik(cik string) ApiGetETFsListRequest {
	r.cik = &cik
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetETFsListRequest) Country(country string) ApiGetETFsListRequest {
	r.country = &country
	return r
}

// Filter by investment company that manages the fund
func (r ApiGetETFsListRequest) FundFamily(fundFamily string) ApiGetETFsListRequest {
	r.fundFamily = &fundFamily
	return r
}

// Filter by the type of fund
func (r ApiGetETFsListRequest) FundType(fundType string) ApiGetETFsListRequest {
	r.fundType = &fundType
	return r
}

// Page number
func (r ApiGetETFsListRequest) Page(page int64) ApiGetETFsListRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r ApiGetETFsListRequest) Outputsize(outputsize int64) ApiGetETFsListRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetETFsListRequest) Execute() (*GetETFsList200Response, *http.Response, error) {
	return r.ApiService.GetETFsListExecute(r)
}

/*
GetETFsList ETFs directory

The ETFs directory endpoint provides a daily updated list of exchange-traded funds, sorted by total assets in descending order. This endpoint is useful for retrieving comprehensive ETF data, including fund names and asset values, to assist users in quickly identifying the ETFs available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetETFsListRequest
*/
func (a *ReferenceDataAPIService) GetETFsList(ctx context.Context) ApiGetETFsListRequest {
	return ApiGetETFsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetETFsList200Response
func (a *ReferenceDataAPIService) GetETFsListExecute(r ApiGetETFsListRequest) (*GetETFsList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetETFsList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetETFsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/etfs/list"

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
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.fundFamily != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_family", r.fundFamily, "form", "")
	}
	if r.fundType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_type", r.fundType, "form", "")
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
        var defaultValue int64 = 50
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

type ApiGetETFsTypeRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	country *string
	fundType *string
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetETFsTypeRequest) Country(country string) ApiGetETFsTypeRequest {
	r.country = &country
	return r
}

// Filter by the type of fund
func (r ApiGetETFsTypeRequest) FundType(fundType string) ApiGetETFsTypeRequest {
	r.fundType = &fundType
	return r
}

func (r ApiGetETFsTypeRequest) Execute() (*GetETFsType200Response, *http.Response, error) {
	return r.ApiService.GetETFsTypeExecute(r)
}

/*
GetETFsType ETFs types

The ETFs Types endpoint provides a concise list of ETF categories by market (e.g., Singapore, United States), including types like "Equity Precious Metals" and "Large Blend." It supports targeted investment research and portfolio diversification.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetETFsTypeRequest
*/
func (a *ReferenceDataAPIService) GetETFsType(ctx context.Context) ApiGetETFsTypeRequest {
	return ApiGetETFsTypeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetETFsType200Response
func (a *ReferenceDataAPIService) GetETFsTypeExecute(r ApiGetETFsTypeRequest) (*GetETFsType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetETFsType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetETFsType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/etfs/type"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.fundType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_type", r.fundType, "form", "")
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

type ApiGetEarliestTimestampRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	interval *string
	symbol *string
	figi *string
	isin *string
	cusip *string
	exchange *string
	micCode *string
	timezone *string
}

// Interval between two consecutive points in time series.
func (r ApiGetEarliestTimestampRequest) Interval(interval string) ApiGetEarliestTimestampRequest {
	r.interval = &interval
	return r
}

// Symbol ticker of the instrument.
func (r ApiGetEarliestTimestampRequest) Symbol(symbol string) ApiGetEarliestTimestampRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI).
func (r ApiGetEarliestTimestampRequest) Figi(figi string) ApiGetEarliestTimestampRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetEarliestTimestampRequest) Isin(isin string) ApiGetEarliestTimestampRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetEarliestTimestampRequest) Cusip(cusip string) ApiGetEarliestTimestampRequest {
	r.cusip = &cusip
	return r
}

// Exchange where instrument is traded.
func (r ApiGetEarliestTimestampRequest) Exchange(exchange string) ApiGetEarliestTimestampRequest {
	r.exchange = &exchange
	return r
}

// Market Identifier Code (MIC) under ISO 10383 standard.
func (r ApiGetEarliestTimestampRequest) MicCode(micCode string) ApiGetEarliestTimestampRequest {
	r.micCode = &micCode
	return r
}

// Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt;
func (r ApiGetEarliestTimestampRequest) Timezone(timezone string) ApiGetEarliestTimestampRequest {
	r.timezone = &timezone
	return r
}

func (r ApiGetEarliestTimestampRequest) Execute() (*GetEarliestTimestamp200Response, *http.Response, error) {
	return r.ApiService.GetEarliestTimestampExecute(r)
}

/*
GetEarliestTimestamp Earliest timestamp

The earliest_timestamp endpoint provides the earliest available date and time for a specified financial instrument at a given data interval. This endpoint is useful for determining the starting point of historical data availability for various assets, such as stocks or currencies, allowing users to understand the time range covered by the data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEarliestTimestampRequest
*/
func (a *ReferenceDataAPIService) GetEarliestTimestamp(ctx context.Context) ApiGetEarliestTimestampRequest {
	return ApiGetEarliestTimestampRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetEarliestTimestamp200Response
func (a *ReferenceDataAPIService) GetEarliestTimestampExecute(r ApiGetEarliestTimestampRequest) (*GetEarliestTimestamp200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEarliestTimestamp200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetEarliestTimestamp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/earliest_timestamp"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
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

type ApiGetEtfRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	cik *string
	exchange *string
	micCode *string
	country *string
	format *string
	delimiter *string
	showPlan *bool
	includeDelisted *bool
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetEtfRequest) Symbol(symbol string) ApiGetEtfRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetEtfRequest) Figi(figi string) ApiGetEtfRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetEtfRequest) Isin(isin string) ApiGetEtfRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetEtfRequest) Cusip(cusip string) ApiGetEtfRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ApiGetEtfRequest) Cik(cik string) ApiGetEtfRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ApiGetEtfRequest) Exchange(exchange string) ApiGetEtfRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetEtfRequest) MicCode(micCode string) ApiGetEtfRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetEtfRequest) Country(country string) ApiGetEtfRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ApiGetEtfRequest) Format(format string) ApiGetEtfRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetEtfRequest) Delimiter(delimiter string) ApiGetEtfRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ApiGetEtfRequest) ShowPlan(showPlan bool) ApiGetEtfRequest {
	r.showPlan = &showPlan
	return r
}

// Include delisted identifiers
func (r ApiGetEtfRequest) IncludeDelisted(includeDelisted bool) ApiGetEtfRequest {
	r.includeDelisted = &includeDelisted
	return r
}

func (r ApiGetEtfRequest) Execute() (*GetEtf200Response, *http.Response, error) {
	return r.ApiService.GetEtfExecute(r)
}

/*
GetEtf ETFs

The ETFs endpoint provides a daily updated list of all available Exchange-Traded Funds. It returns an array containing detailed information about each ETF, including its symbol, name, and other relevant identifiers. This endpoint is useful for retrieving a comprehensive catalog of ETFs for portfolio management, investment tracking, or financial analysis.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEtfRequest
*/
func (a *ReferenceDataAPIService) GetEtf(ctx context.Context) ApiGetEtfRequest {
	return ApiGetEtfRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetEtf200Response
func (a *ReferenceDataAPIService) GetEtfExecute(r ApiGetEtfRequest) (*GetEtf200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetEtf200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetEtf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/etfs"

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
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
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
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
	}
	if r.includeDelisted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", r.includeDelisted, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", defaultValue, "form", "")
        r.includeDelisted = &defaultValue
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

type ApiGetExchangeScheduleRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	micName *string
	micCode *string
	country *string
	date *string
}

// Filter by exchange name
func (r ApiGetExchangeScheduleRequest) MicName(micName string) ApiGetExchangeScheduleRequest {
	r.micName = &micName
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetExchangeScheduleRequest) MicCode(micCode string) ApiGetExchangeScheduleRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetExchangeScheduleRequest) Country(country string) ApiGetExchangeScheduleRequest {
	r.country = &country
	return r
}

// &lt;p&gt; If a date is provided, the API returns the schedule for the specified date; otherwise, it returns the default (common) schedule. &lt;/p&gt; The date can be specified in one of the following formats: &lt;ul&gt; &lt;li&gt;An exact date (e.g., &lt;code&gt;2021-10-27&lt;/code&gt;)&lt;/li&gt; &lt;li&gt;A human-readable keyword: &lt;code&gt;today&lt;/code&gt; or &lt;code&gt;yesterday&lt;/code&gt;&lt;/li&gt; &lt;li&gt;A full datetime string in UTC (e.g., &lt;code&gt;2025-04-11T20:00:00&lt;/code&gt;) to retrieve the schedule corresponding to the day in the specified time.&lt;/li&gt; &lt;/ul&gt; When using a datetime value, the resulting schedule will correspond to the local calendar day at the specified time. For example, &lt;code&gt;2025-04-11T20:00:00 UTC&lt;/code&gt; corresponds to: &lt;ul&gt; &lt;li&gt;&lt;code&gt;2025-04-11&lt;/code&gt; in the &lt;code&gt;America/New_York&lt;/code&gt; timezone&lt;/li&gt; &lt;li&gt;&lt;code&gt;2025-04-12&lt;/code&gt; in the &lt;code&gt;Australia/Sydney&lt;/code&gt; timezone&lt;/li&gt; &lt;/ul&gt;
func (r ApiGetExchangeScheduleRequest) Date(date string) ApiGetExchangeScheduleRequest {
	r.date = &date
	return r
}

func (r ApiGetExchangeScheduleRequest) Execute() (*GetExchangeSchedule200Response, *http.Response, error) {
	return r.ApiService.GetExchangeScheduleExecute(r)
}

/*
GetExchangeSchedule Exchanges schedule

The exchanges schedule endpoint provides detailed information about various stock exchanges, including their trading hours and operational days. This data is essential for users who need to know when specific exchanges are open for trading, allowing them to plan their activities around the availability of these markets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetExchangeScheduleRequest
*/
func (a *ReferenceDataAPIService) GetExchangeSchedule(ctx context.Context) ApiGetExchangeScheduleRequest {
	return ApiGetExchangeScheduleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetExchangeSchedule200Response
func (a *ReferenceDataAPIService) GetExchangeScheduleExecute(r ApiGetExchangeScheduleRequest) (*GetExchangeSchedule200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchangeSchedule200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetExchangeSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchange_schedule"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.micName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_name", r.micName, "form", "")
	}
	if r.micCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mic_code", r.micCode, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
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

type ApiGetExchangesRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	type_ *string
	name *string
	code *string
	country *string
	format *string
	delimiter *string
	showPlan *bool
}

// The asset class to which the instrument belongs
func (r ApiGetExchangesRequest) Type_(type_ string) ApiGetExchangesRequest {
	r.type_ = &type_
	return r
}

// Filter by exchange name
func (r ApiGetExchangesRequest) Name(name string) ApiGetExchangesRequest {
	r.name = &name
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetExchangesRequest) Code(code string) ApiGetExchangesRequest {
	r.code = &code
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetExchangesRequest) Country(country string) ApiGetExchangesRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ApiGetExchangesRequest) Format(format string) ApiGetExchangesRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetExchangesRequest) Delimiter(delimiter string) ApiGetExchangesRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ApiGetExchangesRequest) ShowPlan(showPlan bool) ApiGetExchangesRequest {
	r.showPlan = &showPlan
	return r
}

func (r ApiGetExchangesRequest) Execute() (*GetExchanges200Response, *http.Response, error) {
	return r.ApiService.GetExchangesExecute(r)
}

/*
GetExchanges Exchanges

The exchanges endpoint provides a comprehensive list of all available equity exchanges. It returns an array containing detailed information about each exchange, such as exchange code, name, country, and timezone. This data is updated daily.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetExchangesRequest
*/
func (a *ReferenceDataAPIService) GetExchanges(ctx context.Context) ApiGetExchangesRequest {
	return ApiGetExchangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetExchanges200Response
func (a *ReferenceDataAPIService) GetExchangesExecute(r ApiGetExchangesRequest) (*GetExchanges200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchanges200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetExchanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/exchanges"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
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
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
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

type ApiGetForexPairsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	currencyBase *string
	currencyQuote *string
	format *string
	delimiter *string
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetForexPairsRequest) Symbol(symbol string) ApiGetForexPairsRequest {
	r.symbol = &symbol
	return r
}

// Filter by currency base
func (r ApiGetForexPairsRequest) CurrencyBase(currencyBase string) ApiGetForexPairsRequest {
	r.currencyBase = &currencyBase
	return r
}

// Filter by currency quote
func (r ApiGetForexPairsRequest) CurrencyQuote(currencyQuote string) ApiGetForexPairsRequest {
	r.currencyQuote = &currencyQuote
	return r
}

// The format of the response data
func (r ApiGetForexPairsRequest) Format(format string) ApiGetForexPairsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetForexPairsRequest) Delimiter(delimiter string) ApiGetForexPairsRequest {
	r.delimiter = &delimiter
	return r
}

func (r ApiGetForexPairsRequest) Execute() (*GetForexPairs200Response, *http.Response, error) {
	return r.ApiService.GetForexPairsExecute(r)
}

/*
GetForexPairs Forex pairs

The forex pairs endpoint provides a comprehensive list of all available foreign exchange currency pairs. It returns an array of forex pairs, which is updated daily.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetForexPairsRequest
*/
func (a *ReferenceDataAPIService) GetForexPairs(ctx context.Context) ApiGetForexPairsRequest {
	return ApiGetForexPairsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetForexPairs200Response
func (a *ReferenceDataAPIService) GetForexPairsExecute(r ApiGetForexPairsRequest) (*GetForexPairs200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetForexPairs200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetForexPairs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/forex_pairs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.currencyBase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_base", r.currencyBase, "form", "")
	}
	if r.currencyQuote != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency_quote", r.currencyQuote, "form", "")
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

type ApiGetFundsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	cik *string
	exchange *string
	country *string
	format *string
	delimiter *string
	showPlan *bool
	page *int64
	outputsize *int64
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetFundsRequest) Symbol(symbol string) ApiGetFundsRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetFundsRequest) Figi(figi string) ApiGetFundsRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetFundsRequest) Isin(isin string) ApiGetFundsRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetFundsRequest) Cusip(cusip string) ApiGetFundsRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ApiGetFundsRequest) Cik(cik string) ApiGetFundsRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ApiGetFundsRequest) Exchange(exchange string) ApiGetFundsRequest {
	r.exchange = &exchange
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetFundsRequest) Country(country string) ApiGetFundsRequest {
	r.country = &country
	return r
}

// The format of the response data
func (r ApiGetFundsRequest) Format(format string) ApiGetFundsRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetFundsRequest) Delimiter(delimiter string) ApiGetFundsRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ApiGetFundsRequest) ShowPlan(showPlan bool) ApiGetFundsRequest {
	r.showPlan = &showPlan
	return r
}

// Page number of the results to fetch
func (r ApiGetFundsRequest) Page(page int64) ApiGetFundsRequest {
	r.page = &page
	return r
}

// Determines the number of data points returned in the output
func (r ApiGetFundsRequest) Outputsize(outputsize int64) ApiGetFundsRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetFundsRequest) Execute() (*GetFunds200Response, *http.Response, error) {
	return r.ApiService.GetFundsExecute(r)
}

/*
GetFunds Funds

The funds endpoint provides a daily updated list of available investment funds. It returns an array containing detailed information about each fund, including identifiers, names, and other relevant attributes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFundsRequest
*/
func (a *ReferenceDataAPIService) GetFunds(ctx context.Context) ApiGetFundsRequest {
	return ApiGetFundsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetFunds200Response
func (a *ReferenceDataAPIService) GetFundsExecute(r ApiGetFundsRequest) (*GetFunds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFunds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetFunds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/funds"

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
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
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
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
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
        var defaultValue int64 = 5000
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

type ApiGetInstrumentTypeRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
}

func (r ApiGetInstrumentTypeRequest) Execute() (*GetInstrumentType200Response, *http.Response, error) {
	return r.ApiService.GetInstrumentTypeExecute(r)
}

/*
GetInstrumentType Instrument type

The instrument type endpoint lists all available financial instrument types, such as stocks, ETFs, and cryptos. This information is essential for users to identify and categorize different financial instruments when accessing or analyzing market data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInstrumentTypeRequest
*/
func (a *ReferenceDataAPIService) GetInstrumentType(ctx context.Context) ApiGetInstrumentTypeRequest {
	return ApiGetInstrumentTypeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInstrumentType200Response
func (a *ReferenceDataAPIService) GetInstrumentTypeExecute(r ApiGetInstrumentTypeRequest) (*GetInstrumentType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInstrumentType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetInstrumentType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/instrument_type"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetIntervalsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
}

func (r ApiGetIntervalsRequest) Execute() (*GetIntervals200Response, *http.Response, error) {
	return r.ApiService.GetIntervalsExecute(r)
}

/*
GetIntervals Intervals List

The intervals endpoint provides a list of supported time intervals that can be used for querying financial data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIntervalsRequest
*/
func (a *ReferenceDataAPIService) GetIntervals(ctx context.Context) ApiGetIntervalsRequest {
	return ApiGetIntervalsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetIntervals200Response
func (a *ReferenceDataAPIService) GetIntervalsExecute(r ApiGetIntervalsRequest) (*GetIntervals200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetIntervals200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetIntervals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intervals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetMarketStateRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	exchange *string
	code *string
	country *string
}

// The exchange name where the instrument is traded.
func (r ApiGetMarketStateRequest) Exchange(exchange string) ApiGetMarketStateRequest {
	r.exchange = &exchange
	return r
}

// The Market Identifier Code (MIC) of the exchange where the instrument is traded.
func (r ApiGetMarketStateRequest) Code(code string) ApiGetMarketStateRequest {
	r.code = &code
	return r
}

// The country where the exchange is located. Takes country name or alpha code.
func (r ApiGetMarketStateRequest) Country(country string) ApiGetMarketStateRequest {
	r.country = &country
	return r
}

func (r ApiGetMarketStateRequest) Execute() ([]MarketStateResponseItem, *http.Response, error) {
	return r.ApiService.GetMarketStateExecute(r)
}

/*
GetMarketState Market state

The market state endpoint provides real-time information on the operational status of all available stock exchanges. It returns data on whether each exchange is currently open or closed, along with the time remaining until the next opening or closing. This endpoint is useful for users who need to monitor exchange hours and plan their trading activities accordingly.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMarketStateRequest
*/
func (a *ReferenceDataAPIService) GetMarketState(ctx context.Context) ApiGetMarketStateRequest {
	return ApiGetMarketStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MarketStateResponseItem
func (a *ReferenceDataAPIService) GetMarketStateExecute(r ApiGetMarketStateRequest) ([]MarketStateResponseItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MarketStateResponseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetMarketState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market_state"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.exchange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exchange", r.exchange, "form", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
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

type ApiGetMutualFundsFamilyRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	fundFamily *string
	country *string
}

// Filter by investment company that manages the fund
func (r ApiGetMutualFundsFamilyRequest) FundFamily(fundFamily string) ApiGetMutualFundsFamilyRequest {
	r.fundFamily = &fundFamily
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsFamilyRequest) Country(country string) ApiGetMutualFundsFamilyRequest {
	r.country = &country
	return r
}

func (r ApiGetMutualFundsFamilyRequest) Execute() (*GetMutualFundsFamily200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsFamilyExecute(r)
}

/*
GetMutualFundsFamily MFs families

The mutual funds family endpoint provides a comprehensive list of MF families, which are groups of mutual funds managed by the same investment company. This data is useful for users looking to explore or compare different fund families, understand the range of investment options offered by each, and identify potential investment opportunities within specific fund families.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsFamilyRequest
*/
func (a *ReferenceDataAPIService) GetMutualFundsFamily(ctx context.Context) ApiGetMutualFundsFamilyRequest {
	return ApiGetMutualFundsFamilyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsFamily200Response
func (a *ReferenceDataAPIService) GetMutualFundsFamilyExecute(r ApiGetMutualFundsFamilyRequest) (*GetMutualFundsFamily200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsFamily200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetMutualFundsFamily")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/family"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fundFamily != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_family", r.fundFamily, "form", "")
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

type ApiGetMutualFundsListRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	cik *string
	country *string
	fundFamily *string
	fundType *string
	performanceRating *int64
	riskRating *int64
	page *int64
	outputsize *int64
}

// Filter by symbol
func (r ApiGetMutualFundsListRequest) Symbol(symbol string) ApiGetMutualFundsListRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetMutualFundsListRequest) Figi(figi string) ApiGetMutualFundsListRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetMutualFundsListRequest) Isin(isin string) ApiGetMutualFundsListRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetMutualFundsListRequest) Cusip(cusip string) ApiGetMutualFundsListRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ApiGetMutualFundsListRequest) Cik(cik string) ApiGetMutualFundsListRequest {
	r.cik = &cik
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsListRequest) Country(country string) ApiGetMutualFundsListRequest {
	r.country = &country
	return r
}

// Filter by investment company that manages the fund
func (r ApiGetMutualFundsListRequest) FundFamily(fundFamily string) ApiGetMutualFundsListRequest {
	r.fundFamily = &fundFamily
	return r
}

// Filter by the type of fund
func (r ApiGetMutualFundsListRequest) FundType(fundType string) ApiGetMutualFundsListRequest {
	r.fundType = &fundType
	return r
}

// Filter by performance rating from &#x60;0&#x60; to &#x60;5&#x60;
func (r ApiGetMutualFundsListRequest) PerformanceRating(performanceRating int64) ApiGetMutualFundsListRequest {
	r.performanceRating = &performanceRating
	return r
}

// Filter by risk rating from &#x60;0&#x60; to &#x60;5&#x60;
func (r ApiGetMutualFundsListRequest) RiskRating(riskRating int64) ApiGetMutualFundsListRequest {
	r.riskRating = &riskRating
	return r
}

// Page number
func (r ApiGetMutualFundsListRequest) Page(page int64) ApiGetMutualFundsListRequest {
	r.page = &page
	return r
}

// Number of records in response
func (r ApiGetMutualFundsListRequest) Outputsize(outputsize int64) ApiGetMutualFundsListRequest {
	r.outputsize = &outputsize
	return r
}

func (r ApiGetMutualFundsListRequest) Execute() (*GetMutualFundsList200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsListExecute(r)
}

/*
GetMutualFundsList MFs directory

The mutual funds directory endpoint provides a daily updated list of mutual funds, sorted in descending order by their total assets value. This endpoint is useful for retrieving an organized overview of available mutual funds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsListRequest
*/
func (a *ReferenceDataAPIService) GetMutualFundsList(ctx context.Context) ApiGetMutualFundsListRequest {
	return ApiGetMutualFundsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsList200Response
func (a *ReferenceDataAPIService) GetMutualFundsListExecute(r ApiGetMutualFundsListRequest) (*GetMutualFundsList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetMutualFundsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/list"

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
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
	}
	if r.country != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country", r.country, "form", "")
	}
	if r.fundFamily != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_family", r.fundFamily, "form", "")
	}
	if r.fundType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_type", r.fundType, "form", "")
	}
	if r.performanceRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "performance_rating", r.performanceRating, "form", "")
	}
	if r.riskRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "risk_rating", r.riskRating, "form", "")
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
        var defaultValue int64 = 100
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

type ApiGetMutualFundsTypeRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	fundType *string
	country *string
}

// Filter by the type of fund
func (r ApiGetMutualFundsTypeRequest) FundType(fundType string) ApiGetMutualFundsTypeRequest {
	r.fundType = &fundType
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetMutualFundsTypeRequest) Country(country string) ApiGetMutualFundsTypeRequest {
	r.country = &country
	return r
}

func (r ApiGetMutualFundsTypeRequest) Execute() (*GetMutualFundsType200Response, *http.Response, error) {
	return r.ApiService.GetMutualFundsTypeExecute(r)
}

/*
GetMutualFundsType MFs types

This endpoint provides detailed information on various types of mutual funds, such as equity, bond, and balanced funds, allowing users to understand the different investment options available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMutualFundsTypeRequest
*/
func (a *ReferenceDataAPIService) GetMutualFundsType(ctx context.Context) ApiGetMutualFundsTypeRequest {
	return ApiGetMutualFundsTypeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMutualFundsType200Response
func (a *ReferenceDataAPIService) GetMutualFundsTypeExecute(r ApiGetMutualFundsTypeRequest) (*GetMutualFundsType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMutualFundsType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetMutualFundsType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mutual_funds/type"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fundType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fund_type", r.fundType, "form", "")
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

type ApiGetStocksRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	figi *string
	isin *string
	cusip *string
	cik *string
	exchange *string
	micCode *string
	country *string
	type_ *string
	format *string
	delimiter *string
	showPlan *bool
	includeDelisted *bool
}

// The ticker symbol of an instrument for which data is requested
func (r ApiGetStocksRequest) Symbol(symbol string) ApiGetStocksRequest {
	r.symbol = &symbol
	return r
}

// Filter by financial instrument global identifier (FIGI)
func (r ApiGetStocksRequest) Figi(figi string) ApiGetStocksRequest {
	r.figi = &figi
	return r
}

// Filter by international securities identification number (ISIN)
func (r ApiGetStocksRequest) Isin(isin string) ApiGetStocksRequest {
	r.isin = &isin
	return r
}

// The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Add-ons&lt;/a&gt; section
func (r ApiGetStocksRequest) Cusip(cusip string) ApiGetStocksRequest {
	r.cusip = &cusip
	return r
}

// The CIK of an instrument for which data is requested
func (r ApiGetStocksRequest) Cik(cik string) ApiGetStocksRequest {
	r.cik = &cik
	return r
}

// Filter by exchange name
func (r ApiGetStocksRequest) Exchange(exchange string) ApiGetStocksRequest {
	r.exchange = &exchange
	return r
}

// Filter by market identifier code (MIC) under ISO 10383 standard
func (r ApiGetStocksRequest) MicCode(micCode string) ApiGetStocksRequest {
	r.micCode = &micCode
	return r
}

// Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60;
func (r ApiGetStocksRequest) Country(country string) ApiGetStocksRequest {
	r.country = &country
	return r
}

// The asset class to which the instrument belongs
func (r ApiGetStocksRequest) Type_(type_ string) ApiGetStocksRequest {
	r.type_ = &type_
	return r
}

// The format of the response data
func (r ApiGetStocksRequest) Format(format string) ApiGetStocksRequest {
	r.format = &format
	return r
}

// The separator used in the CSV response data
func (r ApiGetStocksRequest) Delimiter(delimiter string) ApiGetStocksRequest {
	r.delimiter = &delimiter
	return r
}

// Adds info on which plan symbol is available
func (r ApiGetStocksRequest) ShowPlan(showPlan bool) ApiGetStocksRequest {
	r.showPlan = &showPlan
	return r
}

// Include delisted identifiers
func (r ApiGetStocksRequest) IncludeDelisted(includeDelisted bool) ApiGetStocksRequest {
	r.includeDelisted = &includeDelisted
	return r
}

func (r ApiGetStocksRequest) Execute() (*GetStocks200Response, *http.Response, error) {
	return r.ApiService.GetStocksExecute(r)
}

/*
GetStocks Stocks

The stocks endpoint provides a daily updated list of all available stock symbols. It returns an array containing the symbols, which can be used to identify and access specific stock data across various services. This endpoint is essential for users needing to retrieve the latest stock symbol information for further data requests or integration into financial applications.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStocksRequest
*/
func (a *ReferenceDataAPIService) GetStocks(ctx context.Context) ApiGetStocksRequest {
	return ApiGetStocksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetStocks200Response
func (a *ReferenceDataAPIService) GetStocksExecute(r ApiGetStocksRequest) (*GetStocks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStocks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetStocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stocks"

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
	if r.cik != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "form", "")
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
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
	}
	if r.includeDelisted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", r.includeDelisted, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "include_delisted", defaultValue, "form", "")
        r.includeDelisted = &defaultValue
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

type ApiGetSymbolSearchRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
	symbol *string
	outputsize *int64
	showPlan *bool
}

// Symbol to search. Supports: &lt;ul&gt; &lt;li&gt;Ticker symbol of instrument.&lt;/li&gt; &lt;li&gt;International securities identification number (ISIN). &lt;li&gt;Financial instrument global identifier (FIGI). &lt;li&gt;Composite FIGI.&lt;/li&gt; &lt;li&gt;Share Class FIGI.&lt;/li&gt; &lt;/ul&gt;
func (r ApiGetSymbolSearchRequest) Symbol(symbol string) ApiGetSymbolSearchRequest {
	r.symbol = &symbol
	return r
}

// Number of matches in response. Max &lt;code&gt;120&lt;/code&gt;
func (r ApiGetSymbolSearchRequest) Outputsize(outputsize int64) ApiGetSymbolSearchRequest {
	r.outputsize = &outputsize
	return r
}

// Adds info on which plan symbol is available.
func (r ApiGetSymbolSearchRequest) ShowPlan(showPlan bool) ApiGetSymbolSearchRequest {
	r.showPlan = &showPlan
	return r
}

func (r ApiGetSymbolSearchRequest) Execute() (*GetSymbolSearch200Response, *http.Response, error) {
	return r.ApiService.GetSymbolSearchExecute(r)
}

/*
GetSymbolSearch Symbol search

The symbol search endpoint allows users to find financial instruments by name or symbol. It returns a list of matching symbols, ordered by relevance, with the most relevant instrument first. This is useful for quickly locating specific stocks, ETFs, or other financial instruments when only partial information is available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSymbolSearchRequest
*/
func (a *ReferenceDataAPIService) GetSymbolSearch(ctx context.Context) ApiGetSymbolSearchRequest {
	return ApiGetSymbolSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSymbolSearch200Response
func (a *ReferenceDataAPIService) GetSymbolSearchExecute(r ApiGetSymbolSearchRequest) (*GetSymbolSearch200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSymbolSearch200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetSymbolSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/symbol_search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.outputsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", r.outputsize, "form", "")
	} else {
        var defaultValue int64 = 30
        parameterAddToHeaderOrQuery(localVarQueryParams, "outputsize", defaultValue, "form", "")
        r.outputsize = &defaultValue
	}
	if r.showPlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", r.showPlan, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "show_plan", defaultValue, "form", "")
        r.showPlan = &defaultValue
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

type ApiGetTechnicalIndicatorsRequest struct {
	ctx context.Context
	ApiService *ReferenceDataAPIService
}

func (r ApiGetTechnicalIndicatorsRequest) Execute() (*GetTechnicalIndicators200Response, *http.Response, error) {
	return r.ApiService.GetTechnicalIndicatorsExecute(r)
}

/*
GetTechnicalIndicators Technical indicators

The technical indicators endpoint provides a comprehensive list of available technical indicators, each represented as an object. This endpoint is useful for developers looking to integrate a variety of technical analysis tools into their applications, allowing for streamlined access to indicator data without needing to manually configure each one.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTechnicalIndicatorsRequest
*/
func (a *ReferenceDataAPIService) GetTechnicalIndicators(ctx context.Context) ApiGetTechnicalIndicatorsRequest {
	return ApiGetTechnicalIndicatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTechnicalIndicators200Response
func (a *ReferenceDataAPIService) GetTechnicalIndicatorsExecute(r ApiGetTechnicalIndicatorsRequest) (*GetTechnicalIndicators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTechnicalIndicators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceDataAPIService.GetTechnicalIndicators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/technical_indicators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
