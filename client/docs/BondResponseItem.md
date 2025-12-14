# BondResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Bond symbol | [optional] 
**Name** | Pointer to **string** | Full name of the bond | [optional] 
**Country** | Pointer to **string** | Country where the bond is located | [optional] 
**Currency** | Pointer to **string** | Currency of the bond according to the ISO 4217 standard | [optional] 
**Exchange** | Pointer to **string** | Exchange where the bond is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Type** | Pointer to **string** | Type of the bond | [optional] 
**Access** | Pointer to [**BondsResponseItemAccess**](BondsResponseItemAccess.md) |  | [optional] 

## Methods

### NewBondResponseItem

`func NewBondResponseItem() *BondResponseItem`

NewBondResponseItem instantiates a new BondResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondResponseItemWithDefaults

`func NewBondResponseItemWithDefaults() *BondResponseItem`

NewBondResponseItemWithDefaults instantiates a new BondResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *BondResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BondResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BondResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *BondResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *BondResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BondResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountry

`func (o *BondResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BondResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BondResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BondResponseItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *BondResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BondResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BondResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BondResponseItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *BondResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *BondResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *BondResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *BondResponseItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *BondResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *BondResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *BondResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *BondResponseItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *BondResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BondResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BondResponseItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BondResponseItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccess

`func (o *BondResponseItem) GetAccess() BondsResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *BondResponseItem) GetAccessOk() (*BondsResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *BondResponseItem) SetAccess(v BondsResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *BondResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


