# GetTimeSeriesNatr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesNatr200ResponseMeta**](GetTimeSeriesNatr200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesNatr200ResponseValuesInner**](GetTimeSeriesNatr200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesNatr200Response

`func NewGetTimeSeriesNatr200Response() *GetTimeSeriesNatr200Response`

NewGetTimeSeriesNatr200Response instantiates a new GetTimeSeriesNatr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesNatr200ResponseWithDefaults

`func NewGetTimeSeriesNatr200ResponseWithDefaults() *GetTimeSeriesNatr200Response`

NewGetTimeSeriesNatr200ResponseWithDefaults instantiates a new GetTimeSeriesNatr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesNatr200Response) GetMeta() GetTimeSeriesNatr200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesNatr200Response) GetMetaOk() (*GetTimeSeriesNatr200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesNatr200Response) SetMeta(v GetTimeSeriesNatr200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesNatr200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesNatr200Response) GetValues() []GetTimeSeriesNatr200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesNatr200Response) GetValuesOk() (*[]GetTimeSeriesNatr200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesNatr200Response) SetValues(v []GetTimeSeriesNatr200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesNatr200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesNatr200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesNatr200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesNatr200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesNatr200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


