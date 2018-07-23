# DiskEventAnomalyDetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [***ConfigurationMetadata**](ConfigurationMetadata.md) | Metadata useful for debugging. | [optional] [default to null]
**Id** | **string** | The ID of the disk event. | [optional] [default to null]
**Name** | **string** | The name of the disk event. | [default to null]
**Enabled** | **bool** | Whether the disk event is enabled. | [default to null]
**Metric** | **string** | The metric to monitor. | [default to null]
**Threshold** | **float64** | The event threshold. If the metric is &#x60;LowDiskSpace&#x60; or &#x60;LowInodes&#x60;, this is a percentage. For &#x60;ReadTimeExceeding&#x60; or &#x60;WriteTimeExceeding&#x60;, milliseconds. | [default to null]
**ViolatingSamples** | **int32** | The minimum number of violating samples for an event to be triggered. Must not exceed &#x60;samples&#x60;. | [default to null]
**Samples** | **int32** | The number of samples that are inspected. | [default to null]
**DiskNameFilter** | [***DiskNameFilterDto**](DiskNameFilterDto.md) | Only apply this rule to disks matching the given name. | [optional] [default to null]
**TagFilters** | [**[]TagFilterDto**](TagFilterDto.md) | Only apply this rule to hosts having these tags. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


