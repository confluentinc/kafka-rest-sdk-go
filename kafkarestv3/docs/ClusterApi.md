# \ClusterApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdGet**](ClusterApi.md#ClustersClusterIdGet) | **Get** /clusters/{cluster_id} | Get Cluster
[**ClustersGet**](ClusterApi.md#ClustersGet) | **Get** /clusters | List Clusters



## ClustersClusterIdGet

> ClusterData ClustersClusterIdGet(ctx, clusterId)

Get Cluster

Returns the Kafka cluster with the specified ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ClusterData**](ClusterData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersGet

> ClusterDataList ClustersGet(ctx, )

List Clusters

Returns a list of known Kafka clusters. Currently both Kafka and Kafka REST Proxy are only aware of the Kafka cluster pointed at by the ``bootstrap.servers`` configuration. Therefore only one Kafka cluster will be returned in the response.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ClusterDataList**](ClusterDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

