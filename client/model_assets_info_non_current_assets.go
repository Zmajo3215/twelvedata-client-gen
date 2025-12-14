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

// checks if the AssetsInfoNonCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoNonCurrentAssets{}

// AssetsInfoNonCurrentAssets Non-current assets information
type AssetsInfoNonCurrentAssets struct {
	// Total non current assets
	TotalNonCurrentAssets *float64 `json:"total_non_current_assets,omitempty"`
	// Financial assets
	FinancialAssets *float64 `json:"financial_assets,omitempty"`
	// Investments and advances
	InvestmentsAndAdvances *float64 `json:"investments_and_advances,omitempty"`
	// Other investments
	OtherInvestments *float64 `json:"other_investments,omitempty"`
	// Investment in financial assets
	InvestmentInFinancialAssets *float64 `json:"investment_in_financial_assets,omitempty"`
	// Held to maturity securities
	HeldToMaturitySecurities *float64 `json:"held_to_maturity_securities,omitempty"`
	// Available for sale securities
	AvailableForSaleSecurities *float64 `json:"available_for_sale_securities,omitempty"`
	// Financial assets designated as fair value through profit or loss total
	FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal *float64 `json:"financial_assets_designated_as_fair_value_through_profit_or_loss_total,omitempty"`
	// Trading securities
	TradingSecurities *float64 `json:"trading_securities,omitempty"`
	// Long term equity investment
	LongTermEquityInvestment *float64 `json:"long_term_equity_investment,omitempty"`
	// Investments in joint ventures at cost
	InvestmentsInJointVenturesAtCost *float64 `json:"investments_in_joint_ventures_at_cost,omitempty"`
	// Investments in other ventures under equity method
	InvestmentsInOtherVenturesUnderEquityMethod *float64 `json:"investments_in_other_ventures_under_equity_method,omitempty"`
	// Investments in associates at cost
	InvestmentsInAssociatesAtCost *float64 `json:"investments_in_associates_at_cost,omitempty"`
	// Investments in subsidiaries at cost
	InvestmentsInSubsidiariesAtCost *float64 `json:"investments_in_subsidiaries_at_cost,omitempty"`
	// Investment properties
	InvestmentProperties *float64 `json:"investment_properties,omitempty"`
	GoodwillAndOtherIntangibleAssets *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets `json:"goodwill_and_other_intangible_assets,omitempty"`
	// Net ppe
	NetPpe *float64 `json:"net_ppe,omitempty"`
	// Gross ppe
	GrossPpe *float64 `json:"gross_ppe,omitempty"`
	// Accumulated depreciation
	AccumulatedDepreciation *float64 `json:"accumulated_depreciation,omitempty"`
	// Leases
	Leases *float64 `json:"leases,omitempty"`
	// Construction in progress
	ConstructionInProgress *float64 `json:"construction_in_progress,omitempty"`
	// Other properties
	OtherProperties *float64 `json:"other_properties,omitempty"`
	// Machinery furniture equipment
	MachineryFurnitureEquipment *float64 `json:"machinery_furniture_equipment,omitempty"`
	// Buildings and improvements
	BuildingsAndImprovements *float64 `json:"buildings_and_improvements,omitempty"`
	// Land and improvements
	LandAndImprovements *float64 `json:"land_and_improvements,omitempty"`
	// Properties
	Properties *float64 `json:"properties,omitempty"`
	// Non current accounts receivable
	NonCurrentAccountsReceivable *float64 `json:"non_current_accounts_receivable,omitempty"`
	// Non current note receivables
	NonCurrentNoteReceivables *float64 `json:"non_current_note_receivables,omitempty"`
	// Due from related parties non current
	DueFromRelatedPartiesNonCurrent *float64 `json:"due_from_related_parties_non_current,omitempty"`
	// Non current prepaid assets
	NonCurrentPrepaidAssets *float64 `json:"non_current_prepaid_assets,omitempty"`
	// Non current deferred assets
	NonCurrentDeferredAssets *float64 `json:"non_current_deferred_assets,omitempty"`
	// Non current deferred taxes assets
	NonCurrentDeferredTaxesAssets *float64 `json:"non_current_deferred_taxes_assets,omitempty"`
	// Defined pension benefit
	DefinedPensionBenefit *float64 `json:"defined_pension_benefit,omitempty"`
	// Other non current assets
	OtherNonCurrentAssets *float64 `json:"other_non_current_assets,omitempty"`
}

// NewAssetsInfoNonCurrentAssets instantiates a new AssetsInfoNonCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoNonCurrentAssets() *AssetsInfoNonCurrentAssets {
	this := AssetsInfoNonCurrentAssets{}
	return &this
}

// NewAssetsInfoNonCurrentAssetsWithDefaults instantiates a new AssetsInfoNonCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoNonCurrentAssetsWithDefaults() *AssetsInfoNonCurrentAssets {
	this := AssetsInfoNonCurrentAssets{}
	return &this
}

// GetTotalNonCurrentAssets returns the TotalNonCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetTotalNonCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentAssets
}

// GetTotalNonCurrentAssetsOk returns a tuple with the TotalNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetTotalNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		return nil, false
	}
	return o.TotalNonCurrentAssets, true
}

// HasTotalNonCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasTotalNonCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalNonCurrentAssets) {
		return true
	}

	return false
}

// SetTotalNonCurrentAssets gets a reference to the given float64 and assigns it to the TotalNonCurrentAssets field.
func (o *AssetsInfoNonCurrentAssets) SetTotalNonCurrentAssets(v float64) {
	o.TotalNonCurrentAssets = &v
}

// GetFinancialAssets returns the FinancialAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssets() float64 {
	if o == nil || IsNil(o.FinancialAssets) {
		var ret float64
		return ret
	}
	return *o.FinancialAssets
}

// GetFinancialAssetsOk returns a tuple with the FinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssets) {
		return nil, false
	}
	return o.FinancialAssets, true
}

// HasFinancialAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasFinancialAssets() bool {
	if o != nil && !IsNil(o.FinancialAssets) {
		return true
	}

	return false
}

// SetFinancialAssets gets a reference to the given float64 and assigns it to the FinancialAssets field.
func (o *AssetsInfoNonCurrentAssets) SetFinancialAssets(v float64) {
	o.FinancialAssets = &v
}

// GetInvestmentsAndAdvances returns the InvestmentsAndAdvances field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsAndAdvances() float64 {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		var ret float64
		return ret
	}
	return *o.InvestmentsAndAdvances
}

// GetInvestmentsAndAdvancesOk returns a tuple with the InvestmentsAndAdvances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsAndAdvancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		return nil, false
	}
	return o.InvestmentsAndAdvances, true
}

// HasInvestmentsAndAdvances returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsAndAdvances() bool {
	if o != nil && !IsNil(o.InvestmentsAndAdvances) {
		return true
	}

	return false
}

// SetInvestmentsAndAdvances gets a reference to the given float64 and assigns it to the InvestmentsAndAdvances field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsAndAdvances(v float64) {
	o.InvestmentsAndAdvances = &v
}

// GetOtherInvestments returns the OtherInvestments field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherInvestments() float64 {
	if o == nil || IsNil(o.OtherInvestments) {
		var ret float64
		return ret
	}
	return *o.OtherInvestments
}

// GetOtherInvestmentsOk returns a tuple with the OtherInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherInvestments) {
		return nil, false
	}
	return o.OtherInvestments, true
}

// HasOtherInvestments returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherInvestments() bool {
	if o != nil && !IsNil(o.OtherInvestments) {
		return true
	}

	return false
}

// SetOtherInvestments gets a reference to the given float64 and assigns it to the OtherInvestments field.
func (o *AssetsInfoNonCurrentAssets) SetOtherInvestments(v float64) {
	o.OtherInvestments = &v
}

// GetInvestmentInFinancialAssets returns the InvestmentInFinancialAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentInFinancialAssets() float64 {
	if o == nil || IsNil(o.InvestmentInFinancialAssets) {
		var ret float64
		return ret
	}
	return *o.InvestmentInFinancialAssets
}

// GetInvestmentInFinancialAssetsOk returns a tuple with the InvestmentInFinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentInFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentInFinancialAssets) {
		return nil, false
	}
	return o.InvestmentInFinancialAssets, true
}

// HasInvestmentInFinancialAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentInFinancialAssets() bool {
	if o != nil && !IsNil(o.InvestmentInFinancialAssets) {
		return true
	}

	return false
}

// SetInvestmentInFinancialAssets gets a reference to the given float64 and assigns it to the InvestmentInFinancialAssets field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentInFinancialAssets(v float64) {
	o.InvestmentInFinancialAssets = &v
}

// GetHeldToMaturitySecurities returns the HeldToMaturitySecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetHeldToMaturitySecurities() float64 {
	if o == nil || IsNil(o.HeldToMaturitySecurities) {
		var ret float64
		return ret
	}
	return *o.HeldToMaturitySecurities
}

// GetHeldToMaturitySecuritiesOk returns a tuple with the HeldToMaturitySecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetHeldToMaturitySecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.HeldToMaturitySecurities) {
		return nil, false
	}
	return o.HeldToMaturitySecurities, true
}

// HasHeldToMaturitySecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasHeldToMaturitySecurities() bool {
	if o != nil && !IsNil(o.HeldToMaturitySecurities) {
		return true
	}

	return false
}

// SetHeldToMaturitySecurities gets a reference to the given float64 and assigns it to the HeldToMaturitySecurities field.
func (o *AssetsInfoNonCurrentAssets) SetHeldToMaturitySecurities(v float64) {
	o.HeldToMaturitySecurities = &v
}

// GetAvailableForSaleSecurities returns the AvailableForSaleSecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetAvailableForSaleSecurities() float64 {
	if o == nil || IsNil(o.AvailableForSaleSecurities) {
		var ret float64
		return ret
	}
	return *o.AvailableForSaleSecurities
}

// GetAvailableForSaleSecuritiesOk returns a tuple with the AvailableForSaleSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetAvailableForSaleSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.AvailableForSaleSecurities) {
		return nil, false
	}
	return o.AvailableForSaleSecurities, true
}

// HasAvailableForSaleSecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasAvailableForSaleSecurities() bool {
	if o != nil && !IsNil(o.AvailableForSaleSecurities) {
		return true
	}

	return false
}

// SetAvailableForSaleSecurities gets a reference to the given float64 and assigns it to the AvailableForSaleSecurities field.
func (o *AssetsInfoNonCurrentAssets) SetAvailableForSaleSecurities(v float64) {
	o.AvailableForSaleSecurities = &v
}

// GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal returns the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal() float64 {
	if o == nil || IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		var ret float64
		return ret
	}
	return *o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal
}

// GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotalOk returns a tuple with the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		return nil, false
	}
	return o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal, true
}

// HasFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal() bool {
	if o != nil && !IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		return true
	}

	return false
}

// SetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal gets a reference to the given float64 and assigns it to the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field.
func (o *AssetsInfoNonCurrentAssets) SetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal(v float64) {
	o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal = &v
}

// GetTradingSecurities returns the TradingSecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetTradingSecurities() float64 {
	if o == nil || IsNil(o.TradingSecurities) {
		var ret float64
		return ret
	}
	return *o.TradingSecurities
}

// GetTradingSecuritiesOk returns a tuple with the TradingSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetTradingSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TradingSecurities) {
		return nil, false
	}
	return o.TradingSecurities, true
}

// HasTradingSecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasTradingSecurities() bool {
	if o != nil && !IsNil(o.TradingSecurities) {
		return true
	}

	return false
}

// SetTradingSecurities gets a reference to the given float64 and assigns it to the TradingSecurities field.
func (o *AssetsInfoNonCurrentAssets) SetTradingSecurities(v float64) {
	o.TradingSecurities = &v
}

// GetLongTermEquityInvestment returns the LongTermEquityInvestment field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLongTermEquityInvestment() float64 {
	if o == nil || IsNil(o.LongTermEquityInvestment) {
		var ret float64
		return ret
	}
	return *o.LongTermEquityInvestment
}

// GetLongTermEquityInvestmentOk returns a tuple with the LongTermEquityInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLongTermEquityInvestmentOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermEquityInvestment) {
		return nil, false
	}
	return o.LongTermEquityInvestment, true
}

// HasLongTermEquityInvestment returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLongTermEquityInvestment() bool {
	if o != nil && !IsNil(o.LongTermEquityInvestment) {
		return true
	}

	return false
}

// SetLongTermEquityInvestment gets a reference to the given float64 and assigns it to the LongTermEquityInvestment field.
func (o *AssetsInfoNonCurrentAssets) SetLongTermEquityInvestment(v float64) {
	o.LongTermEquityInvestment = &v
}

// GetInvestmentsInJointVenturesAtCost returns the InvestmentsInJointVenturesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInJointVenturesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInJointVenturesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInJointVenturesAtCost
}

// GetInvestmentsInJointVenturesAtCostOk returns a tuple with the InvestmentsInJointVenturesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInJointVenturesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInJointVenturesAtCost) {
		return nil, false
	}
	return o.InvestmentsInJointVenturesAtCost, true
}

// HasInvestmentsInJointVenturesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInJointVenturesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInJointVenturesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInJointVenturesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInJointVenturesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInJointVenturesAtCost(v float64) {
	o.InvestmentsInJointVenturesAtCost = &v
}

// GetInvestmentsInOtherVenturesUnderEquityMethod returns the InvestmentsInOtherVenturesUnderEquityMethod field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInOtherVenturesUnderEquityMethod() float64 {
	if o == nil || IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInOtherVenturesUnderEquityMethod
}

// GetInvestmentsInOtherVenturesUnderEquityMethodOk returns a tuple with the InvestmentsInOtherVenturesUnderEquityMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInOtherVenturesUnderEquityMethodOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		return nil, false
	}
	return o.InvestmentsInOtherVenturesUnderEquityMethod, true
}

// HasInvestmentsInOtherVenturesUnderEquityMethod returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInOtherVenturesUnderEquityMethod() bool {
	if o != nil && !IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		return true
	}

	return false
}

// SetInvestmentsInOtherVenturesUnderEquityMethod gets a reference to the given float64 and assigns it to the InvestmentsInOtherVenturesUnderEquityMethod field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInOtherVenturesUnderEquityMethod(v float64) {
	o.InvestmentsInOtherVenturesUnderEquityMethod = &v
}

// GetInvestmentsInAssociatesAtCost returns the InvestmentsInAssociatesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInAssociatesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInAssociatesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInAssociatesAtCost
}

// GetInvestmentsInAssociatesAtCostOk returns a tuple with the InvestmentsInAssociatesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInAssociatesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInAssociatesAtCost) {
		return nil, false
	}
	return o.InvestmentsInAssociatesAtCost, true
}

// HasInvestmentsInAssociatesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInAssociatesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInAssociatesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInAssociatesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInAssociatesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInAssociatesAtCost(v float64) {
	o.InvestmentsInAssociatesAtCost = &v
}

// GetInvestmentsInSubsidiariesAtCost returns the InvestmentsInSubsidiariesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInSubsidiariesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInSubsidiariesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInSubsidiariesAtCost
}

// GetInvestmentsInSubsidiariesAtCostOk returns a tuple with the InvestmentsInSubsidiariesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInSubsidiariesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInSubsidiariesAtCost) {
		return nil, false
	}
	return o.InvestmentsInSubsidiariesAtCost, true
}

// HasInvestmentsInSubsidiariesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInSubsidiariesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInSubsidiariesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInSubsidiariesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInSubsidiariesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInSubsidiariesAtCost(v float64) {
	o.InvestmentsInSubsidiariesAtCost = &v
}

// GetInvestmentProperties returns the InvestmentProperties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentProperties() float64 {
	if o == nil || IsNil(o.InvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.InvestmentProperties
}

// GetInvestmentPropertiesOk returns a tuple with the InvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentProperties) {
		return nil, false
	}
	return o.InvestmentProperties, true
}

// HasInvestmentProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentProperties() bool {
	if o != nil && !IsNil(o.InvestmentProperties) {
		return true
	}

	return false
}

// SetInvestmentProperties gets a reference to the given float64 and assigns it to the InvestmentProperties field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentProperties(v float64) {
	o.InvestmentProperties = &v
}

// GetGoodwillAndOtherIntangibleAssets returns the GoodwillAndOtherIntangibleAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetGoodwillAndOtherIntangibleAssets() AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	if o == nil || IsNil(o.GoodwillAndOtherIntangibleAssets) {
		var ret AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets
		return ret
	}
	return *o.GoodwillAndOtherIntangibleAssets
}

// GetGoodwillAndOtherIntangibleAssetsOk returns a tuple with the GoodwillAndOtherIntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetGoodwillAndOtherIntangibleAssetsOk() (*AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets, bool) {
	if o == nil || IsNil(o.GoodwillAndOtherIntangibleAssets) {
		return nil, false
	}
	return o.GoodwillAndOtherIntangibleAssets, true
}

// HasGoodwillAndOtherIntangibleAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasGoodwillAndOtherIntangibleAssets() bool {
	if o != nil && !IsNil(o.GoodwillAndOtherIntangibleAssets) {
		return true
	}

	return false
}

// SetGoodwillAndOtherIntangibleAssets gets a reference to the given AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets and assigns it to the GoodwillAndOtherIntangibleAssets field.
func (o *AssetsInfoNonCurrentAssets) SetGoodwillAndOtherIntangibleAssets(v AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) {
	o.GoodwillAndOtherIntangibleAssets = &v
}

// GetNetPpe returns the NetPpe field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNetPpe() float64 {
	if o == nil || IsNil(o.NetPpe) {
		var ret float64
		return ret
	}
	return *o.NetPpe
}

// GetNetPpeOk returns a tuple with the NetPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNetPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetPpe) {
		return nil, false
	}
	return o.NetPpe, true
}

// HasNetPpe returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNetPpe() bool {
	if o != nil && !IsNil(o.NetPpe) {
		return true
	}

	return false
}

// SetNetPpe gets a reference to the given float64 and assigns it to the NetPpe field.
func (o *AssetsInfoNonCurrentAssets) SetNetPpe(v float64) {
	o.NetPpe = &v
}

// GetGrossPpe returns the GrossPpe field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetGrossPpe() float64 {
	if o == nil || IsNil(o.GrossPpe) {
		var ret float64
		return ret
	}
	return *o.GrossPpe
}

// GetGrossPpeOk returns a tuple with the GrossPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetGrossPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossPpe) {
		return nil, false
	}
	return o.GrossPpe, true
}

// HasGrossPpe returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasGrossPpe() bool {
	if o != nil && !IsNil(o.GrossPpe) {
		return true
	}

	return false
}

// SetGrossPpe gets a reference to the given float64 and assigns it to the GrossPpe field.
func (o *AssetsInfoNonCurrentAssets) SetGrossPpe(v float64) {
	o.GrossPpe = &v
}

// GetAccumulatedDepreciation returns the AccumulatedDepreciation field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetAccumulatedDepreciation() float64 {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		var ret float64
		return ret
	}
	return *o.AccumulatedDepreciation
}

// GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetAccumulatedDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		return nil, false
	}
	return o.AccumulatedDepreciation, true
}

// HasAccumulatedDepreciation returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasAccumulatedDepreciation() bool {
	if o != nil && !IsNil(o.AccumulatedDepreciation) {
		return true
	}

	return false
}

// SetAccumulatedDepreciation gets a reference to the given float64 and assigns it to the AccumulatedDepreciation field.
func (o *AssetsInfoNonCurrentAssets) SetAccumulatedDepreciation(v float64) {
	o.AccumulatedDepreciation = &v
}

// GetLeases returns the Leases field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLeases() float64 {
	if o == nil || IsNil(o.Leases) {
		var ret float64
		return ret
	}
	return *o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLeasesOk() (*float64, bool) {
	if o == nil || IsNil(o.Leases) {
		return nil, false
	}
	return o.Leases, true
}

// HasLeases returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLeases() bool {
	if o != nil && !IsNil(o.Leases) {
		return true
	}

	return false
}

// SetLeases gets a reference to the given float64 and assigns it to the Leases field.
func (o *AssetsInfoNonCurrentAssets) SetLeases(v float64) {
	o.Leases = &v
}

// GetConstructionInProgress returns the ConstructionInProgress field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetConstructionInProgress() float64 {
	if o == nil || IsNil(o.ConstructionInProgress) {
		var ret float64
		return ret
	}
	return *o.ConstructionInProgress
}

// GetConstructionInProgressOk returns a tuple with the ConstructionInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetConstructionInProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.ConstructionInProgress) {
		return nil, false
	}
	return o.ConstructionInProgress, true
}

// HasConstructionInProgress returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasConstructionInProgress() bool {
	if o != nil && !IsNil(o.ConstructionInProgress) {
		return true
	}

	return false
}

// SetConstructionInProgress gets a reference to the given float64 and assigns it to the ConstructionInProgress field.
func (o *AssetsInfoNonCurrentAssets) SetConstructionInProgress(v float64) {
	o.ConstructionInProgress = &v
}

// GetOtherProperties returns the OtherProperties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherProperties() float64 {
	if o == nil || IsNil(o.OtherProperties) {
		var ret float64
		return ret
	}
	return *o.OtherProperties
}

// GetOtherPropertiesOk returns a tuple with the OtherProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherProperties) {
		return nil, false
	}
	return o.OtherProperties, true
}

// HasOtherProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherProperties() bool {
	if o != nil && !IsNil(o.OtherProperties) {
		return true
	}

	return false
}

// SetOtherProperties gets a reference to the given float64 and assigns it to the OtherProperties field.
func (o *AssetsInfoNonCurrentAssets) SetOtherProperties(v float64) {
	o.OtherProperties = &v
}

// GetMachineryFurnitureEquipment returns the MachineryFurnitureEquipment field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetMachineryFurnitureEquipment() float64 {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		var ret float64
		return ret
	}
	return *o.MachineryFurnitureEquipment
}

// GetMachineryFurnitureEquipmentOk returns a tuple with the MachineryFurnitureEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetMachineryFurnitureEquipmentOk() (*float64, bool) {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		return nil, false
	}
	return o.MachineryFurnitureEquipment, true
}

// HasMachineryFurnitureEquipment returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasMachineryFurnitureEquipment() bool {
	if o != nil && !IsNil(o.MachineryFurnitureEquipment) {
		return true
	}

	return false
}

// SetMachineryFurnitureEquipment gets a reference to the given float64 and assigns it to the MachineryFurnitureEquipment field.
func (o *AssetsInfoNonCurrentAssets) SetMachineryFurnitureEquipment(v float64) {
	o.MachineryFurnitureEquipment = &v
}

// GetBuildingsAndImprovements returns the BuildingsAndImprovements field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetBuildingsAndImprovements() float64 {
	if o == nil || IsNil(o.BuildingsAndImprovements) {
		var ret float64
		return ret
	}
	return *o.BuildingsAndImprovements
}

// GetBuildingsAndImprovementsOk returns a tuple with the BuildingsAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetBuildingsAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.BuildingsAndImprovements) {
		return nil, false
	}
	return o.BuildingsAndImprovements, true
}

// HasBuildingsAndImprovements returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasBuildingsAndImprovements() bool {
	if o != nil && !IsNil(o.BuildingsAndImprovements) {
		return true
	}

	return false
}

// SetBuildingsAndImprovements gets a reference to the given float64 and assigns it to the BuildingsAndImprovements field.
func (o *AssetsInfoNonCurrentAssets) SetBuildingsAndImprovements(v float64) {
	o.BuildingsAndImprovements = &v
}

// GetLandAndImprovements returns the LandAndImprovements field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLandAndImprovements() float64 {
	if o == nil || IsNil(o.LandAndImprovements) {
		var ret float64
		return ret
	}
	return *o.LandAndImprovements
}

// GetLandAndImprovementsOk returns a tuple with the LandAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLandAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.LandAndImprovements) {
		return nil, false
	}
	return o.LandAndImprovements, true
}

// HasLandAndImprovements returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLandAndImprovements() bool {
	if o != nil && !IsNil(o.LandAndImprovements) {
		return true
	}

	return false
}

// SetLandAndImprovements gets a reference to the given float64 and assigns it to the LandAndImprovements field.
func (o *AssetsInfoNonCurrentAssets) SetLandAndImprovements(v float64) {
	o.LandAndImprovements = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetProperties() float64 {
	if o == nil || IsNil(o.Properties) {
		var ret float64
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given float64 and assigns it to the Properties field.
func (o *AssetsInfoNonCurrentAssets) SetProperties(v float64) {
	o.Properties = &v
}

// GetNonCurrentAccountsReceivable returns the NonCurrentAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentAccountsReceivable() float64 {
	if o == nil || IsNil(o.NonCurrentAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.NonCurrentAccountsReceivable
}

// GetNonCurrentAccountsReceivableOk returns a tuple with the NonCurrentAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentAccountsReceivable) {
		return nil, false
	}
	return o.NonCurrentAccountsReceivable, true
}

// HasNonCurrentAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentAccountsReceivable() bool {
	if o != nil && !IsNil(o.NonCurrentAccountsReceivable) {
		return true
	}

	return false
}

// SetNonCurrentAccountsReceivable gets a reference to the given float64 and assigns it to the NonCurrentAccountsReceivable field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentAccountsReceivable(v float64) {
	o.NonCurrentAccountsReceivable = &v
}

// GetNonCurrentNoteReceivables returns the NonCurrentNoteReceivables field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentNoteReceivables() float64 {
	if o == nil || IsNil(o.NonCurrentNoteReceivables) {
		var ret float64
		return ret
	}
	return *o.NonCurrentNoteReceivables
}

// GetNonCurrentNoteReceivablesOk returns a tuple with the NonCurrentNoteReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentNoteReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentNoteReceivables) {
		return nil, false
	}
	return o.NonCurrentNoteReceivables, true
}

// HasNonCurrentNoteReceivables returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentNoteReceivables() bool {
	if o != nil && !IsNil(o.NonCurrentNoteReceivables) {
		return true
	}

	return false
}

// SetNonCurrentNoteReceivables gets a reference to the given float64 and assigns it to the NonCurrentNoteReceivables field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentNoteReceivables(v float64) {
	o.NonCurrentNoteReceivables = &v
}

// GetDueFromRelatedPartiesNonCurrent returns the DueFromRelatedPartiesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetDueFromRelatedPartiesNonCurrent() float64 {
	if o == nil || IsNil(o.DueFromRelatedPartiesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.DueFromRelatedPartiesNonCurrent
}

// GetDueFromRelatedPartiesNonCurrentOk returns a tuple with the DueFromRelatedPartiesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetDueFromRelatedPartiesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueFromRelatedPartiesNonCurrent) {
		return nil, false
	}
	return o.DueFromRelatedPartiesNonCurrent, true
}

// HasDueFromRelatedPartiesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasDueFromRelatedPartiesNonCurrent() bool {
	if o != nil && !IsNil(o.DueFromRelatedPartiesNonCurrent) {
		return true
	}

	return false
}

// SetDueFromRelatedPartiesNonCurrent gets a reference to the given float64 and assigns it to the DueFromRelatedPartiesNonCurrent field.
func (o *AssetsInfoNonCurrentAssets) SetDueFromRelatedPartiesNonCurrent(v float64) {
	o.DueFromRelatedPartiesNonCurrent = &v
}

// GetNonCurrentPrepaidAssets returns the NonCurrentPrepaidAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentPrepaidAssets() float64 {
	if o == nil || IsNil(o.NonCurrentPrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentPrepaidAssets
}

// GetNonCurrentPrepaidAssetsOk returns a tuple with the NonCurrentPrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentPrepaidAssets) {
		return nil, false
	}
	return o.NonCurrentPrepaidAssets, true
}

// HasNonCurrentPrepaidAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentPrepaidAssets() bool {
	if o != nil && !IsNil(o.NonCurrentPrepaidAssets) {
		return true
	}

	return false
}

// SetNonCurrentPrepaidAssets gets a reference to the given float64 and assigns it to the NonCurrentPrepaidAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentPrepaidAssets(v float64) {
	o.NonCurrentPrepaidAssets = &v
}

// GetNonCurrentDeferredAssets returns the NonCurrentDeferredAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredAssets() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredAssets
}

// GetNonCurrentDeferredAssetsOk returns a tuple with the NonCurrentDeferredAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredAssets) {
		return nil, false
	}
	return o.NonCurrentDeferredAssets, true
}

// HasNonCurrentDeferredAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentDeferredAssets() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredAssets) {
		return true
	}

	return false
}

// SetNonCurrentDeferredAssets gets a reference to the given float64 and assigns it to the NonCurrentDeferredAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentDeferredAssets(v float64) {
	o.NonCurrentDeferredAssets = &v
}

// GetNonCurrentDeferredTaxesAssets returns the NonCurrentDeferredTaxesAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredTaxesAssets() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredTaxesAssets
}

// GetNonCurrentDeferredTaxesAssetsOk returns a tuple with the NonCurrentDeferredTaxesAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredTaxesAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesAssets) {
		return nil, false
	}
	return o.NonCurrentDeferredTaxesAssets, true
}

// HasNonCurrentDeferredTaxesAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentDeferredTaxesAssets() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredTaxesAssets) {
		return true
	}

	return false
}

// SetNonCurrentDeferredTaxesAssets gets a reference to the given float64 and assigns it to the NonCurrentDeferredTaxesAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentDeferredTaxesAssets(v float64) {
	o.NonCurrentDeferredTaxesAssets = &v
}

// GetDefinedPensionBenefit returns the DefinedPensionBenefit field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetDefinedPensionBenefit() float64 {
	if o == nil || IsNil(o.DefinedPensionBenefit) {
		var ret float64
		return ret
	}
	return *o.DefinedPensionBenefit
}

// GetDefinedPensionBenefitOk returns a tuple with the DefinedPensionBenefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetDefinedPensionBenefitOk() (*float64, bool) {
	if o == nil || IsNil(o.DefinedPensionBenefit) {
		return nil, false
	}
	return o.DefinedPensionBenefit, true
}

// HasDefinedPensionBenefit returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasDefinedPensionBenefit() bool {
	if o != nil && !IsNil(o.DefinedPensionBenefit) {
		return true
	}

	return false
}

// SetDefinedPensionBenefit gets a reference to the given float64 and assigns it to the DefinedPensionBenefit field.
func (o *AssetsInfoNonCurrentAssets) SetDefinedPensionBenefit(v float64) {
	o.DefinedPensionBenefit = &v
}

// GetOtherNonCurrentAssets returns the OtherNonCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherNonCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentAssets
}

// GetOtherNonCurrentAssetsOk returns a tuple with the OtherNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		return nil, false
	}
	return o.OtherNonCurrentAssets, true
}

// HasOtherNonCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherNonCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherNonCurrentAssets) {
		return true
	}

	return false
}

// SetOtherNonCurrentAssets gets a reference to the given float64 and assigns it to the OtherNonCurrentAssets field.
func (o *AssetsInfoNonCurrentAssets) SetOtherNonCurrentAssets(v float64) {
	o.OtherNonCurrentAssets = &v
}

func (o AssetsInfoNonCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoNonCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalNonCurrentAssets) {
		toSerialize["total_non_current_assets"] = o.TotalNonCurrentAssets
	}
	if !IsNil(o.FinancialAssets) {
		toSerialize["financial_assets"] = o.FinancialAssets
	}
	if !IsNil(o.InvestmentsAndAdvances) {
		toSerialize["investments_and_advances"] = o.InvestmentsAndAdvances
	}
	if !IsNil(o.OtherInvestments) {
		toSerialize["other_investments"] = o.OtherInvestments
	}
	if !IsNil(o.InvestmentInFinancialAssets) {
		toSerialize["investment_in_financial_assets"] = o.InvestmentInFinancialAssets
	}
	if !IsNil(o.HeldToMaturitySecurities) {
		toSerialize["held_to_maturity_securities"] = o.HeldToMaturitySecurities
	}
	if !IsNil(o.AvailableForSaleSecurities) {
		toSerialize["available_for_sale_securities"] = o.AvailableForSaleSecurities
	}
	if !IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		toSerialize["financial_assets_designated_as_fair_value_through_profit_or_loss_total"] = o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal
	}
	if !IsNil(o.TradingSecurities) {
		toSerialize["trading_securities"] = o.TradingSecurities
	}
	if !IsNil(o.LongTermEquityInvestment) {
		toSerialize["long_term_equity_investment"] = o.LongTermEquityInvestment
	}
	if !IsNil(o.InvestmentsInJointVenturesAtCost) {
		toSerialize["investments_in_joint_ventures_at_cost"] = o.InvestmentsInJointVenturesAtCost
	}
	if !IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		toSerialize["investments_in_other_ventures_under_equity_method"] = o.InvestmentsInOtherVenturesUnderEquityMethod
	}
	if !IsNil(o.InvestmentsInAssociatesAtCost) {
		toSerialize["investments_in_associates_at_cost"] = o.InvestmentsInAssociatesAtCost
	}
	if !IsNil(o.InvestmentsInSubsidiariesAtCost) {
		toSerialize["investments_in_subsidiaries_at_cost"] = o.InvestmentsInSubsidiariesAtCost
	}
	if !IsNil(o.InvestmentProperties) {
		toSerialize["investment_properties"] = o.InvestmentProperties
	}
	if !IsNil(o.GoodwillAndOtherIntangibleAssets) {
		toSerialize["goodwill_and_other_intangible_assets"] = o.GoodwillAndOtherIntangibleAssets
	}
	if !IsNil(o.NetPpe) {
		toSerialize["net_ppe"] = o.NetPpe
	}
	if !IsNil(o.GrossPpe) {
		toSerialize["gross_ppe"] = o.GrossPpe
	}
	if !IsNil(o.AccumulatedDepreciation) {
		toSerialize["accumulated_depreciation"] = o.AccumulatedDepreciation
	}
	if !IsNil(o.Leases) {
		toSerialize["leases"] = o.Leases
	}
	if !IsNil(o.ConstructionInProgress) {
		toSerialize["construction_in_progress"] = o.ConstructionInProgress
	}
	if !IsNil(o.OtherProperties) {
		toSerialize["other_properties"] = o.OtherProperties
	}
	if !IsNil(o.MachineryFurnitureEquipment) {
		toSerialize["machinery_furniture_equipment"] = o.MachineryFurnitureEquipment
	}
	if !IsNil(o.BuildingsAndImprovements) {
		toSerialize["buildings_and_improvements"] = o.BuildingsAndImprovements
	}
	if !IsNil(o.LandAndImprovements) {
		toSerialize["land_and_improvements"] = o.LandAndImprovements
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.NonCurrentAccountsReceivable) {
		toSerialize["non_current_accounts_receivable"] = o.NonCurrentAccountsReceivable
	}
	if !IsNil(o.NonCurrentNoteReceivables) {
		toSerialize["non_current_note_receivables"] = o.NonCurrentNoteReceivables
	}
	if !IsNil(o.DueFromRelatedPartiesNonCurrent) {
		toSerialize["due_from_related_parties_non_current"] = o.DueFromRelatedPartiesNonCurrent
	}
	if !IsNil(o.NonCurrentPrepaidAssets) {
		toSerialize["non_current_prepaid_assets"] = o.NonCurrentPrepaidAssets
	}
	if !IsNil(o.NonCurrentDeferredAssets) {
		toSerialize["non_current_deferred_assets"] = o.NonCurrentDeferredAssets
	}
	if !IsNil(o.NonCurrentDeferredTaxesAssets) {
		toSerialize["non_current_deferred_taxes_assets"] = o.NonCurrentDeferredTaxesAssets
	}
	if !IsNil(o.DefinedPensionBenefit) {
		toSerialize["defined_pension_benefit"] = o.DefinedPensionBenefit
	}
	if !IsNil(o.OtherNonCurrentAssets) {
		toSerialize["other_non_current_assets"] = o.OtherNonCurrentAssets
	}
	return toSerialize, nil
}

type NullableAssetsInfoNonCurrentAssets struct {
	value *AssetsInfoNonCurrentAssets
	isSet bool
}

func (v NullableAssetsInfoNonCurrentAssets) Get() *AssetsInfoNonCurrentAssets {
	return v.value
}

func (v *NullableAssetsInfoNonCurrentAssets) Set(val *AssetsInfoNonCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoNonCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoNonCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoNonCurrentAssets(val *AssetsInfoNonCurrentAssets) *NullableAssetsInfoNonCurrentAssets {
	return &NullableAssetsInfoNonCurrentAssets{value: val, isSet: true}
}

func (v NullableAssetsInfoNonCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoNonCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


