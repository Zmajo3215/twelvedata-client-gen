# GetTimeSeriesKeltner200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the technical indicator | [optional] 
**TimePeriod** | Pointer to **int64** | Number of periods to average over | [optional] 
**AtrTimePeriod** | Pointer to **int64** | The time period used for calculating the Average True Range | [optional] 
**Multiplier** | Pointer to **int64** | The factor used to adjust the indicator&#39;s sensitivity | [optional] 
**SeriesType** | Pointer to **string** | Price type on which technical indicator is calculated | [optional] 
**MaType** | Pointer to **string** | The type of moving average used | [optional] 

## Methods

### NewGetTimeSeriesKeltner200ResponseMetaIndicator

`func NewGetTimeSeriesKeltner200ResponseMetaIndicator() *GetTimeSeriesKeltner200ResponseMetaIndicator`

NewGetTimeSeriesKeltner200ResponseMetaIndicator instantiates a new GetTimeSeriesKeltner200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKeltner200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesKeltner200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesKeltner200ResponseMetaIndicator`

NewGetTimeSeriesKeltner200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesKeltner200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetAtrTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetAtrTimePeriod() int64`

GetAtrTimePeriod returns the AtrTimePeriod field if non-nil, zero value otherwise.

### GetAtrTimePeriodOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetAtrTimePeriodOk() (*int64, bool)`

GetAtrTimePeriodOk returns a tuple with the AtrTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtrTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetAtrTimePeriod(v int64)`

SetAtrTimePeriod sets AtrTimePeriod field to given value.

### HasAtrTimePeriod

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasAtrTimePeriod() bool`

HasAtrTimePeriod returns a boolean if a field has been set.

### GetMultiplier

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMultiplier() int64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMultiplierOk() (*int64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetMultiplier(v int64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetSeriesType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetMaType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMaType() string`

GetMaType returns the MaType field if non-nil, zero value otherwise.

### GetMaTypeOk

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMaTypeOk() (*string, bool)`

GetMaTypeOk returns a tuple with the MaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetMaType(v string)`

SetMaType sets MaType field to given value.

### HasMaType

`func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) HasMaType() bool`

HasMaType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


