# GetInstitutionalHolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetFundHolders200ResponseMeta**](GetFundHolders200ResponseMeta.md) |  | [optional] 
**InstitutionalHolders** | Pointer to [**[]InstitutionalHolderItem**](InstitutionalHolderItem.md) | List of institutional holders for the financial instrument | [optional] 

## Methods

### NewGetInstitutionalHolders200Response

`func NewGetInstitutionalHolders200Response() *GetInstitutionalHolders200Response`

NewGetInstitutionalHolders200Response instantiates a new GetInstitutionalHolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstitutionalHolders200ResponseWithDefaults

`func NewGetInstitutionalHolders200ResponseWithDefaults() *GetInstitutionalHolders200Response`

NewGetInstitutionalHolders200ResponseWithDefaults instantiates a new GetInstitutionalHolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetInstitutionalHolders200Response) GetMeta() GetFundHolders200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInstitutionalHolders200Response) GetMetaOk() (*GetFundHolders200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInstitutionalHolders200Response) SetMeta(v GetFundHolders200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInstitutionalHolders200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetInstitutionalHolders

`func (o *GetInstitutionalHolders200Response) GetInstitutionalHolders() []InstitutionalHolderItem`

GetInstitutionalHolders returns the InstitutionalHolders field if non-nil, zero value otherwise.

### GetInstitutionalHoldersOk

`func (o *GetInstitutionalHolders200Response) GetInstitutionalHoldersOk() (*[]InstitutionalHolderItem, bool)`

GetInstitutionalHoldersOk returns a tuple with the InstitutionalHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalHolders

`func (o *GetInstitutionalHolders200Response) SetInstitutionalHolders(v []InstitutionalHolderItem)`

SetInstitutionalHolders sets InstitutionalHolders field to given value.

### HasInstitutionalHolders

`func (o *GetInstitutionalHolders200Response) HasInstitutionalHolders() bool`

HasInstitutionalHolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


