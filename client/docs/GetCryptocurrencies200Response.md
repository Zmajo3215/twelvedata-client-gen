# GetCryptocurrencies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CryptocurrencyResponseItem**](CryptocurrencyResponseItem.md) | List of cryptocurrencies | [optional] 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetCryptocurrencies200Response

`func NewGetCryptocurrencies200Response() *GetCryptocurrencies200Response`

NewGetCryptocurrencies200Response instantiates a new GetCryptocurrencies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCryptocurrencies200ResponseWithDefaults

`func NewGetCryptocurrencies200ResponseWithDefaults() *GetCryptocurrencies200Response`

NewGetCryptocurrencies200ResponseWithDefaults instantiates a new GetCryptocurrencies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCryptocurrencies200Response) GetData() []CryptocurrencyResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCryptocurrencies200Response) GetDataOk() (*[]CryptocurrencyResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCryptocurrencies200Response) SetData(v []CryptocurrencyResponseItem)`

SetData sets Data field to given value.

### HasData

`func (o *GetCryptocurrencies200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *GetCryptocurrencies200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCryptocurrencies200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCryptocurrencies200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCryptocurrencies200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


