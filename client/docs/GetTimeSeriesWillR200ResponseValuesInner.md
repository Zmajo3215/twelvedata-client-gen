# GetTimeSeriesWillR200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Willr** | Pointer to **string** | Williams %R value | [optional] 

## Methods

### NewGetTimeSeriesWillR200ResponseValuesInner

`func NewGetTimeSeriesWillR200ResponseValuesInner() *GetTimeSeriesWillR200ResponseValuesInner`

NewGetTimeSeriesWillR200ResponseValuesInner instantiates a new GetTimeSeriesWillR200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWillR200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesWillR200ResponseValuesInnerWithDefaults() *GetTimeSeriesWillR200ResponseValuesInner`

NewGetTimeSeriesWillR200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWillR200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesWillR200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesWillR200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesWillR200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesWillR200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetWillr

`func (o *GetTimeSeriesWillR200ResponseValuesInner) GetWillr() string`

GetWillr returns the Willr field if non-nil, zero value otherwise.

### GetWillrOk

`func (o *GetTimeSeriesWillR200ResponseValuesInner) GetWillrOk() (*string, bool)`

GetWillrOk returns a tuple with the Willr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillr

`func (o *GetTimeSeriesWillR200ResponseValuesInner) SetWillr(v string)`

SetWillr sets Willr field to given value.

### HasWillr

`func (o *GetTimeSeriesWillR200ResponseValuesInner) HasWillr() bool`

HasWillr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


