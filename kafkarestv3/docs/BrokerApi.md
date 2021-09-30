# \BrokerApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersdeletePost**](BrokerApi.md#ClustersClusterIdBrokersdeletePost) | **Post** /clusters/{cluster_id}/brokers:delete | Delete several brokers



## ClustersClusterIdBrokersdeletePost

> BrokerRemovalDataList ClustersClusterIdBrokersdeletePost(ctx, clusterId, optional)

Delete several brokers

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdBrokersdeletePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokersdeletePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shouldShutdown** | **optional.Bool**| To shutdown the broker or not, Default: true | 
 **removeBrokersRequestData** | [**optional.Interface of RemoveBrokersRequestData**](RemoveBrokersRequestData.md)| Broker ids to remove | 

### Return type

[**BrokerRemovalDataList**](BrokerRemovalDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

