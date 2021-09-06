# \ConfigsApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdBrokerConfigsGet**](ConfigsApi.md#ClustersClusterIdBrokerConfigsGet) | **Get** /clusters/{cluster_id}/broker-configs | List Cluster Configs
[**ClustersClusterIdBrokerConfigsNameDelete**](ConfigsApi.md#ClustersClusterIdBrokerConfigsNameDelete) | **Delete** /clusters/{cluster_id}/broker-configs/{name} | Reset Cluster Config
[**ClustersClusterIdBrokerConfigsNameGet**](ConfigsApi.md#ClustersClusterIdBrokerConfigsNameGet) | **Get** /clusters/{cluster_id}/broker-configs/{name} | Get Cluster Config
[**ClustersClusterIdBrokerConfigsNamePut**](ConfigsApi.md#ClustersClusterIdBrokerConfigsNamePut) | **Put** /clusters/{cluster_id}/broker-configs/{name} | Update Cluster Config
[**ClustersClusterIdBrokerConfigsalterPost**](ConfigsApi.md#ClustersClusterIdBrokerConfigsalterPost) | **Post** /clusters/{cluster_id}/broker-configs:alter | Batch Alter Cluster Configs
[**ClustersClusterIdBrokersBrokerIdConfigsGet**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs | List Broker Configs
[**ClustersClusterIdBrokersBrokerIdConfigsNameDelete**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNameDelete) | **Delete** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Reset Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNameGet**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNameGet) | **Get** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Get Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsNamePut**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsNamePut) | **Put** /clusters/{cluster_id}/brokers/{broker_id}/configs/{name} | Update Broker Config
[**ClustersClusterIdBrokersBrokerIdConfigsalterPost**](ConfigsApi.md#ClustersClusterIdBrokersBrokerIdConfigsalterPost) | **Post** /clusters/{cluster_id}/brokers/{broker_id}/configs:alter | Batch Alter Broker Configs
[**ClustersClusterIdTopicsTopicNameConfigsGet**](ConfigsApi.md#ClustersClusterIdTopicsTopicNameConfigsGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/configs | List Topic Configs
[**ClustersClusterIdTopicsTopicNameConfigsNameDelete**](ConfigsApi.md#ClustersClusterIdTopicsTopicNameConfigsNameDelete) | **Delete** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Reset Topic Config
[**ClustersClusterIdTopicsTopicNameConfigsNameGet**](ConfigsApi.md#ClustersClusterIdTopicsTopicNameConfigsNameGet) | **Get** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Get Topic Config
[**ClustersClusterIdTopicsTopicNameConfigsNamePut**](ConfigsApi.md#ClustersClusterIdTopicsTopicNameConfigsNamePut) | **Put** /clusters/{cluster_id}/topics/{topic_name}/configs/{name} | Update Topic Config
[**ClustersClusterIdTopicsTopicNameConfigsalterPost**](ConfigsApi.md#ClustersClusterIdTopicsTopicNameConfigsalterPost) | **Post** /clusters/{cluster_id}/topics/{topic_name}/configs:alter | Batch Alter Topic Configs



## ClustersClusterIdBrokerConfigsGet

> ClusterConfigDataList ClustersClusterIdBrokerConfigsGet(ctx, clusterId)

List Cluster Configs

Returns a list of configuration parameters for the specified Kafka cluster.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerConfigsNameDelete

> ClustersClusterIdBrokerConfigsNameDelete(ctx, clusterId, name)

Reset Cluster Config

Resets the configuration parameter specified by ``name`` to its default value.

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerConfigsNameGet

> ClusterConfigData ClustersClusterIdBrokerConfigsNameGet(ctx, clusterId, name)

Get Cluster Config

Returns the configuration parameter specified by ``name``.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerConfigsNamePut

> ClustersClusterIdBrokerConfigsNamePut(ctx, clusterId, name, optional)

Update Cluster Config

Updates the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***ClustersClusterIdBrokerConfigsNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokerConfigsNamePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The cluster configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokerConfigsalterPost

> ClustersClusterIdBrokerConfigsalterPost(ctx, clusterId, optional)

Batch Alter Cluster Configs

Updates or deletes a set of Kafka cluster configuration parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdBrokerConfigsalterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdBrokerConfigsalterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter cluster configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsGet

> BrokerConfigDataList ClustersClusterIdBrokersBrokerIdConfigsGet(ctx, clusterId, brokerId)

List Broker Configs

Return the list of configuration parameters that belong to the specified Kafka broker.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 

### Return type

[**BrokerConfigDataList**](BrokerConfigDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNameDelete

> ClustersClusterIdBrokersBrokerIdConfigsNameDelete(ctx, clusterId, brokerId, name)

Reset Broker Config

Resets the configuration parameter specified by ``name`` to its default value.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNameGet

> BrokerConfigData ClustersClusterIdBrokersBrokerIdConfigsNameGet(ctx, clusterId, brokerId, name)

Get Broker Config

Return the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 
**name** | **string**| The configuration parameter name. | 

### Return type

[**BrokerConfigData**](BrokerConfigData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsNamePut

> ClustersClusterIdBrokersBrokerIdConfigsNamePut(ctx, clusterId, brokerId, name, optional)

Update Broker Config

Updates the configuration parameter specified by ``name``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 
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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdBrokersBrokerIdConfigsalterPost

> ClustersClusterIdBrokersBrokerIdConfigsalterPost(ctx, clusterId, brokerId, optional)

Batch Alter Broker Configs

Updates or deletes a set of broker configuration parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**brokerId** | **int32**| The broker ID. | 
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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameConfigsGet

> TopicConfigDataList ClustersClusterIdTopicsTopicNameConfigsGet(ctx, clusterId, topicName)

List Topic Configs

Return the list of configs that belong to the specified topic.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameConfigsNameDelete

> ClustersClusterIdTopicsTopicNameConfigsNameDelete(ctx, clusterId, topicName, name)

Reset Topic Config

Resets the config with given `name` to its default value.

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameConfigsNameGet

> TopicConfigData ClustersClusterIdTopicsTopicNameConfigsNameGet(ctx, clusterId, topicName, name)

Get Topic Config

Return the config with the given `name`.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameConfigsNamePut

> ClustersClusterIdTopicsTopicNameConfigsNamePut(ctx, clusterId, topicName, name, optional)

Update Topic Config

Updates the config with given `name`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
**name** | **string**| The configuration parameter name. | 
 **optional** | ***ClustersClusterIdTopicsTopicNameConfigsNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsTopicNameConfigsNamePutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateConfigRequestData** | [**optional.Interface of UpdateConfigRequestData**](UpdateConfigRequestData.md)| The topic configuration parameter update request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdTopicsTopicNameConfigsalterPost

> ClustersClusterIdTopicsTopicNameConfigsalterPost(ctx, clusterId, topicName, optional)

Batch Alter Topic Configs

Updates or deletes a set of topic configs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***ClustersClusterIdTopicsTopicNameConfigsalterPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsTopicNameConfigsalterPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)| The alter topic configuration parameter batch request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

