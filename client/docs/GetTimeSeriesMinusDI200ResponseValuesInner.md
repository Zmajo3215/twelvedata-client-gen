# GetTimeSeriesMinusDI200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**MinusDi** | Pointer to **string** | Minus_di value | [optional] 

## Methods

### NewGetTimeSeriesMinusDI200ResponseValuesInner

`func NewGetTimeSeriesMinusDI200ResponseValuesInner() *GetTimeSeriesMinusDI200ResponseValuesInner`

NewGetTimeSeriesMinusDI200ResponseValuesInner instantiates a new GetTimeSeriesMinusDI200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinusDI200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMinusDI200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinusDI200ResponseValuesInner`

NewGetTimeSeriesMinusDI200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinusDI200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetMinusDi

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetMinusDi() string`

GetMinusDi returns the MinusDi field if non-nil, zero value otherwise.

### GetMinusDiOk

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetMinusDiOk() (*string, bool)`

GetMinusDiOk returns a tuple with the MinusDi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinusDi

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) SetMinusDi(v string)`

SetMinusDi sets MinusDi field to given value.

### HasMinusDi

`func (o *GetTimeSeriesMinusDI200ResponseValuesInner) HasMinusDi() bool`

HasMinusDi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


