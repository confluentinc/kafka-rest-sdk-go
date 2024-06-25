# \ConfigsV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokersBrokerIdConfigsGet**](ConfigsV3Api.md#ClustersClusterIdBrokersBrokerIdConfigsGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs | List Broker Configs
[**ClustersClusterIdBrokersBrokerIdConfigsNameDelete**](ConfigsV3Api.md#ClustersClusterIdBrokersBrokerIdConfigsNameDelete) | **Delete** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Reset Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNameGet**](ConfigsV3Api.md#ClustersClusterIdBrokersBrokerIdConfigsNameGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Get Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNamePut**](ConfigsV3Api.md#ClustersClusterIdBrokersBrokerIdConfigsNamePut) | **Put** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Update Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsalterPost**](ConfigsV3Api.md#ClustersClusterIdBrokersBrokerIdConfigsalterPost) | **Post** /clusters/{cluster_id}/brokers/{broker_id}/configs:alter | Batch Alter Broker Configs
[**ClustersClusterIdBrokersConfigsGet**](ConfigsV3Api.md#ClustersClusterIdBrokersConfigsGet) | **Get** /clusters/{cluster_id}/brokers/-/configs | List All Broker Configs
[**DeleteKafkaClusterConfig**](ConfigsV3Api.md#DeleteKafkaClusterConfig) | **Delete** /clusters/{cluster_id}/broker-configs/{name} | Reset Cluster Config
[**DeleteKafkaTopicConfig**](ConfigsV3Api.md#DeleteKafkaTopicConfig) | **Delete** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Reset Topic Config
[**GetKafkaClusterConfig**](ConfigsV3Api.md#GetKafkaClusterConfig) | **Get** /clusters/{cluster_id}/broker-configs/{name} | Get Cluster Config
[**GetKafkaTopicConfig**](ConfigsV3Api.md#GetKafkaTopicConfig) | **Get** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Get Topic Config
[**ListKafkaAllTopicConfigs**](ConfigsV3Api.md#ListKafkaAllTopicConfigs) | **Get** /clusters/{cluster_id}/topics/-/configs | Get All Topic Configs
[**ListKafkaClusterConfigs**](ConfigsV3Api.md#ListKafkaClusterConfigs) | **Get** /clusters/{cluster_id}/broker-configs | List Cluster Configs
[**ListKafkaDefaultTopicConfigs**](ConfigsV3Api.md#ListKafkaDefaultTopicConfigs) | **Get** /clusters/{cluster_id}/topics/{topic_name}/default-configs | List Default Topic Configs
[**ListKafkaTopicConfigs**](ConfigsV3Api.md#ListKafkaTopicConfigs) | **Get** /clusters/{cluster_id}/topics/{topic_name}/configs | List Topic Configs
[**UpdateKafkaClusterConfig**](ConfigsV3Api.md#UpdateKafkaClusterConfig) | **Put** /clusters/{cluster_id}/broker-configs/{name} | Update Cluster Config
[**UpdateKafkaClusterConfigs**](ConfigsV3Api.md#UpdateKafkaClusterConfigs) | **Post** /clusters/{cluster_id}/broker-configs:alter | Batch Alter Cluster Configs
[**UpdateKafkaTopicConfig**](ConfigsV3Api.md#UpdateKafkaTopicConfig) | **Put** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Update Topic Config
[**UpdateKafkaTopicConfigBatch**](ConfigsV3Api.md#UpdateKafkaTopicConfigBatch) | **Post** /clusters/{cluster_id}/topics/{topic_name}/configs:alter | Batch Alter Topic Configs



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


## DeleteKafkaClusterConfig

> DeleteKafkaClusterConfig(ctx, clusterId, name)

Reset Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Resets the configuration parameter specified by ``name`` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
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


## DeleteKafkaTopicConfig

> DeleteKafkaTopicConfig(ctx, clusterId, topicName, name)

Reset Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Resets the config with given `name` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
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


## GetKafkaClusterConfig

> ClusterConfigData GetKafkaClusterConfig(ctx, clusterId, name)

Get Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**ClusterConfigData**](ClusterConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaTopicConfig

> TopicConfigData GetKafkaTopicConfig(ctx, clusterId, topicName, name)

Get Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the config with the given `name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**TopicConfigData**](TopicConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaAllTopicConfigs

> TopicConfigDataList ListKafkaAllTopicConfigs(ctx, clusterId)

Get All Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns all topic configurations for topics hosted by the specified cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

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


## ListKafkaClusterConfigs

> ClusterConfigDataList ListKafkaClusterConfigs(ctx, clusterId)

List Cluster Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of configuration parameters for the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ClusterConfigDataList**](ClusterConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaDefaultTopicConfigs

> TopicConfigDataList ListKafkaDefaultTopicConfigs(ctx, clusterId, topicName)

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


## ListKafkaTopicConfigs

> TopicConfigDataList ListKafkaTopicConfigs(ctx, clusterId, topicName)

List Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of configs that belong to the specified topic.

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


## UpdateKafkaClusterConfig

> UpdateKafkaClusterConfig(ctx, clusterId, name, optional)

Update Cluster Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***UpdateKafkaClusterConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaClusterConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The cluster configuration parameter update request. | 

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


## UpdateKafkaClusterConfigs

> UpdateKafkaClusterConfigs(ctx, clusterId, optional)

Batch Alter Cluster Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates or deletes a set of Kafka cluster configuration parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***UpdateKafkaClusterConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaClusterConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter cluster configuration parameter batch request. | 

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


## UpdateKafkaTopicConfig

> UpdateKafkaTopicConfig(ctx, clusterId, topicName, name, optional)

Update Topic Config

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates the config with given `name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***UpdateKafkaTopicConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaTopicConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The topic configuration parameter update request. | 

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


## UpdateKafkaTopicConfigBatch

> UpdateKafkaTopicConfigBatch(ctx, clusterId, topicName, optional)

Batch Alter Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Updates or deletes a set of topic configs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***UpdateKafkaTopicConfigBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaTopicConfigBatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter topic configuration parameter batch request. | 

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

