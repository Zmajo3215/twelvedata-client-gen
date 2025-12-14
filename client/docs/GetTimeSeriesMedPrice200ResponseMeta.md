# GetTimeSeriesMedPrice200ResponseMeta

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
**Indicator** | Pointer to [**GetTimeSeriesMedPrice200ResponseMetaIndicator**](GetTimeSeriesMedPrice200ResponseMetaIndicator.md) |  | [optional] 

## Methods

### NewGetTimeSeriesMedPrice200ResponseMeta

`func NewGetTimeSeriesMedPrice200ResponseMeta() *GetTimeSeriesMedPrice200ResponseMeta`

NewGetTimeSeriesMedPrice200ResponseMeta instantiates a new GetTimeSeriesMedPrice200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMedPrice200ResponseMetaWithDefaults

`func NewGetTimeSeriesMedPrice200ResponseMetaWithDefaults() *GetTimeSeriesMedPrice200ResponseMeta`

NewGetTimeSeriesMedPrice200ResponseMetaWithDefaults instantiates a new GetTimeSeriesMedPrice200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInterval

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCurrency

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeTimezone

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.

### HasExchangeTimezone

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasExchangeTimezone() bool`

HasExchangeTimezone returns a boolean if a field has been set.

### GetExchange

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndicator

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetIndicator() GetTimeSeriesMedPrice200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesMedPrice200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesMedPrice200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesMedPrice200ResponseMeta) SetIndicator(v GetTimeSeriesMedPrice200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *GetTimeSeriesMedPrice200ResponseMeta) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


