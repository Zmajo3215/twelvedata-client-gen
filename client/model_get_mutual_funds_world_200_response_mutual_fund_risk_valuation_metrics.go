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

// checks if the GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}

// GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics Valuation ratios and metrics of the fund and its category
type GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics struct {
	// Fund price to earnings metric
	PriceToEarnings *float64 `json:"price_to_earnings,omitempty"`
	// Average price to earnings metric of funds in the same category
	PriceToEarningsCategory *float64 `json:"price_to_earnings_category,omitempty"`
	// Fund price to book metric
	PriceToBook *float64 `json:"price_to_book,omitempty"`
	// Average price to book metric of funds in the same category
	PriceToBookCategory *float64 `json:"price_to_book_category,omitempty"`
	// Fund price to sales metric
	PriceToSales *float64 `json:"price_to_sales,omitempty"`
	// Average price to sales metric of funds in the same category
	PriceToSalesCategory *float64 `json:"price_to_sales_category,omitempty"`
	// Fund price to cashflow metric
	PriceToCashflow *float64 `json:"price_to_cashflow,omitempty"`
	// Average price to cashflow metric of funds in the same category
	PriceToCashflowCategory *float64 `json:"price_to_cashflow_category,omitempty"`
	// Median market capitalization of a fund
	MedianMarketCapitalization *int64 `json:"median_market_capitalization,omitempty"`
	// Median market capitalization of funds in the same category
	MedianMarketCapitalizationCategory *int64 `json:"median_market_capitalization_category,omitempty"`
	// Earnings growth over the last three years
	Var3YearEarningsGrowth *float64 `json:"3_year_earnings_growth,omitempty"`
	// Earnings growth over the last three years of funds in the same category
	Var3YearEarningsGrowthsCategory *float64 `json:"3_year_earnings_growths_category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	this := GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetricsWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetricsWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	this := GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}
	return &this
}

// GetPriceToEarnings returns the PriceToEarnings field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarnings() float64 {
	if o == nil || IsNil(o.PriceToEarnings) {
		var ret float64
		return ret
	}
	return *o.PriceToEarnings
}

// GetPriceToEarningsOk returns a tuple with the PriceToEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToEarnings) {
		return nil, false
	}
	return o.PriceToEarnings, true
}

// HasPriceToEarnings returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToEarnings() bool {
	if o != nil && !IsNil(o.PriceToEarnings) {
		return true
	}

	return false
}

// SetPriceToEarnings gets a reference to the given float64 and assigns it to the PriceToEarnings field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToEarnings(v float64) {
	o.PriceToEarnings = &v
}

// GetPriceToEarningsCategory returns the PriceToEarningsCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsCategory() float64 {
	if o == nil || IsNil(o.PriceToEarningsCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToEarningsCategory
}

// GetPriceToEarningsCategoryOk returns a tuple with the PriceToEarningsCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToEarningsCategory) {
		return nil, false
	}
	return o.PriceToEarningsCategory, true
}

// HasPriceToEarningsCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToEarningsCategory() bool {
	if o != nil && !IsNil(o.PriceToEarningsCategory) {
		return true
	}

	return false
}

// SetPriceToEarningsCategory gets a reference to the given float64 and assigns it to the PriceToEarningsCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToEarningsCategory(v float64) {
	o.PriceToEarningsCategory = &v
}

// GetPriceToBook returns the PriceToBook field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBook() float64 {
	if o == nil || IsNil(o.PriceToBook) {
		var ret float64
		return ret
	}
	return *o.PriceToBook
}

// GetPriceToBookOk returns a tuple with the PriceToBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBook) {
		return nil, false
	}
	return o.PriceToBook, true
}

// HasPriceToBook returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToBook() bool {
	if o != nil && !IsNil(o.PriceToBook) {
		return true
	}

	return false
}

// SetPriceToBook gets a reference to the given float64 and assigns it to the PriceToBook field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToBook(v float64) {
	o.PriceToBook = &v
}

// GetPriceToBookCategory returns the PriceToBookCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookCategory() float64 {
	if o == nil || IsNil(o.PriceToBookCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToBookCategory
}

// GetPriceToBookCategoryOk returns a tuple with the PriceToBookCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBookCategory) {
		return nil, false
	}
	return o.PriceToBookCategory, true
}

// HasPriceToBookCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToBookCategory() bool {
	if o != nil && !IsNil(o.PriceToBookCategory) {
		return true
	}

	return false
}

// SetPriceToBookCategory gets a reference to the given float64 and assigns it to the PriceToBookCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToBookCategory(v float64) {
	o.PriceToBookCategory = &v
}

// GetPriceToSales returns the PriceToSales field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSales() float64 {
	if o == nil || IsNil(o.PriceToSales) {
		var ret float64
		return ret
	}
	return *o.PriceToSales
}

// GetPriceToSalesOk returns a tuple with the PriceToSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSales) {
		return nil, false
	}
	return o.PriceToSales, true
}

// HasPriceToSales returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToSales() bool {
	if o != nil && !IsNil(o.PriceToSales) {
		return true
	}

	return false
}

// SetPriceToSales gets a reference to the given float64 and assigns it to the PriceToSales field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToSales(v float64) {
	o.PriceToSales = &v
}

// GetPriceToSalesCategory returns the PriceToSalesCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesCategory() float64 {
	if o == nil || IsNil(o.PriceToSalesCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToSalesCategory
}

// GetPriceToSalesCategoryOk returns a tuple with the PriceToSalesCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSalesCategory) {
		return nil, false
	}
	return o.PriceToSalesCategory, true
}

// HasPriceToSalesCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToSalesCategory() bool {
	if o != nil && !IsNil(o.PriceToSalesCategory) {
		return true
	}

	return false
}

// SetPriceToSalesCategory gets a reference to the given float64 and assigns it to the PriceToSalesCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToSalesCategory(v float64) {
	o.PriceToSalesCategory = &v
}

// GetPriceToCashflow returns the PriceToCashflow field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflow() float64 {
	if o == nil || IsNil(o.PriceToCashflow) {
		var ret float64
		return ret
	}
	return *o.PriceToCashflow
}

// GetPriceToCashflowOk returns a tuple with the PriceToCashflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToCashflow) {
		return nil, false
	}
	return o.PriceToCashflow, true
}

// HasPriceToCashflow returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToCashflow() bool {
	if o != nil && !IsNil(o.PriceToCashflow) {
		return true
	}

	return false
}

// SetPriceToCashflow gets a reference to the given float64 and assigns it to the PriceToCashflow field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToCashflow(v float64) {
	o.PriceToCashflow = &v
}

// GetPriceToCashflowCategory returns the PriceToCashflowCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowCategory() float64 {
	if o == nil || IsNil(o.PriceToCashflowCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToCashflowCategory
}

// GetPriceToCashflowCategoryOk returns a tuple with the PriceToCashflowCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToCashflowCategory) {
		return nil, false
	}
	return o.PriceToCashflowCategory, true
}

// HasPriceToCashflowCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToCashflowCategory() bool {
	if o != nil && !IsNil(o.PriceToCashflowCategory) {
		return true
	}

	return false
}

// SetPriceToCashflowCategory gets a reference to the given float64 and assigns it to the PriceToCashflowCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToCashflowCategory(v float64) {
	o.PriceToCashflowCategory = &v
}

// GetMedianMarketCapitalization returns the MedianMarketCapitalization field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalization() int64 {
	if o == nil || IsNil(o.MedianMarketCapitalization) {
		var ret int64
		return ret
	}
	return *o.MedianMarketCapitalization
}

// GetMedianMarketCapitalizationOk returns a tuple with the MedianMarketCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationOk() (*int64, bool) {
	if o == nil || IsNil(o.MedianMarketCapitalization) {
		return nil, false
	}
	return o.MedianMarketCapitalization, true
}

// HasMedianMarketCapitalization returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasMedianMarketCapitalization() bool {
	if o != nil && !IsNil(o.MedianMarketCapitalization) {
		return true
	}

	return false
}

// SetMedianMarketCapitalization gets a reference to the given int64 and assigns it to the MedianMarketCapitalization field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetMedianMarketCapitalization(v int64) {
	o.MedianMarketCapitalization = &v
}

// GetMedianMarketCapitalizationCategory returns the MedianMarketCapitalizationCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationCategory() int64 {
	if o == nil || IsNil(o.MedianMarketCapitalizationCategory) {
		var ret int64
		return ret
	}
	return *o.MedianMarketCapitalizationCategory
}

// GetMedianMarketCapitalizationCategoryOk returns a tuple with the MedianMarketCapitalizationCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationCategoryOk() (*int64, bool) {
	if o == nil || IsNil(o.MedianMarketCapitalizationCategory) {
		return nil, false
	}
	return o.MedianMarketCapitalizationCategory, true
}

// HasMedianMarketCapitalizationCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasMedianMarketCapitalizationCategory() bool {
	if o != nil && !IsNil(o.MedianMarketCapitalizationCategory) {
		return true
	}

	return false
}

// SetMedianMarketCapitalizationCategory gets a reference to the given int64 and assigns it to the MedianMarketCapitalizationCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetMedianMarketCapitalizationCategory(v int64) {
	o.MedianMarketCapitalizationCategory = &v
}

// GetVar3YearEarningsGrowth returns the Var3YearEarningsGrowth field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowth() float64 {
	if o == nil || IsNil(o.Var3YearEarningsGrowth) {
		var ret float64
		return ret
	}
	return *o.Var3YearEarningsGrowth
}

// GetVar3YearEarningsGrowthOk returns a tuple with the Var3YearEarningsGrowth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthOk() (*float64, bool) {
	if o == nil || IsNil(o.Var3YearEarningsGrowth) {
		return nil, false
	}
	return o.Var3YearEarningsGrowth, true
}

// HasVar3YearEarningsGrowth returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasVar3YearEarningsGrowth() bool {
	if o != nil && !IsNil(o.Var3YearEarningsGrowth) {
		return true
	}

	return false
}

// SetVar3YearEarningsGrowth gets a reference to the given float64 and assigns it to the Var3YearEarningsGrowth field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetVar3YearEarningsGrowth(v float64) {
	o.Var3YearEarningsGrowth = &v
}

// GetVar3YearEarningsGrowthsCategory returns the Var3YearEarningsGrowthsCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthsCategory() float64 {
	if o == nil || IsNil(o.Var3YearEarningsGrowthsCategory) {
		var ret float64
		return ret
	}
	return *o.Var3YearEarningsGrowthsCategory
}

// GetVar3YearEarningsGrowthsCategoryOk returns a tuple with the Var3YearEarningsGrowthsCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthsCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Var3YearEarningsGrowthsCategory) {
		return nil, false
	}
	return o.Var3YearEarningsGrowthsCategory, true
}

// HasVar3YearEarningsGrowthsCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasVar3YearEarningsGrowthsCategory() bool {
	if o != nil && !IsNil(o.Var3YearEarningsGrowthsCategory) {
		return true
	}

	return false
}

// SetVar3YearEarningsGrowthsCategory gets a reference to the given float64 and assigns it to the Var3YearEarningsGrowthsCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetVar3YearEarningsGrowthsCategory(v float64) {
	o.Var3YearEarningsGrowthsCategory = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceToEarnings) {
		toSerialize["price_to_earnings"] = o.PriceToEarnings
	}
	if !IsNil(o.PriceToEarningsCategory) {
		toSerialize["price_to_earnings_category"] = o.PriceToEarningsCategory
	}
	if !IsNil(o.PriceToBook) {
		toSerialize["price_to_book"] = o.PriceToBook
	}
	if !IsNil(o.PriceToBookCategory) {
		toSerialize["price_to_book_category"] = o.PriceToBookCategory
	}
	if !IsNil(o.PriceToSales) {
		toSerialize["price_to_sales"] = o.PriceToSales
	}
	if !IsNil(o.PriceToSalesCategory) {
		toSerialize["price_to_sales_category"] = o.PriceToSalesCategory
	}
	if !IsNil(o.PriceToCashflow) {
		toSerialize["price_to_cashflow"] = o.PriceToCashflow
	}
	if !IsNil(o.PriceToCashflowCategory) {
		toSerialize["price_to_cashflow_category"] = o.PriceToCashflowCategory
	}
	if !IsNil(o.MedianMarketCapitalization) {
		toSerialize["median_market_capitalization"] = o.MedianMarketCapitalization
	}
	if !IsNil(o.MedianMarketCapitalizationCategory) {
		toSerialize["median_market_capitalization_category"] = o.MedianMarketCapitalizationCategory
	}
	if !IsNil(o.Var3YearEarningsGrowth) {
		toSerialize["3_year_earnings_growth"] = o.Var3YearEarningsGrowth
	}
	if !IsNil(o.Var3YearEarningsGrowthsCategory) {
		toSerialize["3_year_earnings_growths_category"] = o.Var3YearEarningsGrowthsCategory
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics struct {
	value *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Get() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Set(val *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics(val *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


