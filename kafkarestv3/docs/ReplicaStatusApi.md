# \ReplicaStatusApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet**](ReplicaStatusApi.md#ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id}/replica-status | List Partition Replica Statuses.
[**ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet**](ReplicaStatusApi.md#ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/-/replica-status | List All Partition Replica Statuses.



## ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet

> ReplicaStatusDataList ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet(ctx, clusterId, topicName, partitionId)

List Partition Replica Statuses.

Return the all the replica statuses for the specified ``topic_name`` and ``partition_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ReplicaStatusDataList**](ReplicaStatusDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet

> ReplicaStatusDataList ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet(ctx, clusterId, topicName)

List All Partition Replica Statuses.

Return the all the replica statuses for the specified ``topic_name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**ReplicaStatusDataList**](ReplicaStatusDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

