# \RecordsApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdTopicsTopicNameRecordsPost**](RecordsApi.md#ClustersClusterIdTopicsTopicNameRecordsPost) | **Post** /clusters/{cluster_id}/topics/{topic_name}/records | Produce records to the given topic.



## ClustersClusterIdTopicsTopicNameRecordsPost

> ProduceResponse ClustersClusterIdTopicsTopicNameRecordsPost(ctx, clusterId, topicName, optional)

Produce records to the given topic.

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Produce records to the given topic, returning delivery reports for each             record produced. This API can be used in streaming mode by setting \"Transfer-Encoding:             chunked\" header. For as long as the connection is kept open, the server will             keep accepting records. For each record sent to the server, the server will             asynchronously send back a delivery report, in the same order. Records are             streamed to and from the server as Concatenated JSON. Errors are reported per             record. The HTTP status code will be HTTP 200 OK as long as the connection is successfully established.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***ClustersClusterIdTopicsTopicNameRecordsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdTopicsTopicNameRecordsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **produceRequest** | [**optional.Interface of ProduceRequest**](ProduceRequest.md)| A single record to be produced to Kafka. To produce multiple records on the same connection, simply concatenate all the records, e.g.: {\&quot;partition_id\&quot;:1}{\&quot;partition_id\&quot;:2}. Delivery reports will be concatenated on the same order as the records are sent. See examples for the options available. | 

### Return type

[**ProduceResponse**](ProduceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

