# ResponseTimeDegradationAutodetectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseTimeDegradationMilliseconds** | **int32** | Alert if the response time degrades by more than *X* milliseconds. | [default to null]
**ResponseTimeDegradationPercent** | **int32** | Alert if the response time degrades by more than *X* %. | [default to null]
**SlowestResponseTimeDegradationMilliseconds** | **int32** | Alert if the response time of the slowest 10% degrades by more than *X* milliseconds. | [default to null]
**SlowestResponseTimeDegradationPercent** | **int32** | Alert if the response time of the slowest 10% degrades by more than *X* %. | [default to null]
**LoadThreshold** | **string** | Mininal service load to detect response time degradation.   Response time degradation of services with smaller load won&#39;t trigger alerts. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


