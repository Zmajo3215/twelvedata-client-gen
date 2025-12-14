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

// checks if the EquityInfoCapitalStock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityInfoCapitalStock{}

// EquityInfoCapitalStock Capital stock information
type EquityInfoCapitalStock struct {
	// Common stock
	CommonStock *float64 `json:"common_stock,omitempty"`
	// Preferred stock
	PreferredStock *float64 `json:"preferred_stock,omitempty"`
	// Total partnership capital
	TotalPartnershipCapital *float64 `json:"total_partnership_capital,omitempty"`
	// General partnership capital
	GeneralPartnershipCapital *float64 `json:"general_partnership_capital,omitempty"`
	// Limited partnership capital
	LimitedPartnershipCapital *float64 `json:"limited_partnership_capital,omitempty"`
	// Capital stock
	CapitalStock *float64 `json:"capital_stock,omitempty"`
	// Other capital stock
	OtherCapitalStock *float64 `json:"other_capital_stock,omitempty"`
	// Additional paid in capital
	AdditionalPaidInCapital *float64 `json:"additional_paid_in_capital,omitempty"`
	// Retained earnings
	RetainedEarnings *float64 `json:"retained_earnings,omitempty"`
	// Treasury stock
	TreasuryStock *float64 `json:"treasury_stock,omitempty"`
	// Treasury shares number
	TreasurySharesNumber *float64 `json:"treasury_shares_number,omitempty"`
	// Ordinary shares number
	OrdinarySharesNumber *float64 `json:"ordinary_shares_number,omitempty"`
	// Preferred shares number
	PreferredSharesNumber *float64 `json:"preferred_shares_number,omitempty"`
	// Share issued
	ShareIssued *float64 `json:"share_issued,omitempty"`
}

// NewEquityInfoCapitalStock instantiates a new EquityInfoCapitalStock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInfoCapitalStock() *EquityInfoCapitalStock {
	this := EquityInfoCapitalStock{}
	return &this
}

// NewEquityInfoCapitalStockWithDefaults instantiates a new EquityInfoCapitalStock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInfoCapitalStockWithDefaults() *EquityInfoCapitalStock {
	this := EquityInfoCapitalStock{}
	return &this
}

// GetCommonStock returns the CommonStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetCommonStock() float64 {
	if o == nil || IsNil(o.CommonStock) {
		var ret float64
		return ret
	}
	return *o.CommonStock
}

// GetCommonStockOk returns a tuple with the CommonStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetCommonStockOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStock) {
		return nil, false
	}
	return o.CommonStock, true
}

// HasCommonStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasCommonStock() bool {
	if o != nil && !IsNil(o.CommonStock) {
		return true
	}

	return false
}

// SetCommonStock gets a reference to the given float64 and assigns it to the CommonStock field.
func (o *EquityInfoCapitalStock) SetCommonStock(v float64) {
	o.CommonStock = &v
}

// GetPreferredStock returns the PreferredStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetPreferredStock() float64 {
	if o == nil || IsNil(o.PreferredStock) {
		var ret float64
		return ret
	}
	return *o.PreferredStock
}

// GetPreferredStockOk returns a tuple with the PreferredStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetPreferredStockOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStock) {
		return nil, false
	}
	return o.PreferredStock, true
}

// HasPreferredStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasPreferredStock() bool {
	if o != nil && !IsNil(o.PreferredStock) {
		return true
	}

	return false
}

// SetPreferredStock gets a reference to the given float64 and assigns it to the PreferredStock field.
func (o *EquityInfoCapitalStock) SetPreferredStock(v float64) {
	o.PreferredStock = &v
}

// GetTotalPartnershipCapital returns the TotalPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTotalPartnershipCapital() float64 {
	if o == nil || IsNil(o.TotalPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.TotalPartnershipCapital
}

// GetTotalPartnershipCapitalOk returns a tuple with the TotalPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTotalPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalPartnershipCapital) {
		return nil, false
	}
	return o.TotalPartnershipCapital, true
}

// HasTotalPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTotalPartnershipCapital() bool {
	if o != nil && !IsNil(o.TotalPartnershipCapital) {
		return true
	}

	return false
}

// SetTotalPartnershipCapital gets a reference to the given float64 and assigns it to the TotalPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetTotalPartnershipCapital(v float64) {
	o.TotalPartnershipCapital = &v
}

// GetGeneralPartnershipCapital returns the GeneralPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetGeneralPartnershipCapital() float64 {
	if o == nil || IsNil(o.GeneralPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.GeneralPartnershipCapital
}

// GetGeneralPartnershipCapitalOk returns a tuple with the GeneralPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetGeneralPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.GeneralPartnershipCapital) {
		return nil, false
	}
	return o.GeneralPartnershipCapital, true
}

// HasGeneralPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasGeneralPartnershipCapital() bool {
	if o != nil && !IsNil(o.GeneralPartnershipCapital) {
		return true
	}

	return false
}

// SetGeneralPartnershipCapital gets a reference to the given float64 and assigns it to the GeneralPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetGeneralPartnershipCapital(v float64) {
	o.GeneralPartnershipCapital = &v
}

// GetLimitedPartnershipCapital returns the LimitedPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetLimitedPartnershipCapital() float64 {
	if o == nil || IsNil(o.LimitedPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.LimitedPartnershipCapital
}

// GetLimitedPartnershipCapitalOk returns a tuple with the LimitedPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetLimitedPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.LimitedPartnershipCapital) {
		return nil, false
	}
	return o.LimitedPartnershipCapital, true
}

// HasLimitedPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasLimitedPartnershipCapital() bool {
	if o != nil && !IsNil(o.LimitedPartnershipCapital) {
		return true
	}

	return false
}

// SetLimitedPartnershipCapital gets a reference to the given float64 and assigns it to the LimitedPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetLimitedPartnershipCapital(v float64) {
	o.LimitedPartnershipCapital = &v
}

// GetCapitalStock returns the CapitalStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetCapitalStock() float64 {
	if o == nil || IsNil(o.CapitalStock) {
		var ret float64
		return ret
	}
	return *o.CapitalStock
}

// GetCapitalStockOk returns a tuple with the CapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalStock) {
		return nil, false
	}
	return o.CapitalStock, true
}

// HasCapitalStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasCapitalStock() bool {
	if o != nil && !IsNil(o.CapitalStock) {
		return true
	}

	return false
}

// SetCapitalStock gets a reference to the given float64 and assigns it to the CapitalStock field.
func (o *EquityInfoCapitalStock) SetCapitalStock(v float64) {
	o.CapitalStock = &v
}

// GetOtherCapitalStock returns the OtherCapitalStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetOtherCapitalStock() float64 {
	if o == nil || IsNil(o.OtherCapitalStock) {
		var ret float64
		return ret
	}
	return *o.OtherCapitalStock
}

// GetOtherCapitalStockOk returns a tuple with the OtherCapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetOtherCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCapitalStock) {
		return nil, false
	}
	return o.OtherCapitalStock, true
}

// HasOtherCapitalStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasOtherCapitalStock() bool {
	if o != nil && !IsNil(o.OtherCapitalStock) {
		return true
	}

	return false
}

// SetOtherCapitalStock gets a reference to the given float64 and assigns it to the OtherCapitalStock field.
func (o *EquityInfoCapitalStock) SetOtherCapitalStock(v float64) {
	o.OtherCapitalStock = &v
}

// GetAdditionalPaidInCapital returns the AdditionalPaidInCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetAdditionalPaidInCapital() float64 {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		var ret float64
		return ret
	}
	return *o.AdditionalPaidInCapital
}

// GetAdditionalPaidInCapitalOk returns a tuple with the AdditionalPaidInCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetAdditionalPaidInCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		return nil, false
	}
	return o.AdditionalPaidInCapital, true
}

// HasAdditionalPaidInCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasAdditionalPaidInCapital() bool {
	if o != nil && !IsNil(o.AdditionalPaidInCapital) {
		return true
	}

	return false
}

// SetAdditionalPaidInCapital gets a reference to the given float64 and assigns it to the AdditionalPaidInCapital field.
func (o *EquityInfoCapitalStock) SetAdditionalPaidInCapital(v float64) {
	o.AdditionalPaidInCapital = &v
}

// GetRetainedEarnings returns the RetainedEarnings field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetRetainedEarnings() float64 {
	if o == nil || IsNil(o.RetainedEarnings) {
		var ret float64
		return ret
	}
	return *o.RetainedEarnings
}

// GetRetainedEarningsOk returns a tuple with the RetainedEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetRetainedEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.RetainedEarnings) {
		return nil, false
	}
	return o.RetainedEarnings, true
}

// HasRetainedEarnings returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasRetainedEarnings() bool {
	if o != nil && !IsNil(o.RetainedEarnings) {
		return true
	}

	return false
}

// SetRetainedEarnings gets a reference to the given float64 and assigns it to the RetainedEarnings field.
func (o *EquityInfoCapitalStock) SetRetainedEarnings(v float64) {
	o.RetainedEarnings = &v
}

// GetTreasuryStock returns the TreasuryStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTreasuryStock() float64 {
	if o == nil || IsNil(o.TreasuryStock) {
		var ret float64
		return ret
	}
	return *o.TreasuryStock
}

// GetTreasuryStockOk returns a tuple with the TreasuryStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTreasuryStockOk() (*float64, bool) {
	if o == nil || IsNil(o.TreasuryStock) {
		return nil, false
	}
	return o.TreasuryStock, true
}

// HasTreasuryStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTreasuryStock() bool {
	if o != nil && !IsNil(o.TreasuryStock) {
		return true
	}

	return false
}

// SetTreasuryStock gets a reference to the given float64 and assigns it to the TreasuryStock field.
func (o *EquityInfoCapitalStock) SetTreasuryStock(v float64) {
	o.TreasuryStock = &v
}

// GetTreasurySharesNumber returns the TreasurySharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTreasurySharesNumber() float64 {
	if o == nil || IsNil(o.TreasurySharesNumber) {
		var ret float64
		return ret
	}
	return *o.TreasurySharesNumber
}

// GetTreasurySharesNumberOk returns a tuple with the TreasurySharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTreasurySharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.TreasurySharesNumber) {
		return nil, false
	}
	return o.TreasurySharesNumber, true
}

// HasTreasurySharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTreasurySharesNumber() bool {
	if o != nil && !IsNil(o.TreasurySharesNumber) {
		return true
	}

	return false
}

// SetTreasurySharesNumber gets a reference to the given float64 and assigns it to the TreasurySharesNumber field.
func (o *EquityInfoCapitalStock) SetTreasurySharesNumber(v float64) {
	o.TreasurySharesNumber = &v
}

// GetOrdinarySharesNumber returns the OrdinarySharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetOrdinarySharesNumber() float64 {
	if o == nil || IsNil(o.OrdinarySharesNumber) {
		var ret float64
		return ret
	}
	return *o.OrdinarySharesNumber
}

// GetOrdinarySharesNumberOk returns a tuple with the OrdinarySharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetOrdinarySharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.OrdinarySharesNumber) {
		return nil, false
	}
	return o.OrdinarySharesNumber, true
}

// HasOrdinarySharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasOrdinarySharesNumber() bool {
	if o != nil && !IsNil(o.OrdinarySharesNumber) {
		return true
	}

	return false
}

// SetOrdinarySharesNumber gets a reference to the given float64 and assigns it to the OrdinarySharesNumber field.
func (o *EquityInfoCapitalStock) SetOrdinarySharesNumber(v float64) {
	o.OrdinarySharesNumber = &v
}

// GetPreferredSharesNumber returns the PreferredSharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetPreferredSharesNumber() float64 {
	if o == nil || IsNil(o.PreferredSharesNumber) {
		var ret float64
		return ret
	}
	return *o.PreferredSharesNumber
}

// GetPreferredSharesNumberOk returns a tuple with the PreferredSharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetPreferredSharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredSharesNumber) {
		return nil, false
	}
	return o.PreferredSharesNumber, true
}

// HasPreferredSharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasPreferredSharesNumber() bool {
	if o != nil && !IsNil(o.PreferredSharesNumber) {
		return true
	}

	return false
}

// SetPreferredSharesNumber gets a reference to the given float64 and assigns it to the PreferredSharesNumber field.
func (o *EquityInfoCapitalStock) SetPreferredSharesNumber(v float64) {
	o.PreferredSharesNumber = &v
}

// GetShareIssued returns the ShareIssued field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetShareIssued() float64 {
	if o == nil || IsNil(o.ShareIssued) {
		var ret float64
		return ret
	}
	return *o.ShareIssued
}

// GetShareIssuedOk returns a tuple with the ShareIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetShareIssuedOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareIssued) {
		return nil, false
	}
	return o.ShareIssued, true
}

// HasShareIssued returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasShareIssued() bool {
	if o != nil && !IsNil(o.ShareIssued) {
		return true
	}

	return false
}

// SetShareIssued gets a reference to the given float64 and assigns it to the ShareIssued field.
func (o *EquityInfoCapitalStock) SetShareIssued(v float64) {
	o.ShareIssued = &v
}

func (o EquityInfoCapitalStock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityInfoCapitalStock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonStock) {
		toSerialize["common_stock"] = o.CommonStock
	}
	if !IsNil(o.PreferredStock) {
		toSerialize["preferred_stock"] = o.PreferredStock
	}
	if !IsNil(o.TotalPartnershipCapital) {
		toSerialize["total_partnership_capital"] = o.TotalPartnershipCapital
	}
	if !IsNil(o.GeneralPartnershipCapital) {
		toSerialize["general_partnership_capital"] = o.GeneralPartnershipCapital
	}
	if !IsNil(o.LimitedPartnershipCapital) {
		toSerialize["limited_partnership_capital"] = o.LimitedPartnershipCapital
	}
	if !IsNil(o.CapitalStock) {
		toSerialize["capital_stock"] = o.CapitalStock
	}
	if !IsNil(o.OtherCapitalStock) {
		toSerialize["other_capital_stock"] = o.OtherCapitalStock
	}
	if !IsNil(o.AdditionalPaidInCapital) {
		toSerialize["additional_paid_in_capital"] = o.AdditionalPaidInCapital
	}
	if !IsNil(o.RetainedEarnings) {
		toSerialize["retained_earnings"] = o.RetainedEarnings
	}
	if !IsNil(o.TreasuryStock) {
		toSerialize["treasury_stock"] = o.TreasuryStock
	}
	if !IsNil(o.TreasurySharesNumber) {
		toSerialize["treasury_shares_number"] = o.TreasurySharesNumber
	}
	if !IsNil(o.OrdinarySharesNumber) {
		toSerialize["ordinary_shares_number"] = o.OrdinarySharesNumber
	}
	if !IsNil(o.PreferredSharesNumber) {
		toSerialize["preferred_shares_number"] = o.PreferredSharesNumber
	}
	if !IsNil(o.ShareIssued) {
		toSerialize["share_issued"] = o.ShareIssued
	}
	return toSerialize, nil
}

type NullableEquityInfoCapitalStock struct {
	value *EquityInfoCapitalStock
	isSet bool
}

func (v NullableEquityInfoCapitalStock) Get() *EquityInfoCapitalStock {
	return v.value
}

func (v *NullableEquityInfoCapitalStock) Set(val *EquityInfoCapitalStock) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInfoCapitalStock) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInfoCapitalStock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInfoCapitalStock(val *EquityInfoCapitalStock) *NullableEquityInfoCapitalStock {
	return &NullableEquityInfoCapitalStock{value: val, isSet: true}
}

func (v NullableEquityInfoCapitalStock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInfoCapitalStock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


