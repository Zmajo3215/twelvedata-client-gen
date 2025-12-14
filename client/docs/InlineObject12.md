# InlineObject12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**InlineObject12Meta**](InlineObject12Meta.md) |  | [optional] 
**Values** | Pointer to [**[]InlineObject12ValuesInner**](InlineObject12ValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewInlineObject12

`func NewInlineObject12() *InlineObject12`

NewInlineObject12 instantiates a new InlineObject12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject12WithDefaults

`func NewInlineObject12WithDefaults() *InlineObject12`

NewInlineObject12WithDefaults instantiates a new InlineObject12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *InlineObject12) GetMeta() InlineObject12Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject12) GetMetaOk() (*InlineObject12Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject12) SetMeta(v InlineObject12Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineObject12) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *InlineObject12) GetValues() []InlineObject12ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject12) GetValuesOk() (*[]InlineObject12ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject12) SetValues(v []InlineObject12ValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *InlineObject12) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *InlineObject12) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject12) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject12) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineObject12) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


