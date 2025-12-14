# GetTimeSeriesIchimoku200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the technical indicator | [optional] 
**ConversionLinePeriod** | Pointer to **int64** | The time period used for generating the conversation line | [optional] 
**BaseLinePeriod** | Pointer to **int64** | The time period used for generating the base line | [optional] 
**LeadingSpanBPeriod** | Pointer to **int64** | The time period used for generating the leading span B line | [optional] 
**LaggingSpanPeriod** | Pointer to **int64** | The time period used for generating the lagging span line | [optional] 
**IncludeAheadSpanPeriod** | Pointer to **bool** | Indicates whether to include ahead span period | [optional] 

## Methods

### NewGetTimeSeriesIchimoku200ResponseMetaIndicator

`func NewGetTimeSeriesIchimoku200ResponseMetaIndicator() *GetTimeSeriesIchimoku200ResponseMetaIndicator`

NewGetTimeSeriesIchimoku200ResponseMetaIndicator instantiates a new GetTimeSeriesIchimoku200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesIchimoku200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesIchimoku200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesIchimoku200ResponseMetaIndicator`

NewGetTimeSeriesIchimoku200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesIchimoku200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConversionLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetConversionLinePeriod() int64`

GetConversionLinePeriod returns the ConversionLinePeriod field if non-nil, zero value otherwise.

### GetConversionLinePeriodOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetConversionLinePeriodOk() (*int64, bool)`

GetConversionLinePeriodOk returns a tuple with the ConversionLinePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetConversionLinePeriod(v int64)`

SetConversionLinePeriod sets ConversionLinePeriod field to given value.

### HasConversionLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasConversionLinePeriod() bool`

HasConversionLinePeriod returns a boolean if a field has been set.

### GetBaseLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetBaseLinePeriod() int64`

GetBaseLinePeriod returns the BaseLinePeriod field if non-nil, zero value otherwise.

### GetBaseLinePeriodOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetBaseLinePeriodOk() (*int64, bool)`

GetBaseLinePeriodOk returns a tuple with the BaseLinePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetBaseLinePeriod(v int64)`

SetBaseLinePeriod sets BaseLinePeriod field to given value.

### HasBaseLinePeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasBaseLinePeriod() bool`

HasBaseLinePeriod returns a boolean if a field has been set.

### GetLeadingSpanBPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLeadingSpanBPeriod() int64`

GetLeadingSpanBPeriod returns the LeadingSpanBPeriod field if non-nil, zero value otherwise.

### GetLeadingSpanBPeriodOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLeadingSpanBPeriodOk() (*int64, bool)`

GetLeadingSpanBPeriodOk returns a tuple with the LeadingSpanBPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadingSpanBPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetLeadingSpanBPeriod(v int64)`

SetLeadingSpanBPeriod sets LeadingSpanBPeriod field to given value.

### HasLeadingSpanBPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasLeadingSpanBPeriod() bool`

HasLeadingSpanBPeriod returns a boolean if a field has been set.

### GetLaggingSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLaggingSpanPeriod() int64`

GetLaggingSpanPeriod returns the LaggingSpanPeriod field if non-nil, zero value otherwise.

### GetLaggingSpanPeriodOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetLaggingSpanPeriodOk() (*int64, bool)`

GetLaggingSpanPeriodOk returns a tuple with the LaggingSpanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaggingSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetLaggingSpanPeriod(v int64)`

SetLaggingSpanPeriod sets LaggingSpanPeriod field to given value.

### HasLaggingSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasLaggingSpanPeriod() bool`

HasLaggingSpanPeriod returns a boolean if a field has been set.

### GetIncludeAheadSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetIncludeAheadSpanPeriod() bool`

GetIncludeAheadSpanPeriod returns the IncludeAheadSpanPeriod field if non-nil, zero value otherwise.

### GetIncludeAheadSpanPeriodOk

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) GetIncludeAheadSpanPeriodOk() (*bool, bool)`

GetIncludeAheadSpanPeriodOk returns a tuple with the IncludeAheadSpanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAheadSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) SetIncludeAheadSpanPeriod(v bool)`

SetIncludeAheadSpanPeriod sets IncludeAheadSpanPeriod field to given value.

### HasIncludeAheadSpanPeriod

`func (o *GetTimeSeriesIchimoku200ResponseMetaIndicator) HasIncludeAheadSpanPeriod() bool`

HasIncludeAheadSpanPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


