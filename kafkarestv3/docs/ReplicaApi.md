# \ReplicaApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet**](ReplicaApi.md#ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id}/replicas/{broker_id} | Get Replica
[**ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet**](ReplicaApi.md#ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id}/replicas | List Replicas



## ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet

> ReplicaData ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet(ctx, clusterId, topicName, partitionId, brokerId)

Get Replica

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the replica for the specified partition assigned to the specified broker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

### Return type

[**ReplicaData**](ReplicaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet

> ReplicaDataList ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet(ctx, clusterId, topicName, partitionId)

List Replicas

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the list of replicas for the specified partition.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ReplicaDataList**](ReplicaDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

