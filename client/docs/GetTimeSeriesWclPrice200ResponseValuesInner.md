# GetTimeSeriesWclPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Wclprice** | Pointer to **float64** | wclprice value | [optional] 

## Methods

### NewGetTimeSeriesWclPrice200ResponseValuesInner

`func NewGetTimeSeriesWclPrice200ResponseValuesInner() *GetTimeSeriesWclPrice200ResponseValuesInner`

NewGetTimeSeriesWclPrice200ResponseValuesInner instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesWclPrice200ResponseValuesInner`

NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetWclprice

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclprice() float64`

GetWclprice returns the Wclprice field if non-nil, zero value otherwise.

### GetWclpriceOk

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclpriceOk() (*float64, bool)`

GetWclpriceOk returns a tuple with the Wclprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWclprice

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetWclprice(v float64)`

SetWclprice sets Wclprice field to given value.

### HasWclprice

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) HasWclprice() bool`

HasWclprice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


