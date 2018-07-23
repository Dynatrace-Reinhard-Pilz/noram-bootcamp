# RequestAttribute

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [***ConfigurationMetadata**](ConfigurationMetadata.md) | Metadata useful for debugging. | [optional] [default to null]
**Id** | **string** | The ID of the request attribute. | [optional] [default to null]
**Name** | **string** | The name of the request attribute. | [default to null]
**Enabled** | **bool** | Request attribute enabled/disabled. | [default to null]
**DataType** | **string** | The data type of the request attribute. | [default to null]
**DataSources** | [**[]DataSource**](DataSource.md) | The list of data sources. | [default to null]
**Normalization** | **string** | String values transformation.    If the **dataType** is not &#x60;string&#x60;, set the &#x60;Original&#x60; here. | [default to null]
**Aggregation** | **string** | Aggregation type for the request values. | [default to null]
**Confidential** | **bool** | Confidential data flag. Set &#x60;true&#x60; to treat the captured data as confidential. | [default to null]
**SkipPersonalDataMasking** | **bool** | Personal data masking flag. Set &#x60;true&#x60; to skip masking. Warning: This will potentially access personalized data. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


