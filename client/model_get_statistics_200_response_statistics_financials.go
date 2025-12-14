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

// checks if the GetStatistics200ResponseStatisticsFinancials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsFinancials{}

// GetStatistics200ResponseStatisticsFinancials Financial information of the company
type GetStatistics200ResponseStatisticsFinancials struct {
	// Refers to the completion of the last 12-month accounting period
	FiscalYearEnds *string `json:"fiscal_year_ends,omitempty"`
	// The most recent quarter (MRQ) refers to the fiscal quarter that most recently ended
	MostRecentQuarter *string `json:"most_recent_quarter,omitempty"`
	// The portion of a company's revenue left over after direct costs are subtracted
	GrossMargin *float64 `json:"gross_margin,omitempty"`
	// Refers to gross profit margin. Calculated by dividing net income by sales revenue
	ProfitMargin *float64 `json:"profit_margin,omitempty"`
	// Operating margin is calculated by dividing operating income by net sales
	OperatingMargin *float64 `json:"operating_margin,omitempty"`
	// Return on assets (ROA) is calculated by dividing net income by total assets
	ReturnOnAssetsTtm *float64 `json:"return_on_assets_ttm,omitempty"`
	// Return on equity (ROE) is calculated by dividing net income by average shareholders' equity
	ReturnOnEquityTtm *float64 `json:"return_on_equity_ttm,omitempty"`
	IncomeStatement *GetStatistics200ResponseStatisticsFinancialsIncomeStatement `json:"income_statement,omitempty"`
	BalanceSheet *GetStatistics200ResponseStatisticsFinancialsBalanceSheet `json:"balance_sheet,omitempty"`
	CashFlow *GetStatistics200ResponseStatisticsFinancialsCashFlow `json:"cash_flow,omitempty"`
}

// NewGetStatistics200ResponseStatisticsFinancials instantiates a new GetStatistics200ResponseStatisticsFinancials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsFinancials() *GetStatistics200ResponseStatisticsFinancials {
	this := GetStatistics200ResponseStatisticsFinancials{}
	return &this
}

// NewGetStatistics200ResponseStatisticsFinancialsWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsFinancialsWithDefaults() *GetStatistics200ResponseStatisticsFinancials {
	this := GetStatistics200ResponseStatisticsFinancials{}
	return &this
}

// GetFiscalYearEnds returns the FiscalYearEnds field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEnds() string {
	if o == nil || IsNil(o.FiscalYearEnds) {
		var ret string
		return ret
	}
	return *o.FiscalYearEnds
}

// GetFiscalYearEndsOk returns a tuple with the FiscalYearEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEndsOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalYearEnds) {
		return nil, false
	}
	return o.FiscalYearEnds, true
}

// HasFiscalYearEnds returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasFiscalYearEnds() bool {
	if o != nil && !IsNil(o.FiscalYearEnds) {
		return true
	}

	return false
}

// SetFiscalYearEnds gets a reference to the given string and assigns it to the FiscalYearEnds field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetFiscalYearEnds(v string) {
	o.FiscalYearEnds = &v
}

// GetMostRecentQuarter returns the MostRecentQuarter field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarter() string {
	if o == nil || IsNil(o.MostRecentQuarter) {
		var ret string
		return ret
	}
	return *o.MostRecentQuarter
}

// GetMostRecentQuarterOk returns a tuple with the MostRecentQuarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarterOk() (*string, bool) {
	if o == nil || IsNil(o.MostRecentQuarter) {
		return nil, false
	}
	return o.MostRecentQuarter, true
}

// HasMostRecentQuarter returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasMostRecentQuarter() bool {
	if o != nil && !IsNil(o.MostRecentQuarter) {
		return true
	}

	return false
}

// SetMostRecentQuarter gets a reference to the given string and assigns it to the MostRecentQuarter field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetMostRecentQuarter(v string) {
	o.MostRecentQuarter = &v
}

// GetGrossMargin returns the GrossMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMargin() float64 {
	if o == nil || IsNil(o.GrossMargin) {
		var ret float64
		return ret
	}
	return *o.GrossMargin
}

// GetGrossMarginOk returns a tuple with the GrossMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossMargin) {
		return nil, false
	}
	return o.GrossMargin, true
}

// HasGrossMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasGrossMargin() bool {
	if o != nil && !IsNil(o.GrossMargin) {
		return true
	}

	return false
}

// SetGrossMargin gets a reference to the given float64 and assigns it to the GrossMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetGrossMargin(v float64) {
	o.GrossMargin = &v
}

// GetProfitMargin returns the ProfitMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMargin() float64 {
	if o == nil || IsNil(o.ProfitMargin) {
		var ret float64
		return ret
	}
	return *o.ProfitMargin
}

// GetProfitMarginOk returns a tuple with the ProfitMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.ProfitMargin) {
		return nil, false
	}
	return o.ProfitMargin, true
}

// HasProfitMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasProfitMargin() bool {
	if o != nil && !IsNil(o.ProfitMargin) {
		return true
	}

	return false
}

// SetProfitMargin gets a reference to the given float64 and assigns it to the ProfitMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetProfitMargin(v float64) {
	o.ProfitMargin = &v
}

// GetOperatingMargin returns the OperatingMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMargin() float64 {
	if o == nil || IsNil(o.OperatingMargin) {
		var ret float64
		return ret
	}
	return *o.OperatingMargin
}

// GetOperatingMarginOk returns a tuple with the OperatingMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingMargin) {
		return nil, false
	}
	return o.OperatingMargin, true
}

// HasOperatingMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasOperatingMargin() bool {
	if o != nil && !IsNil(o.OperatingMargin) {
		return true
	}

	return false
}

// SetOperatingMargin gets a reference to the given float64 and assigns it to the OperatingMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetOperatingMargin(v float64) {
	o.OperatingMargin = &v
}

// GetReturnOnAssetsTtm returns the ReturnOnAssetsTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtm() float64 {
	if o == nil || IsNil(o.ReturnOnAssetsTtm) {
		var ret float64
		return ret
	}
	return *o.ReturnOnAssetsTtm
}

// GetReturnOnAssetsTtmOk returns a tuple with the ReturnOnAssetsTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.ReturnOnAssetsTtm) {
		return nil, false
	}
	return o.ReturnOnAssetsTtm, true
}

// HasReturnOnAssetsTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnAssetsTtm() bool {
	if o != nil && !IsNil(o.ReturnOnAssetsTtm) {
		return true
	}

	return false
}

// SetReturnOnAssetsTtm gets a reference to the given float64 and assigns it to the ReturnOnAssetsTtm field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnAssetsTtm(v float64) {
	o.ReturnOnAssetsTtm = &v
}

// GetReturnOnEquityTtm returns the ReturnOnEquityTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtm() float64 {
	if o == nil || IsNil(o.ReturnOnEquityTtm) {
		var ret float64
		return ret
	}
	return *o.ReturnOnEquityTtm
}

// GetReturnOnEquityTtmOk returns a tuple with the ReturnOnEquityTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.ReturnOnEquityTtm) {
		return nil, false
	}
	return o.ReturnOnEquityTtm, true
}

// HasReturnOnEquityTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnEquityTtm() bool {
	if o != nil && !IsNil(o.ReturnOnEquityTtm) {
		return true
	}

	return false
}

// SetReturnOnEquityTtm gets a reference to the given float64 and assigns it to the ReturnOnEquityTtm field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnEquityTtm(v float64) {
	o.ReturnOnEquityTtm = &v
}

// GetIncomeStatement returns the IncomeStatement field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatement() GetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	if o == nil || IsNil(o.IncomeStatement) {
		var ret GetStatistics200ResponseStatisticsFinancialsIncomeStatement
		return ret
	}
	return *o.IncomeStatement
}

// GetIncomeStatementOk returns a tuple with the IncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatementOk() (*GetStatistics200ResponseStatisticsFinancialsIncomeStatement, bool) {
	if o == nil || IsNil(o.IncomeStatement) {
		return nil, false
	}
	return o.IncomeStatement, true
}

// HasIncomeStatement returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasIncomeStatement() bool {
	if o != nil && !IsNil(o.IncomeStatement) {
		return true
	}

	return false
}

// SetIncomeStatement gets a reference to the given GetStatistics200ResponseStatisticsFinancialsIncomeStatement and assigns it to the IncomeStatement field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetIncomeStatement(v GetStatistics200ResponseStatisticsFinancialsIncomeStatement) {
	o.IncomeStatement = &v
}

// GetBalanceSheet returns the BalanceSheet field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheet() GetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	if o == nil || IsNil(o.BalanceSheet) {
		var ret GetStatistics200ResponseStatisticsFinancialsBalanceSheet
		return ret
	}
	return *o.BalanceSheet
}

// GetBalanceSheetOk returns a tuple with the BalanceSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheetOk() (*GetStatistics200ResponseStatisticsFinancialsBalanceSheet, bool) {
	if o == nil || IsNil(o.BalanceSheet) {
		return nil, false
	}
	return o.BalanceSheet, true
}

// HasBalanceSheet returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasBalanceSheet() bool {
	if o != nil && !IsNil(o.BalanceSheet) {
		return true
	}

	return false
}

// SetBalanceSheet gets a reference to the given GetStatistics200ResponseStatisticsFinancialsBalanceSheet and assigns it to the BalanceSheet field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetBalanceSheet(v GetStatistics200ResponseStatisticsFinancialsBalanceSheet) {
	o.BalanceSheet = &v
}

// GetCashFlow returns the CashFlow field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlow() GetStatistics200ResponseStatisticsFinancialsCashFlow {
	if o == nil || IsNil(o.CashFlow) {
		var ret GetStatistics200ResponseStatisticsFinancialsCashFlow
		return ret
	}
	return *o.CashFlow
}

// GetCashFlowOk returns a tuple with the CashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlowOk() (*GetStatistics200ResponseStatisticsFinancialsCashFlow, bool) {
	if o == nil || IsNil(o.CashFlow) {
		return nil, false
	}
	return o.CashFlow, true
}

// HasCashFlow returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasCashFlow() bool {
	if o != nil && !IsNil(o.CashFlow) {
		return true
	}

	return false
}

// SetCashFlow gets a reference to the given GetStatistics200ResponseStatisticsFinancialsCashFlow and assigns it to the CashFlow field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetCashFlow(v GetStatistics200ResponseStatisticsFinancialsCashFlow) {
	o.CashFlow = &v
}

func (o GetStatistics200ResponseStatisticsFinancials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsFinancials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiscalYearEnds) {
		toSerialize["fiscal_year_ends"] = o.FiscalYearEnds
	}
	if !IsNil(o.MostRecentQuarter) {
		toSerialize["most_recent_quarter"] = o.MostRecentQuarter
	}
	if !IsNil(o.GrossMargin) {
		toSerialize["gross_margin"] = o.GrossMargin
	}
	if !IsNil(o.ProfitMargin) {
		toSerialize["profit_margin"] = o.ProfitMargin
	}
	if !IsNil(o.OperatingMargin) {
		toSerialize["operating_margin"] = o.OperatingMargin
	}
	if !IsNil(o.ReturnOnAssetsTtm) {
		toSerialize["return_on_assets_ttm"] = o.ReturnOnAssetsTtm
	}
	if !IsNil(o.ReturnOnEquityTtm) {
		toSerialize["return_on_equity_ttm"] = o.ReturnOnEquityTtm
	}
	if !IsNil(o.IncomeStatement) {
		toSerialize["income_statement"] = o.IncomeStatement
	}
	if !IsNil(o.BalanceSheet) {
		toSerialize["balance_sheet"] = o.BalanceSheet
	}
	if !IsNil(o.CashFlow) {
		toSerialize["cash_flow"] = o.CashFlow
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsFinancials struct {
	value *GetStatistics200ResponseStatisticsFinancials
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) Get() *GetStatistics200ResponseStatisticsFinancials {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) Set(val *GetStatistics200ResponseStatisticsFinancials) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsFinancials(val *GetStatistics200ResponseStatisticsFinancials) *NullableGetStatistics200ResponseStatisticsFinancials {
	return &NullableGetStatistics200ResponseStatisticsFinancials{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


