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

// checks if the GetQuote200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuote200Response{}

// GetQuote200Response struct for GetQuote200Response
type GetQuote200Response struct {
	// Symbol passed
	Symbol *string `json:"symbol,omitempty"`
	// Name of the instrument
	Name *string `json:"name,omitempty"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard. Available for stocks, ETFs, mutual funds, bonds
	MicCode *string `json:"mic_code,omitempty"`
	// Currency in which the equity is denominated. Available for stocks, ETFs, mutual funds, bonds
	Currency *string `json:"currency,omitempty"`
	// Datetime in defined timezone referring to when the bar with specified interval was opened
	Datetime *string `json:"datetime,omitempty"`
	// Unix timestamp representing the opening candle of the specified interval
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Unix timestamp of last minute candle
	LastQuoteAt *int64 `json:"last_quote_at,omitempty"`
	// Price at the opening of current bar
	Open *string `json:"open,omitempty"`
	// Highest price which occurred during the current bar
	High *string `json:"high,omitempty"`
	// Lowest price which occurred during the current bar
	Low *string `json:"low,omitempty"`
	// Close price at the end of the bar
	Close *string `json:"close,omitempty"`
	// Trading volume during the bar. Available not for all instrument types
	Volume *string `json:"volume,omitempty"`
	// Close price at the end of the previous bar
	PreviousClose *string `json:"previous_close,omitempty"`
	// Close - previous_close
	Change *string `json:"change,omitempty"`
	// (Close - previous_close) / previous_close * 100
	PercentChange *string `json:"percent_change,omitempty"`
	// Average volume of the specified period. Available not for all instrument types
	AverageVolume *string `json:"average_volume,omitempty"`
	// Percent change in price between the current and the backward one, where period is 1 day. Available for crypto
	Rolling1dChange *string `json:"rolling_1d_change,omitempty"`
	// Percent change in price between the current and the backward one, where period is 7 days. Available for crypto
	Rolling7dChange *string `json:"rolling_7d_change,omitempty"`
	// Percent change in price between the current and the backward one, where period specified in request param rolling_period. Available for crypto
	RollingChange *string `json:"rolling_change,omitempty"`
	// True if market is open; false if closed
	IsMarketOpen *bool `json:"is_market_open,omitempty"`
	FiftyTwoWeek *GetQuote200ResponseFiftyTwoWeek `json:"fifty_two_week,omitempty"`
	// Diff between the regular close price and the latest extended price. Displayed only if prepost is true
	ExtendedChange *string `json:"extended_change,omitempty"`
	// Percent change in price between the regular close price and the latest extended price. Displayed only if prepost is true
	ExtendedPercentChange *string `json:"extended_percent_change,omitempty"`
	// Latest extended price. Displayed only if prepost is true
	ExtendedPrice *string `json:"extended_price,omitempty"`
	// Unix timestamp of the last extended price. Displayed only if prepost is true
	ExtendedTimestamp *string `json:"extended_timestamp,omitempty"`
}

// NewGetQuote200Response instantiates a new GetQuote200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuote200Response() *GetQuote200Response {
	this := GetQuote200Response{}
	return &this
}

// NewGetQuote200ResponseWithDefaults instantiates a new GetQuote200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuote200ResponseWithDefaults() *GetQuote200Response {
	this := GetQuote200Response{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetQuote200Response) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetQuote200Response) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetQuote200Response) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetQuote200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetQuote200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetQuote200Response) SetName(v string) {
	o.Name = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetQuote200Response) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetQuote200Response) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetQuote200Response) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetQuote200Response) SetMicCode(v string) {
	o.MicCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetQuote200Response) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetQuote200Response) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetQuote200Response) SetCurrency(v string) {
	o.Currency = &v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *GetQuote200Response) GetDatetime() string {
	if o == nil || IsNil(o.Datetime) {
		var ret string
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetDatetimeOk() (*string, bool) {
	if o == nil || IsNil(o.Datetime) {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *GetQuote200Response) HasDatetime() bool {
	if o != nil && !IsNil(o.Datetime) {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given string and assigns it to the Datetime field.
func (o *GetQuote200Response) SetDatetime(v string) {
	o.Datetime = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetQuote200Response) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetQuote200Response) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetQuote200Response) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetLastQuoteAt returns the LastQuoteAt field value if set, zero value otherwise.
func (o *GetQuote200Response) GetLastQuoteAt() int64 {
	if o == nil || IsNil(o.LastQuoteAt) {
		var ret int64
		return ret
	}
	return *o.LastQuoteAt
}

// GetLastQuoteAtOk returns a tuple with the LastQuoteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetLastQuoteAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastQuoteAt) {
		return nil, false
	}
	return o.LastQuoteAt, true
}

// HasLastQuoteAt returns a boolean if a field has been set.
func (o *GetQuote200Response) HasLastQuoteAt() bool {
	if o != nil && !IsNil(o.LastQuoteAt) {
		return true
	}

	return false
}

// SetLastQuoteAt gets a reference to the given int64 and assigns it to the LastQuoteAt field.
func (o *GetQuote200Response) SetLastQuoteAt(v int64) {
	o.LastQuoteAt = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *GetQuote200Response) GetOpen() string {
	if o == nil || IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetOpenOk() (*string, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *GetQuote200Response) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *GetQuote200Response) SetOpen(v string) {
	o.Open = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *GetQuote200Response) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *GetQuote200Response) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *GetQuote200Response) SetHigh(v string) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *GetQuote200Response) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *GetQuote200Response) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *GetQuote200Response) SetLow(v string) {
	o.Low = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *GetQuote200Response) GetClose() string {
	if o == nil || IsNil(o.Close) {
		var ret string
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetCloseOk() (*string, bool) {
	if o == nil || IsNil(o.Close) {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *GetQuote200Response) HasClose() bool {
	if o != nil && !IsNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given string and assigns it to the Close field.
func (o *GetQuote200Response) SetClose(v string) {
	o.Close = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *GetQuote200Response) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *GetQuote200Response) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *GetQuote200Response) SetVolume(v string) {
	o.Volume = &v
}

// GetPreviousClose returns the PreviousClose field value if set, zero value otherwise.
func (o *GetQuote200Response) GetPreviousClose() string {
	if o == nil || IsNil(o.PreviousClose) {
		var ret string
		return ret
	}
	return *o.PreviousClose
}

// GetPreviousCloseOk returns a tuple with the PreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetPreviousCloseOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousClose) {
		return nil, false
	}
	return o.PreviousClose, true
}

// HasPreviousClose returns a boolean if a field has been set.
func (o *GetQuote200Response) HasPreviousClose() bool {
	if o != nil && !IsNil(o.PreviousClose) {
		return true
	}

	return false
}

// SetPreviousClose gets a reference to the given string and assigns it to the PreviousClose field.
func (o *GetQuote200Response) SetPreviousClose(v string) {
	o.PreviousClose = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *GetQuote200Response) GetChange() string {
	if o == nil || IsNil(o.Change) {
		var ret string
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given string and assigns it to the Change field.
func (o *GetQuote200Response) SetChange(v string) {
	o.Change = &v
}

// GetPercentChange returns the PercentChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetPercentChange() string {
	if o == nil || IsNil(o.PercentChange) {
		var ret string
		return ret
	}
	return *o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetPercentChangeOk() (*string, bool) {
	if o == nil || IsNil(o.PercentChange) {
		return nil, false
	}
	return o.PercentChange, true
}

// HasPercentChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasPercentChange() bool {
	if o != nil && !IsNil(o.PercentChange) {
		return true
	}

	return false
}

// SetPercentChange gets a reference to the given string and assigns it to the PercentChange field.
func (o *GetQuote200Response) SetPercentChange(v string) {
	o.PercentChange = &v
}

// GetAverageVolume returns the AverageVolume field value if set, zero value otherwise.
func (o *GetQuote200Response) GetAverageVolume() string {
	if o == nil || IsNil(o.AverageVolume) {
		var ret string
		return ret
	}
	return *o.AverageVolume
}

// GetAverageVolumeOk returns a tuple with the AverageVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetAverageVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.AverageVolume) {
		return nil, false
	}
	return o.AverageVolume, true
}

// HasAverageVolume returns a boolean if a field has been set.
func (o *GetQuote200Response) HasAverageVolume() bool {
	if o != nil && !IsNil(o.AverageVolume) {
		return true
	}

	return false
}

// SetAverageVolume gets a reference to the given string and assigns it to the AverageVolume field.
func (o *GetQuote200Response) SetAverageVolume(v string) {
	o.AverageVolume = &v
}

// GetRolling1dChange returns the Rolling1dChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRolling1dChange() string {
	if o == nil || IsNil(o.Rolling1dChange) {
		var ret string
		return ret
	}
	return *o.Rolling1dChange
}

// GetRolling1dChangeOk returns a tuple with the Rolling1dChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRolling1dChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Rolling1dChange) {
		return nil, false
	}
	return o.Rolling1dChange, true
}

// HasRolling1dChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRolling1dChange() bool {
	if o != nil && !IsNil(o.Rolling1dChange) {
		return true
	}

	return false
}

// SetRolling1dChange gets a reference to the given string and assigns it to the Rolling1dChange field.
func (o *GetQuote200Response) SetRolling1dChange(v string) {
	o.Rolling1dChange = &v
}

// GetRolling7dChange returns the Rolling7dChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRolling7dChange() string {
	if o == nil || IsNil(o.Rolling7dChange) {
		var ret string
		return ret
	}
	return *o.Rolling7dChange
}

// GetRolling7dChangeOk returns a tuple with the Rolling7dChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRolling7dChangeOk() (*string, bool) {
	if o == nil || IsNil(o.Rolling7dChange) {
		return nil, false
	}
	return o.Rolling7dChange, true
}

// HasRolling7dChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRolling7dChange() bool {
	if o != nil && !IsNil(o.Rolling7dChange) {
		return true
	}

	return false
}

// SetRolling7dChange gets a reference to the given string and assigns it to the Rolling7dChange field.
func (o *GetQuote200Response) SetRolling7dChange(v string) {
	o.Rolling7dChange = &v
}

// GetRollingChange returns the RollingChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetRollingChange() string {
	if o == nil || IsNil(o.RollingChange) {
		var ret string
		return ret
	}
	return *o.RollingChange
}

// GetRollingChangeOk returns a tuple with the RollingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetRollingChangeOk() (*string, bool) {
	if o == nil || IsNil(o.RollingChange) {
		return nil, false
	}
	return o.RollingChange, true
}

// HasRollingChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasRollingChange() bool {
	if o != nil && !IsNil(o.RollingChange) {
		return true
	}

	return false
}

// SetRollingChange gets a reference to the given string and assigns it to the RollingChange field.
func (o *GetQuote200Response) SetRollingChange(v string) {
	o.RollingChange = &v
}

// GetIsMarketOpen returns the IsMarketOpen field value if set, zero value otherwise.
func (o *GetQuote200Response) GetIsMarketOpen() bool {
	if o == nil || IsNil(o.IsMarketOpen) {
		var ret bool
		return ret
	}
	return *o.IsMarketOpen
}

// GetIsMarketOpenOk returns a tuple with the IsMarketOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetIsMarketOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMarketOpen) {
		return nil, false
	}
	return o.IsMarketOpen, true
}

// HasIsMarketOpen returns a boolean if a field has been set.
func (o *GetQuote200Response) HasIsMarketOpen() bool {
	if o != nil && !IsNil(o.IsMarketOpen) {
		return true
	}

	return false
}

// SetIsMarketOpen gets a reference to the given bool and assigns it to the IsMarketOpen field.
func (o *GetQuote200Response) SetIsMarketOpen(v bool) {
	o.IsMarketOpen = &v
}

// GetFiftyTwoWeek returns the FiftyTwoWeek field value if set, zero value otherwise.
func (o *GetQuote200Response) GetFiftyTwoWeek() GetQuote200ResponseFiftyTwoWeek {
	if o == nil || IsNil(o.FiftyTwoWeek) {
		var ret GetQuote200ResponseFiftyTwoWeek
		return ret
	}
	return *o.FiftyTwoWeek
}

// GetFiftyTwoWeekOk returns a tuple with the FiftyTwoWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetFiftyTwoWeekOk() (*GetQuote200ResponseFiftyTwoWeek, bool) {
	if o == nil || IsNil(o.FiftyTwoWeek) {
		return nil, false
	}
	return o.FiftyTwoWeek, true
}

// HasFiftyTwoWeek returns a boolean if a field has been set.
func (o *GetQuote200Response) HasFiftyTwoWeek() bool {
	if o != nil && !IsNil(o.FiftyTwoWeek) {
		return true
	}

	return false
}

// SetFiftyTwoWeek gets a reference to the given GetQuote200ResponseFiftyTwoWeek and assigns it to the FiftyTwoWeek field.
func (o *GetQuote200Response) SetFiftyTwoWeek(v GetQuote200ResponseFiftyTwoWeek) {
	o.FiftyTwoWeek = &v
}

// GetExtendedChange returns the ExtendedChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedChange() string {
	if o == nil || IsNil(o.ExtendedChange) {
		var ret string
		return ret
	}
	return *o.ExtendedChange
}

// GetExtendedChangeOk returns a tuple with the ExtendedChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedChange) {
		return nil, false
	}
	return o.ExtendedChange, true
}

// HasExtendedChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedChange() bool {
	if o != nil && !IsNil(o.ExtendedChange) {
		return true
	}

	return false
}

// SetExtendedChange gets a reference to the given string and assigns it to the ExtendedChange field.
func (o *GetQuote200Response) SetExtendedChange(v string) {
	o.ExtendedChange = &v
}

// GetExtendedPercentChange returns the ExtendedPercentChange field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedPercentChange() string {
	if o == nil || IsNil(o.ExtendedPercentChange) {
		var ret string
		return ret
	}
	return *o.ExtendedPercentChange
}

// GetExtendedPercentChangeOk returns a tuple with the ExtendedPercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedPercentChangeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedPercentChange) {
		return nil, false
	}
	return o.ExtendedPercentChange, true
}

// HasExtendedPercentChange returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedPercentChange() bool {
	if o != nil && !IsNil(o.ExtendedPercentChange) {
		return true
	}

	return false
}

// SetExtendedPercentChange gets a reference to the given string and assigns it to the ExtendedPercentChange field.
func (o *GetQuote200Response) SetExtendedPercentChange(v string) {
	o.ExtendedPercentChange = &v
}

// GetExtendedPrice returns the ExtendedPrice field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedPrice() string {
	if o == nil || IsNil(o.ExtendedPrice) {
		var ret string
		return ret
	}
	return *o.ExtendedPrice
}

// GetExtendedPriceOk returns a tuple with the ExtendedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedPrice) {
		return nil, false
	}
	return o.ExtendedPrice, true
}

// HasExtendedPrice returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedPrice() bool {
	if o != nil && !IsNil(o.ExtendedPrice) {
		return true
	}

	return false
}

// SetExtendedPrice gets a reference to the given string and assigns it to the ExtendedPrice field.
func (o *GetQuote200Response) SetExtendedPrice(v string) {
	o.ExtendedPrice = &v
}

// GetExtendedTimestamp returns the ExtendedTimestamp field value if set, zero value otherwise.
func (o *GetQuote200Response) GetExtendedTimestamp() string {
	if o == nil || IsNil(o.ExtendedTimestamp) {
		var ret string
		return ret
	}
	return *o.ExtendedTimestamp
}

// GetExtendedTimestampOk returns a tuple with the ExtendedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuote200Response) GetExtendedTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedTimestamp) {
		return nil, false
	}
	return o.ExtendedTimestamp, true
}

// HasExtendedTimestamp returns a boolean if a field has been set.
func (o *GetQuote200Response) HasExtendedTimestamp() bool {
	if o != nil && !IsNil(o.ExtendedTimestamp) {
		return true
	}

	return false
}

// SetExtendedTimestamp gets a reference to the given string and assigns it to the ExtendedTimestamp field.
func (o *GetQuote200Response) SetExtendedTimestamp(v string) {
	o.ExtendedTimestamp = &v
}

func (o GetQuote200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuote200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Datetime) {
		toSerialize["datetime"] = o.Datetime
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.LastQuoteAt) {
		toSerialize["last_quote_at"] = o.LastQuoteAt
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.PreviousClose) {
		toSerialize["previous_close"] = o.PreviousClose
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.PercentChange) {
		toSerialize["percent_change"] = o.PercentChange
	}
	if !IsNil(o.AverageVolume) {
		toSerialize["average_volume"] = o.AverageVolume
	}
	if !IsNil(o.Rolling1dChange) {
		toSerialize["rolling_1d_change"] = o.Rolling1dChange
	}
	if !IsNil(o.Rolling7dChange) {
		toSerialize["rolling_7d_change"] = o.Rolling7dChange
	}
	if !IsNil(o.RollingChange) {
		toSerialize["rolling_change"] = o.RollingChange
	}
	if !IsNil(o.IsMarketOpen) {
		toSerialize["is_market_open"] = o.IsMarketOpen
	}
	if !IsNil(o.FiftyTwoWeek) {
		toSerialize["fifty_two_week"] = o.FiftyTwoWeek
	}
	if !IsNil(o.ExtendedChange) {
		toSerialize["extended_change"] = o.ExtendedChange
	}
	if !IsNil(o.ExtendedPercentChange) {
		toSerialize["extended_percent_change"] = o.ExtendedPercentChange
	}
	if !IsNil(o.ExtendedPrice) {
		toSerialize["extended_price"] = o.ExtendedPrice
	}
	if !IsNil(o.ExtendedTimestamp) {
		toSerialize["extended_timestamp"] = o.ExtendedTimestamp
	}
	return toSerialize, nil
}

type NullableGetQuote200Response struct {
	value *GetQuote200Response
	isSet bool
}

func (v NullableGetQuote200Response) Get() *GetQuote200Response {
	return v.value
}

func (v *NullableGetQuote200Response) Set(val *GetQuote200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuote200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuote200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuote200Response(val *GetQuote200Response) *NullableGetQuote200Response {
	return &NullableGetQuote200Response{value: val, isSet: true}
}

func (v NullableGetQuote200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuote200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


