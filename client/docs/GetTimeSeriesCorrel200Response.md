# GetTimeSeriesCorrel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesCorrel200ResponseMeta**](GetTimeSeriesCorrel200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesCorrel200ResponseValuesInner**](GetTimeSeriesCorrel200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesCorrel200Response

`func NewGetTimeSeriesCorrel200Response() *GetTimeSeriesCorrel200Response`

NewGetTimeSeriesCorrel200Response instantiates a new GetTimeSeriesCorrel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCorrel200ResponseWithDefaults

`func NewGetTimeSeriesCorrel200ResponseWithDefaults() *GetTimeSeriesCorrel200Response`

NewGetTimeSeriesCorrel200ResponseWithDefaults instantiates a new GetTimeSeriesCorrel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCorrel200Response) GetMeta() GetTimeSeriesCorrel200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCorrel200Response) GetMetaOk() (*GetTimeSeriesCorrel200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCorrel200Response) SetMeta(v GetTimeSeriesCorrel200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesCorrel200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesCorrel200Response) GetValues() []GetTimeSeriesCorrel200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCorrel200Response) GetValuesOk() (*[]GetTimeSeriesCorrel200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCorrel200Response) SetValues(v []GetTimeSeriesCorrel200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesCorrel200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesCorrel200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCorrel200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCorrel200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesCorrel200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


