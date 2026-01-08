# \ReplicaV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdPartitionReplicasGet**](ReplicaV3Api.md#ClustersClusterIdBrokersBrokerIdPartitionReplicasGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/partition-replicas | Search Replicas by Broker



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

