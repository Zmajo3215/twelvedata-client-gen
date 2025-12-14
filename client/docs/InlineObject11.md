# InlineObject11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**InlineObject11Meta**](InlineObject11Meta.md) |  | [optional] 
**Values** | Pointer to [**[]InlineObject11ValuesInner**](InlineObject11ValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewInlineObject11

`func NewInlineObject11() *InlineObject11`

NewInlineObject11 instantiates a new InlineObject11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11WithDefaults

`func NewInlineObject11WithDefaults() *InlineObject11`

NewInlineObject11WithDefaults instantiates a new InlineObject11 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject11) GetMeta() InlineObject11Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject11) GetMetaOk() (*InlineObject11Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject11) SetMeta(v InlineObject11Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineObject11) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *InlineObject11) GetValues() []InlineObject11ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject11) GetValuesOk() (*[]InlineObject11ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject11) SetValues(v []InlineObject11ValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *InlineObject11) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *InlineObject11) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject11) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject11) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineObject11) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


