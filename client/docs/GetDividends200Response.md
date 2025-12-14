# GetDividends200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetDividends200ResponseMeta**](GetDividends200ResponseMeta.md) |  | [optional] 
**Dividends** | Pointer to [**[]GetDividends200ResponseDividendsInner**](GetDividends200ResponseDividendsInner.md) | List of dividends | [optional] 

## Methods

### NewGetDividends200Response

`func NewGetDividends200Response() *GetDividends200Response`

NewGetDividends200Response instantiates a new GetDividends200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDividends200ResponseWithDefaults

`func NewGetDividends200ResponseWithDefaults() *GetDividends200Response`

NewGetDividends200ResponseWithDefaults instantiates a new GetDividends200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetDividends200Response) GetMeta() GetDividends200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDividends200Response) GetMetaOk() (*GetDividends200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDividends200Response) SetMeta(v GetDividends200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDividends200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDividends

`func (o *GetDividends200Response) GetDividends() []GetDividends200ResponseDividendsInner`

GetDividends returns the Dividends field if non-nil, zero value otherwise.

### GetDividendsOk

`func (o *GetDividends200Response) GetDividendsOk() (*[]GetDividends200ResponseDividendsInner, bool)`

GetDividendsOk returns a tuple with the Dividends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividends

`func (o *GetDividends200Response) SetDividends(v []GetDividends200ResponseDividendsInner)`

SetDividends sets Dividends field to given value.

### HasDividends

`func (o *GetDividends200Response) HasDividends() bool`

HasDividends returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


