# FundResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Instrument symbol (ticker) | [optional] 
**Name** | Pointer to **string** | Full name of the fund | [optional] 
**Country** | Pointer to **string** | Country where the fund is located | [optional] 
**Currency** | Pointer to **string** | Currency of the fund according to the ISO 4217 standard | [optional] 
**Exchange** | Pointer to **string** | Exchange where the fund is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Type** | Pointer to **string** | Type of the fund | [optional] 
**FigiCode** | Pointer to **string** | Financial instrument global identifier (FIGI) | [optional] 
**CfiCode** | Pointer to **string** | Classification of Financial Instruments (CFI) | [optional] 
**Isin** | Pointer to **string** | International securities identification number (ISIN) | [optional] 
**Cusip** | Pointer to **string** | A unique nine-character alphanumeric code used to identify financial securities, ensuring accurate data retrieval for the specified asset | [optional] 
**Access** | Pointer to [**EtfResponseItemAccess**](EtfResponseItemAccess.md) |  | [optional] 

## Methods

### NewFundResponseItem

`func NewFundResponseItem() *FundResponseItem`

NewFundResponseItem instantiates a new FundResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundResponseItemWithDefaults

`func NewFundResponseItemWithDefaults() *FundResponseItem`

NewFundResponseItemWithDefaults instantiates a new FundResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FundResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FundResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FundResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FundResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *FundResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FundResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FundResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FundResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountry

`func (o *FundResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FundResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FundResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FundResponseItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *FundResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FundResponseItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *FundResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *FundResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *FundResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *FundResponseItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *FundResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *FundResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *FundResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *FundResponseItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *FundResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundResponseItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FundResponseItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFigiCode

`func (o *FundResponseItem) GetFigiCode() string`

GetFigiCode returns the FigiCode field if non-nil, zero value otherwise.

### GetFigiCodeOk

`func (o *FundResponseItem) GetFigiCodeOk() (*string, bool)`

GetFigiCodeOk returns a tuple with the FigiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigiCode

`func (o *FundResponseItem) SetFigiCode(v string)`

SetFigiCode sets FigiCode field to given value.

### HasFigiCode

`func (o *FundResponseItem) HasFigiCode() bool`

HasFigiCode returns a boolean if a field has been set.

### GetCfiCode

`func (o *FundResponseItem) GetCfiCode() string`

GetCfiCode returns the CfiCode field if non-nil, zero value otherwise.

### GetCfiCodeOk

`func (o *FundResponseItem) GetCfiCodeOk() (*string, bool)`

GetCfiCodeOk returns a tuple with the CfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfiCode

`func (o *FundResponseItem) SetCfiCode(v string)`

SetCfiCode sets CfiCode field to given value.

### HasCfiCode

`func (o *FundResponseItem) HasCfiCode() bool`

HasCfiCode returns a boolean if a field has been set.

### GetIsin

`func (o *FundResponseItem) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *FundResponseItem) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *FundResponseItem) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *FundResponseItem) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *FundResponseItem) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *FundResponseItem) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *FundResponseItem) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *FundResponseItem) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetAccess

`func (o *FundResponseItem) GetAccess() EtfResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FundResponseItem) GetAccessOk() (*EtfResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FundResponseItem) SetAccess(v EtfResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FundResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


