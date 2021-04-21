# \CLLinkOpApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdLinksGet**](CLLinkOpApi.md#ClustersClusterIdLinksGet) | **Get** /clusters/{cluster_id}/links | List all cluster links in the given cluster
[**ClustersClusterIdLinksLinkNameConfigsConfigNameDelete**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNameDelete) | **Delete** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsConfigNameGet**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNameGet) | **Get** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsConfigNamePut**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNamePut) | **Put** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsGet**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsGet) | **Get** /clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
[**ClustersClusterIdLinksLinkNameConfigsalterPut**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameConfigsalterPut) | **Put** /clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Link Configs
[**ClustersClusterIdLinksLinkNameDelete**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameDelete) | **Delete** /clusters/{cluster_id}/links/{link_name} | Delete the cluster link
[**ClustersClusterIdLinksLinkNameGet**](CLLinkOpApi.md#ClustersClusterIdLinksLinkNameGet) | **Get** /clusters/{cluster_id}/links/{link_name} | Describe the cluster link
[**ClustersClusterIdLinksPost**](CLLinkOpApi.md#ClustersClusterIdLinksPost) | **Post** /clusters/{cluster_id}/links | Create a cluster link



## ClustersClusterIdLinksGet

> ListLinksResponseDataList ClustersClusterIdLinksGet(ctx, clusterId)

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


## ClustersClusterIdLinksLinkNameConfigsConfigNameDelete

> ClustersClusterIdLinksLinkNameConfigsConfigNameDelete(ctx, clusterId, linkName, configName)

Reset the config under the cluster link

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


## ClustersClusterIdLinksLinkNameConfigsConfigNameGet

> ListLinkConfigsResponseData ClustersClusterIdLinksLinkNameConfigsConfigNameGet(ctx, clusterId, linkName, configName)

Describe the config under the cluster link

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


## ClustersClusterIdLinksLinkNameConfigsConfigNamePut

> ClustersClusterIdLinksLinkNameConfigsConfigNamePut(ctx, clusterId, linkName, configName, optional)

Alter the config under the cluster link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 
 **optional** | ***ClustersClusterIdLinksLinkNameConfigsConfigNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameConfigsConfigNamePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**optional.Interface of UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md)| New cluster link config value | 

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


## ClustersClusterIdLinksLinkNameConfigsGet

> ListLinkConfigsResponseDataList ClustersClusterIdLinksLinkNameConfigsGet(ctx, clusterId, linkName)

List all configs of the cluster link

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


## ClustersClusterIdLinksLinkNameConfigsalterPut

> ClustersClusterIdLinksLinkNameConfigsalterPut(ctx, clusterId, linkName, optional)

Batch Alter Link Configs

Batch update configs of the cluster link.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameConfigsalterPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameConfigsalterPutOpts struct


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


## ClustersClusterIdLinksLinkNameDelete

> ClustersClusterIdLinksLinkNameDelete(ctx, clusterId, linkName)

Delete the cluster link

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


## ClustersClusterIdLinksLinkNameGet

> ListLinksResponseData ClustersClusterIdLinksLinkNameGet(ctx, clusterId, linkName)

Describe the cluster link

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


## ClustersClusterIdLinksPost

> ClustersClusterIdLinksPost(ctx, clusterId, linkName, optional)

Create a cluster link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate if the link can be created or not, but not to create it. Default: false | 
 **validateLink** | **optional.Bool**| To synchronously validate that the source cluster ID is expected and the dest cluster has the permission to read topics in the source cluster. default: true | 
 **createLinkRequestData** | [**optional.Interface of CreateLinkRequestData**](CreateLinkRequestData.md)| Source cluster id and cluster link configs | 

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

