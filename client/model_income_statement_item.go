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

// checks if the IncomeStatementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItem{}

// IncomeStatementItem IncomeStatementItem represents a single income statement record
type IncomeStatementItem struct {
	// Date of the income statement release
	FiscalDate *string `json:"fiscal_date,omitempty"`
	// Fiscal year
	Year *int64 `json:"year,omitempty"`
	Revenue *IncomeStatementItemRevenue `json:"revenue,omitempty"`
	GrossProfit *IncomeStatementItemGrossProfit `json:"gross_profit,omitempty"`
	OperatingIncome *IncomeStatementItemOperatingIncome `json:"operating_income,omitempty"`
	NetIncome *IncomeStatementItemNetIncome `json:"net_income,omitempty"`
	EarningsPerShare *IncomeStatementItemEarningsPerShare `json:"earnings_per_share,omitempty"`
	Expenses *IncomeStatementItemExpenses `json:"expenses,omitempty"`
	InterestIncomeAndExpense *IncomeStatementItemInterestIncomeAndExpense `json:"interest_income_and_expense,omitempty"`
	OtherIncomeAndExpenses *IncomeStatementItemOtherIncomeAndExpenses `json:"other_income_and_expenses,omitempty"`
	Taxes *IncomeStatementItemTaxes `json:"taxes,omitempty"`
	DepreciationAndAmortization *IncomeStatementItemDepreciationAndAmortization `json:"depreciation_and_amortization,omitempty"`
	Ebitda *IncomeStatementItemEbitda `json:"ebitda,omitempty"`
	DividendsAndShares *IncomeStatementItemDividendsAndShares `json:"dividends_and_shares,omitempty"`
	UnusualItems *IncomeStatementItemUnusualItems `json:"unusual_items,omitempty"`
	Depreciation *IncomeStatementItemDepreciation `json:"depreciation,omitempty"`
	PretaxIncome *IncomeStatementItemPretaxIncome `json:"pretax_income,omitempty"`
	SpecialIncomeCharges *IncomeStatementItemSpecialIncomeCharges `json:"special_income_charges,omitempty"`
}

// NewIncomeStatementItem instantiates a new IncomeStatementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItem() *IncomeStatementItem {
	this := IncomeStatementItem{}
	return &this
}

// NewIncomeStatementItemWithDefaults instantiates a new IncomeStatementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemWithDefaults() *IncomeStatementItem {
	this := IncomeStatementItem{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetFiscalDate() string {
	if o == nil || IsNil(o.FiscalDate) {
		var ret string
		return ret
	}
	return *o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetFiscalDateOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalDate) {
		return nil, false
	}
	return o.FiscalDate, true
}

// HasFiscalDate returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasFiscalDate() bool {
	if o != nil && !IsNil(o.FiscalDate) {
		return true
	}

	return false
}

// SetFiscalDate gets a reference to the given string and assigns it to the FiscalDate field.
func (o *IncomeStatementItem) SetFiscalDate(v string) {
	o.FiscalDate = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *IncomeStatementItem) SetYear(v int64) {
	o.Year = &v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetRevenue() IncomeStatementItemRevenue {
	if o == nil || IsNil(o.Revenue) {
		var ret IncomeStatementItemRevenue
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetRevenueOk() (*IncomeStatementItemRevenue, bool) {
	if o == nil || IsNil(o.Revenue) {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasRevenue() bool {
	if o != nil && !IsNil(o.Revenue) {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given IncomeStatementItemRevenue and assigns it to the Revenue field.
func (o *IncomeStatementItem) SetRevenue(v IncomeStatementItemRevenue) {
	o.Revenue = &v
}

// GetGrossProfit returns the GrossProfit field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetGrossProfit() IncomeStatementItemGrossProfit {
	if o == nil || IsNil(o.GrossProfit) {
		var ret IncomeStatementItemGrossProfit
		return ret
	}
	return *o.GrossProfit
}

// GetGrossProfitOk returns a tuple with the GrossProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetGrossProfitOk() (*IncomeStatementItemGrossProfit, bool) {
	if o == nil || IsNil(o.GrossProfit) {
		return nil, false
	}
	return o.GrossProfit, true
}

// HasGrossProfit returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasGrossProfit() bool {
	if o != nil && !IsNil(o.GrossProfit) {
		return true
	}

	return false
}

// SetGrossProfit gets a reference to the given IncomeStatementItemGrossProfit and assigns it to the GrossProfit field.
func (o *IncomeStatementItem) SetGrossProfit(v IncomeStatementItemGrossProfit) {
	o.GrossProfit = &v
}

// GetOperatingIncome returns the OperatingIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetOperatingIncome() IncomeStatementItemOperatingIncome {
	if o == nil || IsNil(o.OperatingIncome) {
		var ret IncomeStatementItemOperatingIncome
		return ret
	}
	return *o.OperatingIncome
}

// GetOperatingIncomeOk returns a tuple with the OperatingIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetOperatingIncomeOk() (*IncomeStatementItemOperatingIncome, bool) {
	if o == nil || IsNil(o.OperatingIncome) {
		return nil, false
	}
	return o.OperatingIncome, true
}

// HasOperatingIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasOperatingIncome() bool {
	if o != nil && !IsNil(o.OperatingIncome) {
		return true
	}

	return false
}

// SetOperatingIncome gets a reference to the given IncomeStatementItemOperatingIncome and assigns it to the OperatingIncome field.
func (o *IncomeStatementItem) SetOperatingIncome(v IncomeStatementItemOperatingIncome) {
	o.OperatingIncome = &v
}

// GetNetIncome returns the NetIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetNetIncome() IncomeStatementItemNetIncome {
	if o == nil || IsNil(o.NetIncome) {
		var ret IncomeStatementItemNetIncome
		return ret
	}
	return *o.NetIncome
}

// GetNetIncomeOk returns a tuple with the NetIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetNetIncomeOk() (*IncomeStatementItemNetIncome, bool) {
	if o == nil || IsNil(o.NetIncome) {
		return nil, false
	}
	return o.NetIncome, true
}

// HasNetIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasNetIncome() bool {
	if o != nil && !IsNil(o.NetIncome) {
		return true
	}

	return false
}

// SetNetIncome gets a reference to the given IncomeStatementItemNetIncome and assigns it to the NetIncome field.
func (o *IncomeStatementItem) SetNetIncome(v IncomeStatementItemNetIncome) {
	o.NetIncome = &v
}

// GetEarningsPerShare returns the EarningsPerShare field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetEarningsPerShare() IncomeStatementItemEarningsPerShare {
	if o == nil || IsNil(o.EarningsPerShare) {
		var ret IncomeStatementItemEarningsPerShare
		return ret
	}
	return *o.EarningsPerShare
}

// GetEarningsPerShareOk returns a tuple with the EarningsPerShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetEarningsPerShareOk() (*IncomeStatementItemEarningsPerShare, bool) {
	if o == nil || IsNil(o.EarningsPerShare) {
		return nil, false
	}
	return o.EarningsPerShare, true
}

// HasEarningsPerShare returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasEarningsPerShare() bool {
	if o != nil && !IsNil(o.EarningsPerShare) {
		return true
	}

	return false
}

// SetEarningsPerShare gets a reference to the given IncomeStatementItemEarningsPerShare and assigns it to the EarningsPerShare field.
func (o *IncomeStatementItem) SetEarningsPerShare(v IncomeStatementItemEarningsPerShare) {
	o.EarningsPerShare = &v
}

// GetExpenses returns the Expenses field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetExpenses() IncomeStatementItemExpenses {
	if o == nil || IsNil(o.Expenses) {
		var ret IncomeStatementItemExpenses
		return ret
	}
	return *o.Expenses
}

// GetExpensesOk returns a tuple with the Expenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetExpensesOk() (*IncomeStatementItemExpenses, bool) {
	if o == nil || IsNil(o.Expenses) {
		return nil, false
	}
	return o.Expenses, true
}

// HasExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasExpenses() bool {
	if o != nil && !IsNil(o.Expenses) {
		return true
	}

	return false
}

// SetExpenses gets a reference to the given IncomeStatementItemExpenses and assigns it to the Expenses field.
func (o *IncomeStatementItem) SetExpenses(v IncomeStatementItemExpenses) {
	o.Expenses = &v
}

// GetInterestIncomeAndExpense returns the InterestIncomeAndExpense field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetInterestIncomeAndExpense() IncomeStatementItemInterestIncomeAndExpense {
	if o == nil || IsNil(o.InterestIncomeAndExpense) {
		var ret IncomeStatementItemInterestIncomeAndExpense
		return ret
	}
	return *o.InterestIncomeAndExpense
}

// GetInterestIncomeAndExpenseOk returns a tuple with the InterestIncomeAndExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetInterestIncomeAndExpenseOk() (*IncomeStatementItemInterestIncomeAndExpense, bool) {
	if o == nil || IsNil(o.InterestIncomeAndExpense) {
		return nil, false
	}
	return o.InterestIncomeAndExpense, true
}

// HasInterestIncomeAndExpense returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasInterestIncomeAndExpense() bool {
	if o != nil && !IsNil(o.InterestIncomeAndExpense) {
		return true
	}

	return false
}

// SetInterestIncomeAndExpense gets a reference to the given IncomeStatementItemInterestIncomeAndExpense and assigns it to the InterestIncomeAndExpense field.
func (o *IncomeStatementItem) SetInterestIncomeAndExpense(v IncomeStatementItemInterestIncomeAndExpense) {
	o.InterestIncomeAndExpense = &v
}

// GetOtherIncomeAndExpenses returns the OtherIncomeAndExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetOtherIncomeAndExpenses() IncomeStatementItemOtherIncomeAndExpenses {
	if o == nil || IsNil(o.OtherIncomeAndExpenses) {
		var ret IncomeStatementItemOtherIncomeAndExpenses
		return ret
	}
	return *o.OtherIncomeAndExpenses
}

// GetOtherIncomeAndExpensesOk returns a tuple with the OtherIncomeAndExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetOtherIncomeAndExpensesOk() (*IncomeStatementItemOtherIncomeAndExpenses, bool) {
	if o == nil || IsNil(o.OtherIncomeAndExpenses) {
		return nil, false
	}
	return o.OtherIncomeAndExpenses, true
}

// HasOtherIncomeAndExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasOtherIncomeAndExpenses() bool {
	if o != nil && !IsNil(o.OtherIncomeAndExpenses) {
		return true
	}

	return false
}

// SetOtherIncomeAndExpenses gets a reference to the given IncomeStatementItemOtherIncomeAndExpenses and assigns it to the OtherIncomeAndExpenses field.
func (o *IncomeStatementItem) SetOtherIncomeAndExpenses(v IncomeStatementItemOtherIncomeAndExpenses) {
	o.OtherIncomeAndExpenses = &v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetTaxes() IncomeStatementItemTaxes {
	if o == nil || IsNil(o.Taxes) {
		var ret IncomeStatementItemTaxes
		return ret
	}
	return *o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetTaxesOk() (*IncomeStatementItemTaxes, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given IncomeStatementItemTaxes and assigns it to the Taxes field.
func (o *IncomeStatementItem) SetTaxes(v IncomeStatementItemTaxes) {
	o.Taxes = &v
}

// GetDepreciationAndAmortization returns the DepreciationAndAmortization field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDepreciationAndAmortization() IncomeStatementItemDepreciationAndAmortization {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		var ret IncomeStatementItemDepreciationAndAmortization
		return ret
	}
	return *o.DepreciationAndAmortization
}

// GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDepreciationAndAmortizationOk() (*IncomeStatementItemDepreciationAndAmortization, bool) {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		return nil, false
	}
	return o.DepreciationAndAmortization, true
}

// HasDepreciationAndAmortization returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDepreciationAndAmortization() bool {
	if o != nil && !IsNil(o.DepreciationAndAmortization) {
		return true
	}

	return false
}

// SetDepreciationAndAmortization gets a reference to the given IncomeStatementItemDepreciationAndAmortization and assigns it to the DepreciationAndAmortization field.
func (o *IncomeStatementItem) SetDepreciationAndAmortization(v IncomeStatementItemDepreciationAndAmortization) {
	o.DepreciationAndAmortization = &v
}

// GetEbitda returns the Ebitda field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetEbitda() IncomeStatementItemEbitda {
	if o == nil || IsNil(o.Ebitda) {
		var ret IncomeStatementItemEbitda
		return ret
	}
	return *o.Ebitda
}

// GetEbitdaOk returns a tuple with the Ebitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetEbitdaOk() (*IncomeStatementItemEbitda, bool) {
	if o == nil || IsNil(o.Ebitda) {
		return nil, false
	}
	return o.Ebitda, true
}

// HasEbitda returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasEbitda() bool {
	if o != nil && !IsNil(o.Ebitda) {
		return true
	}

	return false
}

// SetEbitda gets a reference to the given IncomeStatementItemEbitda and assigns it to the Ebitda field.
func (o *IncomeStatementItem) SetEbitda(v IncomeStatementItemEbitda) {
	o.Ebitda = &v
}

// GetDividendsAndShares returns the DividendsAndShares field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDividendsAndShares() IncomeStatementItemDividendsAndShares {
	if o == nil || IsNil(o.DividendsAndShares) {
		var ret IncomeStatementItemDividendsAndShares
		return ret
	}
	return *o.DividendsAndShares
}

// GetDividendsAndSharesOk returns a tuple with the DividendsAndShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDividendsAndSharesOk() (*IncomeStatementItemDividendsAndShares, bool) {
	if o == nil || IsNil(o.DividendsAndShares) {
		return nil, false
	}
	return o.DividendsAndShares, true
}

// HasDividendsAndShares returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDividendsAndShares() bool {
	if o != nil && !IsNil(o.DividendsAndShares) {
		return true
	}

	return false
}

// SetDividendsAndShares gets a reference to the given IncomeStatementItemDividendsAndShares and assigns it to the DividendsAndShares field.
func (o *IncomeStatementItem) SetDividendsAndShares(v IncomeStatementItemDividendsAndShares) {
	o.DividendsAndShares = &v
}

// GetUnusualItems returns the UnusualItems field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetUnusualItems() IncomeStatementItemUnusualItems {
	if o == nil || IsNil(o.UnusualItems) {
		var ret IncomeStatementItemUnusualItems
		return ret
	}
	return *o.UnusualItems
}

// GetUnusualItemsOk returns a tuple with the UnusualItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetUnusualItemsOk() (*IncomeStatementItemUnusualItems, bool) {
	if o == nil || IsNil(o.UnusualItems) {
		return nil, false
	}
	return o.UnusualItems, true
}

// HasUnusualItems returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasUnusualItems() bool {
	if o != nil && !IsNil(o.UnusualItems) {
		return true
	}

	return false
}

// SetUnusualItems gets a reference to the given IncomeStatementItemUnusualItems and assigns it to the UnusualItems field.
func (o *IncomeStatementItem) SetUnusualItems(v IncomeStatementItemUnusualItems) {
	o.UnusualItems = &v
}

// GetDepreciation returns the Depreciation field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDepreciation() IncomeStatementItemDepreciation {
	if o == nil || IsNil(o.Depreciation) {
		var ret IncomeStatementItemDepreciation
		return ret
	}
	return *o.Depreciation
}

// GetDepreciationOk returns a tuple with the Depreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDepreciationOk() (*IncomeStatementItemDepreciation, bool) {
	if o == nil || IsNil(o.Depreciation) {
		return nil, false
	}
	return o.Depreciation, true
}

// HasDepreciation returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDepreciation() bool {
	if o != nil && !IsNil(o.Depreciation) {
		return true
	}

	return false
}

// SetDepreciation gets a reference to the given IncomeStatementItemDepreciation and assigns it to the Depreciation field.
func (o *IncomeStatementItem) SetDepreciation(v IncomeStatementItemDepreciation) {
	o.Depreciation = &v
}

// GetPretaxIncome returns the PretaxIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetPretaxIncome() IncomeStatementItemPretaxIncome {
	if o == nil || IsNil(o.PretaxIncome) {
		var ret IncomeStatementItemPretaxIncome
		return ret
	}
	return *o.PretaxIncome
}

// GetPretaxIncomeOk returns a tuple with the PretaxIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetPretaxIncomeOk() (*IncomeStatementItemPretaxIncome, bool) {
	if o == nil || IsNil(o.PretaxIncome) {
		return nil, false
	}
	return o.PretaxIncome, true
}

// HasPretaxIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasPretaxIncome() bool {
	if o != nil && !IsNil(o.PretaxIncome) {
		return true
	}

	return false
}

// SetPretaxIncome gets a reference to the given IncomeStatementItemPretaxIncome and assigns it to the PretaxIncome field.
func (o *IncomeStatementItem) SetPretaxIncome(v IncomeStatementItemPretaxIncome) {
	o.PretaxIncome = &v
}

// GetSpecialIncomeCharges returns the SpecialIncomeCharges field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetSpecialIncomeCharges() IncomeStatementItemSpecialIncomeCharges {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		var ret IncomeStatementItemSpecialIncomeCharges
		return ret
	}
	return *o.SpecialIncomeCharges
}

// GetSpecialIncomeChargesOk returns a tuple with the SpecialIncomeCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetSpecialIncomeChargesOk() (*IncomeStatementItemSpecialIncomeCharges, bool) {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		return nil, false
	}
	return o.SpecialIncomeCharges, true
}

// HasSpecialIncomeCharges returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasSpecialIncomeCharges() bool {
	if o != nil && !IsNil(o.SpecialIncomeCharges) {
		return true
	}

	return false
}

// SetSpecialIncomeCharges gets a reference to the given IncomeStatementItemSpecialIncomeCharges and assigns it to the SpecialIncomeCharges field.
func (o *IncomeStatementItem) SetSpecialIncomeCharges(v IncomeStatementItemSpecialIncomeCharges) {
	o.SpecialIncomeCharges = &v
}

func (o IncomeStatementItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiscalDate) {
		toSerialize["fiscal_date"] = o.FiscalDate
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Revenue) {
		toSerialize["revenue"] = o.Revenue
	}
	if !IsNil(o.GrossProfit) {
		toSerialize["gross_profit"] = o.GrossProfit
	}
	if !IsNil(o.OperatingIncome) {
		toSerialize["operating_income"] = o.OperatingIncome
	}
	if !IsNil(o.NetIncome) {
		toSerialize["net_income"] = o.NetIncome
	}
	if !IsNil(o.EarningsPerShare) {
		toSerialize["earnings_per_share"] = o.EarningsPerShare
	}
	if !IsNil(o.Expenses) {
		toSerialize["expenses"] = o.Expenses
	}
	if !IsNil(o.InterestIncomeAndExpense) {
		toSerialize["interest_income_and_expense"] = o.InterestIncomeAndExpense
	}
	if !IsNil(o.OtherIncomeAndExpenses) {
		toSerialize["other_income_and_expenses"] = o.OtherIncomeAndExpenses
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !IsNil(o.DepreciationAndAmortization) {
		toSerialize["depreciation_and_amortization"] = o.DepreciationAndAmortization
	}
	if !IsNil(o.Ebitda) {
		toSerialize["ebitda"] = o.Ebitda
	}
	if !IsNil(o.DividendsAndShares) {
		toSerialize["dividends_and_shares"] = o.DividendsAndShares
	}
	if !IsNil(o.UnusualItems) {
		toSerialize["unusual_items"] = o.UnusualItems
	}
	if !IsNil(o.Depreciation) {
		toSerialize["depreciation"] = o.Depreciation
	}
	if !IsNil(o.PretaxIncome) {
		toSerialize["pretax_income"] = o.PretaxIncome
	}
	if !IsNil(o.SpecialIncomeCharges) {
		toSerialize["special_income_charges"] = o.SpecialIncomeCharges
	}
	return toSerialize, nil
}

type NullableIncomeStatementItem struct {
	value *IncomeStatementItem
	isSet bool
}

func (v NullableIncomeStatementItem) Get() *IncomeStatementItem {
	return v.value
}

func (v *NullableIncomeStatementItem) Set(val *IncomeStatementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItem(val *IncomeStatementItem) *NullableIncomeStatementItem {
	return &NullableIncomeStatementItem{value: val, isSet: true}
}

func (v NullableIncomeStatementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


