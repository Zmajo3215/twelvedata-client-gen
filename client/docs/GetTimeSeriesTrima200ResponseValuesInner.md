# GetTimeSeriesTrima200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Trima** | Pointer to **string** | TRIMA value | [optional] 

## Methods

### NewGetTimeSeriesTrima200ResponseValuesInner

`func NewGetTimeSeriesTrima200ResponseValuesInner() *GetTimeSeriesTrima200ResponseValuesInner`

NewGetTimeSeriesTrima200ResponseValuesInner instantiates a new GetTimeSeriesTrima200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTrima200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesTrima200ResponseValuesInnerWithDefaults() *GetTimeSeriesTrima200ResponseValuesInner`

NewGetTimeSeriesTrima200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTrima200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesTrima200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesTrima200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesTrima200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesTrima200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetTrima

`func (o *GetTimeSeriesTrima200ResponseValuesInner) GetTrima() string`

GetTrima returns the Trima field if non-nil, zero value otherwise.

### GetTrimaOk

`func (o *GetTimeSeriesTrima200ResponseValuesInner) GetTrimaOk() (*string, bool)`

GetTrimaOk returns a tuple with the Trima field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrima

`func (o *GetTimeSeriesTrima200ResponseValuesInner) SetTrima(v string)`

SetTrima sets Trima field to given value.

### HasTrima

`func (o *GetTimeSeriesTrima200ResponseValuesInner) HasTrima() bool`

HasTrima returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


