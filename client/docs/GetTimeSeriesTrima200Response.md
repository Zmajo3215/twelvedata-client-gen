# GetTimeSeriesTrima200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesTrima200ResponseMeta**](GetTimeSeriesTrima200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesTrima200ResponseValuesInner**](GetTimeSeriesTrima200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesTrima200Response

`func NewGetTimeSeriesTrima200Response() *GetTimeSeriesTrima200Response`

NewGetTimeSeriesTrima200Response instantiates a new GetTimeSeriesTrima200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTrima200ResponseWithDefaults

`func NewGetTimeSeriesTrima200ResponseWithDefaults() *GetTimeSeriesTrima200Response`

NewGetTimeSeriesTrima200ResponseWithDefaults instantiates a new GetTimeSeriesTrima200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesTrima200Response) GetMeta() GetTimeSeriesTrima200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesTrima200Response) GetMetaOk() (*GetTimeSeriesTrima200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesTrima200Response) SetMeta(v GetTimeSeriesTrima200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesTrima200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesTrima200Response) GetValues() []GetTimeSeriesTrima200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesTrima200Response) GetValuesOk() (*[]GetTimeSeriesTrima200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesTrima200Response) SetValues(v []GetTimeSeriesTrima200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesTrima200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesTrima200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesTrima200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesTrima200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesTrima200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


