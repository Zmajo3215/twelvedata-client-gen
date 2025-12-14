# GetQuote200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Symbol passed | [optional] 
**Name** | Pointer to **string** | Name of the instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard. Available for stocks, ETFs, mutual funds, bonds | [optional] 
**Currency** | Pointer to **string** | Currency in which the equity is denominated. Available for stocks, ETFs, mutual funds, bonds | [optional] 
**Datetime** | Pointer to **string** | Datetime in defined timezone referring to when the bar with specified interval was opened | [optional] 
**Timestamp** | Pointer to **int64** | Unix timestamp representing the opening candle of the specified interval | [optional] 
**LastQuoteAt** | Pointer to **int64** | Unix timestamp of last minute candle | [optional] 
**Open** | Pointer to **string** | Price at the opening of current bar | [optional] 
**High** | Pointer to **string** | Highest price which occurred during the current bar | [optional] 
**Low** | Pointer to **string** | Lowest price which occurred during the current bar | [optional] 
**Close** | Pointer to **string** | Close price at the end of the bar | [optional] 
**Volume** | Pointer to **string** | Trading volume during the bar. Available not for all instrument types | [optional] 
**PreviousClose** | Pointer to **string** | Close price at the end of the previous bar | [optional] 
**Change** | Pointer to **string** | Close - previous_close | [optional] 
**PercentChange** | Pointer to **string** | (Close - previous_close) / previous_close * 100 | [optional] 
**AverageVolume** | Pointer to **string** | Average volume of the specified period. Available not for all instrument types | [optional] 
**Rolling1dChange** | Pointer to **string** | Percent change in price between the current and the backward one, where period is 1 day. Available for crypto | [optional] 
**Rolling7dChange** | Pointer to **string** | Percent change in price between the current and the backward one, where period is 7 days. Available for crypto | [optional] 
**RollingChange** | Pointer to **string** | Percent change in price between the current and the backward one, where period specified in request param rolling_period. Available for crypto | [optional] 
**IsMarketOpen** | Pointer to **bool** | True if market is open; false if closed | [optional] 
**FiftyTwoWeek** | Pointer to [**GetQuote200ResponseFiftyTwoWeek**](GetQuote200ResponseFiftyTwoWeek.md) |  | [optional] 
**ExtendedChange** | Pointer to **string** | Diff between the regular close price and the latest extended price. Displayed only if prepost is true | [optional] 
**ExtendedPercentChange** | Pointer to **string** | Percent change in price between the regular close price and the latest extended price. Displayed only if prepost is true | [optional] 
**ExtendedPrice** | Pointer to **string** | Latest extended price. Displayed only if prepost is true | [optional] 
**ExtendedTimestamp** | Pointer to **string** | Unix timestamp of the last extended price. Displayed only if prepost is true | [optional] 

## Methods

### NewGetQuote200Response

`func NewGetQuote200Response() *GetQuote200Response`

NewGetQuote200Response instantiates a new GetQuote200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuote200ResponseWithDefaults

`func NewGetQuote200ResponseWithDefaults() *GetQuote200Response`

NewGetQuote200ResponseWithDefaults instantiates a new GetQuote200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetQuote200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetQuote200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetQuote200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetQuote200Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *GetQuote200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetQuote200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetQuote200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetQuote200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchange

`func (o *GetQuote200Response) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetQuote200Response) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetQuote200Response) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetQuote200Response) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *GetQuote200Response) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetQuote200Response) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetQuote200Response) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetQuote200Response) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetCurrency

`func (o *GetQuote200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetQuote200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetQuote200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetQuote200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDatetime

`func (o *GetQuote200Response) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetQuote200Response) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetQuote200Response) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetQuote200Response) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetQuote200Response) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetQuote200Response) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetQuote200Response) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetQuote200Response) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLastQuoteAt

`func (o *GetQuote200Response) GetLastQuoteAt() int64`

GetLastQuoteAt returns the LastQuoteAt field if non-nil, zero value otherwise.

### GetLastQuoteAtOk

`func (o *GetQuote200Response) GetLastQuoteAtOk() (*int64, bool)`

GetLastQuoteAtOk returns a tuple with the LastQuoteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuoteAt

`func (o *GetQuote200Response) SetLastQuoteAt(v int64)`

SetLastQuoteAt sets LastQuoteAt field to given value.

### HasLastQuoteAt

`func (o *GetQuote200Response) HasLastQuoteAt() bool`

HasLastQuoteAt returns a boolean if a field has been set.

### GetOpen

`func (o *GetQuote200Response) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *GetQuote200Response) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *GetQuote200Response) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *GetQuote200Response) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetHigh

`func (o *GetQuote200Response) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *GetQuote200Response) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *GetQuote200Response) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *GetQuote200Response) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *GetQuote200Response) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *GetQuote200Response) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *GetQuote200Response) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *GetQuote200Response) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetClose

`func (o *GetQuote200Response) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *GetQuote200Response) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *GetQuote200Response) SetClose(v string)`

SetClose sets Close field to given value.

### HasClose

`func (o *GetQuote200Response) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetVolume

`func (o *GetQuote200Response) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetQuote200Response) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetQuote200Response) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetQuote200Response) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetPreviousClose

`func (o *GetQuote200Response) GetPreviousClose() string`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *GetQuote200Response) GetPreviousCloseOk() (*string, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *GetQuote200Response) SetPreviousClose(v string)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *GetQuote200Response) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetChange

`func (o *GetQuote200Response) GetChange() string`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *GetQuote200Response) GetChangeOk() (*string, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *GetQuote200Response) SetChange(v string)`

SetChange sets Change field to given value.

### HasChange

`func (o *GetQuote200Response) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetPercentChange

`func (o *GetQuote200Response) GetPercentChange() string`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *GetQuote200Response) GetPercentChangeOk() (*string, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *GetQuote200Response) SetPercentChange(v string)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *GetQuote200Response) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetAverageVolume

`func (o *GetQuote200Response) GetAverageVolume() string`

GetAverageVolume returns the AverageVolume field if non-nil, zero value otherwise.

### GetAverageVolumeOk

`func (o *GetQuote200Response) GetAverageVolumeOk() (*string, bool)`

GetAverageVolumeOk returns a tuple with the AverageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageVolume

`func (o *GetQuote200Response) SetAverageVolume(v string)`

SetAverageVolume sets AverageVolume field to given value.

### HasAverageVolume

`func (o *GetQuote200Response) HasAverageVolume() bool`

HasAverageVolume returns a boolean if a field has been set.

### GetRolling1dChange

`func (o *GetQuote200Response) GetRolling1dChange() string`

GetRolling1dChange returns the Rolling1dChange field if non-nil, zero value otherwise.

### GetRolling1dChangeOk

`func (o *GetQuote200Response) GetRolling1dChangeOk() (*string, bool)`

GetRolling1dChangeOk returns a tuple with the Rolling1dChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolling1dChange

`func (o *GetQuote200Response) SetRolling1dChange(v string)`

SetRolling1dChange sets Rolling1dChange field to given value.

### HasRolling1dChange

`func (o *GetQuote200Response) HasRolling1dChange() bool`

HasRolling1dChange returns a boolean if a field has been set.

### GetRolling7dChange

`func (o *GetQuote200Response) GetRolling7dChange() string`

GetRolling7dChange returns the Rolling7dChange field if non-nil, zero value otherwise.

### GetRolling7dChangeOk

`func (o *GetQuote200Response) GetRolling7dChangeOk() (*string, bool)`

GetRolling7dChangeOk returns a tuple with the Rolling7dChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolling7dChange

`func (o *GetQuote200Response) SetRolling7dChange(v string)`

SetRolling7dChange sets Rolling7dChange field to given value.

### HasRolling7dChange

`func (o *GetQuote200Response) HasRolling7dChange() bool`

HasRolling7dChange returns a boolean if a field has been set.

### GetRollingChange

`func (o *GetQuote200Response) GetRollingChange() string`

GetRollingChange returns the RollingChange field if non-nil, zero value otherwise.

### GetRollingChangeOk

`func (o *GetQuote200Response) GetRollingChangeOk() (*string, bool)`

GetRollingChangeOk returns a tuple with the RollingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingChange

`func (o *GetQuote200Response) SetRollingChange(v string)`

SetRollingChange sets RollingChange field to given value.

### HasRollingChange

`func (o *GetQuote200Response) HasRollingChange() bool`

HasRollingChange returns a boolean if a field has been set.

### GetIsMarketOpen

`func (o *GetQuote200Response) GetIsMarketOpen() bool`

GetIsMarketOpen returns the IsMarketOpen field if non-nil, zero value otherwise.

### GetIsMarketOpenOk

`func (o *GetQuote200Response) GetIsMarketOpenOk() (*bool, bool)`

GetIsMarketOpenOk returns a tuple with the IsMarketOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketOpen

`func (o *GetQuote200Response) SetIsMarketOpen(v bool)`

SetIsMarketOpen sets IsMarketOpen field to given value.

### HasIsMarketOpen

`func (o *GetQuote200Response) HasIsMarketOpen() bool`

HasIsMarketOpen returns a boolean if a field has been set.

### GetFiftyTwoWeek

`func (o *GetQuote200Response) GetFiftyTwoWeek() GetQuote200ResponseFiftyTwoWeek`

GetFiftyTwoWeek returns the FiftyTwoWeek field if non-nil, zero value otherwise.

### GetFiftyTwoWeekOk

`func (o *GetQuote200Response) GetFiftyTwoWeekOk() (*GetQuote200ResponseFiftyTwoWeek, bool)`

GetFiftyTwoWeekOk returns a tuple with the FiftyTwoWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeek

`func (o *GetQuote200Response) SetFiftyTwoWeek(v GetQuote200ResponseFiftyTwoWeek)`

SetFiftyTwoWeek sets FiftyTwoWeek field to given value.

### HasFiftyTwoWeek

`func (o *GetQuote200Response) HasFiftyTwoWeek() bool`

HasFiftyTwoWeek returns a boolean if a field has been set.

### GetExtendedChange

`func (o *GetQuote200Response) GetExtendedChange() string`

GetExtendedChange returns the ExtendedChange field if non-nil, zero value otherwise.

### GetExtendedChangeOk

`func (o *GetQuote200Response) GetExtendedChangeOk() (*string, bool)`

GetExtendedChangeOk returns a tuple with the ExtendedChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedChange

`func (o *GetQuote200Response) SetExtendedChange(v string)`

SetExtendedChange sets ExtendedChange field to given value.

### HasExtendedChange

`func (o *GetQuote200Response) HasExtendedChange() bool`

HasExtendedChange returns a boolean if a field has been set.

### GetExtendedPercentChange

`func (o *GetQuote200Response) GetExtendedPercentChange() string`

GetExtendedPercentChange returns the ExtendedPercentChange field if non-nil, zero value otherwise.

### GetExtendedPercentChangeOk

`func (o *GetQuote200Response) GetExtendedPercentChangeOk() (*string, bool)`

GetExtendedPercentChangeOk returns a tuple with the ExtendedPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPercentChange

`func (o *GetQuote200Response) SetExtendedPercentChange(v string)`

SetExtendedPercentChange sets ExtendedPercentChange field to given value.

### HasExtendedPercentChange

`func (o *GetQuote200Response) HasExtendedPercentChange() bool`

HasExtendedPercentChange returns a boolean if a field has been set.

### GetExtendedPrice

`func (o *GetQuote200Response) GetExtendedPrice() string`

GetExtendedPrice returns the ExtendedPrice field if non-nil, zero value otherwise.

### GetExtendedPriceOk

`func (o *GetQuote200Response) GetExtendedPriceOk() (*string, bool)`

GetExtendedPriceOk returns a tuple with the ExtendedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedPrice

`func (o *GetQuote200Response) SetExtendedPrice(v string)`

SetExtendedPrice sets ExtendedPrice field to given value.

### HasExtendedPrice

`func (o *GetQuote200Response) HasExtendedPrice() bool`

HasExtendedPrice returns a boolean if a field has been set.

### GetExtendedTimestamp

`func (o *GetQuote200Response) GetExtendedTimestamp() string`

GetExtendedTimestamp returns the ExtendedTimestamp field if non-nil, zero value otherwise.

### GetExtendedTimestampOk

`func (o *GetQuote200Response) GetExtendedTimestampOk() (*string, bool)`

GetExtendedTimestampOk returns a tuple with the ExtendedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTimestamp

`func (o *GetQuote200Response) SetExtendedTimestamp(v string)`

SetExtendedTimestamp sets ExtendedTimestamp field to given value.

### HasExtendedTimestamp

`func (o *GetQuote200Response) HasExtendedTimestamp() bool`

HasExtendedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


