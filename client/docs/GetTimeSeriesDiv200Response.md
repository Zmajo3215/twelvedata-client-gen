# GetTimeSeriesDiv200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesDiv200ResponseMeta**](GetTimeSeriesDiv200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesDiv200ResponseValuesInner**](GetTimeSeriesDiv200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesDiv200Response

`func NewGetTimeSeriesDiv200Response() *GetTimeSeriesDiv200Response`

NewGetTimeSeriesDiv200Response instantiates a new GetTimeSeriesDiv200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDiv200ResponseWithDefaults

`func NewGetTimeSeriesDiv200ResponseWithDefaults() *GetTimeSeriesDiv200Response`

NewGetTimeSeriesDiv200ResponseWithDefaults instantiates a new GetTimeSeriesDiv200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesDiv200Response) GetMeta() GetTimeSeriesDiv200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesDiv200Response) GetMetaOk() (*GetTimeSeriesDiv200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesDiv200Response) SetMeta(v GetTimeSeriesDiv200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesDiv200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesDiv200Response) GetValues() []GetTimeSeriesDiv200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesDiv200Response) GetValuesOk() (*[]GetTimeSeriesDiv200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesDiv200Response) SetValues(v []GetTimeSeriesDiv200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesDiv200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesDiv200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesDiv200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesDiv200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesDiv200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


