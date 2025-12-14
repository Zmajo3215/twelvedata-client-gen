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

// checks if the EquityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityInfo{}

// EquityInfo EquityInfo represents equity information
type EquityInfo struct {
	// Total equity gross minority interest
	TotalEquityGrossMinorityInterest *float64 `json:"total_equity_gross_minority_interest,omitempty"`
	// Stockholders equity
	StockholdersEquity *float64 `json:"stockholders_equity,omitempty"`
	// Common stock equity
	CommonStockEquity *float64 `json:"common_stock_equity,omitempty"`
	// Preferred stock equity
	PreferredStockEquity *float64 `json:"preferred_stock_equity,omitempty"`
	// Other equity interest
	OtherEquityInterest *float64 `json:"other_equity_interest,omitempty"`
	// Minority interest
	MinorityInterest *float64 `json:"minority_interest,omitempty"`
	// Total capitalization
	TotalCapitalization *float64 `json:"total_capitalization,omitempty"`
	// Net tangible assets
	NetTangibleAssets *float64 `json:"net_tangible_assets,omitempty"`
	// Tangible book value
	TangibleBookValue *float64 `json:"tangible_book_value,omitempty"`
	// Invested capital
	InvestedCapital *float64 `json:"invested_capital,omitempty"`
	// Working capital
	WorkingCapital *float64 `json:"working_capital,omitempty"`
	CapitalStock *EquityInfoCapitalStock `json:"capital_stock,omitempty"`
	EquityAdjustments *EquityInfoEquityAdjustments `json:"equity_adjustments,omitempty"`
	// Net debt
	NetDebt *float64 `json:"net_debt,omitempty"`
	// Total debt
	TotalDebt *float64 `json:"total_debt,omitempty"`
}

// NewEquityInfo instantiates a new EquityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInfo() *EquityInfo {
	this := EquityInfo{}
	return &this
}

// NewEquityInfoWithDefaults instantiates a new EquityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInfoWithDefaults() *EquityInfo {
	this := EquityInfo{}
	return &this
}

// GetTotalEquityGrossMinorityInterest returns the TotalEquityGrossMinorityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalEquityGrossMinorityInterest() float64 {
	if o == nil || IsNil(o.TotalEquityGrossMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.TotalEquityGrossMinorityInterest
}

// GetTotalEquityGrossMinorityInterestOk returns a tuple with the TotalEquityGrossMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalEquityGrossMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalEquityGrossMinorityInterest) {
		return nil, false
	}
	return o.TotalEquityGrossMinorityInterest, true
}

// HasTotalEquityGrossMinorityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalEquityGrossMinorityInterest() bool {
	if o != nil && !IsNil(o.TotalEquityGrossMinorityInterest) {
		return true
	}

	return false
}

// SetTotalEquityGrossMinorityInterest gets a reference to the given float64 and assigns it to the TotalEquityGrossMinorityInterest field.
func (o *EquityInfo) SetTotalEquityGrossMinorityInterest(v float64) {
	o.TotalEquityGrossMinorityInterest = &v
}

// GetStockholdersEquity returns the StockholdersEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetStockholdersEquity() float64 {
	if o == nil || IsNil(o.StockholdersEquity) {
		var ret float64
		return ret
	}
	return *o.StockholdersEquity
}

// GetStockholdersEquityOk returns a tuple with the StockholdersEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetStockholdersEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.StockholdersEquity) {
		return nil, false
	}
	return o.StockholdersEquity, true
}

// HasStockholdersEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasStockholdersEquity() bool {
	if o != nil && !IsNil(o.StockholdersEquity) {
		return true
	}

	return false
}

// SetStockholdersEquity gets a reference to the given float64 and assigns it to the StockholdersEquity field.
func (o *EquityInfo) SetStockholdersEquity(v float64) {
	o.StockholdersEquity = &v
}

// GetCommonStockEquity returns the CommonStockEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetCommonStockEquity() float64 {
	if o == nil || IsNil(o.CommonStockEquity) {
		var ret float64
		return ret
	}
	return *o.CommonStockEquity
}

// GetCommonStockEquityOk returns a tuple with the CommonStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetCommonStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockEquity) {
		return nil, false
	}
	return o.CommonStockEquity, true
}

// HasCommonStockEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasCommonStockEquity() bool {
	if o != nil && !IsNil(o.CommonStockEquity) {
		return true
	}

	return false
}

// SetCommonStockEquity gets a reference to the given float64 and assigns it to the CommonStockEquity field.
func (o *EquityInfo) SetCommonStockEquity(v float64) {
	o.CommonStockEquity = &v
}

// GetPreferredStockEquity returns the PreferredStockEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetPreferredStockEquity() float64 {
	if o == nil || IsNil(o.PreferredStockEquity) {
		var ret float64
		return ret
	}
	return *o.PreferredStockEquity
}

// GetPreferredStockEquityOk returns a tuple with the PreferredStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetPreferredStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockEquity) {
		return nil, false
	}
	return o.PreferredStockEquity, true
}

// HasPreferredStockEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasPreferredStockEquity() bool {
	if o != nil && !IsNil(o.PreferredStockEquity) {
		return true
	}

	return false
}

// SetPreferredStockEquity gets a reference to the given float64 and assigns it to the PreferredStockEquity field.
func (o *EquityInfo) SetPreferredStockEquity(v float64) {
	o.PreferredStockEquity = &v
}

// GetOtherEquityInterest returns the OtherEquityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetOtherEquityInterest() float64 {
	if o == nil || IsNil(o.OtherEquityInterest) {
		var ret float64
		return ret
	}
	return *o.OtherEquityInterest
}

// GetOtherEquityInterestOk returns a tuple with the OtherEquityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetOtherEquityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherEquityInterest) {
		return nil, false
	}
	return o.OtherEquityInterest, true
}

// HasOtherEquityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasOtherEquityInterest() bool {
	if o != nil && !IsNil(o.OtherEquityInterest) {
		return true
	}

	return false
}

// SetOtherEquityInterest gets a reference to the given float64 and assigns it to the OtherEquityInterest field.
func (o *EquityInfo) SetOtherEquityInterest(v float64) {
	o.OtherEquityInterest = &v
}

// GetMinorityInterest returns the MinorityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetMinorityInterest() float64 {
	if o == nil || IsNil(o.MinorityInterest) {
		var ret float64
		return ret
	}
	return *o.MinorityInterest
}

// GetMinorityInterestOk returns a tuple with the MinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.MinorityInterest) {
		return nil, false
	}
	return o.MinorityInterest, true
}

// HasMinorityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasMinorityInterest() bool {
	if o != nil && !IsNil(o.MinorityInterest) {
		return true
	}

	return false
}

// SetMinorityInterest gets a reference to the given float64 and assigns it to the MinorityInterest field.
func (o *EquityInfo) SetMinorityInterest(v float64) {
	o.MinorityInterest = &v
}

// GetTotalCapitalization returns the TotalCapitalization field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalCapitalization() float64 {
	if o == nil || IsNil(o.TotalCapitalization) {
		var ret float64
		return ret
	}
	return *o.TotalCapitalization
}

// GetTotalCapitalizationOk returns a tuple with the TotalCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalCapitalizationOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCapitalization) {
		return nil, false
	}
	return o.TotalCapitalization, true
}

// HasTotalCapitalization returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalCapitalization() bool {
	if o != nil && !IsNil(o.TotalCapitalization) {
		return true
	}

	return false
}

// SetTotalCapitalization gets a reference to the given float64 and assigns it to the TotalCapitalization field.
func (o *EquityInfo) SetTotalCapitalization(v float64) {
	o.TotalCapitalization = &v
}

// GetNetTangibleAssets returns the NetTangibleAssets field value if set, zero value otherwise.
func (o *EquityInfo) GetNetTangibleAssets() float64 {
	if o == nil || IsNil(o.NetTangibleAssets) {
		var ret float64
		return ret
	}
	return *o.NetTangibleAssets
}

// GetNetTangibleAssetsOk returns a tuple with the NetTangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetNetTangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetTangibleAssets) {
		return nil, false
	}
	return o.NetTangibleAssets, true
}

// HasNetTangibleAssets returns a boolean if a field has been set.
func (o *EquityInfo) HasNetTangibleAssets() bool {
	if o != nil && !IsNil(o.NetTangibleAssets) {
		return true
	}

	return false
}

// SetNetTangibleAssets gets a reference to the given float64 and assigns it to the NetTangibleAssets field.
func (o *EquityInfo) SetNetTangibleAssets(v float64) {
	o.NetTangibleAssets = &v
}

// GetTangibleBookValue returns the TangibleBookValue field value if set, zero value otherwise.
func (o *EquityInfo) GetTangibleBookValue() float64 {
	if o == nil || IsNil(o.TangibleBookValue) {
		var ret float64
		return ret
	}
	return *o.TangibleBookValue
}

// GetTangibleBookValueOk returns a tuple with the TangibleBookValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTangibleBookValueOk() (*float64, bool) {
	if o == nil || IsNil(o.TangibleBookValue) {
		return nil, false
	}
	return o.TangibleBookValue, true
}

// HasTangibleBookValue returns a boolean if a field has been set.
func (o *EquityInfo) HasTangibleBookValue() bool {
	if o != nil && !IsNil(o.TangibleBookValue) {
		return true
	}

	return false
}

// SetTangibleBookValue gets a reference to the given float64 and assigns it to the TangibleBookValue field.
func (o *EquityInfo) SetTangibleBookValue(v float64) {
	o.TangibleBookValue = &v
}

// GetInvestedCapital returns the InvestedCapital field value if set, zero value otherwise.
func (o *EquityInfo) GetInvestedCapital() float64 {
	if o == nil || IsNil(o.InvestedCapital) {
		var ret float64
		return ret
	}
	return *o.InvestedCapital
}

// GetInvestedCapitalOk returns a tuple with the InvestedCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetInvestedCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestedCapital) {
		return nil, false
	}
	return o.InvestedCapital, true
}

// HasInvestedCapital returns a boolean if a field has been set.
func (o *EquityInfo) HasInvestedCapital() bool {
	if o != nil && !IsNil(o.InvestedCapital) {
		return true
	}

	return false
}

// SetInvestedCapital gets a reference to the given float64 and assigns it to the InvestedCapital field.
func (o *EquityInfo) SetInvestedCapital(v float64) {
	o.InvestedCapital = &v
}

// GetWorkingCapital returns the WorkingCapital field value if set, zero value otherwise.
func (o *EquityInfo) GetWorkingCapital() float64 {
	if o == nil || IsNil(o.WorkingCapital) {
		var ret float64
		return ret
	}
	return *o.WorkingCapital
}

// GetWorkingCapitalOk returns a tuple with the WorkingCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetWorkingCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.WorkingCapital) {
		return nil, false
	}
	return o.WorkingCapital, true
}

// HasWorkingCapital returns a boolean if a field has been set.
func (o *EquityInfo) HasWorkingCapital() bool {
	if o != nil && !IsNil(o.WorkingCapital) {
		return true
	}

	return false
}

// SetWorkingCapital gets a reference to the given float64 and assigns it to the WorkingCapital field.
func (o *EquityInfo) SetWorkingCapital(v float64) {
	o.WorkingCapital = &v
}

// GetCapitalStock returns the CapitalStock field value if set, zero value otherwise.
func (o *EquityInfo) GetCapitalStock() EquityInfoCapitalStock {
	if o == nil || IsNil(o.CapitalStock) {
		var ret EquityInfoCapitalStock
		return ret
	}
	return *o.CapitalStock
}

// GetCapitalStockOk returns a tuple with the CapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetCapitalStockOk() (*EquityInfoCapitalStock, bool) {
	if o == nil || IsNil(o.CapitalStock) {
		return nil, false
	}
	return o.CapitalStock, true
}

// HasCapitalStock returns a boolean if a field has been set.
func (o *EquityInfo) HasCapitalStock() bool {
	if o != nil && !IsNil(o.CapitalStock) {
		return true
	}

	return false
}

// SetCapitalStock gets a reference to the given EquityInfoCapitalStock and assigns it to the CapitalStock field.
func (o *EquityInfo) SetCapitalStock(v EquityInfoCapitalStock) {
	o.CapitalStock = &v
}

// GetEquityAdjustments returns the EquityAdjustments field value if set, zero value otherwise.
func (o *EquityInfo) GetEquityAdjustments() EquityInfoEquityAdjustments {
	if o == nil || IsNil(o.EquityAdjustments) {
		var ret EquityInfoEquityAdjustments
		return ret
	}
	return *o.EquityAdjustments
}

// GetEquityAdjustmentsOk returns a tuple with the EquityAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetEquityAdjustmentsOk() (*EquityInfoEquityAdjustments, bool) {
	if o == nil || IsNil(o.EquityAdjustments) {
		return nil, false
	}
	return o.EquityAdjustments, true
}

// HasEquityAdjustments returns a boolean if a field has been set.
func (o *EquityInfo) HasEquityAdjustments() bool {
	if o != nil && !IsNil(o.EquityAdjustments) {
		return true
	}

	return false
}

// SetEquityAdjustments gets a reference to the given EquityInfoEquityAdjustments and assigns it to the EquityAdjustments field.
func (o *EquityInfo) SetEquityAdjustments(v EquityInfoEquityAdjustments) {
	o.EquityAdjustments = &v
}

// GetNetDebt returns the NetDebt field value if set, zero value otherwise.
func (o *EquityInfo) GetNetDebt() float64 {
	if o == nil || IsNil(o.NetDebt) {
		var ret float64
		return ret
	}
	return *o.NetDebt
}

// GetNetDebtOk returns a tuple with the NetDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetNetDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.NetDebt) {
		return nil, false
	}
	return o.NetDebt, true
}

// HasNetDebt returns a boolean if a field has been set.
func (o *EquityInfo) HasNetDebt() bool {
	if o != nil && !IsNil(o.NetDebt) {
		return true
	}

	return false
}

// SetNetDebt gets a reference to the given float64 and assigns it to the NetDebt field.
func (o *EquityInfo) SetNetDebt(v float64) {
	o.NetDebt = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalDebt() float64 {
	if o == nil || IsNil(o.TotalDebt) {
		var ret float64
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalDebt() bool {
	if o != nil && !IsNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given float64 and assigns it to the TotalDebt field.
func (o *EquityInfo) SetTotalDebt(v float64) {
	o.TotalDebt = &v
}

func (o EquityInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalEquityGrossMinorityInterest) {
		toSerialize["total_equity_gross_minority_interest"] = o.TotalEquityGrossMinorityInterest
	}
	if !IsNil(o.StockholdersEquity) {
		toSerialize["stockholders_equity"] = o.StockholdersEquity
	}
	if !IsNil(o.CommonStockEquity) {
		toSerialize["common_stock_equity"] = o.CommonStockEquity
	}
	if !IsNil(o.PreferredStockEquity) {
		toSerialize["preferred_stock_equity"] = o.PreferredStockEquity
	}
	if !IsNil(o.OtherEquityInterest) {
		toSerialize["other_equity_interest"] = o.OtherEquityInterest
	}
	if !IsNil(o.MinorityInterest) {
		toSerialize["minority_interest"] = o.MinorityInterest
	}
	if !IsNil(o.TotalCapitalization) {
		toSerialize["total_capitalization"] = o.TotalCapitalization
	}
	if !IsNil(o.NetTangibleAssets) {
		toSerialize["net_tangible_assets"] = o.NetTangibleAssets
	}
	if !IsNil(o.TangibleBookValue) {
		toSerialize["tangible_book_value"] = o.TangibleBookValue
	}
	if !IsNil(o.InvestedCapital) {
		toSerialize["invested_capital"] = o.InvestedCapital
	}
	if !IsNil(o.WorkingCapital) {
		toSerialize["working_capital"] = o.WorkingCapital
	}
	if !IsNil(o.CapitalStock) {
		toSerialize["capital_stock"] = o.CapitalStock
	}
	if !IsNil(o.EquityAdjustments) {
		toSerialize["equity_adjustments"] = o.EquityAdjustments
	}
	if !IsNil(o.NetDebt) {
		toSerialize["net_debt"] = o.NetDebt
	}
	if !IsNil(o.TotalDebt) {
		toSerialize["total_debt"] = o.TotalDebt
	}
	return toSerialize, nil
}

type NullableEquityInfo struct {
	value *EquityInfo
	isSet bool
}

func (v NullableEquityInfo) Get() *EquityInfo {
	return v.value
}

func (v *NullableEquityInfo) Set(val *EquityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInfo(val *EquityInfo) *NullableEquityInfo {
	return &NullableEquityInfo{value: val, isSet: true}
}

func (v NullableEquityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


