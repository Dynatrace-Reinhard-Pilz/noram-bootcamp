# CustomService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [***ConfigurationMetadata**](ConfigurationMetadata.md) | Metadata useful for debugging. | [optional] [default to null]
**Id** | **string** | The ID of the custom service. | [optional] [default to null]
**Name** | **string** | The name of the custom service, displayed in the UI. | [default to null]
**Enabled** | **bool** | Custom service enabled/disabled. | [default to null]
**Rules** | [**[]DetectionRule**](DetectionRule.md) | The list of rules defining the custom service. | [default to null]
**QueueEntryPoint** | **bool** | The queue entry point flag.   Set to &#x60;true&#x60; for custom messaging services. | [default to null]
**ProcessGroups** | **[]string** | The list of process groups the custom service should belong to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


