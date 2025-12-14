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

// checks if the ResponseMutualFundWorldSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldSummary{}

// ResponseMutualFundWorldSummary A brief summary of a mutual fund
type ResponseMutualFundWorldSummary struct {
	// All available fund types segmented by country
	Symbol *string `json:"symbol,omitempty"`
	// Fund name
	Name *string `json:"name,omitempty"`
	// Investment company that manages the fund
	FundFamily *string `json:"fund_family,omitempty"`
	// Type of the fund
	FundType *string `json:"fund_type,omitempty"`
	// Currency of fund price
	Currency *string `json:"currency,omitempty"`
	// The date from which the fund started operations and the returns are calculated
	ShareClassInceptionDate *string `json:"share_class_inception_date,omitempty"`
	// Percentage of profit of the fund since the first trading day of the current calendar year
	YtdReturn *float64 `json:"ytd_return,omitempty"`
	// Percentage of mutual fund assets steered toward a fund's operating expenses and fund management fees
	ExpenseRatioNet *float64 `json:"expense_ratio_net,omitempty"`
	// Income returned to its investors through interest and dividends generated by the fund's investments
	Yield *float64 `json:"yield,omitempty"`
	// Net Asset Value: fund value minus liabilities
	Nav *float64 `json:"nav,omitempty"`
	// Investment minimum
	MinInvestment *int64 `json:"min_investment,omitempty"`
	// Percentage rate at which mutual fund replaces its holdings on investment every year
	TurnoverRate *float64 `json:"turnover_rate,omitempty"`
	// Total assets of a fund minus its total liabilities
	NetAssets *int64 `json:"net_assets,omitempty"`
	// Description of the fund
	Overview *string `json:"overview,omitempty"`
	// Information about the fund’s managers
	People []GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner `json:"people,omitempty"`
}

// NewResponseMutualFundWorldSummary instantiates a new ResponseMutualFundWorldSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldSummary() *ResponseMutualFundWorldSummary {
	this := ResponseMutualFundWorldSummary{}
	return &this
}

// NewResponseMutualFundWorldSummaryWithDefaults instantiates a new ResponseMutualFundWorldSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldSummaryWithDefaults() *ResponseMutualFundWorldSummary {
	this := ResponseMutualFundWorldSummary{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ResponseMutualFundWorldSummary) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponseMutualFundWorldSummary) SetName(v string) {
	o.Name = &v
}

// GetFundFamily returns the FundFamily field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetFundFamily() string {
	if o == nil || IsNil(o.FundFamily) {
		var ret string
		return ret
	}
	return *o.FundFamily
}

// GetFundFamilyOk returns a tuple with the FundFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetFundFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.FundFamily) {
		return nil, false
	}
	return o.FundFamily, true
}

// HasFundFamily returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasFundFamily() bool {
	if o != nil && !IsNil(o.FundFamily) {
		return true
	}

	return false
}

// SetFundFamily gets a reference to the given string and assigns it to the FundFamily field.
func (o *ResponseMutualFundWorldSummary) SetFundFamily(v string) {
	o.FundFamily = &v
}

// GetFundType returns the FundType field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetFundType() string {
	if o == nil || IsNil(o.FundType) {
		var ret string
		return ret
	}
	return *o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetFundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundType) {
		return nil, false
	}
	return o.FundType, true
}

// HasFundType returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasFundType() bool {
	if o != nil && !IsNil(o.FundType) {
		return true
	}

	return false
}

// SetFundType gets a reference to the given string and assigns it to the FundType field.
func (o *ResponseMutualFundWorldSummary) SetFundType(v string) {
	o.FundType = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ResponseMutualFundWorldSummary) SetCurrency(v string) {
	o.Currency = &v
}

// GetShareClassInceptionDate returns the ShareClassInceptionDate field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetShareClassInceptionDate() string {
	if o == nil || IsNil(o.ShareClassInceptionDate) {
		var ret string
		return ret
	}
	return *o.ShareClassInceptionDate
}

// GetShareClassInceptionDateOk returns a tuple with the ShareClassInceptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetShareClassInceptionDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShareClassInceptionDate) {
		return nil, false
	}
	return o.ShareClassInceptionDate, true
}

// HasShareClassInceptionDate returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasShareClassInceptionDate() bool {
	if o != nil && !IsNil(o.ShareClassInceptionDate) {
		return true
	}

	return false
}

// SetShareClassInceptionDate gets a reference to the given string and assigns it to the ShareClassInceptionDate field.
func (o *ResponseMutualFundWorldSummary) SetShareClassInceptionDate(v string) {
	o.ShareClassInceptionDate = &v
}

// GetYtdReturn returns the YtdReturn field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetYtdReturn() float64 {
	if o == nil || IsNil(o.YtdReturn) {
		var ret float64
		return ret
	}
	return *o.YtdReturn
}

// GetYtdReturnOk returns a tuple with the YtdReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetYtdReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.YtdReturn) {
		return nil, false
	}
	return o.YtdReturn, true
}

// HasYtdReturn returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasYtdReturn() bool {
	if o != nil && !IsNil(o.YtdReturn) {
		return true
	}

	return false
}

// SetYtdReturn gets a reference to the given float64 and assigns it to the YtdReturn field.
func (o *ResponseMutualFundWorldSummary) SetYtdReturn(v float64) {
	o.YtdReturn = &v
}

// GetExpenseRatioNet returns the ExpenseRatioNet field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetExpenseRatioNet() float64 {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		var ret float64
		return ret
	}
	return *o.ExpenseRatioNet
}

// GetExpenseRatioNetOk returns a tuple with the ExpenseRatioNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetExpenseRatioNetOk() (*float64, bool) {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		return nil, false
	}
	return o.ExpenseRatioNet, true
}

// HasExpenseRatioNet returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasExpenseRatioNet() bool {
	if o != nil && !IsNil(o.ExpenseRatioNet) {
		return true
	}

	return false
}

// SetExpenseRatioNet gets a reference to the given float64 and assigns it to the ExpenseRatioNet field.
func (o *ResponseMutualFundWorldSummary) SetExpenseRatioNet(v float64) {
	o.ExpenseRatioNet = &v
}

// GetYield returns the Yield field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetYield() float64 {
	if o == nil || IsNil(o.Yield) {
		var ret float64
		return ret
	}
	return *o.Yield
}

// GetYieldOk returns a tuple with the Yield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.Yield) {
		return nil, false
	}
	return o.Yield, true
}

// HasYield returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasYield() bool {
	if o != nil && !IsNil(o.Yield) {
		return true
	}

	return false
}

// SetYield gets a reference to the given float64 and assigns it to the Yield field.
func (o *ResponseMutualFundWorldSummary) SetYield(v float64) {
	o.Yield = &v
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetNav() float64 {
	if o == nil || IsNil(o.Nav) {
		var ret float64
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetNavOk() (*float64, bool) {
	if o == nil || IsNil(o.Nav) {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasNav() bool {
	if o != nil && !IsNil(o.Nav) {
		return true
	}

	return false
}

// SetNav gets a reference to the given float64 and assigns it to the Nav field.
func (o *ResponseMutualFundWorldSummary) SetNav(v float64) {
	o.Nav = &v
}

// GetMinInvestment returns the MinInvestment field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetMinInvestment() int64 {
	if o == nil || IsNil(o.MinInvestment) {
		var ret int64
		return ret
	}
	return *o.MinInvestment
}

// GetMinInvestmentOk returns a tuple with the MinInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetMinInvestmentOk() (*int64, bool) {
	if o == nil || IsNil(o.MinInvestment) {
		return nil, false
	}
	return o.MinInvestment, true
}

// HasMinInvestment returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasMinInvestment() bool {
	if o != nil && !IsNil(o.MinInvestment) {
		return true
	}

	return false
}

// SetMinInvestment gets a reference to the given int64 and assigns it to the MinInvestment field.
func (o *ResponseMutualFundWorldSummary) SetMinInvestment(v int64) {
	o.MinInvestment = &v
}

// GetTurnoverRate returns the TurnoverRate field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetTurnoverRate() float64 {
	if o == nil || IsNil(o.TurnoverRate) {
		var ret float64
		return ret
	}
	return *o.TurnoverRate
}

// GetTurnoverRateOk returns a tuple with the TurnoverRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetTurnoverRateOk() (*float64, bool) {
	if o == nil || IsNil(o.TurnoverRate) {
		return nil, false
	}
	return o.TurnoverRate, true
}

// HasTurnoverRate returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasTurnoverRate() bool {
	if o != nil && !IsNil(o.TurnoverRate) {
		return true
	}

	return false
}

// SetTurnoverRate gets a reference to the given float64 and assigns it to the TurnoverRate field.
func (o *ResponseMutualFundWorldSummary) SetTurnoverRate(v float64) {
	o.TurnoverRate = &v
}

// GetNetAssets returns the NetAssets field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetNetAssets() int64 {
	if o == nil || IsNil(o.NetAssets) {
		var ret int64
		return ret
	}
	return *o.NetAssets
}

// GetNetAssetsOk returns a tuple with the NetAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetNetAssetsOk() (*int64, bool) {
	if o == nil || IsNil(o.NetAssets) {
		return nil, false
	}
	return o.NetAssets, true
}

// HasNetAssets returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasNetAssets() bool {
	if o != nil && !IsNil(o.NetAssets) {
		return true
	}

	return false
}

// SetNetAssets gets a reference to the given int64 and assigns it to the NetAssets field.
func (o *ResponseMutualFundWorldSummary) SetNetAssets(v int64) {
	o.NetAssets = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *ResponseMutualFundWorldSummary) SetOverview(v string) {
	o.Overview = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldSummary) GetPeople() []GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner {
	if o == nil || IsNil(o.People) {
		var ret []GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner
		return ret
	}
	return o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldSummary) GetPeopleOk() ([]GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldSummary) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner and assigns it to the People field.
func (o *ResponseMutualFundWorldSummary) SetPeople(v []GetMutualFundsWorld200ResponseMutualFundSummaryPeopleInner) {
	o.People = v
}

func (o ResponseMutualFundWorldSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FundFamily) {
		toSerialize["fund_family"] = o.FundFamily
	}
	if !IsNil(o.FundType) {
		toSerialize["fund_type"] = o.FundType
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ShareClassInceptionDate) {
		toSerialize["share_class_inception_date"] = o.ShareClassInceptionDate
	}
	if !IsNil(o.YtdReturn) {
		toSerialize["ytd_return"] = o.YtdReturn
	}
	if !IsNil(o.ExpenseRatioNet) {
		toSerialize["expense_ratio_net"] = o.ExpenseRatioNet
	}
	if !IsNil(o.Yield) {
		toSerialize["yield"] = o.Yield
	}
	if !IsNil(o.Nav) {
		toSerialize["nav"] = o.Nav
	}
	if !IsNil(o.MinInvestment) {
		toSerialize["min_investment"] = o.MinInvestment
	}
	if !IsNil(o.TurnoverRate) {
		toSerialize["turnover_rate"] = o.TurnoverRate
	}
	if !IsNil(o.NetAssets) {
		toSerialize["net_assets"] = o.NetAssets
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldSummary struct {
	value *ResponseMutualFundWorldSummary
	isSet bool
}

func (v NullableResponseMutualFundWorldSummary) Get() *ResponseMutualFundWorldSummary {
	return v.value
}

func (v *NullableResponseMutualFundWorldSummary) Set(val *ResponseMutualFundWorldSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldSummary(val *ResponseMutualFundWorldSummary) *NullableResponseMutualFundWorldSummary {
	return &NullableResponseMutualFundWorldSummary{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


