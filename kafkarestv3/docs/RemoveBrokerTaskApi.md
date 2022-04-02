# \RemoveBrokerTaskApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdRemoveBrokerTasksBrokerIdGet**](RemoveBrokerTaskApi.md#ClustersClusterIdRemoveBrokerTasksBrokerIdGet) | **Get** /clusters/{cluster_id}/remove-broker-tasks/{broker_id} | Get Remove Broker Task
[**ClustersClusterIdRemoveBrokerTasksGet**](RemoveBrokerTaskApi.md#ClustersClusterIdRemoveBrokerTasksGet) | **Get** /clusters/{cluster_id}/remove-broker-tasks | List Remove Broker Tasks



## ClustersClusterIdRemoveBrokerTasksBrokerIdGet

> RemoveBrokerTaskData ClustersClusterIdRemoveBrokerTasksBrokerIdGet(ctx, clusterId, brokerId)

Get Remove Broker Task

Return the remove broker task for the specified ``broker_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 

### Return type

[**RemoveBrokerTaskData**](RemoveBrokerTaskData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdRemoveBrokerTasksGet

> RemoveBrokerTaskDataList ClustersClusterIdRemoveBrokerTasksGet(ctx, clusterId)

List Remove Broker Tasks

Returns a list of remove-broker-tasks for the specified Kafka cluster. ``/remove-broker-tasks`` is deprecated and may be removed in a future release. Use the new ``/tasks`` API instead.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**RemoveBrokerTaskDataList**](RemoveBrokerTaskDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

