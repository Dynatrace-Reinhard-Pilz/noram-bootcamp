# AdvancedJavaScriptTagSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncBeaconFirefox** | **bool** | Send the beacon signal as a synchronous XMLHttpRequest using Firefox enabled/disabled. | [default to null]
**SyncBeaconInternetExplorer** | **bool** | Send the beacon signal as a synchronous XMLHttpRequest using Internet Explorer enabled/disabled. | [default to null]
**InstrumentUnsupportedAjaxFrameworks** | **bool** | Instrumentation of unsupported Ajax frameworks enabled/disabled. | [default to null]
**SpecialCharactersToEscape** | **string** | Additional special characters that are to be escaped using non-alphanumeric characters in HTML escape format. | [default to null]
**MaxActionNameLength** | **int32** | Maximum character length for action names. Valid values range from 5 to 10000. | [default to null]
**MaxErrorsToCapture** | **int32** | Maximum number of errors to be captured per page. Valid values range from 0 to 50. | [default to null]
**AdditionalEventHandlers** | [***AdditionalEventHandlers**](AdditionalEventHandlers.md) | Additional event handlers and wrappers. | [default to null]
**EventWrapperSettings** | [***EventWrapperSettings**](EventWrapperSettings.md) | In addition to the event handlers, events called using addEventListener or attachEvent can be captured. Be careful with this option! Event wrappers can conflict with the JavaScript code on a web page. | [default to null]
**GlobalEventCaptureSettings** | [***GlobalEventCaptureSettings**](GlobalEventCaptureSettings.md) | Global event capture. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


