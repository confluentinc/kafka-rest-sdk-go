# \ReplicationFactorChangeV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelKafkaReplicationFactorChange**](ReplicationFactorChangeV3Api.md#CancelKafkaReplicationFactorChange) | **Patch** /clusters/{cluster_id}/topics/-/replication-factor-changes:cancel | Cancel Replication Factor Change
[**ChangeKafkaReplicationFactor**](ReplicationFactorChangeV3Api.md#ChangeKafkaReplicationFactor) | **Patch** /clusters/{cluster_id}/topics/-/replication-factor-changes | Change Replication Factor
[**ListKafkaReplicationFactorChanges**](ReplicationFactorChangeV3Api.md#ListKafkaReplicationFactorChanges) | **Get** /clusters/{cluster_id}/topics/-/replication-factor-changes | List Replication Factor Changes



## CancelKafkaReplicationFactorChange

> ReplicationFactorChangeCancellationDataList CancelKafkaReplicationFactorChange(ctx, clusterId, optional)

Cancel Replication Factor Change

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Cancel the in-flight Replication Factor Change for the topics specified in the request body, in the cluster specified with ``cluster_id``. Cancellation does not roll back replicas already applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***CancelKafkaReplicationFactorChangeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelKafkaReplicationFactorChangeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelReplicationFactorChangeRequestData** | [**optional.Interface of CancelReplicationFactorChangeRequestData**](CancelReplicationFactorChangeRequestData.md)| Topic names whose Replication Factor Change should be canceled. | 

### Return type

[**ReplicationFactorChangeCancellationDataList**](ReplicationFactorChangeCancellationDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeKafkaReplicationFactor

> ChangeKafkaReplicationFactor(ctx, clusterId, optional)

Change Replication Factor

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Submit a Replication Factor change for one or more topics in the cluster specified with ``cluster_id``. The change is applied asynchronously; poll the List Replication Factor Changes API to follow progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ChangeKafkaReplicationFactorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ChangeKafkaReplicationFactorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeReplicationFactorRequestData** | [**optional.Interface of ChangeReplicationFactorRequestData**](ChangeReplicationFactorRequestData.md)| Replication Factor changes to submit, one entry per topic. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkaReplicationFactorChanges

> ReplicationFactorChangeDataList ListKafkaReplicationFactorChanges(ctx, clusterId)

List Replication Factor Changes

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return all tracked Replication Factor Changes in the cluster specified with ``cluster_id``.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**ReplicationFactorChangeDataList**](ReplicationFactorChangeDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

