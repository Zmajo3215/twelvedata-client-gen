# GetTimeSeriesMinMax200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesMinMax200ResponseMeta**](GetTimeSeriesMinMax200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesMinMax200ResponseValuesInner**](GetTimeSeriesMinMax200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesMinMax200Response

`func NewGetTimeSeriesMinMax200Response() *GetTimeSeriesMinMax200Response`

NewGetTimeSeriesMinMax200Response instantiates a new GetTimeSeriesMinMax200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinMax200ResponseWithDefaults

`func NewGetTimeSeriesMinMax200ResponseWithDefaults() *GetTimeSeriesMinMax200Response`

NewGetTimeSeriesMinMax200ResponseWithDefaults instantiates a new GetTimeSeriesMinMax200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMinMax200Response) GetMeta() GetTimeSeriesMinMax200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMinMax200Response) GetMetaOk() (*GetTimeSeriesMinMax200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMinMax200Response) SetMeta(v GetTimeSeriesMinMax200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesMinMax200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesMinMax200Response) GetValues() []GetTimeSeriesMinMax200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMinMax200Response) GetValuesOk() (*[]GetTimeSeriesMinMax200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMinMax200Response) SetValues(v []GetTimeSeriesMinMax200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesMinMax200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesMinMax200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMinMax200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMinMax200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesMinMax200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


