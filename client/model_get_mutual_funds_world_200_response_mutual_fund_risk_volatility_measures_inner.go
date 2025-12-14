/*
Twelve Data API

## Overview  Welcome to Twelve Data developer docs — your gateway to comprehensive financial market data through a powerful and easy-to-use API. Twelve Data provides access to financial markets across over 50 global countries, covering more than 1 million public instruments, including stocks, forex, ETFs, mutual funds, commodities, and cryptocurrencies.  ## Quickstart  To get started, you'll need to sign up for an API key. Once you have your API key, you can start making requests to the API.  ### Step 1: Create Twelve Data account  Sign up on the Twelve Data website to create your account [here](https://twelvedata.com/register). This gives you access to the API dashboard and your API key.  ### Step 2: Get your API key  After signing in, navigate to your [dashboard](https://twelvedata.com/account/api-keys) to find your unique API key. This key is required to authenticate all API and WebSocket requests.  ### Step 3: Make your first request  Try a simple API call with cURL to fetch the latest price for Apple (AAPL):  ``` curl \"https://api.twelvedata.com/price?symbol=AAPL&apikey=your_api_key\" ```  ### Step 4: Make a request from Python or Javascript  Use our client libraries or standard HTTP clients to make API calls programmatically. Here’s an example in [Python](https://github.com/twelvedata/twelvedata-python) and JavaScript:  #### Python (using official Twelve Data SDK):  ```python from twelvedata import TDClient  # Initialize client with your API key td = TDClient(apikey=\"your_api_key\")  # Get latest price for Apple price = td.price(symbol=\"AAPL\").as_json()  print(price) ```  #### JavaScript (Node.js):  ```javascript const fetch = require('node-fetch');  fetch('https://api.twelvedata.com/price?symbol=AAPL&apikey=your_api_key') &nbsp;&nbsp;.then(response => response.json()) &nbsp;&nbsp;.then(data => console.log(data)); ```  ### Step 5: Perform correlation analysis between Tesla and Microsoft prices  Fetch historical price data for Tesla (TSLA) and Microsoft (MSFT) and calculate the correlation of their closing prices:  ```python from twelvedata import TDClient import pandas as pd  # Initialize client with your API key td = TDClient(apikey=\"your_api_key\")  # Fetch historical price data for Tesla tsla_ts = td.time_series( &nbsp;&nbsp;&nbsp;&nbsp;symbol=\"TSLA\", &nbsp;&nbsp;&nbsp;&nbsp;interval=\"1day\", &nbsp;&nbsp;&nbsp;&nbsp;outputsize=100 ).as_pandas()  # Fetch historical price data for Microsoft msft_ts = td.time_series( &nbsp;&nbsp;&nbsp;&nbsp;symbol=\"MSFT\", &nbsp;&nbsp;&nbsp;&nbsp;interval=\"1day\", &nbsp;&nbsp;&nbsp;&nbsp;outputsize=100 ).as_pandas()  # Align data on datetime index combined = pd.concat( &nbsp;&nbsp;&nbsp;&nbsp;[tsla_ts['close'].astype(float), msft_ts['close'].astype(float)], &nbsp;&nbsp;&nbsp;&nbsp;axis=1, &nbsp;&nbsp;&nbsp;&nbsp;keys=[\"TSLA\", \"MSFT\"] ).dropna()  # Calculate correlation correlation = combined[\"TSLA\"].corr(combined[\"MSFT\"]) print(f\"Correlation of closing prices between TSLA and MSFT: {correlation:.2f}\") ```  ### Authentication  Authenticate your requests using one of these methods:  #### Query parameter method ``` GET https://api.twelvedata.com/endpoint?symbol=AAPL&apikey=your_api_key ```  #### HTTP header method (recommended) ``` Authorization: apikey your_api_key ```  ##### API key useful information <ul> <li> Demo API key (<code>apikey=demo</code>) available for demo requests</li> <li> Personal API key required for full access</li> <li> Premium endpoints and data require higher-tier plans (testable with <a href=\"https://twelvedata.com/exchanges\">trial symbols</a>)</li> </ul>  ### API endpoints   Service | Base URL | ---------|----------|  REST API | `https://api.twelvedata.com` |  WebSocket | `wss://ws.twelvedata.com` |  ### Parameter guidelines <ul> <li><b>Separator:</b> Use <code>&</code> to separate multiple parameters</li> <li><b>Case sensitivity:</b> Parameter names are case-insensitive</li>  <ul><li><code>symbol=AAPL</code> = <code>symbol=aapl</code></li></ul>  <li><b>Multiple values:</b> Separate with commas where supported</li> </ul>  ### Response handling  #### Default format All responses return JSON format by default unless otherwise specified.  #### Null values <b>Important:</b> Some response fields may contain `null` values when data is unavailable for specific metrics. This is expected behavior, not an error.  ##### Best Practices: <ul> <li>Always implement <code>null</code> value handling in your application</li> <li>Use defensive programming techniques for data processing</li> <li>Consider fallback values or error handling for critical metrics</li> </ul>  #### Error handling Structure your code to gracefully handle: <ul> <li>Network timeouts</li> <li>Rate limiting responses</li> <li>Invalid parameter errors</li> <li>Data unavailability periods</li> </ul>  ##### Best practices <ul> <li><b>Rate limits:</b> Adhere to your plan’s rate limits to avoid throttling. Check your dashboard for details.</li> <li><b>Error handling:</b> Implement retry logic for transient errors (e.g., <code>429 Too Many Requests</code>).</li> <li><b>Caching:</b> Cache responses for frequently accessed data to reduce API calls and improve performance.</li> <li><b>Secure storage:</b> Store your API key securely and never expose it in client-side code or public repositories.</li> </ul>  ## Errors  Twelve Data API employs a standardized error response format, delivering a JSON object with `code`, `message`, and `status` keys for clear and consistent error communication.  ### Codes  Below is a table of possible error codes, their HTTP status, meanings, and resolution steps:   Code | status | Meaning | Resolution |  --- | --- | --- | --- |  **400** | Bad Request | Invalid or incorrect parameter(s) provided. | Check the `message` in the response for details. Refer to the API Documenta­tion to correct the input. |  **401** | Unauthor­ized | Invalid or incorrect API key. | Verify your API key is correct. Sign up for a key <a href=\"https://twelvedata.com/account/api-keys\">here</a>. |  **403** | Forbidden | API key lacks permissions for the requested resource (upgrade required). | Upgrade your plan <a href=\"https://twelvedata.com/pricing\">here</a>. |  **404** | Not Found | Requested data could not be found. | Adjust parameters to be less strict as they may be too restrictive. |  **414** | Parameter Too Long | Input parameter array exceeds the allowed length. | Follow the `message` guidance to adjust the parameter length. |  **429** | Too Many Requests | API request limit reached for your key. | Wait briefly or upgrade your plan <a href=\"https://twelvedata.com/pricing\">here</a>. |  **500** | Internal Server Error | Server-side issue occurred; retry later. | Contact support <a href=\"https://twelvedata.com/contact\">here</a> for assistance. |  ### Example error response  Consider the following invalid request:  ``` https://api.twelvedata.com/time_series?symbol=AAPL&interval=0.99min&apikey=your_api_key ```  Due to the incorrect `interval` value, the API returns:  ```json { &nbsp;&nbsp;\"code\": 400, &nbsp;&nbsp;\"message\": \"Invalid **interval** provided: 0.99min. Supported intervals: 1min, 5min, 15min, 30min, 45min, 1h, 2h, 4h, 8h, 1day, 1week, 1month\", &nbsp;&nbsp;\"status\": \"error\" } ```  Refer to the API Documentation for valid parameter values to resolve such errors.  ## Libraries  Twelve Data provides a growing ecosystem of libraries and integrations to help you build faster and smarter in your preferred environment. Official libraries are actively maintained by the Twelve Data team, while selected community-built libraries offer additional flexibility.  A full list is available on our [GitHub profile](https://github.com/search?q=twelvedata).  ### Official SDKs <ul> <li><b>Python:</b> <a href=\"https://github.com/twelvedata/twelvedata-python\">twelvedata-python</a></li> <li><b>R:</b> <a href=\"https://github.com/twelvedata/twelvedata-r-sdk\">twelvedata-r-sdk</a></li> </ul>  ### AI integrations <ul> <li><b>Twelve Data MCP Server:</b> <a href=\"https://github.com/twelvedata/mcp\">Repository</a> — Model Context Protocol (MCP) server that provides seamless integration with AI assistants and language models, enabling direct access to Twelve Data's financial market data within conversational interfaces and AI workflows.</li> </ul>  ### Spreadsheet add-ons <ul> <li><b>Excel:</b> <a href=\"https://twelvedata.com/excel\">Excel Add-in</a></li> <li><b>Google Sheets:</b> <a href=\"https://twelvedata.com/google-sheets\">Google Sheets Add-on</a></li> </ul>  ### Community libraries  The community has developed libraries in several popular languages. You can explore more community libraries on [GitHub](https://github.com/search?q=twelvedata). <ul> <li><b>C#:</b> <a href=\"https://github.com/pseudomarkets/TwelveDataSharp\">TwelveDataSharp</a></li> <li><b>JavaScript:</b> <a href=\"https://github.com/evzaboun/twelvedata\">twelvedata</a></li> <li><b>PHP:</b> <a href=\"https://github.com/ingelby/twelvedata\">twelvedata</a></li> <li><b>Go:</b> <a href=\"https://github.com/soulgarden/twelvedata\">twelvedata</a></li> <li><b>TypeScript:</b> <a href=\"https://github.com/Clyde-Goodall/twelve-data-wrapper\">twelve-data-wrapper</a></li> </ul>  ### Other Twelve Data repositories <ul> <li><b>searchindex</b> <i>(Go)</i>: <a href=\"https://github.com/twelvedata/searchindex\">Repository</a> — In-memory search index by strings</li> <li><b>ws-tools</b> <i>(Python)</i>: <a href=\"https://github.com/twelvedata/ws-tools\">Repository</a> — Utility tools for WebSocket stream handling</li> </ul>  ### API specification <ul> <li><b>OpenAPI / Swagger:</b> Access the <a href=\"https://api.twelvedata.com/doc/swagger/openapi.json\">complete API specification</a> in OpenAPI format. You can use this file to automatically generate client libraries in your preferred programming language, explore the API interactively via Swagger tools, or integrate Twelve Data seamlessly into your AI and LLM workflows.</li> </ul>

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}

// GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct for GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
type GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct {
	// Period of a measure
	Period *string `json:"period,omitempty"`
	// Alpha score of a fund
	Alpha *float64 `json:"alpha,omitempty"`
	// Average alpha score of a fund's category
	AlphaCategory *float64 `json:"alpha_category,omitempty"`
	// Beta score of a fund
	Beta *float64 `json:"beta,omitempty"`
	// Average beta score of a fund's category
	BetaCategory *float64 `json:"beta_category,omitempty"`
	// Mean annual return of a fund
	MeanAnnualReturn *float64 `json:"mean_annual_return,omitempty"`
	// Average mean annual return of a fund's category
	MeanAnnualReturnCategory *float64 `json:"mean_annual_return_category,omitempty"`
	// R-squared metric of a fund
	RSquared *float64 `json:"r_squared,omitempty"`
	// Average r-squared metric of a fund's category
	RSquaredCategory *float64 `json:"r_squared_category,omitempty"`
	// Standard deviation of a fund
	Std *float64 `json:"std,omitempty"`
	// Average standard deviation of a fund's category
	StdCategory *float64 `json:"std_category,omitempty"`
	// Sharpe ratio of a fund
	SharpeRatio *float64 `json:"sharpe_ratio,omitempty"`
	// Average sharpe ratio of a fund's category
	SharpeRatioCategory *float64 `json:"sharpe_ratio_category,omitempty"`
	// Treynor ratio of a fund
	TreynorRatio *float64 `json:"treynor_ratio,omitempty"`
	// Average treynor ratio of a fund's category
	TreynorRatioCategory *float64 `json:"treynor_ratio_category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	this := GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	this := GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetPeriod(v string) {
	o.Period = &v
}

// GetAlpha returns the Alpha field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlpha() float64 {
	if o == nil || IsNil(o.Alpha) {
		var ret float64
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaOk() (*float64, bool) {
	if o == nil || IsNil(o.Alpha) {
		return nil, false
	}
	return o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasAlpha() bool {
	if o != nil && !IsNil(o.Alpha) {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given float64 and assigns it to the Alpha field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetAlpha(v float64) {
	o.Alpha = &v
}

// GetAlphaCategory returns the AlphaCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaCategory() float64 {
	if o == nil || IsNil(o.AlphaCategory) {
		var ret float64
		return ret
	}
	return *o.AlphaCategory
}

// GetAlphaCategoryOk returns a tuple with the AlphaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetAlphaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.AlphaCategory) {
		return nil, false
	}
	return o.AlphaCategory, true
}

// HasAlphaCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasAlphaCategory() bool {
	if o != nil && !IsNil(o.AlphaCategory) {
		return true
	}

	return false
}

// SetAlphaCategory gets a reference to the given float64 and assigns it to the AlphaCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetAlphaCategory(v float64) {
	o.AlphaCategory = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBeta() float64 {
	if o == nil || IsNil(o.Beta) {
		var ret float64
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaOk() (*float64, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given float64 and assigns it to the Beta field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetBeta(v float64) {
	o.Beta = &v
}

// GetBetaCategory returns the BetaCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaCategory() float64 {
	if o == nil || IsNil(o.BetaCategory) {
		var ret float64
		return ret
	}
	return *o.BetaCategory
}

// GetBetaCategoryOk returns a tuple with the BetaCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetBetaCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.BetaCategory) {
		return nil, false
	}
	return o.BetaCategory, true
}

// HasBetaCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasBetaCategory() bool {
	if o != nil && !IsNil(o.BetaCategory) {
		return true
	}

	return false
}

// SetBetaCategory gets a reference to the given float64 and assigns it to the BetaCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetBetaCategory(v float64) {
	o.BetaCategory = &v
}

// GetMeanAnnualReturn returns the MeanAnnualReturn field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturn() float64 {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturn
}

// GetMeanAnnualReturnOk returns a tuple with the MeanAnnualReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturn) {
		return nil, false
	}
	return o.MeanAnnualReturn, true
}

// HasMeanAnnualReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasMeanAnnualReturn() bool {
	if o != nil && !IsNil(o.MeanAnnualReturn) {
		return true
	}

	return false
}

// SetMeanAnnualReturn gets a reference to the given float64 and assigns it to the MeanAnnualReturn field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetMeanAnnualReturn(v float64) {
	o.MeanAnnualReturn = &v
}

// GetMeanAnnualReturnCategory returns the MeanAnnualReturnCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategory() float64 {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		var ret float64
		return ret
	}
	return *o.MeanAnnualReturnCategory
}

// GetMeanAnnualReturnCategoryOk returns a tuple with the MeanAnnualReturnCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetMeanAnnualReturnCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.MeanAnnualReturnCategory) {
		return nil, false
	}
	return o.MeanAnnualReturnCategory, true
}

// HasMeanAnnualReturnCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasMeanAnnualReturnCategory() bool {
	if o != nil && !IsNil(o.MeanAnnualReturnCategory) {
		return true
	}

	return false
}

// SetMeanAnnualReturnCategory gets a reference to the given float64 and assigns it to the MeanAnnualReturnCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetMeanAnnualReturnCategory(v float64) {
	o.MeanAnnualReturnCategory = &v
}

// GetRSquared returns the RSquared field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquared() float64 {
	if o == nil || IsNil(o.RSquared) {
		var ret float64
		return ret
	}
	return *o.RSquared
}

// GetRSquaredOk returns a tuple with the RSquared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquared) {
		return nil, false
	}
	return o.RSquared, true
}

// HasRSquared returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasRSquared() bool {
	if o != nil && !IsNil(o.RSquared) {
		return true
	}

	return false
}

// SetRSquared gets a reference to the given float64 and assigns it to the RSquared field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetRSquared(v float64) {
	o.RSquared = &v
}

// GetRSquaredCategory returns the RSquaredCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredCategory() float64 {
	if o == nil || IsNil(o.RSquaredCategory) {
		var ret float64
		return ret
	}
	return *o.RSquaredCategory
}

// GetRSquaredCategoryOk returns a tuple with the RSquaredCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetRSquaredCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.RSquaredCategory) {
		return nil, false
	}
	return o.RSquaredCategory, true
}

// HasRSquaredCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasRSquaredCategory() bool {
	if o != nil && !IsNil(o.RSquaredCategory) {
		return true
	}

	return false
}

// SetRSquaredCategory gets a reference to the given float64 and assigns it to the RSquaredCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetRSquaredCategory(v float64) {
	o.RSquaredCategory = &v
}

// GetStd returns the Std field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStd() float64 {
	if o == nil || IsNil(o.Std) {
		var ret float64
		return ret
	}
	return *o.Std
}

// GetStdOk returns a tuple with the Std field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdOk() (*float64, bool) {
	if o == nil || IsNil(o.Std) {
		return nil, false
	}
	return o.Std, true
}

// HasStd returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasStd() bool {
	if o != nil && !IsNil(o.Std) {
		return true
	}

	return false
}

// SetStd gets a reference to the given float64 and assigns it to the Std field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetStd(v float64) {
	o.Std = &v
}

// GetStdCategory returns the StdCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdCategory() float64 {
	if o == nil || IsNil(o.StdCategory) {
		var ret float64
		return ret
	}
	return *o.StdCategory
}

// GetStdCategoryOk returns a tuple with the StdCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetStdCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.StdCategory) {
		return nil, false
	}
	return o.StdCategory, true
}

// HasStdCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasStdCategory() bool {
	if o != nil && !IsNil(o.StdCategory) {
		return true
	}

	return false
}

// SetStdCategory gets a reference to the given float64 and assigns it to the StdCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetStdCategory(v float64) {
	o.StdCategory = &v
}

// GetSharpeRatio returns the SharpeRatio field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatio() float64 {
	if o == nil || IsNil(o.SharpeRatio) {
		var ret float64
		return ret
	}
	return *o.SharpeRatio
}

// GetSharpeRatioOk returns a tuple with the SharpeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatio) {
		return nil, false
	}
	return o.SharpeRatio, true
}

// HasSharpeRatio returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasSharpeRatio() bool {
	if o != nil && !IsNil(o.SharpeRatio) {
		return true
	}

	return false
}

// SetSharpeRatio gets a reference to the given float64 and assigns it to the SharpeRatio field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetSharpeRatio(v float64) {
	o.SharpeRatio = &v
}

// GetSharpeRatioCategory returns the SharpeRatioCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioCategory() float64 {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		var ret float64
		return ret
	}
	return *o.SharpeRatioCategory
}

// GetSharpeRatioCategoryOk returns a tuple with the SharpeRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetSharpeRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.SharpeRatioCategory) {
		return nil, false
	}
	return o.SharpeRatioCategory, true
}

// HasSharpeRatioCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasSharpeRatioCategory() bool {
	if o != nil && !IsNil(o.SharpeRatioCategory) {
		return true
	}

	return false
}

// SetSharpeRatioCategory gets a reference to the given float64 and assigns it to the SharpeRatioCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetSharpeRatioCategory(v float64) {
	o.SharpeRatioCategory = &v
}

// GetTreynorRatio returns the TreynorRatio field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatio() float64 {
	if o == nil || IsNil(o.TreynorRatio) {
		var ret float64
		return ret
	}
	return *o.TreynorRatio
}

// GetTreynorRatioOk returns a tuple with the TreynorRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatio) {
		return nil, false
	}
	return o.TreynorRatio, true
}

// HasTreynorRatio returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasTreynorRatio() bool {
	if o != nil && !IsNil(o.TreynorRatio) {
		return true
	}

	return false
}

// SetTreynorRatio gets a reference to the given float64 and assigns it to the TreynorRatio field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetTreynorRatio(v float64) {
	o.TreynorRatio = &v
}

// GetTreynorRatioCategory returns the TreynorRatioCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioCategory() float64 {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		var ret float64
		return ret
	}
	return *o.TreynorRatioCategory
}

// GetTreynorRatioCategoryOk returns a tuple with the TreynorRatioCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) GetTreynorRatioCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.TreynorRatioCategory) {
		return nil, false
	}
	return o.TreynorRatioCategory, true
}

// HasTreynorRatioCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) HasTreynorRatioCategory() bool {
	if o != nil && !IsNil(o.TreynorRatioCategory) {
		return true
	}

	return false
}

// SetTreynorRatioCategory gets a reference to the given float64 and assigns it to the TreynorRatioCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) SetTreynorRatioCategory(v float64) {
	o.TreynorRatioCategory = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Alpha) {
		toSerialize["alpha"] = o.Alpha
	}
	if !IsNil(o.AlphaCategory) {
		toSerialize["alpha_category"] = o.AlphaCategory
	}
	if !IsNil(o.Beta) {
		toSerialize["beta"] = o.Beta
	}
	if !IsNil(o.BetaCategory) {
		toSerialize["beta_category"] = o.BetaCategory
	}
	if !IsNil(o.MeanAnnualReturn) {
		toSerialize["mean_annual_return"] = o.MeanAnnualReturn
	}
	if !IsNil(o.MeanAnnualReturnCategory) {
		toSerialize["mean_annual_return_category"] = o.MeanAnnualReturnCategory
	}
	if !IsNil(o.RSquared) {
		toSerialize["r_squared"] = o.RSquared
	}
	if !IsNil(o.RSquaredCategory) {
		toSerialize["r_squared_category"] = o.RSquaredCategory
	}
	if !IsNil(o.Std) {
		toSerialize["std"] = o.Std
	}
	if !IsNil(o.StdCategory) {
		toSerialize["std_category"] = o.StdCategory
	}
	if !IsNil(o.SharpeRatio) {
		toSerialize["sharpe_ratio"] = o.SharpeRatio
	}
	if !IsNil(o.SharpeRatioCategory) {
		toSerialize["sharpe_ratio_category"] = o.SharpeRatioCategory
	}
	if !IsNil(o.TreynorRatio) {
		toSerialize["treynor_ratio"] = o.TreynorRatio
	}
	if !IsNil(o.TreynorRatioCategory) {
		toSerialize["treynor_ratio_category"] = o.TreynorRatioCategory
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Get() *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Set(val *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner(val *GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


