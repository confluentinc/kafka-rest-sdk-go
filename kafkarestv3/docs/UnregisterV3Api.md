# \UnregisterV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdunregisterPost**](UnregisterV3Api.md#ClustersClusterIdBrokersBrokerIdunregisterPost) | **Post** /clusters/{cluster_id}/brokers/{broker_id}:unregister | Unregister a Broker



## ClustersClusterIdBrokersBrokerIdunregisterPost

> UnregisterBrokerData ClustersClusterIdBrokersBrokerIdunregisterPost(ctx, clusterId, brokerId)

Unregister a Broker

Unregister a broker from the cluster. This API is only supported for Kafka clusters running in KRaft mode. If run against a cluster running with ZooKeeper, a 400 response with an unsupported version error code will be returned (see BadRequestErrorResponse for more detail).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

### Return type

[**UnregisterBrokerData**](UnregisterBrokerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

