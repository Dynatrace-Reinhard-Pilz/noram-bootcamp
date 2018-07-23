# DataSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Data source enbled/disabled. | [default to null]
**Source** | **string** | The source of the attribute to capture. Works in conjunction with **parameterName** or **methods**. | [default to null]
**ValueProcessing** | [***ValueProcessing**](ValueProcessing.md) | Process values as specified. | [optional] [default to null]
**Methods** | [**[]CapturedMethod**](CapturedMethod.md) | The method specification if the **source** value is &#x60;METHOD_PARAM&#x60;   Not applicable in other cases. | [optional] [default to null]
**ParameterName** | **string** | The name of the web request parameter to capture.   Required if the **source** is one of the following: &#x60;POST_PARAMETER&#x60;, &#x60;GET_PARAMETER&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;.   Not applicable in other cases. | [optional] [default to null]
**Scope** | [***ScopeConditions**](ScopeConditions.md) | Conditions for data capturing. | [optional] [default to null]
**CapturingAndStorageLocation** | **string** | Specifies the location where the values are captured and stored.   Required if the **source** is one of the following: &#x60;GET_PARAMETER&#x60;, &#x60;URI&#x60;, &#x60;REQUEST_HEADER&#x60;, &#x60;RESPONSE_HEADER&#x60;.    Not applicable in other cases.    If the **source** value is &#x60;REQUEST_HEADER&#x60; or &#x60;RESPONSE_HEADER&#x60;, the &#x60;CAPTURE_AND_STORE_ON_BOTH&#x60; location is not allowed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


