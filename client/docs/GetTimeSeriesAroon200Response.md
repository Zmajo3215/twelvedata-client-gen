# GetTimeSeriesAroon200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesAroon200ResponseMeta**](GetTimeSeriesAroon200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesAroon200ResponseValuesInner**](GetTimeSeriesAroon200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesAroon200Response

`func NewGetTimeSeriesAroon200Response() *GetTimeSeriesAroon200Response`

NewGetTimeSeriesAroon200Response instantiates a new GetTimeSeriesAroon200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAroon200ResponseWithDefaults

`func NewGetTimeSeriesAroon200ResponseWithDefaults() *GetTimeSeriesAroon200Response`

NewGetTimeSeriesAroon200ResponseWithDefaults instantiates a new GetTimeSeriesAroon200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAroon200Response) GetMeta() GetTimeSeriesAroon200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAroon200Response) GetMetaOk() (*GetTimeSeriesAroon200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAroon200Response) SetMeta(v GetTimeSeriesAroon200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesAroon200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesAroon200Response) GetValues() []GetTimeSeriesAroon200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAroon200Response) GetValuesOk() (*[]GetTimeSeriesAroon200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAroon200Response) SetValues(v []GetTimeSeriesAroon200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesAroon200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesAroon200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAroon200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAroon200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesAroon200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


