# GetTimeSeriesMacdSlope200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetTimeSeriesMacdSlope200ResponseMeta**](GetTimeSeriesMacdSlope200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]GetTimeSeriesMacdSlope200ResponseValuesInner**](GetTimeSeriesMacdSlope200ResponseValuesInner.md) | Array of time series data points | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetTimeSeriesMacdSlope200Response

`func NewGetTimeSeriesMacdSlope200Response() *GetTimeSeriesMacdSlope200Response`

NewGetTimeSeriesMacdSlope200Response instantiates a new GetTimeSeriesMacdSlope200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacdSlope200ResponseWithDefaults

`func NewGetTimeSeriesMacdSlope200ResponseWithDefaults() *GetTimeSeriesMacdSlope200Response`

NewGetTimeSeriesMacdSlope200ResponseWithDefaults instantiates a new GetTimeSeriesMacdSlope200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMacdSlope200Response) GetMeta() GetTimeSeriesMacdSlope200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMacdSlope200Response) GetMetaOk() (*GetTimeSeriesMacdSlope200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMacdSlope200Response) SetMeta(v GetTimeSeriesMacdSlope200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTimeSeriesMacdSlope200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetTimeSeriesMacdSlope200Response) GetValues() []GetTimeSeriesMacdSlope200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMacdSlope200Response) GetValuesOk() (*[]GetTimeSeriesMacdSlope200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMacdSlope200Response) SetValues(v []GetTimeSeriesMacdSlope200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetTimeSeriesMacdSlope200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *GetTimeSeriesMacdSlope200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMacdSlope200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMacdSlope200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTimeSeriesMacdSlope200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


