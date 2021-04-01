# \BrokerTaskApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdTasksGet**](BrokerTaskApi.md#ClustersClusterIdBrokersBrokerIdTasksGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/tasks | List Broker Tasks of a specific Broker
[**ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet**](BrokerTaskApi.md#ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/tasks/{task_type} | Get single Broker Task.
[**ClustersClusterIdBrokersTasksGet**](BrokerTaskApi.md#ClustersClusterIdBrokersTasksGet) | **Get** /clusters/{cluster_id}/brokers/-/tasks | List Broker Tasks
[**ClustersClusterIdBrokersTasksTaskTypeGet**](BrokerTaskApi.md#ClustersClusterIdBrokersTasksTaskTypeGet) | **Get** /clusters/{cluster_id}/brokers/-/tasks/{task_type} | List Broker Tasks of a specific TaskType



## ClustersClusterIdBrokersBrokerIdTasksGet

> BrokerTaskDataList ClustersClusterIdBrokersBrokerIdTasksGet(ctx, clusterId, brokerId)

List Broker Tasks of a specific Broker

Returns a list of all broker tasks for broker specified with ``broker_id`` in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

### Return type

[**BrokerTaskDataList**](BrokerTaskDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet

> BrokerTaskData ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet(ctx, clusterId, brokerId, taskType)

Get single Broker Task.

Returns a single Broker Task specified with ``task_type`` for broker specified with ``broker_id`` in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
**taskType** | [**BrokerTaskType**](.md)| The Kafka broker task type. | 

### Return type

[**BrokerTaskData**](BrokerTaskData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersTasksGet

> BrokerTaskDataList ClustersClusterIdBrokersTasksGet(ctx, clusterId)

List Broker Tasks

Returns a list of all tasks for all brokers in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**BrokerTaskDataList**](BrokerTaskDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersTasksTaskTypeGet

> BrokerTaskDataList ClustersClusterIdBrokersTasksTaskTypeGet(ctx, clusterId, taskType)

List Broker Tasks of a specific TaskType

Returns a list of all broker tasks of specified ``task_type`` in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**taskType** | [**BrokerTaskType**](.md)| The Kafka broker task type. | 

### Return type

[**BrokerTaskDataList**](BrokerTaskDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

