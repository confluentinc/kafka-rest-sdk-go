# \BalancerStatusApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBalancerAnyUnevenLoadGet**](BalancerStatusApi.md#ClustersClusterIdBalancerAnyUnevenLoadGet) | **Get** /clusters/{cluster_id}/balancer/any-uneven-load | Get AnyUnevenLoad status
[**ClustersClusterIdBalancerGet**](BalancerStatusApi.md#ClustersClusterIdBalancerGet) | **Get** /clusters/{cluster_id}/balancer | Get status of the balancer



## ClustersClusterIdBalancerAnyUnevenLoadGet

> AnyUnevenLoadData ClustersClusterIdBalancerAnyUnevenLoadGet(ctx, clusterId)

Get AnyUnevenLoad status

Returns status of the AnyUnevenLoad for the cluster specified by ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**AnyUnevenLoadData**](AnyUnevenLoadData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBalancerGet

> BalancerStatusData ClustersClusterIdBalancerGet(ctx, clusterId)

Get status of the balancer

Returns status about the balancer component for the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**BalancerStatusData**](BalancerStatusData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

