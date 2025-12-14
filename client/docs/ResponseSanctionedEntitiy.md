# ResponseSanctionedEntitiy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The instrument symbol ticker | [optional] 
**Name** | Pointer to **string** | The instrument name | [optional] 
**MicCode** | Pointer to **string** | Market Identifier Code (MIC) under ISO 10383 standard | [optional] 
**Country** | Pointer to **string** | Country name | [optional] 
**Sanction** | Pointer to [**ResponseSanctionItem**](ResponseSanctionItem.md) |  | [optional] 

## Methods

### NewResponseSanctionedEntitiy

`func NewResponseSanctionedEntitiy() *ResponseSanctionedEntitiy`

NewResponseSanctionedEntitiy instantiates a new ResponseSanctionedEntitiy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSanctionedEntitiyWithDefaults

`func NewResponseSanctionedEntitiyWithDefaults() *ResponseSanctionedEntitiy`

NewResponseSanctionedEntitiyWithDefaults instantiates a new ResponseSanctionedEntitiy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ResponseSanctionedEntitiy) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ResponseSanctionedEntitiy) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ResponseSanctionedEntitiy) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ResponseSanctionedEntitiy) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *ResponseSanctionedEntitiy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseSanctionedEntitiy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseSanctionedEntitiy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseSanctionedEntitiy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMicCode

`func (o *ResponseSanctionedEntitiy) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *ResponseSanctionedEntitiy) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *ResponseSanctionedEntitiy) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *ResponseSanctionedEntitiy) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseSanctionedEntitiy) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseSanctionedEntitiy) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseSanctionedEntitiy) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseSanctionedEntitiy) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSanction

`func (o *ResponseSanctionedEntitiy) GetSanction() ResponseSanctionItem`

GetSanction returns the Sanction field if non-nil, zero value otherwise.

### GetSanctionOk

`func (o *ResponseSanctionedEntitiy) GetSanctionOk() (*ResponseSanctionItem, bool)`

GetSanctionOk returns a tuple with the Sanction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanction

`func (o *ResponseSanctionedEntitiy) SetSanction(v ResponseSanctionItem)`

SetSanction sets Sanction field to given value.

### HasSanction

`func (o *ResponseSanctionedEntitiy) HasSanction() bool`

HasSanction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


