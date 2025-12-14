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

// checks if the IncomeStatementItemOtherIncomeAndExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemOtherIncomeAndExpenses{}

// IncomeStatementItemOtherIncomeAndExpenses Other income and expenses information
type IncomeStatementItemOtherIncomeAndExpenses struct {
	// Other income expense
	OtherIncomeExpense *float64 `json:"other_income_expense,omitempty"`
	// Other non operating income expenses
	OtherNonOperatingIncomeExpenses *float64 `json:"other_non_operating_income_expenses,omitempty"`
	// Special income charges
	SpecialIncomeCharges *float64 `json:"special_income_charges,omitempty"`
	// Gain on sale of PPE
	GainOnSaleOfPpe *float64 `json:"gain_on_sale_of_ppe,omitempty"`
	// Gain on sale of business
	GainOnSaleOfBusiness *float64 `json:"gain_on_sale_of_business,omitempty"`
	// Gain on sale of security
	GainOnSaleOfSecurity *float64 `json:"gain_on_sale_of_security,omitempty"`
	// Other special charges
	OtherSpecialCharges *float64 `json:"other_special_charges,omitempty"`
	// Write off
	WriteOff *float64 `json:"write_off,omitempty"`
	// Impairment of capital assets
	ImpairmentOfCapitalAssets *float64 `json:"impairment_of_capital_assets,omitempty"`
	// Restructuring and merger acquisition
	RestructuringAndMergerAcquisition *float64 `json:"restructuring_and_merger_acquisition,omitempty"`
	// Securities amortization
	SecuritiesAmortization *float64 `json:"securities_amortization,omitempty"`
	// Earnings from equity interest
	EarningsFromEquityInterest *float64 `json:"earnings_from_equity_interest,omitempty"`
	// Earnings from equity interest net of tax
	EarningsFromEquityInterestNetOfTax *float64 `json:"earnings_from_equity_interest_net_of_tax,omitempty"`
	// Total other finance cost
	TotalOtherFinanceCost *float64 `json:"total_other_finance_cost,omitempty"`
}

// NewIncomeStatementItemOtherIncomeAndExpenses instantiates a new IncomeStatementItemOtherIncomeAndExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemOtherIncomeAndExpenses() *IncomeStatementItemOtherIncomeAndExpenses {
	this := IncomeStatementItemOtherIncomeAndExpenses{}
	return &this
}

// NewIncomeStatementItemOtherIncomeAndExpensesWithDefaults instantiates a new IncomeStatementItemOtherIncomeAndExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemOtherIncomeAndExpensesWithDefaults() *IncomeStatementItemOtherIncomeAndExpenses {
	this := IncomeStatementItemOtherIncomeAndExpenses{}
	return &this
}

// GetOtherIncomeExpense returns the OtherIncomeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherIncomeExpense() float64 {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		var ret float64
		return ret
	}
	return *o.OtherIncomeExpense
}

// GetOtherIncomeExpenseOk returns a tuple with the OtherIncomeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherIncomeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		return nil, false
	}
	return o.OtherIncomeExpense, true
}

// HasOtherIncomeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherIncomeExpense() bool {
	if o != nil && !IsNil(o.OtherIncomeExpense) {
		return true
	}

	return false
}

// SetOtherIncomeExpense gets a reference to the given float64 and assigns it to the OtherIncomeExpense field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherIncomeExpense(v float64) {
	o.OtherIncomeExpense = &v
}

// GetOtherNonOperatingIncomeExpenses returns the OtherNonOperatingIncomeExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherNonOperatingIncomeExpenses() float64 {
	if o == nil || IsNil(o.OtherNonOperatingIncomeExpenses) {
		var ret float64
		return ret
	}
	return *o.OtherNonOperatingIncomeExpenses
}

// GetOtherNonOperatingIncomeExpensesOk returns a tuple with the OtherNonOperatingIncomeExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherNonOperatingIncomeExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonOperatingIncomeExpenses) {
		return nil, false
	}
	return o.OtherNonOperatingIncomeExpenses, true
}

// HasOtherNonOperatingIncomeExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherNonOperatingIncomeExpenses() bool {
	if o != nil && !IsNil(o.OtherNonOperatingIncomeExpenses) {
		return true
	}

	return false
}

// SetOtherNonOperatingIncomeExpenses gets a reference to the given float64 and assigns it to the OtherNonOperatingIncomeExpenses field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherNonOperatingIncomeExpenses(v float64) {
	o.OtherNonOperatingIncomeExpenses = &v
}

// GetSpecialIncomeCharges returns the SpecialIncomeCharges field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSpecialIncomeCharges() float64 {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		var ret float64
		return ret
	}
	return *o.SpecialIncomeCharges
}

// GetSpecialIncomeChargesOk returns a tuple with the SpecialIncomeCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSpecialIncomeChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		return nil, false
	}
	return o.SpecialIncomeCharges, true
}

// HasSpecialIncomeCharges returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasSpecialIncomeCharges() bool {
	if o != nil && !IsNil(o.SpecialIncomeCharges) {
		return true
	}

	return false
}

// SetSpecialIncomeCharges gets a reference to the given float64 and assigns it to the SpecialIncomeCharges field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetSpecialIncomeCharges(v float64) {
	o.SpecialIncomeCharges = &v
}

// GetGainOnSaleOfPpe returns the GainOnSaleOfPpe field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfPpe() float64 {
	if o == nil || IsNil(o.GainOnSaleOfPpe) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfPpe
}

// GetGainOnSaleOfPpeOk returns a tuple with the GainOnSaleOfPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfPpe) {
		return nil, false
	}
	return o.GainOnSaleOfPpe, true
}

// HasGainOnSaleOfPpe returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfPpe() bool {
	if o != nil && !IsNil(o.GainOnSaleOfPpe) {
		return true
	}

	return false
}

// SetGainOnSaleOfPpe gets a reference to the given float64 and assigns it to the GainOnSaleOfPpe field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfPpe(v float64) {
	o.GainOnSaleOfPpe = &v
}

// GetGainOnSaleOfBusiness returns the GainOnSaleOfBusiness field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfBusiness() float64 {
	if o == nil || IsNil(o.GainOnSaleOfBusiness) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfBusiness
}

// GetGainOnSaleOfBusinessOk returns a tuple with the GainOnSaleOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfBusinessOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfBusiness) {
		return nil, false
	}
	return o.GainOnSaleOfBusiness, true
}

// HasGainOnSaleOfBusiness returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfBusiness() bool {
	if o != nil && !IsNil(o.GainOnSaleOfBusiness) {
		return true
	}

	return false
}

// SetGainOnSaleOfBusiness gets a reference to the given float64 and assigns it to the GainOnSaleOfBusiness field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfBusiness(v float64) {
	o.GainOnSaleOfBusiness = &v
}

// GetGainOnSaleOfSecurity returns the GainOnSaleOfSecurity field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfSecurity() float64 {
	if o == nil || IsNil(o.GainOnSaleOfSecurity) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfSecurity
}

// GetGainOnSaleOfSecurityOk returns a tuple with the GainOnSaleOfSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfSecurityOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfSecurity) {
		return nil, false
	}
	return o.GainOnSaleOfSecurity, true
}

// HasGainOnSaleOfSecurity returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfSecurity() bool {
	if o != nil && !IsNil(o.GainOnSaleOfSecurity) {
		return true
	}

	return false
}

// SetGainOnSaleOfSecurity gets a reference to the given float64 and assigns it to the GainOnSaleOfSecurity field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfSecurity(v float64) {
	o.GainOnSaleOfSecurity = &v
}

// GetOtherSpecialCharges returns the OtherSpecialCharges field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherSpecialCharges() float64 {
	if o == nil || IsNil(o.OtherSpecialCharges) {
		var ret float64
		return ret
	}
	return *o.OtherSpecialCharges
}

// GetOtherSpecialChargesOk returns a tuple with the OtherSpecialCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherSpecialChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherSpecialCharges) {
		return nil, false
	}
	return o.OtherSpecialCharges, true
}

// HasOtherSpecialCharges returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherSpecialCharges() bool {
	if o != nil && !IsNil(o.OtherSpecialCharges) {
		return true
	}

	return false
}

// SetOtherSpecialCharges gets a reference to the given float64 and assigns it to the OtherSpecialCharges field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherSpecialCharges(v float64) {
	o.OtherSpecialCharges = &v
}

// GetWriteOff returns the WriteOff field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetWriteOff() float64 {
	if o == nil || IsNil(o.WriteOff) {
		var ret float64
		return ret
	}
	return *o.WriteOff
}

// GetWriteOffOk returns a tuple with the WriteOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetWriteOffOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteOff) {
		return nil, false
	}
	return o.WriteOff, true
}

// HasWriteOff returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasWriteOff() bool {
	if o != nil && !IsNil(o.WriteOff) {
		return true
	}

	return false
}

// SetWriteOff gets a reference to the given float64 and assigns it to the WriteOff field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetWriteOff(v float64) {
	o.WriteOff = &v
}

// GetImpairmentOfCapitalAssets returns the ImpairmentOfCapitalAssets field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetImpairmentOfCapitalAssets() float64 {
	if o == nil || IsNil(o.ImpairmentOfCapitalAssets) {
		var ret float64
		return ret
	}
	return *o.ImpairmentOfCapitalAssets
}

// GetImpairmentOfCapitalAssetsOk returns a tuple with the ImpairmentOfCapitalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetImpairmentOfCapitalAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.ImpairmentOfCapitalAssets) {
		return nil, false
	}
	return o.ImpairmentOfCapitalAssets, true
}

// HasImpairmentOfCapitalAssets returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasImpairmentOfCapitalAssets() bool {
	if o != nil && !IsNil(o.ImpairmentOfCapitalAssets) {
		return true
	}

	return false
}

// SetImpairmentOfCapitalAssets gets a reference to the given float64 and assigns it to the ImpairmentOfCapitalAssets field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetImpairmentOfCapitalAssets(v float64) {
	o.ImpairmentOfCapitalAssets = &v
}

// GetRestructuringAndMergerAcquisition returns the RestructuringAndMergerAcquisition field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetRestructuringAndMergerAcquisition() float64 {
	if o == nil || IsNil(o.RestructuringAndMergerAcquisition) {
		var ret float64
		return ret
	}
	return *o.RestructuringAndMergerAcquisition
}

// GetRestructuringAndMergerAcquisitionOk returns a tuple with the RestructuringAndMergerAcquisition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetRestructuringAndMergerAcquisitionOk() (*float64, bool) {
	if o == nil || IsNil(o.RestructuringAndMergerAcquisition) {
		return nil, false
	}
	return o.RestructuringAndMergerAcquisition, true
}

// HasRestructuringAndMergerAcquisition returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasRestructuringAndMergerAcquisition() bool {
	if o != nil && !IsNil(o.RestructuringAndMergerAcquisition) {
		return true
	}

	return false
}

// SetRestructuringAndMergerAcquisition gets a reference to the given float64 and assigns it to the RestructuringAndMergerAcquisition field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetRestructuringAndMergerAcquisition(v float64) {
	o.RestructuringAndMergerAcquisition = &v
}

// GetSecuritiesAmortization returns the SecuritiesAmortization field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSecuritiesAmortization() float64 {
	if o == nil || IsNil(o.SecuritiesAmortization) {
		var ret float64
		return ret
	}
	return *o.SecuritiesAmortization
}

// GetSecuritiesAmortizationOk returns a tuple with the SecuritiesAmortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSecuritiesAmortizationOk() (*float64, bool) {
	if o == nil || IsNil(o.SecuritiesAmortization) {
		return nil, false
	}
	return o.SecuritiesAmortization, true
}

// HasSecuritiesAmortization returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasSecuritiesAmortization() bool {
	if o != nil && !IsNil(o.SecuritiesAmortization) {
		return true
	}

	return false
}

// SetSecuritiesAmortization gets a reference to the given float64 and assigns it to the SecuritiesAmortization field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetSecuritiesAmortization(v float64) {
	o.SecuritiesAmortization = &v
}

// GetEarningsFromEquityInterest returns the EarningsFromEquityInterest field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterest() float64 {
	if o == nil || IsNil(o.EarningsFromEquityInterest) {
		var ret float64
		return ret
	}
	return *o.EarningsFromEquityInterest
}

// GetEarningsFromEquityInterestOk returns a tuple with the EarningsFromEquityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.EarningsFromEquityInterest) {
		return nil, false
	}
	return o.EarningsFromEquityInterest, true
}

// HasEarningsFromEquityInterest returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasEarningsFromEquityInterest() bool {
	if o != nil && !IsNil(o.EarningsFromEquityInterest) {
		return true
	}

	return false
}

// SetEarningsFromEquityInterest gets a reference to the given float64 and assigns it to the EarningsFromEquityInterest field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetEarningsFromEquityInterest(v float64) {
	o.EarningsFromEquityInterest = &v
}

// GetEarningsFromEquityInterestNetOfTax returns the EarningsFromEquityInterestNetOfTax field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestNetOfTax() float64 {
	if o == nil || IsNil(o.EarningsFromEquityInterestNetOfTax) {
		var ret float64
		return ret
	}
	return *o.EarningsFromEquityInterestNetOfTax
}

// GetEarningsFromEquityInterestNetOfTaxOk returns a tuple with the EarningsFromEquityInterestNetOfTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestNetOfTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.EarningsFromEquityInterestNetOfTax) {
		return nil, false
	}
	return o.EarningsFromEquityInterestNetOfTax, true
}

// HasEarningsFromEquityInterestNetOfTax returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasEarningsFromEquityInterestNetOfTax() bool {
	if o != nil && !IsNil(o.EarningsFromEquityInterestNetOfTax) {
		return true
	}

	return false
}

// SetEarningsFromEquityInterestNetOfTax gets a reference to the given float64 and assigns it to the EarningsFromEquityInterestNetOfTax field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetEarningsFromEquityInterestNetOfTax(v float64) {
	o.EarningsFromEquityInterestNetOfTax = &v
}

// GetTotalOtherFinanceCost returns the TotalOtherFinanceCost field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetTotalOtherFinanceCost() float64 {
	if o == nil || IsNil(o.TotalOtherFinanceCost) {
		var ret float64
		return ret
	}
	return *o.TotalOtherFinanceCost
}

// GetTotalOtherFinanceCostOk returns a tuple with the TotalOtherFinanceCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetTotalOtherFinanceCostOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalOtherFinanceCost) {
		return nil, false
	}
	return o.TotalOtherFinanceCost, true
}

// HasTotalOtherFinanceCost returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasTotalOtherFinanceCost() bool {
	if o != nil && !IsNil(o.TotalOtherFinanceCost) {
		return true
	}

	return false
}

// SetTotalOtherFinanceCost gets a reference to the given float64 and assigns it to the TotalOtherFinanceCost field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetTotalOtherFinanceCost(v float64) {
	o.TotalOtherFinanceCost = &v
}

func (o IncomeStatementItemOtherIncomeAndExpenses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemOtherIncomeAndExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OtherIncomeExpense) {
		toSerialize["other_income_expense"] = o.OtherIncomeExpense
	}
	if !IsNil(o.OtherNonOperatingIncomeExpenses) {
		toSerialize["other_non_operating_income_expenses"] = o.OtherNonOperatingIncomeExpenses
	}
	if !IsNil(o.SpecialIncomeCharges) {
		toSerialize["special_income_charges"] = o.SpecialIncomeCharges
	}
	if !IsNil(o.GainOnSaleOfPpe) {
		toSerialize["gain_on_sale_of_ppe"] = o.GainOnSaleOfPpe
	}
	if !IsNil(o.GainOnSaleOfBusiness) {
		toSerialize["gain_on_sale_of_business"] = o.GainOnSaleOfBusiness
	}
	if !IsNil(o.GainOnSaleOfSecurity) {
		toSerialize["gain_on_sale_of_security"] = o.GainOnSaleOfSecurity
	}
	if !IsNil(o.OtherSpecialCharges) {
		toSerialize["other_special_charges"] = o.OtherSpecialCharges
	}
	if !IsNil(o.WriteOff) {
		toSerialize["write_off"] = o.WriteOff
	}
	if !IsNil(o.ImpairmentOfCapitalAssets) {
		toSerialize["impairment_of_capital_assets"] = o.ImpairmentOfCapitalAssets
	}
	if !IsNil(o.RestructuringAndMergerAcquisition) {
		toSerialize["restructuring_and_merger_acquisition"] = o.RestructuringAndMergerAcquisition
	}
	if !IsNil(o.SecuritiesAmortization) {
		toSerialize["securities_amortization"] = o.SecuritiesAmortization
	}
	if !IsNil(o.EarningsFromEquityInterest) {
		toSerialize["earnings_from_equity_interest"] = o.EarningsFromEquityInterest
	}
	if !IsNil(o.EarningsFromEquityInterestNetOfTax) {
		toSerialize["earnings_from_equity_interest_net_of_tax"] = o.EarningsFromEquityInterestNetOfTax
	}
	if !IsNil(o.TotalOtherFinanceCost) {
		toSerialize["total_other_finance_cost"] = o.TotalOtherFinanceCost
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemOtherIncomeAndExpenses struct {
	value *IncomeStatementItemOtherIncomeAndExpenses
	isSet bool
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) Get() *IncomeStatementItemOtherIncomeAndExpenses {
	return v.value
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) Set(val *IncomeStatementItemOtherIncomeAndExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemOtherIncomeAndExpenses(val *IncomeStatementItemOtherIncomeAndExpenses) *NullableIncomeStatementItemOtherIncomeAndExpenses {
	return &NullableIncomeStatementItemOtherIncomeAndExpenses{value: val, isSet: true}
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


