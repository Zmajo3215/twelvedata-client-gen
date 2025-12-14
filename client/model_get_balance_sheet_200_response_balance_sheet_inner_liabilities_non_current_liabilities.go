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

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}

// GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities Non-current liabilities section
type GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities struct {
	// Represents money set aside for employee benefits such as gratuity
	LongTermProvisions *float64 `json:"long_term_provisions,omitempty"`
	// Represents amount of outstanding debt that has a maturity of 12 months or longer
	LongTermDebt *float64 `json:"long_term_debt,omitempty"`
	// Represents funds set aside as assets to pay for anticipated future losses
	ProvisionForRisksAndCharges *float64 `json:"provision_for_risks_and_charges,omitempty"`
	// Represents revenue producing activity for which revenue has not yet been recognized, and is not expected to be recognized in the next 12 months
	DeferredLiabilities *float64 `json:"deferred_liabilities,omitempty"`
	// Represents the value of derivative financial instruments that a company has issued
	DerivativeProductLiabilities *float64 `json:"derivative_product_liabilities,omitempty"`
	// Represents other non-current liabilities
	OtherNonCurrentLiabilities *float64 `json:"other_non_current_liabilities,omitempty"`
	// Represents total non-current liabilities
	TotalNonCurrentLiabilities *float64 `json:"total_non_current_liabilities,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilitiesWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilitiesWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}
	return &this
}

// GetLongTermProvisions returns the LongTermProvisions field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermProvisions() float64 {
	if o == nil || IsNil(o.LongTermProvisions) {
		var ret float64
		return ret
	}
	return *o.LongTermProvisions
}

// GetLongTermProvisionsOk returns a tuple with the LongTermProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermProvisions) {
		return nil, false
	}
	return o.LongTermProvisions, true
}

// HasLongTermProvisions returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasLongTermProvisions() bool {
	if o != nil && !IsNil(o.LongTermProvisions) {
		return true
	}

	return false
}

// SetLongTermProvisions gets a reference to the given float64 and assigns it to the LongTermProvisions field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetLongTermProvisions(v float64) {
	o.LongTermProvisions = &v
}

// GetLongTermDebt returns the LongTermDebt field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermDebt() float64 {
	if o == nil || IsNil(o.LongTermDebt) {
		var ret float64
		return ret
	}
	return *o.LongTermDebt
}

// GetLongTermDebtOk returns a tuple with the LongTermDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebt) {
		return nil, false
	}
	return o.LongTermDebt, true
}

// HasLongTermDebt returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasLongTermDebt() bool {
	if o != nil && !IsNil(o.LongTermDebt) {
		return true
	}

	return false
}

// SetLongTermDebt gets a reference to the given float64 and assigns it to the LongTermDebt field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetLongTermDebt(v float64) {
	o.LongTermDebt = &v
}

// GetProvisionForRisksAndCharges returns the ProvisionForRisksAndCharges field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetProvisionForRisksAndCharges() float64 {
	if o == nil || IsNil(o.ProvisionForRisksAndCharges) {
		var ret float64
		return ret
	}
	return *o.ProvisionForRisksAndCharges
}

// GetProvisionForRisksAndChargesOk returns a tuple with the ProvisionForRisksAndCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetProvisionForRisksAndChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionForRisksAndCharges) {
		return nil, false
	}
	return o.ProvisionForRisksAndCharges, true
}

// HasProvisionForRisksAndCharges returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasProvisionForRisksAndCharges() bool {
	if o != nil && !IsNil(o.ProvisionForRisksAndCharges) {
		return true
	}

	return false
}

// SetProvisionForRisksAndCharges gets a reference to the given float64 and assigns it to the ProvisionForRisksAndCharges field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetProvisionForRisksAndCharges(v float64) {
	o.ProvisionForRisksAndCharges = &v
}

// GetDeferredLiabilities returns the DeferredLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDeferredLiabilities() float64 {
	if o == nil || IsNil(o.DeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.DeferredLiabilities
}

// GetDeferredLiabilitiesOk returns a tuple with the DeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DeferredLiabilities) {
		return nil, false
	}
	return o.DeferredLiabilities, true
}

// HasDeferredLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasDeferredLiabilities() bool {
	if o != nil && !IsNil(o.DeferredLiabilities) {
		return true
	}

	return false
}

// SetDeferredLiabilities gets a reference to the given float64 and assigns it to the DeferredLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetDeferredLiabilities(v float64) {
	o.DeferredLiabilities = &v
}

// GetDerivativeProductLiabilities returns the DerivativeProductLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilities() float64 {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		var ret float64
		return ret
	}
	return *o.DerivativeProductLiabilities
}

// GetDerivativeProductLiabilitiesOk returns a tuple with the DerivativeProductLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		return nil, false
	}
	return o.DerivativeProductLiabilities, true
}

// HasDerivativeProductLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasDerivativeProductLiabilities() bool {
	if o != nil && !IsNil(o.DerivativeProductLiabilities) {
		return true
	}

	return false
}

// SetDerivativeProductLiabilities gets a reference to the given float64 and assigns it to the DerivativeProductLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetDerivativeProductLiabilities(v float64) {
	o.DerivativeProductLiabilities = &v
}

// GetOtherNonCurrentLiabilities returns the OtherNonCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentLiabilities
}

// GetOtherNonCurrentLiabilitiesOk returns a tuple with the OtherNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		return nil, false
	}
	return o.OtherNonCurrentLiabilities, true
}

// HasOtherNonCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasOtherNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherNonCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherNonCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetOtherNonCurrentLiabilities(v float64) {
	o.OtherNonCurrentLiabilities = &v
}

// GetTotalNonCurrentLiabilities returns the TotalNonCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.TotalNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentLiabilities
}

// GetTotalNonCurrentLiabilitiesOk returns a tuple with the TotalNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentLiabilities) {
		return nil, false
	}
	return o.TotalNonCurrentLiabilities, true
}

// HasTotalNonCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasTotalNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.TotalNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetTotalNonCurrentLiabilities gets a reference to the given float64 and assigns it to the TotalNonCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetTotalNonCurrentLiabilities(v float64) {
	o.TotalNonCurrentLiabilities = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongTermProvisions) {
		toSerialize["long_term_provisions"] = o.LongTermProvisions
	}
	if !IsNil(o.LongTermDebt) {
		toSerialize["long_term_debt"] = o.LongTermDebt
	}
	if !IsNil(o.ProvisionForRisksAndCharges) {
		toSerialize["provision_for_risks_and_charges"] = o.ProvisionForRisksAndCharges
	}
	if !IsNil(o.DeferredLiabilities) {
		toSerialize["deferred_liabilities"] = o.DeferredLiabilities
	}
	if !IsNil(o.DerivativeProductLiabilities) {
		toSerialize["derivative_product_liabilities"] = o.DerivativeProductLiabilities
	}
	if !IsNil(o.OtherNonCurrentLiabilities) {
		toSerialize["other_non_current_liabilities"] = o.OtherNonCurrentLiabilities
	}
	if !IsNil(o.TotalNonCurrentLiabilities) {
		toSerialize["total_non_current_liabilities"] = o.TotalNonCurrentLiabilities
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Get() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


