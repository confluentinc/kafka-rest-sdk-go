# \PartitionApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsPartitionsReassignmentGet**](PartitionApi.md#ClustersClusterIdTopicsPartitionsReassignmentGet) | **Get** /clusters/{cluster_id}/topics/-/partitions/-/reassignment | List All Replica Reassignments
[**ClustersClusterIdTopicsTopicNamePartitionsGet**](PartitionApi.md#ClustersClusterIdTopicsTopicNamePartitionsGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions | List Partitions
[**ClustersClusterIdTopicsTopicNamePartitionsPartitionIdGet**](PartitionApi.md#ClustersClusterIdTopicsTopicNamePartitionsPartitionIdGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id} | Get Partition
[**ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReassignmentGet**](PartitionApi.md#ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReassignmentGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/{partition_id}/reassignment | Get Replica Reassignments
[**ClustersClusterIdTopicsTopicNamePartitionsReassignmentGet**](PartitionApi.md#ClustersClusterIdTopicsTopicNamePartitionsReassignmentGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/partitions/-/reassignment | Search Replica Reassignments By Topic



## ClustersClusterIdTopicsPartitionsReassignmentGet

> ReassignmentDataList ClustersClusterIdTopicsPartitionsReassignmentGet(ctx, clusterId)

List All Replica Reassignments

Returns the list of all ongoing replica reassignments in the given Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ReassignmentDataList**](ReassignmentDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsGet

> PartitionDataList ClustersClusterIdTopicsTopicNamePartitionsGet(ctx, clusterId, topicName)

List Partitions

Returns the list of partitions that belong to the specified topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**PartitionDataList**](PartitionDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsPartitionIdGet

> PartitionData ClustersClusterIdTopicsTopicNamePartitionsPartitionIdGet(ctx, clusterId, topicName, partitionId)

Get Partition

Returns the partition with the given `partition_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**PartitionData**](PartitionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReassignmentGet

> ReassignmentData ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReassignmentGet(ctx, clusterId, topicName, partitionId)

Get Replica Reassignments

Returns the list of ongoing replica reassignments for the given partition.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**partitionId** | **int32**| The partition ID. | 

### Return type

[**ReassignmentData**](ReassignmentData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNamePartitionsReassignmentGet

> ReassignmentDataList ClustersClusterIdTopicsTopicNamePartitionsReassignmentGet(ctx, clusterId, topicName)

Search Replica Reassignments By Topic

Returns the list of ongoing replica reassignments for the given topic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**ReassignmentDataList**](ReassignmentDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

