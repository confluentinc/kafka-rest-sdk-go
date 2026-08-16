# \LicenseV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicense**](LicenseV3Api.md#GetLicense) | **Get** /clusters/{cluster_id}/licenses/{category_short_name} | Get license
[**GetLicenses**](LicenseV3Api.md#GetLicenses) | **Get** /clusters/{cluster_id}/licenses | List licenses
[**UpdateLicense**](LicenseV3Api.md#UpdateLicense) | **Put** /clusters/{cluster_id}/licenses | Update (validate and store) a license



## GetLicense

> LicenseData GetLicense(ctx, clusterId, categoryShortName)

Get license

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Get the license with the specified category short name for the cluster specified by ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**categoryShortName** | **string**| The short name of the license category. | 

### Return type

[**LicenseData**](LicenseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenses

> LicenseDataList GetLicenses(ctx, clusterId)

List licenses

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  List the licenses for the cluster specified by ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**LicenseDataList**](LicenseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> UpdateLicenseResponseData UpdateLicense(ctx, clusterId, updateLicenseRequestData, optional)

Update (validate and store) a license

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Validate and store a license for the cluster specified by ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**updateLicenseRequestData** | [**UpdateLicenseRequestData**](UpdateLicenseRequestData.md)| The license to validate and store. | 
 **optional** | ***UpdateLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLicenseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dryRun** | **optional.Bool**| To validate the license without storing it. Default: false | [default to false]

### Return type

[**UpdateLicenseResponseData**](UpdateLicenseResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

