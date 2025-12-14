# GetTimeSeriesMacd200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Macd** | Pointer to **string** | MACD value | [optional] 
**MacdSignal** | Pointer to **string** | MACD signal line value | [optional] 
**MacdHist** | Pointer to **string** | MACD histogram value | [optional] 

## Methods

### NewGetTimeSeriesMacd200ResponseValuesInner

`func NewGetTimeSeriesMacd200ResponseValuesInner() *GetTimeSeriesMacd200ResponseValuesInner`

NewGetTimeSeriesMacd200ResponseValuesInner instantiates a new GetTimeSeriesMacd200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacd200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMacd200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacd200ResponseValuesInner`

NewGetTimeSeriesMacd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacd200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMacd200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesMacd200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetMacd

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacd() string`

GetMacd returns the Macd field if non-nil, zero value otherwise.

### GetMacdOk

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdOk() (*string, bool)`

GetMacdOk returns a tuple with the Macd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacd

`func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacd(v string)`

SetMacd sets Macd field to given value.

### HasMacd

`func (o *GetTimeSeriesMacd200ResponseValuesInner) HasMacd() bool`

HasMacd returns a boolean if a field has been set.

### GetMacdSignal

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdSignal() string`

GetMacdSignal returns the MacdSignal field if non-nil, zero value otherwise.

### GetMacdSignalOk

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdSignalOk() (*string, bool)`

GetMacdSignalOk returns a tuple with the MacdSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdSignal

`func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacdSignal(v string)`

SetMacdSignal sets MacdSignal field to given value.

### HasMacdSignal

`func (o *GetTimeSeriesMacd200ResponseValuesInner) HasMacdSignal() bool`

HasMacdSignal returns a boolean if a field has been set.

### GetMacdHist

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdHist() string`

GetMacdHist returns the MacdHist field if non-nil, zero value otherwise.

### GetMacdHistOk

`func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdHistOk() (*string, bool)`

GetMacdHistOk returns a tuple with the MacdHist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdHist

`func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacdHist(v string)`

SetMacdHist sets MacdHist field to given value.

### HasMacdHist

`func (o *GetTimeSeriesMacd200ResponseValuesInner) HasMacdHist() bool`

HasMacdHist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


