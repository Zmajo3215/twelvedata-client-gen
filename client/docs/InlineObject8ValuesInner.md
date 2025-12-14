# InlineObject8ValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Acos** | Pointer to **string** | ACOS value | [optional] 

## Methods

### NewInlineObject8ValuesInner

`func NewInlineObject8ValuesInner() *InlineObject8ValuesInner`

NewInlineObject8ValuesInner instantiates a new InlineObject8ValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject8ValuesInnerWithDefaults

`func NewInlineObject8ValuesInnerWithDefaults() *InlineObject8ValuesInner`

NewInlineObject8ValuesInnerWithDefaults instantiates a new InlineObject8ValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *InlineObject8ValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *InlineObject8ValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *InlineObject8ValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *InlineObject8ValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetAcos

`func (o *InlineObject8ValuesInner) GetAcos() string`

GetAcos returns the Acos field if non-nil, zero value otherwise.

### GetAcosOk

`func (o *InlineObject8ValuesInner) GetAcosOk() (*string, bool)`

GetAcosOk returns a tuple with the Acos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcos

`func (o *InlineObject8ValuesInner) SetAcos(v string)`

SetAcos sets Acos field to given value.

### HasAcos

`func (o *InlineObject8ValuesInner) HasAcos() bool`

HasAcos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


