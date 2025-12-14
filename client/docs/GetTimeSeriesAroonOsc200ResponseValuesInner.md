# GetTimeSeriesAroonOsc200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | [optional] 
**Aroonosc** | Pointer to **string** | Aroon oscillator value | [optional] 

## Methods

### NewGetTimeSeriesAroonOsc200ResponseValuesInner

`func NewGetTimeSeriesAroonOsc200ResponseValuesInner() *GetTimeSeriesAroonOsc200ResponseValuesInner`

NewGetTimeSeriesAroonOsc200ResponseValuesInner instantiates a new GetTimeSeriesAroonOsc200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAroonOsc200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAroonOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesAroonOsc200ResponseValuesInner`

NewGetTimeSeriesAroonOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAroonOsc200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetAroonosc

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetAroonosc() string`

GetAroonosc returns the Aroonosc field if non-nil, zero value otherwise.

### GetAroonoscOk

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetAroonoscOk() (*string, bool)`

GetAroonoscOk returns a tuple with the Aroonosc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAroonosc

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) SetAroonosc(v string)`

SetAroonosc sets Aroonosc field to given value.

### HasAroonosc

`func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) HasAroonosc() bool`

HasAroonosc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


