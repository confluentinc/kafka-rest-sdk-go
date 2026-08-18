# \ReplicationFactorChangeV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsReplicationFactorChangesGet**](ReplicationFactorChangeV3Api.md#ClustersClusterIdTopicsReplicationFactorChangesGet) | **Get** /clusters/{cluster_id}/topics/-/replication-factor-changes | List Replication Factor Changes
[**ClustersClusterIdTopicsReplicationFactorChangesPatch**](ReplicationFactorChangeV3Api.md#ClustersClusterIdTopicsReplicationFactorChangesPatch) | **Patch** /clusters/{cluster_id}/topics/-/replication-factor-changes | Change Replication Factor
[**ClustersClusterIdTopicsReplicationFactorChangescancelPatch**](ReplicationFactorChangeV3Api.md#ClustersClusterIdTopicsReplicationFactorChangescancelPatch) | **Patch** /clusters/{cluster_id}/topics/-/replication-factor-changes:cancel | Cancel Replication Factor Change



## ClustersClusterIdTopicsReplicationFactorChangesGet

> ReplicationFactorChangeDataList ClustersClusterIdTopicsReplicationFactorChangesGet(ctx, clusterId)

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


## ClustersClusterIdTopicsReplicationFactorChangesPatch

> ClustersClusterIdTopicsReplicationFactorChangesPatch(ctx, clusterId, optional)

Change Replication Factor

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Submit a Replication Factor change for one or more topics in the cluster specified with ``cluster_id``. The change is applied asynchronously; poll the List Replication Factor Changes API to follow progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdTopicsReplicationFactorChangesPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsReplicationFactorChangesPatchOpts struct


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


## ClustersClusterIdTopicsReplicationFactorChangescancelPatch

> ReplicationFactorChangeCancellationDataList ClustersClusterIdTopicsReplicationFactorChangescancelPatch(ctx, clusterId, optional)

Cancel Replication Factor Change

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Cancel the in-flight Replication Factor Change for the topics specified in the request body, in the cluster specified with ``cluster_id``. Cancellation does not roll back replicas already applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdTopicsReplicationFactorChangescancelPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsReplicationFactorChangescancelPatchOpts struct


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

