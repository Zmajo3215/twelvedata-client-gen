# CrossListingsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Number of cross listings found | [optional] 
**List** | Pointer to [**[]CrossListingsItem**](CrossListingsItem.md) | List of cross listings | [optional] 

## Methods

### NewCrossListingsResult

`func NewCrossListingsResult() *CrossListingsResult`

NewCrossListingsResult instantiates a new CrossListingsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossListingsResultWithDefaults

`func NewCrossListingsResultWithDefaults() *CrossListingsResult`

NewCrossListingsResultWithDefaults instantiates a new CrossListingsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CrossListingsResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CrossListingsResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CrossListingsResult) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CrossListingsResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetList

`func (o *CrossListingsResult) GetList() []CrossListingsItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *CrossListingsResult) GetListOk() (*[]CrossListingsItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *CrossListingsResult) SetList(v []CrossListingsItem)`

SetList sets List field to given value.

### HasList

`func (o *CrossListingsResult) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


