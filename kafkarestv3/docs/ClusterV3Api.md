# \ClusterV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaCluster**](ClusterV3Api.md#GetKafkaCluster) | **Get** /clusters/{cluster_id} | Get Cluster
[**ListKafkaClusters**](ClusterV3Api.md#ListKafkaClusters) | **Get** /clusters | List Clusters



## GetKafkaCluster

> ClusterData GetKafkaCluster(ctx, clusterId)

Get Cluster

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the Kafka cluster with the specified ``cluster_id``.

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
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaClusters

> ClusterDataList ListKafkaClusters(ctx, )

List Clusters

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  'Return a list of known Kafka clusters. Currently both Kafka and Kafka REST Proxy are only aware of the Kafka cluster pointed at by the ``bootstrap.servers`` configuration. Therefore only one Kafka cluster will be returned in the response.'

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ClusterDataList**](ClusterDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

