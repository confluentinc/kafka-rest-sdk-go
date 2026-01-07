# \ConsumerGroupV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaConsumer**](ConsumerGroupV3Api.md#GetKafkaConsumer) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
[**GetKafkaConsumerAssignment**](ConsumerGroupV3Api.md#GetKafkaConsumerAssignment) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments/{topic_name}/partitions/{partition_id} | Get Consumer Assignment
[**GetKafkaConsumerGroup**](ConsumerGroupV3Api.md#GetKafkaConsumerGroup) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
[**GetKafkaConsumerGroupLagSummary**](ConsumerGroupV3Api.md#GetKafkaConsumerGroupLagSummary) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag-summary | Get Consumer Group Lag Summary
[**GetKafkaConsumerLag**](ConsumerGroupV3Api.md#GetKafkaConsumerLag) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags/{topic_name}/partitions/{partition_id} | Get Consumer Lag
[**ListKafkaConsumerAssignment**](ConsumerGroupV3Api.md#ListKafkaConsumerAssignment) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments | List Consumer Assignments
[**ListKafkaConsumerGroups**](ConsumerGroupV3Api.md#ListKafkaConsumerGroups) | **Get** /clusters/{cluster_id}/consumer-groups | List Consumer Groups
[**ListKafkaConsumerLags**](ConsumerGroupV3Api.md#ListKafkaConsumerLags) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
[**ListKafkaConsumers**](ConsumerGroupV3Api.md#ListKafkaConsumers) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers



## GetKafkaConsumer

> ConsumerData GetKafkaConsumer(ctx, clusterId, consumerGroupId, consumerId)

Get Consumer

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the consumer specified by the ``consumer_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 
**consumerId** | **string**| The consumer ID. | 

### Return type

[**ConsumerData**](ConsumerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerAssignment

> ConsumerAssignmentData GetKafkaConsumerAssignment(ctx, clusterId, consumerGroupId, consumerId, topicName, partitionId)

Get Consumer Assignment

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return information about the assignment for the specified consumer to the specified partition.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 
**consumerId** | **string**| The consumer ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ConsumerAssignmentData**](ConsumerAssignmentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroup

> ConsumerGroupData GetKafkaConsumerGroup(ctx, clusterId, consumerGroupId)

Get Consumer Group

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the consumer group specified by the ``consumer_group_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 

### Return type

[**ConsumerGroupData**](ConsumerGroupData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroupLagSummary

> ConsumerGroupLagSummaryData GetKafkaConsumerGroupLagSummary(ctx, clusterId, consumerGroupId)

Get Consumer Group Lag Summary

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Available in dedicated clusters only](https://img.shields.io/badge/-Available%20in%20dedicated%20clusters%20only-%23bc8540)](https://docs.confluent.io/cloud/current/clusters/cluster-types.html#dedicated-cluster)  Return the maximum and total lag of the consumers belonging to the specified consumer group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 

### Return type

[**ConsumerGroupLagSummaryData**](ConsumerGroupLagSummaryData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerLag

> ConsumerLagData GetKafkaConsumerLag(ctx, clusterId, consumerGroupId, topicName, partitionId)

Get Consumer Lag

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Available in dedicated clusters only](https://img.shields.io/badge/-Available%20in%20dedicated%20clusters%20only-%23bc8540)](https://docs.confluent.io/cloud/current/clusters/cluster-types.html#dedicated-cluster)  Return the consumer lag on a partition with the given `partition_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ConsumerLagData**](ConsumerLagData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerAssignment

> ConsumerAssignmentDataList ListKafkaConsumerAssignment(ctx, clusterId, consumerGroupId, consumerId)

List Consumer Assignments

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return a list of partition assignments for the specified consumer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 
**consumerId** | **string**| The consumer ID. | 

### Return type

[**ConsumerAssignmentDataList**](ConsumerAssignmentDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerGroups

> ConsumerGroupDataList ListKafkaConsumerGroups(ctx, clusterId)

List Consumer Groups

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of consumer groups that belong to the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ConsumerGroupDataList**](ConsumerGroupDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerLags

> ConsumerLagDataList ListKafkaConsumerLags(ctx, clusterId, consumerGroupId)

List Consumer Lags

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Available in dedicated clusters only](https://img.shields.io/badge/-Available%20in%20dedicated%20clusters%20only-%23bc8540)](https://docs.confluent.io/cloud/current/clusters/cluster-types.html#dedicated-cluster)  Return a list of consumer lags of the consumers belonging to the specified consumer group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 

### Return type

[**ConsumerLagDataList**](ConsumerLagDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumers

> ConsumerDataList ListKafkaConsumers(ctx, clusterId, consumerGroupId)

List Consumers

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return a list of consumers that belong to the specified consumer group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 

### Return type

[**ConsumerDataList**](ConsumerDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

