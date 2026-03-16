# \ShareGroupV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaShareGroup**](ShareGroupV3Api.md#GetKafkaShareGroup) | **Get** /clusters/{cluster_id}/share-groups/{group_id} | Get Share Group
[**GetKafkaShareGroupConsumer**](ShareGroupV3Api.md#GetKafkaShareGroupConsumer) | **Get** /clusters/{cluster_id}/share-groups/{group_id}/consumers/{consumer_id} | Get Share Group Consumer
[**ListKafkaShareGroupConsumerAssignments**](ShareGroupV3Api.md#ListKafkaShareGroupConsumerAssignments) | **Get** /clusters/{cluster_id}/share-groups/{group_id}/consumers/{consumer_id}/assignments | List Share Group Consumer Assignments
[**ListKafkaShareGroupConsumers**](ShareGroupV3Api.md#ListKafkaShareGroupConsumers) | **Get** /clusters/{cluster_id}/share-groups/{group_id}/consumers | List Share Group Consumers
[**ListKafkaShareGroups**](ShareGroupV3Api.md#ListKafkaShareGroups) | **Get** /clusters/{cluster_id}/share-groups | List Share Groups



## GetKafkaShareGroup

> ShareGroupData GetKafkaShareGroup(ctx, clusterId, groupId)

Get Share Group

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the share group specified by the ``group_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**groupId** | **string**| The group ID. | 

### Return type

[**ShareGroupData**](ShareGroupData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaShareGroupConsumer

> ShareGroupConsumerData GetKafkaShareGroupConsumer(ctx, clusterId, groupId, consumerId)

Get Share Group Consumer

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the consumer specified by the ``consumer_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**groupId** | **string**| The group ID. | 
**consumerId** | **string**| The consumer ID. | 

### Return type

[**ShareGroupConsumerData**](ShareGroupConsumerData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroupConsumerAssignments

> ShareGroupConsumerAssignmentDataList ListKafkaShareGroupConsumerAssignments(ctx, clusterId, groupId, consumerId)

List Share Group Consumer Assignments

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the consumer assignments specified by the ``consumer_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**groupId** | **string**| The group ID. | 
**consumerId** | **string**| The consumer ID. | 

### Return type

[**ShareGroupConsumerAssignmentDataList**](ShareGroupConsumerAssignmentDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroupConsumers

> ShareGroupConsumerDataList ListKafkaShareGroupConsumers(ctx, clusterId, groupId)

List Share Group Consumers

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return a list of consumers that belong to the specified share group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**groupId** | **string**| The group ID. | 

### Return type

[**ShareGroupConsumerDataList**](ShareGroupConsumerDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaShareGroups

> ShareGroupDataList ListKafkaShareGroups(ctx, clusterId)

List Share Groups

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return the list of share groups that belong to the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ShareGroupDataList**](ShareGroupDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

