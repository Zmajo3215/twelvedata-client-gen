# GetTimeSeriesT3ma200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**T3ma** | Pointer to **string** | T3MA value | [optional] 

## Methods

### NewGetTimeSeriesT3ma200ResponseValuesInner

`func NewGetTimeSeriesT3ma200ResponseValuesInner() *GetTimeSeriesT3ma200ResponseValuesInner`

NewGetTimeSeriesT3ma200ResponseValuesInner instantiates a new GetTimeSeriesT3ma200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesT3ma200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesT3ma200ResponseValuesInnerWithDefaults() *GetTimeSeriesT3ma200ResponseValuesInner`

NewGetTimeSeriesT3ma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesT3ma200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetT3ma

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetT3ma() string`

GetT3ma returns the T3ma field if non-nil, zero value otherwise.

### GetT3maOk

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetT3maOk() (*string, bool)`

GetT3maOk returns a tuple with the T3ma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT3ma

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) SetT3ma(v string)`

SetT3ma sets T3ma field to given value.

### HasT3ma

`func (o *GetTimeSeriesT3ma200ResponseValuesInner) HasT3ma() bool`

HasT3ma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


