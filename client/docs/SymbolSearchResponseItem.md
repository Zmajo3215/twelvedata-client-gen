# SymbolSearchResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Ticker symbol of instrument | [optional] 
**InstrumentName** | Pointer to **string** | Name of exchange | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**ExchangeTimezone** | Pointer to **string** | Time zone where exchange is located | [optional] 
**InstrumentType** | Pointer to **string** | Type of instrument | [optional] 
**Country** | Pointer to **string** | Country to which stock exchange belongs to | [optional] 
**Currency** | Pointer to **string** | Currency in which the instrument is traded | [optional] 
**Access** | Pointer to [**SymbolSearchResponseItemAccess**](SymbolSearchResponseItemAccess.md) |  | [optional] 

## Methods

### NewSymbolSearchResponseItem

`func NewSymbolSearchResponseItem() *SymbolSearchResponseItem`

NewSymbolSearchResponseItem instantiates a new SymbolSearchResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolSearchResponseItemWithDefaults

`func NewSymbolSearchResponseItemWithDefaults() *SymbolSearchResponseItem`

NewSymbolSearchResponseItemWithDefaults instantiates a new SymbolSearchResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SymbolSearchResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SymbolSearchResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SymbolSearchResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SymbolSearchResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInstrumentName

`func (o *SymbolSearchResponseItem) GetInstrumentName() string`

GetInstrumentName returns the InstrumentName field if non-nil, zero value otherwise.

### GetInstrumentNameOk

`func (o *SymbolSearchResponseItem) GetInstrumentNameOk() (*string, bool)`

GetInstrumentNameOk returns a tuple with the InstrumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentName

`func (o *SymbolSearchResponseItem) SetInstrumentName(v string)`

SetInstrumentName sets InstrumentName field to given value.

### HasInstrumentName

`func (o *SymbolSearchResponseItem) HasInstrumentName() bool`

HasInstrumentName returns a boolean if a field has been set.

### GetExchange

`func (o *SymbolSearchResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SymbolSearchResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SymbolSearchResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *SymbolSearchResponseItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *SymbolSearchResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *SymbolSearchResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *SymbolSearchResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *SymbolSearchResponseItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetExchangeTimezone

`func (o *SymbolSearchResponseItem) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *SymbolSearchResponseItem) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *SymbolSearchResponseItem) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.

### HasExchangeTimezone

`func (o *SymbolSearchResponseItem) HasExchangeTimezone() bool`

HasExchangeTimezone returns a boolean if a field has been set.

### GetInstrumentType

`func (o *SymbolSearchResponseItem) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *SymbolSearchResponseItem) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *SymbolSearchResponseItem) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *SymbolSearchResponseItem) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetCountry

`func (o *SymbolSearchResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SymbolSearchResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SymbolSearchResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SymbolSearchResponseItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *SymbolSearchResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SymbolSearchResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SymbolSearchResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SymbolSearchResponseItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAccess

`func (o *SymbolSearchResponseItem) GetAccess() SymbolSearchResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SymbolSearchResponseItem) GetAccessOk() (*SymbolSearchResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SymbolSearchResponseItem) SetAccess(v SymbolSearchResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SymbolSearchResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


