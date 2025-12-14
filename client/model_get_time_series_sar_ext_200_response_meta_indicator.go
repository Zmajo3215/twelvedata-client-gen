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

// checks if the GetTimeSeriesSarExt200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSarExt200ResponseMetaIndicator{}

// GetTimeSeriesSarExt200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesSarExt200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name *string `json:"name,omitempty"`
	// The initial value for the indicator calculation
	StartValue *float64 `json:"start_value,omitempty"`
	// The adjustment applied when the indicator's direction changes
	OffsetOnReverse *float64 `json:"offset_on_reverse,omitempty"`
	// The maximum acceleration value for long positions
	AccelerationLimitLong *float64 `json:"acceleration_limit_long,omitempty"`
	// The acceleration value for long positions
	AccelerationLong *float64 `json:"acceleration_long,omitempty"`
	// The highest allowed acceleration for long positions
	AccelerationMaxLong *float64 `json:"acceleration_max_long,omitempty"`
	// The maximum acceleration value for short positions
	AccelerationLimitShort *float64 `json:"acceleration_limit_short,omitempty"`
	// The acceleration value for short positions
	AccelerationShort *float64 `json:"acceleration_short,omitempty"`
	// The highest allowed acceleration for short positions
	AccelerationMaxShort *float64 `json:"acceleration_max_short,omitempty"`
}

// NewGetTimeSeriesSarExt200ResponseMetaIndicator instantiates a new GetTimeSeriesSarExt200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSarExt200ResponseMetaIndicator() *GetTimeSeriesSarExt200ResponseMetaIndicator {
	this := GetTimeSeriesSarExt200ResponseMetaIndicator{}
	return &this
}

// NewGetTimeSeriesSarExt200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesSarExt200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSarExt200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesSarExt200ResponseMetaIndicator {
	this := GetTimeSeriesSarExt200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetName(v string) {
	o.Name = &v
}

// GetStartValue returns the StartValue field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetStartValue() float64 {
	if o == nil || IsNil(o.StartValue) {
		var ret float64
		return ret
	}
	return *o.StartValue
}

// GetStartValueOk returns a tuple with the StartValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetStartValueOk() (*float64, bool) {
	if o == nil || IsNil(o.StartValue) {
		return nil, false
	}
	return o.StartValue, true
}

// HasStartValue returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasStartValue() bool {
	if o != nil && !IsNil(o.StartValue) {
		return true
	}

	return false
}

// SetStartValue gets a reference to the given float64 and assigns it to the StartValue field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetStartValue(v float64) {
	o.StartValue = &v
}

// GetOffsetOnReverse returns the OffsetOnReverse field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetOffsetOnReverse() float64 {
	if o == nil || IsNil(o.OffsetOnReverse) {
		var ret float64
		return ret
	}
	return *o.OffsetOnReverse
}

// GetOffsetOnReverseOk returns a tuple with the OffsetOnReverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetOffsetOnReverseOk() (*float64, bool) {
	if o == nil || IsNil(o.OffsetOnReverse) {
		return nil, false
	}
	return o.OffsetOnReverse, true
}

// HasOffsetOnReverse returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasOffsetOnReverse() bool {
	if o != nil && !IsNil(o.OffsetOnReverse) {
		return true
	}

	return false
}

// SetOffsetOnReverse gets a reference to the given float64 and assigns it to the OffsetOnReverse field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetOffsetOnReverse(v float64) {
	o.OffsetOnReverse = &v
}

// GetAccelerationLimitLong returns the AccelerationLimitLong field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitLong() float64 {
	if o == nil || IsNil(o.AccelerationLimitLong) {
		var ret float64
		return ret
	}
	return *o.AccelerationLimitLong
}

// GetAccelerationLimitLongOk returns a tuple with the AccelerationLimitLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitLongOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationLimitLong) {
		return nil, false
	}
	return o.AccelerationLimitLong, true
}

// HasAccelerationLimitLong returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationLimitLong() bool {
	if o != nil && !IsNil(o.AccelerationLimitLong) {
		return true
	}

	return false
}

// SetAccelerationLimitLong gets a reference to the given float64 and assigns it to the AccelerationLimitLong field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLimitLong(v float64) {
	o.AccelerationLimitLong = &v
}

// GetAccelerationLong returns the AccelerationLong field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLong() float64 {
	if o == nil || IsNil(o.AccelerationLong) {
		var ret float64
		return ret
	}
	return *o.AccelerationLong
}

// GetAccelerationLongOk returns a tuple with the AccelerationLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLongOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationLong) {
		return nil, false
	}
	return o.AccelerationLong, true
}

// HasAccelerationLong returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationLong() bool {
	if o != nil && !IsNil(o.AccelerationLong) {
		return true
	}

	return false
}

// SetAccelerationLong gets a reference to the given float64 and assigns it to the AccelerationLong field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLong(v float64) {
	o.AccelerationLong = &v
}

// GetAccelerationMaxLong returns the AccelerationMaxLong field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxLong() float64 {
	if o == nil || IsNil(o.AccelerationMaxLong) {
		var ret float64
		return ret
	}
	return *o.AccelerationMaxLong
}

// GetAccelerationMaxLongOk returns a tuple with the AccelerationMaxLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxLongOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationMaxLong) {
		return nil, false
	}
	return o.AccelerationMaxLong, true
}

// HasAccelerationMaxLong returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationMaxLong() bool {
	if o != nil && !IsNil(o.AccelerationMaxLong) {
		return true
	}

	return false
}

// SetAccelerationMaxLong gets a reference to the given float64 and assigns it to the AccelerationMaxLong field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationMaxLong(v float64) {
	o.AccelerationMaxLong = &v
}

// GetAccelerationLimitShort returns the AccelerationLimitShort field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitShort() float64 {
	if o == nil || IsNil(o.AccelerationLimitShort) {
		var ret float64
		return ret
	}
	return *o.AccelerationLimitShort
}

// GetAccelerationLimitShortOk returns a tuple with the AccelerationLimitShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitShortOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationLimitShort) {
		return nil, false
	}
	return o.AccelerationLimitShort, true
}

// HasAccelerationLimitShort returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationLimitShort() bool {
	if o != nil && !IsNil(o.AccelerationLimitShort) {
		return true
	}

	return false
}

// SetAccelerationLimitShort gets a reference to the given float64 and assigns it to the AccelerationLimitShort field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLimitShort(v float64) {
	o.AccelerationLimitShort = &v
}

// GetAccelerationShort returns the AccelerationShort field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationShort() float64 {
	if o == nil || IsNil(o.AccelerationShort) {
		var ret float64
		return ret
	}
	return *o.AccelerationShort
}

// GetAccelerationShortOk returns a tuple with the AccelerationShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationShortOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationShort) {
		return nil, false
	}
	return o.AccelerationShort, true
}

// HasAccelerationShort returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationShort() bool {
	if o != nil && !IsNil(o.AccelerationShort) {
		return true
	}

	return false
}

// SetAccelerationShort gets a reference to the given float64 and assigns it to the AccelerationShort field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationShort(v float64) {
	o.AccelerationShort = &v
}

// GetAccelerationMaxShort returns the AccelerationMaxShort field value if set, zero value otherwise.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxShort() float64 {
	if o == nil || IsNil(o.AccelerationMaxShort) {
		var ret float64
		return ret
	}
	return *o.AccelerationMaxShort
}

// GetAccelerationMaxShortOk returns a tuple with the AccelerationMaxShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxShortOk() (*float64, bool) {
	if o == nil || IsNil(o.AccelerationMaxShort) {
		return nil, false
	}
	return o.AccelerationMaxShort, true
}

// HasAccelerationMaxShort returns a boolean if a field has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) HasAccelerationMaxShort() bool {
	if o != nil && !IsNil(o.AccelerationMaxShort) {
		return true
	}

	return false
}

// SetAccelerationMaxShort gets a reference to the given float64 and assigns it to the AccelerationMaxShort field.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationMaxShort(v float64) {
	o.AccelerationMaxShort = &v
}

func (o GetTimeSeriesSarExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSarExt200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartValue) {
		toSerialize["start_value"] = o.StartValue
	}
	if !IsNil(o.OffsetOnReverse) {
		toSerialize["offset_on_reverse"] = o.OffsetOnReverse
	}
	if !IsNil(o.AccelerationLimitLong) {
		toSerialize["acceleration_limit_long"] = o.AccelerationLimitLong
	}
	if !IsNil(o.AccelerationLong) {
		toSerialize["acceleration_long"] = o.AccelerationLong
	}
	if !IsNil(o.AccelerationMaxLong) {
		toSerialize["acceleration_max_long"] = o.AccelerationMaxLong
	}
	if !IsNil(o.AccelerationLimitShort) {
		toSerialize["acceleration_limit_short"] = o.AccelerationLimitShort
	}
	if !IsNil(o.AccelerationShort) {
		toSerialize["acceleration_short"] = o.AccelerationShort
	}
	if !IsNil(o.AccelerationMaxShort) {
		toSerialize["acceleration_max_short"] = o.AccelerationMaxShort
	}
	return toSerialize, nil
}

type NullableGetTimeSeriesSarExt200ResponseMetaIndicator struct {
	value *GetTimeSeriesSarExt200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Get() *GetTimeSeriesSarExt200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Set(val *GetTimeSeriesSarExt200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSarExt200ResponseMetaIndicator(val *GetTimeSeriesSarExt200ResponseMetaIndicator) *NullableGetTimeSeriesSarExt200ResponseMetaIndicator {
	return &NullableGetTimeSeriesSarExt200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


