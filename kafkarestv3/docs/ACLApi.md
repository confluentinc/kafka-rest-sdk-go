# \ACLApi

All URIs are relative to *http://localhost:8082/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterIdAclsDelete**](ACLApi.md#ClustersClusterIdAclsDelete) | **Delete** /clusters/{cluster_id}/acls | Delete ACLs
[**ClustersClusterIdAclsGet**](ACLApi.md#ClustersClusterIdAclsGet) | **Get** /clusters/{cluster_id}/acls | Search ACLs
[**ClustersClusterIdAclsPost**](ACLApi.md#ClustersClusterIdAclsPost) | **Post** /clusters/{cluster_id}/acls | Create ACLs



## ClustersClusterIdAclsDelete

> InlineResponse200 ClustersClusterIdAclsDelete(ctx, clusterId, optional)

Delete ACLs

Deletes the list of ACLs that matches the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdAclsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdAclsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**optional.Interface of AclResourceType**](.md)| The ACL resource type. | 
 **resourceName** | **optional.String**| The ACL resource name. | 
 **patternType** | [**optional.Interface of AclPatternType**](.md)| The ACL pattern type. | 
 **principal** | **optional.String**| The ACL principal. | 
 **host** | **optional.String**| The ACL host. | 
 **operation** | [**optional.Interface of AclOperation**](.md)| The ACL operation. | 
 **permission** | [**optional.Interface of AclPermission**](.md)| The ACL permission. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdAclsGet

> AclDataList ClustersClusterIdAclsGet(ctx, clusterId, optional)

Search ACLs

Returns a list of ACLs that match the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdAclsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdAclsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**optional.Interface of AclResourceType**](.md)| The ACL resource type. | 
 **resourceName** | **optional.String**| The ACL resource name. | 
 **patternType** | [**optional.Interface of AclPatternType**](.md)| The ACL pattern type. | 
 **principal** | **optional.String**| The ACL principal. | 
 **host** | **optional.String**| The ACL host. | 
 **operation** | [**optional.Interface of AclOperation**](.md)| The ACL operation. | 
 **permission** | [**optional.Interface of AclPermission**](.md)| The ACL permission. | 

### Return type

[**AclDataList**](AclDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersClusterIdAclsPost

> ClustersClusterIdAclsPost(ctx, clusterId, optional)

Create ACLs

Creates an ACL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***ClustersClusterIdAclsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ClustersClusterIdAclsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAclRequestData** | [**optional.Interface of CreateAclRequestData**](CreateAclRequestData.md)| The ACL creation request. | 

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

