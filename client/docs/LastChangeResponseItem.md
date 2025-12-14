# LastChangeResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Ticker of the company | [optional] 
**MicCode** | Pointer to **string** | Market Identifier Code (MIC) under ISO 10383 standard | [optional] 
**LastChange** | Pointer to **time.Time** | The date and time of last changes, in &#x60;2006-01-02 15:04:05&#x60; format | [optional] 

## Methods

### NewLastChangeResponseItem

`func NewLastChangeResponseItem() *LastChangeResponseItem`

NewLastChangeResponseItem instantiates a new LastChangeResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastChangeResponseItemWithDefaults

`func NewLastChangeResponseItemWithDefaults() *LastChangeResponseItem`

NewLastChangeResponseItemWithDefaults instantiates a new LastChangeResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LastChangeResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LastChangeResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LastChangeResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *LastChangeResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMicCode

`func (o *LastChangeResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *LastChangeResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *LastChangeResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *LastChangeResponseItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetLastChange

`func (o *LastChangeResponseItem) GetLastChange() time.Time`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *LastChangeResponseItem) GetLastChangeOk() (*time.Time, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *LastChangeResponseItem) SetLastChange(v time.Time)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *LastChangeResponseItem) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


