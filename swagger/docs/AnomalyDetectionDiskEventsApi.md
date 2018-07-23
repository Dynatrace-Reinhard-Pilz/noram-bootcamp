# \AnomalyDetectionDiskEventsApi

All URIs are relative to *https://siz65484.live.dynatrace.com/api/config/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#CreateDiskEventConfig1) | **Post** /anomalyDetection/diskEvents | Create a new disk event rule.
[**DeleteDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#DeleteDiskEventConfig1) | **Delete** /anomalyDetection/diskEvents/{id} | Delete the disk alert rule with the given id.
[**GetDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#GetDiskEventConfig1) | **Get** /anomalyDetection/diskEvents/{id} | Get the disk event rule with the given id.
[**ListDiskEventConfigs1**](AnomalyDetectionDiskEventsApi.md#ListDiskEventConfigs1) | **Get** /anomalyDetection/diskEvents | Get the list of disk event rules.
[**UpdateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#UpdateDiskEventConfig1) | **Put** /anomalyDetection/diskEvents/{id} | Update the disk event rule with the given id.
[**ValidateDiskEventConfig1**](AnomalyDetectionDiskEventsApi.md#ValidateDiskEventConfig1) | **Post** /anomalyDetection/diskEvents/validator | Validate configuration of disk event rule.


# **CreateDiskEventConfig1**
> DiskEventAnomalyDetectionConfig CreateDiskEventConfig1(ctx, optional)
Create a new disk event rule.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

[**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDiskEventConfig1**
> DeleteDiskEventConfig1(ctx, id)
Delete the disk alert rule with the given id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskEventConfig1**
> DiskEventAnomalyDetectionConfig GetDiskEventConfig1(ctx, id)
Get the disk event rule with the given id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiskEventConfigs1**
> DiskEventAnomalyDetectionConfigList ListDiskEventConfigs1(ctx, )
Get the list of disk event rules.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiskEventAnomalyDetectionConfigList**](DiskEventAnomalyDetectionConfigList.md)

### Authorization

[ReadConfigToken](../README.md#ReadConfigToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDiskEventConfig1**
> UpdateDiskEventConfig1(ctx, id, optional)
Update the disk event rule with the given id.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | [**string**](.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**string**](.md)|  | 
 **body** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDiskEventConfig1**
> ValidateDiskEventConfig1(ctx, optional)
Validate configuration of disk event rule.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DiskEventAnomalyDetectionConfig**](DiskEventAnomalyDetectionConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[WriteConfigToken](../README.md#WriteConfigToken)

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

