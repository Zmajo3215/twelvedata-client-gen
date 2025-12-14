# ExchangeScheduleResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Official name of exchange | [optional] 
**Name** | Pointer to **string** | Name of exchange | [optional] 
**Code** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Country** | Pointer to **string** | Country to which stock exchange belongs to | [optional] 
**TimeZone** | Pointer to **string** | Time zone where exchange is located | [optional] 
**Sessions** | Pointer to [**[]ExchangeScheduleSession**](ExchangeScheduleSession.md) | Exchange trading hours | [optional] 

## Methods

### NewExchangeScheduleResponseItem

`func NewExchangeScheduleResponseItem() *ExchangeScheduleResponseItem`

NewExchangeScheduleResponseItem instantiates a new ExchangeScheduleResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeScheduleResponseItemWithDefaults

`func NewExchangeScheduleResponseItemWithDefaults() *ExchangeScheduleResponseItem`

NewExchangeScheduleResponseItemWithDefaults instantiates a new ExchangeScheduleResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ExchangeScheduleResponseItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExchangeScheduleResponseItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExchangeScheduleResponseItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExchangeScheduleResponseItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *ExchangeScheduleResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeScheduleResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeScheduleResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeScheduleResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ExchangeScheduleResponseItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExchangeScheduleResponseItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExchangeScheduleResponseItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExchangeScheduleResponseItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCountry

`func (o *ExchangeScheduleResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ExchangeScheduleResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ExchangeScheduleResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ExchangeScheduleResponseItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTimeZone

`func (o *ExchangeScheduleResponseItem) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ExchangeScheduleResponseItem) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ExchangeScheduleResponseItem) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ExchangeScheduleResponseItem) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSessions

`func (o *ExchangeScheduleResponseItem) GetSessions() []ExchangeScheduleSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ExchangeScheduleResponseItem) GetSessionsOk() (*[]ExchangeScheduleSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ExchangeScheduleResponseItem) SetSessions(v []ExchangeScheduleSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ExchangeScheduleResponseItem) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


