# GetTaxInfo200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The ticker symbol of an instrument for which data was requested | [optional] 
**Name** | Pointer to **string** | The instrument name | [optional] 
**Exchange** | Pointer to **string** | The exchange name where the instrument is traded | [optional] 
**MicCode** | Pointer to **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded | [optional] 
**Country** | Pointer to **string** | The instrument country name | [optional] 

## Methods

### NewGetTaxInfo200ResponseMeta

`func NewGetTaxInfo200ResponseMeta() *GetTaxInfo200ResponseMeta`

NewGetTaxInfo200ResponseMeta instantiates a new GetTaxInfo200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxInfo200ResponseMetaWithDefaults

`func NewGetTaxInfo200ResponseMetaWithDefaults() *GetTaxInfo200ResponseMeta`

NewGetTaxInfo200ResponseMetaWithDefaults instantiates a new GetTaxInfo200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTaxInfo200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTaxInfo200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTaxInfo200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetTaxInfo200ResponseMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *GetTaxInfo200ResponseMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaxInfo200ResponseMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaxInfo200ResponseMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTaxInfo200ResponseMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchange

`func (o *GetTaxInfo200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTaxInfo200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTaxInfo200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetTaxInfo200ResponseMeta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *GetTaxInfo200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTaxInfo200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTaxInfo200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetTaxInfo200ResponseMeta) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetCountry

`func (o *GetTaxInfo200ResponseMeta) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetTaxInfo200ResponseMeta) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetTaxInfo200ResponseMeta) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetTaxInfo200ResponseMeta) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


