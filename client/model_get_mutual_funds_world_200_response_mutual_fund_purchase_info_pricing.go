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

// checks if the GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}

// GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing Pricing information for the mutual fund
type GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing struct {
	// Net Asset Value: fund value minus liabilities
	Nav *float64 `json:"nav,omitempty"`
	// Lowest price of the fund over the last year
	Var12MonthLow *float64 `json:"12_month_low,omitempty"`
	// Highest price of the fund over the last year
	Var12MonthHigh *float64 `json:"12_month_high,omitempty"`
	// Fund price at the end of the last month
	LastMonth *float64 `json:"last_month,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricingWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricingWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}
	return &this
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetNav() float64 {
	if o == nil || IsNil(o.Nav) {
		var ret float64
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetNavOk() (*float64, bool) {
	if o == nil || IsNil(o.Nav) {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasNav() bool {
	if o != nil && !IsNil(o.Nav) {
		return true
	}

	return false
}

// SetNav gets a reference to the given float64 and assigns it to the Nav field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetNav(v float64) {
	o.Nav = &v
}

// GetVar12MonthLow returns the Var12MonthLow field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthLow() float64 {
	if o == nil || IsNil(o.Var12MonthLow) {
		var ret float64
		return ret
	}
	return *o.Var12MonthLow
}

// GetVar12MonthLowOk returns a tuple with the Var12MonthLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthLowOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthLow) {
		return nil, false
	}
	return o.Var12MonthLow, true
}

// HasVar12MonthLow returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasVar12MonthLow() bool {
	if o != nil && !IsNil(o.Var12MonthLow) {
		return true
	}

	return false
}

// SetVar12MonthLow gets a reference to the given float64 and assigns it to the Var12MonthLow field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetVar12MonthLow(v float64) {
	o.Var12MonthLow = &v
}

// GetVar12MonthHigh returns the Var12MonthHigh field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthHigh() float64 {
	if o == nil || IsNil(o.Var12MonthHigh) {
		var ret float64
		return ret
	}
	return *o.Var12MonthHigh
}

// GetVar12MonthHighOk returns a tuple with the Var12MonthHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthHighOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthHigh) {
		return nil, false
	}
	return o.Var12MonthHigh, true
}

// HasVar12MonthHigh returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasVar12MonthHigh() bool {
	if o != nil && !IsNil(o.Var12MonthHigh) {
		return true
	}

	return false
}

// SetVar12MonthHigh gets a reference to the given float64 and assigns it to the Var12MonthHigh field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetVar12MonthHigh(v float64) {
	o.Var12MonthHigh = &v
}

// GetLastMonth returns the LastMonth field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetLastMonth() float64 {
	if o == nil || IsNil(o.LastMonth) {
		var ret float64
		return ret
	}
	return *o.LastMonth
}

// GetLastMonthOk returns a tuple with the LastMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetLastMonthOk() (*float64, bool) {
	if o == nil || IsNil(o.LastMonth) {
		return nil, false
	}
	return o.LastMonth, true
}

// HasLastMonth returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasLastMonth() bool {
	if o != nil && !IsNil(o.LastMonth) {
		return true
	}

	return false
}

// SetLastMonth gets a reference to the given float64 and assigns it to the LastMonth field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetLastMonth(v float64) {
	o.LastMonth = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nav) {
		toSerialize["nav"] = o.Nav
	}
	if !IsNil(o.Var12MonthLow) {
		toSerialize["12_month_low"] = o.Var12MonthLow
	}
	if !IsNil(o.Var12MonthHigh) {
		toSerialize["12_month_high"] = o.Var12MonthHigh
	}
	if !IsNil(o.LastMonth) {
		toSerialize["last_month"] = o.LastMonth
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing struct {
	value *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Get() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Set(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


