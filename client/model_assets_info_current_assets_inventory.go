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

// checks if the AssetsInfoCurrentAssetsInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoCurrentAssetsInventory{}

// AssetsInfoCurrentAssetsInventory Inventory information
type AssetsInfoCurrentAssetsInventory struct {
	// Total inventory
	TotalInventory *float64 `json:"total_inventory,omitempty"`
	// Inventories adjustments allowances
	InventoriesAdjustmentsAllowances *float64 `json:"inventories_adjustments_allowances,omitempty"`
	// Other inventories
	OtherInventories *float64 `json:"other_inventories,omitempty"`
	// Finished goods
	FinishedGoods *float64 `json:"finished_goods,omitempty"`
	// Work in process
	WorkInProcess *float64 `json:"work_in_process,omitempty"`
	// Raw materials
	RawMaterials *float64 `json:"raw_materials,omitempty"`
}

// NewAssetsInfoCurrentAssetsInventory instantiates a new AssetsInfoCurrentAssetsInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoCurrentAssetsInventory() *AssetsInfoCurrentAssetsInventory {
	this := AssetsInfoCurrentAssetsInventory{}
	return &this
}

// NewAssetsInfoCurrentAssetsInventoryWithDefaults instantiates a new AssetsInfoCurrentAssetsInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoCurrentAssetsInventoryWithDefaults() *AssetsInfoCurrentAssetsInventory {
	this := AssetsInfoCurrentAssetsInventory{}
	return &this
}

// GetTotalInventory returns the TotalInventory field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetTotalInventory() float64 {
	if o == nil || IsNil(o.TotalInventory) {
		var ret float64
		return ret
	}
	return *o.TotalInventory
}

// GetTotalInventoryOk returns a tuple with the TotalInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetTotalInventoryOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalInventory) {
		return nil, false
	}
	return o.TotalInventory, true
}

// HasTotalInventory returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasTotalInventory() bool {
	if o != nil && !IsNil(o.TotalInventory) {
		return true
	}

	return false
}

// SetTotalInventory gets a reference to the given float64 and assigns it to the TotalInventory field.
func (o *AssetsInfoCurrentAssetsInventory) SetTotalInventory(v float64) {
	o.TotalInventory = &v
}

// GetInventoriesAdjustmentsAllowances returns the InventoriesAdjustmentsAllowances field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetInventoriesAdjustmentsAllowances() float64 {
	if o == nil || IsNil(o.InventoriesAdjustmentsAllowances) {
		var ret float64
		return ret
	}
	return *o.InventoriesAdjustmentsAllowances
}

// GetInventoriesAdjustmentsAllowancesOk returns a tuple with the InventoriesAdjustmentsAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetInventoriesAdjustmentsAllowancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InventoriesAdjustmentsAllowances) {
		return nil, false
	}
	return o.InventoriesAdjustmentsAllowances, true
}

// HasInventoriesAdjustmentsAllowances returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasInventoriesAdjustmentsAllowances() bool {
	if o != nil && !IsNil(o.InventoriesAdjustmentsAllowances) {
		return true
	}

	return false
}

// SetInventoriesAdjustmentsAllowances gets a reference to the given float64 and assigns it to the InventoriesAdjustmentsAllowances field.
func (o *AssetsInfoCurrentAssetsInventory) SetInventoriesAdjustmentsAllowances(v float64) {
	o.InventoriesAdjustmentsAllowances = &v
}

// GetOtherInventories returns the OtherInventories field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetOtherInventories() float64 {
	if o == nil || IsNil(o.OtherInventories) {
		var ret float64
		return ret
	}
	return *o.OtherInventories
}

// GetOtherInventoriesOk returns a tuple with the OtherInventories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetOtherInventoriesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherInventories) {
		return nil, false
	}
	return o.OtherInventories, true
}

// HasOtherInventories returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasOtherInventories() bool {
	if o != nil && !IsNil(o.OtherInventories) {
		return true
	}

	return false
}

// SetOtherInventories gets a reference to the given float64 and assigns it to the OtherInventories field.
func (o *AssetsInfoCurrentAssetsInventory) SetOtherInventories(v float64) {
	o.OtherInventories = &v
}

// GetFinishedGoods returns the FinishedGoods field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetFinishedGoods() float64 {
	if o == nil || IsNil(o.FinishedGoods) {
		var ret float64
		return ret
	}
	return *o.FinishedGoods
}

// GetFinishedGoodsOk returns a tuple with the FinishedGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetFinishedGoodsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinishedGoods) {
		return nil, false
	}
	return o.FinishedGoods, true
}

// HasFinishedGoods returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasFinishedGoods() bool {
	if o != nil && !IsNil(o.FinishedGoods) {
		return true
	}

	return false
}

// SetFinishedGoods gets a reference to the given float64 and assigns it to the FinishedGoods field.
func (o *AssetsInfoCurrentAssetsInventory) SetFinishedGoods(v float64) {
	o.FinishedGoods = &v
}

// GetWorkInProcess returns the WorkInProcess field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetWorkInProcess() float64 {
	if o == nil || IsNil(o.WorkInProcess) {
		var ret float64
		return ret
	}
	return *o.WorkInProcess
}

// GetWorkInProcessOk returns a tuple with the WorkInProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetWorkInProcessOk() (*float64, bool) {
	if o == nil || IsNil(o.WorkInProcess) {
		return nil, false
	}
	return o.WorkInProcess, true
}

// HasWorkInProcess returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasWorkInProcess() bool {
	if o != nil && !IsNil(o.WorkInProcess) {
		return true
	}

	return false
}

// SetWorkInProcess gets a reference to the given float64 and assigns it to the WorkInProcess field.
func (o *AssetsInfoCurrentAssetsInventory) SetWorkInProcess(v float64) {
	o.WorkInProcess = &v
}

// GetRawMaterials returns the RawMaterials field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsInventory) GetRawMaterials() float64 {
	if o == nil || IsNil(o.RawMaterials) {
		var ret float64
		return ret
	}
	return *o.RawMaterials
}

// GetRawMaterialsOk returns a tuple with the RawMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsInventory) GetRawMaterialsOk() (*float64, bool) {
	if o == nil || IsNil(o.RawMaterials) {
		return nil, false
	}
	return o.RawMaterials, true
}

// HasRawMaterials returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsInventory) HasRawMaterials() bool {
	if o != nil && !IsNil(o.RawMaterials) {
		return true
	}

	return false
}

// SetRawMaterials gets a reference to the given float64 and assigns it to the RawMaterials field.
func (o *AssetsInfoCurrentAssetsInventory) SetRawMaterials(v float64) {
	o.RawMaterials = &v
}

func (o AssetsInfoCurrentAssetsInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoCurrentAssetsInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalInventory) {
		toSerialize["total_inventory"] = o.TotalInventory
	}
	if !IsNil(o.InventoriesAdjustmentsAllowances) {
		toSerialize["inventories_adjustments_allowances"] = o.InventoriesAdjustmentsAllowances
	}
	if !IsNil(o.OtherInventories) {
		toSerialize["other_inventories"] = o.OtherInventories
	}
	if !IsNil(o.FinishedGoods) {
		toSerialize["finished_goods"] = o.FinishedGoods
	}
	if !IsNil(o.WorkInProcess) {
		toSerialize["work_in_process"] = o.WorkInProcess
	}
	if !IsNil(o.RawMaterials) {
		toSerialize["raw_materials"] = o.RawMaterials
	}
	return toSerialize, nil
}

type NullableAssetsInfoCurrentAssetsInventory struct {
	value *AssetsInfoCurrentAssetsInventory
	isSet bool
}

func (v NullableAssetsInfoCurrentAssetsInventory) Get() *AssetsInfoCurrentAssetsInventory {
	return v.value
}

func (v *NullableAssetsInfoCurrentAssetsInventory) Set(val *AssetsInfoCurrentAssetsInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoCurrentAssetsInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoCurrentAssetsInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoCurrentAssetsInventory(val *AssetsInfoCurrentAssetsInventory) *NullableAssetsInfoCurrentAssetsInventory {
	return &NullableAssetsInfoCurrentAssetsInventory{value: val, isSet: true}
}

func (v NullableAssetsInfoCurrentAssetsInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoCurrentAssetsInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


