# ResponseTimeDegradationDetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMode** | **string** | How to detect response time degradation: automatically, based on fixed thresholds, or do not detect. | [default to null]
**AutomaticDetection** | [***ResponseTimeDegradationAutodetectionConfig**](ResponseTimeDegradationAutodetectionConfig.md) | Parameters of the automatic response time degradation detection. | [optional] [default to null]
**Thresholds** | [***ResponseTimeDegradationThresholdConfig**](ResponseTimeDegradationThresholdConfig.md) | Fixed thresholds for response time degradation detection. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


