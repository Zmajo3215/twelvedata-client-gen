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

// checks if the AssetsInfoLiabilitiesCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesCurrentLiabilities{}

// AssetsInfoLiabilitiesCurrentLiabilities Current liabilities information
type AssetsInfoLiabilitiesCurrentLiabilities struct {
	// Total current liabilities
	TotalCurrentLiabilities *float64 `json:"total_current_liabilities,omitempty"`
	// Current debt and capital lease obligation
	CurrentDebtAndCapitalLeaseObligation *float64 `json:"current_debt_and_capital_lease_obligation,omitempty"`
	// Current debt
	CurrentDebt *float64 `json:"current_debt,omitempty"`
	// Current capital lease obligation
	CurrentCapitalLeaseObligation *float64 `json:"current_capital_lease_obligation,omitempty"`
	// Other current borrowings
	OtherCurrentBorrowings *float64 `json:"other_current_borrowings,omitempty"`
	// Line of credit
	LineOfCredit *float64 `json:"line_of_credit,omitempty"`
	// Commercial paper
	CommercialPaper *float64 `json:"commercial_paper,omitempty"`
	// Current notes payable
	CurrentNotesPayable *float64 `json:"current_notes_payable,omitempty"`
	// Current provisions
	CurrentProvisions *float64 `json:"current_provisions,omitempty"`
	PayablesAndAccruedExpenses *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses `json:"payables_and_accrued_expenses,omitempty"`
	// Pension and other post retirement benefit plans current
	PensionAndOtherPostRetirementBenefitPlansCurrent *float64 `json:"pension_and_other_post_retirement_benefit_plans_current,omitempty"`
	// Employee benefits
	EmployeeBenefits *float64 `json:"employee_benefits,omitempty"`
	// Current deferred liabilities
	CurrentDeferredLiabilities *float64 `json:"current_deferred_liabilities,omitempty"`
	// Current deferred revenue
	CurrentDeferredRevenue *float64 `json:"current_deferred_revenue,omitempty"`
	// Current deferred taxes liabilities
	CurrentDeferredTaxesLiabilities *float64 `json:"current_deferred_taxes_liabilities,omitempty"`
	// Other current liabilities
	OtherCurrentLiabilities *float64 `json:"other_current_liabilities,omitempty"`
	// Liabilities held for sale non current
	LiabilitiesHeldForSaleNonCurrent *float64 `json:"liabilities_held_for_sale_non_current,omitempty"`
}

// NewAssetsInfoLiabilitiesCurrentLiabilities instantiates a new AssetsInfoLiabilitiesCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesCurrentLiabilities() *AssetsInfoLiabilitiesCurrentLiabilities {
	this := AssetsInfoLiabilitiesCurrentLiabilities{}
	return &this
}

// NewAssetsInfoLiabilitiesCurrentLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilitiesCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesCurrentLiabilitiesWithDefaults() *AssetsInfoLiabilitiesCurrentLiabilities {
	this := AssetsInfoLiabilitiesCurrentLiabilities{}
	return &this
}

// GetTotalCurrentLiabilities returns the TotalCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilities() float64 {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentLiabilities
}

// GetTotalCurrentLiabilitiesOk returns a tuple with the TotalCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		return nil, false
	}
	return o.TotalCurrentLiabilities, true
}

// HasTotalCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasTotalCurrentLiabilities() bool {
	if o != nil && !IsNil(o.TotalCurrentLiabilities) {
		return true
	}

	return false
}

// SetTotalCurrentLiabilities gets a reference to the given float64 and assigns it to the TotalCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetTotalCurrentLiabilities(v float64) {
	o.TotalCurrentLiabilities = &v
}

// GetCurrentDebtAndCapitalLeaseObligation returns the CurrentDebtAndCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtAndCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.CurrentDebtAndCapitalLeaseObligation
}

// GetCurrentDebtAndCapitalLeaseObligationOk returns a tuple with the CurrentDebtAndCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtAndCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		return nil, false
	}
	return o.CurrentDebtAndCapitalLeaseObligation, true
}

// HasCurrentDebtAndCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDebtAndCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetCurrentDebtAndCapitalLeaseObligation gets a reference to the given float64 and assigns it to the CurrentDebtAndCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDebtAndCapitalLeaseObligation(v float64) {
	o.CurrentDebtAndCapitalLeaseObligation = &v
}

// GetCurrentDebt returns the CurrentDebt field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebt() float64 {
	if o == nil || IsNil(o.CurrentDebt) {
		var ret float64
		return ret
	}
	return *o.CurrentDebt
}

// GetCurrentDebtOk returns a tuple with the CurrentDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDebt) {
		return nil, false
	}
	return o.CurrentDebt, true
}

// HasCurrentDebt returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDebt() bool {
	if o != nil && !IsNil(o.CurrentDebt) {
		return true
	}

	return false
}

// SetCurrentDebt gets a reference to the given float64 and assigns it to the CurrentDebt field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDebt(v float64) {
	o.CurrentDebt = &v
}

// GetCurrentCapitalLeaseObligation returns the CurrentCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.CurrentCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.CurrentCapitalLeaseObligation
}

// GetCurrentCapitalLeaseObligationOk returns a tuple with the CurrentCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentCapitalLeaseObligation) {
		return nil, false
	}
	return o.CurrentCapitalLeaseObligation, true
}

// HasCurrentCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.CurrentCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetCurrentCapitalLeaseObligation gets a reference to the given float64 and assigns it to the CurrentCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentCapitalLeaseObligation(v float64) {
	o.CurrentCapitalLeaseObligation = &v
}

// GetOtherCurrentBorrowings returns the OtherCurrentBorrowings field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentBorrowings() float64 {
	if o == nil || IsNil(o.OtherCurrentBorrowings) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentBorrowings
}

// GetOtherCurrentBorrowingsOk returns a tuple with the OtherCurrentBorrowings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentBorrowingsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentBorrowings) {
		return nil, false
	}
	return o.OtherCurrentBorrowings, true
}

// HasOtherCurrentBorrowings returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasOtherCurrentBorrowings() bool {
	if o != nil && !IsNil(o.OtherCurrentBorrowings) {
		return true
	}

	return false
}

// SetOtherCurrentBorrowings gets a reference to the given float64 and assigns it to the OtherCurrentBorrowings field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetOtherCurrentBorrowings(v float64) {
	o.OtherCurrentBorrowings = &v
}

// GetLineOfCredit returns the LineOfCredit field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLineOfCredit() float64 {
	if o == nil || IsNil(o.LineOfCredit) {
		var ret float64
		return ret
	}
	return *o.LineOfCredit
}

// GetLineOfCreditOk returns a tuple with the LineOfCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLineOfCreditOk() (*float64, bool) {
	if o == nil || IsNil(o.LineOfCredit) {
		return nil, false
	}
	return o.LineOfCredit, true
}

// HasLineOfCredit returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasLineOfCredit() bool {
	if o != nil && !IsNil(o.LineOfCredit) {
		return true
	}

	return false
}

// SetLineOfCredit gets a reference to the given float64 and assigns it to the LineOfCredit field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetLineOfCredit(v float64) {
	o.LineOfCredit = &v
}

// GetCommercialPaper returns the CommercialPaper field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCommercialPaper() float64 {
	if o == nil || IsNil(o.CommercialPaper) {
		var ret float64
		return ret
	}
	return *o.CommercialPaper
}

// GetCommercialPaperOk returns a tuple with the CommercialPaper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCommercialPaperOk() (*float64, bool) {
	if o == nil || IsNil(o.CommercialPaper) {
		return nil, false
	}
	return o.CommercialPaper, true
}

// HasCommercialPaper returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCommercialPaper() bool {
	if o != nil && !IsNil(o.CommercialPaper) {
		return true
	}

	return false
}

// SetCommercialPaper gets a reference to the given float64 and assigns it to the CommercialPaper field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCommercialPaper(v float64) {
	o.CommercialPaper = &v
}

// GetCurrentNotesPayable returns the CurrentNotesPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentNotesPayable() float64 {
	if o == nil || IsNil(o.CurrentNotesPayable) {
		var ret float64
		return ret
	}
	return *o.CurrentNotesPayable
}

// GetCurrentNotesPayableOk returns a tuple with the CurrentNotesPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentNotesPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentNotesPayable) {
		return nil, false
	}
	return o.CurrentNotesPayable, true
}

// HasCurrentNotesPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentNotesPayable() bool {
	if o != nil && !IsNil(o.CurrentNotesPayable) {
		return true
	}

	return false
}

// SetCurrentNotesPayable gets a reference to the given float64 and assigns it to the CurrentNotesPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentNotesPayable(v float64) {
	o.CurrentNotesPayable = &v
}

// GetCurrentProvisions returns the CurrentProvisions field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentProvisions() float64 {
	if o == nil || IsNil(o.CurrentProvisions) {
		var ret float64
		return ret
	}
	return *o.CurrentProvisions
}

// GetCurrentProvisionsOk returns a tuple with the CurrentProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentProvisions) {
		return nil, false
	}
	return o.CurrentProvisions, true
}

// HasCurrentProvisions returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentProvisions() bool {
	if o != nil && !IsNil(o.CurrentProvisions) {
		return true
	}

	return false
}

// SetCurrentProvisions gets a reference to the given float64 and assigns it to the CurrentProvisions field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentProvisions(v float64) {
	o.CurrentProvisions = &v
}

// GetPayablesAndAccruedExpenses returns the PayablesAndAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPayablesAndAccruedExpenses() AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	if o == nil || IsNil(o.PayablesAndAccruedExpenses) {
		var ret AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses
		return ret
	}
	return *o.PayablesAndAccruedExpenses
}

// GetPayablesAndAccruedExpensesOk returns a tuple with the PayablesAndAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPayablesAndAccruedExpensesOk() (*AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses, bool) {
	if o == nil || IsNil(o.PayablesAndAccruedExpenses) {
		return nil, false
	}
	return o.PayablesAndAccruedExpenses, true
}

// HasPayablesAndAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasPayablesAndAccruedExpenses() bool {
	if o != nil && !IsNil(o.PayablesAndAccruedExpenses) {
		return true
	}

	return false
}

// SetPayablesAndAccruedExpenses gets a reference to the given AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses and assigns it to the PayablesAndAccruedExpenses field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetPayablesAndAccruedExpenses(v AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) {
	o.PayablesAndAccruedExpenses = &v
}

// GetPensionAndOtherPostRetirementBenefitPlansCurrent returns the PensionAndOtherPostRetirementBenefitPlansCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPensionAndOtherPostRetirementBenefitPlansCurrent() float64 {
	if o == nil || IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		var ret float64
		return ret
	}
	return *o.PensionAndOtherPostRetirementBenefitPlansCurrent
}

// GetPensionAndOtherPostRetirementBenefitPlansCurrentOk returns a tuple with the PensionAndOtherPostRetirementBenefitPlansCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPensionAndOtherPostRetirementBenefitPlansCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		return nil, false
	}
	return o.PensionAndOtherPostRetirementBenefitPlansCurrent, true
}

// HasPensionAndOtherPostRetirementBenefitPlansCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasPensionAndOtherPostRetirementBenefitPlansCurrent() bool {
	if o != nil && !IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		return true
	}

	return false
}

// SetPensionAndOtherPostRetirementBenefitPlansCurrent gets a reference to the given float64 and assigns it to the PensionAndOtherPostRetirementBenefitPlansCurrent field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetPensionAndOtherPostRetirementBenefitPlansCurrent(v float64) {
	o.PensionAndOtherPostRetirementBenefitPlansCurrent = &v
}

// GetEmployeeBenefits returns the EmployeeBenefits field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetEmployeeBenefits() float64 {
	if o == nil || IsNil(o.EmployeeBenefits) {
		var ret float64
		return ret
	}
	return *o.EmployeeBenefits
}

// GetEmployeeBenefitsOk returns a tuple with the EmployeeBenefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetEmployeeBenefitsOk() (*float64, bool) {
	if o == nil || IsNil(o.EmployeeBenefits) {
		return nil, false
	}
	return o.EmployeeBenefits, true
}

// HasEmployeeBenefits returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasEmployeeBenefits() bool {
	if o != nil && !IsNil(o.EmployeeBenefits) {
		return true
	}

	return false
}

// SetEmployeeBenefits gets a reference to the given float64 and assigns it to the EmployeeBenefits field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetEmployeeBenefits(v float64) {
	o.EmployeeBenefits = &v
}

// GetCurrentDeferredLiabilities returns the CurrentDeferredLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredLiabilities() float64 {
	if o == nil || IsNil(o.CurrentDeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredLiabilities
}

// GetCurrentDeferredLiabilitiesOk returns a tuple with the CurrentDeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredLiabilities) {
		return nil, false
	}
	return o.CurrentDeferredLiabilities, true
}

// HasCurrentDeferredLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredLiabilities() bool {
	if o != nil && !IsNil(o.CurrentDeferredLiabilities) {
		return true
	}

	return false
}

// SetCurrentDeferredLiabilities gets a reference to the given float64 and assigns it to the CurrentDeferredLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredLiabilities(v float64) {
	o.CurrentDeferredLiabilities = &v
}

// GetCurrentDeferredRevenue returns the CurrentDeferredRevenue field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredRevenue() float64 {
	if o == nil || IsNil(o.CurrentDeferredRevenue) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredRevenue
}

// GetCurrentDeferredRevenueOk returns a tuple with the CurrentDeferredRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredRevenue) {
		return nil, false
	}
	return o.CurrentDeferredRevenue, true
}

// HasCurrentDeferredRevenue returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredRevenue() bool {
	if o != nil && !IsNil(o.CurrentDeferredRevenue) {
		return true
	}

	return false
}

// SetCurrentDeferredRevenue gets a reference to the given float64 and assigns it to the CurrentDeferredRevenue field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredRevenue(v float64) {
	o.CurrentDeferredRevenue = &v
}

// GetCurrentDeferredTaxesLiabilities returns the CurrentDeferredTaxesLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredTaxesLiabilities() float64 {
	if o == nil || IsNil(o.CurrentDeferredTaxesLiabilities) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredTaxesLiabilities
}

// GetCurrentDeferredTaxesLiabilitiesOk returns a tuple with the CurrentDeferredTaxesLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredTaxesLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredTaxesLiabilities) {
		return nil, false
	}
	return o.CurrentDeferredTaxesLiabilities, true
}

// HasCurrentDeferredTaxesLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredTaxesLiabilities() bool {
	if o != nil && !IsNil(o.CurrentDeferredTaxesLiabilities) {
		return true
	}

	return false
}

// SetCurrentDeferredTaxesLiabilities gets a reference to the given float64 and assigns it to the CurrentDeferredTaxesLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredTaxesLiabilities(v float64) {
	o.CurrentDeferredTaxesLiabilities = &v
}

// GetOtherCurrentLiabilities returns the OtherCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentLiabilities
}

// GetOtherCurrentLiabilitiesOk returns a tuple with the OtherCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		return nil, false
	}
	return o.OtherCurrentLiabilities, true
}

// HasOtherCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasOtherCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetOtherCurrentLiabilities(v float64) {
	o.OtherCurrentLiabilities = &v
}

// GetLiabilitiesHeldForSaleNonCurrent returns the LiabilitiesHeldForSaleNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLiabilitiesHeldForSaleNonCurrent() float64 {
	if o == nil || IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		var ret float64
		return ret
	}
	return *o.LiabilitiesHeldForSaleNonCurrent
}

// GetLiabilitiesHeldForSaleNonCurrentOk returns a tuple with the LiabilitiesHeldForSaleNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLiabilitiesHeldForSaleNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		return nil, false
	}
	return o.LiabilitiesHeldForSaleNonCurrent, true
}

// HasLiabilitiesHeldForSaleNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasLiabilitiesHeldForSaleNonCurrent() bool {
	if o != nil && !IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		return true
	}

	return false
}

// SetLiabilitiesHeldForSaleNonCurrent gets a reference to the given float64 and assigns it to the LiabilitiesHeldForSaleNonCurrent field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetLiabilitiesHeldForSaleNonCurrent(v float64) {
	o.LiabilitiesHeldForSaleNonCurrent = &v
}

func (o AssetsInfoLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCurrentLiabilities) {
		toSerialize["total_current_liabilities"] = o.TotalCurrentLiabilities
	}
	if !IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		toSerialize["current_debt_and_capital_lease_obligation"] = o.CurrentDebtAndCapitalLeaseObligation
	}
	if !IsNil(o.CurrentDebt) {
		toSerialize["current_debt"] = o.CurrentDebt
	}
	if !IsNil(o.CurrentCapitalLeaseObligation) {
		toSerialize["current_capital_lease_obligation"] = o.CurrentCapitalLeaseObligation
	}
	if !IsNil(o.OtherCurrentBorrowings) {
		toSerialize["other_current_borrowings"] = o.OtherCurrentBorrowings
	}
	if !IsNil(o.LineOfCredit) {
		toSerialize["line_of_credit"] = o.LineOfCredit
	}
	if !IsNil(o.CommercialPaper) {
		toSerialize["commercial_paper"] = o.CommercialPaper
	}
	if !IsNil(o.CurrentNotesPayable) {
		toSerialize["current_notes_payable"] = o.CurrentNotesPayable
	}
	if !IsNil(o.CurrentProvisions) {
		toSerialize["current_provisions"] = o.CurrentProvisions
	}
	if !IsNil(o.PayablesAndAccruedExpenses) {
		toSerialize["payables_and_accrued_expenses"] = o.PayablesAndAccruedExpenses
	}
	if !IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		toSerialize["pension_and_other_post_retirement_benefit_plans_current"] = o.PensionAndOtherPostRetirementBenefitPlansCurrent
	}
	if !IsNil(o.EmployeeBenefits) {
		toSerialize["employee_benefits"] = o.EmployeeBenefits
	}
	if !IsNil(o.CurrentDeferredLiabilities) {
		toSerialize["current_deferred_liabilities"] = o.CurrentDeferredLiabilities
	}
	if !IsNil(o.CurrentDeferredRevenue) {
		toSerialize["current_deferred_revenue"] = o.CurrentDeferredRevenue
	}
	if !IsNil(o.CurrentDeferredTaxesLiabilities) {
		toSerialize["current_deferred_taxes_liabilities"] = o.CurrentDeferredTaxesLiabilities
	}
	if !IsNil(o.OtherCurrentLiabilities) {
		toSerialize["other_current_liabilities"] = o.OtherCurrentLiabilities
	}
	if !IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		toSerialize["liabilities_held_for_sale_non_current"] = o.LiabilitiesHeldForSaleNonCurrent
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesCurrentLiabilities struct {
	value *AssetsInfoLiabilitiesCurrentLiabilities
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) Get() *AssetsInfoLiabilitiesCurrentLiabilities {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) Set(val *AssetsInfoLiabilitiesCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesCurrentLiabilities(val *AssetsInfoLiabilitiesCurrentLiabilities) *NullableAssetsInfoLiabilitiesCurrentLiabilities {
	return &NullableAssetsInfoLiabilitiesCurrentLiabilities{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


