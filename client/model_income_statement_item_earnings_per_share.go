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

// checks if the IncomeStatementItemEarningsPerShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemEarningsPerShare{}

// IncomeStatementItemEarningsPerShare Earnings per share information
type IncomeStatementItemEarningsPerShare struct {
	// Diluted EPS
	DilutedEps *float64 `json:"diluted_eps,omitempty"`
	// Basic EPS
	BasicEps *float64 `json:"basic_eps,omitempty"`
	// Continuing and discontinued diluted EPS
	ContinuingAndDiscontinuedDilutedEps *float64 `json:"continuing_and_discontinued_diluted_eps,omitempty"`
	// Continuing and discontinued basic EPS
	ContinuingAndDiscontinuedBasicEps *float64 `json:"continuing_and_discontinued_basic_eps,omitempty"`
	// Normalized diluted EPS
	NormalizedDilutedEps *float64 `json:"normalized_diluted_eps,omitempty"`
	// Normalized basic EPS
	NormalizedBasicEps *float64 `json:"normalized_basic_eps,omitempty"`
	// Reported normalized diluted EPS
	ReportedNormalizedDilutedEps *float64 `json:"reported_normalized_diluted_eps,omitempty"`
	// Reported normalized basic EPS
	ReportedNormalizedBasicEps *float64 `json:"reported_normalized_basic_eps,omitempty"`
	// Diluted EPS other gains losses
	DilutedEpsOtherGainsLosses *float64 `json:"diluted_eps_other_gains_losses,omitempty"`
	// Tax loss carryforward diluted EPS
	TaxLossCarryforwardDilutedEps *float64 `json:"tax_loss_carryforward_diluted_eps,omitempty"`
	// Diluted accounting change
	DilutedAccountingChange *float64 `json:"diluted_accounting_change,omitempty"`
	// Diluted extraordinary
	DilutedExtraordinary *float64 `json:"diluted_extraordinary,omitempty"`
	// Diluted discontinuous operations
	DilutedDiscontinuousOperations *float64 `json:"diluted_discontinuous_operations,omitempty"`
	// Diluted continuous operations
	DilutedContinuousOperations *float64 `json:"diluted_continuous_operations,omitempty"`
	// Basic EPS other gains losses
	BasicEpsOtherGainsLosses *float64 `json:"basic_eps_other_gains_losses,omitempty"`
	// Tax loss carryforward basic EPS
	TaxLossCarryforwardBasicEps *float64 `json:"tax_loss_carryforward_basic_eps,omitempty"`
	// Basic accounting change
	BasicAccountingChange *float64 `json:"basic_accounting_change,omitempty"`
	// Basic extraordinary
	BasicExtraordinary *float64 `json:"basic_extraordinary,omitempty"`
	// Basic discontinuous operations
	BasicDiscontinuousOperations *float64 `json:"basic_discontinuous_operations,omitempty"`
	// Basic continuous operations
	BasicContinuousOperations *float64 `json:"basic_continuous_operations,omitempty"`
	// Diluted NI available to common stockholders
	DilutedNiAvailToCommonStockholders *float64 `json:"diluted_ni_avail_to_common_stockholders,omitempty"`
	// Average dilution earnings
	AverageDilutionEarnings *float64 `json:"average_dilution_earnings,omitempty"`
}

// NewIncomeStatementItemEarningsPerShare instantiates a new IncomeStatementItemEarningsPerShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemEarningsPerShare() *IncomeStatementItemEarningsPerShare {
	this := IncomeStatementItemEarningsPerShare{}
	return &this
}

// NewIncomeStatementItemEarningsPerShareWithDefaults instantiates a new IncomeStatementItemEarningsPerShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemEarningsPerShareWithDefaults() *IncomeStatementItemEarningsPerShare {
	this := IncomeStatementItemEarningsPerShare{}
	return &this
}

// GetDilutedEps returns the DilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEps() float64 {
	if o == nil || IsNil(o.DilutedEps) {
		var ret float64
		return ret
	}
	return *o.DilutedEps
}

// GetDilutedEpsOk returns a tuple with the DilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedEps) {
		return nil, false
	}
	return o.DilutedEps, true
}

// HasDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedEps() bool {
	if o != nil && !IsNil(o.DilutedEps) {
		return true
	}

	return false
}

// SetDilutedEps gets a reference to the given float64 and assigns it to the DilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedEps(v float64) {
	o.DilutedEps = &v
}

// GetBasicEps returns the BasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEps() float64 {
	if o == nil || IsNil(o.BasicEps) {
		var ret float64
		return ret
	}
	return *o.BasicEps
}

// GetBasicEpsOk returns a tuple with the BasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicEps) {
		return nil, false
	}
	return o.BasicEps, true
}

// HasBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicEps() bool {
	if o != nil && !IsNil(o.BasicEps) {
		return true
	}

	return false
}

// SetBasicEps gets a reference to the given float64 and assigns it to the BasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicEps(v float64) {
	o.BasicEps = &v
}

// GetContinuingAndDiscontinuedDilutedEps returns the ContinuingAndDiscontinuedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedDilutedEps() float64 {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.ContinuingAndDiscontinuedDilutedEps
}

// GetContinuingAndDiscontinuedDilutedEpsOk returns a tuple with the ContinuingAndDiscontinuedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		return nil, false
	}
	return o.ContinuingAndDiscontinuedDilutedEps, true
}

// HasContinuingAndDiscontinuedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasContinuingAndDiscontinuedDilutedEps() bool {
	if o != nil && !IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		return true
	}

	return false
}

// SetContinuingAndDiscontinuedDilutedEps gets a reference to the given float64 and assigns it to the ContinuingAndDiscontinuedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetContinuingAndDiscontinuedDilutedEps(v float64) {
	o.ContinuingAndDiscontinuedDilutedEps = &v
}

// GetContinuingAndDiscontinuedBasicEps returns the ContinuingAndDiscontinuedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedBasicEps() float64 {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		var ret float64
		return ret
	}
	return *o.ContinuingAndDiscontinuedBasicEps
}

// GetContinuingAndDiscontinuedBasicEpsOk returns a tuple with the ContinuingAndDiscontinuedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetContinuingAndDiscontinuedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		return nil, false
	}
	return o.ContinuingAndDiscontinuedBasicEps, true
}

// HasContinuingAndDiscontinuedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasContinuingAndDiscontinuedBasicEps() bool {
	if o != nil && !IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		return true
	}

	return false
}

// SetContinuingAndDiscontinuedBasicEps gets a reference to the given float64 and assigns it to the ContinuingAndDiscontinuedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetContinuingAndDiscontinuedBasicEps(v float64) {
	o.ContinuingAndDiscontinuedBasicEps = &v
}

// GetNormalizedDilutedEps returns the NormalizedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedDilutedEps() float64 {
	if o == nil || IsNil(o.NormalizedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.NormalizedDilutedEps
}

// GetNormalizedDilutedEpsOk returns a tuple with the NormalizedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedDilutedEps) {
		return nil, false
	}
	return o.NormalizedDilutedEps, true
}

// HasNormalizedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasNormalizedDilutedEps() bool {
	if o != nil && !IsNil(o.NormalizedDilutedEps) {
		return true
	}

	return false
}

// SetNormalizedDilutedEps gets a reference to the given float64 and assigns it to the NormalizedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetNormalizedDilutedEps(v float64) {
	o.NormalizedDilutedEps = &v
}

// GetNormalizedBasicEps returns the NormalizedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedBasicEps() float64 {
	if o == nil || IsNil(o.NormalizedBasicEps) {
		var ret float64
		return ret
	}
	return *o.NormalizedBasicEps
}

// GetNormalizedBasicEpsOk returns a tuple with the NormalizedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetNormalizedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedBasicEps) {
		return nil, false
	}
	return o.NormalizedBasicEps, true
}

// HasNormalizedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasNormalizedBasicEps() bool {
	if o != nil && !IsNil(o.NormalizedBasicEps) {
		return true
	}

	return false
}

// SetNormalizedBasicEps gets a reference to the given float64 and assigns it to the NormalizedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetNormalizedBasicEps(v float64) {
	o.NormalizedBasicEps = &v
}

// GetReportedNormalizedDilutedEps returns the ReportedNormalizedDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedDilutedEps() float64 {
	if o == nil || IsNil(o.ReportedNormalizedDilutedEps) {
		var ret float64
		return ret
	}
	return *o.ReportedNormalizedDilutedEps
}

// GetReportedNormalizedDilutedEpsOk returns a tuple with the ReportedNormalizedDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReportedNormalizedDilutedEps) {
		return nil, false
	}
	return o.ReportedNormalizedDilutedEps, true
}

// HasReportedNormalizedDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasReportedNormalizedDilutedEps() bool {
	if o != nil && !IsNil(o.ReportedNormalizedDilutedEps) {
		return true
	}

	return false
}

// SetReportedNormalizedDilutedEps gets a reference to the given float64 and assigns it to the ReportedNormalizedDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetReportedNormalizedDilutedEps(v float64) {
	o.ReportedNormalizedDilutedEps = &v
}

// GetReportedNormalizedBasicEps returns the ReportedNormalizedBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedBasicEps() float64 {
	if o == nil || IsNil(o.ReportedNormalizedBasicEps) {
		var ret float64
		return ret
	}
	return *o.ReportedNormalizedBasicEps
}

// GetReportedNormalizedBasicEpsOk returns a tuple with the ReportedNormalizedBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetReportedNormalizedBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReportedNormalizedBasicEps) {
		return nil, false
	}
	return o.ReportedNormalizedBasicEps, true
}

// HasReportedNormalizedBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasReportedNormalizedBasicEps() bool {
	if o != nil && !IsNil(o.ReportedNormalizedBasicEps) {
		return true
	}

	return false
}

// SetReportedNormalizedBasicEps gets a reference to the given float64 and assigns it to the ReportedNormalizedBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetReportedNormalizedBasicEps(v float64) {
	o.ReportedNormalizedBasicEps = &v
}

// GetDilutedEpsOtherGainsLosses returns the DilutedEpsOtherGainsLosses field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOtherGainsLosses() float64 {
	if o == nil || IsNil(o.DilutedEpsOtherGainsLosses) {
		var ret float64
		return ret
	}
	return *o.DilutedEpsOtherGainsLosses
}

// GetDilutedEpsOtherGainsLossesOk returns a tuple with the DilutedEpsOtherGainsLosses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedEpsOtherGainsLossesOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedEpsOtherGainsLosses) {
		return nil, false
	}
	return o.DilutedEpsOtherGainsLosses, true
}

// HasDilutedEpsOtherGainsLosses returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedEpsOtherGainsLosses() bool {
	if o != nil && !IsNil(o.DilutedEpsOtherGainsLosses) {
		return true
	}

	return false
}

// SetDilutedEpsOtherGainsLosses gets a reference to the given float64 and assigns it to the DilutedEpsOtherGainsLosses field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedEpsOtherGainsLosses(v float64) {
	o.DilutedEpsOtherGainsLosses = &v
}

// GetTaxLossCarryforwardDilutedEps returns the TaxLossCarryforwardDilutedEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardDilutedEps() float64 {
	if o == nil || IsNil(o.TaxLossCarryforwardDilutedEps) {
		var ret float64
		return ret
	}
	return *o.TaxLossCarryforwardDilutedEps
}

// GetTaxLossCarryforwardDilutedEpsOk returns a tuple with the TaxLossCarryforwardDilutedEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardDilutedEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxLossCarryforwardDilutedEps) {
		return nil, false
	}
	return o.TaxLossCarryforwardDilutedEps, true
}

// HasTaxLossCarryforwardDilutedEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasTaxLossCarryforwardDilutedEps() bool {
	if o != nil && !IsNil(o.TaxLossCarryforwardDilutedEps) {
		return true
	}

	return false
}

// SetTaxLossCarryforwardDilutedEps gets a reference to the given float64 and assigns it to the TaxLossCarryforwardDilutedEps field.
func (o *IncomeStatementItemEarningsPerShare) SetTaxLossCarryforwardDilutedEps(v float64) {
	o.TaxLossCarryforwardDilutedEps = &v
}

// GetDilutedAccountingChange returns the DilutedAccountingChange field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedAccountingChange() float64 {
	if o == nil || IsNil(o.DilutedAccountingChange) {
		var ret float64
		return ret
	}
	return *o.DilutedAccountingChange
}

// GetDilutedAccountingChangeOk returns a tuple with the DilutedAccountingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedAccountingChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedAccountingChange) {
		return nil, false
	}
	return o.DilutedAccountingChange, true
}

// HasDilutedAccountingChange returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedAccountingChange() bool {
	if o != nil && !IsNil(o.DilutedAccountingChange) {
		return true
	}

	return false
}

// SetDilutedAccountingChange gets a reference to the given float64 and assigns it to the DilutedAccountingChange field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedAccountingChange(v float64) {
	o.DilutedAccountingChange = &v
}

// GetDilutedExtraordinary returns the DilutedExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedExtraordinary() float64 {
	if o == nil || IsNil(o.DilutedExtraordinary) {
		var ret float64
		return ret
	}
	return *o.DilutedExtraordinary
}

// GetDilutedExtraordinaryOk returns a tuple with the DilutedExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedExtraordinary) {
		return nil, false
	}
	return o.DilutedExtraordinary, true
}

// HasDilutedExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedExtraordinary() bool {
	if o != nil && !IsNil(o.DilutedExtraordinary) {
		return true
	}

	return false
}

// SetDilutedExtraordinary gets a reference to the given float64 and assigns it to the DilutedExtraordinary field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedExtraordinary(v float64) {
	o.DilutedExtraordinary = &v
}

// GetDilutedDiscontinuousOperations returns the DilutedDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.DilutedDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.DilutedDiscontinuousOperations
}

// GetDilutedDiscontinuousOperationsOk returns a tuple with the DilutedDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedDiscontinuousOperations) {
		return nil, false
	}
	return o.DilutedDiscontinuousOperations, true
}

// HasDilutedDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.DilutedDiscontinuousOperations) {
		return true
	}

	return false
}

// SetDilutedDiscontinuousOperations gets a reference to the given float64 and assigns it to the DilutedDiscontinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedDiscontinuousOperations(v float64) {
	o.DilutedDiscontinuousOperations = &v
}

// GetDilutedContinuousOperations returns the DilutedContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedContinuousOperations() float64 {
	if o == nil || IsNil(o.DilutedContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.DilutedContinuousOperations
}

// GetDilutedContinuousOperationsOk returns a tuple with the DilutedContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedContinuousOperations) {
		return nil, false
	}
	return o.DilutedContinuousOperations, true
}

// HasDilutedContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedContinuousOperations() bool {
	if o != nil && !IsNil(o.DilutedContinuousOperations) {
		return true
	}

	return false
}

// SetDilutedContinuousOperations gets a reference to the given float64 and assigns it to the DilutedContinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedContinuousOperations(v float64) {
	o.DilutedContinuousOperations = &v
}

// GetBasicEpsOtherGainsLosses returns the BasicEpsOtherGainsLosses field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOtherGainsLosses() float64 {
	if o == nil || IsNil(o.BasicEpsOtherGainsLosses) {
		var ret float64
		return ret
	}
	return *o.BasicEpsOtherGainsLosses
}

// GetBasicEpsOtherGainsLossesOk returns a tuple with the BasicEpsOtherGainsLosses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicEpsOtherGainsLossesOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicEpsOtherGainsLosses) {
		return nil, false
	}
	return o.BasicEpsOtherGainsLosses, true
}

// HasBasicEpsOtherGainsLosses returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicEpsOtherGainsLosses() bool {
	if o != nil && !IsNil(o.BasicEpsOtherGainsLosses) {
		return true
	}

	return false
}

// SetBasicEpsOtherGainsLosses gets a reference to the given float64 and assigns it to the BasicEpsOtherGainsLosses field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicEpsOtherGainsLosses(v float64) {
	o.BasicEpsOtherGainsLosses = &v
}

// GetTaxLossCarryforwardBasicEps returns the TaxLossCarryforwardBasicEps field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardBasicEps() float64 {
	if o == nil || IsNil(o.TaxLossCarryforwardBasicEps) {
		var ret float64
		return ret
	}
	return *o.TaxLossCarryforwardBasicEps
}

// GetTaxLossCarryforwardBasicEpsOk returns a tuple with the TaxLossCarryforwardBasicEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetTaxLossCarryforwardBasicEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxLossCarryforwardBasicEps) {
		return nil, false
	}
	return o.TaxLossCarryforwardBasicEps, true
}

// HasTaxLossCarryforwardBasicEps returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasTaxLossCarryforwardBasicEps() bool {
	if o != nil && !IsNil(o.TaxLossCarryforwardBasicEps) {
		return true
	}

	return false
}

// SetTaxLossCarryforwardBasicEps gets a reference to the given float64 and assigns it to the TaxLossCarryforwardBasicEps field.
func (o *IncomeStatementItemEarningsPerShare) SetTaxLossCarryforwardBasicEps(v float64) {
	o.TaxLossCarryforwardBasicEps = &v
}

// GetBasicAccountingChange returns the BasicAccountingChange field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicAccountingChange() float64 {
	if o == nil || IsNil(o.BasicAccountingChange) {
		var ret float64
		return ret
	}
	return *o.BasicAccountingChange
}

// GetBasicAccountingChangeOk returns a tuple with the BasicAccountingChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicAccountingChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicAccountingChange) {
		return nil, false
	}
	return o.BasicAccountingChange, true
}

// HasBasicAccountingChange returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicAccountingChange() bool {
	if o != nil && !IsNil(o.BasicAccountingChange) {
		return true
	}

	return false
}

// SetBasicAccountingChange gets a reference to the given float64 and assigns it to the BasicAccountingChange field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicAccountingChange(v float64) {
	o.BasicAccountingChange = &v
}

// GetBasicExtraordinary returns the BasicExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicExtraordinary() float64 {
	if o == nil || IsNil(o.BasicExtraordinary) {
		var ret float64
		return ret
	}
	return *o.BasicExtraordinary
}

// GetBasicExtraordinaryOk returns a tuple with the BasicExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicExtraordinary) {
		return nil, false
	}
	return o.BasicExtraordinary, true
}

// HasBasicExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicExtraordinary() bool {
	if o != nil && !IsNil(o.BasicExtraordinary) {
		return true
	}

	return false
}

// SetBasicExtraordinary gets a reference to the given float64 and assigns it to the BasicExtraordinary field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicExtraordinary(v float64) {
	o.BasicExtraordinary = &v
}

// GetBasicDiscontinuousOperations returns the BasicDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.BasicDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.BasicDiscontinuousOperations
}

// GetBasicDiscontinuousOperationsOk returns a tuple with the BasicDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicDiscontinuousOperations) {
		return nil, false
	}
	return o.BasicDiscontinuousOperations, true
}

// HasBasicDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.BasicDiscontinuousOperations) {
		return true
	}

	return false
}

// SetBasicDiscontinuousOperations gets a reference to the given float64 and assigns it to the BasicDiscontinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicDiscontinuousOperations(v float64) {
	o.BasicDiscontinuousOperations = &v
}

// GetBasicContinuousOperations returns the BasicContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetBasicContinuousOperations() float64 {
	if o == nil || IsNil(o.BasicContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.BasicContinuousOperations
}

// GetBasicContinuousOperationsOk returns a tuple with the BasicContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetBasicContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.BasicContinuousOperations) {
		return nil, false
	}
	return o.BasicContinuousOperations, true
}

// HasBasicContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasBasicContinuousOperations() bool {
	if o != nil && !IsNil(o.BasicContinuousOperations) {
		return true
	}

	return false
}

// SetBasicContinuousOperations gets a reference to the given float64 and assigns it to the BasicContinuousOperations field.
func (o *IncomeStatementItemEarningsPerShare) SetBasicContinuousOperations(v float64) {
	o.BasicContinuousOperations = &v
}

// GetDilutedNiAvailToCommonStockholders returns the DilutedNiAvailToCommonStockholders field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedNiAvailToCommonStockholders() float64 {
	if o == nil || IsNil(o.DilutedNiAvailToCommonStockholders) {
		var ret float64
		return ret
	}
	return *o.DilutedNiAvailToCommonStockholders
}

// GetDilutedNiAvailToCommonStockholdersOk returns a tuple with the DilutedNiAvailToCommonStockholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetDilutedNiAvailToCommonStockholdersOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedNiAvailToCommonStockholders) {
		return nil, false
	}
	return o.DilutedNiAvailToCommonStockholders, true
}

// HasDilutedNiAvailToCommonStockholders returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasDilutedNiAvailToCommonStockholders() bool {
	if o != nil && !IsNil(o.DilutedNiAvailToCommonStockholders) {
		return true
	}

	return false
}

// SetDilutedNiAvailToCommonStockholders gets a reference to the given float64 and assigns it to the DilutedNiAvailToCommonStockholders field.
func (o *IncomeStatementItemEarningsPerShare) SetDilutedNiAvailToCommonStockholders(v float64) {
	o.DilutedNiAvailToCommonStockholders = &v
}

// GetAverageDilutionEarnings returns the AverageDilutionEarnings field value if set, zero value otherwise.
func (o *IncomeStatementItemEarningsPerShare) GetAverageDilutionEarnings() float64 {
	if o == nil || IsNil(o.AverageDilutionEarnings) {
		var ret float64
		return ret
	}
	return *o.AverageDilutionEarnings
}

// GetAverageDilutionEarningsOk returns a tuple with the AverageDilutionEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEarningsPerShare) GetAverageDilutionEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageDilutionEarnings) {
		return nil, false
	}
	return o.AverageDilutionEarnings, true
}

// HasAverageDilutionEarnings returns a boolean if a field has been set.
func (o *IncomeStatementItemEarningsPerShare) HasAverageDilutionEarnings() bool {
	if o != nil && !IsNil(o.AverageDilutionEarnings) {
		return true
	}

	return false
}

// SetAverageDilutionEarnings gets a reference to the given float64 and assigns it to the AverageDilutionEarnings field.
func (o *IncomeStatementItemEarningsPerShare) SetAverageDilutionEarnings(v float64) {
	o.AverageDilutionEarnings = &v
}

func (o IncomeStatementItemEarningsPerShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemEarningsPerShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DilutedEps) {
		toSerialize["diluted_eps"] = o.DilutedEps
	}
	if !IsNil(o.BasicEps) {
		toSerialize["basic_eps"] = o.BasicEps
	}
	if !IsNil(o.ContinuingAndDiscontinuedDilutedEps) {
		toSerialize["continuing_and_discontinued_diluted_eps"] = o.ContinuingAndDiscontinuedDilutedEps
	}
	if !IsNil(o.ContinuingAndDiscontinuedBasicEps) {
		toSerialize["continuing_and_discontinued_basic_eps"] = o.ContinuingAndDiscontinuedBasicEps
	}
	if !IsNil(o.NormalizedDilutedEps) {
		toSerialize["normalized_diluted_eps"] = o.NormalizedDilutedEps
	}
	if !IsNil(o.NormalizedBasicEps) {
		toSerialize["normalized_basic_eps"] = o.NormalizedBasicEps
	}
	if !IsNil(o.ReportedNormalizedDilutedEps) {
		toSerialize["reported_normalized_diluted_eps"] = o.ReportedNormalizedDilutedEps
	}
	if !IsNil(o.ReportedNormalizedBasicEps) {
		toSerialize["reported_normalized_basic_eps"] = o.ReportedNormalizedBasicEps
	}
	if !IsNil(o.DilutedEpsOtherGainsLosses) {
		toSerialize["diluted_eps_other_gains_losses"] = o.DilutedEpsOtherGainsLosses
	}
	if !IsNil(o.TaxLossCarryforwardDilutedEps) {
		toSerialize["tax_loss_carryforward_diluted_eps"] = o.TaxLossCarryforwardDilutedEps
	}
	if !IsNil(o.DilutedAccountingChange) {
		toSerialize["diluted_accounting_change"] = o.DilutedAccountingChange
	}
	if !IsNil(o.DilutedExtraordinary) {
		toSerialize["diluted_extraordinary"] = o.DilutedExtraordinary
	}
	if !IsNil(o.DilutedDiscontinuousOperations) {
		toSerialize["diluted_discontinuous_operations"] = o.DilutedDiscontinuousOperations
	}
	if !IsNil(o.DilutedContinuousOperations) {
		toSerialize["diluted_continuous_operations"] = o.DilutedContinuousOperations
	}
	if !IsNil(o.BasicEpsOtherGainsLosses) {
		toSerialize["basic_eps_other_gains_losses"] = o.BasicEpsOtherGainsLosses
	}
	if !IsNil(o.TaxLossCarryforwardBasicEps) {
		toSerialize["tax_loss_carryforward_basic_eps"] = o.TaxLossCarryforwardBasicEps
	}
	if !IsNil(o.BasicAccountingChange) {
		toSerialize["basic_accounting_change"] = o.BasicAccountingChange
	}
	if !IsNil(o.BasicExtraordinary) {
		toSerialize["basic_extraordinary"] = o.BasicExtraordinary
	}
	if !IsNil(o.BasicDiscontinuousOperations) {
		toSerialize["basic_discontinuous_operations"] = o.BasicDiscontinuousOperations
	}
	if !IsNil(o.BasicContinuousOperations) {
		toSerialize["basic_continuous_operations"] = o.BasicContinuousOperations
	}
	if !IsNil(o.DilutedNiAvailToCommonStockholders) {
		toSerialize["diluted_ni_avail_to_common_stockholders"] = o.DilutedNiAvailToCommonStockholders
	}
	if !IsNil(o.AverageDilutionEarnings) {
		toSerialize["average_dilution_earnings"] = o.AverageDilutionEarnings
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemEarningsPerShare struct {
	value *IncomeStatementItemEarningsPerShare
	isSet bool
}

func (v NullableIncomeStatementItemEarningsPerShare) Get() *IncomeStatementItemEarningsPerShare {
	return v.value
}

func (v *NullableIncomeStatementItemEarningsPerShare) Set(val *IncomeStatementItemEarningsPerShare) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemEarningsPerShare) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemEarningsPerShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemEarningsPerShare(val *IncomeStatementItemEarningsPerShare) *NullableIncomeStatementItemEarningsPerShare {
	return &NullableIncomeStatementItemEarningsPerShare{value: val, isSet: true}
}

func (v NullableIncomeStatementItemEarningsPerShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemEarningsPerShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


