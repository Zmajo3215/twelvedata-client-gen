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

// checks if the GetStatistics200ResponseStatisticsValuationsMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsValuationsMetrics{}

// GetStatistics200ResponseStatisticsValuationsMetrics Valuation metrics of the company
type GetStatistics200ResponseStatisticsValuationsMetrics struct {
	// Refers to the market value of the company's outstanding shares
	MarketCapitalization *float64 `json:"market_capitalization,omitempty"`
	// Refers to enterprise value (EV) of the company, often used as a more comprehensive alternative to market capitalization
	EnterpriseValue *float64 `json:"enterprise_value,omitempty"`
	// Refers to the trailing price-to-earnings (P/E). It is calculated by taking the current stock price and dividing it by the trailing earnings per share (EPS) for the past 12 months
	TrailingPe *float64 `json:"trailing_pe,omitempty"`
	// Refers to the forward price-to-earnings ratio. It is calculated by dividing the current share price by the estimated future earnings per share
	ForwardPe *float64 `json:"forward_pe,omitempty"`
	// The price/earnings to growth (PEG) ratio is a price-to-earnings ratio divided by the growth rate of the earnings
	PegRatio *float64 `json:"peg_ratio,omitempty"`
	// The price-to-sales (P/S) ratio is a valuation ratio that compares the market capitalization to its revenues over the last 12 months
	PriceToSalesTtm *float64 `json:"price_to_sales_ttm,omitempty"`
	// The price-to-book (P/B) ratio is equal to the current share price divided by the book value of all shares (BVPS) over the last quarter
	PriceToBookMrq *float64 `json:"price_to_book_mrq,omitempty"`
	// The enterprise value-to-revenue multiple (EV/R) is a measure that compares enterprise value to revenue
	EnterpriseToRevenue *float64 `json:"enterprise_to_revenue,omitempty"`
	// The enterprise value-to-ebitda multiple (EV/EBITDA) is a measure that compares enterprise value to EBITDA
	EnterpriseToEbitda *float64 `json:"enterprise_to_ebitda,omitempty"`
}

// NewGetStatistics200ResponseStatisticsValuationsMetrics instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsValuationsMetrics() *GetStatistics200ResponseStatisticsValuationsMetrics {
	this := GetStatistics200ResponseStatisticsValuationsMetrics{}
	return &this
}

// NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults() *GetStatistics200ResponseStatisticsValuationsMetrics {
	this := GetStatistics200ResponseStatisticsValuationsMetrics{}
	return &this
}

// GetMarketCapitalization returns the MarketCapitalization field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalization() float64 {
	if o == nil || IsNil(o.MarketCapitalization) {
		var ret float64
		return ret
	}
	return *o.MarketCapitalization
}

// GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalizationOk() (*float64, bool) {
	if o == nil || IsNil(o.MarketCapitalization) {
		return nil, false
	}
	return o.MarketCapitalization, true
}

// HasMarketCapitalization returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasMarketCapitalization() bool {
	if o != nil && !IsNil(o.MarketCapitalization) {
		return true
	}

	return false
}

// SetMarketCapitalization gets a reference to the given float64 and assigns it to the MarketCapitalization field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetMarketCapitalization(v float64) {
	o.MarketCapitalization = &v
}

// GetEnterpriseValue returns the EnterpriseValue field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValue() float64 {
	if o == nil || IsNil(o.EnterpriseValue) {
		var ret float64
		return ret
	}
	return *o.EnterpriseValue
}

// GetEnterpriseValueOk returns a tuple with the EnterpriseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValueOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseValue) {
		return nil, false
	}
	return o.EnterpriseValue, true
}

// HasEnterpriseValue returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseValue() bool {
	if o != nil && !IsNil(o.EnterpriseValue) {
		return true
	}

	return false
}

// SetEnterpriseValue gets a reference to the given float64 and assigns it to the EnterpriseValue field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseValue(v float64) {
	o.EnterpriseValue = &v
}

// GetTrailingPe returns the TrailingPe field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPe() float64 {
	if o == nil || IsNil(o.TrailingPe) {
		var ret float64
		return ret
	}
	return *o.TrailingPe
}

// GetTrailingPeOk returns a tuple with the TrailingPe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPeOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingPe) {
		return nil, false
	}
	return o.TrailingPe, true
}

// HasTrailingPe returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasTrailingPe() bool {
	if o != nil && !IsNil(o.TrailingPe) {
		return true
	}

	return false
}

// SetTrailingPe gets a reference to the given float64 and assigns it to the TrailingPe field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetTrailingPe(v float64) {
	o.TrailingPe = &v
}

// GetForwardPe returns the ForwardPe field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPe() float64 {
	if o == nil || IsNil(o.ForwardPe) {
		var ret float64
		return ret
	}
	return *o.ForwardPe
}

// GetForwardPeOk returns a tuple with the ForwardPe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPeOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardPe) {
		return nil, false
	}
	return o.ForwardPe, true
}

// HasForwardPe returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasForwardPe() bool {
	if o != nil && !IsNil(o.ForwardPe) {
		return true
	}

	return false
}

// SetForwardPe gets a reference to the given float64 and assigns it to the ForwardPe field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetForwardPe(v float64) {
	o.ForwardPe = &v
}

// GetPegRatio returns the PegRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatio() float64 {
	if o == nil || IsNil(o.PegRatio) {
		var ret float64
		return ret
	}
	return *o.PegRatio
}

// GetPegRatioOk returns a tuple with the PegRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.PegRatio) {
		return nil, false
	}
	return o.PegRatio, true
}

// HasPegRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPegRatio() bool {
	if o != nil && !IsNil(o.PegRatio) {
		return true
	}

	return false
}

// SetPegRatio gets a reference to the given float64 and assigns it to the PegRatio field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPegRatio(v float64) {
	o.PegRatio = &v
}

// GetPriceToSalesTtm returns the PriceToSalesTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtm() float64 {
	if o == nil || IsNil(o.PriceToSalesTtm) {
		var ret float64
		return ret
	}
	return *o.PriceToSalesTtm
}

// GetPriceToSalesTtmOk returns a tuple with the PriceToSalesTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSalesTtm) {
		return nil, false
	}
	return o.PriceToSalesTtm, true
}

// HasPriceToSalesTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToSalesTtm() bool {
	if o != nil && !IsNil(o.PriceToSalesTtm) {
		return true
	}

	return false
}

// SetPriceToSalesTtm gets a reference to the given float64 and assigns it to the PriceToSalesTtm field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToSalesTtm(v float64) {
	o.PriceToSalesTtm = &v
}

// GetPriceToBookMrq returns the PriceToBookMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrq() float64 {
	if o == nil || IsNil(o.PriceToBookMrq) {
		var ret float64
		return ret
	}
	return *o.PriceToBookMrq
}

// GetPriceToBookMrqOk returns a tuple with the PriceToBookMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBookMrq) {
		return nil, false
	}
	return o.PriceToBookMrq, true
}

// HasPriceToBookMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToBookMrq() bool {
	if o != nil && !IsNil(o.PriceToBookMrq) {
		return true
	}

	return false
}

// SetPriceToBookMrq gets a reference to the given float64 and assigns it to the PriceToBookMrq field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToBookMrq(v float64) {
	o.PriceToBookMrq = &v
}

// GetEnterpriseToRevenue returns the EnterpriseToRevenue field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenue() float64 {
	if o == nil || IsNil(o.EnterpriseToRevenue) {
		var ret float64
		return ret
	}
	return *o.EnterpriseToRevenue
}

// GetEnterpriseToRevenueOk returns a tuple with the EnterpriseToRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseToRevenue) {
		return nil, false
	}
	return o.EnterpriseToRevenue, true
}

// HasEnterpriseToRevenue returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToRevenue() bool {
	if o != nil && !IsNil(o.EnterpriseToRevenue) {
		return true
	}

	return false
}

// SetEnterpriseToRevenue gets a reference to the given float64 and assigns it to the EnterpriseToRevenue field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToRevenue(v float64) {
	o.EnterpriseToRevenue = &v
}

// GetEnterpriseToEbitda returns the EnterpriseToEbitda field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitda() float64 {
	if o == nil || IsNil(o.EnterpriseToEbitda) {
		var ret float64
		return ret
	}
	return *o.EnterpriseToEbitda
}

// GetEnterpriseToEbitdaOk returns a tuple with the EnterpriseToEbitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitdaOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseToEbitda) {
		return nil, false
	}
	return o.EnterpriseToEbitda, true
}

// HasEnterpriseToEbitda returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToEbitda() bool {
	if o != nil && !IsNil(o.EnterpriseToEbitda) {
		return true
	}

	return false
}

// SetEnterpriseToEbitda gets a reference to the given float64 and assigns it to the EnterpriseToEbitda field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToEbitda(v float64) {
	o.EnterpriseToEbitda = &v
}

func (o GetStatistics200ResponseStatisticsValuationsMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsValuationsMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketCapitalization) {
		toSerialize["market_capitalization"] = o.MarketCapitalization
	}
	if !IsNil(o.EnterpriseValue) {
		toSerialize["enterprise_value"] = o.EnterpriseValue
	}
	if !IsNil(o.TrailingPe) {
		toSerialize["trailing_pe"] = o.TrailingPe
	}
	if !IsNil(o.ForwardPe) {
		toSerialize["forward_pe"] = o.ForwardPe
	}
	if !IsNil(o.PegRatio) {
		toSerialize["peg_ratio"] = o.PegRatio
	}
	if !IsNil(o.PriceToSalesTtm) {
		toSerialize["price_to_sales_ttm"] = o.PriceToSalesTtm
	}
	if !IsNil(o.PriceToBookMrq) {
		toSerialize["price_to_book_mrq"] = o.PriceToBookMrq
	}
	if !IsNil(o.EnterpriseToRevenue) {
		toSerialize["enterprise_to_revenue"] = o.EnterpriseToRevenue
	}
	if !IsNil(o.EnterpriseToEbitda) {
		toSerialize["enterprise_to_ebitda"] = o.EnterpriseToEbitda
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsValuationsMetrics struct {
	value *GetStatistics200ResponseStatisticsValuationsMetrics
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) Get() *GetStatistics200ResponseStatisticsValuationsMetrics {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) Set(val *GetStatistics200ResponseStatisticsValuationsMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsValuationsMetrics(val *GetStatistics200ResponseStatisticsValuationsMetrics) *NullableGetStatistics200ResponseStatisticsValuationsMetrics {
	return &NullableGetStatistics200ResponseStatisticsValuationsMetrics{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


