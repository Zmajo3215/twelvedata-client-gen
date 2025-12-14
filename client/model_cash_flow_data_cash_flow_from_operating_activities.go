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

// checks if the CashFlowDataCashFlowFromOperatingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataCashFlowFromOperatingActivities{}

// CashFlowDataCashFlowFromOperatingActivities Cash flow from operating activities
type CashFlowDataCashFlowFromOperatingActivities struct {
	// Net income from continuing operations
	NetIncomeFromContinuingOperations *float64 `json:"net_income_from_continuing_operations,omitempty"`
	// Operating cash flow
	OperatingCashFlow *float64 `json:"operating_cash_flow,omitempty"`
	// Cash flow from continuing operating activities
	CashFlowFromContinuingOperatingActivities *float64 `json:"cash_flow_from_continuing_operating_activities,omitempty"`
	// Cash from discontinued operating activities
	CashFromDiscontinuedOperatingActivities *float64 `json:"cash_from_discontinued_operating_activities,omitempty"`
	// Cash flow from discontinued operation
	CashFlowFromDiscontinuedOperation *float64 `json:"cash_flow_from_discontinued_operation,omitempty"`
	// Free cash flow
	FreeCashFlow *float64 `json:"free_cash_flow,omitempty"`
	// Cash flows from used in operating activities direct
	CashFlowsFromUsedInOperatingActivitiesDirect *float64 `json:"cash_flows_from_used_in_operating_activities_direct,omitempty"`
	// Taxes refund paid
	TaxesRefundPaid *float64 `json:"taxes_refund_paid,omitempty"`
	// Taxes refund paid direct
	TaxesRefundPaidDirect *float64 `json:"taxes_refund_paid_direct,omitempty"`
	// Interest received
	InterestReceived *float64 `json:"interest_received,omitempty"`
	// Interest received direct
	InterestReceivedDirect *float64 `json:"interest_received_direct,omitempty"`
	// Interest paid
	InterestPaid *float64 `json:"interest_paid,omitempty"`
	// Interest paid direct
	InterestPaidDirect *float64 `json:"interest_paid_direct,omitempty"`
	// Dividend received
	DividendReceived *float64 `json:"dividend_received,omitempty"`
	// Dividend received direct
	DividendReceivedDirect *float64 `json:"dividend_received_direct,omitempty"`
	// Dividend paid
	DividendPaid *float64 `json:"dividend_paid,omitempty"`
	// Dividend paid direct
	DividendPaidDirect *float64 `json:"dividend_paid_direct,omitempty"`
	// Change in working capital
	ChangeInWorkingCapital *float64 `json:"change_in_working_capital,omitempty"`
	// Change in other working capital
	ChangeInOtherWorkingCapital *float64 `json:"change_in_other_working_capital,omitempty"`
	// Change in receivables
	ChangeInReceivables *float64 `json:"change_in_receivables,omitempty"`
	// Changes in account receivables
	ChangesInAccountReceivables *float64 `json:"changes_in_account_receivables,omitempty"`
	// Change in payables and accrued expense
	ChangeInPayablesAndAccruedExpense *float64 `json:"change_in_payables_and_accrued_expense,omitempty"`
	// Change in accrued expense
	ChangeInAccruedExpense *float64 `json:"change_in_accrued_expense,omitempty"`
	// Change in payable
	ChangeInPayable *float64 `json:"change_in_payable,omitempty"`
	// Change in dividend payable
	ChangeInDividendPayable *float64 `json:"change_in_dividend_payable,omitempty"`
	// Change in account payable
	ChangeInAccountPayable *float64 `json:"change_in_account_payable,omitempty"`
	// Change in tax payable
	ChangeInTaxPayable *float64 `json:"change_in_tax_payable,omitempty"`
	// Change in income tax payable
	ChangeInIncomeTaxPayable *float64 `json:"change_in_income_tax_payable,omitempty"`
	// Change in interest payable
	ChangeInInterestPayable *float64 `json:"change_in_interest_payable,omitempty"`
	// Change in other current liabilities
	ChangeInOtherCurrentLiabilities *float64 `json:"change_in_other_current_liabilities,omitempty"`
	// Change in other current assets
	ChangeInOtherCurrentAssets *float64 `json:"change_in_other_current_assets,omitempty"`
	// Change in inventory
	ChangeInInventory *float64 `json:"change_in_inventory,omitempty"`
	// Change in prepaid assets
	ChangeInPrepaidAssets *float64 `json:"change_in_prepaid_assets,omitempty"`
	// Other non cash items
	OtherNonCashItems *float64 `json:"other_non_cash_items,omitempty"`
	// Excess tax benefit from stock based compensation
	ExcessTaxBenefitFromStockBasedCompensation *float64 `json:"excess_tax_benefit_from_stock_based_compensation,omitempty"`
	// Stock based compensation
	StockBasedCompensation *float64 `json:"stock_based_compensation,omitempty"`
	// Unrealized gain loss on investment securities
	UnrealizedGainLossOnInvestmentSecurities *float64 `json:"unrealized_gain_loss_on_investment_securities,omitempty"`
	// Provision and write off of assets
	ProvisionAndWriteOffOfAssets *float64 `json:"provision_and_write_off_of_assets,omitempty"`
	// Asset impairment charge
	AssetImpairmentCharge *float64 `json:"asset_impairment_charge,omitempty"`
	// Amortization of securities
	AmortizationOfSecurities *float64 `json:"amortization_of_securities,omitempty"`
	// Deferred tax
	DeferredTax *float64 `json:"deferred_tax,omitempty"`
	// Deferred income tax
	DeferredIncomeTax *float64 `json:"deferred_income_tax,omitempty"`
	// Depreciation amortization depletion
	DepreciationAmortizationDepletion *float64 `json:"depreciation_amortization_depletion,omitempty"`
	// Depletion
	Depletion *float64 `json:"depletion,omitempty"`
	// Depreciation and amortization
	DepreciationAndAmortization *float64 `json:"depreciation_and_amortization,omitempty"`
	// Amortization cash flow
	AmortizationCashFlow *float64 `json:"amortization_cash_flow,omitempty"`
	// Amortization of intangibles
	AmortizationOfIntangibles *float64 `json:"amortization_of_intangibles,omitempty"`
	// Depreciation
	Depreciation *float64 `json:"depreciation,omitempty"`
	// Operating gains losses
	OperatingGainsLosses *float64 `json:"operating_gains_losses,omitempty"`
	// Pension and employee benefit expense
	PensionAndEmployeeBenefitExpense *float64 `json:"pension_and_employee_benefit_expense,omitempty"`
	// Earnings losses from equity investments
	EarningsLossesFromEquityInvestments *float64 `json:"earnings_losses_from_equity_investments,omitempty"`
	// Gain loss on investment securities
	GainLossOnInvestmentSecurities *float64 `json:"gain_loss_on_investment_securities,omitempty"`
	// Net foreign currency exchange gain loss
	NetForeignCurrencyExchangeGainLoss *float64 `json:"net_foreign_currency_exchange_gain_loss,omitempty"`
	// Gain loss on sale of ppe
	GainLossOnSaleOfPpe *float64 `json:"gain_loss_on_sale_of_ppe,omitempty"`
	// Gain loss on sale of business
	GainLossOnSaleOfBusiness *float64 `json:"gain_loss_on_sale_of_business,omitempty"`
}

// NewCashFlowDataCashFlowFromOperatingActivities instantiates a new CashFlowDataCashFlowFromOperatingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataCashFlowFromOperatingActivities() *CashFlowDataCashFlowFromOperatingActivities {
	this := CashFlowDataCashFlowFromOperatingActivities{}
	return &this
}

// NewCashFlowDataCashFlowFromOperatingActivitiesWithDefaults instantiates a new CashFlowDataCashFlowFromOperatingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataCashFlowFromOperatingActivitiesWithDefaults() *CashFlowDataCashFlowFromOperatingActivities {
	this := CashFlowDataCashFlowFromOperatingActivities{}
	return &this
}

// GetNetIncomeFromContinuingOperations returns the NetIncomeFromContinuingOperations field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetNetIncomeFromContinuingOperations() float64 {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperations) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromContinuingOperations
}

// GetNetIncomeFromContinuingOperationsOk returns a tuple with the NetIncomeFromContinuingOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetNetIncomeFromContinuingOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperations) {
		return nil, false
	}
	return o.NetIncomeFromContinuingOperations, true
}

// HasNetIncomeFromContinuingOperations returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasNetIncomeFromContinuingOperations() bool {
	if o != nil && !IsNil(o.NetIncomeFromContinuingOperations) {
		return true
	}

	return false
}

// SetNetIncomeFromContinuingOperations gets a reference to the given float64 and assigns it to the NetIncomeFromContinuingOperations field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetNetIncomeFromContinuingOperations(v float64) {
	o.NetIncomeFromContinuingOperations = &v
}

// GetOperatingCashFlow returns the OperatingCashFlow field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOperatingCashFlow() float64 {
	if o == nil || IsNil(o.OperatingCashFlow) {
		var ret float64
		return ret
	}
	return *o.OperatingCashFlow
}

// GetOperatingCashFlowOk returns a tuple with the OperatingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOperatingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingCashFlow) {
		return nil, false
	}
	return o.OperatingCashFlow, true
}

// HasOperatingCashFlow returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasOperatingCashFlow() bool {
	if o != nil && !IsNil(o.OperatingCashFlow) {
		return true
	}

	return false
}

// SetOperatingCashFlow gets a reference to the given float64 and assigns it to the OperatingCashFlow field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetOperatingCashFlow(v float64) {
	o.OperatingCashFlow = &v
}

// GetCashFlowFromContinuingOperatingActivities returns the CashFlowFromContinuingOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowFromContinuingOperatingActivities() float64 {
	if o == nil || IsNil(o.CashFlowFromContinuingOperatingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFlowFromContinuingOperatingActivities
}

// GetCashFlowFromContinuingOperatingActivitiesOk returns a tuple with the CashFlowFromContinuingOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowFromContinuingOperatingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFlowFromContinuingOperatingActivities) {
		return nil, false
	}
	return o.CashFlowFromContinuingOperatingActivities, true
}

// HasCashFlowFromContinuingOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasCashFlowFromContinuingOperatingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromContinuingOperatingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromContinuingOperatingActivities gets a reference to the given float64 and assigns it to the CashFlowFromContinuingOperatingActivities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetCashFlowFromContinuingOperatingActivities(v float64) {
	o.CashFlowFromContinuingOperatingActivities = &v
}

// GetCashFromDiscontinuedOperatingActivities returns the CashFromDiscontinuedOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFromDiscontinuedOperatingActivities() float64 {
	if o == nil || IsNil(o.CashFromDiscontinuedOperatingActivities) {
		var ret float64
		return ret
	}
	return *o.CashFromDiscontinuedOperatingActivities
}

// GetCashFromDiscontinuedOperatingActivitiesOk returns a tuple with the CashFromDiscontinuedOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFromDiscontinuedOperatingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFromDiscontinuedOperatingActivities) {
		return nil, false
	}
	return o.CashFromDiscontinuedOperatingActivities, true
}

// HasCashFromDiscontinuedOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasCashFromDiscontinuedOperatingActivities() bool {
	if o != nil && !IsNil(o.CashFromDiscontinuedOperatingActivities) {
		return true
	}

	return false
}

// SetCashFromDiscontinuedOperatingActivities gets a reference to the given float64 and assigns it to the CashFromDiscontinuedOperatingActivities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetCashFromDiscontinuedOperatingActivities(v float64) {
	o.CashFromDiscontinuedOperatingActivities = &v
}

// GetCashFlowFromDiscontinuedOperation returns the CashFlowFromDiscontinuedOperation field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowFromDiscontinuedOperation() float64 {
	if o == nil || IsNil(o.CashFlowFromDiscontinuedOperation) {
		var ret float64
		return ret
	}
	return *o.CashFlowFromDiscontinuedOperation
}

// GetCashFlowFromDiscontinuedOperationOk returns a tuple with the CashFlowFromDiscontinuedOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowFromDiscontinuedOperationOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFlowFromDiscontinuedOperation) {
		return nil, false
	}
	return o.CashFlowFromDiscontinuedOperation, true
}

// HasCashFlowFromDiscontinuedOperation returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasCashFlowFromDiscontinuedOperation() bool {
	if o != nil && !IsNil(o.CashFlowFromDiscontinuedOperation) {
		return true
	}

	return false
}

// SetCashFlowFromDiscontinuedOperation gets a reference to the given float64 and assigns it to the CashFlowFromDiscontinuedOperation field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetCashFlowFromDiscontinuedOperation(v float64) {
	o.CashFlowFromDiscontinuedOperation = &v
}

// GetFreeCashFlow returns the FreeCashFlow field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetFreeCashFlow() float64 {
	if o == nil || IsNil(o.FreeCashFlow) {
		var ret float64
		return ret
	}
	return *o.FreeCashFlow
}

// GetFreeCashFlowOk returns a tuple with the FreeCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetFreeCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.FreeCashFlow) {
		return nil, false
	}
	return o.FreeCashFlow, true
}

// HasFreeCashFlow returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasFreeCashFlow() bool {
	if o != nil && !IsNil(o.FreeCashFlow) {
		return true
	}

	return false
}

// SetFreeCashFlow gets a reference to the given float64 and assigns it to the FreeCashFlow field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetFreeCashFlow(v float64) {
	o.FreeCashFlow = &v
}

// GetCashFlowsFromUsedInOperatingActivitiesDirect returns the CashFlowsFromUsedInOperatingActivitiesDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowsFromUsedInOperatingActivitiesDirect() float64 {
	if o == nil || IsNil(o.CashFlowsFromUsedInOperatingActivitiesDirect) {
		var ret float64
		return ret
	}
	return *o.CashFlowsFromUsedInOperatingActivitiesDirect
}

// GetCashFlowsFromUsedInOperatingActivitiesDirectOk returns a tuple with the CashFlowsFromUsedInOperatingActivitiesDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetCashFlowsFromUsedInOperatingActivitiesDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFlowsFromUsedInOperatingActivitiesDirect) {
		return nil, false
	}
	return o.CashFlowsFromUsedInOperatingActivitiesDirect, true
}

// HasCashFlowsFromUsedInOperatingActivitiesDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasCashFlowsFromUsedInOperatingActivitiesDirect() bool {
	if o != nil && !IsNil(o.CashFlowsFromUsedInOperatingActivitiesDirect) {
		return true
	}

	return false
}

// SetCashFlowsFromUsedInOperatingActivitiesDirect gets a reference to the given float64 and assigns it to the CashFlowsFromUsedInOperatingActivitiesDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetCashFlowsFromUsedInOperatingActivitiesDirect(v float64) {
	o.CashFlowsFromUsedInOperatingActivitiesDirect = &v
}

// GetTaxesRefundPaid returns the TaxesRefundPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetTaxesRefundPaid() float64 {
	if o == nil || IsNil(o.TaxesRefundPaid) {
		var ret float64
		return ret
	}
	return *o.TaxesRefundPaid
}

// GetTaxesRefundPaidOk returns a tuple with the TaxesRefundPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetTaxesRefundPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxesRefundPaid) {
		return nil, false
	}
	return o.TaxesRefundPaid, true
}

// HasTaxesRefundPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasTaxesRefundPaid() bool {
	if o != nil && !IsNil(o.TaxesRefundPaid) {
		return true
	}

	return false
}

// SetTaxesRefundPaid gets a reference to the given float64 and assigns it to the TaxesRefundPaid field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetTaxesRefundPaid(v float64) {
	o.TaxesRefundPaid = &v
}

// GetTaxesRefundPaidDirect returns the TaxesRefundPaidDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetTaxesRefundPaidDirect() float64 {
	if o == nil || IsNil(o.TaxesRefundPaidDirect) {
		var ret float64
		return ret
	}
	return *o.TaxesRefundPaidDirect
}

// GetTaxesRefundPaidDirectOk returns a tuple with the TaxesRefundPaidDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetTaxesRefundPaidDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxesRefundPaidDirect) {
		return nil, false
	}
	return o.TaxesRefundPaidDirect, true
}

// HasTaxesRefundPaidDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasTaxesRefundPaidDirect() bool {
	if o != nil && !IsNil(o.TaxesRefundPaidDirect) {
		return true
	}

	return false
}

// SetTaxesRefundPaidDirect gets a reference to the given float64 and assigns it to the TaxesRefundPaidDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetTaxesRefundPaidDirect(v float64) {
	o.TaxesRefundPaidDirect = &v
}

// GetInterestReceived returns the InterestReceived field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestReceived() float64 {
	if o == nil || IsNil(o.InterestReceived) {
		var ret float64
		return ret
	}
	return *o.InterestReceived
}

// GetInterestReceivedOk returns a tuple with the InterestReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestReceivedOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestReceived) {
		return nil, false
	}
	return o.InterestReceived, true
}

// HasInterestReceived returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasInterestReceived() bool {
	if o != nil && !IsNil(o.InterestReceived) {
		return true
	}

	return false
}

// SetInterestReceived gets a reference to the given float64 and assigns it to the InterestReceived field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetInterestReceived(v float64) {
	o.InterestReceived = &v
}

// GetInterestReceivedDirect returns the InterestReceivedDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestReceivedDirect() float64 {
	if o == nil || IsNil(o.InterestReceivedDirect) {
		var ret float64
		return ret
	}
	return *o.InterestReceivedDirect
}

// GetInterestReceivedDirectOk returns a tuple with the InterestReceivedDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestReceivedDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestReceivedDirect) {
		return nil, false
	}
	return o.InterestReceivedDirect, true
}

// HasInterestReceivedDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasInterestReceivedDirect() bool {
	if o != nil && !IsNil(o.InterestReceivedDirect) {
		return true
	}

	return false
}

// SetInterestReceivedDirect gets a reference to the given float64 and assigns it to the InterestReceivedDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetInterestReceivedDirect(v float64) {
	o.InterestReceivedDirect = &v
}

// GetInterestPaid returns the InterestPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestPaid() float64 {
	if o == nil || IsNil(o.InterestPaid) {
		var ret float64
		return ret
	}
	return *o.InterestPaid
}

// GetInterestPaidOk returns a tuple with the InterestPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPaid) {
		return nil, false
	}
	return o.InterestPaid, true
}

// HasInterestPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasInterestPaid() bool {
	if o != nil && !IsNil(o.InterestPaid) {
		return true
	}

	return false
}

// SetInterestPaid gets a reference to the given float64 and assigns it to the InterestPaid field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetInterestPaid(v float64) {
	o.InterestPaid = &v
}

// GetInterestPaidDirect returns the InterestPaidDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestPaidDirect() float64 {
	if o == nil || IsNil(o.InterestPaidDirect) {
		var ret float64
		return ret
	}
	return *o.InterestPaidDirect
}

// GetInterestPaidDirectOk returns a tuple with the InterestPaidDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetInterestPaidDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPaidDirect) {
		return nil, false
	}
	return o.InterestPaidDirect, true
}

// HasInterestPaidDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasInterestPaidDirect() bool {
	if o != nil && !IsNil(o.InterestPaidDirect) {
		return true
	}

	return false
}

// SetInterestPaidDirect gets a reference to the given float64 and assigns it to the InterestPaidDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetInterestPaidDirect(v float64) {
	o.InterestPaidDirect = &v
}

// GetDividendReceived returns the DividendReceived field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendReceived() float64 {
	if o == nil || IsNil(o.DividendReceived) {
		var ret float64
		return ret
	}
	return *o.DividendReceived
}

// GetDividendReceivedOk returns a tuple with the DividendReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendReceivedOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendReceived) {
		return nil, false
	}
	return o.DividendReceived, true
}

// HasDividendReceived returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDividendReceived() bool {
	if o != nil && !IsNil(o.DividendReceived) {
		return true
	}

	return false
}

// SetDividendReceived gets a reference to the given float64 and assigns it to the DividendReceived field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDividendReceived(v float64) {
	o.DividendReceived = &v
}

// GetDividendReceivedDirect returns the DividendReceivedDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendReceivedDirect() float64 {
	if o == nil || IsNil(o.DividendReceivedDirect) {
		var ret float64
		return ret
	}
	return *o.DividendReceivedDirect
}

// GetDividendReceivedDirectOk returns a tuple with the DividendReceivedDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendReceivedDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendReceivedDirect) {
		return nil, false
	}
	return o.DividendReceivedDirect, true
}

// HasDividendReceivedDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDividendReceivedDirect() bool {
	if o != nil && !IsNil(o.DividendReceivedDirect) {
		return true
	}

	return false
}

// SetDividendReceivedDirect gets a reference to the given float64 and assigns it to the DividendReceivedDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDividendReceivedDirect(v float64) {
	o.DividendReceivedDirect = &v
}

// GetDividendPaid returns the DividendPaid field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendPaid() float64 {
	if o == nil || IsNil(o.DividendPaid) {
		var ret float64
		return ret
	}
	return *o.DividendPaid
}

// GetDividendPaidOk returns a tuple with the DividendPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendPaid) {
		return nil, false
	}
	return o.DividendPaid, true
}

// HasDividendPaid returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDividendPaid() bool {
	if o != nil && !IsNil(o.DividendPaid) {
		return true
	}

	return false
}

// SetDividendPaid gets a reference to the given float64 and assigns it to the DividendPaid field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDividendPaid(v float64) {
	o.DividendPaid = &v
}

// GetDividendPaidDirect returns the DividendPaidDirect field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendPaidDirect() float64 {
	if o == nil || IsNil(o.DividendPaidDirect) {
		var ret float64
		return ret
	}
	return *o.DividendPaidDirect
}

// GetDividendPaidDirectOk returns a tuple with the DividendPaidDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDividendPaidDirectOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendPaidDirect) {
		return nil, false
	}
	return o.DividendPaidDirect, true
}

// HasDividendPaidDirect returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDividendPaidDirect() bool {
	if o != nil && !IsNil(o.DividendPaidDirect) {
		return true
	}

	return false
}

// SetDividendPaidDirect gets a reference to the given float64 and assigns it to the DividendPaidDirect field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDividendPaidDirect(v float64) {
	o.DividendPaidDirect = &v
}

// GetChangeInWorkingCapital returns the ChangeInWorkingCapital field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInWorkingCapital() float64 {
	if o == nil || IsNil(o.ChangeInWorkingCapital) {
		var ret float64
		return ret
	}
	return *o.ChangeInWorkingCapital
}

// GetChangeInWorkingCapitalOk returns a tuple with the ChangeInWorkingCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInWorkingCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInWorkingCapital) {
		return nil, false
	}
	return o.ChangeInWorkingCapital, true
}

// HasChangeInWorkingCapital returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInWorkingCapital() bool {
	if o != nil && !IsNil(o.ChangeInWorkingCapital) {
		return true
	}

	return false
}

// SetChangeInWorkingCapital gets a reference to the given float64 and assigns it to the ChangeInWorkingCapital field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInWorkingCapital(v float64) {
	o.ChangeInWorkingCapital = &v
}

// GetChangeInOtherWorkingCapital returns the ChangeInOtherWorkingCapital field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherWorkingCapital() float64 {
	if o == nil || IsNil(o.ChangeInOtherWorkingCapital) {
		var ret float64
		return ret
	}
	return *o.ChangeInOtherWorkingCapital
}

// GetChangeInOtherWorkingCapitalOk returns a tuple with the ChangeInOtherWorkingCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherWorkingCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInOtherWorkingCapital) {
		return nil, false
	}
	return o.ChangeInOtherWorkingCapital, true
}

// HasChangeInOtherWorkingCapital returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInOtherWorkingCapital() bool {
	if o != nil && !IsNil(o.ChangeInOtherWorkingCapital) {
		return true
	}

	return false
}

// SetChangeInOtherWorkingCapital gets a reference to the given float64 and assigns it to the ChangeInOtherWorkingCapital field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInOtherWorkingCapital(v float64) {
	o.ChangeInOtherWorkingCapital = &v
}

// GetChangeInReceivables returns the ChangeInReceivables field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInReceivables() float64 {
	if o == nil || IsNil(o.ChangeInReceivables) {
		var ret float64
		return ret
	}
	return *o.ChangeInReceivables
}

// GetChangeInReceivablesOk returns a tuple with the ChangeInReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInReceivables) {
		return nil, false
	}
	return o.ChangeInReceivables, true
}

// HasChangeInReceivables returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInReceivables() bool {
	if o != nil && !IsNil(o.ChangeInReceivables) {
		return true
	}

	return false
}

// SetChangeInReceivables gets a reference to the given float64 and assigns it to the ChangeInReceivables field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInReceivables(v float64) {
	o.ChangeInReceivables = &v
}

// GetChangesInAccountReceivables returns the ChangesInAccountReceivables field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangesInAccountReceivables() float64 {
	if o == nil || IsNil(o.ChangesInAccountReceivables) {
		var ret float64
		return ret
	}
	return *o.ChangesInAccountReceivables
}

// GetChangesInAccountReceivablesOk returns a tuple with the ChangesInAccountReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangesInAccountReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangesInAccountReceivables) {
		return nil, false
	}
	return o.ChangesInAccountReceivables, true
}

// HasChangesInAccountReceivables returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangesInAccountReceivables() bool {
	if o != nil && !IsNil(o.ChangesInAccountReceivables) {
		return true
	}

	return false
}

// SetChangesInAccountReceivables gets a reference to the given float64 and assigns it to the ChangesInAccountReceivables field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangesInAccountReceivables(v float64) {
	o.ChangesInAccountReceivables = &v
}

// GetChangeInPayablesAndAccruedExpense returns the ChangeInPayablesAndAccruedExpense field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPayablesAndAccruedExpense() float64 {
	if o == nil || IsNil(o.ChangeInPayablesAndAccruedExpense) {
		var ret float64
		return ret
	}
	return *o.ChangeInPayablesAndAccruedExpense
}

// GetChangeInPayablesAndAccruedExpenseOk returns a tuple with the ChangeInPayablesAndAccruedExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPayablesAndAccruedExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInPayablesAndAccruedExpense) {
		return nil, false
	}
	return o.ChangeInPayablesAndAccruedExpense, true
}

// HasChangeInPayablesAndAccruedExpense returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInPayablesAndAccruedExpense() bool {
	if o != nil && !IsNil(o.ChangeInPayablesAndAccruedExpense) {
		return true
	}

	return false
}

// SetChangeInPayablesAndAccruedExpense gets a reference to the given float64 and assigns it to the ChangeInPayablesAndAccruedExpense field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInPayablesAndAccruedExpense(v float64) {
	o.ChangeInPayablesAndAccruedExpense = &v
}

// GetChangeInAccruedExpense returns the ChangeInAccruedExpense field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInAccruedExpense() float64 {
	if o == nil || IsNil(o.ChangeInAccruedExpense) {
		var ret float64
		return ret
	}
	return *o.ChangeInAccruedExpense
}

// GetChangeInAccruedExpenseOk returns a tuple with the ChangeInAccruedExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInAccruedExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInAccruedExpense) {
		return nil, false
	}
	return o.ChangeInAccruedExpense, true
}

// HasChangeInAccruedExpense returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInAccruedExpense() bool {
	if o != nil && !IsNil(o.ChangeInAccruedExpense) {
		return true
	}

	return false
}

// SetChangeInAccruedExpense gets a reference to the given float64 and assigns it to the ChangeInAccruedExpense field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInAccruedExpense(v float64) {
	o.ChangeInAccruedExpense = &v
}

// GetChangeInPayable returns the ChangeInPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPayable() float64 {
	if o == nil || IsNil(o.ChangeInPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInPayable
}

// GetChangeInPayableOk returns a tuple with the ChangeInPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInPayable) {
		return nil, false
	}
	return o.ChangeInPayable, true
}

// HasChangeInPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInPayable() bool {
	if o != nil && !IsNil(o.ChangeInPayable) {
		return true
	}

	return false
}

// SetChangeInPayable gets a reference to the given float64 and assigns it to the ChangeInPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInPayable(v float64) {
	o.ChangeInPayable = &v
}

// GetChangeInDividendPayable returns the ChangeInDividendPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInDividendPayable() float64 {
	if o == nil || IsNil(o.ChangeInDividendPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInDividendPayable
}

// GetChangeInDividendPayableOk returns a tuple with the ChangeInDividendPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInDividendPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInDividendPayable) {
		return nil, false
	}
	return o.ChangeInDividendPayable, true
}

// HasChangeInDividendPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInDividendPayable() bool {
	if o != nil && !IsNil(o.ChangeInDividendPayable) {
		return true
	}

	return false
}

// SetChangeInDividendPayable gets a reference to the given float64 and assigns it to the ChangeInDividendPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInDividendPayable(v float64) {
	o.ChangeInDividendPayable = &v
}

// GetChangeInAccountPayable returns the ChangeInAccountPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInAccountPayable() float64 {
	if o == nil || IsNil(o.ChangeInAccountPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInAccountPayable
}

// GetChangeInAccountPayableOk returns a tuple with the ChangeInAccountPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInAccountPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInAccountPayable) {
		return nil, false
	}
	return o.ChangeInAccountPayable, true
}

// HasChangeInAccountPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInAccountPayable() bool {
	if o != nil && !IsNil(o.ChangeInAccountPayable) {
		return true
	}

	return false
}

// SetChangeInAccountPayable gets a reference to the given float64 and assigns it to the ChangeInAccountPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInAccountPayable(v float64) {
	o.ChangeInAccountPayable = &v
}

// GetChangeInTaxPayable returns the ChangeInTaxPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInTaxPayable() float64 {
	if o == nil || IsNil(o.ChangeInTaxPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInTaxPayable
}

// GetChangeInTaxPayableOk returns a tuple with the ChangeInTaxPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInTaxPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInTaxPayable) {
		return nil, false
	}
	return o.ChangeInTaxPayable, true
}

// HasChangeInTaxPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInTaxPayable() bool {
	if o != nil && !IsNil(o.ChangeInTaxPayable) {
		return true
	}

	return false
}

// SetChangeInTaxPayable gets a reference to the given float64 and assigns it to the ChangeInTaxPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInTaxPayable(v float64) {
	o.ChangeInTaxPayable = &v
}

// GetChangeInIncomeTaxPayable returns the ChangeInIncomeTaxPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInIncomeTaxPayable() float64 {
	if o == nil || IsNil(o.ChangeInIncomeTaxPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInIncomeTaxPayable
}

// GetChangeInIncomeTaxPayableOk returns a tuple with the ChangeInIncomeTaxPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInIncomeTaxPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInIncomeTaxPayable) {
		return nil, false
	}
	return o.ChangeInIncomeTaxPayable, true
}

// HasChangeInIncomeTaxPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInIncomeTaxPayable() bool {
	if o != nil && !IsNil(o.ChangeInIncomeTaxPayable) {
		return true
	}

	return false
}

// SetChangeInIncomeTaxPayable gets a reference to the given float64 and assigns it to the ChangeInIncomeTaxPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInIncomeTaxPayable(v float64) {
	o.ChangeInIncomeTaxPayable = &v
}

// GetChangeInInterestPayable returns the ChangeInInterestPayable field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInInterestPayable() float64 {
	if o == nil || IsNil(o.ChangeInInterestPayable) {
		var ret float64
		return ret
	}
	return *o.ChangeInInterestPayable
}

// GetChangeInInterestPayableOk returns a tuple with the ChangeInInterestPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInInterestPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInInterestPayable) {
		return nil, false
	}
	return o.ChangeInInterestPayable, true
}

// HasChangeInInterestPayable returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInInterestPayable() bool {
	if o != nil && !IsNil(o.ChangeInInterestPayable) {
		return true
	}

	return false
}

// SetChangeInInterestPayable gets a reference to the given float64 and assigns it to the ChangeInInterestPayable field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInInterestPayable(v float64) {
	o.ChangeInInterestPayable = &v
}

// GetChangeInOtherCurrentLiabilities returns the ChangeInOtherCurrentLiabilities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherCurrentLiabilities() float64 {
	if o == nil || IsNil(o.ChangeInOtherCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.ChangeInOtherCurrentLiabilities
}

// GetChangeInOtherCurrentLiabilitiesOk returns a tuple with the ChangeInOtherCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInOtherCurrentLiabilities) {
		return nil, false
	}
	return o.ChangeInOtherCurrentLiabilities, true
}

// HasChangeInOtherCurrentLiabilities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInOtherCurrentLiabilities() bool {
	if o != nil && !IsNil(o.ChangeInOtherCurrentLiabilities) {
		return true
	}

	return false
}

// SetChangeInOtherCurrentLiabilities gets a reference to the given float64 and assigns it to the ChangeInOtherCurrentLiabilities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInOtherCurrentLiabilities(v float64) {
	o.ChangeInOtherCurrentLiabilities = &v
}

// GetChangeInOtherCurrentAssets returns the ChangeInOtherCurrentAssets field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherCurrentAssets() float64 {
	if o == nil || IsNil(o.ChangeInOtherCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.ChangeInOtherCurrentAssets
}

// GetChangeInOtherCurrentAssetsOk returns a tuple with the ChangeInOtherCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInOtherCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInOtherCurrentAssets) {
		return nil, false
	}
	return o.ChangeInOtherCurrentAssets, true
}

// HasChangeInOtherCurrentAssets returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInOtherCurrentAssets() bool {
	if o != nil && !IsNil(o.ChangeInOtherCurrentAssets) {
		return true
	}

	return false
}

// SetChangeInOtherCurrentAssets gets a reference to the given float64 and assigns it to the ChangeInOtherCurrentAssets field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInOtherCurrentAssets(v float64) {
	o.ChangeInOtherCurrentAssets = &v
}

// GetChangeInInventory returns the ChangeInInventory field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInInventory() float64 {
	if o == nil || IsNil(o.ChangeInInventory) {
		var ret float64
		return ret
	}
	return *o.ChangeInInventory
}

// GetChangeInInventoryOk returns a tuple with the ChangeInInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInInventoryOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInInventory) {
		return nil, false
	}
	return o.ChangeInInventory, true
}

// HasChangeInInventory returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInInventory() bool {
	if o != nil && !IsNil(o.ChangeInInventory) {
		return true
	}

	return false
}

// SetChangeInInventory gets a reference to the given float64 and assigns it to the ChangeInInventory field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInInventory(v float64) {
	o.ChangeInInventory = &v
}

// GetChangeInPrepaidAssets returns the ChangeInPrepaidAssets field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPrepaidAssets() float64 {
	if o == nil || IsNil(o.ChangeInPrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.ChangeInPrepaidAssets
}

// GetChangeInPrepaidAssetsOk returns a tuple with the ChangeInPrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetChangeInPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeInPrepaidAssets) {
		return nil, false
	}
	return o.ChangeInPrepaidAssets, true
}

// HasChangeInPrepaidAssets returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasChangeInPrepaidAssets() bool {
	if o != nil && !IsNil(o.ChangeInPrepaidAssets) {
		return true
	}

	return false
}

// SetChangeInPrepaidAssets gets a reference to the given float64 and assigns it to the ChangeInPrepaidAssets field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetChangeInPrepaidAssets(v float64) {
	o.ChangeInPrepaidAssets = &v
}

// GetOtherNonCashItems returns the OtherNonCashItems field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOtherNonCashItems() float64 {
	if o == nil || IsNil(o.OtherNonCashItems) {
		var ret float64
		return ret
	}
	return *o.OtherNonCashItems
}

// GetOtherNonCashItemsOk returns a tuple with the OtherNonCashItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOtherNonCashItemsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCashItems) {
		return nil, false
	}
	return o.OtherNonCashItems, true
}

// HasOtherNonCashItems returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasOtherNonCashItems() bool {
	if o != nil && !IsNil(o.OtherNonCashItems) {
		return true
	}

	return false
}

// SetOtherNonCashItems gets a reference to the given float64 and assigns it to the OtherNonCashItems field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetOtherNonCashItems(v float64) {
	o.OtherNonCashItems = &v
}

// GetExcessTaxBenefitFromStockBasedCompensation returns the ExcessTaxBenefitFromStockBasedCompensation field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetExcessTaxBenefitFromStockBasedCompensation() float64 {
	if o == nil || IsNil(o.ExcessTaxBenefitFromStockBasedCompensation) {
		var ret float64
		return ret
	}
	return *o.ExcessTaxBenefitFromStockBasedCompensation
}

// GetExcessTaxBenefitFromStockBasedCompensationOk returns a tuple with the ExcessTaxBenefitFromStockBasedCompensation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetExcessTaxBenefitFromStockBasedCompensationOk() (*float64, bool) {
	if o == nil || IsNil(o.ExcessTaxBenefitFromStockBasedCompensation) {
		return nil, false
	}
	return o.ExcessTaxBenefitFromStockBasedCompensation, true
}

// HasExcessTaxBenefitFromStockBasedCompensation returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasExcessTaxBenefitFromStockBasedCompensation() bool {
	if o != nil && !IsNil(o.ExcessTaxBenefitFromStockBasedCompensation) {
		return true
	}

	return false
}

// SetExcessTaxBenefitFromStockBasedCompensation gets a reference to the given float64 and assigns it to the ExcessTaxBenefitFromStockBasedCompensation field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetExcessTaxBenefitFromStockBasedCompensation(v float64) {
	o.ExcessTaxBenefitFromStockBasedCompensation = &v
}

// GetStockBasedCompensation returns the StockBasedCompensation field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetStockBasedCompensation() float64 {
	if o == nil || IsNil(o.StockBasedCompensation) {
		var ret float64
		return ret
	}
	return *o.StockBasedCompensation
}

// GetStockBasedCompensationOk returns a tuple with the StockBasedCompensation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetStockBasedCompensationOk() (*float64, bool) {
	if o == nil || IsNil(o.StockBasedCompensation) {
		return nil, false
	}
	return o.StockBasedCompensation, true
}

// HasStockBasedCompensation returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasStockBasedCompensation() bool {
	if o != nil && !IsNil(o.StockBasedCompensation) {
		return true
	}

	return false
}

// SetStockBasedCompensation gets a reference to the given float64 and assigns it to the StockBasedCompensation field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetStockBasedCompensation(v float64) {
	o.StockBasedCompensation = &v
}

// GetUnrealizedGainLossOnInvestmentSecurities returns the UnrealizedGainLossOnInvestmentSecurities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetUnrealizedGainLossOnInvestmentSecurities() float64 {
	if o == nil || IsNil(o.UnrealizedGainLossOnInvestmentSecurities) {
		var ret float64
		return ret
	}
	return *o.UnrealizedGainLossOnInvestmentSecurities
}

// GetUnrealizedGainLossOnInvestmentSecuritiesOk returns a tuple with the UnrealizedGainLossOnInvestmentSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetUnrealizedGainLossOnInvestmentSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.UnrealizedGainLossOnInvestmentSecurities) {
		return nil, false
	}
	return o.UnrealizedGainLossOnInvestmentSecurities, true
}

// HasUnrealizedGainLossOnInvestmentSecurities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasUnrealizedGainLossOnInvestmentSecurities() bool {
	if o != nil && !IsNil(o.UnrealizedGainLossOnInvestmentSecurities) {
		return true
	}

	return false
}

// SetUnrealizedGainLossOnInvestmentSecurities gets a reference to the given float64 and assigns it to the UnrealizedGainLossOnInvestmentSecurities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetUnrealizedGainLossOnInvestmentSecurities(v float64) {
	o.UnrealizedGainLossOnInvestmentSecurities = &v
}

// GetProvisionAndWriteOffOfAssets returns the ProvisionAndWriteOffOfAssets field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetProvisionAndWriteOffOfAssets() float64 {
	if o == nil || IsNil(o.ProvisionAndWriteOffOfAssets) {
		var ret float64
		return ret
	}
	return *o.ProvisionAndWriteOffOfAssets
}

// GetProvisionAndWriteOffOfAssetsOk returns a tuple with the ProvisionAndWriteOffOfAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetProvisionAndWriteOffOfAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionAndWriteOffOfAssets) {
		return nil, false
	}
	return o.ProvisionAndWriteOffOfAssets, true
}

// HasProvisionAndWriteOffOfAssets returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasProvisionAndWriteOffOfAssets() bool {
	if o != nil && !IsNil(o.ProvisionAndWriteOffOfAssets) {
		return true
	}

	return false
}

// SetProvisionAndWriteOffOfAssets gets a reference to the given float64 and assigns it to the ProvisionAndWriteOffOfAssets field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetProvisionAndWriteOffOfAssets(v float64) {
	o.ProvisionAndWriteOffOfAssets = &v
}

// GetAssetImpairmentCharge returns the AssetImpairmentCharge field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAssetImpairmentCharge() float64 {
	if o == nil || IsNil(o.AssetImpairmentCharge) {
		var ret float64
		return ret
	}
	return *o.AssetImpairmentCharge
}

// GetAssetImpairmentChargeOk returns a tuple with the AssetImpairmentCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAssetImpairmentChargeOk() (*float64, bool) {
	if o == nil || IsNil(o.AssetImpairmentCharge) {
		return nil, false
	}
	return o.AssetImpairmentCharge, true
}

// HasAssetImpairmentCharge returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasAssetImpairmentCharge() bool {
	if o != nil && !IsNil(o.AssetImpairmentCharge) {
		return true
	}

	return false
}

// SetAssetImpairmentCharge gets a reference to the given float64 and assigns it to the AssetImpairmentCharge field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetAssetImpairmentCharge(v float64) {
	o.AssetImpairmentCharge = &v
}

// GetAmortizationOfSecurities returns the AmortizationOfSecurities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationOfSecurities() float64 {
	if o == nil || IsNil(o.AmortizationOfSecurities) {
		var ret float64
		return ret
	}
	return *o.AmortizationOfSecurities
}

// GetAmortizationOfSecuritiesOk returns a tuple with the AmortizationOfSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationOfSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.AmortizationOfSecurities) {
		return nil, false
	}
	return o.AmortizationOfSecurities, true
}

// HasAmortizationOfSecurities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasAmortizationOfSecurities() bool {
	if o != nil && !IsNil(o.AmortizationOfSecurities) {
		return true
	}

	return false
}

// SetAmortizationOfSecurities gets a reference to the given float64 and assigns it to the AmortizationOfSecurities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetAmortizationOfSecurities(v float64) {
	o.AmortizationOfSecurities = &v
}

// GetDeferredTax returns the DeferredTax field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDeferredTax() float64 {
	if o == nil || IsNil(o.DeferredTax) {
		var ret float64
		return ret
	}
	return *o.DeferredTax
}

// GetDeferredTaxOk returns a tuple with the DeferredTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDeferredTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.DeferredTax) {
		return nil, false
	}
	return o.DeferredTax, true
}

// HasDeferredTax returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDeferredTax() bool {
	if o != nil && !IsNil(o.DeferredTax) {
		return true
	}

	return false
}

// SetDeferredTax gets a reference to the given float64 and assigns it to the DeferredTax field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDeferredTax(v float64) {
	o.DeferredTax = &v
}

// GetDeferredIncomeTax returns the DeferredIncomeTax field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDeferredIncomeTax() float64 {
	if o == nil || IsNil(o.DeferredIncomeTax) {
		var ret float64
		return ret
	}
	return *o.DeferredIncomeTax
}

// GetDeferredIncomeTaxOk returns a tuple with the DeferredIncomeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDeferredIncomeTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.DeferredIncomeTax) {
		return nil, false
	}
	return o.DeferredIncomeTax, true
}

// HasDeferredIncomeTax returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDeferredIncomeTax() bool {
	if o != nil && !IsNil(o.DeferredIncomeTax) {
		return true
	}

	return false
}

// SetDeferredIncomeTax gets a reference to the given float64 and assigns it to the DeferredIncomeTax field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDeferredIncomeTax(v float64) {
	o.DeferredIncomeTax = &v
}

// GetDepreciationAmortizationDepletion returns the DepreciationAmortizationDepletion field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciationAmortizationDepletion() float64 {
	if o == nil || IsNil(o.DepreciationAmortizationDepletion) {
		var ret float64
		return ret
	}
	return *o.DepreciationAmortizationDepletion
}

// GetDepreciationAmortizationDepletionOk returns a tuple with the DepreciationAmortizationDepletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciationAmortizationDepletionOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAmortizationDepletion) {
		return nil, false
	}
	return o.DepreciationAmortizationDepletion, true
}

// HasDepreciationAmortizationDepletion returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDepreciationAmortizationDepletion() bool {
	if o != nil && !IsNil(o.DepreciationAmortizationDepletion) {
		return true
	}

	return false
}

// SetDepreciationAmortizationDepletion gets a reference to the given float64 and assigns it to the DepreciationAmortizationDepletion field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDepreciationAmortizationDepletion(v float64) {
	o.DepreciationAmortizationDepletion = &v
}

// GetDepletion returns the Depletion field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepletion() float64 {
	if o == nil || IsNil(o.Depletion) {
		var ret float64
		return ret
	}
	return *o.Depletion
}

// GetDepletionOk returns a tuple with the Depletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Depletion) {
		return nil, false
	}
	return o.Depletion, true
}

// HasDepletion returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDepletion() bool {
	if o != nil && !IsNil(o.Depletion) {
		return true
	}

	return false
}

// SetDepletion gets a reference to the given float64 and assigns it to the Depletion field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDepletion(v float64) {
	o.Depletion = &v
}

// GetDepreciationAndAmortization returns the DepreciationAndAmortization field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciationAndAmortization() float64 {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		var ret float64
		return ret
	}
	return *o.DepreciationAndAmortization
}

// GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciationAndAmortizationOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		return nil, false
	}
	return o.DepreciationAndAmortization, true
}

// HasDepreciationAndAmortization returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDepreciationAndAmortization() bool {
	if o != nil && !IsNil(o.DepreciationAndAmortization) {
		return true
	}

	return false
}

// SetDepreciationAndAmortization gets a reference to the given float64 and assigns it to the DepreciationAndAmortization field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDepreciationAndAmortization(v float64) {
	o.DepreciationAndAmortization = &v
}

// GetAmortizationCashFlow returns the AmortizationCashFlow field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationCashFlow() float64 {
	if o == nil || IsNil(o.AmortizationCashFlow) {
		var ret float64
		return ret
	}
	return *o.AmortizationCashFlow
}

// GetAmortizationCashFlowOk returns a tuple with the AmortizationCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.AmortizationCashFlow) {
		return nil, false
	}
	return o.AmortizationCashFlow, true
}

// HasAmortizationCashFlow returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasAmortizationCashFlow() bool {
	if o != nil && !IsNil(o.AmortizationCashFlow) {
		return true
	}

	return false
}

// SetAmortizationCashFlow gets a reference to the given float64 and assigns it to the AmortizationCashFlow field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetAmortizationCashFlow(v float64) {
	o.AmortizationCashFlow = &v
}

// GetAmortizationOfIntangibles returns the AmortizationOfIntangibles field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationOfIntangibles() float64 {
	if o == nil || IsNil(o.AmortizationOfIntangibles) {
		var ret float64
		return ret
	}
	return *o.AmortizationOfIntangibles
}

// GetAmortizationOfIntangiblesOk returns a tuple with the AmortizationOfIntangibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetAmortizationOfIntangiblesOk() (*float64, bool) {
	if o == nil || IsNil(o.AmortizationOfIntangibles) {
		return nil, false
	}
	return o.AmortizationOfIntangibles, true
}

// HasAmortizationOfIntangibles returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasAmortizationOfIntangibles() bool {
	if o != nil && !IsNil(o.AmortizationOfIntangibles) {
		return true
	}

	return false
}

// SetAmortizationOfIntangibles gets a reference to the given float64 and assigns it to the AmortizationOfIntangibles field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetAmortizationOfIntangibles(v float64) {
	o.AmortizationOfIntangibles = &v
}

// GetDepreciation returns the Depreciation field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciation() float64 {
	if o == nil || IsNil(o.Depreciation) {
		var ret float64
		return ret
	}
	return *o.Depreciation
}

// GetDepreciationOk returns a tuple with the Depreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.Depreciation) {
		return nil, false
	}
	return o.Depreciation, true
}

// HasDepreciation returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasDepreciation() bool {
	if o != nil && !IsNil(o.Depreciation) {
		return true
	}

	return false
}

// SetDepreciation gets a reference to the given float64 and assigns it to the Depreciation field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetDepreciation(v float64) {
	o.Depreciation = &v
}

// GetOperatingGainsLosses returns the OperatingGainsLosses field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOperatingGainsLosses() float64 {
	if o == nil || IsNil(o.OperatingGainsLosses) {
		var ret float64
		return ret
	}
	return *o.OperatingGainsLosses
}

// GetOperatingGainsLossesOk returns a tuple with the OperatingGainsLosses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetOperatingGainsLossesOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingGainsLosses) {
		return nil, false
	}
	return o.OperatingGainsLosses, true
}

// HasOperatingGainsLosses returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasOperatingGainsLosses() bool {
	if o != nil && !IsNil(o.OperatingGainsLosses) {
		return true
	}

	return false
}

// SetOperatingGainsLosses gets a reference to the given float64 and assigns it to the OperatingGainsLosses field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetOperatingGainsLosses(v float64) {
	o.OperatingGainsLosses = &v
}

// GetPensionAndEmployeeBenefitExpense returns the PensionAndEmployeeBenefitExpense field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetPensionAndEmployeeBenefitExpense() float64 {
	if o == nil || IsNil(o.PensionAndEmployeeBenefitExpense) {
		var ret float64
		return ret
	}
	return *o.PensionAndEmployeeBenefitExpense
}

// GetPensionAndEmployeeBenefitExpenseOk returns a tuple with the PensionAndEmployeeBenefitExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetPensionAndEmployeeBenefitExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.PensionAndEmployeeBenefitExpense) {
		return nil, false
	}
	return o.PensionAndEmployeeBenefitExpense, true
}

// HasPensionAndEmployeeBenefitExpense returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasPensionAndEmployeeBenefitExpense() bool {
	if o != nil && !IsNil(o.PensionAndEmployeeBenefitExpense) {
		return true
	}

	return false
}

// SetPensionAndEmployeeBenefitExpense gets a reference to the given float64 and assigns it to the PensionAndEmployeeBenefitExpense field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetPensionAndEmployeeBenefitExpense(v float64) {
	o.PensionAndEmployeeBenefitExpense = &v
}

// GetEarningsLossesFromEquityInvestments returns the EarningsLossesFromEquityInvestments field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetEarningsLossesFromEquityInvestments() float64 {
	if o == nil || IsNil(o.EarningsLossesFromEquityInvestments) {
		var ret float64
		return ret
	}
	return *o.EarningsLossesFromEquityInvestments
}

// GetEarningsLossesFromEquityInvestmentsOk returns a tuple with the EarningsLossesFromEquityInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetEarningsLossesFromEquityInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.EarningsLossesFromEquityInvestments) {
		return nil, false
	}
	return o.EarningsLossesFromEquityInvestments, true
}

// HasEarningsLossesFromEquityInvestments returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasEarningsLossesFromEquityInvestments() bool {
	if o != nil && !IsNil(o.EarningsLossesFromEquityInvestments) {
		return true
	}

	return false
}

// SetEarningsLossesFromEquityInvestments gets a reference to the given float64 and assigns it to the EarningsLossesFromEquityInvestments field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetEarningsLossesFromEquityInvestments(v float64) {
	o.EarningsLossesFromEquityInvestments = &v
}

// GetGainLossOnInvestmentSecurities returns the GainLossOnInvestmentSecurities field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnInvestmentSecurities() float64 {
	if o == nil || IsNil(o.GainLossOnInvestmentSecurities) {
		var ret float64
		return ret
	}
	return *o.GainLossOnInvestmentSecurities
}

// GetGainLossOnInvestmentSecuritiesOk returns a tuple with the GainLossOnInvestmentSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnInvestmentSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.GainLossOnInvestmentSecurities) {
		return nil, false
	}
	return o.GainLossOnInvestmentSecurities, true
}

// HasGainLossOnInvestmentSecurities returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasGainLossOnInvestmentSecurities() bool {
	if o != nil && !IsNil(o.GainLossOnInvestmentSecurities) {
		return true
	}

	return false
}

// SetGainLossOnInvestmentSecurities gets a reference to the given float64 and assigns it to the GainLossOnInvestmentSecurities field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetGainLossOnInvestmentSecurities(v float64) {
	o.GainLossOnInvestmentSecurities = &v
}

// GetNetForeignCurrencyExchangeGainLoss returns the NetForeignCurrencyExchangeGainLoss field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetNetForeignCurrencyExchangeGainLoss() float64 {
	if o == nil || IsNil(o.NetForeignCurrencyExchangeGainLoss) {
		var ret float64
		return ret
	}
	return *o.NetForeignCurrencyExchangeGainLoss
}

// GetNetForeignCurrencyExchangeGainLossOk returns a tuple with the NetForeignCurrencyExchangeGainLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetNetForeignCurrencyExchangeGainLossOk() (*float64, bool) {
	if o == nil || IsNil(o.NetForeignCurrencyExchangeGainLoss) {
		return nil, false
	}
	return o.NetForeignCurrencyExchangeGainLoss, true
}

// HasNetForeignCurrencyExchangeGainLoss returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasNetForeignCurrencyExchangeGainLoss() bool {
	if o != nil && !IsNil(o.NetForeignCurrencyExchangeGainLoss) {
		return true
	}

	return false
}

// SetNetForeignCurrencyExchangeGainLoss gets a reference to the given float64 and assigns it to the NetForeignCurrencyExchangeGainLoss field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetNetForeignCurrencyExchangeGainLoss(v float64) {
	o.NetForeignCurrencyExchangeGainLoss = &v
}

// GetGainLossOnSaleOfPpe returns the GainLossOnSaleOfPpe field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnSaleOfPpe() float64 {
	if o == nil || IsNil(o.GainLossOnSaleOfPpe) {
		var ret float64
		return ret
	}
	return *o.GainLossOnSaleOfPpe
}

// GetGainLossOnSaleOfPpeOk returns a tuple with the GainLossOnSaleOfPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnSaleOfPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.GainLossOnSaleOfPpe) {
		return nil, false
	}
	return o.GainLossOnSaleOfPpe, true
}

// HasGainLossOnSaleOfPpe returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasGainLossOnSaleOfPpe() bool {
	if o != nil && !IsNil(o.GainLossOnSaleOfPpe) {
		return true
	}

	return false
}

// SetGainLossOnSaleOfPpe gets a reference to the given float64 and assigns it to the GainLossOnSaleOfPpe field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetGainLossOnSaleOfPpe(v float64) {
	o.GainLossOnSaleOfPpe = &v
}

// GetGainLossOnSaleOfBusiness returns the GainLossOnSaleOfBusiness field value if set, zero value otherwise.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnSaleOfBusiness() float64 {
	if o == nil || IsNil(o.GainLossOnSaleOfBusiness) {
		var ret float64
		return ret
	}
	return *o.GainLossOnSaleOfBusiness
}

// GetGainLossOnSaleOfBusinessOk returns a tuple with the GainLossOnSaleOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) GetGainLossOnSaleOfBusinessOk() (*float64, bool) {
	if o == nil || IsNil(o.GainLossOnSaleOfBusiness) {
		return nil, false
	}
	return o.GainLossOnSaleOfBusiness, true
}

// HasGainLossOnSaleOfBusiness returns a boolean if a field has been set.
func (o *CashFlowDataCashFlowFromOperatingActivities) HasGainLossOnSaleOfBusiness() bool {
	if o != nil && !IsNil(o.GainLossOnSaleOfBusiness) {
		return true
	}

	return false
}

// SetGainLossOnSaleOfBusiness gets a reference to the given float64 and assigns it to the GainLossOnSaleOfBusiness field.
func (o *CashFlowDataCashFlowFromOperatingActivities) SetGainLossOnSaleOfBusiness(v float64) {
	o.GainLossOnSaleOfBusiness = &v
}

func (o CashFlowDataCashFlowFromOperatingActivities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataCashFlowFromOperatingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetIncomeFromContinuingOperations) {
		toSerialize["net_income_from_continuing_operations"] = o.NetIncomeFromContinuingOperations
	}
	if !IsNil(o.OperatingCashFlow) {
		toSerialize["operating_cash_flow"] = o.OperatingCashFlow
	}
	if !IsNil(o.CashFlowFromContinuingOperatingActivities) {
		toSerialize["cash_flow_from_continuing_operating_activities"] = o.CashFlowFromContinuingOperatingActivities
	}
	if !IsNil(o.CashFromDiscontinuedOperatingActivities) {
		toSerialize["cash_from_discontinued_operating_activities"] = o.CashFromDiscontinuedOperatingActivities
	}
	if !IsNil(o.CashFlowFromDiscontinuedOperation) {
		toSerialize["cash_flow_from_discontinued_operation"] = o.CashFlowFromDiscontinuedOperation
	}
	if !IsNil(o.FreeCashFlow) {
		toSerialize["free_cash_flow"] = o.FreeCashFlow
	}
	if !IsNil(o.CashFlowsFromUsedInOperatingActivitiesDirect) {
		toSerialize["cash_flows_from_used_in_operating_activities_direct"] = o.CashFlowsFromUsedInOperatingActivitiesDirect
	}
	if !IsNil(o.TaxesRefundPaid) {
		toSerialize["taxes_refund_paid"] = o.TaxesRefundPaid
	}
	if !IsNil(o.TaxesRefundPaidDirect) {
		toSerialize["taxes_refund_paid_direct"] = o.TaxesRefundPaidDirect
	}
	if !IsNil(o.InterestReceived) {
		toSerialize["interest_received"] = o.InterestReceived
	}
	if !IsNil(o.InterestReceivedDirect) {
		toSerialize["interest_received_direct"] = o.InterestReceivedDirect
	}
	if !IsNil(o.InterestPaid) {
		toSerialize["interest_paid"] = o.InterestPaid
	}
	if !IsNil(o.InterestPaidDirect) {
		toSerialize["interest_paid_direct"] = o.InterestPaidDirect
	}
	if !IsNil(o.DividendReceived) {
		toSerialize["dividend_received"] = o.DividendReceived
	}
	if !IsNil(o.DividendReceivedDirect) {
		toSerialize["dividend_received_direct"] = o.DividendReceivedDirect
	}
	if !IsNil(o.DividendPaid) {
		toSerialize["dividend_paid"] = o.DividendPaid
	}
	if !IsNil(o.DividendPaidDirect) {
		toSerialize["dividend_paid_direct"] = o.DividendPaidDirect
	}
	if !IsNil(o.ChangeInWorkingCapital) {
		toSerialize["change_in_working_capital"] = o.ChangeInWorkingCapital
	}
	if !IsNil(o.ChangeInOtherWorkingCapital) {
		toSerialize["change_in_other_working_capital"] = o.ChangeInOtherWorkingCapital
	}
	if !IsNil(o.ChangeInReceivables) {
		toSerialize["change_in_receivables"] = o.ChangeInReceivables
	}
	if !IsNil(o.ChangesInAccountReceivables) {
		toSerialize["changes_in_account_receivables"] = o.ChangesInAccountReceivables
	}
	if !IsNil(o.ChangeInPayablesAndAccruedExpense) {
		toSerialize["change_in_payables_and_accrued_expense"] = o.ChangeInPayablesAndAccruedExpense
	}
	if !IsNil(o.ChangeInAccruedExpense) {
		toSerialize["change_in_accrued_expense"] = o.ChangeInAccruedExpense
	}
	if !IsNil(o.ChangeInPayable) {
		toSerialize["change_in_payable"] = o.ChangeInPayable
	}
	if !IsNil(o.ChangeInDividendPayable) {
		toSerialize["change_in_dividend_payable"] = o.ChangeInDividendPayable
	}
	if !IsNil(o.ChangeInAccountPayable) {
		toSerialize["change_in_account_payable"] = o.ChangeInAccountPayable
	}
	if !IsNil(o.ChangeInTaxPayable) {
		toSerialize["change_in_tax_payable"] = o.ChangeInTaxPayable
	}
	if !IsNil(o.ChangeInIncomeTaxPayable) {
		toSerialize["change_in_income_tax_payable"] = o.ChangeInIncomeTaxPayable
	}
	if !IsNil(o.ChangeInInterestPayable) {
		toSerialize["change_in_interest_payable"] = o.ChangeInInterestPayable
	}
	if !IsNil(o.ChangeInOtherCurrentLiabilities) {
		toSerialize["change_in_other_current_liabilities"] = o.ChangeInOtherCurrentLiabilities
	}
	if !IsNil(o.ChangeInOtherCurrentAssets) {
		toSerialize["change_in_other_current_assets"] = o.ChangeInOtherCurrentAssets
	}
	if !IsNil(o.ChangeInInventory) {
		toSerialize["change_in_inventory"] = o.ChangeInInventory
	}
	if !IsNil(o.ChangeInPrepaidAssets) {
		toSerialize["change_in_prepaid_assets"] = o.ChangeInPrepaidAssets
	}
	if !IsNil(o.OtherNonCashItems) {
		toSerialize["other_non_cash_items"] = o.OtherNonCashItems
	}
	if !IsNil(o.ExcessTaxBenefitFromStockBasedCompensation) {
		toSerialize["excess_tax_benefit_from_stock_based_compensation"] = o.ExcessTaxBenefitFromStockBasedCompensation
	}
	if !IsNil(o.StockBasedCompensation) {
		toSerialize["stock_based_compensation"] = o.StockBasedCompensation
	}
	if !IsNil(o.UnrealizedGainLossOnInvestmentSecurities) {
		toSerialize["unrealized_gain_loss_on_investment_securities"] = o.UnrealizedGainLossOnInvestmentSecurities
	}
	if !IsNil(o.ProvisionAndWriteOffOfAssets) {
		toSerialize["provision_and_write_off_of_assets"] = o.ProvisionAndWriteOffOfAssets
	}
	if !IsNil(o.AssetImpairmentCharge) {
		toSerialize["asset_impairment_charge"] = o.AssetImpairmentCharge
	}
	if !IsNil(o.AmortizationOfSecurities) {
		toSerialize["amortization_of_securities"] = o.AmortizationOfSecurities
	}
	if !IsNil(o.DeferredTax) {
		toSerialize["deferred_tax"] = o.DeferredTax
	}
	if !IsNil(o.DeferredIncomeTax) {
		toSerialize["deferred_income_tax"] = o.DeferredIncomeTax
	}
	if !IsNil(o.DepreciationAmortizationDepletion) {
		toSerialize["depreciation_amortization_depletion"] = o.DepreciationAmortizationDepletion
	}
	if !IsNil(o.Depletion) {
		toSerialize["depletion"] = o.Depletion
	}
	if !IsNil(o.DepreciationAndAmortization) {
		toSerialize["depreciation_and_amortization"] = o.DepreciationAndAmortization
	}
	if !IsNil(o.AmortizationCashFlow) {
		toSerialize["amortization_cash_flow"] = o.AmortizationCashFlow
	}
	if !IsNil(o.AmortizationOfIntangibles) {
		toSerialize["amortization_of_intangibles"] = o.AmortizationOfIntangibles
	}
	if !IsNil(o.Depreciation) {
		toSerialize["depreciation"] = o.Depreciation
	}
	if !IsNil(o.OperatingGainsLosses) {
		toSerialize["operating_gains_losses"] = o.OperatingGainsLosses
	}
	if !IsNil(o.PensionAndEmployeeBenefitExpense) {
		toSerialize["pension_and_employee_benefit_expense"] = o.PensionAndEmployeeBenefitExpense
	}
	if !IsNil(o.EarningsLossesFromEquityInvestments) {
		toSerialize["earnings_losses_from_equity_investments"] = o.EarningsLossesFromEquityInvestments
	}
	if !IsNil(o.GainLossOnInvestmentSecurities) {
		toSerialize["gain_loss_on_investment_securities"] = o.GainLossOnInvestmentSecurities
	}
	if !IsNil(o.NetForeignCurrencyExchangeGainLoss) {
		toSerialize["net_foreign_currency_exchange_gain_loss"] = o.NetForeignCurrencyExchangeGainLoss
	}
	if !IsNil(o.GainLossOnSaleOfPpe) {
		toSerialize["gain_loss_on_sale_of_ppe"] = o.GainLossOnSaleOfPpe
	}
	if !IsNil(o.GainLossOnSaleOfBusiness) {
		toSerialize["gain_loss_on_sale_of_business"] = o.GainLossOnSaleOfBusiness
	}
	return toSerialize, nil
}

type NullableCashFlowDataCashFlowFromOperatingActivities struct {
	value *CashFlowDataCashFlowFromOperatingActivities
	isSet bool
}

func (v NullableCashFlowDataCashFlowFromOperatingActivities) Get() *CashFlowDataCashFlowFromOperatingActivities {
	return v.value
}

func (v *NullableCashFlowDataCashFlowFromOperatingActivities) Set(val *CashFlowDataCashFlowFromOperatingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataCashFlowFromOperatingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataCashFlowFromOperatingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataCashFlowFromOperatingActivities(val *CashFlowDataCashFlowFromOperatingActivities) *NullableCashFlowDataCashFlowFromOperatingActivities {
	return &NullableCashFlowDataCashFlowFromOperatingActivities{value: val, isSet: true}
}

func (v NullableCashFlowDataCashFlowFromOperatingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataCashFlowFromOperatingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


