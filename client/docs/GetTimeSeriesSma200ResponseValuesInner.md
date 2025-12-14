# GetTimeSeriesSma200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Sma** | Pointer to **string** | SMA value | [optional] 

## Methods

### NewGetTimeSeriesSma200ResponseValuesInner

`func NewGetTimeSeriesSma200ResponseValuesInner() *GetTimeSeriesSma200ResponseValuesInner`

NewGetTimeSeriesSma200ResponseValuesInner instantiates a new GetTimeSeriesSma200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSma200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSma200ResponseValuesInnerWithDefaults() *GetTimeSeriesSma200ResponseValuesInner`

NewGetTimeSeriesSma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSma200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSma200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSma200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSma200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesSma200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetSma

`func (o *GetTimeSeriesSma200ResponseValuesInner) GetSma() string`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *GetTimeSeriesSma200ResponseValuesInner) GetSmaOk() (*string, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *GetTimeSeriesSma200ResponseValuesInner) SetSma(v string)`

SetSma sets Sma field to given value.

### HasSma

`func (o *GetTimeSeriesSma200ResponseValuesInner) HasSma() bool`

HasSma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


