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

// checks if the GetStatistics200ResponseStatisticsDividendsAndSplits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsDividendsAndSplits{}

// GetStatistics200ResponseStatisticsDividendsAndSplits Dividends and splits information of the company
type GetStatistics200ResponseStatisticsDividendsAndSplits struct {
	// Refers to the forward dividend yield estimation in the currency of instrument
	ForwardAnnualDividendRate *float64 `json:"forward_annual_dividend_rate,omitempty"`
	// Refers to the forward dividend yield percentage relative to stock price
	ForwardAnnualDividendYield *float64 `json:"forward_annual_dividend_yield,omitempty"`
	// Refers to the trailing dividend yield rate in the currency of instrument over the last 12 months
	TrailingAnnualDividendRate *float64 `json:"trailing_annual_dividend_rate,omitempty"`
	// Refers to the trailing dividend yield percentage relative to stock price
	TrailingAnnualDividendYield *float64 `json:"trailing_annual_dividend_yield,omitempty"`
	// Refers to the average 5 years dividend yield
	Var5YearAverageDividendYield *float64 `json:"5_year_average_dividend_yield,omitempty"`
	// Refers to payout ratio, showing the proportion of earnings a company pays its shareholders in the form of dividends
	PayoutRatio *float64 `json:"payout_ratio,omitempty"`
	// Refers to how often a stock or fund pays dividends
	DividendFrequency *string `json:"dividend_frequency,omitempty"`
	// Refers to the last dividend payout date
	DividendDate *string `json:"dividend_date,omitempty"`
	// Refers to the last ex-dividend payout date
	ExDividendDate *string `json:"ex_dividend_date,omitempty"`
	// Specification of the last split event
	LastSplitFactor *string `json:"last_split_factor,omitempty"`
	// Refers for the last split date
	LastSplitDate *string `json:"last_split_date,omitempty"`
}

// NewGetStatistics200ResponseStatisticsDividendsAndSplits instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsDividendsAndSplits() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	this := GetStatistics200ResponseStatisticsDividendsAndSplits{}
	return &this
}

// NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	this := GetStatistics200ResponseStatisticsDividendsAndSplits{}
	return &this
}

// GetForwardAnnualDividendRate returns the ForwardAnnualDividendRate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRate() float64 {
	if o == nil || IsNil(o.ForwardAnnualDividendRate) {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendRate
}

// GetForwardAnnualDividendRateOk returns a tuple with the ForwardAnnualDividendRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRateOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardAnnualDividendRate) {
		return nil, false
	}
	return o.ForwardAnnualDividendRate, true
}

// HasForwardAnnualDividendRate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendRate() bool {
	if o != nil && !IsNil(o.ForwardAnnualDividendRate) {
		return true
	}

	return false
}

// SetForwardAnnualDividendRate gets a reference to the given float64 and assigns it to the ForwardAnnualDividendRate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendRate(v float64) {
	o.ForwardAnnualDividendRate = &v
}

// GetForwardAnnualDividendYield returns the ForwardAnnualDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYield() float64 {
	if o == nil || IsNil(o.ForwardAnnualDividendYield) {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendYield
}

// GetForwardAnnualDividendYieldOk returns a tuple with the ForwardAnnualDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardAnnualDividendYield) {
		return nil, false
	}
	return o.ForwardAnnualDividendYield, true
}

// HasForwardAnnualDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendYield() bool {
	if o != nil && !IsNil(o.ForwardAnnualDividendYield) {
		return true
	}

	return false
}

// SetForwardAnnualDividendYield gets a reference to the given float64 and assigns it to the ForwardAnnualDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendYield(v float64) {
	o.ForwardAnnualDividendYield = &v
}

// GetTrailingAnnualDividendRate returns the TrailingAnnualDividendRate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRate() float64 {
	if o == nil || IsNil(o.TrailingAnnualDividendRate) {
		var ret float64
		return ret
	}
	return *o.TrailingAnnualDividendRate
}

// GetTrailingAnnualDividendRateOk returns a tuple with the TrailingAnnualDividendRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRateOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingAnnualDividendRate) {
		return nil, false
	}
	return o.TrailingAnnualDividendRate, true
}

// HasTrailingAnnualDividendRate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendRate() bool {
	if o != nil && !IsNil(o.TrailingAnnualDividendRate) {
		return true
	}

	return false
}

// SetTrailingAnnualDividendRate gets a reference to the given float64 and assigns it to the TrailingAnnualDividendRate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendRate(v float64) {
	o.TrailingAnnualDividendRate = &v
}

// GetTrailingAnnualDividendYield returns the TrailingAnnualDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYield() float64 {
	if o == nil || IsNil(o.TrailingAnnualDividendYield) {
		var ret float64
		return ret
	}
	return *o.TrailingAnnualDividendYield
}

// GetTrailingAnnualDividendYieldOk returns a tuple with the TrailingAnnualDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingAnnualDividendYield) {
		return nil, false
	}
	return o.TrailingAnnualDividendYield, true
}

// HasTrailingAnnualDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendYield() bool {
	if o != nil && !IsNil(o.TrailingAnnualDividendYield) {
		return true
	}

	return false
}

// SetTrailingAnnualDividendYield gets a reference to the given float64 and assigns it to the TrailingAnnualDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendYield(v float64) {
	o.TrailingAnnualDividendYield = &v
}

// GetVar5YearAverageDividendYield returns the Var5YearAverageDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYield() float64 {
	if o == nil || IsNil(o.Var5YearAverageDividendYield) {
		var ret float64
		return ret
	}
	return *o.Var5YearAverageDividendYield
}

// GetVar5YearAverageDividendYieldOk returns a tuple with the Var5YearAverageDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.Var5YearAverageDividendYield) {
		return nil, false
	}
	return o.Var5YearAverageDividendYield, true
}

// HasVar5YearAverageDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasVar5YearAverageDividendYield() bool {
	if o != nil && !IsNil(o.Var5YearAverageDividendYield) {
		return true
	}

	return false
}

// SetVar5YearAverageDividendYield gets a reference to the given float64 and assigns it to the Var5YearAverageDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetVar5YearAverageDividendYield(v float64) {
	o.Var5YearAverageDividendYield = &v
}

// GetPayoutRatio returns the PayoutRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatio() float64 {
	if o == nil || IsNil(o.PayoutRatio) {
		var ret float64
		return ret
	}
	return *o.PayoutRatio
}

// GetPayoutRatioOk returns a tuple with the PayoutRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.PayoutRatio) {
		return nil, false
	}
	return o.PayoutRatio, true
}

// HasPayoutRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasPayoutRatio() bool {
	if o != nil && !IsNil(o.PayoutRatio) {
		return true
	}

	return false
}

// SetPayoutRatio gets a reference to the given float64 and assigns it to the PayoutRatio field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetPayoutRatio(v float64) {
	o.PayoutRatio = &v
}

// GetDividendFrequency returns the DividendFrequency field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequency() string {
	if o == nil || IsNil(o.DividendFrequency) {
		var ret string
		return ret
	}
	return *o.DividendFrequency
}

// GetDividendFrequencyOk returns a tuple with the DividendFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.DividendFrequency) {
		return nil, false
	}
	return o.DividendFrequency, true
}

// HasDividendFrequency returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendFrequency() bool {
	if o != nil && !IsNil(o.DividendFrequency) {
		return true
	}

	return false
}

// SetDividendFrequency gets a reference to the given string and assigns it to the DividendFrequency field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendFrequency(v string) {
	o.DividendFrequency = &v
}

// GetDividendDate returns the DividendDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDate() string {
	if o == nil || IsNil(o.DividendDate) {
		var ret string
		return ret
	}
	return *o.DividendDate
}

// GetDividendDateOk returns a tuple with the DividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDateOk() (*string, bool) {
	if o == nil || IsNil(o.DividendDate) {
		return nil, false
	}
	return o.DividendDate, true
}

// HasDividendDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendDate() bool {
	if o != nil && !IsNil(o.DividendDate) {
		return true
	}

	return false
}

// SetDividendDate gets a reference to the given string and assigns it to the DividendDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendDate(v string) {
	o.DividendDate = &v
}

// GetExDividendDate returns the ExDividendDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDate() string {
	if o == nil || IsNil(o.ExDividendDate) {
		var ret string
		return ret
	}
	return *o.ExDividendDate
}

// GetExDividendDateOk returns a tuple with the ExDividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExDividendDate) {
		return nil, false
	}
	return o.ExDividendDate, true
}

// HasExDividendDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasExDividendDate() bool {
	if o != nil && !IsNil(o.ExDividendDate) {
		return true
	}

	return false
}

// SetExDividendDate gets a reference to the given string and assigns it to the ExDividendDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetExDividendDate(v string) {
	o.ExDividendDate = &v
}

// GetLastSplitFactor returns the LastSplitFactor field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactor() string {
	if o == nil || IsNil(o.LastSplitFactor) {
		var ret string
		return ret
	}
	return *o.LastSplitFactor
}

// GetLastSplitFactorOk returns a tuple with the LastSplitFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactorOk() (*string, bool) {
	if o == nil || IsNil(o.LastSplitFactor) {
		return nil, false
	}
	return o.LastSplitFactor, true
}

// HasLastSplitFactor returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitFactor() bool {
	if o != nil && !IsNil(o.LastSplitFactor) {
		return true
	}

	return false
}

// SetLastSplitFactor gets a reference to the given string and assigns it to the LastSplitFactor field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitFactor(v string) {
	o.LastSplitFactor = &v
}

// GetLastSplitDate returns the LastSplitDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDate() string {
	if o == nil || IsNil(o.LastSplitDate) {
		var ret string
		return ret
	}
	return *o.LastSplitDate
}

// GetLastSplitDateOk returns a tuple with the LastSplitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSplitDate) {
		return nil, false
	}
	return o.LastSplitDate, true
}

// HasLastSplitDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitDate() bool {
	if o != nil && !IsNil(o.LastSplitDate) {
		return true
	}

	return false
}

// SetLastSplitDate gets a reference to the given string and assigns it to the LastSplitDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitDate(v string) {
	o.LastSplitDate = &v
}

func (o GetStatistics200ResponseStatisticsDividendsAndSplits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsDividendsAndSplits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForwardAnnualDividendRate) {
		toSerialize["forward_annual_dividend_rate"] = o.ForwardAnnualDividendRate
	}
	if !IsNil(o.ForwardAnnualDividendYield) {
		toSerialize["forward_annual_dividend_yield"] = o.ForwardAnnualDividendYield
	}
	if !IsNil(o.TrailingAnnualDividendRate) {
		toSerialize["trailing_annual_dividend_rate"] = o.TrailingAnnualDividendRate
	}
	if !IsNil(o.TrailingAnnualDividendYield) {
		toSerialize["trailing_annual_dividend_yield"] = o.TrailingAnnualDividendYield
	}
	if !IsNil(o.Var5YearAverageDividendYield) {
		toSerialize["5_year_average_dividend_yield"] = o.Var5YearAverageDividendYield
	}
	if !IsNil(o.PayoutRatio) {
		toSerialize["payout_ratio"] = o.PayoutRatio
	}
	if !IsNil(o.DividendFrequency) {
		toSerialize["dividend_frequency"] = o.DividendFrequency
	}
	if !IsNil(o.DividendDate) {
		toSerialize["dividend_date"] = o.DividendDate
	}
	if !IsNil(o.ExDividendDate) {
		toSerialize["ex_dividend_date"] = o.ExDividendDate
	}
	if !IsNil(o.LastSplitFactor) {
		toSerialize["last_split_factor"] = o.LastSplitFactor
	}
	if !IsNil(o.LastSplitDate) {
		toSerialize["last_split_date"] = o.LastSplitDate
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsDividendsAndSplits struct {
	value *GetStatistics200ResponseStatisticsDividendsAndSplits
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Get() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Set(val *GetStatistics200ResponseStatisticsDividendsAndSplits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsDividendsAndSplits(val *GetStatistics200ResponseStatisticsDividendsAndSplits) *NullableGetStatistics200ResponseStatisticsDividendsAndSplits {
	return &NullableGetStatistics200ResponseStatisticsDividendsAndSplits{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


