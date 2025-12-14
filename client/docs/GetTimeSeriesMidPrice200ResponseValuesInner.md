# GetTimeSeriesMidPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Midprice** | Pointer to **string** | Midprice value | [optional] 

## Methods

### NewGetTimeSeriesMidPrice200ResponseValuesInner

`func NewGetTimeSeriesMidPrice200ResponseValuesInner() *GetTimeSeriesMidPrice200ResponseValuesInner`

NewGetTimeSeriesMidPrice200ResponseValuesInner instantiates a new GetTimeSeriesMidPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMidPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMidPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesMidPrice200ResponseValuesInner`

NewGetTimeSeriesMidPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMidPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetMidprice

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetMidprice() string`

GetMidprice returns the Midprice field if non-nil, zero value otherwise.

### GetMidpriceOk

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetMidpriceOk() (*string, bool)`

GetMidpriceOk returns a tuple with the Midprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidprice

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) SetMidprice(v string)`

SetMidprice sets Midprice field to given value.

### HasMidprice

`func (o *GetTimeSeriesMidPrice200ResponseValuesInner) HasMidprice() bool`

HasMidprice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


