# GetTimeSeriesRvol200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Rvol** | Pointer to **string** | RVOL value | [optional] 

## Methods

### NewGetTimeSeriesRvol200ResponseValuesInner

`func NewGetTimeSeriesRvol200ResponseValuesInner() *GetTimeSeriesRvol200ResponseValuesInner`

NewGetTimeSeriesRvol200ResponseValuesInner instantiates a new GetTimeSeriesRvol200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRvol200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRvol200ResponseValuesInnerWithDefaults() *GetTimeSeriesRvol200ResponseValuesInner`

NewGetTimeSeriesRvol200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRvol200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRvol200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRvol200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRvol200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesRvol200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetRvol

`func (o *GetTimeSeriesRvol200ResponseValuesInner) GetRvol() string`

GetRvol returns the Rvol field if non-nil, zero value otherwise.

### GetRvolOk

`func (o *GetTimeSeriesRvol200ResponseValuesInner) GetRvolOk() (*string, bool)`

GetRvolOk returns a tuple with the Rvol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRvol

`func (o *GetTimeSeriesRvol200ResponseValuesInner) SetRvol(v string)`

SetRvol sets Rvol field to given value.

### HasRvol

`func (o *GetTimeSeriesRvol200ResponseValuesInner) HasRvol() bool`

HasRvol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


