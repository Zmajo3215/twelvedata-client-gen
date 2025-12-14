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

// checks if the CashFlowStructFinancingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowStructFinancingActivities{}

// CashFlowStructFinancingActivities Financing activities section
type CashFlowStructFinancingActivities struct {
	// Refers to the issuance of any financial obligations that extend beyond a 12 months period
	LongTermDebtIssuance *float64 `json:"long_term_debt_issuance,omitempty"`
	// Refers to the payments of any financial obligations that extend beyond a 12 months period
	LongTermDebtPayments *float64 `json:"long_term_debt_payments,omitempty"`
	// Refers to the issuance of any financial obligations that are expected to be paid off within a year
	ShortTermDebtIssuance *float64 `json:"short_term_debt_issuance,omitempty"`
	// Represents a transaction whereby a company issues its own shares to the marketplace
	CommonStockIssuance *float64 `json:"common_stock_issuance,omitempty"`
	// Represents a transaction whereby a company buys back its own shares from the marketplace
	CommonStockRepurchase *float64 `json:"common_stock_repurchase,omitempty"`
	// Returns value of payment doled out by a company to its stockholders in the form of periodic distributions of cash
	CommonDividends *float64 `json:"common_dividends,omitempty"`
	// Represents other financing charges
	OtherFinancingCharges *float64 `json:"other_financing_charges,omitempty"`
	// Returns cash flow from financing activities (CFF), which shows the net flows of cash that are used to fund the company
	FinancingCashFlow *float64 `json:"financing_cash_flow,omitempty"`
}

// NewCashFlowStructFinancingActivities instantiates a new CashFlowStructFinancingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowStructFinancingActivities() *CashFlowStructFinancingActivities {
	this := CashFlowStructFinancingActivities{}
	return &this
}

// NewCashFlowStructFinancingActivitiesWithDefaults instantiates a new CashFlowStructFinancingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowStructFinancingActivitiesWithDefaults() *CashFlowStructFinancingActivities {
	this := CashFlowStructFinancingActivities{}
	return &this
}

// GetLongTermDebtIssuance returns the LongTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuance() float64 {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtIssuance
}

// GetLongTermDebtIssuanceOk returns a tuple with the LongTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		return nil, false
	}
	return o.LongTermDebtIssuance, true
}

// HasLongTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasLongTermDebtIssuance() bool {
	if o != nil && !IsNil(o.LongTermDebtIssuance) {
		return true
	}

	return false
}

// SetLongTermDebtIssuance gets a reference to the given float64 and assigns it to the LongTermDebtIssuance field.
func (o *CashFlowStructFinancingActivities) SetLongTermDebtIssuance(v float64) {
	o.LongTermDebtIssuance = &v
}

// GetLongTermDebtPayments returns the LongTermDebtPayments field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtPayments() float64 {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtPayments
}

// GetLongTermDebtPaymentsOk returns a tuple with the LongTermDebtPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		return nil, false
	}
	return o.LongTermDebtPayments, true
}

// HasLongTermDebtPayments returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasLongTermDebtPayments() bool {
	if o != nil && !IsNil(o.LongTermDebtPayments) {
		return true
	}

	return false
}

// SetLongTermDebtPayments gets a reference to the given float64 and assigns it to the LongTermDebtPayments field.
func (o *CashFlowStructFinancingActivities) SetLongTermDebtPayments(v float64) {
	o.LongTermDebtPayments = &v
}

// GetShortTermDebtIssuance returns the ShortTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuance() float64 {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.ShortTermDebtIssuance
}

// GetShortTermDebtIssuanceOk returns a tuple with the ShortTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		return nil, false
	}
	return o.ShortTermDebtIssuance, true
}

// HasShortTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasShortTermDebtIssuance() bool {
	if o != nil && !IsNil(o.ShortTermDebtIssuance) {
		return true
	}

	return false
}

// SetShortTermDebtIssuance gets a reference to the given float64 and assigns it to the ShortTermDebtIssuance field.
func (o *CashFlowStructFinancingActivities) SetShortTermDebtIssuance(v float64) {
	o.ShortTermDebtIssuance = &v
}

// GetCommonStockIssuance returns the CommonStockIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonStockIssuance() float64 {
	if o == nil || IsNil(o.CommonStockIssuance) {
		var ret float64
		return ret
	}
	return *o.CommonStockIssuance
}

// GetCommonStockIssuanceOk returns a tuple with the CommonStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockIssuance) {
		return nil, false
	}
	return o.CommonStockIssuance, true
}

// HasCommonStockIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonStockIssuance() bool {
	if o != nil && !IsNil(o.CommonStockIssuance) {
		return true
	}

	return false
}

// SetCommonStockIssuance gets a reference to the given float64 and assigns it to the CommonStockIssuance field.
func (o *CashFlowStructFinancingActivities) SetCommonStockIssuance(v float64) {
	o.CommonStockIssuance = &v
}

// GetCommonStockRepurchase returns the CommonStockRepurchase field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchase() float64 {
	if o == nil || IsNil(o.CommonStockRepurchase) {
		var ret float64
		return ret
	}
	return *o.CommonStockRepurchase
}

// GetCommonStockRepurchaseOk returns a tuple with the CommonStockRepurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchaseOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockRepurchase) {
		return nil, false
	}
	return o.CommonStockRepurchase, true
}

// HasCommonStockRepurchase returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonStockRepurchase() bool {
	if o != nil && !IsNil(o.CommonStockRepurchase) {
		return true
	}

	return false
}

// SetCommonStockRepurchase gets a reference to the given float64 and assigns it to the CommonStockRepurchase field.
func (o *CashFlowStructFinancingActivities) SetCommonStockRepurchase(v float64) {
	o.CommonStockRepurchase = &v
}

// GetCommonDividends returns the CommonDividends field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonDividends() float64 {
	if o == nil || IsNil(o.CommonDividends) {
		var ret float64
		return ret
	}
	return *o.CommonDividends
}

// GetCommonDividendsOk returns a tuple with the CommonDividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonDividendsOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonDividends) {
		return nil, false
	}
	return o.CommonDividends, true
}

// HasCommonDividends returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonDividends() bool {
	if o != nil && !IsNil(o.CommonDividends) {
		return true
	}

	return false
}

// SetCommonDividends gets a reference to the given float64 and assigns it to the CommonDividends field.
func (o *CashFlowStructFinancingActivities) SetCommonDividends(v float64) {
	o.CommonDividends = &v
}

// GetOtherFinancingCharges returns the OtherFinancingCharges field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetOtherFinancingCharges() float64 {
	if o == nil || IsNil(o.OtherFinancingCharges) {
		var ret float64
		return ret
	}
	return *o.OtherFinancingCharges
}

// GetOtherFinancingChargesOk returns a tuple with the OtherFinancingCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetOtherFinancingChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherFinancingCharges) {
		return nil, false
	}
	return o.OtherFinancingCharges, true
}

// HasOtherFinancingCharges returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasOtherFinancingCharges() bool {
	if o != nil && !IsNil(o.OtherFinancingCharges) {
		return true
	}

	return false
}

// SetOtherFinancingCharges gets a reference to the given float64 and assigns it to the OtherFinancingCharges field.
func (o *CashFlowStructFinancingActivities) SetOtherFinancingCharges(v float64) {
	o.OtherFinancingCharges = &v
}

// GetFinancingCashFlow returns the FinancingCashFlow field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetFinancingCashFlow() float64 {
	if o == nil || IsNil(o.FinancingCashFlow) {
		var ret float64
		return ret
	}
	return *o.FinancingCashFlow
}

// GetFinancingCashFlowOk returns a tuple with the FinancingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetFinancingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancingCashFlow) {
		return nil, false
	}
	return o.FinancingCashFlow, true
}

// HasFinancingCashFlow returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasFinancingCashFlow() bool {
	if o != nil && !IsNil(o.FinancingCashFlow) {
		return true
	}

	return false
}

// SetFinancingCashFlow gets a reference to the given float64 and assigns it to the FinancingCashFlow field.
func (o *CashFlowStructFinancingActivities) SetFinancingCashFlow(v float64) {
	o.FinancingCashFlow = &v
}

func (o CashFlowStructFinancingActivities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowStructFinancingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongTermDebtIssuance) {
		toSerialize["long_term_debt_issuance"] = o.LongTermDebtIssuance
	}
	if !IsNil(o.LongTermDebtPayments) {
		toSerialize["long_term_debt_payments"] = o.LongTermDebtPayments
	}
	if !IsNil(o.ShortTermDebtIssuance) {
		toSerialize["short_term_debt_issuance"] = o.ShortTermDebtIssuance
	}
	if !IsNil(o.CommonStockIssuance) {
		toSerialize["common_stock_issuance"] = o.CommonStockIssuance
	}
	if !IsNil(o.CommonStockRepurchase) {
		toSerialize["common_stock_repurchase"] = o.CommonStockRepurchase
	}
	if !IsNil(o.CommonDividends) {
		toSerialize["common_dividends"] = o.CommonDividends
	}
	if !IsNil(o.OtherFinancingCharges) {
		toSerialize["other_financing_charges"] = o.OtherFinancingCharges
	}
	if !IsNil(o.FinancingCashFlow) {
		toSerialize["financing_cash_flow"] = o.FinancingCashFlow
	}
	return toSerialize, nil
}

type NullableCashFlowStructFinancingActivities struct {
	value *CashFlowStructFinancingActivities
	isSet bool
}

func (v NullableCashFlowStructFinancingActivities) Get() *CashFlowStructFinancingActivities {
	return v.value
}

func (v *NullableCashFlowStructFinancingActivities) Set(val *CashFlowStructFinancingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowStructFinancingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowStructFinancingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowStructFinancingActivities(val *CashFlowStructFinancingActivities) *NullableCashFlowStructFinancingActivities {
	return &NullableCashFlowStructFinancingActivities{value: val, isSet: true}
}

func (v NullableCashFlowStructFinancingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowStructFinancingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


