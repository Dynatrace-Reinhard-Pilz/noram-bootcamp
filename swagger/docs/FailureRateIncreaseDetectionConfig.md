# FailureRateIncreaseDetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMode** | **string** | How to detect failure rate increase: automatically, based on fixed thresholds, or do not detect. | [default to null]
**AutomaticDetection** | [***FailureRateIncreaseAutodetectionConfig**](FailureRateIncreaseAutodetectionConfig.md) | Parameters of the automatic failure rate increase detection. | [optional] [default to null]
**Thresholds** | [***FailureRateIncreaseThresholdConfig**](FailureRateIncreaseThresholdConfig.md) | Fixed thresholds for failure rate increase detection. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


