# \ReplicationFactorChangeV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdReplicationFactorChangesGet**](ReplicationFactorChangeV3Api.md#ClustersClusterIdReplicationFactorChangesGet) | **Get** /clusters/{cluster_id}/replication-factor-changes | List Replication Factor Changes
[**ClustersClusterIdReplicationFactorChangescancelPost**](ReplicationFactorChangeV3Api.md#ClustersClusterIdReplicationFactorChangescancelPost) | **Post** /clusters/{cluster_id}/replication-factor-changes:cancel | Cancel Replication Factor Change
[**ClustersClusterIdReplicationFactorChangescreatePost**](ReplicationFactorChangeV3Api.md#ClustersClusterIdReplicationFactorChangescreatePost) | **Post** /clusters/{cluster_id}/replication-factor-changes:create | Change Replication Factor



## ClustersClusterIdReplicationFactorChangesGet

> ReplicationFactorChangeDataList ClustersClusterIdReplicationFactorChangesGet(ctx, clusterId)

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


## ClustersClusterIdReplicationFactorChangescancelPost

> ReplicationFactorChangeCancelDataList ClustersClusterIdReplicationFactorChangescancelPost(ctx, clusterId, optional)

Cancel Replication Factor Change

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Cancel the in-flight Replication Factor Change for the topics specified in the request body, in the cluster specified with ``cluster_id``. Cancellation does not roll back replicas already applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdReplicationFactorChangescancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdReplicationFactorChangescancelPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelReplicationFactorChangeRequestData** | [**optional.Interface of CancelReplicationFactorChangeRequestData**](CancelReplicationFactorChangeRequestData.md)| Topic names whose Replication Factor Change should be canceled. | 

### Return type

[**ReplicationFactorChangeCancelDataList**](ReplicationFactorChangeCancelDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdReplicationFactorChangescreatePost

> ClustersClusterIdReplicationFactorChangescreatePost(ctx, clusterId, optional)

Change Replication Factor

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Submit a Replication Factor change for one or more topics in the cluster specified with ``cluster_id``. The change is applied asynchronously; poll the List Replication Factor Changes API to follow progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdReplicationFactorChangescreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdReplicationFactorChangescreatePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeReplicationFactorBatchRequestData** | [**optional.Interface of ChangeReplicationFactorBatchRequestData**](ChangeReplicationFactorBatchRequestData.md)| Replication Factor changes to submit, one entry per topic. | 

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

