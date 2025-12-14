# GetTimeSeriesHtTrendline200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**HtTrendline** | Pointer to **string** | HT_TRENDLINE value | [optional] 

## Methods

### NewGetTimeSeriesHtTrendline200ResponseValuesInner

`func NewGetTimeSeriesHtTrendline200ResponseValuesInner() *GetTimeSeriesHtTrendline200ResponseValuesInner`

NewGetTimeSeriesHtTrendline200ResponseValuesInner instantiates a new GetTimeSeriesHtTrendline200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtTrendline200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtTrendline200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtTrendline200ResponseValuesInner`

NewGetTimeSeriesHtTrendline200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtTrendline200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetHtTrendline

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetHtTrendline() string`

GetHtTrendline returns the HtTrendline field if non-nil, zero value otherwise.

### GetHtTrendlineOk

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetHtTrendlineOk() (*string, bool)`

GetHtTrendlineOk returns a tuple with the HtTrendline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtTrendline

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) SetHtTrendline(v string)`

SetHtTrendline sets HtTrendline field to given value.

### HasHtTrendline

`func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) HasHtTrendline() bool`

HasHtTrendline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


