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

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}

// GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets Current assets section
type GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets struct {
	// Cash includes currency, bank accounts, and undeposited checks
	Cash *float64 `json:"cash,omitempty"`
	// Represents cash equivalents that have high credit quality and are highly liquid
	CashEquivalents *float64 `json:"cash_equivalents,omitempty"`
	// Combined value of cash and cash equivalents when company doesn't report a breakdown
	CashAndCashEquivalents *float64 `json:"cash_and_cash_equivalents,omitempty"`
	// Represents other short term investments
	OtherShortTermInvestments *float64 `json:"other_short_term_investments,omitempty"`
	// Represents the balance of money due to a firm for goods or services delivered or used but not yet paid for by customers
	AccountsReceivable *float64 `json:"accounts_receivable,omitempty"`
	// Represents other receivables
	OtherReceivables *float64 `json:"other_receivables,omitempty"`
	// Represents the goods available for sale and raw materials used to produce goods available for sale
	Inventory *float64 `json:"inventory,omitempty"`
	// Represents expense that has already been paid for, but which has not yet been consumed
	PrepaidAssets *float64 `json:"prepaid_assets,omitempty"`
	// Represents money that is held for a specific purpose and thus not available to the company for immediate or general business use
	RestrictedCash *float64 `json:"restricted_cash,omitempty"`
	// Represents assets which company plans to sell
	AssetsHeldForSale *float64 `json:"assets_held_for_sale,omitempty"`
	// Represents money that is spent on hedging assets
	HedgingAssets *float64 `json:"hedging_assets,omitempty"`
	// Represents other current assets
	OtherCurrentAssets *float64 `json:"other_current_assets,omitempty"`
	// All current assets values in a total
	TotalCurrentAssets *float64 `json:"total_current_assets,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}
	return &this
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCash() float64 {
	if o == nil || IsNil(o.Cash) {
		var ret float64
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashOk() (*float64, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given float64 and assigns it to the Cash field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCash(v float64) {
	o.Cash = &v
}

// GetCashEquivalents returns the CashEquivalents field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashEquivalents() float64 {
	if o == nil || IsNil(o.CashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashEquivalents
}

// GetCashEquivalentsOk returns a tuple with the CashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashEquivalents) {
		return nil, false
	}
	return o.CashEquivalents, true
}

// HasCashEquivalents returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCashEquivalents() bool {
	if o != nil && !IsNil(o.CashEquivalents) {
		return true
	}

	return false
}

// SetCashEquivalents gets a reference to the given float64 and assigns it to the CashEquivalents field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCashEquivalents(v float64) {
	o.CashEquivalents = &v
}

// GetCashAndCashEquivalents returns the CashAndCashEquivalents field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashAndCashEquivalents() float64 {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashAndCashEquivalents
}

// GetCashAndCashEquivalentsOk returns a tuple with the CashAndCashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashAndCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		return nil, false
	}
	return o.CashAndCashEquivalents, true
}

// HasCashAndCashEquivalents returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCashAndCashEquivalents() bool {
	if o != nil && !IsNil(o.CashAndCashEquivalents) {
		return true
	}

	return false
}

// SetCashAndCashEquivalents gets a reference to the given float64 and assigns it to the CashAndCashEquivalents field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCashAndCashEquivalents(v float64) {
	o.CashAndCashEquivalents = &v
}

// GetOtherShortTermInvestments returns the OtherShortTermInvestments field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherShortTermInvestments() float64 {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		var ret float64
		return ret
	}
	return *o.OtherShortTermInvestments
}

// GetOtherShortTermInvestmentsOk returns a tuple with the OtherShortTermInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherShortTermInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		return nil, false
	}
	return o.OtherShortTermInvestments, true
}

// HasOtherShortTermInvestments returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherShortTermInvestments() bool {
	if o != nil && !IsNil(o.OtherShortTermInvestments) {
		return true
	}

	return false
}

// SetOtherShortTermInvestments gets a reference to the given float64 and assigns it to the OtherShortTermInvestments field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherShortTermInvestments(v float64) {
	o.OtherShortTermInvestments = &v
}

// GetAccountsReceivable returns the AccountsReceivable field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAccountsReceivable() float64 {
	if o == nil || IsNil(o.AccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AccountsReceivable
}

// GetAccountsReceivableOk returns a tuple with the AccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsReceivable) {
		return nil, false
	}
	return o.AccountsReceivable, true
}

// HasAccountsReceivable returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasAccountsReceivable() bool {
	if o != nil && !IsNil(o.AccountsReceivable) {
		return true
	}

	return false
}

// SetAccountsReceivable gets a reference to the given float64 and assigns it to the AccountsReceivable field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetAccountsReceivable(v float64) {
	o.AccountsReceivable = &v
}

// GetOtherReceivables returns the OtherReceivables field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherReceivables() float64 {
	if o == nil || IsNil(o.OtherReceivables) {
		var ret float64
		return ret
	}
	return *o.OtherReceivables
}

// GetOtherReceivablesOk returns a tuple with the OtherReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherReceivables) {
		return nil, false
	}
	return o.OtherReceivables, true
}

// HasOtherReceivables returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherReceivables() bool {
	if o != nil && !IsNil(o.OtherReceivables) {
		return true
	}

	return false
}

// SetOtherReceivables gets a reference to the given float64 and assigns it to the OtherReceivables field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherReceivables(v float64) {
	o.OtherReceivables = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetInventory() float64 {
	if o == nil || IsNil(o.Inventory) {
		var ret float64
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetInventoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given float64 and assigns it to the Inventory field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetInventory(v float64) {
	o.Inventory = &v
}

// GetPrepaidAssets returns the PrepaidAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetPrepaidAssets() float64 {
	if o == nil || IsNil(o.PrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.PrepaidAssets
}

// GetPrepaidAssetsOk returns a tuple with the PrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.PrepaidAssets) {
		return nil, false
	}
	return o.PrepaidAssets, true
}

// HasPrepaidAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasPrepaidAssets() bool {
	if o != nil && !IsNil(o.PrepaidAssets) {
		return true
	}

	return false
}

// SetPrepaidAssets gets a reference to the given float64 and assigns it to the PrepaidAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetPrepaidAssets(v float64) {
	o.PrepaidAssets = &v
}

// GetRestrictedCash returns the RestrictedCash field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetRestrictedCash() float64 {
	if o == nil || IsNil(o.RestrictedCash) {
		var ret float64
		return ret
	}
	return *o.RestrictedCash
}

// GetRestrictedCashOk returns a tuple with the RestrictedCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetRestrictedCashOk() (*float64, bool) {
	if o == nil || IsNil(o.RestrictedCash) {
		return nil, false
	}
	return o.RestrictedCash, true
}

// HasRestrictedCash returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasRestrictedCash() bool {
	if o != nil && !IsNil(o.RestrictedCash) {
		return true
	}

	return false
}

// SetRestrictedCash gets a reference to the given float64 and assigns it to the RestrictedCash field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetRestrictedCash(v float64) {
	o.RestrictedCash = &v
}

// GetAssetsHeldForSale returns the AssetsHeldForSale field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAssetsHeldForSale() float64 {
	if o == nil || IsNil(o.AssetsHeldForSale) {
		var ret float64
		return ret
	}
	return *o.AssetsHeldForSale
}

// GetAssetsHeldForSaleOk returns a tuple with the AssetsHeldForSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAssetsHeldForSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.AssetsHeldForSale) {
		return nil, false
	}
	return o.AssetsHeldForSale, true
}

// HasAssetsHeldForSale returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasAssetsHeldForSale() bool {
	if o != nil && !IsNil(o.AssetsHeldForSale) {
		return true
	}

	return false
}

// SetAssetsHeldForSale gets a reference to the given float64 and assigns it to the AssetsHeldForSale field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetAssetsHeldForSale(v float64) {
	o.AssetsHeldForSale = &v
}

// GetHedgingAssets returns the HedgingAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetHedgingAssets() float64 {
	if o == nil || IsNil(o.HedgingAssets) {
		var ret float64
		return ret
	}
	return *o.HedgingAssets
}

// GetHedgingAssetsOk returns a tuple with the HedgingAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetHedgingAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.HedgingAssets) {
		return nil, false
	}
	return o.HedgingAssets, true
}

// HasHedgingAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasHedgingAssets() bool {
	if o != nil && !IsNil(o.HedgingAssets) {
		return true
	}

	return false
}

// SetHedgingAssets gets a reference to the given float64 and assigns it to the HedgingAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetHedgingAssets(v float64) {
	o.HedgingAssets = &v
}

// GetOtherCurrentAssets returns the OtherCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentAssets
}

// GetOtherCurrentAssetsOk returns a tuple with the OtherCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		return nil, false
	}
	return o.OtherCurrentAssets, true
}

// HasOtherCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherCurrentAssets) {
		return true
	}

	return false
}

// SetOtherCurrentAssets gets a reference to the given float64 and assigns it to the OtherCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherCurrentAssets(v float64) {
	o.OtherCurrentAssets = &v
}

// GetTotalCurrentAssets returns the TotalCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetTotalCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentAssets
}

// GetTotalCurrentAssetsOk returns a tuple with the TotalCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetTotalCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		return nil, false
	}
	return o.TotalCurrentAssets, true
}

// HasTotalCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasTotalCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalCurrentAssets) {
		return true
	}

	return false
}

// SetTotalCurrentAssets gets a reference to the given float64 and assigns it to the TotalCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetTotalCurrentAssets(v float64) {
	o.TotalCurrentAssets = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cash) {
		toSerialize["cash"] = o.Cash
	}
	if !IsNil(o.CashEquivalents) {
		toSerialize["cash_equivalents"] = o.CashEquivalents
	}
	if !IsNil(o.CashAndCashEquivalents) {
		toSerialize["cash_and_cash_equivalents"] = o.CashAndCashEquivalents
	}
	if !IsNil(o.OtherShortTermInvestments) {
		toSerialize["other_short_term_investments"] = o.OtherShortTermInvestments
	}
	if !IsNil(o.AccountsReceivable) {
		toSerialize["accounts_receivable"] = o.AccountsReceivable
	}
	if !IsNil(o.OtherReceivables) {
		toSerialize["other_receivables"] = o.OtherReceivables
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.PrepaidAssets) {
		toSerialize["prepaid_assets"] = o.PrepaidAssets
	}
	if !IsNil(o.RestrictedCash) {
		toSerialize["restricted_cash"] = o.RestrictedCash
	}
	if !IsNil(o.AssetsHeldForSale) {
		toSerialize["assets_held_for_sale"] = o.AssetsHeldForSale
	}
	if !IsNil(o.HedgingAssets) {
		toSerialize["hedging_assets"] = o.HedgingAssets
	}
	if !IsNil(o.OtherCurrentAssets) {
		toSerialize["other_current_assets"] = o.OtherCurrentAssets
	}
	if !IsNil(o.TotalCurrentAssets) {
		toSerialize["total_current_assets"] = o.TotalCurrentAssets
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Get() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


