# \ConfigsApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdConfigsGet**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs | List Broker Configs
[**ClustersClusterIdBrokersBrokerIdConfigsNameDelete**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNameDelete) | **Delete** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Reset Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNameGet**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNameGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Get Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNamePut**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNamePut) | **Put** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Update Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsalterPost**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsalterPost) | **Post** /clusters/{cluster_id}/brokers/{broker_id}/configs:alter | Batch Alter Broker Configs
[**ClustersClusterIdBrokersConfigsGet**](ConfigsApi.md#ClustersClusterIdBrokersConfigsGet) | **Get** /clusters/{cluster_id}/brokers/-/configs | List All Broker Configs
[**ListKafkaV3DefaultTopicConfigs**](ConfigsApi.md#ListKafkaV3DefaultTopicConfigs) | **Get** /clusters/{cluster_id}/topics/{topic_name}/default-configs | List Default Topic Configs



## ClustersClusterIdBrokersBrokerIdConfigsGet

> BrokerConfigDataList ClustersClusterIdBrokersBrokerIdConfigsGet(ctx, clusterId, brokerId)

List Broker Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of configuration parameters that belong to the specified Kafka broker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 

### Return type

[**BrokerConfigDataList**](BrokerConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNameDelete

> ClustersClusterIdBrokersBrokerIdConfigsNameDelete(ctx, clusterId, brokerId, name)

Reset Broker Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Resets the configuration parameter specified by ``name`` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNameGet

> BrokerConfigData ClustersClusterIdBrokersBrokerIdConfigsNameGet(ctx, clusterId, brokerId, name)

Get Broker Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**BrokerConfigData**](BrokerConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNamePut

> ClustersClusterIdBrokersBrokerIdConfigsNamePut(ctx, clusterId, brokerId, name, optional)

Update Broker Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***ClustersClusterIdBrokersBrokerIdConfigsNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokersBrokerIdConfigsNamePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The broker configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsalterPost

> ClustersClusterIdBrokersBrokerIdConfigsalterPost(ctx, clusterId, brokerId, optional)

Batch Alter Broker Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates or deletes a set of broker configuration parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The Kafka broker ID. | 
 **optional** | ***ClustersClusterIdBrokersBrokerIdConfigsalterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokersBrokerIdConfigsalterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter broker configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersConfigsGet

> BrokerConfigDataList ClustersClusterIdBrokersConfigsGet(ctx, clusterId)

List All Broker Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of configuration parameters for all the brokers in the given Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**BrokerConfigDataList**](BrokerConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaV3DefaultTopicConfigs

> TopicConfigDataList ListKafkaV3DefaultTopicConfigs(ctx, clusterId, topicName)

List Default Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 

### Return type

[**TopicConfigDataList**](TopicConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

