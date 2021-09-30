# \BrokerV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdDelete**](BrokerV3Api.md#ClustersClusterIdBrokersBrokerIdDelete) | **Delete** /clusters/{cluster_id}/brokers/{broker_id} | Delete Broker
[**ClustersClusterIdBrokersBrokerIdGet**](BrokerV3Api.md#ClustersClusterIdBrokersBrokerIdGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id} | Get Broker
[**ClustersClusterIdBrokersBrokerIdPartitionReplicasGet**](BrokerV3Api.md#ClustersClusterIdBrokersBrokerIdPartitionReplicasGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/partition-replicas | Search Replicas by Broker
[**ClustersClusterIdBrokersGet**](BrokerV3Api.md#ClustersClusterIdBrokersGet) | **Get** /clusters/{cluster_id}/brokers | List Brokers



## ClustersClusterIdBrokersBrokerIdDelete

> BrokerRemovalData ClustersClusterIdBrokersBrokerIdDelete(ctx, clusterId, brokerId, optional)

Delete Broker

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Deletes the broker that is specified by ``broker_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
 **optional** | ***ClustersClusterIdBrokersBrokerIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokersBrokerIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shouldShutdown** | **optional.Bool**| To shutdown the broker or not, Default: true | 

### Return type

[**BrokerRemovalData**](BrokerRemovalData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdGet

> BrokerData ClustersClusterIdBrokersBrokerIdGet(ctx, clusterId, brokerId)

Get Broker

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the broker specified by ``broker_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

### Return type

[**BrokerData**](BrokerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdPartitionReplicasGet

> ReplicaDataList ClustersClusterIdBrokersBrokerIdPartitionReplicasGet(ctx, clusterId, brokerId)

Search Replicas by Broker

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the list of replicas assigned to the specified broker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

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


## ClustersClusterIdBrokersGet

> BrokerDataList ClustersClusterIdBrokersGet(ctx, clusterId)

List Brokers

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return a list of brokers that belong to the specified Kafka cluster.

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
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

