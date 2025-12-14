# GetTimeSeriesHtTrendline200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The ticker symbol of an instrument for which data was requested. | [optional] 
**Interval** | Pointer to **string** | The time gap between consecutive data points. | [optional] 
**Currency** | Pointer to **string** | The currency of a traded instrument. | [optional] 
**ExchangeTimezone** | Pointer to **string** | The timezone of the exchange where the instrument is traded. | [optional] 
**Exchange** | Pointer to **string** | The exchange name where the instrument is traded. | [optional] 
**MicCode** | Pointer to **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded. | [optional] 
**Type** | Pointer to **string** | The asset class to which the instrument belongs. | [optional] 
**Indicator** | Pointer to [**GetTimeSeriesHtTrendline200ResponseMetaIndicator**](GetTimeSeriesHtTrendline200ResponseMetaIndicator.md) |  | [optional] 

## Methods

### NewGetTimeSeriesHtTrendline200ResponseMeta

`func NewGetTimeSeriesHtTrendline200ResponseMeta() *GetTimeSeriesHtTrendline200ResponseMeta`

NewGetTimeSeriesHtTrendline200ResponseMeta instantiates a new GetTimeSeriesHtTrendline200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtTrendline200ResponseMetaWithDefaults

`func NewGetTimeSeriesHtTrendline200ResponseMetaWithDefaults() *GetTimeSeriesHtTrendline200ResponseMeta`

NewGetTimeSeriesHtTrendline200ResponseMetaWithDefaults instantiates a new GetTimeSeriesHtTrendline200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInterval

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCurrency

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeTimezone

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.

### HasExchangeTimezone

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasExchangeTimezone() bool`

HasExchangeTimezone returns a boolean if a field has been set.

### GetExchange

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndicator

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetIndicator() GetTimeSeriesHtTrendline200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesHtTrendline200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) SetIndicator(v GetTimeSeriesHtTrendline200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *GetTimeSeriesHtTrendline200ResponseMeta) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


