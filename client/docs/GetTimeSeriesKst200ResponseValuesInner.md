# GetTimeSeriesKst200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Kst** | Pointer to **string** | KST value | [optional] 
**KstSignal** | Pointer to **string** | KST signal value | [optional] 

## Methods

### NewGetTimeSeriesKst200ResponseValuesInner

`func NewGetTimeSeriesKst200ResponseValuesInner() *GetTimeSeriesKst200ResponseValuesInner`

NewGetTimeSeriesKst200ResponseValuesInner instantiates a new GetTimeSeriesKst200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKst200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesKst200ResponseValuesInnerWithDefaults() *GetTimeSeriesKst200ResponseValuesInner`

NewGetTimeSeriesKst200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKst200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesKst200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesKst200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetKst

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetKst() string`

GetKst returns the Kst field if non-nil, zero value otherwise.

### GetKstOk

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstOk() (*string, bool)`

GetKstOk returns a tuple with the Kst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKst

`func (o *GetTimeSeriesKst200ResponseValuesInner) SetKst(v string)`

SetKst sets Kst field to given value.

### HasKst

`func (o *GetTimeSeriesKst200ResponseValuesInner) HasKst() bool`

HasKst returns a boolean if a field has been set.

### GetKstSignal

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstSignal() string`

GetKstSignal returns the KstSignal field if non-nil, zero value otherwise.

### GetKstSignalOk

`func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstSignalOk() (*string, bool)`

GetKstSignalOk returns a tuple with the KstSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKstSignal

`func (o *GetTimeSeriesKst200ResponseValuesInner) SetKstSignal(v string)`

SetKstSignal sets KstSignal field to given value.

### HasKstSignal

`func (o *GetTimeSeriesKst200ResponseValuesInner) HasKstSignal() bool`

HasKstSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


