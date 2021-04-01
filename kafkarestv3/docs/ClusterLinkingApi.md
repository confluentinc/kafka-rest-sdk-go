# \ClusterLinkingApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdLinksGet**](ClusterLinkingApi.md#ClustersClusterIdLinksGet) | **Get** /clusters/{cluster_id}/links | List all cluster links in the given cluster
[**ClustersClusterIdLinksLinkNameConfigsConfigNameDelete**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNameDelete) | **Delete** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsConfigNameGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNameGet) | **Get** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsConfigNamePut**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameConfigsConfigNamePut) | **Put** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
[**ClustersClusterIdLinksLinkNameConfigsGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameConfigsGet) | **Get** /clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
[**ClustersClusterIdLinksLinkNameConfigsalterPut**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameConfigsalterPut) | **Put** /clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Link Configs
[**ClustersClusterIdLinksLinkNameDelete**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameDelete) | **Delete** /clusters/{cluster_id}/links/{link_name} | Delete the cluster link
[**ClustersClusterIdLinksLinkNameGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameGet) | **Get** /clusters/{cluster_id}/links/{link_name} | Describe the cluster link
[**ClustersClusterIdLinksLinkNameMirrorsFailoverPost**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsFailoverPost) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors/failover | Failover the mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsGet) | **Get** /clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsMirrorTopicNameGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsMirrorTopicNameGet) | **Get** /clusters/{cluster_id}/links/{link_name}/mirrors/{mirror_topic_name} | Describe the mirror topic
[**ClustersClusterIdLinksLinkNameMirrorsPausePost**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsPausePost) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors/pause | Pause the mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsPost**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsPost) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic
[**ClustersClusterIdLinksLinkNameMirrorsPromotePost**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsPromotePost) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors/promote | Promote the mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsResumePost**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsResumePost) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors/resume | Resume the mirror topics
[**ClustersClusterIdLinksMirrorsGet**](ClusterLinkingApi.md#ClustersClusterIdLinksMirrorsGet) | **Get** /clusters/{cluster_id}/links/-/mirrors | List mirror topics
[**ClustersClusterIdLinksPost**](ClusterLinkingApi.md#ClustersClusterIdLinksPost) | **Post** /clusters/{cluster_id}/links | Create a cluster link



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


## ClustersClusterIdLinksLinkNameMirrorsFailoverPost

> AlterMirrorStatusResponseDataList ClustersClusterIdLinksLinkNameMirrorsFailoverPost(ctx, clusterId, linkName, optional)

Failover the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| Is validate-only or not. default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksLinkNameMirrorsGet

> ListMirrorTopicsResponseDataList ClustersClusterIdLinksLinkNameMirrorsGet(ctx, clusterId, linkName, optional)

List mirror topics

List all mirror topics under the link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksLinkNameMirrorsMirrorTopicNameGet

> ListMirrorTopicsResponseData ClustersClusterIdLinksLinkNameMirrorsMirrorTopicNameGet(ctx, clusterId, linkName, mirrorTopicName)

Describe the mirror topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**mirrorTopicName** | **string**| Cluster Linking mirror topic name | 

### Return type

[**ListMirrorTopicsResponseData**](ListMirrorTopicsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksLinkNameMirrorsPausePost

> AlterMirrorStatusResponseDataList ClustersClusterIdLinksLinkNameMirrorsPausePost(ctx, clusterId, linkName, optional)

Pause the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsPausePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsPausePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| Is validate-only or not. default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksLinkNameMirrorsPost

> ClustersClusterIdLinksLinkNameMirrorsPost(ctx, clusterId, linkName, optional)

Create a mirror topic

Create a topic in the destination cluster mirroring a topic in the source cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**optional.Interface of CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md)| Name and configs of the topics mirroring from and mirroring to | 

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


## ClustersClusterIdLinksLinkNameMirrorsPromotePost

> AlterMirrorStatusResponseDataList ClustersClusterIdLinksLinkNameMirrorsPromotePost(ctx, clusterId, linkName, optional)

Promote the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| Is validate-only or not. default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksLinkNameMirrorsResumePost

> AlterMirrorStatusResponseDataList ClustersClusterIdLinksLinkNameMirrorsResumePost(ctx, clusterId, linkName, optional)

Resume the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsResumePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsResumePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| Is validate-only or not. default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdLinksMirrorsGet

> ListMirrorTopicsResponseDataList ClustersClusterIdLinksMirrorsGet(ctx, clusterId, optional)

List mirror topics

List all mirror topics in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdLinksMirrorsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksMirrorsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

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

