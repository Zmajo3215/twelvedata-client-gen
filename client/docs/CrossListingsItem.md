# CrossListingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Ticker symbol of instrument | [optional] 
**Name** | Pointer to **string** | Name of symbol | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 

## Methods

### NewCrossListingsItem

`func NewCrossListingsItem() *CrossListingsItem`

NewCrossListingsItem instantiates a new CrossListingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossListingsItemWithDefaults

`func NewCrossListingsItemWithDefaults() *CrossListingsItem`

NewCrossListingsItemWithDefaults instantiates a new CrossListingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CrossListingsItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CrossListingsItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CrossListingsItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CrossListingsItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *CrossListingsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrossListingsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrossListingsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrossListingsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchange

`func (o *CrossListingsItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *CrossListingsItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *CrossListingsItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *CrossListingsItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *CrossListingsItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *CrossListingsItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *CrossListingsItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *CrossListingsItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


