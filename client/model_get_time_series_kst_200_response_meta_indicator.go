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

// checks if the GetTimeSeriesKst200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKst200ResponseMetaIndicator{}

// GetTimeSeriesKst200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesKst200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name *string `json:"name,omitempty"`
	// The time period for the first Rate of Change calculation
	RocPeriod1 *int64 `json:"roc_period_1,omitempty"`
	// The time period for the second Rate of Change calculation
	RocPeriod2 *int64 `json:"roc_period_2,omitempty"`
	// The time period for the third Rate of Change calculation
	RocPeriod3 *int64 `json:"roc_period_3,omitempty"`
	// The time period for the forth Rate of Change calculation
	RocPeriod4 *int64 `json:"roc_period_4,omitempty"`
	// The time period for the first Simple Moving Average
	SmaPeriod1 *int64 `json:"sma_period_1,omitempty"`
	// The time period for the second Simple Moving Average
	SmaPeriod2 *int64 `json:"sma_period_2,omitempty"`
	// The time period for the third Simple Moving Average
	SmaPeriod3 *int64 `json:"sma_period_3,omitempty"`
	// The time period for the forth Simple Moving Average
	SmaPeriod4 *int64 `json:"sma_period_4,omitempty"`
	// The time period used for generating the signal line
	SignalPeriod *int64 `json:"signal_period,omitempty"`
}

// NewGetTimeSeriesKst200ResponseMetaIndicator instantiates a new GetTimeSeriesKst200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKst200ResponseMetaIndicator() *GetTimeSeriesKst200ResponseMetaIndicator {
	this := GetTimeSeriesKst200ResponseMetaIndicator{}
	return &this
}

// NewGetTimeSeriesKst200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesKst200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKst200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesKst200ResponseMetaIndicator {
	this := GetTimeSeriesKst200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetName(v string) {
	o.Name = &v
}

// GetRocPeriod1 returns the RocPeriod1 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod1() int64 {
	if o == nil || IsNil(o.RocPeriod1) {
		var ret int64
		return ret
	}
	return *o.RocPeriod1
}

// GetRocPeriod1Ok returns a tuple with the RocPeriod1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod1Ok() (*int64, bool) {
	if o == nil || IsNil(o.RocPeriod1) {
		return nil, false
	}
	return o.RocPeriod1, true
}

// HasRocPeriod1 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasRocPeriod1() bool {
	if o != nil && !IsNil(o.RocPeriod1) {
		return true
	}

	return false
}

// SetRocPeriod1 gets a reference to the given int64 and assigns it to the RocPeriod1 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod1(v int64) {
	o.RocPeriod1 = &v
}

// GetRocPeriod2 returns the RocPeriod2 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod2() int64 {
	if o == nil || IsNil(o.RocPeriod2) {
		var ret int64
		return ret
	}
	return *o.RocPeriod2
}

// GetRocPeriod2Ok returns a tuple with the RocPeriod2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod2Ok() (*int64, bool) {
	if o == nil || IsNil(o.RocPeriod2) {
		return nil, false
	}
	return o.RocPeriod2, true
}

// HasRocPeriod2 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasRocPeriod2() bool {
	if o != nil && !IsNil(o.RocPeriod2) {
		return true
	}

	return false
}

// SetRocPeriod2 gets a reference to the given int64 and assigns it to the RocPeriod2 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod2(v int64) {
	o.RocPeriod2 = &v
}

// GetRocPeriod3 returns the RocPeriod3 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod3() int64 {
	if o == nil || IsNil(o.RocPeriod3) {
		var ret int64
		return ret
	}
	return *o.RocPeriod3
}

// GetRocPeriod3Ok returns a tuple with the RocPeriod3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod3Ok() (*int64, bool) {
	if o == nil || IsNil(o.RocPeriod3) {
		return nil, false
	}
	return o.RocPeriod3, true
}

// HasRocPeriod3 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasRocPeriod3() bool {
	if o != nil && !IsNil(o.RocPeriod3) {
		return true
	}

	return false
}

// SetRocPeriod3 gets a reference to the given int64 and assigns it to the RocPeriod3 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod3(v int64) {
	o.RocPeriod3 = &v
}

// GetRocPeriod4 returns the RocPeriod4 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod4() int64 {
	if o == nil || IsNil(o.RocPeriod4) {
		var ret int64
		return ret
	}
	return *o.RocPeriod4
}

// GetRocPeriod4Ok returns a tuple with the RocPeriod4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod4Ok() (*int64, bool) {
	if o == nil || IsNil(o.RocPeriod4) {
		return nil, false
	}
	return o.RocPeriod4, true
}

// HasRocPeriod4 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasRocPeriod4() bool {
	if o != nil && !IsNil(o.RocPeriod4) {
		return true
	}

	return false
}

// SetRocPeriod4 gets a reference to the given int64 and assigns it to the RocPeriod4 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod4(v int64) {
	o.RocPeriod4 = &v
}

// GetSmaPeriod1 returns the SmaPeriod1 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod1() int64 {
	if o == nil || IsNil(o.SmaPeriod1) {
		var ret int64
		return ret
	}
	return *o.SmaPeriod1
}

// GetSmaPeriod1Ok returns a tuple with the SmaPeriod1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod1Ok() (*int64, bool) {
	if o == nil || IsNil(o.SmaPeriod1) {
		return nil, false
	}
	return o.SmaPeriod1, true
}

// HasSmaPeriod1 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasSmaPeriod1() bool {
	if o != nil && !IsNil(o.SmaPeriod1) {
		return true
	}

	return false
}

// SetSmaPeriod1 gets a reference to the given int64 and assigns it to the SmaPeriod1 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod1(v int64) {
	o.SmaPeriod1 = &v
}

// GetSmaPeriod2 returns the SmaPeriod2 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod2() int64 {
	if o == nil || IsNil(o.SmaPeriod2) {
		var ret int64
		return ret
	}
	return *o.SmaPeriod2
}

// GetSmaPeriod2Ok returns a tuple with the SmaPeriod2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod2Ok() (*int64, bool) {
	if o == nil || IsNil(o.SmaPeriod2) {
		return nil, false
	}
	return o.SmaPeriod2, true
}

// HasSmaPeriod2 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasSmaPeriod2() bool {
	if o != nil && !IsNil(o.SmaPeriod2) {
		return true
	}

	return false
}

// SetSmaPeriod2 gets a reference to the given int64 and assigns it to the SmaPeriod2 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod2(v int64) {
	o.SmaPeriod2 = &v
}

// GetSmaPeriod3 returns the SmaPeriod3 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod3() int64 {
	if o == nil || IsNil(o.SmaPeriod3) {
		var ret int64
		return ret
	}
	return *o.SmaPeriod3
}

// GetSmaPeriod3Ok returns a tuple with the SmaPeriod3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod3Ok() (*int64, bool) {
	if o == nil || IsNil(o.SmaPeriod3) {
		return nil, false
	}
	return o.SmaPeriod3, true
}

// HasSmaPeriod3 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasSmaPeriod3() bool {
	if o != nil && !IsNil(o.SmaPeriod3) {
		return true
	}

	return false
}

// SetSmaPeriod3 gets a reference to the given int64 and assigns it to the SmaPeriod3 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod3(v int64) {
	o.SmaPeriod3 = &v
}

// GetSmaPeriod4 returns the SmaPeriod4 field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod4() int64 {
	if o == nil || IsNil(o.SmaPeriod4) {
		var ret int64
		return ret
	}
	return *o.SmaPeriod4
}

// GetSmaPeriod4Ok returns a tuple with the SmaPeriod4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod4Ok() (*int64, bool) {
	if o == nil || IsNil(o.SmaPeriod4) {
		return nil, false
	}
	return o.SmaPeriod4, true
}

// HasSmaPeriod4 returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasSmaPeriod4() bool {
	if o != nil && !IsNil(o.SmaPeriod4) {
		return true
	}

	return false
}

// SetSmaPeriod4 gets a reference to the given int64 and assigns it to the SmaPeriod4 field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod4(v int64) {
	o.SmaPeriod4 = &v
}

// GetSignalPeriod returns the SignalPeriod field value if set, zero value otherwise.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSignalPeriod() int64 {
	if o == nil || IsNil(o.SignalPeriod) {
		var ret int64
		return ret
	}
	return *o.SignalPeriod
}

// GetSignalPeriodOk returns a tuple with the SignalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.SignalPeriod) {
		return nil, false
	}
	return o.SignalPeriod, true
}

// HasSignalPeriod returns a boolean if a field has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) HasSignalPeriod() bool {
	if o != nil && !IsNil(o.SignalPeriod) {
		return true
	}

	return false
}

// SetSignalPeriod gets a reference to the given int64 and assigns it to the SignalPeriod field.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSignalPeriod(v int64) {
	o.SignalPeriod = &v
}

func (o GetTimeSeriesKst200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKst200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RocPeriod1) {
		toSerialize["roc_period_1"] = o.RocPeriod1
	}
	if !IsNil(o.RocPeriod2) {
		toSerialize["roc_period_2"] = o.RocPeriod2
	}
	if !IsNil(o.RocPeriod3) {
		toSerialize["roc_period_3"] = o.RocPeriod3
	}
	if !IsNil(o.RocPeriod4) {
		toSerialize["roc_period_4"] = o.RocPeriod4
	}
	if !IsNil(o.SmaPeriod1) {
		toSerialize["sma_period_1"] = o.SmaPeriod1
	}
	if !IsNil(o.SmaPeriod2) {
		toSerialize["sma_period_2"] = o.SmaPeriod2
	}
	if !IsNil(o.SmaPeriod3) {
		toSerialize["sma_period_3"] = o.SmaPeriod3
	}
	if !IsNil(o.SmaPeriod4) {
		toSerialize["sma_period_4"] = o.SmaPeriod4
	}
	if !IsNil(o.SignalPeriod) {
		toSerialize["signal_period"] = o.SignalPeriod
	}
	return toSerialize, nil
}

type NullableGetTimeSeriesKst200ResponseMetaIndicator struct {
	value *GetTimeSeriesKst200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) Get() *GetTimeSeriesKst200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) Set(val *GetTimeSeriesKst200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKst200ResponseMetaIndicator(val *GetTimeSeriesKst200ResponseMetaIndicator) *NullableGetTimeSeriesKst200ResponseMetaIndicator {
	return &NullableGetTimeSeriesKst200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


