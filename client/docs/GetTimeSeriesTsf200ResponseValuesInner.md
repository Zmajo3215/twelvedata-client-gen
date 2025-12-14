# GetTimeSeriesTsf200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Tsf** | Pointer to **string** | TSF value | [optional] 

## Methods

### NewGetTimeSeriesTsf200ResponseValuesInner

`func NewGetTimeSeriesTsf200ResponseValuesInner() *GetTimeSeriesTsf200ResponseValuesInner`

NewGetTimeSeriesTsf200ResponseValuesInner instantiates a new GetTimeSeriesTsf200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTsf200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesTsf200ResponseValuesInnerWithDefaults() *GetTimeSeriesTsf200ResponseValuesInner`

NewGetTimeSeriesTsf200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTsf200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesTsf200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesTsf200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesTsf200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesTsf200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetTsf

`func (o *GetTimeSeriesTsf200ResponseValuesInner) GetTsf() string`

GetTsf returns the Tsf field if non-nil, zero value otherwise.

### GetTsfOk

`func (o *GetTimeSeriesTsf200ResponseValuesInner) GetTsfOk() (*string, bool)`

GetTsfOk returns a tuple with the Tsf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsf

`func (o *GetTimeSeriesTsf200ResponseValuesInner) SetTsf(v string)`

SetTsf sets Tsf field to given value.

### HasTsf

`func (o *GetTimeSeriesTsf200ResponseValuesInner) HasTsf() bool`

HasTsf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


