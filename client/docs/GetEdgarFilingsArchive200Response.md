# GetEdgarFilingsArchive200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetEdgarFilingsArchive200ResponseMeta**](GetEdgarFilingsArchive200ResponseMeta.md) |  | [optional] 
**Values** | Pointer to [**[]EdgarFilingValue**](EdgarFilingValue.md) | List of filings | [optional] 

## Methods

### NewGetEdgarFilingsArchive200Response

`func NewGetEdgarFilingsArchive200Response() *GetEdgarFilingsArchive200Response`

NewGetEdgarFilingsArchive200Response instantiates a new GetEdgarFilingsArchive200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEdgarFilingsArchive200ResponseWithDefaults

`func NewGetEdgarFilingsArchive200ResponseWithDefaults() *GetEdgarFilingsArchive200Response`

NewGetEdgarFilingsArchive200ResponseWithDefaults instantiates a new GetEdgarFilingsArchive200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetEdgarFilingsArchive200Response) GetMeta() GetEdgarFilingsArchive200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetEdgarFilingsArchive200Response) GetMetaOk() (*GetEdgarFilingsArchive200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetEdgarFilingsArchive200Response) SetMeta(v GetEdgarFilingsArchive200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetEdgarFilingsArchive200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValues

`func (o *GetEdgarFilingsArchive200Response) GetValues() []EdgarFilingValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetEdgarFilingsArchive200Response) GetValuesOk() (*[]EdgarFilingValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetEdgarFilingsArchive200Response) SetValues(v []EdgarFilingValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetEdgarFilingsArchive200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


