# ExchangesResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of exchange | [optional] 
**Name** | Pointer to **string** | Name of exchange | [optional] 
**Code** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Country** | Pointer to **string** | Country to which stock exchange belongs to | [optional] 
**Timezone** | Pointer to **string** | Time zone where exchange is located | [optional] 
**Access** | Pointer to [**ExchangesResponseItemAccess**](ExchangesResponseItemAccess.md) |  | [optional] 

## Methods

### NewExchangesResponseItem

`func NewExchangesResponseItem() *ExchangesResponseItem`

NewExchangesResponseItem instantiates a new ExchangesResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangesResponseItemWithDefaults

`func NewExchangesResponseItemWithDefaults() *ExchangesResponseItem`

NewExchangesResponseItemWithDefaults instantiates a new ExchangesResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ExchangesResponseItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExchangesResponseItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExchangesResponseItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExchangesResponseItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *ExchangesResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangesResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangesResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangesResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ExchangesResponseItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExchangesResponseItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExchangesResponseItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExchangesResponseItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCountry

`func (o *ExchangesResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ExchangesResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ExchangesResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ExchangesResponseItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTimezone

`func (o *ExchangesResponseItem) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExchangesResponseItem) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExchangesResponseItem) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExchangesResponseItem) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAccess

`func (o *ExchangesResponseItem) GetAccess() ExchangesResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ExchangesResponseItem) GetAccessOk() (*ExchangesResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ExchangesResponseItem) SetAccess(v ExchangesResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ExchangesResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


