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

// checks if the IncomeStatementItemExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemExpenses{}

// IncomeStatementItemExpenses Expenses information
type IncomeStatementItemExpenses struct {
	// Total expenses
	TotalExpenses *float64 `json:"total_expenses,omitempty"`
	// Selling general and administration expense
	SellingGeneralAndAdministrationExpense *float64 `json:"selling_general_and_administration_expense,omitempty"`
	// Selling and marketing expense
	SellingAndMarketingExpense *float64 `json:"selling_and_marketing_expense,omitempty"`
	// General and administrative expense
	GeneralAndAdministrativeExpense *float64 `json:"general_and_administrative_expense,omitempty"`
	// Other general and administrative expense
	OtherGeneralAndAdministrativeExpense *float64 `json:"other_general_and_administrative_expense,omitempty"`
	// Depreciation amortization depletion income statement
	DepreciationAmortizationDepletionIncomeStatement *float64 `json:"depreciation_amortization_depletion_income_statement,omitempty"`
	// Research and development expense
	ResearchAndDevelopmentExpense *float64 `json:"research_and_development_expense,omitempty"`
	// Insurance and claims expense
	InsuranceAndClaimsExpense *float64 `json:"insurance_and_claims_expense,omitempty"`
	// Rent and landing fees
	RentAndLandingFees *float64 `json:"rent_and_landing_fees,omitempty"`
	// Salaries and wages expense
	SalariesAndWagesExpense *float64 `json:"salaries_and_wages_expense,omitempty"`
	// Rent expense supplemental
	RentExpenseSupplemental *float64 `json:"rent_expense_supplemental,omitempty"`
	// Provision for doubtful accounts
	ProvisionForDoubtfulAccounts *float64 `json:"provision_for_doubtful_accounts,omitempty"`
}

// NewIncomeStatementItemExpenses instantiates a new IncomeStatementItemExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemExpenses() *IncomeStatementItemExpenses {
	this := IncomeStatementItemExpenses{}
	return &this
}

// NewIncomeStatementItemExpensesWithDefaults instantiates a new IncomeStatementItemExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemExpensesWithDefaults() *IncomeStatementItemExpenses {
	this := IncomeStatementItemExpenses{}
	return &this
}

// GetTotalExpenses returns the TotalExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetTotalExpenses() float64 {
	if o == nil || IsNil(o.TotalExpenses) {
		var ret float64
		return ret
	}
	return *o.TotalExpenses
}

// GetTotalExpensesOk returns a tuple with the TotalExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetTotalExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalExpenses) {
		return nil, false
	}
	return o.TotalExpenses, true
}

// HasTotalExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasTotalExpenses() bool {
	if o != nil && !IsNil(o.TotalExpenses) {
		return true
	}

	return false
}

// SetTotalExpenses gets a reference to the given float64 and assigns it to the TotalExpenses field.
func (o *IncomeStatementItemExpenses) SetTotalExpenses(v float64) {
	o.TotalExpenses = &v
}

// GetSellingGeneralAndAdministrationExpense returns the SellingGeneralAndAdministrationExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSellingGeneralAndAdministrationExpense() float64 {
	if o == nil || IsNil(o.SellingGeneralAndAdministrationExpense) {
		var ret float64
		return ret
	}
	return *o.SellingGeneralAndAdministrationExpense
}

// GetSellingGeneralAndAdministrationExpenseOk returns a tuple with the SellingGeneralAndAdministrationExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSellingGeneralAndAdministrationExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SellingGeneralAndAdministrationExpense) {
		return nil, false
	}
	return o.SellingGeneralAndAdministrationExpense, true
}

// HasSellingGeneralAndAdministrationExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSellingGeneralAndAdministrationExpense() bool {
	if o != nil && !IsNil(o.SellingGeneralAndAdministrationExpense) {
		return true
	}

	return false
}

// SetSellingGeneralAndAdministrationExpense gets a reference to the given float64 and assigns it to the SellingGeneralAndAdministrationExpense field.
func (o *IncomeStatementItemExpenses) SetSellingGeneralAndAdministrationExpense(v float64) {
	o.SellingGeneralAndAdministrationExpense = &v
}

// GetSellingAndMarketingExpense returns the SellingAndMarketingExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSellingAndMarketingExpense() float64 {
	if o == nil || IsNil(o.SellingAndMarketingExpense) {
		var ret float64
		return ret
	}
	return *o.SellingAndMarketingExpense
}

// GetSellingAndMarketingExpenseOk returns a tuple with the SellingAndMarketingExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSellingAndMarketingExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SellingAndMarketingExpense) {
		return nil, false
	}
	return o.SellingAndMarketingExpense, true
}

// HasSellingAndMarketingExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSellingAndMarketingExpense() bool {
	if o != nil && !IsNil(o.SellingAndMarketingExpense) {
		return true
	}

	return false
}

// SetSellingAndMarketingExpense gets a reference to the given float64 and assigns it to the SellingAndMarketingExpense field.
func (o *IncomeStatementItemExpenses) SetSellingAndMarketingExpense(v float64) {
	o.SellingAndMarketingExpense = &v
}

// GetGeneralAndAdministrativeExpense returns the GeneralAndAdministrativeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetGeneralAndAdministrativeExpense() float64 {
	if o == nil || IsNil(o.GeneralAndAdministrativeExpense) {
		var ret float64
		return ret
	}
	return *o.GeneralAndAdministrativeExpense
}

// GetGeneralAndAdministrativeExpenseOk returns a tuple with the GeneralAndAdministrativeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetGeneralAndAdministrativeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.GeneralAndAdministrativeExpense) {
		return nil, false
	}
	return o.GeneralAndAdministrativeExpense, true
}

// HasGeneralAndAdministrativeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasGeneralAndAdministrativeExpense() bool {
	if o != nil && !IsNil(o.GeneralAndAdministrativeExpense) {
		return true
	}

	return false
}

// SetGeneralAndAdministrativeExpense gets a reference to the given float64 and assigns it to the GeneralAndAdministrativeExpense field.
func (o *IncomeStatementItemExpenses) SetGeneralAndAdministrativeExpense(v float64) {
	o.GeneralAndAdministrativeExpense = &v
}

// GetOtherGeneralAndAdministrativeExpense returns the OtherGeneralAndAdministrativeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetOtherGeneralAndAdministrativeExpense() float64 {
	if o == nil || IsNil(o.OtherGeneralAndAdministrativeExpense) {
		var ret float64
		return ret
	}
	return *o.OtherGeneralAndAdministrativeExpense
}

// GetOtherGeneralAndAdministrativeExpenseOk returns a tuple with the OtherGeneralAndAdministrativeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetOtherGeneralAndAdministrativeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherGeneralAndAdministrativeExpense) {
		return nil, false
	}
	return o.OtherGeneralAndAdministrativeExpense, true
}

// HasOtherGeneralAndAdministrativeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasOtherGeneralAndAdministrativeExpense() bool {
	if o != nil && !IsNil(o.OtherGeneralAndAdministrativeExpense) {
		return true
	}

	return false
}

// SetOtherGeneralAndAdministrativeExpense gets a reference to the given float64 and assigns it to the OtherGeneralAndAdministrativeExpense field.
func (o *IncomeStatementItemExpenses) SetOtherGeneralAndAdministrativeExpense(v float64) {
	o.OtherGeneralAndAdministrativeExpense = &v
}

// GetDepreciationAmortizationDepletionIncomeStatement returns the DepreciationAmortizationDepletionIncomeStatement field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetDepreciationAmortizationDepletionIncomeStatement() float64 {
	if o == nil || IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		var ret float64
		return ret
	}
	return *o.DepreciationAmortizationDepletionIncomeStatement
}

// GetDepreciationAmortizationDepletionIncomeStatementOk returns a tuple with the DepreciationAmortizationDepletionIncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetDepreciationAmortizationDepletionIncomeStatementOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		return nil, false
	}
	return o.DepreciationAmortizationDepletionIncomeStatement, true
}

// HasDepreciationAmortizationDepletionIncomeStatement returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasDepreciationAmortizationDepletionIncomeStatement() bool {
	if o != nil && !IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		return true
	}

	return false
}

// SetDepreciationAmortizationDepletionIncomeStatement gets a reference to the given float64 and assigns it to the DepreciationAmortizationDepletionIncomeStatement field.
func (o *IncomeStatementItemExpenses) SetDepreciationAmortizationDepletionIncomeStatement(v float64) {
	o.DepreciationAmortizationDepletionIncomeStatement = &v
}

// GetResearchAndDevelopmentExpense returns the ResearchAndDevelopmentExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetResearchAndDevelopmentExpense() float64 {
	if o == nil || IsNil(o.ResearchAndDevelopmentExpense) {
		var ret float64
		return ret
	}
	return *o.ResearchAndDevelopmentExpense
}

// GetResearchAndDevelopmentExpenseOk returns a tuple with the ResearchAndDevelopmentExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetResearchAndDevelopmentExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.ResearchAndDevelopmentExpense) {
		return nil, false
	}
	return o.ResearchAndDevelopmentExpense, true
}

// HasResearchAndDevelopmentExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasResearchAndDevelopmentExpense() bool {
	if o != nil && !IsNil(o.ResearchAndDevelopmentExpense) {
		return true
	}

	return false
}

// SetResearchAndDevelopmentExpense gets a reference to the given float64 and assigns it to the ResearchAndDevelopmentExpense field.
func (o *IncomeStatementItemExpenses) SetResearchAndDevelopmentExpense(v float64) {
	o.ResearchAndDevelopmentExpense = &v
}

// GetInsuranceAndClaimsExpense returns the InsuranceAndClaimsExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetInsuranceAndClaimsExpense() float64 {
	if o == nil || IsNil(o.InsuranceAndClaimsExpense) {
		var ret float64
		return ret
	}
	return *o.InsuranceAndClaimsExpense
}

// GetInsuranceAndClaimsExpenseOk returns a tuple with the InsuranceAndClaimsExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetInsuranceAndClaimsExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.InsuranceAndClaimsExpense) {
		return nil, false
	}
	return o.InsuranceAndClaimsExpense, true
}

// HasInsuranceAndClaimsExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasInsuranceAndClaimsExpense() bool {
	if o != nil && !IsNil(o.InsuranceAndClaimsExpense) {
		return true
	}

	return false
}

// SetInsuranceAndClaimsExpense gets a reference to the given float64 and assigns it to the InsuranceAndClaimsExpense field.
func (o *IncomeStatementItemExpenses) SetInsuranceAndClaimsExpense(v float64) {
	o.InsuranceAndClaimsExpense = &v
}

// GetRentAndLandingFees returns the RentAndLandingFees field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetRentAndLandingFees() float64 {
	if o == nil || IsNil(o.RentAndLandingFees) {
		var ret float64
		return ret
	}
	return *o.RentAndLandingFees
}

// GetRentAndLandingFeesOk returns a tuple with the RentAndLandingFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetRentAndLandingFeesOk() (*float64, bool) {
	if o == nil || IsNil(o.RentAndLandingFees) {
		return nil, false
	}
	return o.RentAndLandingFees, true
}

// HasRentAndLandingFees returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasRentAndLandingFees() bool {
	if o != nil && !IsNil(o.RentAndLandingFees) {
		return true
	}

	return false
}

// SetRentAndLandingFees gets a reference to the given float64 and assigns it to the RentAndLandingFees field.
func (o *IncomeStatementItemExpenses) SetRentAndLandingFees(v float64) {
	o.RentAndLandingFees = &v
}

// GetSalariesAndWagesExpense returns the SalariesAndWagesExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSalariesAndWagesExpense() float64 {
	if o == nil || IsNil(o.SalariesAndWagesExpense) {
		var ret float64
		return ret
	}
	return *o.SalariesAndWagesExpense
}

// GetSalariesAndWagesExpenseOk returns a tuple with the SalariesAndWagesExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSalariesAndWagesExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SalariesAndWagesExpense) {
		return nil, false
	}
	return o.SalariesAndWagesExpense, true
}

// HasSalariesAndWagesExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSalariesAndWagesExpense() bool {
	if o != nil && !IsNil(o.SalariesAndWagesExpense) {
		return true
	}

	return false
}

// SetSalariesAndWagesExpense gets a reference to the given float64 and assigns it to the SalariesAndWagesExpense field.
func (o *IncomeStatementItemExpenses) SetSalariesAndWagesExpense(v float64) {
	o.SalariesAndWagesExpense = &v
}

// GetRentExpenseSupplemental returns the RentExpenseSupplemental field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetRentExpenseSupplemental() float64 {
	if o == nil || IsNil(o.RentExpenseSupplemental) {
		var ret float64
		return ret
	}
	return *o.RentExpenseSupplemental
}

// GetRentExpenseSupplementalOk returns a tuple with the RentExpenseSupplemental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetRentExpenseSupplementalOk() (*float64, bool) {
	if o == nil || IsNil(o.RentExpenseSupplemental) {
		return nil, false
	}
	return o.RentExpenseSupplemental, true
}

// HasRentExpenseSupplemental returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasRentExpenseSupplemental() bool {
	if o != nil && !IsNil(o.RentExpenseSupplemental) {
		return true
	}

	return false
}

// SetRentExpenseSupplemental gets a reference to the given float64 and assigns it to the RentExpenseSupplemental field.
func (o *IncomeStatementItemExpenses) SetRentExpenseSupplemental(v float64) {
	o.RentExpenseSupplemental = &v
}

// GetProvisionForDoubtfulAccounts returns the ProvisionForDoubtfulAccounts field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetProvisionForDoubtfulAccounts() float64 {
	if o == nil || IsNil(o.ProvisionForDoubtfulAccounts) {
		var ret float64
		return ret
	}
	return *o.ProvisionForDoubtfulAccounts
}

// GetProvisionForDoubtfulAccountsOk returns a tuple with the ProvisionForDoubtfulAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetProvisionForDoubtfulAccountsOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionForDoubtfulAccounts) {
		return nil, false
	}
	return o.ProvisionForDoubtfulAccounts, true
}

// HasProvisionForDoubtfulAccounts returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasProvisionForDoubtfulAccounts() bool {
	if o != nil && !IsNil(o.ProvisionForDoubtfulAccounts) {
		return true
	}

	return false
}

// SetProvisionForDoubtfulAccounts gets a reference to the given float64 and assigns it to the ProvisionForDoubtfulAccounts field.
func (o *IncomeStatementItemExpenses) SetProvisionForDoubtfulAccounts(v float64) {
	o.ProvisionForDoubtfulAccounts = &v
}

func (o IncomeStatementItemExpenses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalExpenses) {
		toSerialize["total_expenses"] = o.TotalExpenses
	}
	if !IsNil(o.SellingGeneralAndAdministrationExpense) {
		toSerialize["selling_general_and_administration_expense"] = o.SellingGeneralAndAdministrationExpense
	}
	if !IsNil(o.SellingAndMarketingExpense) {
		toSerialize["selling_and_marketing_expense"] = o.SellingAndMarketingExpense
	}
	if !IsNil(o.GeneralAndAdministrativeExpense) {
		toSerialize["general_and_administrative_expense"] = o.GeneralAndAdministrativeExpense
	}
	if !IsNil(o.OtherGeneralAndAdministrativeExpense) {
		toSerialize["other_general_and_administrative_expense"] = o.OtherGeneralAndAdministrativeExpense
	}
	if !IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		toSerialize["depreciation_amortization_depletion_income_statement"] = o.DepreciationAmortizationDepletionIncomeStatement
	}
	if !IsNil(o.ResearchAndDevelopmentExpense) {
		toSerialize["research_and_development_expense"] = o.ResearchAndDevelopmentExpense
	}
	if !IsNil(o.InsuranceAndClaimsExpense) {
		toSerialize["insurance_and_claims_expense"] = o.InsuranceAndClaimsExpense
	}
	if !IsNil(o.RentAndLandingFees) {
		toSerialize["rent_and_landing_fees"] = o.RentAndLandingFees
	}
	if !IsNil(o.SalariesAndWagesExpense) {
		toSerialize["salaries_and_wages_expense"] = o.SalariesAndWagesExpense
	}
	if !IsNil(o.RentExpenseSupplemental) {
		toSerialize["rent_expense_supplemental"] = o.RentExpenseSupplemental
	}
	if !IsNil(o.ProvisionForDoubtfulAccounts) {
		toSerialize["provision_for_doubtful_accounts"] = o.ProvisionForDoubtfulAccounts
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemExpenses struct {
	value *IncomeStatementItemExpenses
	isSet bool
}

func (v NullableIncomeStatementItemExpenses) Get() *IncomeStatementItemExpenses {
	return v.value
}

func (v *NullableIncomeStatementItemExpenses) Set(val *IncomeStatementItemExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemExpenses(val *IncomeStatementItemExpenses) *NullableIncomeStatementItemExpenses {
	return &NullableIncomeStatementItemExpenses{value: val, isSet: true}
}

func (v NullableIncomeStatementItemExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


