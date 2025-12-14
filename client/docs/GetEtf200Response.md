# GetEtf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EtfResponseItem**](EtfResponseItem.md) | List of ETFs | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetEtf200Response

`func NewGetEtf200Response() *GetEtf200Response`

NewGetEtf200Response instantiates a new GetEtf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEtf200ResponseWithDefaults

`func NewGetEtf200ResponseWithDefaults() *GetEtf200Response`

NewGetEtf200ResponseWithDefaults instantiates a new GetEtf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEtf200Response) GetData() []EtfResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEtf200Response) GetDataOk() (*[]EtfResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEtf200Response) SetData(v []EtfResponseItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetEtf200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetEtf200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEtf200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEtf200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEtf200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


