# \AnomalyDetectionServicesApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration1**](AnomalyDetectionServicesApi.md#GetConfiguration1) | **Get** /anomalyDetection/services | Gets the service anomaly detection configuration
[**UpdateConfiguration1**](AnomalyDetectionServicesApi.md#UpdateConfiguration1) | **Put** /anomalyDetection/services | Updates the service anomaly detection configuration
[**ValidateConfiguration1**](AnomalyDetectionServicesApi.md#ValidateConfiguration1) | **Post** /anomalyDetection/services/validator | Validates the service anomaly detection configuration


# **GetConfiguration1**
> ServiceAnomalyDetectionConfig GetConfiguration1(ctx, )
Gets the service anomaly detection configuration



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfiguration1**
> ServiceAnomalyDetectionConfig UpdateConfiguration1(ctx, optional)
Updates the service anomaly detection configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)|  | 

### Return type

[**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConfiguration1**
> ValidateConfiguration1(ctx, optional)
Validates the service anomaly detection configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

