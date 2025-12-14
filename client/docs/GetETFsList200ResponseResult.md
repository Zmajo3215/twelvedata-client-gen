# GetETFsList200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Total number of matching funds | [optional] 
**List** | Pointer to [**[]ETFsListResponseItem**](ETFsListResponseItem.md) | List of ETFs | [optional] 

## Methods

### NewGetETFsList200ResponseResult

`func NewGetETFsList200ResponseResult() *GetETFsList200ResponseResult`

NewGetETFsList200ResponseResult instantiates a new GetETFsList200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsList200ResponseResultWithDefaults

`func NewGetETFsList200ResponseResultWithDefaults() *GetETFsList200ResponseResult`

NewGetETFsList200ResponseResultWithDefaults instantiates a new GetETFsList200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetETFsList200ResponseResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetETFsList200ResponseResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetETFsList200ResponseResult) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetETFsList200ResponseResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetList

`func (o *GetETFsList200ResponseResult) GetList() []ETFsListResponseItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetETFsList200ResponseResult) GetListOk() (*[]ETFsListResponseItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetETFsList200ResponseResult) SetList(v []ETFsListResponseItem)`

SetList sets List field to given value.

### HasList

`func (o *GetETFsList200ResponseResult) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


