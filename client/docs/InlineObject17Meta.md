# InlineObject17Meta

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
**Indicator** | Pointer to [**InlineObject17MetaIndicator**](InlineObject17MetaIndicator.md) |  | [optional] 

## Methods

### NewInlineObject17Meta

`func NewInlineObject17Meta() *InlineObject17Meta`

NewInlineObject17Meta instantiates a new InlineObject17Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject17MetaWithDefaults

`func NewInlineObject17MetaWithDefaults() *InlineObject17Meta`

NewInlineObject17MetaWithDefaults instantiates a new InlineObject17Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineObject17Meta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineObject17Meta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineObject17Meta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InlineObject17Meta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInterval

`func (o *InlineObject17Meta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineObject17Meta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineObject17Meta) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InlineObject17Meta) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCurrency

`func (o *InlineObject17Meta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineObject17Meta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineObject17Meta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineObject17Meta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeTimezone

`func (o *InlineObject17Meta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *InlineObject17Meta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *InlineObject17Meta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.

### HasExchangeTimezone

`func (o *InlineObject17Meta) HasExchangeTimezone() bool`

HasExchangeTimezone returns a boolean if a field has been set.

### GetExchange

`func (o *InlineObject17Meta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineObject17Meta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineObject17Meta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InlineObject17Meta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *InlineObject17Meta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *InlineObject17Meta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *InlineObject17Meta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *InlineObject17Meta) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *InlineObject17Meta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject17Meta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject17Meta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject17Meta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIndicator

`func (o *InlineObject17Meta) GetIndicator() InlineObject17MetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *InlineObject17Meta) GetIndicatorOk() (*InlineObject17MetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *InlineObject17Meta) SetIndicator(v InlineObject17MetaIndicator)`

SetIndicator sets Indicator field to given value.

### HasIndicator

`func (o *InlineObject17Meta) HasIndicator() bool`

HasIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


