# GetTimeSeriesPlusDM200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesPlusDM200ResponseMeta**](GetTimeSeriesPlusDM200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesPlusDM200ResponseValuesInner**](GetTimeSeriesPlusDM200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesPlusDM200Response

`func NewGetTimeSeriesPlusDM200Response() *GetTimeSeriesPlusDM200Response`

NewGetTimeSeriesPlusDM200Response instantiates a new GetTimeSeriesPlusDM200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPlusDM200ResponseWithDefaults

`func NewGetTimeSeriesPlusDM200ResponseWithDefaults() *GetTimeSeriesPlusDM200Response`

NewGetTimeSeriesPlusDM200ResponseWithDefaults instantiates a new GetTimeSeriesPlusDM200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPlusDM200Response) GetMeta() GetTimeSeriesPlusDM200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPlusDM200Response) GetMetaOk() (*GetTimeSeriesPlusDM200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPlusDM200Response) SetMeta(v GetTimeSeriesPlusDM200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesPlusDM200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesPlusDM200Response) GetValues() []GetTimeSeriesPlusDM200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPlusDM200Response) GetValuesOk() (*[]GetTimeSeriesPlusDM200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPlusDM200Response) SetValues(v []GetTimeSeriesPlusDM200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesPlusDM200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesPlusDM200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPlusDM200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPlusDM200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesPlusDM200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


