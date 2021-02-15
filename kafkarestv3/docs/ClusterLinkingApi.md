# \ClusterLinkingApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterClusterIdLinksLinkNameMirrorsDestinationTopicNameGet**](ClusterLinkingApi.md#ClusterClusterIdLinksLinkNameMirrorsDestinationTopicNameGet) | **Get** /cluster/{cluster_id}/links/{link_name}/mirrors/{destination_topic_name} | Describe the mirror topic
[**ClusterClusterIdLinksLinkNameMirrorsFailoverPost**](ClusterLinkingApi.md#ClusterClusterIdLinksLinkNameMirrorsFailoverPost) | **Post** /cluster/{cluster_id}/links/{link_name}/mirrors/failover | Failover the mirror topics
[**ClusterClusterIdLinksLinkNameMirrorsPausePost**](ClusterLinkingApi.md#ClusterClusterIdLinksLinkNameMirrorsPausePost) | **Post** /cluster/{cluster_id}/links/{link_name}/mirrors/pause | Pause the mirror topics
[**ClusterClusterIdLinksLinkNameMirrorsPromotePost**](ClusterLinkingApi.md#ClusterClusterIdLinksLinkNameMirrorsPromotePost) | **Post** /cluster/{cluster_id}/links/{link_name}/mirrors/promote | Promote the mirror topics
[**ClusterClusterIdLinksLinkNameMirrorsResumePost**](ClusterLinkingApi.md#ClusterClusterIdLinksLinkNameMirrorsResumePost) | **Post** /cluster/{cluster_id}/links/{link_name}/mirrors/resume | Resume the mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsGet**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsGet) | **Get** /clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
[**ClustersClusterIdLinksLinkNameMirrorsPut**](ClusterLinkingApi.md#ClustersClusterIdLinksLinkNameMirrorsPut) | **Put** /clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic



## ClusterClusterIdLinksLinkNameMirrorsDestinationTopicNameGet

> ListMirrorTopicsResponseData ClusterClusterIdLinksLinkNameMirrorsDestinationTopicNameGet(ctx, clusterId, linkName, destinationTopicName)

Describe the mirror topic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**destinationTopicName** | **string**| Cluster Linking destination topic name | 

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


## ClusterClusterIdLinksLinkNameMirrorsFailoverPost

> AlterMirrorStatusResponseDataList ClusterClusterIdLinksLinkNameMirrorsFailoverPost(ctx, clusterId, linkName, destinationTopicName, body)

Failover the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**destinationTopicName** | **string**| Cluster Linking destination topic name | 
**body** | **AlterMirrorsRequest**|  | 

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


## ClusterClusterIdLinksLinkNameMirrorsPausePost

> AlterMirrorStatusResponseDataList ClusterClusterIdLinksLinkNameMirrorsPausePost(ctx, clusterId, linkName, destinationTopicName, body)

Pause the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**destinationTopicName** | **string**| Cluster Linking destination topic name | 
**body** | **AlterMirrorsRequest**|  | 

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


## ClusterClusterIdLinksLinkNameMirrorsPromotePost

> AlterMirrorStatusResponseDataList ClusterClusterIdLinksLinkNameMirrorsPromotePost(ctx, clusterId, linkName, destinationTopicName, body)

Promote the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**destinationTopicName** | **string**| Cluster Linking destination topic name | 
**body** | **AlterMirrorsRequest**|  | 

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


## ClusterClusterIdLinksLinkNameMirrorsResumePost

> AlterMirrorStatusResponseDataList ClusterClusterIdLinksLinkNameMirrorsResumePost(ctx, clusterId, linkName, destinationTopicName, body)

Resume the mirror topics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**destinationTopicName** | **string**| Cluster Linking destination topic name | 
**body** | **AlterMirrorsRequest**|  | 

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


## ClustersClusterIdLinksLinkNameMirrorsPut

> ClustersClusterIdLinksLinkNameMirrorsPut(ctx, clusterId, linkName, optional)

Create a mirror topic

Create a topic in the destination cluster mirroring a topic in the source cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ClustersClusterIdLinksLinkNameMirrorsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdLinksLinkNameMirrorsPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**optional.Interface of CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md)| Name of the topics mirroring from and mirroring to | 

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

