# GetTimeSeriesPercentB200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesPercentB200ResponseMeta**](GetTimeSeriesPercentB200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesPercentB200ResponseValuesInner**](GetTimeSeriesPercentB200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesPercentB200Response

`func NewGetTimeSeriesPercentB200Response() *GetTimeSeriesPercentB200Response`

NewGetTimeSeriesPercentB200Response instantiates a new GetTimeSeriesPercentB200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPercentB200ResponseWithDefaults

`func NewGetTimeSeriesPercentB200ResponseWithDefaults() *GetTimeSeriesPercentB200Response`

NewGetTimeSeriesPercentB200ResponseWithDefaults instantiates a new GetTimeSeriesPercentB200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPercentB200Response) GetMeta() GetTimeSeriesPercentB200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPercentB200Response) GetMetaOk() (*GetTimeSeriesPercentB200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPercentB200Response) SetMeta(v GetTimeSeriesPercentB200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesPercentB200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesPercentB200Response) GetValues() []GetTimeSeriesPercentB200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPercentB200Response) GetValuesOk() (*[]GetTimeSeriesPercentB200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPercentB200Response) SetValues(v []GetTimeSeriesPercentB200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesPercentB200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesPercentB200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPercentB200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPercentB200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesPercentB200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


