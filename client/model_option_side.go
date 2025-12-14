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

// checks if the OptionSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionSide{}

// OptionSide struct for OptionSide
type OptionSide struct {
	Ask *float64 `json:"ask,omitempty"`
	Bid *float64 `json:"bid,omitempty"`
	Change *float64 `json:"change,omitempty"`
	ContractName *string `json:"contract_name,omitempty"`
	ImpliedVolatility *float64 `json:"implied_volatility,omitempty"`
	InTheMoney *bool `json:"in_the_money,omitempty"`
	LastPrice *float64 `json:"last_price,omitempty"`
	LastTradeDate *string `json:"last_trade_date,omitempty"`
	OpenInterest *int64 `json:"open_interest,omitempty"`
	OptionId *string `json:"option_id,omitempty"`
	PercentChange *float64 `json:"percent_change,omitempty"`
	Strike *float64 `json:"strike,omitempty"`
	Volume *int64 `json:"volume,omitempty"`
}

// NewOptionSide instantiates a new OptionSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionSide() *OptionSide {
	this := OptionSide{}
	return &this
}

// NewOptionSideWithDefaults instantiates a new OptionSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionSideWithDefaults() *OptionSide {
	this := OptionSide{}
	return &this
}

// GetAsk returns the Ask field value if set, zero value otherwise.
func (o *OptionSide) GetAsk() float64 {
	if o == nil || IsNil(o.Ask) {
		var ret float64
		return ret
	}
	return *o.Ask
}

// GetAskOk returns a tuple with the Ask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetAskOk() (*float64, bool) {
	if o == nil || IsNil(o.Ask) {
		return nil, false
	}
	return o.Ask, true
}

// HasAsk returns a boolean if a field has been set.
func (o *OptionSide) HasAsk() bool {
	if o != nil && !IsNil(o.Ask) {
		return true
	}

	return false
}

// SetAsk gets a reference to the given float64 and assigns it to the Ask field.
func (o *OptionSide) SetAsk(v float64) {
	o.Ask = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *OptionSide) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *OptionSide) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *OptionSide) SetBid(v float64) {
	o.Bid = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *OptionSide) GetChange() float64 {
	if o == nil || IsNil(o.Change) {
		var ret float64
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *OptionSide) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given float64 and assigns it to the Change field.
func (o *OptionSide) SetChange(v float64) {
	o.Change = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *OptionSide) GetContractName() string {
	if o == nil || IsNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetContractNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractName) {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *OptionSide) HasContractName() bool {
	if o != nil && !IsNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *OptionSide) SetContractName(v string) {
	o.ContractName = &v
}

// GetImpliedVolatility returns the ImpliedVolatility field value if set, zero value otherwise.
func (o *OptionSide) GetImpliedVolatility() float64 {
	if o == nil || IsNil(o.ImpliedVolatility) {
		var ret float64
		return ret
	}
	return *o.ImpliedVolatility
}

// GetImpliedVolatilityOk returns a tuple with the ImpliedVolatility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetImpliedVolatilityOk() (*float64, bool) {
	if o == nil || IsNil(o.ImpliedVolatility) {
		return nil, false
	}
	return o.ImpliedVolatility, true
}

// HasImpliedVolatility returns a boolean if a field has been set.
func (o *OptionSide) HasImpliedVolatility() bool {
	if o != nil && !IsNil(o.ImpliedVolatility) {
		return true
	}

	return false
}

// SetImpliedVolatility gets a reference to the given float64 and assigns it to the ImpliedVolatility field.
func (o *OptionSide) SetImpliedVolatility(v float64) {
	o.ImpliedVolatility = &v
}

// GetInTheMoney returns the InTheMoney field value if set, zero value otherwise.
func (o *OptionSide) GetInTheMoney() bool {
	if o == nil || IsNil(o.InTheMoney) {
		var ret bool
		return ret
	}
	return *o.InTheMoney
}

// GetInTheMoneyOk returns a tuple with the InTheMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetInTheMoneyOk() (*bool, bool) {
	if o == nil || IsNil(o.InTheMoney) {
		return nil, false
	}
	return o.InTheMoney, true
}

// HasInTheMoney returns a boolean if a field has been set.
func (o *OptionSide) HasInTheMoney() bool {
	if o != nil && !IsNil(o.InTheMoney) {
		return true
	}

	return false
}

// SetInTheMoney gets a reference to the given bool and assigns it to the InTheMoney field.
func (o *OptionSide) SetInTheMoney(v bool) {
	o.InTheMoney = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *OptionSide) GetLastPrice() float64 {
	if o == nil || IsNil(o.LastPrice) {
		var ret float64
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetLastPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *OptionSide) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given float64 and assigns it to the LastPrice field.
func (o *OptionSide) SetLastPrice(v float64) {
	o.LastPrice = &v
}

// GetLastTradeDate returns the LastTradeDate field value if set, zero value otherwise.
func (o *OptionSide) GetLastTradeDate() string {
	if o == nil || IsNil(o.LastTradeDate) {
		var ret string
		return ret
	}
	return *o.LastTradeDate
}

// GetLastTradeDateOk returns a tuple with the LastTradeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetLastTradeDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastTradeDate) {
		return nil, false
	}
	return o.LastTradeDate, true
}

// HasLastTradeDate returns a boolean if a field has been set.
func (o *OptionSide) HasLastTradeDate() bool {
	if o != nil && !IsNil(o.LastTradeDate) {
		return true
	}

	return false
}

// SetLastTradeDate gets a reference to the given string and assigns it to the LastTradeDate field.
func (o *OptionSide) SetLastTradeDate(v string) {
	o.LastTradeDate = &v
}

// GetOpenInterest returns the OpenInterest field value if set, zero value otherwise.
func (o *OptionSide) GetOpenInterest() int64 {
	if o == nil || IsNil(o.OpenInterest) {
		var ret int64
		return ret
	}
	return *o.OpenInterest
}

// GetOpenInterestOk returns a tuple with the OpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetOpenInterestOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenInterest) {
		return nil, false
	}
	return o.OpenInterest, true
}

// HasOpenInterest returns a boolean if a field has been set.
func (o *OptionSide) HasOpenInterest() bool {
	if o != nil && !IsNil(o.OpenInterest) {
		return true
	}

	return false
}

// SetOpenInterest gets a reference to the given int64 and assigns it to the OpenInterest field.
func (o *OptionSide) SetOpenInterest(v int64) {
	o.OpenInterest = &v
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *OptionSide) GetOptionId() string {
	if o == nil || IsNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionId) {
		return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *OptionSide) HasOptionId() bool {
	if o != nil && !IsNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *OptionSide) SetOptionId(v string) {
	o.OptionId = &v
}

// GetPercentChange returns the PercentChange field value if set, zero value otherwise.
func (o *OptionSide) GetPercentChange() float64 {
	if o == nil || IsNil(o.PercentChange) {
		var ret float64
		return ret
	}
	return *o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetPercentChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentChange) {
		return nil, false
	}
	return o.PercentChange, true
}

// HasPercentChange returns a boolean if a field has been set.
func (o *OptionSide) HasPercentChange() bool {
	if o != nil && !IsNil(o.PercentChange) {
		return true
	}

	return false
}

// SetPercentChange gets a reference to the given float64 and assigns it to the PercentChange field.
func (o *OptionSide) SetPercentChange(v float64) {
	o.PercentChange = &v
}

// GetStrike returns the Strike field value if set, zero value otherwise.
func (o *OptionSide) GetStrike() float64 {
	if o == nil || IsNil(o.Strike) {
		var ret float64
		return ret
	}
	return *o.Strike
}

// GetStrikeOk returns a tuple with the Strike field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetStrikeOk() (*float64, bool) {
	if o == nil || IsNil(o.Strike) {
		return nil, false
	}
	return o.Strike, true
}

// HasStrike returns a boolean if a field has been set.
func (o *OptionSide) HasStrike() bool {
	if o != nil && !IsNil(o.Strike) {
		return true
	}

	return false
}

// SetStrike gets a reference to the given float64 and assigns it to the Strike field.
func (o *OptionSide) SetStrike(v float64) {
	o.Strike = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *OptionSide) GetVolume() int64 {
	if o == nil || IsNil(o.Volume) {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionSide) GetVolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *OptionSide) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *OptionSide) SetVolume(v int64) {
	o.Volume = &v
}

func (o OptionSide) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ask) {
		toSerialize["ask"] = o.Ask
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.ContractName) {
		toSerialize["contract_name"] = o.ContractName
	}
	if !IsNil(o.ImpliedVolatility) {
		toSerialize["implied_volatility"] = o.ImpliedVolatility
	}
	if !IsNil(o.InTheMoney) {
		toSerialize["in_the_money"] = o.InTheMoney
	}
	if !IsNil(o.LastPrice) {
		toSerialize["last_price"] = o.LastPrice
	}
	if !IsNil(o.LastTradeDate) {
		toSerialize["last_trade_date"] = o.LastTradeDate
	}
	if !IsNil(o.OpenInterest) {
		toSerialize["open_interest"] = o.OpenInterest
	}
	if !IsNil(o.OptionId) {
		toSerialize["option_id"] = o.OptionId
	}
	if !IsNil(o.PercentChange) {
		toSerialize["percent_change"] = o.PercentChange
	}
	if !IsNil(o.Strike) {
		toSerialize["strike"] = o.Strike
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableOptionSide struct {
	value *OptionSide
	isSet bool
}

func (v NullableOptionSide) Get() *OptionSide {
	return v.value
}

func (v *NullableOptionSide) Set(val *OptionSide) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionSide) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionSide(val *OptionSide) *NullableOptionSide {
	return &NullableOptionSide{value: val, isSet: true}
}

func (v NullableOptionSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


