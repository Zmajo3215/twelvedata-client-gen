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

// checks if the AssetsInfoLiabilitiesNonCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesNonCurrentLiabilities{}

// AssetsInfoLiabilitiesNonCurrentLiabilities Non-current liabilities information
type AssetsInfoLiabilitiesNonCurrentLiabilities struct {
	// Total non current liabilities net minority interest
	TotalNonCurrentLiabilitiesNetMinorityInterest *float64 `json:"total_non_current_liabilities_net_minority_interest,omitempty"`
	LongTermDebtAndCapitalLeaseObligation *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation `json:"long_term_debt_and_capital_lease_obligation,omitempty"`
	// Long term provisions
	LongTermProvisions *float64 `json:"long_term_provisions,omitempty"`
	// Non current pension and other postretirement benefit plans
	NonCurrentPensionAndOtherPostretirementBenefitPlans *float64 `json:"non_current_pension_and_other_postretirement_benefit_plans,omitempty"`
	// Non current accrued expenses
	NonCurrentAccruedExpenses *float64 `json:"non_current_accrued_expenses,omitempty"`
	// Due to related parties non current
	DueToRelatedPartiesNonCurrent *float64 `json:"due_to_related_parties_non_current,omitempty"`
	// Trade and other payables non current
	TradeAndOtherPayablesNonCurrent *float64 `json:"trade_and_other_payables_non_current,omitempty"`
	// Non current deferred liabilities
	NonCurrentDeferredLiabilities *float64 `json:"non_current_deferred_liabilities,omitempty"`
	// Non current deferred revenue
	NonCurrentDeferredRevenue *float64 `json:"non_current_deferred_revenue,omitempty"`
	// Non current deferred taxes liabilities
	NonCurrentDeferredTaxesLiabilities *float64 `json:"non_current_deferred_taxes_liabilities,omitempty"`
	// Other non current liabilities
	OtherNonCurrentLiabilities *float64 `json:"other_non_current_liabilities,omitempty"`
	// Preferred securities outside stock equity
	PreferredSecuritiesOutsideStockEquity *float64 `json:"preferred_securities_outside_stock_equity,omitempty"`
	// Derivative product liabilities
	DerivativeProductLiabilities *float64 `json:"derivative_product_liabilities,omitempty"`
	// Capital lease obligations
	CapitalLeaseObligations *float64 `json:"capital_lease_obligations,omitempty"`
	// Restricted common stock
	RestrictedCommonStock *float64 `json:"restricted_common_stock,omitempty"`
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilities instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesNonCurrentLiabilities() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	this := AssetsInfoLiabilitiesNonCurrentLiabilities{}
	return &this
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesNonCurrentLiabilitiesWithDefaults() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	this := AssetsInfoLiabilitiesNonCurrentLiabilities{}
	return &this
}

// GetTotalNonCurrentLiabilitiesNetMinorityInterest returns the TotalNonCurrentLiabilitiesNetMinorityInterest field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesNetMinorityInterest() float64 {
	if o == nil || IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentLiabilitiesNetMinorityInterest
}

// GetTotalNonCurrentLiabilitiesNetMinorityInterestOk returns a tuple with the TotalNonCurrentLiabilitiesNetMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesNetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		return nil, false
	}
	return o.TotalNonCurrentLiabilitiesNetMinorityInterest, true
}

// HasTotalNonCurrentLiabilitiesNetMinorityInterest returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasTotalNonCurrentLiabilitiesNetMinorityInterest() bool {
	if o != nil && !IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		return true
	}

	return false
}

// SetTotalNonCurrentLiabilitiesNetMinorityInterest gets a reference to the given float64 and assigns it to the TotalNonCurrentLiabilitiesNetMinorityInterest field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetTotalNonCurrentLiabilitiesNetMinorityInterest(v float64) {
	o.TotalNonCurrentLiabilitiesNetMinorityInterest = &v
}

// GetLongTermDebtAndCapitalLeaseObligation returns the LongTermDebtAndCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermDebtAndCapitalLeaseObligation() AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	if o == nil || IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		var ret AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation
		return ret
	}
	return *o.LongTermDebtAndCapitalLeaseObligation
}

// GetLongTermDebtAndCapitalLeaseObligationOk returns a tuple with the LongTermDebtAndCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermDebtAndCapitalLeaseObligationOk() (*AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation, bool) {
	if o == nil || IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		return nil, false
	}
	return o.LongTermDebtAndCapitalLeaseObligation, true
}

// HasLongTermDebtAndCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasLongTermDebtAndCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetLongTermDebtAndCapitalLeaseObligation gets a reference to the given AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation and assigns it to the LongTermDebtAndCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetLongTermDebtAndCapitalLeaseObligation(v AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) {
	o.LongTermDebtAndCapitalLeaseObligation = &v
}

// GetLongTermProvisions returns the LongTermProvisions field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermProvisions() float64 {
	if o == nil || IsNil(o.LongTermProvisions) {
		var ret float64
		return ret
	}
	return *o.LongTermProvisions
}

// GetLongTermProvisionsOk returns a tuple with the LongTermProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermProvisions) {
		return nil, false
	}
	return o.LongTermProvisions, true
}

// HasLongTermProvisions returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasLongTermProvisions() bool {
	if o != nil && !IsNil(o.LongTermProvisions) {
		return true
	}

	return false
}

// SetLongTermProvisions gets a reference to the given float64 and assigns it to the LongTermProvisions field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetLongTermProvisions(v float64) {
	o.LongTermProvisions = &v
}

// GetNonCurrentPensionAndOtherPostretirementBenefitPlans returns the NonCurrentPensionAndOtherPostretirementBenefitPlans field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentPensionAndOtherPostretirementBenefitPlans() float64 {
	if o == nil || IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		var ret float64
		return ret
	}
	return *o.NonCurrentPensionAndOtherPostretirementBenefitPlans
}

// GetNonCurrentPensionAndOtherPostretirementBenefitPlansOk returns a tuple with the NonCurrentPensionAndOtherPostretirementBenefitPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentPensionAndOtherPostretirementBenefitPlansOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		return nil, false
	}
	return o.NonCurrentPensionAndOtherPostretirementBenefitPlans, true
}

// HasNonCurrentPensionAndOtherPostretirementBenefitPlans returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentPensionAndOtherPostretirementBenefitPlans() bool {
	if o != nil && !IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		return true
	}

	return false
}

// SetNonCurrentPensionAndOtherPostretirementBenefitPlans gets a reference to the given float64 and assigns it to the NonCurrentPensionAndOtherPostretirementBenefitPlans field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentPensionAndOtherPostretirementBenefitPlans(v float64) {
	o.NonCurrentPensionAndOtherPostretirementBenefitPlans = &v
}

// GetNonCurrentAccruedExpenses returns the NonCurrentAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentAccruedExpenses() float64 {
	if o == nil || IsNil(o.NonCurrentAccruedExpenses) {
		var ret float64
		return ret
	}
	return *o.NonCurrentAccruedExpenses
}

// GetNonCurrentAccruedExpensesOk returns a tuple with the NonCurrentAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentAccruedExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentAccruedExpenses) {
		return nil, false
	}
	return o.NonCurrentAccruedExpenses, true
}

// HasNonCurrentAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentAccruedExpenses() bool {
	if o != nil && !IsNil(o.NonCurrentAccruedExpenses) {
		return true
	}

	return false
}

// SetNonCurrentAccruedExpenses gets a reference to the given float64 and assigns it to the NonCurrentAccruedExpenses field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentAccruedExpenses(v float64) {
	o.NonCurrentAccruedExpenses = &v
}

// GetDueToRelatedPartiesNonCurrent returns the DueToRelatedPartiesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDueToRelatedPartiesNonCurrent() float64 {
	if o == nil || IsNil(o.DueToRelatedPartiesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.DueToRelatedPartiesNonCurrent
}

// GetDueToRelatedPartiesNonCurrentOk returns a tuple with the DueToRelatedPartiesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDueToRelatedPartiesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueToRelatedPartiesNonCurrent) {
		return nil, false
	}
	return o.DueToRelatedPartiesNonCurrent, true
}

// HasDueToRelatedPartiesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasDueToRelatedPartiesNonCurrent() bool {
	if o != nil && !IsNil(o.DueToRelatedPartiesNonCurrent) {
		return true
	}

	return false
}

// SetDueToRelatedPartiesNonCurrent gets a reference to the given float64 and assigns it to the DueToRelatedPartiesNonCurrent field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetDueToRelatedPartiesNonCurrent(v float64) {
	o.DueToRelatedPartiesNonCurrent = &v
}

// GetTradeAndOtherPayablesNonCurrent returns the TradeAndOtherPayablesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTradeAndOtherPayablesNonCurrent() float64 {
	if o == nil || IsNil(o.TradeAndOtherPayablesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.TradeAndOtherPayablesNonCurrent
}

// GetTradeAndOtherPayablesNonCurrentOk returns a tuple with the TradeAndOtherPayablesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTradeAndOtherPayablesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.TradeAndOtherPayablesNonCurrent) {
		return nil, false
	}
	return o.TradeAndOtherPayablesNonCurrent, true
}

// HasTradeAndOtherPayablesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasTradeAndOtherPayablesNonCurrent() bool {
	if o != nil && !IsNil(o.TradeAndOtherPayablesNonCurrent) {
		return true
	}

	return false
}

// SetTradeAndOtherPayablesNonCurrent gets a reference to the given float64 and assigns it to the TradeAndOtherPayablesNonCurrent field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetTradeAndOtherPayablesNonCurrent(v float64) {
	o.TradeAndOtherPayablesNonCurrent = &v
}

// GetNonCurrentDeferredLiabilities returns the NonCurrentDeferredLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredLiabilities() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredLiabilities
}

// GetNonCurrentDeferredLiabilitiesOk returns a tuple with the NonCurrentDeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredLiabilities) {
		return nil, false
	}
	return o.NonCurrentDeferredLiabilities, true
}

// HasNonCurrentDeferredLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentDeferredLiabilities gets a reference to the given float64 and assigns it to the NonCurrentDeferredLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredLiabilities(v float64) {
	o.NonCurrentDeferredLiabilities = &v
}

// GetNonCurrentDeferredRevenue returns the NonCurrentDeferredRevenue field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredRevenue() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredRevenue) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredRevenue
}

// GetNonCurrentDeferredRevenueOk returns a tuple with the NonCurrentDeferredRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredRevenue) {
		return nil, false
	}
	return o.NonCurrentDeferredRevenue, true
}

// HasNonCurrentDeferredRevenue returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredRevenue() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredRevenue) {
		return true
	}

	return false
}

// SetNonCurrentDeferredRevenue gets a reference to the given float64 and assigns it to the NonCurrentDeferredRevenue field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredRevenue(v float64) {
	o.NonCurrentDeferredRevenue = &v
}

// GetNonCurrentDeferredTaxesLiabilities returns the NonCurrentDeferredTaxesLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredTaxesLiabilities() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredTaxesLiabilities
}

// GetNonCurrentDeferredTaxesLiabilitiesOk returns a tuple with the NonCurrentDeferredTaxesLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredTaxesLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		return nil, false
	}
	return o.NonCurrentDeferredTaxesLiabilities, true
}

// HasNonCurrentDeferredTaxesLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredTaxesLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentDeferredTaxesLiabilities gets a reference to the given float64 and assigns it to the NonCurrentDeferredTaxesLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredTaxesLiabilities(v float64) {
	o.NonCurrentDeferredTaxesLiabilities = &v
}

// GetOtherNonCurrentLiabilities returns the OtherNonCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentLiabilities
}

// GetOtherNonCurrentLiabilitiesOk returns a tuple with the OtherNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		return nil, false
	}
	return o.OtherNonCurrentLiabilities, true
}

// HasOtherNonCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasOtherNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherNonCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherNonCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetOtherNonCurrentLiabilities(v float64) {
	o.OtherNonCurrentLiabilities = &v
}

// GetPreferredSecuritiesOutsideStockEquity returns the PreferredSecuritiesOutsideStockEquity field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetPreferredSecuritiesOutsideStockEquity() float64 {
	if o == nil || IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		var ret float64
		return ret
	}
	return *o.PreferredSecuritiesOutsideStockEquity
}

// GetPreferredSecuritiesOutsideStockEquityOk returns a tuple with the PreferredSecuritiesOutsideStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetPreferredSecuritiesOutsideStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		return nil, false
	}
	return o.PreferredSecuritiesOutsideStockEquity, true
}

// HasPreferredSecuritiesOutsideStockEquity returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasPreferredSecuritiesOutsideStockEquity() bool {
	if o != nil && !IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		return true
	}

	return false
}

// SetPreferredSecuritiesOutsideStockEquity gets a reference to the given float64 and assigns it to the PreferredSecuritiesOutsideStockEquity field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetPreferredSecuritiesOutsideStockEquity(v float64) {
	o.PreferredSecuritiesOutsideStockEquity = &v
}

// GetDerivativeProductLiabilities returns the DerivativeProductLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilities() float64 {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		var ret float64
		return ret
	}
	return *o.DerivativeProductLiabilities
}

// GetDerivativeProductLiabilitiesOk returns a tuple with the DerivativeProductLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		return nil, false
	}
	return o.DerivativeProductLiabilities, true
}

// HasDerivativeProductLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasDerivativeProductLiabilities() bool {
	if o != nil && !IsNil(o.DerivativeProductLiabilities) {
		return true
	}

	return false
}

// SetDerivativeProductLiabilities gets a reference to the given float64 and assigns it to the DerivativeProductLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetDerivativeProductLiabilities(v float64) {
	o.DerivativeProductLiabilities = &v
}

// GetCapitalLeaseObligations returns the CapitalLeaseObligations field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetCapitalLeaseObligations() float64 {
	if o == nil || IsNil(o.CapitalLeaseObligations) {
		var ret float64
		return ret
	}
	return *o.CapitalLeaseObligations
}

// GetCapitalLeaseObligationsOk returns a tuple with the CapitalLeaseObligations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetCapitalLeaseObligationsOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalLeaseObligations) {
		return nil, false
	}
	return o.CapitalLeaseObligations, true
}

// HasCapitalLeaseObligations returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasCapitalLeaseObligations() bool {
	if o != nil && !IsNil(o.CapitalLeaseObligations) {
		return true
	}

	return false
}

// SetCapitalLeaseObligations gets a reference to the given float64 and assigns it to the CapitalLeaseObligations field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetCapitalLeaseObligations(v float64) {
	o.CapitalLeaseObligations = &v
}

// GetRestrictedCommonStock returns the RestrictedCommonStock field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetRestrictedCommonStock() float64 {
	if o == nil || IsNil(o.RestrictedCommonStock) {
		var ret float64
		return ret
	}
	return *o.RestrictedCommonStock
}

// GetRestrictedCommonStockOk returns a tuple with the RestrictedCommonStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetRestrictedCommonStockOk() (*float64, bool) {
	if o == nil || IsNil(o.RestrictedCommonStock) {
		return nil, false
	}
	return o.RestrictedCommonStock, true
}

// HasRestrictedCommonStock returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasRestrictedCommonStock() bool {
	if o != nil && !IsNil(o.RestrictedCommonStock) {
		return true
	}

	return false
}

// SetRestrictedCommonStock gets a reference to the given float64 and assigns it to the RestrictedCommonStock field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetRestrictedCommonStock(v float64) {
	o.RestrictedCommonStock = &v
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		toSerialize["total_non_current_liabilities_net_minority_interest"] = o.TotalNonCurrentLiabilitiesNetMinorityInterest
	}
	if !IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		toSerialize["long_term_debt_and_capital_lease_obligation"] = o.LongTermDebtAndCapitalLeaseObligation
	}
	if !IsNil(o.LongTermProvisions) {
		toSerialize["long_term_provisions"] = o.LongTermProvisions
	}
	if !IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		toSerialize["non_current_pension_and_other_postretirement_benefit_plans"] = o.NonCurrentPensionAndOtherPostretirementBenefitPlans
	}
	if !IsNil(o.NonCurrentAccruedExpenses) {
		toSerialize["non_current_accrued_expenses"] = o.NonCurrentAccruedExpenses
	}
	if !IsNil(o.DueToRelatedPartiesNonCurrent) {
		toSerialize["due_to_related_parties_non_current"] = o.DueToRelatedPartiesNonCurrent
	}
	if !IsNil(o.TradeAndOtherPayablesNonCurrent) {
		toSerialize["trade_and_other_payables_non_current"] = o.TradeAndOtherPayablesNonCurrent
	}
	if !IsNil(o.NonCurrentDeferredLiabilities) {
		toSerialize["non_current_deferred_liabilities"] = o.NonCurrentDeferredLiabilities
	}
	if !IsNil(o.NonCurrentDeferredRevenue) {
		toSerialize["non_current_deferred_revenue"] = o.NonCurrentDeferredRevenue
	}
	if !IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		toSerialize["non_current_deferred_taxes_liabilities"] = o.NonCurrentDeferredTaxesLiabilities
	}
	if !IsNil(o.OtherNonCurrentLiabilities) {
		toSerialize["other_non_current_liabilities"] = o.OtherNonCurrentLiabilities
	}
	if !IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		toSerialize["preferred_securities_outside_stock_equity"] = o.PreferredSecuritiesOutsideStockEquity
	}
	if !IsNil(o.DerivativeProductLiabilities) {
		toSerialize["derivative_product_liabilities"] = o.DerivativeProductLiabilities
	}
	if !IsNil(o.CapitalLeaseObligations) {
		toSerialize["capital_lease_obligations"] = o.CapitalLeaseObligations
	}
	if !IsNil(o.RestrictedCommonStock) {
		toSerialize["restricted_common_stock"] = o.RestrictedCommonStock
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesNonCurrentLiabilities struct {
	value *AssetsInfoLiabilitiesNonCurrentLiabilities
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Get() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Set(val *AssetsInfoLiabilitiesNonCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesNonCurrentLiabilities(val *AssetsInfoLiabilitiesNonCurrentLiabilities) *NullableAssetsInfoLiabilitiesNonCurrentLiabilities {
	return &NullableAssetsInfoLiabilitiesNonCurrentLiabilities{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


