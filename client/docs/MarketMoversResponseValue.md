# MarketMoversResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The exchange symbol ticker | [optional] 
**Name** | Pointer to **string** | The official name of the instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Datetime** | Pointer to **string** | The last updated datetime timestamp | [optional] 
**Last** | Pointer to **float64** | The latest available price for the symbol today | [optional] 
**High** | Pointer to **float64** | The highest price for the symbol today | [optional] 
**Low** | Pointer to **float64** | The lowest price for the symbol today | [optional] 
**Volume** | Pointer to **int64** | The trading volume of the symbol today | [optional] 
**Change** | Pointer to **float64** | The value of the change since the previous day | [optional] 
**PercentChange** | Pointer to **float64** | The percentage change since the previous day | [optional] 

## Methods

### NewMarketMoversResponseValue

`func NewMarketMoversResponseValue() *MarketMoversResponseValue`

NewMarketMoversResponseValue instantiates a new MarketMoversResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMoversResponseValueWithDefaults

`func NewMarketMoversResponseValueWithDefaults() *MarketMoversResponseValue`

NewMarketMoversResponseValueWithDefaults instantiates a new MarketMoversResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarketMoversResponseValue) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketMoversResponseValue) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketMoversResponseValue) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarketMoversResponseValue) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *MarketMoversResponseValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketMoversResponseValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketMoversResponseValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketMoversResponseValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchange

`func (o *MarketMoversResponseValue) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketMoversResponseValue) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketMoversResponseValue) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *MarketMoversResponseValue) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *MarketMoversResponseValue) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *MarketMoversResponseValue) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *MarketMoversResponseValue) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *MarketMoversResponseValue) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetDatetime

`func (o *MarketMoversResponseValue) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MarketMoversResponseValue) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MarketMoversResponseValue) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *MarketMoversResponseValue) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetLast

`func (o *MarketMoversResponseValue) GetLast() float64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *MarketMoversResponseValue) GetLastOk() (*float64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *MarketMoversResponseValue) SetLast(v float64)`

SetLast sets Last field to given value.

### HasLast

`func (o *MarketMoversResponseValue) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetHigh

`func (o *MarketMoversResponseValue) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *MarketMoversResponseValue) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *MarketMoversResponseValue) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *MarketMoversResponseValue) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *MarketMoversResponseValue) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *MarketMoversResponseValue) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *MarketMoversResponseValue) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *MarketMoversResponseValue) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetVolume

`func (o *MarketMoversResponseValue) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketMoversResponseValue) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketMoversResponseValue) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketMoversResponseValue) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetChange

`func (o *MarketMoversResponseValue) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *MarketMoversResponseValue) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *MarketMoversResponseValue) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *MarketMoversResponseValue) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetPercentChange

`func (o *MarketMoversResponseValue) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *MarketMoversResponseValue) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *MarketMoversResponseValue) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *MarketMoversResponseValue) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


