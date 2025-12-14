# GetSplits200ResponseSplitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Stands for the split date | [optional] 
**Description** | Pointer to **string** | Specification of the split event | [optional] 
**Ratio** | Pointer to **float64** | The ratio by which the number of a company&#39;s outstanding shares of stock are increased following a stock split. For example, a &#x60;4-for-1 split&#x60; results in four times as many outstanding shares, with each share selling at one forth of its pre-split price | [optional] 
**FromFactor** | Pointer to **float64** | From factor of the split | [optional] 
**ToFactor** | Pointer to **float64** | To factor of the split | [optional] 

## Methods

### NewGetSplits200ResponseSplitsInner

`func NewGetSplits200ResponseSplitsInner() *GetSplits200ResponseSplitsInner`

NewGetSplits200ResponseSplitsInner instantiates a new GetSplits200ResponseSplitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSplits200ResponseSplitsInnerWithDefaults

`func NewGetSplits200ResponseSplitsInnerWithDefaults() *GetSplits200ResponseSplitsInner`

NewGetSplits200ResponseSplitsInnerWithDefaults instantiates a new GetSplits200ResponseSplitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetSplits200ResponseSplitsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetSplits200ResponseSplitsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetSplits200ResponseSplitsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetSplits200ResponseSplitsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *GetSplits200ResponseSplitsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSplits200ResponseSplitsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSplits200ResponseSplitsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSplits200ResponseSplitsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRatio

`func (o *GetSplits200ResponseSplitsInner) GetRatio() float64`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetSplits200ResponseSplitsInner) GetRatioOk() (*float64, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetSplits200ResponseSplitsInner) SetRatio(v float64)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetSplits200ResponseSplitsInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetFromFactor

`func (o *GetSplits200ResponseSplitsInner) GetFromFactor() float64`

GetFromFactor returns the FromFactor field if non-nil, zero value otherwise.

### GetFromFactorOk

`func (o *GetSplits200ResponseSplitsInner) GetFromFactorOk() (*float64, bool)`

GetFromFactorOk returns a tuple with the FromFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFactor

`func (o *GetSplits200ResponseSplitsInner) SetFromFactor(v float64)`

SetFromFactor sets FromFactor field to given value.

### HasFromFactor

`func (o *GetSplits200ResponseSplitsInner) HasFromFactor() bool`

HasFromFactor returns a boolean if a field has been set.

### GetToFactor

`func (o *GetSplits200ResponseSplitsInner) GetToFactor() float64`

GetToFactor returns the ToFactor field if non-nil, zero value otherwise.

### GetToFactorOk

`func (o *GetSplits200ResponseSplitsInner) GetToFactorOk() (*float64, bool)`

GetToFactorOk returns a tuple with the ToFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFactor

`func (o *GetSplits200ResponseSplitsInner) SetToFactor(v float64)`

SetToFactor sets ToFactor field to given value.

### HasToFactor

`func (o *GetSplits200ResponseSplitsInner) HasToFactor() bool`

HasToFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


