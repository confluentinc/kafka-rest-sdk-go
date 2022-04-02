# \BrokerReplicaExclusionApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet**](BrokerReplicaExclusionApi.md#ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet) | **Get** /clusters/{cluster_id}/broker-replica-exclusions/{broker_id} | Get a Broker Replica Exclusions.
[**ClustersClusterIdBrokerReplicaExclusionsGet**](BrokerReplicaExclusionApi.md#ClustersClusterIdBrokerReplicaExclusionsGet) | **Get** /clusters/{cluster_id}/broker-replica-exclusions | Get all Broker Replica Exclusions.
[**ClustersClusterIdBrokerReplicaExclusionscreatePost**](BrokerReplicaExclusionApi.md#ClustersClusterIdBrokerReplicaExclusionscreatePost) | **Post** /clusters/{cluster_id}/broker-replica-exclusions:create | Create Broker Replica Exclusions
[**ClustersClusterIdBrokerReplicaExclusionsdeletePost**](BrokerReplicaExclusionApi.md#ClustersClusterIdBrokerReplicaExclusionsdeletePost) | **Post** /clusters/{cluster_id}/broker-replica-exclusions:delete | Delete Broker Replica Exclusions



## ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet

> BrokerReplicaExclusionData ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet(ctx, clusterId, brokerId)

Get a Broker Replica Exclusions.

Returns a Broker Replica Exclusions in the cluster specified with ``cluster_id`` and broker specified with ``broker_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 

### Return type

[**BrokerReplicaExclusionData**](BrokerReplicaExclusionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerReplicaExclusionsGet

> BrokerReplicaExclusionDataList ClustersClusterIdBrokerReplicaExclusionsGet(ctx, clusterId)

Get all Broker Replica Exclusions.

Returns all Broker Replica Exclusions in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**BrokerReplicaExclusionDataList**](BrokerReplicaExclusionDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerReplicaExclusionscreatePost

> AlterBrokerReplicaExclusionDataList ClustersClusterIdBrokerReplicaExclusionscreatePost(ctx, clusterId, optional)

Create Broker Replica Exclusions

Create Broker Replica Exclusions for brokers in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdBrokerReplicaExclusionscreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokerReplicaExclusionscreatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brokerReplicaExclusionBatchRequestData** | [**optional.Interface of BrokerReplicaExclusionBatchRequestData**](BrokerReplicaExclusionBatchRequestData.md)| Alter Broker Replica Exclusions. | 

### Return type

[**AlterBrokerReplicaExclusionDataList**](AlterBrokerReplicaExclusionDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerReplicaExclusionsdeletePost

> AlterBrokerReplicaExclusionDataList ClustersClusterIdBrokerReplicaExclusionsdeletePost(ctx, clusterId, optional)

Delete Broker Replica Exclusions

Delete Broker Replica Exclusions for brokers in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brokerReplicaExclusionBatchRequestData** | [**optional.Interface of BrokerReplicaExclusionBatchRequestData**](BrokerReplicaExclusionBatchRequestData.md)| Alter Broker Replica Exclusions. | 

### Return type

[**AlterBrokerReplicaExclusionDataList**](AlterBrokerReplicaExclusionDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

