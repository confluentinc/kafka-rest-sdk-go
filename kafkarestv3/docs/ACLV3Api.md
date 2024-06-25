# \ACLV3Api

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCreateKafkaAcls**](ACLV3Api.md#BatchCreateKafkaAcls) | **Post** /clusters/{cluster_id}/acls:batch | Batch Create ACLs
[**CreateKafkaAcls**](ACLV3Api.md#CreateKafkaAcls) | **Post** /clusters/{cluster_id}/acls | Create an ACL
[**DeleteKafkaAcls**](ACLV3Api.md#DeleteKafkaAcls) | **Delete** /clusters/{cluster_id}/acls | Delete ACLs
[**GetKafkaAcls**](ACLV3Api.md#GetKafkaAcls) | **Get** /clusters/{cluster_id}/acls | List ACLs



## BatchCreateKafkaAcls

> BatchCreateKafkaAcls(ctx, clusterId, optional)

Batch Create ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create ACLs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***BatchCreateKafkaAclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BatchCreateKafkaAclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAclRequestDataList** | [**optional.Interface of CreateAclRequestDataList**](CreateAclRequestDataList.md)| The batch ACL creation request. | 

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


## CreateKafkaAcls

> CreateKafkaAcls(ctx, clusterId, optional)

Create an ACL

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create an Apache Kafka ACL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***CreateKafkaAclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaAclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAclRequestData** | [**optional.Interface of CreateAclRequestData**](CreateAclRequestData.md)| The ACL creation request. | 

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


## DeleteKafkaAcls

> InlineResponse200 DeleteKafkaAcls(ctx, clusterId, resourceType, patternType, operation, permission, optional)

Delete ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Delete the list of Apache Kafka ACLs that matches the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
**resourceType** | [**AclResourceType**](.md)| The ACL resource type. | 
**patternType** | **string**| The ACL pattern type. | 
**operation** | **string**| The ACL operation. | 
**permission** | **string**| The ACL permission. | 
 **optional** | ***DeleteKafkaAclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteKafkaAclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **resourceName** | **optional.String**| The ACL resource name. | 
 **principal** | **optional.String**| The ACL principal. This is the Service Account name or user name. | 
 **host** | **optional.String**| The ACL host. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaAcls

> AclDataList GetKafkaAcls(ctx, clusterId, optional)

List ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Return a list of ACLs that match the search criteria. These are Apache Kafka ACLs, which differ from the Confluent Metadata Service (MDS) based, centralized ACLs created with the Confluent CLI. MDS has a separate API for ACLs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***GetKafkaAclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKafkaAclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**optional.Interface of AclResourceType**](.md)| The ACL resource type. | 
 **resourceName** | **optional.String**| The ACL resource name. | 
 **patternType** | **optional.String**| The ACL pattern type. | 
 **principal** | **optional.String**| The ACL principal. This is the Service Account name or user name. | 
 **host** | **optional.String**| The ACL host. | 
 **operation** | **optional.String**| The ACL operation. | 
 **permission** | **optional.String**| The ACL permission. | 

### Return type

[**AclDataList**](AclDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

