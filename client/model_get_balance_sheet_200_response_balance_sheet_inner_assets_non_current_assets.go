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

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}

// GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets Non-current assets section
type GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets struct {
	// Represents property owned
	Properties *float64 `json:"properties,omitempty"`
	// Represents land and improvements owned
	LandAndImprovements *float64 `json:"land_and_improvements,omitempty"`
	// Represents office equipment, furniture, and vehicles owned
	MachineryFurnitureEquipment *float64 `json:"machinery_furniture_equipment,omitempty"`
	// Represents the cost of construction work, which is not yet completed
	ConstructionInProgress *float64 `json:"construction_in_progress,omitempty"`
	// Represents operating and financial leases
	Leases *float64 `json:"leases,omitempty"`
	// Represents the cumulative depreciation of an asset that has been recorded
	AccumulatedDepreciation *float64 `json:"accumulated_depreciation,omitempty"`
	// Represents the value of a brand name, solid customer base, good customer relations, good employee relations, and proprietary technology
	Goodwill *float64 `json:"goodwill,omitempty"`
	// Represents real estate property purchased with the intention of earning a return on the investment
	InvestmentProperties *float64 `json:"investment_properties,omitempty"`
	// Represents liquid asset that gets its value from a contractual right or ownership claim
	FinancialAssets *float64 `json:"financial_assets,omitempty"`
	// Represents the patents, trademarks, and other intellectual properties
	IntangibleAssets *float64 `json:"intangible_assets,omitempty"`
	// Represents available for sale financial securities
	InvestmentsAndAdvances *float64 `json:"investments_and_advances,omitempty"`
	// Represents other long-term assets
	OtherNonCurrentAssets *float64 `json:"other_non_current_assets,omitempty"`
	// All long-term assets values in total
	TotalNonCurrentAssets *float64 `json:"total_non_current_assets,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetProperties() float64 {
	if o == nil || IsNil(o.Properties) {
		var ret float64
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given float64 and assigns it to the Properties field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetProperties(v float64) {
	o.Properties = &v
}

// GetLandAndImprovements returns the LandAndImprovements field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLandAndImprovements() float64 {
	if o == nil || IsNil(o.LandAndImprovements) {
		var ret float64
		return ret
	}
	return *o.LandAndImprovements
}

// GetLandAndImprovementsOk returns a tuple with the LandAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLandAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.LandAndImprovements) {
		return nil, false
	}
	return o.LandAndImprovements, true
}

// HasLandAndImprovements returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasLandAndImprovements() bool {
	if o != nil && !IsNil(o.LandAndImprovements) {
		return true
	}

	return false
}

// SetLandAndImprovements gets a reference to the given float64 and assigns it to the LandAndImprovements field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetLandAndImprovements(v float64) {
	o.LandAndImprovements = &v
}

// GetMachineryFurnitureEquipment returns the MachineryFurnitureEquipment field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetMachineryFurnitureEquipment() float64 {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		var ret float64
		return ret
	}
	return *o.MachineryFurnitureEquipment
}

// GetMachineryFurnitureEquipmentOk returns a tuple with the MachineryFurnitureEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetMachineryFurnitureEquipmentOk() (*float64, bool) {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		return nil, false
	}
	return o.MachineryFurnitureEquipment, true
}

// HasMachineryFurnitureEquipment returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasMachineryFurnitureEquipment() bool {
	if o != nil && !IsNil(o.MachineryFurnitureEquipment) {
		return true
	}

	return false
}

// SetMachineryFurnitureEquipment gets a reference to the given float64 and assigns it to the MachineryFurnitureEquipment field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetMachineryFurnitureEquipment(v float64) {
	o.MachineryFurnitureEquipment = &v
}

// GetConstructionInProgress returns the ConstructionInProgress field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetConstructionInProgress() float64 {
	if o == nil || IsNil(o.ConstructionInProgress) {
		var ret float64
		return ret
	}
	return *o.ConstructionInProgress
}

// GetConstructionInProgressOk returns a tuple with the ConstructionInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetConstructionInProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.ConstructionInProgress) {
		return nil, false
	}
	return o.ConstructionInProgress, true
}

// HasConstructionInProgress returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasConstructionInProgress() bool {
	if o != nil && !IsNil(o.ConstructionInProgress) {
		return true
	}

	return false
}

// SetConstructionInProgress gets a reference to the given float64 and assigns it to the ConstructionInProgress field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetConstructionInProgress(v float64) {
	o.ConstructionInProgress = &v
}

// GetLeases returns the Leases field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLeases() float64 {
	if o == nil || IsNil(o.Leases) {
		var ret float64
		return ret
	}
	return *o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetLeasesOk() (*float64, bool) {
	if o == nil || IsNil(o.Leases) {
		return nil, false
	}
	return o.Leases, true
}

// HasLeases returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasLeases() bool {
	if o != nil && !IsNil(o.Leases) {
		return true
	}

	return false
}

// SetLeases gets a reference to the given float64 and assigns it to the Leases field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetLeases(v float64) {
	o.Leases = &v
}

// GetAccumulatedDepreciation returns the AccumulatedDepreciation field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetAccumulatedDepreciation() float64 {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		var ret float64
		return ret
	}
	return *o.AccumulatedDepreciation
}

// GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetAccumulatedDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		return nil, false
	}
	return o.AccumulatedDepreciation, true
}

// HasAccumulatedDepreciation returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasAccumulatedDepreciation() bool {
	if o != nil && !IsNil(o.AccumulatedDepreciation) {
		return true
	}

	return false
}

// SetAccumulatedDepreciation gets a reference to the given float64 and assigns it to the AccumulatedDepreciation field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetAccumulatedDepreciation(v float64) {
	o.AccumulatedDepreciation = &v
}

// GetGoodwill returns the Goodwill field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetGoodwill() float64 {
	if o == nil || IsNil(o.Goodwill) {
		var ret float64
		return ret
	}
	return *o.Goodwill
}

// GetGoodwillOk returns a tuple with the Goodwill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetGoodwillOk() (*float64, bool) {
	if o == nil || IsNil(o.Goodwill) {
		return nil, false
	}
	return o.Goodwill, true
}

// HasGoodwill returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasGoodwill() bool {
	if o != nil && !IsNil(o.Goodwill) {
		return true
	}

	return false
}

// SetGoodwill gets a reference to the given float64 and assigns it to the Goodwill field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetGoodwill(v float64) {
	o.Goodwill = &v
}

// GetInvestmentProperties returns the InvestmentProperties field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentProperties() float64 {
	if o == nil || IsNil(o.InvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.InvestmentProperties
}

// GetInvestmentPropertiesOk returns a tuple with the InvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentProperties) {
		return nil, false
	}
	return o.InvestmentProperties, true
}

// HasInvestmentProperties returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasInvestmentProperties() bool {
	if o != nil && !IsNil(o.InvestmentProperties) {
		return true
	}

	return false
}

// SetInvestmentProperties gets a reference to the given float64 and assigns it to the InvestmentProperties field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetInvestmentProperties(v float64) {
	o.InvestmentProperties = &v
}

// GetFinancialAssets returns the FinancialAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetFinancialAssets() float64 {
	if o == nil || IsNil(o.FinancialAssets) {
		var ret float64
		return ret
	}
	return *o.FinancialAssets
}

// GetFinancialAssetsOk returns a tuple with the FinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssets) {
		return nil, false
	}
	return o.FinancialAssets, true
}

// HasFinancialAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasFinancialAssets() bool {
	if o != nil && !IsNil(o.FinancialAssets) {
		return true
	}

	return false
}

// SetFinancialAssets gets a reference to the given float64 and assigns it to the FinancialAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetFinancialAssets(v float64) {
	o.FinancialAssets = &v
}

// GetIntangibleAssets returns the IntangibleAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetIntangibleAssets() float64 {
	if o == nil || IsNil(o.IntangibleAssets) {
		var ret float64
		return ret
	}
	return *o.IntangibleAssets
}

// GetIntangibleAssetsOk returns a tuple with the IntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetIntangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.IntangibleAssets) {
		return nil, false
	}
	return o.IntangibleAssets, true
}

// HasIntangibleAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasIntangibleAssets() bool {
	if o != nil && !IsNil(o.IntangibleAssets) {
		return true
	}

	return false
}

// SetIntangibleAssets gets a reference to the given float64 and assigns it to the IntangibleAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetIntangibleAssets(v float64) {
	o.IntangibleAssets = &v
}

// GetInvestmentsAndAdvances returns the InvestmentsAndAdvances field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentsAndAdvances() float64 {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		var ret float64
		return ret
	}
	return *o.InvestmentsAndAdvances
}

// GetInvestmentsAndAdvancesOk returns a tuple with the InvestmentsAndAdvances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetInvestmentsAndAdvancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		return nil, false
	}
	return o.InvestmentsAndAdvances, true
}

// HasInvestmentsAndAdvances returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasInvestmentsAndAdvances() bool {
	if o != nil && !IsNil(o.InvestmentsAndAdvances) {
		return true
	}

	return false
}

// SetInvestmentsAndAdvances gets a reference to the given float64 and assigns it to the InvestmentsAndAdvances field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetInvestmentsAndAdvances(v float64) {
	o.InvestmentsAndAdvances = &v
}

// GetOtherNonCurrentAssets returns the OtherNonCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetOtherNonCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentAssets
}

// GetOtherNonCurrentAssetsOk returns a tuple with the OtherNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetOtherNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		return nil, false
	}
	return o.OtherNonCurrentAssets, true
}

// HasOtherNonCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasOtherNonCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherNonCurrentAssets) {
		return true
	}

	return false
}

// SetOtherNonCurrentAssets gets a reference to the given float64 and assigns it to the OtherNonCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetOtherNonCurrentAssets(v float64) {
	o.OtherNonCurrentAssets = &v
}

// GetTotalNonCurrentAssets returns the TotalNonCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetTotalNonCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentAssets
}

// GetTotalNonCurrentAssetsOk returns a tuple with the TotalNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) GetTotalNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		return nil, false
	}
	return o.TotalNonCurrentAssets, true
}

// HasTotalNonCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) HasTotalNonCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalNonCurrentAssets) {
		return true
	}

	return false
}

// SetTotalNonCurrentAssets gets a reference to the given float64 and assigns it to the TotalNonCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) SetTotalNonCurrentAssets(v float64) {
	o.TotalNonCurrentAssets = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.LandAndImprovements) {
		toSerialize["land_and_improvements"] = o.LandAndImprovements
	}
	if !IsNil(o.MachineryFurnitureEquipment) {
		toSerialize["machinery_furniture_equipment"] = o.MachineryFurnitureEquipment
	}
	if !IsNil(o.ConstructionInProgress) {
		toSerialize["construction_in_progress"] = o.ConstructionInProgress
	}
	if !IsNil(o.Leases) {
		toSerialize["leases"] = o.Leases
	}
	if !IsNil(o.AccumulatedDepreciation) {
		toSerialize["accumulated_depreciation"] = o.AccumulatedDepreciation
	}
	if !IsNil(o.Goodwill) {
		toSerialize["goodwill"] = o.Goodwill
	}
	if !IsNil(o.InvestmentProperties) {
		toSerialize["investment_properties"] = o.InvestmentProperties
	}
	if !IsNil(o.FinancialAssets) {
		toSerialize["financial_assets"] = o.FinancialAssets
	}
	if !IsNil(o.IntangibleAssets) {
		toSerialize["intangible_assets"] = o.IntangibleAssets
	}
	if !IsNil(o.InvestmentsAndAdvances) {
		toSerialize["investments_and_advances"] = o.InvestmentsAndAdvances
	}
	if !IsNil(o.OtherNonCurrentAssets) {
		toSerialize["other_non_current_assets"] = o.OtherNonCurrentAssets
	}
	if !IsNil(o.TotalNonCurrentAssets) {
		toSerialize["total_non_current_assets"] = o.TotalNonCurrentAssets
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Get() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


