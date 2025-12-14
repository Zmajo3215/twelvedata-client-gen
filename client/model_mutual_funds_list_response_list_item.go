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

// checks if the MutualFundsListResponseListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutualFundsListResponseListItem{}

// MutualFundsListResponseListItem struct for MutualFundsListResponseListItem
type MutualFundsListResponseListItem struct {
	// Fund symbol ticker
	Symbol *string `json:"symbol,omitempty"`
	// Fund name
	Name *string `json:"name,omitempty"`
	// Country of fund incorporation
	Country *string `json:"country,omitempty"`
	// Investment company that manages the fund
	FundFamily *string `json:"fund_family,omitempty"`
	// Type of fund
	FundType *string `json:"fund_type,omitempty"`
	// Performance rating from `0` to `5`
	PerformanceRating *int64 `json:"performance_rating,omitempty"`
	// Risk rating from `0` to `5`
	RiskRating *int64 `json:"risk_rating,omitempty"`
	// Currency code in which the fund is denominated
	Currency *string `json:"currency,omitempty"`
	// Exchange name where the fund is listed
	Exchange *string `json:"exchange,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
}

// NewMutualFundsListResponseListItem instantiates a new MutualFundsListResponseListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFundsListResponseListItem() *MutualFundsListResponseListItem {
	this := MutualFundsListResponseListItem{}
	return &this
}

// NewMutualFundsListResponseListItemWithDefaults instantiates a new MutualFundsListResponseListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFundsListResponseListItemWithDefaults() *MutualFundsListResponseListItem {
	this := MutualFundsListResponseListItem{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MutualFundsListResponseListItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MutualFundsListResponseListItem) SetName(v string) {
	o.Name = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *MutualFundsListResponseListItem) SetCountry(v string) {
	o.Country = &v
}

// GetFundFamily returns the FundFamily field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetFundFamily() string {
	if o == nil || IsNil(o.FundFamily) {
		var ret string
		return ret
	}
	return *o.FundFamily
}

// GetFundFamilyOk returns a tuple with the FundFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetFundFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.FundFamily) {
		return nil, false
	}
	return o.FundFamily, true
}

// HasFundFamily returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasFundFamily() bool {
	if o != nil && !IsNil(o.FundFamily) {
		return true
	}

	return false
}

// SetFundFamily gets a reference to the given string and assigns it to the FundFamily field.
func (o *MutualFundsListResponseListItem) SetFundFamily(v string) {
	o.FundFamily = &v
}

// GetFundType returns the FundType field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetFundType() string {
	if o == nil || IsNil(o.FundType) {
		var ret string
		return ret
	}
	return *o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetFundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundType) {
		return nil, false
	}
	return o.FundType, true
}

// HasFundType returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasFundType() bool {
	if o != nil && !IsNil(o.FundType) {
		return true
	}

	return false
}

// SetFundType gets a reference to the given string and assigns it to the FundType field.
func (o *MutualFundsListResponseListItem) SetFundType(v string) {
	o.FundType = &v
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetPerformanceRating() int64 {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret int64
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetPerformanceRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given int64 and assigns it to the PerformanceRating field.
func (o *MutualFundsListResponseListItem) SetPerformanceRating(v int64) {
	o.PerformanceRating = &v
}

// GetRiskRating returns the RiskRating field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetRiskRating() int64 {
	if o == nil || IsNil(o.RiskRating) {
		var ret int64
		return ret
	}
	return *o.RiskRating
}

// GetRiskRatingOk returns a tuple with the RiskRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetRiskRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.RiskRating) {
		return nil, false
	}
	return o.RiskRating, true
}

// HasRiskRating returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasRiskRating() bool {
	if o != nil && !IsNil(o.RiskRating) {
		return true
	}

	return false
}

// SetRiskRating gets a reference to the given int64 and assigns it to the RiskRating field.
func (o *MutualFundsListResponseListItem) SetRiskRating(v int64) {
	o.RiskRating = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MutualFundsListResponseListItem) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *MutualFundsListResponseListItem) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *MutualFundsListResponseListItem) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundsListResponseListItem) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *MutualFundsListResponseListItem) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *MutualFundsListResponseListItem) SetMicCode(v string) {
	o.MicCode = &v
}

func (o MutualFundsListResponseListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutualFundsListResponseListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.FundFamily) {
		toSerialize["fund_family"] = o.FundFamily
	}
	if !IsNil(o.FundType) {
		toSerialize["fund_type"] = o.FundType
	}
	if !IsNil(o.PerformanceRating) {
		toSerialize["performance_rating"] = o.PerformanceRating
	}
	if !IsNil(o.RiskRating) {
		toSerialize["risk_rating"] = o.RiskRating
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	return toSerialize, nil
}

type NullableMutualFundsListResponseListItem struct {
	value *MutualFundsListResponseListItem
	isSet bool
}

func (v NullableMutualFundsListResponseListItem) Get() *MutualFundsListResponseListItem {
	return v.value
}

func (v *NullableMutualFundsListResponseListItem) Set(val *MutualFundsListResponseListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundsListResponseListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundsListResponseListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundsListResponseListItem(val *MutualFundsListResponseListItem) *NullableMutualFundsListResponseListItem {
	return &NullableMutualFundsListResponseListItem{value: val, isSet: true}
}

func (v NullableMutualFundsListResponseListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundsListResponseListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


