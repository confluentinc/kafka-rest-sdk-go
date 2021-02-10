# \ConsumerGroupApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments | List Consumer Assignments
[**ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments/{topic_name}/partitions/{partition_id} | Get Consumer Assignment
[**ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id} | Get Consumer
[**ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers | List Consumers
[**ClustersClusterIdConsumerGroupsConsumerGroupIdGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id} | Get Consumer Group
[**ClustersClusterIdConsumerGroupsConsumerGroupIdLagGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdLagGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lag | Get Consumer Group Lag.
[**ClustersClusterIdConsumerGroupsConsumerGroupIdLagsGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsConsumerGroupIdLagsGet) | **Get** /clusters/{cluster_id}/consumer-groups/{consumer_group_id}/lags | List Consumer Lags
[**ClustersClusterIdConsumerGroupsGet**](ConsumerGroupApi.md#ClustersClusterIdConsumerGroupsGet) | **Get** /clusters/{cluster_id}/consumer-groups | List Consumer Groups



## ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet

> ConsumerAssignmentDataList ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet(ctx, clusterId, consumerGroupId, consumerId)

List Consumer Assignments

Returns a list of partition assignments for the specified consumer.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet

> ConsumerAssignmentData ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet(ctx, clusterId, consumerGroupId, consumerId, topicName, partitionId)

Get Consumer Assignment

Returns information about the assignment for the specified consumer to the specified partition.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet

> ConsumerData ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet(ctx, clusterId, consumerGroupId, consumerId)

Get Consumer

Returns the consumer specified by the ``consumer_id``.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet

> ConsumerDataList ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet(ctx, clusterId, consumerGroupId)

List Consumers

Returns a list of consumers that belong to the specified consumer group.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdGet

> ConsumerGroupData ClustersClusterIdConsumerGroupsConsumerGroupIdGet(ctx, clusterId, consumerGroupId)

Get Consumer Group

Returns the consumer group specified by the ``consumer_group_id``.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdLagGet

> ConsumerGroupLagData ClustersClusterIdConsumerGroupsConsumerGroupIdLagGet(ctx, clusterId, consumerGroupId)

Get Consumer Group Lag.

Returns the max and total lag of the consumers belonging to the specified consumer group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**consumerGroupId** | **string**| The consumer group ID. | 

### Return type

[**ConsumerGroupLagData**](ConsumerGroupLagData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsConsumerGroupIdLagsGet

> ConsumerLagDataList ClustersClusterIdConsumerGroupsConsumerGroupIdLagsGet(ctx, clusterId, consumerGroupId)

List Consumer Lags

Returns a list of consumer lags of the consumers belonging to the specified consumer group.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdConsumerGroupsGet

> ConsumerGroupDataList ClustersClusterIdConsumerGroupsGet(ctx, clusterId)

List Consumer Groups

Returns the list of consumer groups that belong to the specified Kafka cluster.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

