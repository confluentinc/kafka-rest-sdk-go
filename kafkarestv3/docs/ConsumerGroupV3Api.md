# \ConsumerGroupV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaConsumer**](ConsumerGroupV3Api.md#GetKafkaConsumer) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
[**GetKafkaConsumerAssignment**](ConsumerGroupV3Api.md#GetKafkaConsumerAssignment) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments/{topic_name}/partitions/{partition_id} | Get Consumer Assignment
[**GetKafkaConsumerGroup**](ConsumerGroupV3Api.md#GetKafkaConsumerGroup) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
[**GetKafkaConsumerGroupLagSummary**](ConsumerGroupV3Api.md#GetKafkaConsumerGroupLagSummary) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag-summary | Get Consumer Group Lag Summary.
[**ListKafkaConsumerAssignment**](ConsumerGroupV3Api.md#ListKafkaConsumerAssignment) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments | List Consumer Assignments
[**ListKafkaConsumerGroups**](ConsumerGroupV3Api.md#ListKafkaConsumerGroups) | **Get** /clusters/{cluster_id}/consumer-groups | List Consumer Groups
[**ListKafkaConsumerLags**](ConsumerGroupV3Api.md#ListKafkaConsumerLags) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
[**ListKafkaConsumers**](ConsumerGroupV3Api.md#ListKafkaConsumers) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers



## GetKafkaConsumer

> ConsumerData GetKafkaConsumer(ctx, clusterId, consumerGroupId, consumerId)

Get Consumer

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Cluster Admin for Kafka (v3)](https://img.shields.io/badge/-Request%20Access%20To%20Cluster%20Admin%20For%20Kafka%20v3-%23bc8540)](mailto:ccloud-rest-api+consumer-lag-earlyaccess@confluent.io?subject=Request%20to%20join%20v3%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cluster%20Admin%20For%20Kafka%20v3%20Early%20Access%20to%20provide%20early%20feedback%20on%20consumer%20lag%20apis%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)  Returns the consumer specified by the ``consumer_id``.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerAssignment

> ConsumerAssignmentData GetKafkaConsumerAssignment(ctx, clusterId, consumerGroupId, consumerId, topicName, partitionId)

Get Consumer Assignment

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns information about the assignment for the specified consumer to the specified partition.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroup

> ConsumerGroupData GetKafkaConsumerGroup(ctx, clusterId, consumerGroupId)

Get Consumer Group

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the consumer group specified by the ``consumer_group_id``.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaConsumerGroupLagSummary

> ConsumerGroupLagSummaryData GetKafkaConsumerGroupLagSummary(ctx, clusterId, consumerGroupId)

Get Consumer Group Lag Summary.

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Cluster Admin for Kafka (v3)](https://img.shields.io/badge/-Request%20Access%20To%20Cluster%20Admin%20For%20Kafka%20v3-%23bc8540)](mailto:ccloud-rest-api+consumer-lag-earlyaccess@confluent.io?subject=Request%20to%20join%20v3%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cluster%20Admin%20For%20Kafka%20v3%20Early%20Access%20to%20provide%20early%20feedback%20on%20consumer%20lag%20apis%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)  Returns the max and total lag of the consumers belonging to the specified consumer group.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerAssignment

> ConsumerAssignmentDataList ListKafkaConsumerAssignment(ctx, clusterId, consumerGroupId, consumerId)

List Consumer Assignments

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of partition assignments for the specified consumer.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerGroups

> ConsumerGroupDataList ListKafkaConsumerGroups(ctx, clusterId)

List Consumer Groups

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the list of consumer groups that belong to the specified Kafka cluster.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumerLags

> ConsumerLagDataList ListKafkaConsumerLags(ctx, clusterId, consumerGroupId)

List Consumer Lags

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Cluster Admin for Kafka (v3)](https://img.shields.io/badge/-Request%20Access%20To%20Cluster%20Admin%20For%20Kafka%20v3-%23bc8540)](mailto:ccloud-rest-api+consumer-lag-earlyaccess@confluent.io?subject=Request%20to%20join%20v3%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cluster%20Admin%20For%20Kafka%20v3%20Early%20Access%20to%20provide%20early%20feedback%20on%20consumer%20lag%20apis%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)  Returns a list of consumer lags of the consumers belonging to the specified consumer group.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaConsumers

> ConsumerDataList ListKafkaConsumers(ctx, clusterId, consumerGroupId)

List Consumers

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of consumers that belong to the specified consumer group.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

