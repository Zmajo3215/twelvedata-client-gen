# GetTimeSeriesStochRsi200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the technical indicator | [optional] 
**SeriesType** | Pointer to **string** | Price type on which technical indicator is calculated | [optional] 
**RsiLength** | Pointer to **int64** | Length of period for calculating the RSI component | [optional] 
**StochLength** | Pointer to **int64** | Period length for computing the stochastic oscillator of the RSI | [optional] 
**KPeriod** | Pointer to **int64** | Period for smoothing the %K line | [optional] 
**DPeriod** | Pointer to **int64** | Period for smoothing the %D line, which is a moving average of %K | [optional] 

## Methods

### NewGetTimeSeriesStochRsi200ResponseMetaIndicator

`func NewGetTimeSeriesStochRsi200ResponseMetaIndicator() *GetTimeSeriesStochRsi200ResponseMetaIndicator`

NewGetTimeSeriesStochRsi200ResponseMetaIndicator instantiates a new GetTimeSeriesStochRsi200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStochRsi200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesStochRsi200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStochRsi200ResponseMetaIndicator`

NewGetTimeSeriesStochRsi200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStochRsi200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeriesType

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetRsiLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetRsiLength() int64`

GetRsiLength returns the RsiLength field if non-nil, zero value otherwise.

### GetRsiLengthOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetRsiLengthOk() (*int64, bool)`

GetRsiLengthOk returns a tuple with the RsiLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsiLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetRsiLength(v int64)`

SetRsiLength sets RsiLength field to given value.

### HasRsiLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasRsiLength() bool`

HasRsiLength returns a boolean if a field has been set.

### GetStochLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetStochLength() int64`

GetStochLength returns the StochLength field if non-nil, zero value otherwise.

### GetStochLengthOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetStochLengthOk() (*int64, bool)`

GetStochLengthOk returns a tuple with the StochLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStochLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetStochLength(v int64)`

SetStochLength sets StochLength field to given value.

### HasStochLength

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasStochLength() bool`

HasStochLength returns a boolean if a field has been set.

### GetKPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetKPeriod() int64`

GetKPeriod returns the KPeriod field if non-nil, zero value otherwise.

### GetKPeriodOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetKPeriodOk() (*int64, bool)`

GetKPeriodOk returns a tuple with the KPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetKPeriod(v int64)`

SetKPeriod sets KPeriod field to given value.

### HasKPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasKPeriod() bool`

HasKPeriod returns a boolean if a field has been set.

### GetDPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetDPeriod() int64`

GetDPeriod returns the DPeriod field if non-nil, zero value otherwise.

### GetDPeriodOk

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) GetDPeriodOk() (*int64, bool)`

GetDPeriodOk returns a tuple with the DPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) SetDPeriod(v int64)`

SetDPeriod sets DPeriod field to given value.

### HasDPeriod

`func (o *GetTimeSeriesStochRsi200ResponseMetaIndicator) HasDPeriod() bool`

HasDPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


