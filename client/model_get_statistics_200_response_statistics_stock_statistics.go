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

// checks if the GetStatistics200ResponseStatisticsStockStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsStockStatistics{}

// GetStatistics200ResponseStatisticsStockStatistics Stock statistics of the company
type GetStatistics200ResponseStatisticsStockStatistics struct {
	// Refers for the shares outstanding currently held by all its shareholders
	SharesOutstanding *float64 `json:"shares_outstanding,omitempty"`
	// Refers to floating stock is the number of public shares a company has available for trading on the open market
	FloatShares *float64 `json:"float_shares,omitempty"`
	// Refers to the average 10 days volume
	Avg10Volume *int64 `json:"avg_10_volume,omitempty"`
	// Refers to the average 90 days volume
	Avg90Volume *int64 `json:"avg_90_volume,omitempty"`
	// Refers to the number of shares that have been shorted
	SharesShort *int64 `json:"shares_short,omitempty"`
	// Refers to short ratio measure
	ShortRatio *float64 `json:"short_ratio,omitempty"`
	// Refers to the number of shorted shares divided by the number of shares outstanding
	ShortPercentOfSharesOutstanding *float64 `json:"short_percent_of_shares_outstanding,omitempty"`
	// Refers to the percentage of shares held by the company insiders
	PercentHeldByInsiders *float64 `json:"percent_held_by_insiders,omitempty"`
	// Refers to the percentage of shares held by the institutions
	PercentHeldByInstitutions *float64 `json:"percent_held_by_institutions,omitempty"`
}

// NewGetStatistics200ResponseStatisticsStockStatistics instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsStockStatistics() *GetStatistics200ResponseStatisticsStockStatistics {
	this := GetStatistics200ResponseStatisticsStockStatistics{}
	return &this
}

// NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults instantiates a new GetStatistics200ResponseStatisticsStockStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsStockStatisticsWithDefaults() *GetStatistics200ResponseStatisticsStockStatistics {
	this := GetStatistics200ResponseStatisticsStockStatistics{}
	return &this
}

// GetSharesOutstanding returns the SharesOutstanding field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstanding() float64 {
	if o == nil || IsNil(o.SharesOutstanding) {
		var ret float64
		return ret
	}
	return *o.SharesOutstanding
}

// GetSharesOutstandingOk returns a tuple with the SharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesOutstandingOk() (*float64, bool) {
	if o == nil || IsNil(o.SharesOutstanding) {
		return nil, false
	}
	return o.SharesOutstanding, true
}

// HasSharesOutstanding returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesOutstanding() bool {
	if o != nil && !IsNil(o.SharesOutstanding) {
		return true
	}

	return false
}

// SetSharesOutstanding gets a reference to the given float64 and assigns it to the SharesOutstanding field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesOutstanding(v float64) {
	o.SharesOutstanding = &v
}

// GetFloatShares returns the FloatShares field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatShares() float64 {
	if o == nil || IsNil(o.FloatShares) {
		var ret float64
		return ret
	}
	return *o.FloatShares
}

// GetFloatSharesOk returns a tuple with the FloatShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetFloatSharesOk() (*float64, bool) {
	if o == nil || IsNil(o.FloatShares) {
		return nil, false
	}
	return o.FloatShares, true
}

// HasFloatShares returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasFloatShares() bool {
	if o != nil && !IsNil(o.FloatShares) {
		return true
	}

	return false
}

// SetFloatShares gets a reference to the given float64 and assigns it to the FloatShares field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetFloatShares(v float64) {
	o.FloatShares = &v
}

// GetAvg10Volume returns the Avg10Volume field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10Volume() int64 {
	if o == nil || IsNil(o.Avg10Volume) {
		var ret int64
		return ret
	}
	return *o.Avg10Volume
}

// GetAvg10VolumeOk returns a tuple with the Avg10Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg10VolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Avg10Volume) {
		return nil, false
	}
	return o.Avg10Volume, true
}

// HasAvg10Volume returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg10Volume() bool {
	if o != nil && !IsNil(o.Avg10Volume) {
		return true
	}

	return false
}

// SetAvg10Volume gets a reference to the given int64 and assigns it to the Avg10Volume field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg10Volume(v int64) {
	o.Avg10Volume = &v
}

// GetAvg90Volume returns the Avg90Volume field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90Volume() int64 {
	if o == nil || IsNil(o.Avg90Volume) {
		var ret int64
		return ret
	}
	return *o.Avg90Volume
}

// GetAvg90VolumeOk returns a tuple with the Avg90Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetAvg90VolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Avg90Volume) {
		return nil, false
	}
	return o.Avg90Volume, true
}

// HasAvg90Volume returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasAvg90Volume() bool {
	if o != nil && !IsNil(o.Avg90Volume) {
		return true
	}

	return false
}

// SetAvg90Volume gets a reference to the given int64 and assigns it to the Avg90Volume field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetAvg90Volume(v int64) {
	o.Avg90Volume = &v
}

// GetSharesShort returns the SharesShort field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShort() int64 {
	if o == nil || IsNil(o.SharesShort) {
		var ret int64
		return ret
	}
	return *o.SharesShort
}

// GetSharesShortOk returns a tuple with the SharesShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetSharesShortOk() (*int64, bool) {
	if o == nil || IsNil(o.SharesShort) {
		return nil, false
	}
	return o.SharesShort, true
}

// HasSharesShort returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasSharesShort() bool {
	if o != nil && !IsNil(o.SharesShort) {
		return true
	}

	return false
}

// SetSharesShort gets a reference to the given int64 and assigns it to the SharesShort field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetSharesShort(v int64) {
	o.SharesShort = &v
}

// GetShortRatio returns the ShortRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatio() float64 {
	if o == nil || IsNil(o.ShortRatio) {
		var ret float64
		return ret
	}
	return *o.ShortRatio
}

// GetShortRatioOk returns a tuple with the ShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortRatio) {
		return nil, false
	}
	return o.ShortRatio, true
}

// HasShortRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortRatio() bool {
	if o != nil && !IsNil(o.ShortRatio) {
		return true
	}

	return false
}

// SetShortRatio gets a reference to the given float64 and assigns it to the ShortRatio field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortRatio(v float64) {
	o.ShortRatio = &v
}

// GetShortPercentOfSharesOutstanding returns the ShortPercentOfSharesOutstanding field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstanding() float64 {
	if o == nil || IsNil(o.ShortPercentOfSharesOutstanding) {
		var ret float64
		return ret
	}
	return *o.ShortPercentOfSharesOutstanding
}

// GetShortPercentOfSharesOutstandingOk returns a tuple with the ShortPercentOfSharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetShortPercentOfSharesOutstandingOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortPercentOfSharesOutstanding) {
		return nil, false
	}
	return o.ShortPercentOfSharesOutstanding, true
}

// HasShortPercentOfSharesOutstanding returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasShortPercentOfSharesOutstanding() bool {
	if o != nil && !IsNil(o.ShortPercentOfSharesOutstanding) {
		return true
	}

	return false
}

// SetShortPercentOfSharesOutstanding gets a reference to the given float64 and assigns it to the ShortPercentOfSharesOutstanding field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetShortPercentOfSharesOutstanding(v float64) {
	o.ShortPercentOfSharesOutstanding = &v
}

// GetPercentHeldByInsiders returns the PercentHeldByInsiders field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsiders() float64 {
	if o == nil || IsNil(o.PercentHeldByInsiders) {
		var ret float64
		return ret
	}
	return *o.PercentHeldByInsiders
}

// GetPercentHeldByInsidersOk returns a tuple with the PercentHeldByInsiders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInsidersOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentHeldByInsiders) {
		return nil, false
	}
	return o.PercentHeldByInsiders, true
}

// HasPercentHeldByInsiders returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInsiders() bool {
	if o != nil && !IsNil(o.PercentHeldByInsiders) {
		return true
	}

	return false
}

// SetPercentHeldByInsiders gets a reference to the given float64 and assigns it to the PercentHeldByInsiders field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInsiders(v float64) {
	o.PercentHeldByInsiders = &v
}

// GetPercentHeldByInstitutions returns the PercentHeldByInstitutions field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutions() float64 {
	if o == nil || IsNil(o.PercentHeldByInstitutions) {
		var ret float64
		return ret
	}
	return *o.PercentHeldByInstitutions
}

// GetPercentHeldByInstitutionsOk returns a tuple with the PercentHeldByInstitutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) GetPercentHeldByInstitutionsOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentHeldByInstitutions) {
		return nil, false
	}
	return o.PercentHeldByInstitutions, true
}

// HasPercentHeldByInstitutions returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockStatistics) HasPercentHeldByInstitutions() bool {
	if o != nil && !IsNil(o.PercentHeldByInstitutions) {
		return true
	}

	return false
}

// SetPercentHeldByInstitutions gets a reference to the given float64 and assigns it to the PercentHeldByInstitutions field.
func (o *GetStatistics200ResponseStatisticsStockStatistics) SetPercentHeldByInstitutions(v float64) {
	o.PercentHeldByInstitutions = &v
}

func (o GetStatistics200ResponseStatisticsStockStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsStockStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SharesOutstanding) {
		toSerialize["shares_outstanding"] = o.SharesOutstanding
	}
	if !IsNil(o.FloatShares) {
		toSerialize["float_shares"] = o.FloatShares
	}
	if !IsNil(o.Avg10Volume) {
		toSerialize["avg_10_volume"] = o.Avg10Volume
	}
	if !IsNil(o.Avg90Volume) {
		toSerialize["avg_90_volume"] = o.Avg90Volume
	}
	if !IsNil(o.SharesShort) {
		toSerialize["shares_short"] = o.SharesShort
	}
	if !IsNil(o.ShortRatio) {
		toSerialize["short_ratio"] = o.ShortRatio
	}
	if !IsNil(o.ShortPercentOfSharesOutstanding) {
		toSerialize["short_percent_of_shares_outstanding"] = o.ShortPercentOfSharesOutstanding
	}
	if !IsNil(o.PercentHeldByInsiders) {
		toSerialize["percent_held_by_insiders"] = o.PercentHeldByInsiders
	}
	if !IsNil(o.PercentHeldByInstitutions) {
		toSerialize["percent_held_by_institutions"] = o.PercentHeldByInstitutions
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsStockStatistics struct {
	value *GetStatistics200ResponseStatisticsStockStatistics
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) Get() *GetStatistics200ResponseStatisticsStockStatistics {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) Set(val *GetStatistics200ResponseStatisticsStockStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsStockStatistics(val *GetStatistics200ResponseStatisticsStockStatistics) *NullableGetStatistics200ResponseStatisticsStockStatistics {
	return &NullableGetStatistics200ResponseStatisticsStockStatistics{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsStockStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsStockStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


