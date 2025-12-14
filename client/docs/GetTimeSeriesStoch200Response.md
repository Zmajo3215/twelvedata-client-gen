# GetTimeSeriesStoch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesStoch200ResponseMeta**](GetTimeSeriesStoch200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesStoch200ResponseValuesInner**](GetTimeSeriesStoch200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesStoch200Response

`func NewGetTimeSeriesStoch200Response() *GetTimeSeriesStoch200Response`

NewGetTimeSeriesStoch200Response instantiates a new GetTimeSeriesStoch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStoch200ResponseWithDefaults

`func NewGetTimeSeriesStoch200ResponseWithDefaults() *GetTimeSeriesStoch200Response`

NewGetTimeSeriesStoch200ResponseWithDefaults instantiates a new GetTimeSeriesStoch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesStoch200Response) GetMeta() GetTimeSeriesStoch200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesStoch200Response) GetMetaOk() (*GetTimeSeriesStoch200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesStoch200Response) SetMeta(v GetTimeSeriesStoch200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesStoch200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesStoch200Response) GetValues() []GetTimeSeriesStoch200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesStoch200Response) GetValuesOk() (*[]GetTimeSeriesStoch200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesStoch200Response) SetValues(v []GetTimeSeriesStoch200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesStoch200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesStoch200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesStoch200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesStoch200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesStoch200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


