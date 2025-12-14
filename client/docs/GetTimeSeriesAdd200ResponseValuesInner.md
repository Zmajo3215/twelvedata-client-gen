# GetTimeSeriesAdd200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Add** | Pointer to **string** | Add value | [optional] 

## Methods

### NewGetTimeSeriesAdd200ResponseValuesInner

`func NewGetTimeSeriesAdd200ResponseValuesInner() *GetTimeSeriesAdd200ResponseValuesInner`

NewGetTimeSeriesAdd200ResponseValuesInner instantiates a new GetTimeSeriesAdd200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdd200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAdd200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdd200ResponseValuesInner`

NewGetTimeSeriesAdd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdd200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAdd200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAdd200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAdd200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesAdd200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetAdd

`func (o *GetTimeSeriesAdd200ResponseValuesInner) GetAdd() string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *GetTimeSeriesAdd200ResponseValuesInner) GetAddOk() (*string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *GetTimeSeriesAdd200ResponseValuesInner) SetAdd(v string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *GetTimeSeriesAdd200ResponseValuesInner) HasAdd() bool`

HasAdd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


