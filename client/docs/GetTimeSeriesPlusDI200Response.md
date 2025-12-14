# GetTimeSeriesPlusDI200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesPlusDI200ResponseMeta**](GetTimeSeriesPlusDI200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesPlusDI200ResponseValuesInner**](GetTimeSeriesPlusDI200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesPlusDI200Response

`func NewGetTimeSeriesPlusDI200Response() *GetTimeSeriesPlusDI200Response`

NewGetTimeSeriesPlusDI200Response instantiates a new GetTimeSeriesPlusDI200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPlusDI200ResponseWithDefaults

`func NewGetTimeSeriesPlusDI200ResponseWithDefaults() *GetTimeSeriesPlusDI200Response`

NewGetTimeSeriesPlusDI200ResponseWithDefaults instantiates a new GetTimeSeriesPlusDI200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPlusDI200Response) GetMeta() GetTimeSeriesPlusDI200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPlusDI200Response) GetMetaOk() (*GetTimeSeriesPlusDI200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPlusDI200Response) SetMeta(v GetTimeSeriesPlusDI200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesPlusDI200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesPlusDI200Response) GetValues() []GetTimeSeriesPlusDI200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPlusDI200Response) GetValuesOk() (*[]GetTimeSeriesPlusDI200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPlusDI200Response) SetValues(v []GetTimeSeriesPlusDI200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesPlusDI200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesPlusDI200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPlusDI200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPlusDI200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesPlusDI200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


