# GetTimeSeriesMaxIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesMaxIndex200ResponseMeta**](GetTimeSeriesMaxIndex200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesMaxIndex200ResponseValuesInner**](GetTimeSeriesMaxIndex200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesMaxIndex200Response

`func NewGetTimeSeriesMaxIndex200Response() *GetTimeSeriesMaxIndex200Response`

NewGetTimeSeriesMaxIndex200Response instantiates a new GetTimeSeriesMaxIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMaxIndex200ResponseWithDefaults

`func NewGetTimeSeriesMaxIndex200ResponseWithDefaults() *GetTimeSeriesMaxIndex200Response`

NewGetTimeSeriesMaxIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMaxIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMaxIndex200Response) GetMeta() GetTimeSeriesMaxIndex200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMaxIndex200Response) GetMetaOk() (*GetTimeSeriesMaxIndex200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMaxIndex200Response) SetMeta(v GetTimeSeriesMaxIndex200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesMaxIndex200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesMaxIndex200Response) GetValues() []GetTimeSeriesMaxIndex200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMaxIndex200Response) GetValuesOk() (*[]GetTimeSeriesMaxIndex200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMaxIndex200Response) SetValues(v []GetTimeSeriesMaxIndex200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesMaxIndex200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesMaxIndex200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMaxIndex200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMaxIndex200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesMaxIndex200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


