# \BrokerApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdGet**](BrokerApi.md#ClustersClusterIdBrokersBrokerIdGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id} | Get Broker
[**ClustersClusterIdBrokersBrokerIdPartitionReplicasGet**](BrokerApi.md#ClustersClusterIdBrokersBrokerIdPartitionReplicasGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/partition-replicas | Search Replicas by Broker
[**ClustersClusterIdBrokersGet**](BrokerApi.md#ClustersClusterIdBrokersGet) | **Get** /clusters/{cluster_id}/brokers | List Brokers



## ClustersClusterIdBrokersBrokerIdGet

> BrokerData ClustersClusterIdBrokersBrokerIdGet(ctx, clusterId, brokerId)

Get Broker

Returns the broker specified by ``broker_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 

### Return type

[**BrokerData**](BrokerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdPartitionReplicasGet

> ReplicaDataList ClustersClusterIdBrokersBrokerIdPartitionReplicasGet(ctx, clusterId, brokerId)

Search Replicas by Broker

Returns the list of replicas assigned to the specified broker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 

### Return type

[**ReplicaDataList**](ReplicaDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersGet

> BrokerDataList ClustersClusterIdBrokersGet(ctx, clusterId)

List Brokers

Return a list of brokers that belong to the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**BrokerDataList**](BrokerDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

