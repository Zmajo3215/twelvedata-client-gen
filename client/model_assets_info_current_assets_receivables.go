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

// checks if the AssetsInfoCurrentAssetsReceivables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoCurrentAssetsReceivables{}

// AssetsInfoCurrentAssetsReceivables Receivables information
type AssetsInfoCurrentAssetsReceivables struct {
	// Total receivables
	TotalReceivables *float64 `json:"total_receivables,omitempty"`
	// Accounts receivable
	AccountsReceivable *float64 `json:"accounts_receivable,omitempty"`
	// Gross accounts receivable
	GrossAccountsReceivable *float64 `json:"gross_accounts_receivable,omitempty"`
	// Allowance for doubtful accounts receivable
	AllowanceForDoubtfulAccountsReceivable *float64 `json:"allowance_for_doubtful_accounts_receivable,omitempty"`
	// Receivables adjustments allowances
	ReceivablesAdjustmentsAllowances *float64 `json:"receivables_adjustments_allowances,omitempty"`
	// Other receivables
	OtherReceivables *float64 `json:"other_receivables,omitempty"`
	// Due from related parties current
	DueFromRelatedPartiesCurrent *float64 `json:"due_from_related_parties_current,omitempty"`
	// Taxes receivable
	TaxesReceivable *float64 `json:"taxes_receivable,omitempty"`
	// Accrued interest receivable
	AccruedInterestReceivable *float64 `json:"accrued_interest_receivable,omitempty"`
	// Notes receivable
	NotesReceivable *float64 `json:"notes_receivable,omitempty"`
	// Loans receivable
	LoansReceivable *float64 `json:"loans_receivable,omitempty"`
}

// NewAssetsInfoCurrentAssetsReceivables instantiates a new AssetsInfoCurrentAssetsReceivables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoCurrentAssetsReceivables() *AssetsInfoCurrentAssetsReceivables {
	this := AssetsInfoCurrentAssetsReceivables{}
	return &this
}

// NewAssetsInfoCurrentAssetsReceivablesWithDefaults instantiates a new AssetsInfoCurrentAssetsReceivables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoCurrentAssetsReceivablesWithDefaults() *AssetsInfoCurrentAssetsReceivables {
	this := AssetsInfoCurrentAssetsReceivables{}
	return &this
}

// GetTotalReceivables returns the TotalReceivables field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetTotalReceivables() float64 {
	if o == nil || IsNil(o.TotalReceivables) {
		var ret float64
		return ret
	}
	return *o.TotalReceivables
}

// GetTotalReceivablesOk returns a tuple with the TotalReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetTotalReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalReceivables) {
		return nil, false
	}
	return o.TotalReceivables, true
}

// HasTotalReceivables returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasTotalReceivables() bool {
	if o != nil && !IsNil(o.TotalReceivables) {
		return true
	}

	return false
}

// SetTotalReceivables gets a reference to the given float64 and assigns it to the TotalReceivables field.
func (o *AssetsInfoCurrentAssetsReceivables) SetTotalReceivables(v float64) {
	o.TotalReceivables = &v
}

// GetAccountsReceivable returns the AccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccountsReceivable() float64 {
	if o == nil || IsNil(o.AccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AccountsReceivable
}

// GetAccountsReceivableOk returns a tuple with the AccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsReceivable) {
		return nil, false
	}
	return o.AccountsReceivable, true
}

// HasAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAccountsReceivable() bool {
	if o != nil && !IsNil(o.AccountsReceivable) {
		return true
	}

	return false
}

// SetAccountsReceivable gets a reference to the given float64 and assigns it to the AccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAccountsReceivable(v float64) {
	o.AccountsReceivable = &v
}

// GetGrossAccountsReceivable returns the GrossAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetGrossAccountsReceivable() float64 {
	if o == nil || IsNil(o.GrossAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.GrossAccountsReceivable
}

// GetGrossAccountsReceivableOk returns a tuple with the GrossAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetGrossAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossAccountsReceivable) {
		return nil, false
	}
	return o.GrossAccountsReceivable, true
}

// HasGrossAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasGrossAccountsReceivable() bool {
	if o != nil && !IsNil(o.GrossAccountsReceivable) {
		return true
	}

	return false
}

// SetGrossAccountsReceivable gets a reference to the given float64 and assigns it to the GrossAccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetGrossAccountsReceivable(v float64) {
	o.GrossAccountsReceivable = &v
}

// GetAllowanceForDoubtfulAccountsReceivable returns the AllowanceForDoubtfulAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAllowanceForDoubtfulAccountsReceivable() float64 {
	if o == nil || IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AllowanceForDoubtfulAccountsReceivable
}

// GetAllowanceForDoubtfulAccountsReceivableOk returns a tuple with the AllowanceForDoubtfulAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAllowanceForDoubtfulAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		return nil, false
	}
	return o.AllowanceForDoubtfulAccountsReceivable, true
}

// HasAllowanceForDoubtfulAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAllowanceForDoubtfulAccountsReceivable() bool {
	if o != nil && !IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		return true
	}

	return false
}

// SetAllowanceForDoubtfulAccountsReceivable gets a reference to the given float64 and assigns it to the AllowanceForDoubtfulAccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAllowanceForDoubtfulAccountsReceivable(v float64) {
	o.AllowanceForDoubtfulAccountsReceivable = &v
}

// GetReceivablesAdjustmentsAllowances returns the ReceivablesAdjustmentsAllowances field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetReceivablesAdjustmentsAllowances() float64 {
	if o == nil || IsNil(o.ReceivablesAdjustmentsAllowances) {
		var ret float64
		return ret
	}
	return *o.ReceivablesAdjustmentsAllowances
}

// GetReceivablesAdjustmentsAllowancesOk returns a tuple with the ReceivablesAdjustmentsAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetReceivablesAdjustmentsAllowancesOk() (*float64, bool) {
	if o == nil || IsNil(o.ReceivablesAdjustmentsAllowances) {
		return nil, false
	}
	return o.ReceivablesAdjustmentsAllowances, true
}

// HasReceivablesAdjustmentsAllowances returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasReceivablesAdjustmentsAllowances() bool {
	if o != nil && !IsNil(o.ReceivablesAdjustmentsAllowances) {
		return true
	}

	return false
}

// SetReceivablesAdjustmentsAllowances gets a reference to the given float64 and assigns it to the ReceivablesAdjustmentsAllowances field.
func (o *AssetsInfoCurrentAssetsReceivables) SetReceivablesAdjustmentsAllowances(v float64) {
	o.ReceivablesAdjustmentsAllowances = &v
}

// GetOtherReceivables returns the OtherReceivables field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetOtherReceivables() float64 {
	if o == nil || IsNil(o.OtherReceivables) {
		var ret float64
		return ret
	}
	return *o.OtherReceivables
}

// GetOtherReceivablesOk returns a tuple with the OtherReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetOtherReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherReceivables) {
		return nil, false
	}
	return o.OtherReceivables, true
}

// HasOtherReceivables returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasOtherReceivables() bool {
	if o != nil && !IsNil(o.OtherReceivables) {
		return true
	}

	return false
}

// SetOtherReceivables gets a reference to the given float64 and assigns it to the OtherReceivables field.
func (o *AssetsInfoCurrentAssetsReceivables) SetOtherReceivables(v float64) {
	o.OtherReceivables = &v
}

// GetDueFromRelatedPartiesCurrent returns the DueFromRelatedPartiesCurrent field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetDueFromRelatedPartiesCurrent() float64 {
	if o == nil || IsNil(o.DueFromRelatedPartiesCurrent) {
		var ret float64
		return ret
	}
	return *o.DueFromRelatedPartiesCurrent
}

// GetDueFromRelatedPartiesCurrentOk returns a tuple with the DueFromRelatedPartiesCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetDueFromRelatedPartiesCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueFromRelatedPartiesCurrent) {
		return nil, false
	}
	return o.DueFromRelatedPartiesCurrent, true
}

// HasDueFromRelatedPartiesCurrent returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasDueFromRelatedPartiesCurrent() bool {
	if o != nil && !IsNil(o.DueFromRelatedPartiesCurrent) {
		return true
	}

	return false
}

// SetDueFromRelatedPartiesCurrent gets a reference to the given float64 and assigns it to the DueFromRelatedPartiesCurrent field.
func (o *AssetsInfoCurrentAssetsReceivables) SetDueFromRelatedPartiesCurrent(v float64) {
	o.DueFromRelatedPartiesCurrent = &v
}

// GetTaxesReceivable returns the TaxesReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetTaxesReceivable() float64 {
	if o == nil || IsNil(o.TaxesReceivable) {
		var ret float64
		return ret
	}
	return *o.TaxesReceivable
}

// GetTaxesReceivableOk returns a tuple with the TaxesReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetTaxesReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxesReceivable) {
		return nil, false
	}
	return o.TaxesReceivable, true
}

// HasTaxesReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasTaxesReceivable() bool {
	if o != nil && !IsNil(o.TaxesReceivable) {
		return true
	}

	return false
}

// SetTaxesReceivable gets a reference to the given float64 and assigns it to the TaxesReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetTaxesReceivable(v float64) {
	o.TaxesReceivable = &v
}

// GetAccruedInterestReceivable returns the AccruedInterestReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccruedInterestReceivable() float64 {
	if o == nil || IsNil(o.AccruedInterestReceivable) {
		var ret float64
		return ret
	}
	return *o.AccruedInterestReceivable
}

// GetAccruedInterestReceivableOk returns a tuple with the AccruedInterestReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccruedInterestReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccruedInterestReceivable) {
		return nil, false
	}
	return o.AccruedInterestReceivable, true
}

// HasAccruedInterestReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAccruedInterestReceivable() bool {
	if o != nil && !IsNil(o.AccruedInterestReceivable) {
		return true
	}

	return false
}

// SetAccruedInterestReceivable gets a reference to the given float64 and assigns it to the AccruedInterestReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAccruedInterestReceivable(v float64) {
	o.AccruedInterestReceivable = &v
}

// GetNotesReceivable returns the NotesReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetNotesReceivable() float64 {
	if o == nil || IsNil(o.NotesReceivable) {
		var ret float64
		return ret
	}
	return *o.NotesReceivable
}

// GetNotesReceivableOk returns a tuple with the NotesReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetNotesReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.NotesReceivable) {
		return nil, false
	}
	return o.NotesReceivable, true
}

// HasNotesReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasNotesReceivable() bool {
	if o != nil && !IsNil(o.NotesReceivable) {
		return true
	}

	return false
}

// SetNotesReceivable gets a reference to the given float64 and assigns it to the NotesReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetNotesReceivable(v float64) {
	o.NotesReceivable = &v
}

// GetLoansReceivable returns the LoansReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetLoansReceivable() float64 {
	if o == nil || IsNil(o.LoansReceivable) {
		var ret float64
		return ret
	}
	return *o.LoansReceivable
}

// GetLoansReceivableOk returns a tuple with the LoansReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetLoansReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.LoansReceivable) {
		return nil, false
	}
	return o.LoansReceivable, true
}

// HasLoansReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasLoansReceivable() bool {
	if o != nil && !IsNil(o.LoansReceivable) {
		return true
	}

	return false
}

// SetLoansReceivable gets a reference to the given float64 and assigns it to the LoansReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetLoansReceivable(v float64) {
	o.LoansReceivable = &v
}

func (o AssetsInfoCurrentAssetsReceivables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoCurrentAssetsReceivables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalReceivables) {
		toSerialize["total_receivables"] = o.TotalReceivables
	}
	if !IsNil(o.AccountsReceivable) {
		toSerialize["accounts_receivable"] = o.AccountsReceivable
	}
	if !IsNil(o.GrossAccountsReceivable) {
		toSerialize["gross_accounts_receivable"] = o.GrossAccountsReceivable
	}
	if !IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		toSerialize["allowance_for_doubtful_accounts_receivable"] = o.AllowanceForDoubtfulAccountsReceivable
	}
	if !IsNil(o.ReceivablesAdjustmentsAllowances) {
		toSerialize["receivables_adjustments_allowances"] = o.ReceivablesAdjustmentsAllowances
	}
	if !IsNil(o.OtherReceivables) {
		toSerialize["other_receivables"] = o.OtherReceivables
	}
	if !IsNil(o.DueFromRelatedPartiesCurrent) {
		toSerialize["due_from_related_parties_current"] = o.DueFromRelatedPartiesCurrent
	}
	if !IsNil(o.TaxesReceivable) {
		toSerialize["taxes_receivable"] = o.TaxesReceivable
	}
	if !IsNil(o.AccruedInterestReceivable) {
		toSerialize["accrued_interest_receivable"] = o.AccruedInterestReceivable
	}
	if !IsNil(o.NotesReceivable) {
		toSerialize["notes_receivable"] = o.NotesReceivable
	}
	if !IsNil(o.LoansReceivable) {
		toSerialize["loans_receivable"] = o.LoansReceivable
	}
	return toSerialize, nil
}

type NullableAssetsInfoCurrentAssetsReceivables struct {
	value *AssetsInfoCurrentAssetsReceivables
	isSet bool
}

func (v NullableAssetsInfoCurrentAssetsReceivables) Get() *AssetsInfoCurrentAssetsReceivables {
	return v.value
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) Set(val *AssetsInfoCurrentAssetsReceivables) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoCurrentAssetsReceivables) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoCurrentAssetsReceivables(val *AssetsInfoCurrentAssetsReceivables) *NullableAssetsInfoCurrentAssetsReceivables {
	return &NullableAssetsInfoCurrentAssetsReceivables{value: val, isSet: true}
}

func (v NullableAssetsInfoCurrentAssetsReceivables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


