# GetMutualFundsType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **map[string][]string** | List of fund types by country | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetMutualFundsType200Response

`func NewGetMutualFundsType200Response() *GetMutualFundsType200Response`

NewGetMutualFundsType200Response instantiates a new GetMutualFundsType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMutualFundsType200ResponseWithDefaults

`func NewGetMutualFundsType200ResponseWithDefaults() *GetMutualFundsType200Response`

NewGetMutualFundsType200ResponseWithDefaults instantiates a new GetMutualFundsType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetMutualFundsType200Response) GetResult() map[string][]string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMutualFundsType200Response) GetResultOk() (*map[string][]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMutualFundsType200Response) SetResult(v map[string][]string)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMutualFundsType200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *GetMutualFundsType200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMutualFundsType200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMutualFundsType200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMutualFundsType200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


