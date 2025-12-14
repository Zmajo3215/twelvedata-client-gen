# GetTimeSeriesSuperTrend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesSuperTrend200ResponseMeta**](GetTimeSeriesSuperTrend200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesSuperTrend200ResponseValuesInner**](GetTimeSeriesSuperTrend200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesSuperTrend200Response

`func NewGetTimeSeriesSuperTrend200Response() *GetTimeSeriesSuperTrend200Response`

NewGetTimeSeriesSuperTrend200Response instantiates a new GetTimeSeriesSuperTrend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSuperTrend200ResponseWithDefaults

`func NewGetTimeSeriesSuperTrend200ResponseWithDefaults() *GetTimeSeriesSuperTrend200Response`

NewGetTimeSeriesSuperTrend200ResponseWithDefaults instantiates a new GetTimeSeriesSuperTrend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSuperTrend200Response) GetMeta() GetTimeSeriesSuperTrend200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSuperTrend200Response) GetMetaOk() (*GetTimeSeriesSuperTrend200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSuperTrend200Response) SetMeta(v GetTimeSeriesSuperTrend200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesSuperTrend200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesSuperTrend200Response) GetValues() []GetTimeSeriesSuperTrend200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSuperTrend200Response) GetValuesOk() (*[]GetTimeSeriesSuperTrend200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSuperTrend200Response) SetValues(v []GetTimeSeriesSuperTrend200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesSuperTrend200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesSuperTrend200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSuperTrend200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSuperTrend200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesSuperTrend200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


