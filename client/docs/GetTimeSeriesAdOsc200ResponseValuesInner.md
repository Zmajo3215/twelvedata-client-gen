# GetTimeSeriesAdOsc200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Adosc** | Pointer to **string** | Adosc value | [optional] 

## Methods

### NewGetTimeSeriesAdOsc200ResponseValuesInner

`func NewGetTimeSeriesAdOsc200ResponseValuesInner() *GetTimeSeriesAdOsc200ResponseValuesInner`

NewGetTimeSeriesAdOsc200ResponseValuesInner instantiates a new GetTimeSeriesAdOsc200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdOsc200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAdOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdOsc200ResponseValuesInner`

NewGetTimeSeriesAdOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdOsc200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetAdosc

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetAdosc() string`

GetAdosc returns the Adosc field if non-nil, zero value otherwise.

### GetAdoscOk

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetAdoscOk() (*string, bool)`

GetAdoscOk returns a tuple with the Adosc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdosc

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) SetAdosc(v string)`

SetAdosc sets Adosc field to given value.

### HasAdosc

`func (o *GetTimeSeriesAdOsc200ResponseValuesInner) HasAdosc() bool`

HasAdosc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


