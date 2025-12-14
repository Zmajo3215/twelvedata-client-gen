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

// checks if the IncomeStatementItemNetIncome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemNetIncome{}

// IncomeStatementItemNetIncome Net income information
type IncomeStatementItemNetIncome struct {
	// Net income value
	NetIncomeValue *float64 `json:"net_income_value,omitempty"`
	// Net income common stockholders
	NetIncomeCommonStockholders *float64 `json:"net_income_common_stockholders,omitempty"`
	// Net income including noncontrolling interests
	NetIncomeIncludingNoncontrollingInterests *float64 `json:"net_income_including_noncontrolling_interests,omitempty"`
	// Net income from tax loss carryforward
	NetIncomeFromTaxLossCarryforward *float64 `json:"net_income_from_tax_loss_carryforward,omitempty"`
	// Net income extraordinary
	NetIncomeExtraordinary *float64 `json:"net_income_extraordinary,omitempty"`
	// Net income discontinuous operations
	NetIncomeDiscontinuousOperations *float64 `json:"net_income_discontinuous_operations,omitempty"`
	// Net income continuous operations
	NetIncomeContinuousOperations *float64 `json:"net_income_continuous_operations,omitempty"`
	// Net income from continuing operation net minority interest
	NetIncomeFromContinuingOperationNetMinorityInterest *float64 `json:"net_income_from_continuing_operation_net_minority_interest,omitempty"`
	// Net income from continuing and discontinued operation
	NetIncomeFromContinuingAndDiscontinuedOperation *float64 `json:"net_income_from_continuing_and_discontinued_operation,omitempty"`
	// Normalized income
	NormalizedIncome *float64 `json:"normalized_income,omitempty"`
	// Minority interests
	MinorityInterests *float64 `json:"minority_interests,omitempty"`
}

// NewIncomeStatementItemNetIncome instantiates a new IncomeStatementItemNetIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemNetIncome() *IncomeStatementItemNetIncome {
	this := IncomeStatementItemNetIncome{}
	return &this
}

// NewIncomeStatementItemNetIncomeWithDefaults instantiates a new IncomeStatementItemNetIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemNetIncomeWithDefaults() *IncomeStatementItemNetIncome {
	this := IncomeStatementItemNetIncome{}
	return &this
}

// GetNetIncomeValue returns the NetIncomeValue field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeValue() float64 {
	if o == nil || IsNil(o.NetIncomeValue) {
		var ret float64
		return ret
	}
	return *o.NetIncomeValue
}

// GetNetIncomeValueOk returns a tuple with the NetIncomeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeValueOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeValue) {
		return nil, false
	}
	return o.NetIncomeValue, true
}

// HasNetIncomeValue returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeValue() bool {
	if o != nil && !IsNil(o.NetIncomeValue) {
		return true
	}

	return false
}

// SetNetIncomeValue gets a reference to the given float64 and assigns it to the NetIncomeValue field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeValue(v float64) {
	o.NetIncomeValue = &v
}

// GetNetIncomeCommonStockholders returns the NetIncomeCommonStockholders field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeCommonStockholders() float64 {
	if o == nil || IsNil(o.NetIncomeCommonStockholders) {
		var ret float64
		return ret
	}
	return *o.NetIncomeCommonStockholders
}

// GetNetIncomeCommonStockholdersOk returns a tuple with the NetIncomeCommonStockholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeCommonStockholdersOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeCommonStockholders) {
		return nil, false
	}
	return o.NetIncomeCommonStockholders, true
}

// HasNetIncomeCommonStockholders returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeCommonStockholders() bool {
	if o != nil && !IsNil(o.NetIncomeCommonStockholders) {
		return true
	}

	return false
}

// SetNetIncomeCommonStockholders gets a reference to the given float64 and assigns it to the NetIncomeCommonStockholders field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeCommonStockholders(v float64) {
	o.NetIncomeCommonStockholders = &v
}

// GetNetIncomeIncludingNoncontrollingInterests returns the NetIncomeIncludingNoncontrollingInterests field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeIncludingNoncontrollingInterests() float64 {
	if o == nil || IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		var ret float64
		return ret
	}
	return *o.NetIncomeIncludingNoncontrollingInterests
}

// GetNetIncomeIncludingNoncontrollingInterestsOk returns a tuple with the NetIncomeIncludingNoncontrollingInterests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeIncludingNoncontrollingInterestsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		return nil, false
	}
	return o.NetIncomeIncludingNoncontrollingInterests, true
}

// HasNetIncomeIncludingNoncontrollingInterests returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeIncludingNoncontrollingInterests() bool {
	if o != nil && !IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		return true
	}

	return false
}

// SetNetIncomeIncludingNoncontrollingInterests gets a reference to the given float64 and assigns it to the NetIncomeIncludingNoncontrollingInterests field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeIncludingNoncontrollingInterests(v float64) {
	o.NetIncomeIncludingNoncontrollingInterests = &v
}

// GetNetIncomeFromTaxLossCarryforward returns the NetIncomeFromTaxLossCarryforward field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromTaxLossCarryforward() float64 {
	if o == nil || IsNil(o.NetIncomeFromTaxLossCarryforward) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromTaxLossCarryforward
}

// GetNetIncomeFromTaxLossCarryforwardOk returns a tuple with the NetIncomeFromTaxLossCarryforward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromTaxLossCarryforwardOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromTaxLossCarryforward) {
		return nil, false
	}
	return o.NetIncomeFromTaxLossCarryforward, true
}

// HasNetIncomeFromTaxLossCarryforward returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromTaxLossCarryforward() bool {
	if o != nil && !IsNil(o.NetIncomeFromTaxLossCarryforward) {
		return true
	}

	return false
}

// SetNetIncomeFromTaxLossCarryforward gets a reference to the given float64 and assigns it to the NetIncomeFromTaxLossCarryforward field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromTaxLossCarryforward(v float64) {
	o.NetIncomeFromTaxLossCarryforward = &v
}

// GetNetIncomeExtraordinary returns the NetIncomeExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeExtraordinary() float64 {
	if o == nil || IsNil(o.NetIncomeExtraordinary) {
		var ret float64
		return ret
	}
	return *o.NetIncomeExtraordinary
}

// GetNetIncomeExtraordinaryOk returns a tuple with the NetIncomeExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeExtraordinary) {
		return nil, false
	}
	return o.NetIncomeExtraordinary, true
}

// HasNetIncomeExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeExtraordinary() bool {
	if o != nil && !IsNil(o.NetIncomeExtraordinary) {
		return true
	}

	return false
}

// SetNetIncomeExtraordinary gets a reference to the given float64 and assigns it to the NetIncomeExtraordinary field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeExtraordinary(v float64) {
	o.NetIncomeExtraordinary = &v
}

// GetNetIncomeDiscontinuousOperations returns the NetIncomeDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.NetIncomeDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.NetIncomeDiscontinuousOperations
}

// GetNetIncomeDiscontinuousOperationsOk returns a tuple with the NetIncomeDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeDiscontinuousOperations) {
		return nil, false
	}
	return o.NetIncomeDiscontinuousOperations, true
}

// HasNetIncomeDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.NetIncomeDiscontinuousOperations) {
		return true
	}

	return false
}

// SetNetIncomeDiscontinuousOperations gets a reference to the given float64 and assigns it to the NetIncomeDiscontinuousOperations field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeDiscontinuousOperations(v float64) {
	o.NetIncomeDiscontinuousOperations = &v
}

// GetNetIncomeContinuousOperations returns the NetIncomeContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeContinuousOperations() float64 {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.NetIncomeContinuousOperations
}

// GetNetIncomeContinuousOperationsOk returns a tuple with the NetIncomeContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		return nil, false
	}
	return o.NetIncomeContinuousOperations, true
}

// HasNetIncomeContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeContinuousOperations() bool {
	if o != nil && !IsNil(o.NetIncomeContinuousOperations) {
		return true
	}

	return false
}

// SetNetIncomeContinuousOperations gets a reference to the given float64 and assigns it to the NetIncomeContinuousOperations field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeContinuousOperations(v float64) {
	o.NetIncomeContinuousOperations = &v
}

// GetNetIncomeFromContinuingOperationNetMinorityInterest returns the NetIncomeFromContinuingOperationNetMinorityInterest field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingOperationNetMinorityInterest() float64 {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromContinuingOperationNetMinorityInterest
}

// GetNetIncomeFromContinuingOperationNetMinorityInterestOk returns a tuple with the NetIncomeFromContinuingOperationNetMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingOperationNetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		return nil, false
	}
	return o.NetIncomeFromContinuingOperationNetMinorityInterest, true
}

// HasNetIncomeFromContinuingOperationNetMinorityInterest returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromContinuingOperationNetMinorityInterest() bool {
	if o != nil && !IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		return true
	}

	return false
}

// SetNetIncomeFromContinuingOperationNetMinorityInterest gets a reference to the given float64 and assigns it to the NetIncomeFromContinuingOperationNetMinorityInterest field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromContinuingOperationNetMinorityInterest(v float64) {
	o.NetIncomeFromContinuingOperationNetMinorityInterest = &v
}

// GetNetIncomeFromContinuingAndDiscontinuedOperation returns the NetIncomeFromContinuingAndDiscontinuedOperation field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingAndDiscontinuedOperation() float64 {
	if o == nil || IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromContinuingAndDiscontinuedOperation
}

// GetNetIncomeFromContinuingAndDiscontinuedOperationOk returns a tuple with the NetIncomeFromContinuingAndDiscontinuedOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingAndDiscontinuedOperationOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		return nil, false
	}
	return o.NetIncomeFromContinuingAndDiscontinuedOperation, true
}

// HasNetIncomeFromContinuingAndDiscontinuedOperation returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromContinuingAndDiscontinuedOperation() bool {
	if o != nil && !IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		return true
	}

	return false
}

// SetNetIncomeFromContinuingAndDiscontinuedOperation gets a reference to the given float64 and assigns it to the NetIncomeFromContinuingAndDiscontinuedOperation field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromContinuingAndDiscontinuedOperation(v float64) {
	o.NetIncomeFromContinuingAndDiscontinuedOperation = &v
}

// GetNormalizedIncome returns the NormalizedIncome field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNormalizedIncome() float64 {
	if o == nil || IsNil(o.NormalizedIncome) {
		var ret float64
		return ret
	}
	return *o.NormalizedIncome
}

// GetNormalizedIncomeOk returns a tuple with the NormalizedIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNormalizedIncomeOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedIncome) {
		return nil, false
	}
	return o.NormalizedIncome, true
}

// HasNormalizedIncome returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNormalizedIncome() bool {
	if o != nil && !IsNil(o.NormalizedIncome) {
		return true
	}

	return false
}

// SetNormalizedIncome gets a reference to the given float64 and assigns it to the NormalizedIncome field.
func (o *IncomeStatementItemNetIncome) SetNormalizedIncome(v float64) {
	o.NormalizedIncome = &v
}

// GetMinorityInterests returns the MinorityInterests field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetMinorityInterests() float64 {
	if o == nil || IsNil(o.MinorityInterests) {
		var ret float64
		return ret
	}
	return *o.MinorityInterests
}

// GetMinorityInterestsOk returns a tuple with the MinorityInterests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetMinorityInterestsOk() (*float64, bool) {
	if o == nil || IsNil(o.MinorityInterests) {
		return nil, false
	}
	return o.MinorityInterests, true
}

// HasMinorityInterests returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasMinorityInterests() bool {
	if o != nil && !IsNil(o.MinorityInterests) {
		return true
	}

	return false
}

// SetMinorityInterests gets a reference to the given float64 and assigns it to the MinorityInterests field.
func (o *IncomeStatementItemNetIncome) SetMinorityInterests(v float64) {
	o.MinorityInterests = &v
}

func (o IncomeStatementItemNetIncome) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemNetIncome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetIncomeValue) {
		toSerialize["net_income_value"] = o.NetIncomeValue
	}
	if !IsNil(o.NetIncomeCommonStockholders) {
		toSerialize["net_income_common_stockholders"] = o.NetIncomeCommonStockholders
	}
	if !IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		toSerialize["net_income_including_noncontrolling_interests"] = o.NetIncomeIncludingNoncontrollingInterests
	}
	if !IsNil(o.NetIncomeFromTaxLossCarryforward) {
		toSerialize["net_income_from_tax_loss_carryforward"] = o.NetIncomeFromTaxLossCarryforward
	}
	if !IsNil(o.NetIncomeExtraordinary) {
		toSerialize["net_income_extraordinary"] = o.NetIncomeExtraordinary
	}
	if !IsNil(o.NetIncomeDiscontinuousOperations) {
		toSerialize["net_income_discontinuous_operations"] = o.NetIncomeDiscontinuousOperations
	}
	if !IsNil(o.NetIncomeContinuousOperations) {
		toSerialize["net_income_continuous_operations"] = o.NetIncomeContinuousOperations
	}
	if !IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		toSerialize["net_income_from_continuing_operation_net_minority_interest"] = o.NetIncomeFromContinuingOperationNetMinorityInterest
	}
	if !IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		toSerialize["net_income_from_continuing_and_discontinued_operation"] = o.NetIncomeFromContinuingAndDiscontinuedOperation
	}
	if !IsNil(o.NormalizedIncome) {
		toSerialize["normalized_income"] = o.NormalizedIncome
	}
	if !IsNil(o.MinorityInterests) {
		toSerialize["minority_interests"] = o.MinorityInterests
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemNetIncome struct {
	value *IncomeStatementItemNetIncome
	isSet bool
}

func (v NullableIncomeStatementItemNetIncome) Get() *IncomeStatementItemNetIncome {
	return v.value
}

func (v *NullableIncomeStatementItemNetIncome) Set(val *IncomeStatementItemNetIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemNetIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemNetIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemNetIncome(val *IncomeStatementItemNetIncome) *NullableIncomeStatementItemNetIncome {
	return &NullableIncomeStatementItemNetIncome{value: val, isSet: true}
}

func (v NullableIncomeStatementItemNetIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemNetIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


