# \TopicV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaTopic**](TopicV3Api.md#CreateKafkaTopic) | **Post** /clusters/{cluster_id}/topics | Create Topic
[**DeleteKafkaTopic**](TopicV3Api.md#DeleteKafkaTopic) | **Delete** /clusters/{cluster_id}/topics/{topic_name} | Delete Topic
[**GetKafkaTopic**](TopicV3Api.md#GetKafkaTopic) | **Get** /clusters/{cluster_id}/topics/{topic_name} | Get Topic
[**ListKafkaTopics**](TopicV3Api.md#ListKafkaTopics) | **Get** /clusters/{cluster_id}/topics | List Topics
[**UpdatePartitionCountKafkaTopic**](TopicV3Api.md#UpdatePartitionCountKafkaTopic) | **Patch** /clusters/{cluster_id}/topics/{topic_name} | Update Partition Count



## CreateKafkaTopic

> TopicData CreateKafkaTopic(ctx, clusterId, optional)

Create Topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create a topic. Also supports a dry-run mode that only validates whether the topic creation would succeed if the ``validate_only`` request property is explicitly specified and set to true.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***CreateKafkaTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequestData** | [**optional.Interface of CreateTopicRequestData**](CreateTopicRequestData.md)| The topic creation request. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaTopic

> DeleteKafkaTopic(ctx, clusterId, topicName)

Delete Topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Delete the topic with the given `topic_name`.

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
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaTopic

> TopicData GetKafkaTopic(ctx, clusterId, topicName, optional)

Get Topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the topic with the given `topic_name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***GetKafkaTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKafkaTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAuthorizedOperations** | **optional.Bool**| Specify if authorized operations should be included in the response. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaTopics

> TopicDataList ListKafkaTopics(ctx, clusterId)

List Topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of topics that belong to the specified Kafka cluster.

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
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartitionCountKafkaTopic

> TopicData UpdatePartitionCountKafkaTopic(ctx, clusterId, topicName, optional)

Update Partition Count

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Increase the number of partitions for a topic.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***UpdatePartitionCountKafkaTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePartitionCountKafkaTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePartitionCountRequestData** | [**optional.Interface of UpdatePartitionCountRequestData**](UpdatePartitionCountRequestData.md)| The number of partitions to increase the partition count to. | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

