# GetETFsType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **map[string][]string** | List of ETFs by market | [optional] 
**Status** | Pointer to **string** | Status of the response | [optional] 

## Methods

### NewGetETFsType200Response

`func NewGetETFsType200Response() *GetETFsType200Response`

NewGetETFsType200Response instantiates a new GetETFsType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsType200ResponseWithDefaults

`func NewGetETFsType200ResponseWithDefaults() *GetETFsType200Response`

NewGetETFsType200ResponseWithDefaults instantiates a new GetETFsType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetETFsType200Response) GetResult() map[string][]string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetETFsType200Response) GetResultOk() (*map[string][]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetETFsType200Response) SetResult(v map[string][]string)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetETFsType200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *GetETFsType200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsType200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsType200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetETFsType200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


