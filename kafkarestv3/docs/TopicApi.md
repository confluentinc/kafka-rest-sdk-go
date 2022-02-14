# \TopicApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsGet**](TopicApi.md#ClustersClusterIdTopicsGet) | **Get** /clusters/{cluster_id}/topics | List Topics
[**ClustersClusterIdTopicsPost**](TopicApi.md#ClustersClusterIdTopicsPost) | **Post** /clusters/{cluster_id}/topics | Create Topic
[**ClustersClusterIdTopicsTopicNameDelete**](TopicApi.md#ClustersClusterIdTopicsTopicNameDelete) | **Delete** /clusters/{cluster_id}/topics/{topic_name} | Delete Topic
[**ClustersClusterIdTopicsTopicNameGet**](TopicApi.md#ClustersClusterIdTopicsTopicNameGet) | **Get** /clusters/{cluster_id}/topics/{topic_name} | Get Topic



## ClustersClusterIdTopicsGet

> TopicDataList ClustersClusterIdTopicsGet(ctx, clusterId)

List Topics

Returns the list of topics that belong to the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**TopicDataList**](TopicDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsPost

> TopicData ClustersClusterIdTopicsPost(ctx, clusterId, optional)

Create Topic

Creates a topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdTopicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequestData** | [**optional.Interface of CreateTopicRequestData**](CreateTopicRequestData.md)| The topic creation request. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameDelete

> ClustersClusterIdTopicsTopicNameDelete(ctx, clusterId, topicName)

Delete Topic

Deletes the topic with the given `topic_name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameGet

> TopicData ClustersClusterIdTopicsTopicNameGet(ctx, clusterId, topicName)

Get Topic

Returns the topic with the given `topic_name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

