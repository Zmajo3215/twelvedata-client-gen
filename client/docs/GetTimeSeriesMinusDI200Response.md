# GetTimeSeriesMinusDI200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesMinusDI200ResponseMeta**](GetTimeSeriesMinusDI200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesMinusDI200ResponseValuesInner**](GetTimeSeriesMinusDI200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesMinusDI200Response

`func NewGetTimeSeriesMinusDI200Response() *GetTimeSeriesMinusDI200Response`

NewGetTimeSeriesMinusDI200Response instantiates a new GetTimeSeriesMinusDI200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinusDI200ResponseWithDefaults

`func NewGetTimeSeriesMinusDI200ResponseWithDefaults() *GetTimeSeriesMinusDI200Response`

NewGetTimeSeriesMinusDI200ResponseWithDefaults instantiates a new GetTimeSeriesMinusDI200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMinusDI200Response) GetMeta() GetTimeSeriesMinusDI200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMinusDI200Response) GetMetaOk() (*GetTimeSeriesMinusDI200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMinusDI200Response) SetMeta(v GetTimeSeriesMinusDI200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesMinusDI200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesMinusDI200Response) GetValues() []GetTimeSeriesMinusDI200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMinusDI200Response) GetValuesOk() (*[]GetTimeSeriesMinusDI200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMinusDI200Response) SetValues(v []GetTimeSeriesMinusDI200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesMinusDI200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesMinusDI200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMinusDI200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMinusDI200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesMinusDI200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


