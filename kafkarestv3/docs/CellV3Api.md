# \CellV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeCells**](CellV3Api.md#DescribeCells) | **Get** /clusters/{cluster_id}/cells | Get Cluster Cell Configuration



## DescribeCells

> CellConfigDataAllOf DescribeCells(ctx, clusterId)

Get Cluster Cell Configuration

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of cells that belong to the specified Kafka cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 

### Return type

[**CellConfigDataAllOf**](CellConfigData_allOf.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

