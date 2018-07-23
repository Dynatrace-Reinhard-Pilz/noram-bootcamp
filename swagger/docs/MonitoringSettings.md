# MonitoringSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchRequests** | **bool** | Capture fetch() request enabled/disabled | [default to null]
**XmlHttpRequest** | **bool** | Support for XmlHttpRequest enabled/disabled | [default to null]
**JavaScriptFrameworkSupport** | [***JavaScriptFrameworkSupport**](JavaScriptFrameworkSupport.md) | JavaScript framework support. | [default to null]
**ContentCapture** | [***ContentCapture**](ContentCapture.md) | Contains all settings regarding content capturing, e.g. resource timings. | [default to null]
**ExcludeXhrRegex** | **string** | A regular expression to match all URLs that should be excluded from becoming XHR actions.  An empty value disables the feature. | [default to null]
**InjectionMode** | **string** | JavaScript injection mode. | [default to null]
**LibraryFileLocation** | **string** | The source path for placement of your application’s custom JavaScript library file.  An empty value sets it to the root directory of your web server. | [default to null]
**MonitoringDataPath** | **string** | The path where the JavaScript tag should send monitoring data.  Specify either a relative or an absolute URL. If you enter an absolute URL, data will be sent using CORS. | [default to null]
**CustomConfigurationProperties** | **string** | Additional JavaScript tag properties that are specific to your application. To do this, type key-value pairs defined using (&#x3D;) and separated using a (|) symbol. | [default to null]
**ServerRequestPathId** | **string** | Path that is to be used to identify the server’s request ID. | [default to null]
**SecureCookieAttribute** | **bool** | Use the secure cookie attribute for cookies set by Dynatrace enabled/disabled. | [default to null]
**CookiePlacementDomain** | **string** | Domain to be used for cookie placement. | [default to null]
**CacheControlHeaderOptimizations** | **bool** | Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled. | [default to null]
**AdvancedJavaScriptTagSettings** | [***AdvancedJavaScriptTagSettings**](AdvancedJavaScriptTagSettings.md) | Advanced JavaScript tag settings. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


