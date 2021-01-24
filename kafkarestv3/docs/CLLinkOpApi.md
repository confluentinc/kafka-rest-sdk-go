# \CLLinkOpApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdLinksLinkNameConfigsalterPost**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsalterPost) | **Post** /clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Topic Configs
[**V3ClustersClusterIdLinksGet**](CLLinkOpApi.md#V3ClustersClusterIdLinksGet) | **Get** /v3/clusters/{cluster_id}/links | List all cluster links in the given cluster
[**V3ClustersClusterIdLinksLinkNameConfigsConfigNameDelete**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameConfigsConfigNameDelete) | **Delete** /v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the given config to default value
[**V3ClustersClusterIdLinksLinkNameConfigsConfigNameGet**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameConfigsConfigNameGet) | **Get** /v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | List the given config
[**V3ClustersClusterIdLinksLinkNameConfigsConfigNamePut**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameConfigsConfigNamePut) | **Put** /v3/clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Update the given config
[**V3ClustersClusterIdLinksLinkNameConfigsGet**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameConfigsGet) | **Get** /v3/clusters/{cluster_id}/links/{link_name}/configs | List the configs of the given link
[**V3ClustersClusterIdLinksLinkNameDelete**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameDelete) | **Delete** /v3/clusters/{cluster_id}/links/{link_name} | Delete the given link
[**V3ClustersClusterIdLinksLinkNameGet**](CLLinkOpApi.md#V3ClustersClusterIdLinksLinkNameGet) | **Get** /v3/clusters/{cluster_id}/links/{link_name} | List the info of the given link
[**V3ClustersClusterIdLinksPost**](CLLinkOpApi.md#V3ClustersClusterIdLinksPost) | **Post** /v3/clusters/{cluster_id}/links | Create a new link



## ClustersClusterIdLinksLinkNameConfigsalterPost

> ClustersClusterIdLinksLinkNameConfigsalterPost(ctx, clusterId, linkName, optional)

Batch Alter Topic Configs

Updates or deletes a set of link configs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameConfigsalterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameConfigsalterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksGet

> ListLinksResponseDataList V3ClustersClusterIdLinksGet(ctx, clusterId)

List all cluster links in the given cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ListLinksResponseDataList**](ListLinksResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameConfigsConfigNameDelete

> V3ClustersClusterIdLinksLinkNameConfigsConfigNameDelete(ctx, clusterId, linkName, configName)

Reset the given config to default value

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameConfigsConfigNameGet

> ListLinkConfigsResponseData V3ClustersClusterIdLinksLinkNameConfigsConfigNameGet(ctx, clusterId, linkName, configName)

List the given config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

### Return type

[**ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameConfigsConfigNamePut

> V3ClustersClusterIdLinksLinkNameConfigsConfigNamePut(ctx, clusterId, linkName, configName, optional)

Update the given config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 
 **optional** | ***V3ClustersClusterIdLinksLinkNameConfigsConfigNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V3ClustersClusterIdLinksLinkNameConfigsConfigNamePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**optional.Interface of UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md)| Link config value to update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameConfigsGet

> ListLinkConfigsResponseDataList V3ClustersClusterIdLinksLinkNameConfigsGet(ctx, clusterId, linkName)

List the configs of the given link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinkConfigsResponseDataList**](ListLinkConfigsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameDelete

> V3ClustersClusterIdLinksLinkNameDelete(ctx, clusterId, linkName)

Delete the given link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksLinkNameGet

> ListLinksResponseData V3ClustersClusterIdLinksLinkNameGet(ctx, clusterId, linkName)

List the info of the given link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinksResponseData**](ListLinksResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ClustersClusterIdLinksPost

> V3ClustersClusterIdLinksPost(ctx, clusterId, optional)

Create a new link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***V3ClustersClusterIdLinksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V3ClustersClusterIdLinksPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLinkRequestData** | [**optional.Interface of CreateLinkRequestData**](CreateLinkRequestData.md)| Create a cluster link | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

