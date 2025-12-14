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

// checks if the CashFlowDataCashFlowFromFinancingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataCashFlowFromFinancingActivities{}

// CashFlowDataCashFlowFromFinancingActivities Cash flow from financing activities
type CashFlowDataCashFlowFromFinancingActivities struct {
	// Financing cash flow
	FinancingCashFlow *float64 `json:"financing_cash_flow,omitempty"`
	// Cash flow from continuing financing activities
	CashFlowFromContinuingFinancingActivities *float64 `json:"cash_flow_from_continuing_financing_activities,omitempty"`
	// Cash from discontinued financing activities
	CashFromDiscontinuedFinancingActivities *float64 `json:"cash_from_discontinued_financing_activities,omitempty"`
	// Net other financing charges
	NetOtherFinancingCharges *float64 `json:"net_other_financing_charges,omitempty"`
	// Interest paid cff
	InterestPaidCff *float64 `json:"interest_paid_cff,omitempty"`
	// Proceeds from stock option exercised
	ProceedsFromStockOptionExercised *float64 `json:"proceeds_from_stock_option_exercised,omitempty"`
	// Cash dividends paid
	CashDividendsPaid *float64 `json:"cash_dividends_paid,omitempty"`
	// Preferred stock dividend paid
	PreferredStockDividendPaid *float64 `json:"preferred_stock_dividend_paid,omitempty"`
	// Common stock dividend paid
	CommonStockDividendPaid *float64 `json:"common_stock_dividend_paid,omitempty"`
	// Net preferred stock issuance
	NetPreferredStockIssuance *float64 `json:"net_preferred_stock_issuance,omitempty"`
	// Preferred stock payments
	PreferredStockPayments *float64 `json:"preferred_stock_payments,omitempty"`
	// Preferred stock issuance
	PreferredStockIssuance *float64 `json:"preferred_stock_issuance,omitempty"`
	// Net common stock issuance
	NetCommonStockIssuance *float64 `json:"net_common_stock_issuance,omitempty"`
	// Common stock payments
	CommonStockPayments *float64 `json:"common_stock_payments,omitempty"`
	// Common stock issuance
	CommonStockIssuance *float64 `json:"common_stock_issuance,omitempty"`
	// Repurchase of capital stock
	RepurchaseOfCapitalStock *float64 `json:"repurchase_of_capital_stock,omitempty"`
	// Net issuance payments of debt
	NetIssuancePaymentsOfDebt *float64 `json:"net_issuance_payments_of_debt,omitempty"`
	// Net short term debt issuance
	NetShortTermDebtIssuance *float64 `json:"net_short_term_debt_issuance,omitempty"`
	// Short term debt payments
	ShortTermDebtPayments *float64 `json:"short_term_debt_payments,omitempty"`
	// Short term debt issuance
	ShortTermDebtIssuance *float64 `json:"short_term_debt_issuance,omitempty"`
	// Net long term debt issuance
	NetLongTermDebtIssuance *float64 `json:"net_long_term_debt_issuance,omitempty"`
	// Long term debt payments
	LongTermDebtPayments *float64 `json:"long_term_debt_payments,omitempty"`
	// Long term debt issuance
	LongTermDebtIssuance *float64 `json:"long_term_debt_issuance,omitempty"`
	// Issuance of debt
	IssuanceOfDebt *float64 `json:"issuance_of_debt,omitempty"`
	// Repayment of debt
	RepaymentOfDebt *float64 `json:"repayment_of_debt,omitempty"`
	// Issuance of capital stock
	IssuanceOfCapitalStock *float64 `json:"issuance_of_capital_stock,omitempty"`
}

// NewCashFlowDataCashFlowFromFinancingActivities instantiates a new CashFlowDataCashFlowFromFinancingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataCashFlowFromFinancingActivities() *CashFlowDataCashFlowFromFinancingActivities {
	this := CashFlowDataCashFlowFromFinancingActivities{}
	return &this
}

// NewCashFlowDataCashFlowFromFinancingActivitiesWithDefaults instantiates a new CashFlowDataCashFlowFromFinancingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataCashFlowFromFinancingActivitiesWithDefaults() *CashFlowDataCashFlowFromFinancingActivities {
	this := CashFlowDataCashFlowFromFinancingActivities{}
	return &this
}

// GetFinancingCashFlow returns the FinancingCashFlow field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetFinancingCashFlow() float64 {
	if o == nil || IsNil(o.FinancingCashFlow) {
		var ret float64
		return ret
	}
	return *o.FinancingCashFlow
}

// GetFinancingCashFlowOk returns a tuple with the FinancingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetFinancingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancingCashFlow) {
		return nil, false
	}
	return o.FinancingCashFlow, true
}

// HasFinancingCashFlow returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasFinancingCashFlow() bool {
	if o != nil && !IsNil(o.FinancingCashFlow) {
		return true
	}

	return false
}

// SetFinancingCashFlow gets a reference to the given float64 and assigns it to the FinancingCashFlow field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetFinancingCashFlow(v float64) {
	o.FinancingCashFlow = &v
}

// GetCashFlowFromContinuingFinancingActivities returns the CashFlowFromContinuingFinancingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashFlowFromContinuingFinancingActivities() float64 {
	if o == nil || IsNil(o.CashFlowFromContinuingFinancingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFlowFromContinuingFinancingActivities
}

// GetCashFlowFromContinuingFinancingActivitiesOk returns a tuple with the CashFlowFromContinuingFinancingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashFlowFromContinuingFinancingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFlowFromContinuingFinancingActivities) {
		return nil, false
	}
	return o.CashFlowFromContinuingFinancingActivities, true
}

// HasCashFlowFromContinuingFinancingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCashFlowFromContinuingFinancingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromContinuingFinancingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromContinuingFinancingActivities gets a reference to the given float64 and assigns it to the CashFlowFromContinuingFinancingActivities field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCashFlowFromContinuingFinancingActivities(v float64) {
	o.CashFlowFromContinuingFinancingActivities = &v
}

// GetCashFromDiscontinuedFinancingActivities returns the CashFromDiscontinuedFinancingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashFromDiscontinuedFinancingActivities() float64 {
	if o == nil || IsNil(o.CashFromDiscontinuedFinancingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFromDiscontinuedFinancingActivities
}

// GetCashFromDiscontinuedFinancingActivitiesOk returns a tuple with the CashFromDiscontinuedFinancingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashFromDiscontinuedFinancingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFromDiscontinuedFinancingActivities) {
		return nil, false
	}
	return o.CashFromDiscontinuedFinancingActivities, true
}

// HasCashFromDiscontinuedFinancingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCashFromDiscontinuedFinancingActivities() bool {
	if o != nil && !IsNil(o.CashFromDiscontinuedFinancingActivities) {
		return true
	}

	return false
}

// SetCashFromDiscontinuedFinancingActivities gets a reference to the given float64 and assigns it to the CashFromDiscontinuedFinancingActivities field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCashFromDiscontinuedFinancingActivities(v float64) {
	o.CashFromDiscontinuedFinancingActivities = &v
}

// GetNetOtherFinancingCharges returns the NetOtherFinancingCharges field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetOtherFinancingCharges() float64 {
	if o == nil || IsNil(o.NetOtherFinancingCharges) {
		var ret float64
		return ret
	}
	return *o.NetOtherFinancingCharges
}

// GetNetOtherFinancingChargesOk returns a tuple with the NetOtherFinancingCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetOtherFinancingChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.NetOtherFinancingCharges) {
		return nil, false
	}
	return o.NetOtherFinancingCharges, true
}

// HasNetOtherFinancingCharges returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetOtherFinancingCharges() bool {
	if o != nil && !IsNil(o.NetOtherFinancingCharges) {
		return true
	}

	return false
}

// SetNetOtherFinancingCharges gets a reference to the given float64 and assigns it to the NetOtherFinancingCharges field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetOtherFinancingCharges(v float64) {
	o.NetOtherFinancingCharges = &v
}

// GetInterestPaidCff returns the InterestPaidCff field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetInterestPaidCff() float64 {
	if o == nil || IsNil(o.InterestPaidCff) {
		var ret float64
		return ret
	}
	return *o.InterestPaidCff
}

// GetInterestPaidCffOk returns a tuple with the InterestPaidCff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetInterestPaidCffOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPaidCff) {
		return nil, false
	}
	return o.InterestPaidCff, true
}

// HasInterestPaidCff returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasInterestPaidCff() bool {
	if o != nil && !IsNil(o.InterestPaidCff) {
		return true
	}

	return false
}

// SetInterestPaidCff gets a reference to the given float64 and assigns it to the InterestPaidCff field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetInterestPaidCff(v float64) {
	o.InterestPaidCff = &v
}

// GetProceedsFromStockOptionExercised returns the ProceedsFromStockOptionExercised field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetProceedsFromStockOptionExercised() float64 {
	if o == nil || IsNil(o.ProceedsFromStockOptionExercised) {
		var ret float64
		return ret
	}
	return *o.ProceedsFromStockOptionExercised
}

// GetProceedsFromStockOptionExercisedOk returns a tuple with the ProceedsFromStockOptionExercised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetProceedsFromStockOptionExercisedOk() (*float64, bool) {
	if o == nil || IsNil(o.ProceedsFromStockOptionExercised) {
		return nil, false
	}
	return o.ProceedsFromStockOptionExercised, true
}

// HasProceedsFromStockOptionExercised returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasProceedsFromStockOptionExercised() bool {
	if o != nil && !IsNil(o.ProceedsFromStockOptionExercised) {
		return true
	}

	return false
}

// SetProceedsFromStockOptionExercised gets a reference to the given float64 and assigns it to the ProceedsFromStockOptionExercised field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetProceedsFromStockOptionExercised(v float64) {
	o.ProceedsFromStockOptionExercised = &v
}

// GetCashDividendsPaid returns the CashDividendsPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashDividendsPaid() float64 {
	if o == nil || IsNil(o.CashDividendsPaid) {
		var ret float64
		return ret
	}
	return *o.CashDividendsPaid
}

// GetCashDividendsPaidOk returns a tuple with the CashDividendsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCashDividendsPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.CashDividendsPaid) {
		return nil, false
	}
	return o.CashDividendsPaid, true
}

// HasCashDividendsPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCashDividendsPaid() bool {
	if o != nil && !IsNil(o.CashDividendsPaid) {
		return true
	}

	return false
}

// SetCashDividendsPaid gets a reference to the given float64 and assigns it to the CashDividendsPaid field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCashDividendsPaid(v float64) {
	o.CashDividendsPaid = &v
}

// GetPreferredStockDividendPaid returns the PreferredStockDividendPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockDividendPaid() float64 {
	if o == nil || IsNil(o.PreferredStockDividendPaid) {
		var ret float64
		return ret
	}
	return *o.PreferredStockDividendPaid
}

// GetPreferredStockDividendPaidOk returns a tuple with the PreferredStockDividendPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockDividendPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockDividendPaid) {
		return nil, false
	}
	return o.PreferredStockDividendPaid, true
}

// HasPreferredStockDividendPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasPreferredStockDividendPaid() bool {
	if o != nil && !IsNil(o.PreferredStockDividendPaid) {
		return true
	}

	return false
}

// SetPreferredStockDividendPaid gets a reference to the given float64 and assigns it to the PreferredStockDividendPaid field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetPreferredStockDividendPaid(v float64) {
	o.PreferredStockDividendPaid = &v
}

// GetCommonStockDividendPaid returns the CommonStockDividendPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockDividendPaid() float64 {
	if o == nil || IsNil(o.CommonStockDividendPaid) {
		var ret float64
		return ret
	}
	return *o.CommonStockDividendPaid
}

// GetCommonStockDividendPaidOk returns a tuple with the CommonStockDividendPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockDividendPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockDividendPaid) {
		return nil, false
	}
	return o.CommonStockDividendPaid, true
}

// HasCommonStockDividendPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCommonStockDividendPaid() bool {
	if o != nil && !IsNil(o.CommonStockDividendPaid) {
		return true
	}

	return false
}

// SetCommonStockDividendPaid gets a reference to the given float64 and assigns it to the CommonStockDividendPaid field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCommonStockDividendPaid(v float64) {
	o.CommonStockDividendPaid = &v
}

// GetNetPreferredStockIssuance returns the NetPreferredStockIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetPreferredStockIssuance() float64 {
	if o == nil || IsNil(o.NetPreferredStockIssuance) {
		var ret float64
		return ret
	}
	return *o.NetPreferredStockIssuance
}

// GetNetPreferredStockIssuanceOk returns a tuple with the NetPreferredStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetPreferredStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.NetPreferredStockIssuance) {
		return nil, false
	}
	return o.NetPreferredStockIssuance, true
}

// HasNetPreferredStockIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetPreferredStockIssuance() bool {
	if o != nil && !IsNil(o.NetPreferredStockIssuance) {
		return true
	}

	return false
}

// SetNetPreferredStockIssuance gets a reference to the given float64 and assigns it to the NetPreferredStockIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetPreferredStockIssuance(v float64) {
	o.NetPreferredStockIssuance = &v
}

// GetPreferredStockPayments returns the PreferredStockPayments field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockPayments() float64 {
	if o == nil || IsNil(o.PreferredStockPayments) {
		var ret float64
		return ret
	}
	return *o.PreferredStockPayments
}

// GetPreferredStockPaymentsOk returns a tuple with the PreferredStockPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockPayments) {
		return nil, false
	}
	return o.PreferredStockPayments, true
}

// HasPreferredStockPayments returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasPreferredStockPayments() bool {
	if o != nil && !IsNil(o.PreferredStockPayments) {
		return true
	}

	return false
}

// SetPreferredStockPayments gets a reference to the given float64 and assigns it to the PreferredStockPayments field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetPreferredStockPayments(v float64) {
	o.PreferredStockPayments = &v
}

// GetPreferredStockIssuance returns the PreferredStockIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockIssuance() float64 {
	if o == nil || IsNil(o.PreferredStockIssuance) {
		var ret float64
		return ret
	}
	return *o.PreferredStockIssuance
}

// GetPreferredStockIssuanceOk returns a tuple with the PreferredStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetPreferredStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockIssuance) {
		return nil, false
	}
	return o.PreferredStockIssuance, true
}

// HasPreferredStockIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasPreferredStockIssuance() bool {
	if o != nil && !IsNil(o.PreferredStockIssuance) {
		return true
	}

	return false
}

// SetPreferredStockIssuance gets a reference to the given float64 and assigns it to the PreferredStockIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetPreferredStockIssuance(v float64) {
	o.PreferredStockIssuance = &v
}

// GetNetCommonStockIssuance returns the NetCommonStockIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetCommonStockIssuance() float64 {
	if o == nil || IsNil(o.NetCommonStockIssuance) {
		var ret float64
		return ret
	}
	return *o.NetCommonStockIssuance
}

// GetNetCommonStockIssuanceOk returns a tuple with the NetCommonStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetCommonStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.NetCommonStockIssuance) {
		return nil, false
	}
	return o.NetCommonStockIssuance, true
}

// HasNetCommonStockIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetCommonStockIssuance() bool {
	if o != nil && !IsNil(o.NetCommonStockIssuance) {
		return true
	}

	return false
}

// SetNetCommonStockIssuance gets a reference to the given float64 and assigns it to the NetCommonStockIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetCommonStockIssuance(v float64) {
	o.NetCommonStockIssuance = &v
}

// GetCommonStockPayments returns the CommonStockPayments field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockPayments() float64 {
	if o == nil || IsNil(o.CommonStockPayments) {
		var ret float64
		return ret
	}
	return *o.CommonStockPayments
}

// GetCommonStockPaymentsOk returns a tuple with the CommonStockPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockPayments) {
		return nil, false
	}
	return o.CommonStockPayments, true
}

// HasCommonStockPayments returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCommonStockPayments() bool {
	if o != nil && !IsNil(o.CommonStockPayments) {
		return true
	}

	return false
}

// SetCommonStockPayments gets a reference to the given float64 and assigns it to the CommonStockPayments field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCommonStockPayments(v float64) {
	o.CommonStockPayments = &v
}

// GetCommonStockIssuance returns the CommonStockIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockIssuance() float64 {
	if o == nil || IsNil(o.CommonStockIssuance) {
		var ret float64
		return ret
	}
	return *o.CommonStockIssuance
}

// GetCommonStockIssuanceOk returns a tuple with the CommonStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetCommonStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockIssuance) {
		return nil, false
	}
	return o.CommonStockIssuance, true
}

// HasCommonStockIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasCommonStockIssuance() bool {
	if o != nil && !IsNil(o.CommonStockIssuance) {
		return true
	}

	return false
}

// SetCommonStockIssuance gets a reference to the given float64 and assigns it to the CommonStockIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetCommonStockIssuance(v float64) {
	o.CommonStockIssuance = &v
}

// GetRepurchaseOfCapitalStock returns the RepurchaseOfCapitalStock field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetRepurchaseOfCapitalStock() float64 {
	if o == nil || IsNil(o.RepurchaseOfCapitalStock) {
		var ret float64
		return ret
	}
	return *o.RepurchaseOfCapitalStock
}

// GetRepurchaseOfCapitalStockOk returns a tuple with the RepurchaseOfCapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetRepurchaseOfCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.RepurchaseOfCapitalStock) {
		return nil, false
	}
	return o.RepurchaseOfCapitalStock, true
}

// HasRepurchaseOfCapitalStock returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasRepurchaseOfCapitalStock() bool {
	if o != nil && !IsNil(o.RepurchaseOfCapitalStock) {
		return true
	}

	return false
}

// SetRepurchaseOfCapitalStock gets a reference to the given float64 and assigns it to the RepurchaseOfCapitalStock field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetRepurchaseOfCapitalStock(v float64) {
	o.RepurchaseOfCapitalStock = &v
}

// GetNetIssuancePaymentsOfDebt returns the NetIssuancePaymentsOfDebt field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetIssuancePaymentsOfDebt() float64 {
	if o == nil || IsNil(o.NetIssuancePaymentsOfDebt) {
		var ret float64
		return ret
	}
	return *o.NetIssuancePaymentsOfDebt
}

// GetNetIssuancePaymentsOfDebtOk returns a tuple with the NetIssuancePaymentsOfDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetIssuancePaymentsOfDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIssuancePaymentsOfDebt) {
		return nil, false
	}
	return o.NetIssuancePaymentsOfDebt, true
}

// HasNetIssuancePaymentsOfDebt returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetIssuancePaymentsOfDebt() bool {
	if o != nil && !IsNil(o.NetIssuancePaymentsOfDebt) {
		return true
	}

	return false
}

// SetNetIssuancePaymentsOfDebt gets a reference to the given float64 and assigns it to the NetIssuancePaymentsOfDebt field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetIssuancePaymentsOfDebt(v float64) {
	o.NetIssuancePaymentsOfDebt = &v
}

// GetNetShortTermDebtIssuance returns the NetShortTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetShortTermDebtIssuance() float64 {
	if o == nil || IsNil(o.NetShortTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.NetShortTermDebtIssuance
}

// GetNetShortTermDebtIssuanceOk returns a tuple with the NetShortTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetShortTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.NetShortTermDebtIssuance) {
		return nil, false
	}
	return o.NetShortTermDebtIssuance, true
}

// HasNetShortTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetShortTermDebtIssuance() bool {
	if o != nil && !IsNil(o.NetShortTermDebtIssuance) {
		return true
	}

	return false
}

// SetNetShortTermDebtIssuance gets a reference to the given float64 and assigns it to the NetShortTermDebtIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetShortTermDebtIssuance(v float64) {
	o.NetShortTermDebtIssuance = &v
}

// GetShortTermDebtPayments returns the ShortTermDebtPayments field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetShortTermDebtPayments() float64 {
	if o == nil || IsNil(o.ShortTermDebtPayments) {
		var ret float64
		return ret
	}
	return *o.ShortTermDebtPayments
}

// GetShortTermDebtPaymentsOk returns a tuple with the ShortTermDebtPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetShortTermDebtPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortTermDebtPayments) {
		return nil, false
	}
	return o.ShortTermDebtPayments, true
}

// HasShortTermDebtPayments returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasShortTermDebtPayments() bool {
	if o != nil && !IsNil(o.ShortTermDebtPayments) {
		return true
	}

	return false
}

// SetShortTermDebtPayments gets a reference to the given float64 and assigns it to the ShortTermDebtPayments field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetShortTermDebtPayments(v float64) {
	o.ShortTermDebtPayments = &v
}

// GetShortTermDebtIssuance returns the ShortTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetShortTermDebtIssuance() float64 {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.ShortTermDebtIssuance
}

// GetShortTermDebtIssuanceOk returns a tuple with the ShortTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetShortTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		return nil, false
	}
	return o.ShortTermDebtIssuance, true
}

// HasShortTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasShortTermDebtIssuance() bool {
	if o != nil && !IsNil(o.ShortTermDebtIssuance) {
		return true
	}

	return false
}

// SetShortTermDebtIssuance gets a reference to the given float64 and assigns it to the ShortTermDebtIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetShortTermDebtIssuance(v float64) {
	o.ShortTermDebtIssuance = &v
}

// GetNetLongTermDebtIssuance returns the NetLongTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetLongTermDebtIssuance() float64 {
	if o == nil || IsNil(o.NetLongTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.NetLongTermDebtIssuance
}

// GetNetLongTermDebtIssuanceOk returns a tuple with the NetLongTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetNetLongTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.NetLongTermDebtIssuance) {
		return nil, false
	}
	return o.NetLongTermDebtIssuance, true
}

// HasNetLongTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasNetLongTermDebtIssuance() bool {
	if o != nil && !IsNil(o.NetLongTermDebtIssuance) {
		return true
	}

	return false
}

// SetNetLongTermDebtIssuance gets a reference to the given float64 and assigns it to the NetLongTermDebtIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetNetLongTermDebtIssuance(v float64) {
	o.NetLongTermDebtIssuance = &v
}

// GetLongTermDebtPayments returns the LongTermDebtPayments field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetLongTermDebtPayments() float64 {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtPayments
}

// GetLongTermDebtPaymentsOk returns a tuple with the LongTermDebtPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetLongTermDebtPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		return nil, false
	}
	return o.LongTermDebtPayments, true
}

// HasLongTermDebtPayments returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasLongTermDebtPayments() bool {
	if o != nil && !IsNil(o.LongTermDebtPayments) {
		return true
	}

	return false
}

// SetLongTermDebtPayments gets a reference to the given float64 and assigns it to the LongTermDebtPayments field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetLongTermDebtPayments(v float64) {
	o.LongTermDebtPayments = &v
}

// GetLongTermDebtIssuance returns the LongTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetLongTermDebtIssuance() float64 {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtIssuance
}

// GetLongTermDebtIssuanceOk returns a tuple with the LongTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetLongTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		return nil, false
	}
	return o.LongTermDebtIssuance, true
}

// HasLongTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasLongTermDebtIssuance() bool {
	if o != nil && !IsNil(o.LongTermDebtIssuance) {
		return true
	}

	return false
}

// SetLongTermDebtIssuance gets a reference to the given float64 and assigns it to the LongTermDebtIssuance field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetLongTermDebtIssuance(v float64) {
	o.LongTermDebtIssuance = &v
}

// GetIssuanceOfDebt returns the IssuanceOfDebt field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetIssuanceOfDebt() float64 {
	if o == nil || IsNil(o.IssuanceOfDebt) {
		var ret float64
		return ret
	}
	return *o.IssuanceOfDebt
}

// GetIssuanceOfDebtOk returns a tuple with the IssuanceOfDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetIssuanceOfDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.IssuanceOfDebt) {
		return nil, false
	}
	return o.IssuanceOfDebt, true
}

// HasIssuanceOfDebt returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasIssuanceOfDebt() bool {
	if o != nil && !IsNil(o.IssuanceOfDebt) {
		return true
	}

	return false
}

// SetIssuanceOfDebt gets a reference to the given float64 and assigns it to the IssuanceOfDebt field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetIssuanceOfDebt(v float64) {
	o.IssuanceOfDebt = &v
}

// GetRepaymentOfDebt returns the RepaymentOfDebt field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetRepaymentOfDebt() float64 {
	if o == nil || IsNil(o.RepaymentOfDebt) {
		var ret float64
		return ret
	}
	return *o.RepaymentOfDebt
}

// GetRepaymentOfDebtOk returns a tuple with the RepaymentOfDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetRepaymentOfDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.RepaymentOfDebt) {
		return nil, false
	}
	return o.RepaymentOfDebt, true
}

// HasRepaymentOfDebt returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasRepaymentOfDebt() bool {
	if o != nil && !IsNil(o.RepaymentOfDebt) {
		return true
	}

	return false
}

// SetRepaymentOfDebt gets a reference to the given float64 and assigns it to the RepaymentOfDebt field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetRepaymentOfDebt(v float64) {
	o.RepaymentOfDebt = &v
}

// GetIssuanceOfCapitalStock returns the IssuanceOfCapitalStock field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetIssuanceOfCapitalStock() float64 {
	if o == nil || IsNil(o.IssuanceOfCapitalStock) {
		var ret float64
		return ret
	}
	return *o.IssuanceOfCapitalStock
}

// GetIssuanceOfCapitalStockOk returns a tuple with the IssuanceOfCapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) GetIssuanceOfCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.IssuanceOfCapitalStock) {
		return nil, false
	}
	return o.IssuanceOfCapitalStock, true
}

// HasIssuanceOfCapitalStock returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromFinancingActivities) HasIssuanceOfCapitalStock() bool {
	if o != nil && !IsNil(o.IssuanceOfCapitalStock) {
		return true
	}

	return false
}

// SetIssuanceOfCapitalStock gets a reference to the given float64 and assigns it to the IssuanceOfCapitalStock field.
func (o *CashFlowDataCashFlowFromFinancingActivities) SetIssuanceOfCapitalStock(v float64) {
	o.IssuanceOfCapitalStock = &v
}

func (o CashFlowDataCashFlowFromFinancingActivities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataCashFlowFromFinancingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FinancingCashFlow) {
		toSerialize["financing_cash_flow"] = o.FinancingCashFlow
	}
	if !IsNil(o.CashFlowFromContinuingFinancingActivities) {
		toSerialize["cash_flow_from_continuing_financing_activities"] = o.CashFlowFromContinuingFinancingActivities
	}
	if !IsNil(o.CashFromDiscontinuedFinancingActivities) {
		toSerialize["cash_from_discontinued_financing_activities"] = o.CashFromDiscontinuedFinancingActivities
	}
	if !IsNil(o.NetOtherFinancingCharges) {
		toSerialize["net_other_financing_charges"] = o.NetOtherFinancingCharges
	}
	if !IsNil(o.InterestPaidCff) {
		toSerialize["interest_paid_cff"] = o.InterestPaidCff
	}
	if !IsNil(o.ProceedsFromStockOptionExercised) {
		toSerialize["proceeds_from_stock_option_exercised"] = o.ProceedsFromStockOptionExercised
	}
	if !IsNil(o.CashDividendsPaid) {
		toSerialize["cash_dividends_paid"] = o.CashDividendsPaid
	}
	if !IsNil(o.PreferredStockDividendPaid) {
		toSerialize["preferred_stock_dividend_paid"] = o.PreferredStockDividendPaid
	}
	if !IsNil(o.CommonStockDividendPaid) {
		toSerialize["common_stock_dividend_paid"] = o.CommonStockDividendPaid
	}
	if !IsNil(o.NetPreferredStockIssuance) {
		toSerialize["net_preferred_stock_issuance"] = o.NetPreferredStockIssuance
	}
	if !IsNil(o.PreferredStockPayments) {
		toSerialize["preferred_stock_payments"] = o.PreferredStockPayments
	}
	if !IsNil(o.PreferredStockIssuance) {
		toSerialize["preferred_stock_issuance"] = o.PreferredStockIssuance
	}
	if !IsNil(o.NetCommonStockIssuance) {
		toSerialize["net_common_stock_issuance"] = o.NetCommonStockIssuance
	}
	if !IsNil(o.CommonStockPayments) {
		toSerialize["common_stock_payments"] = o.CommonStockPayments
	}
	if !IsNil(o.CommonStockIssuance) {
		toSerialize["common_stock_issuance"] = o.CommonStockIssuance
	}
	if !IsNil(o.RepurchaseOfCapitalStock) {
		toSerialize["repurchase_of_capital_stock"] = o.RepurchaseOfCapitalStock
	}
	if !IsNil(o.NetIssuancePaymentsOfDebt) {
		toSerialize["net_issuance_payments_of_debt"] = o.NetIssuancePaymentsOfDebt
	}
	if !IsNil(o.NetShortTermDebtIssuance) {
		toSerialize["net_short_term_debt_issuance"] = o.NetShortTermDebtIssuance
	}
	if !IsNil(o.ShortTermDebtPayments) {
		toSerialize["short_term_debt_payments"] = o.ShortTermDebtPayments
	}
	if !IsNil(o.ShortTermDebtIssuance) {
		toSerialize["short_term_debt_issuance"] = o.ShortTermDebtIssuance
	}
	if !IsNil(o.NetLongTermDebtIssuance) {
		toSerialize["net_long_term_debt_issuance"] = o.NetLongTermDebtIssuance
	}
	if !IsNil(o.LongTermDebtPayments) {
		toSerialize["long_term_debt_payments"] = o.LongTermDebtPayments
	}
	if !IsNil(o.LongTermDebtIssuance) {
		toSerialize["long_term_debt_issuance"] = o.LongTermDebtIssuance
	}
	if !IsNil(o.IssuanceOfDebt) {
		toSerialize["issuance_of_debt"] = o.IssuanceOfDebt
	}
	if !IsNil(o.RepaymentOfDebt) {
		toSerialize["repayment_of_debt"] = o.RepaymentOfDebt
	}
	if !IsNil(o.IssuanceOfCapitalStock) {
		toSerialize["issuance_of_capital_stock"] = o.IssuanceOfCapitalStock
	}
	return toSerialize, nil
}

type NullableCashFlowDataCashFlowFromFinancingActivities struct {
	value *CashFlowDataCashFlowFromFinancingActivities
	isSet bool
}

func (v NullableCashFlowDataCashFlowFromFinancingActivities) Get() *CashFlowDataCashFlowFromFinancingActivities {
	return v.value
}

func (v *NullableCashFlowDataCashFlowFromFinancingActivities) Set(val *CashFlowDataCashFlowFromFinancingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataCashFlowFromFinancingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataCashFlowFromFinancingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataCashFlowFromFinancingActivities(val *CashFlowDataCashFlowFromFinancingActivities) *NullableCashFlowDataCashFlowFromFinancingActivities {
	return &NullableCashFlowDataCashFlowFromFinancingActivities{value: val, isSet: true}
}

func (v NullableCashFlowDataCashFlowFromFinancingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataCashFlowFromFinancingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


