# GetTimeSeriesBeta200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesBeta200ResponseMeta**](GetTimeSeriesBeta200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesBeta200ResponseValuesInner**](GetTimeSeriesBeta200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesBeta200Response

`func NewGetTimeSeriesBeta200Response() *GetTimeSeriesBeta200Response`

NewGetTimeSeriesBeta200Response instantiates a new GetTimeSeriesBeta200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBeta200ResponseWithDefaults

`func NewGetTimeSeriesBeta200ResponseWithDefaults() *GetTimeSeriesBeta200Response`

NewGetTimeSeriesBeta200ResponseWithDefaults instantiates a new GetTimeSeriesBeta200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesBeta200Response) GetMeta() GetTimeSeriesBeta200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesBeta200Response) GetMetaOk() (*GetTimeSeriesBeta200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesBeta200Response) SetMeta(v GetTimeSeriesBeta200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesBeta200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesBeta200Response) GetValues() []GetTimeSeriesBeta200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesBeta200Response) GetValuesOk() (*[]GetTimeSeriesBeta200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesBeta200Response) SetValues(v []GetTimeSeriesBeta200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesBeta200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesBeta200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesBeta200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesBeta200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesBeta200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


