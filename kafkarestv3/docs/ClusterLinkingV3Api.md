# \ClusterLinkingV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaLink**](ClusterLinkingV3Api.md#CreateKafkaLink) | **Post** /clusters/{cluster_id}/links | Create a cluster link
[**CreateKafkaMirrorTopic**](ClusterLinkingV3Api.md#CreateKafkaMirrorTopic) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors | Create a mirror topic
[**DeleteKafkaLink**](ClusterLinkingV3Api.md#DeleteKafkaLink) | **Delete** /clusters/{cluster_id}/links/{link_name} | Delete the cluster link
[**DeleteKafkaLinkConfig**](ClusterLinkingV3Api.md#DeleteKafkaLinkConfig) | **Delete** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Reset the given config to default value
[**GetKafkaLink**](ClusterLinkingV3Api.md#GetKafkaLink) | **Get** /clusters/{cluster_id}/links/{link_name} | Describe the cluster link
[**GetKafkaLinkConfigs**](ClusterLinkingV3Api.md#GetKafkaLinkConfigs) | **Get** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Describe the config under the cluster link
[**ListKafkaLinkConfigs**](ClusterLinkingV3Api.md#ListKafkaLinkConfigs) | **Get** /clusters/{cluster_id}/links/{link_name}/configs | List all configs of the cluster link
[**ListKafkaLinks**](ClusterLinkingV3Api.md#ListKafkaLinks) | **Get** /clusters/{cluster_id}/links | List all cluster links in the dest cluster
[**ListKafkaMirrorTopics**](ClusterLinkingV3Api.md#ListKafkaMirrorTopics) | **Get** /clusters/{cluster_id}/links/-/mirrors | List mirror topics
[**ListKafkaMirrorTopicsUnderLink**](ClusterLinkingV3Api.md#ListKafkaMirrorTopicsUnderLink) | **Get** /clusters/{cluster_id}/links/{link_name}/mirrors | List mirror topics
[**ReadKafkaMirrorTopic**](ClusterLinkingV3Api.md#ReadKafkaMirrorTopic) | **Get** /clusters/{cluster_id}/links/{link_name}/mirrors/{mirror_topic_name} | Describe the mirror topic
[**UpdateKafkaLinkConfig**](ClusterLinkingV3Api.md#UpdateKafkaLinkConfig) | **Put** /clusters/{cluster_id}/links/{link_name}/configs/{config_name} | Alter the config under the cluster link
[**UpdateKafkaLinkConfigBatch**](ClusterLinkingV3Api.md#UpdateKafkaLinkConfigBatch) | **Post** /clusters/{cluster_id}/links/{link_name}/configs:alter | Batch Alter Topic Configs
[**UpdateKafkaMirrorTopicsFailover**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsFailover) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors:failover | Failover the mirror topics
[**UpdateKafkaMirrorTopicsPause**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsPause) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors:pause | Pause the mirror topics
[**UpdateKafkaMirrorTopicsPromote**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsPromote) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors:promote | Promote the mirror topics
[**UpdateKafkaMirrorTopicsResume**](ClusterLinkingV3Api.md#UpdateKafkaMirrorTopicsResume) | **Post** /clusters/{cluster_id}/links/{link_name}/mirrors:resume | Resume the mirror topics



## CreateKafkaLink

> CreateKafkaLink(ctx, clusterId, linkName, optional)

Create a cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***CreateKafkaLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **validateLink** | **optional.Bool**| To synchronously validate that the source cluster ID is expected and the dest cluster has the permission to read topics in the source cluster. Default: true | 
 **createLinkRequestData** | [**optional.Interface of CreateLinkRequestData**](CreateLinkRequestData.md)| Create a cluster link | 

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


## CreateKafkaMirrorTopic

> CreateKafkaMirrorTopic(ctx, clusterId, linkName, optional)

Create a mirror topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create a topic in the destination cluster mirroring a topic in the source cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***CreateKafkaMirrorTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaMirrorTopicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMirrorTopicRequestData** | [**optional.Interface of CreateMirrorTopicRequestData**](CreateMirrorTopicRequestData.md)| Name and configs of the topics mirroring from and mirroring to | 

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


## DeleteKafkaLink

> DeleteKafkaLink(ctx, clusterId, linkName, optional)

Delete the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***DeleteKafkaLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteKafkaLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Force the action. Default: false | 
 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 

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


## DeleteKafkaLinkConfig

> DeleteKafkaLinkConfig(ctx, clusterId, linkName, configName)

Reset the given config to default value

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

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


## GetKafkaLink

> ListLinksResponseData GetKafkaLink(ctx, clusterId, linkName)

Describe the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  ``link_id`` in ``ListLinksResponseData`` is deprecated and may be removed in a future release. Use the new ``cluster_link_id`` instead.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinksResponseData**](ListLinksResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaLinkConfigs

> ListLinkConfigsResponseData GetKafkaLinkConfigs(ctx, clusterId, linkName, configName)

Describe the config under the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 

### Return type

[**ListLinkConfigsResponseData**](ListLinkConfigsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaLinkConfigs

> ListLinkConfigsResponseDataList ListKafkaLinkConfigs(ctx, clusterId, linkName)

List all configs of the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 

### Return type

[**ListLinkConfigsResponseDataList**](ListLinkConfigsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaLinks

> ListLinksResponseDataList ListKafkaLinks(ctx, clusterId)

List all cluster links in the dest cluster

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  ``link_id`` in ``ListLinksResponseData`` is deprecated and may be removed in a future release. Use the new ``cluster_link_id`` instead.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ListLinksResponseDataList**](ListLinksResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaMirrorTopics

> ListMirrorTopicsResponseDataList ListKafkaMirrorTopics(ctx, clusterId, optional)

List mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  List all mirror topics in the cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ListKafkaMirrorTopicsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKafkaMirrorTopicsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaMirrorTopicsUnderLink

> ListMirrorTopicsResponseDataList ListKafkaMirrorTopicsUnderLink(ctx, clusterId, linkName, optional)

List mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  List all mirror topics under the link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***ListKafkaMirrorTopicsUnderLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKafkaMirrorTopicsUnderLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mirrorStatus** | [**optional.Interface of MirrorTopicStatus**](.md)| The status of the mirror topic. If not specified, all mirror topics will be returned. | 

### Return type

[**ListMirrorTopicsResponseDataList**](ListMirrorTopicsResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKafkaMirrorTopic

> ListMirrorTopicsResponseData ReadKafkaMirrorTopic(ctx, clusterId, linkName, mirrorTopicName)

Describe the mirror topic

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**mirrorTopicName** | **string**| Cluster Linking mirror topic name | 

### Return type

[**ListMirrorTopicsResponseData**](ListMirrorTopicsResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaLinkConfig

> UpdateKafkaLinkConfig(ctx, clusterId, linkName, configName, optional)

Alter the config under the cluster link

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
**configName** | **string**| The link config name | 
 **optional** | ***UpdateKafkaLinkConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaLinkConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateLinkConfigRequestData** | [**optional.Interface of UpdateLinkConfigRequestData**](UpdateLinkConfigRequestData.md)| Link config value to update | 

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


## UpdateKafkaLinkConfigBatch

> UpdateKafkaLinkConfigBatch(ctx, clusterId, linkName, optional)

Batch Alter Topic Configs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Batch Alter Link Configs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaLinkConfigBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaLinkConfigBatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **alterConfigBatchRequestData** | [**optional.Interface of AlterConfigBatchRequestData**](AlterConfigBatchRequestData.md)|  | 

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


## UpdateKafkaMirrorTopicsFailover

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsFailover(ctx, clusterId, linkName, optional)

Failover the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaMirrorTopicsFailoverOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaMirrorTopicsFailoverOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsPause

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsPause(ctx, clusterId, linkName, optional)

Pause the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaMirrorTopicsPauseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaMirrorTopicsPauseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsPromote

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsPromote(ctx, clusterId, linkName, optional)

Promote the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaMirrorTopicsPromoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaMirrorTopicsPromoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKafkaMirrorTopicsResume

> AlterMirrorStatusResponseDataList UpdateKafkaMirrorTopicsResume(ctx, clusterId, linkName, optional)

Resume the mirror topics

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**linkName** | **string**| The link name | 
 **optional** | ***UpdateKafkaMirrorTopicsResumeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKafkaMirrorTopicsResumeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateOnly** | **optional.Bool**| To validate the action can be performed successfully or not. Default: false | 
 **alterMirrorsRequestData** | [**optional.Interface of AlterMirrorsRequestData**](AlterMirrorsRequestData.md)| Name of the topics to apply the changes | 

### Return type

[**AlterMirrorStatusResponseDataList**](AlterMirrorStatusResponseDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

