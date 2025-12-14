# GetSourceSanctionedEntities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sanctions** | Pointer to [**[]ResponseSanctionedEntitiy**](ResponseSanctionedEntitiy.md) | List of sanctioned entities | [optional] 
**Count** | Pointer to **int64** | Total number of sanctioned entities | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetSourceSanctionedEntities200Response

`func NewGetSourceSanctionedEntities200Response() *GetSourceSanctionedEntities200Response`

NewGetSourceSanctionedEntities200Response instantiates a new GetSourceSanctionedEntities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceSanctionedEntities200ResponseWithDefaults

`func NewGetSourceSanctionedEntities200ResponseWithDefaults() *GetSourceSanctionedEntities200Response`

NewGetSourceSanctionedEntities200ResponseWithDefaults instantiates a new GetSourceSanctionedEntities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSanctions

`func (o *GetSourceSanctionedEntities200Response) GetSanctions() []ResponseSanctionedEntitiy`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *GetSourceSanctionedEntities200Response) GetSanctionsOk() (*[]ResponseSanctionedEntitiy, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *GetSourceSanctionedEntities200Response) SetSanctions(v []ResponseSanctionedEntitiy)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *GetSourceSanctionedEntities200Response) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.

### GetCount

`func (o *GetSourceSanctionedEntities200Response) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetSourceSanctionedEntities200Response) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetSourceSanctionedEntities200Response) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetSourceSanctionedEntities200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStatus

`func (o *GetSourceSanctionedEntities200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSourceSanctionedEntities200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSourceSanctionedEntities200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSourceSanctionedEntities200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


