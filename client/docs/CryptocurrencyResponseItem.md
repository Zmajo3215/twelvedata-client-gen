# CryptocurrencyResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Cryptocurrency pair codes with slash(/) delimiter | [optional] 
**AvailableExchanges** | Pointer to **[]string** | List of exchanges where the cryptocurrency is available | [optional] 
**CurrencyBase** | Pointer to **string** | Base currency of the cryptocurrency pair | [optional] 
**CurrencyQuote** | Pointer to **string** | Quote currency of the cryptocurrency pair | [optional] 

## Methods

### NewCryptocurrencyResponseItem

`func NewCryptocurrencyResponseItem() *CryptocurrencyResponseItem`

NewCryptocurrencyResponseItem instantiates a new CryptocurrencyResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptocurrencyResponseItemWithDefaults

`func NewCryptocurrencyResponseItemWithDefaults() *CryptocurrencyResponseItem`

NewCryptocurrencyResponseItemWithDefaults instantiates a new CryptocurrencyResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CryptocurrencyResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CryptocurrencyResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CryptocurrencyResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CryptocurrencyResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAvailableExchanges

`func (o *CryptocurrencyResponseItem) GetAvailableExchanges() []string`

GetAvailableExchanges returns the AvailableExchanges field if non-nil, zero value otherwise.

### GetAvailableExchangesOk

`func (o *CryptocurrencyResponseItem) GetAvailableExchangesOk() (*[]string, bool)`

GetAvailableExchangesOk returns a tuple with the AvailableExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableExchanges

`func (o *CryptocurrencyResponseItem) SetAvailableExchanges(v []string)`

SetAvailableExchanges sets AvailableExchanges field to given value.

### HasAvailableExchanges

`func (o *CryptocurrencyResponseItem) HasAvailableExchanges() bool`

HasAvailableExchanges returns a boolean if a field has been set.

### GetCurrencyBase

`func (o *CryptocurrencyResponseItem) GetCurrencyBase() string`

GetCurrencyBase returns the CurrencyBase field if non-nil, zero value otherwise.

### GetCurrencyBaseOk

`func (o *CryptocurrencyResponseItem) GetCurrencyBaseOk() (*string, bool)`

GetCurrencyBaseOk returns a tuple with the CurrencyBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyBase

`func (o *CryptocurrencyResponseItem) SetCurrencyBase(v string)`

SetCurrencyBase sets CurrencyBase field to given value.

### HasCurrencyBase

`func (o *CryptocurrencyResponseItem) HasCurrencyBase() bool`

HasCurrencyBase returns a boolean if a field has been set.

### GetCurrencyQuote

`func (o *CryptocurrencyResponseItem) GetCurrencyQuote() string`

GetCurrencyQuote returns the CurrencyQuote field if non-nil, zero value otherwise.

### GetCurrencyQuoteOk

`func (o *CryptocurrencyResponseItem) GetCurrencyQuoteOk() (*string, bool)`

GetCurrencyQuoteOk returns a tuple with the CurrencyQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyQuote

`func (o *CryptocurrencyResponseItem) SetCurrencyQuote(v string)`

SetCurrencyQuote sets CurrencyQuote field to given value.

### HasCurrencyQuote

`func (o *CryptocurrencyResponseItem) HasCurrencyQuote() bool`

HasCurrencyQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


