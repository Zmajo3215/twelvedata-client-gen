# CrossMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseInstrument** | Pointer to **string** | Base instrument symbol | [optional] 
**BaseCurrency** | Pointer to **string** | Base currency | [optional] 
**BaseExchange** | Pointer to **string** | Base exchange | [optional] 
**Interval** | Pointer to **string** | Interval between two consecutive points in time series | [optional] 
**QuoteInstrument** | Pointer to **string** | Quote instrument symbol | [optional] 
**QuoteCurrency** | Pointer to **string** | Quote currency | [optional] 
**QuoteExchange** | Pointer to **string** | Quote exchange | [optional] 

## Methods

### NewCrossMeta

`func NewCrossMeta() *CrossMeta`

NewCrossMeta instantiates a new CrossMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossMetaWithDefaults

`func NewCrossMetaWithDefaults() *CrossMeta`

NewCrossMetaWithDefaults instantiates a new CrossMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseInstrument

`func (o *CrossMeta) GetBaseInstrument() string`

GetBaseInstrument returns the BaseInstrument field if non-nil, zero value otherwise.

### GetBaseInstrumentOk

`func (o *CrossMeta) GetBaseInstrumentOk() (*string, bool)`

GetBaseInstrumentOk returns a tuple with the BaseInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseInstrument

`func (o *CrossMeta) SetBaseInstrument(v string)`

SetBaseInstrument sets BaseInstrument field to given value.

### HasBaseInstrument

`func (o *CrossMeta) HasBaseInstrument() bool`

HasBaseInstrument returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *CrossMeta) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CrossMeta) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CrossMeta) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *CrossMeta) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetBaseExchange

`func (o *CrossMeta) GetBaseExchange() string`

GetBaseExchange returns the BaseExchange field if non-nil, zero value otherwise.

### GetBaseExchangeOk

`func (o *CrossMeta) GetBaseExchangeOk() (*string, bool)`

GetBaseExchangeOk returns a tuple with the BaseExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseExchange

`func (o *CrossMeta) SetBaseExchange(v string)`

SetBaseExchange sets BaseExchange field to given value.

### HasBaseExchange

`func (o *CrossMeta) HasBaseExchange() bool`

HasBaseExchange returns a boolean if a field has been set.

### GetInterval

`func (o *CrossMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CrossMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CrossMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CrossMeta) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetQuoteInstrument

`func (o *CrossMeta) GetQuoteInstrument() string`

GetQuoteInstrument returns the QuoteInstrument field if non-nil, zero value otherwise.

### GetQuoteInstrumentOk

`func (o *CrossMeta) GetQuoteInstrumentOk() (*string, bool)`

GetQuoteInstrumentOk returns a tuple with the QuoteInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteInstrument

`func (o *CrossMeta) SetQuoteInstrument(v string)`

SetQuoteInstrument sets QuoteInstrument field to given value.

### HasQuoteInstrument

`func (o *CrossMeta) HasQuoteInstrument() bool`

HasQuoteInstrument returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *CrossMeta) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *CrossMeta) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *CrossMeta) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *CrossMeta) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetQuoteExchange

`func (o *CrossMeta) GetQuoteExchange() string`

GetQuoteExchange returns the QuoteExchange field if non-nil, zero value otherwise.

### GetQuoteExchangeOk

`func (o *CrossMeta) GetQuoteExchangeOk() (*string, bool)`

GetQuoteExchangeOk returns a tuple with the QuoteExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExchange

`func (o *CrossMeta) SetQuoteExchange(v string)`

SetQuoteExchange sets QuoteExchange field to given value.

### HasQuoteExchange

`func (o *CrossMeta) HasQuoteExchange() bool`

HasQuoteExchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


