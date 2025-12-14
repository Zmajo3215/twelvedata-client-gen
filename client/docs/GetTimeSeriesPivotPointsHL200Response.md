# GetTimeSeriesPivotPointsHL200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesPivotPointsHL200ResponseMeta**](GetTimeSeriesPivotPointsHL200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesPivotPointsHL200ResponseValuesInner**](GetTimeSeriesPivotPointsHL200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesPivotPointsHL200Response

`func NewGetTimeSeriesPivotPointsHL200Response() *GetTimeSeriesPivotPointsHL200Response`

NewGetTimeSeriesPivotPointsHL200Response instantiates a new GetTimeSeriesPivotPointsHL200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults

`func NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults() *GetTimeSeriesPivotPointsHL200Response`

NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults instantiates a new GetTimeSeriesPivotPointsHL200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPivotPointsHL200Response) GetMeta() GetTimeSeriesPivotPointsHL200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetMetaOk() (*GetTimeSeriesPivotPointsHL200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPivotPointsHL200Response) SetMeta(v GetTimeSeriesPivotPointsHL200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesPivotPointsHL200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesPivotPointsHL200Response) GetValues() []GetTimeSeriesPivotPointsHL200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetValuesOk() (*[]GetTimeSeriesPivotPointsHL200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPivotPointsHL200Response) SetValues(v []GetTimeSeriesPivotPointsHL200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesPivotPointsHL200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesPivotPointsHL200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPivotPointsHL200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesPivotPointsHL200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


