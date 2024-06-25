# \RecordsV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProduceRecord**](RecordsV3Api.md#ProduceRecord) | **Post** /clusters/{cluster_id}/topics/{topic_name}/records | Produce Records



## ProduceRecord

> ProduceResponse ProduceRecord(ctx, clusterId, topicName, optional)

Produce Records

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Produce records to the given topic, returning delivery reports for each record produced. This API can be used in streaming mode by setting \"Transfer-Encoding: chunked\" header. For as long as the connection is kept open, the server will keep accepting records. For each record sent to the server, the server will asynchronously send back a delivery report, in the same order. Records are streamed to and from the server as Concatenated JSON. Errors are reported per record. The HTTP status code will be HTTP 200 OK as long as the connection is successfully established.  Note that the cluster_id is validated only when running in Confluent Cloud.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**topicName** | **string**| The topic name. | 
 **optional** | ***ProduceRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProduceRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **produceRequest** | [**optional.Interface of ProduceRequest**](ProduceRequest.md)| A single record to be produced to Kafka. To produce multiple records in the same request, simply concatenate the records. The delivery reports are concatenated in the same order as the records are sent. | 

### Return type

[**ProduceResponse**](ProduceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

