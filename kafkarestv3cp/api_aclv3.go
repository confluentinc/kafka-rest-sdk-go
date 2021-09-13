// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * REST Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Contact: kafka-clients-proxy-team@confluent.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3cp

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ACLV3Api interface {

	/*
	 * CreateKafkaV3Acls Create ACLs
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Creates an ACL.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param optional nil or *CreateKafkaV3AclsOpts - Optional Parameters:
	 * @param "CreateAclRequestData" (optional.Interface of CreateAclRequestData) -  The ACL creation request.
	 */
	CreateKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *CreateKafkaV3AclsOpts) (*_nethttp.Response, error)

	/*
	 * DeleteKafkaV3Acls Delete ACLs
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Deletes the list of ACLs that matches the search criteria.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param optional nil or *DeleteKafkaV3AclsOpts - Optional Parameters:
	 * @param "ResourceType" (optional.Interface of AclResourceType) -  The ACL resource type.
	 * @param "ResourceName" (optional.String) -  The ACL resource name.
	 * @param "PatternType" (optional.String) -  The ACL pattern type.
	 * @param "Principal" (optional.String) -  The ACL principal.
	 * @param "Host" (optional.String) -  The ACL host.
	 * @param "Operation" (optional.String) -  The ACL operation.
	 * @param "Permission" (optional.String) -  The ACL permission.
	 * @return InlineResponse200
	 */
	DeleteKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *DeleteKafkaV3AclsOpts) (InlineResponse200, *_nethttp.Response, error)

	/*
	 * GetKafkaV3Acls Search ACLs
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of ACLs that match the search criteria.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param optional nil or *GetKafkaV3AclsOpts - Optional Parameters:
	 * @param "ResourceType" (optional.Interface of AclResourceType) -  The ACL resource type.
	 * @param "ResourceName" (optional.String) -  The ACL resource name.
	 * @param "PatternType" (optional.String) -  The ACL pattern type.
	 * @param "Principal" (optional.String) -  The ACL principal.
	 * @param "Host" (optional.String) -  The ACL host.
	 * @param "Operation" (optional.String) -  The ACL operation.
	 * @param "Permission" (optional.String) -  The ACL permission.
	 * @return AclDataList
	 */
	GetKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *GetKafkaV3AclsOpts) (AclDataList, *_nethttp.Response, error)
}

// ACLV3ApiService ACLV3Api service
type ACLV3ApiService service

// CreateKafkaV3AclsOpts Optional parameters for the method 'CreateKafkaV3Acls'
type CreateKafkaV3AclsOpts struct {
	CreateAclRequestData optional.Interface
}

/*
 * CreateKafkaV3Acls Create ACLs
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Creates an ACL.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param optional nil or *CreateKafkaV3AclsOpts - Optional Parameters:
 * @param "CreateAclRequestData" (optional.Interface of CreateAclRequestData) -  The ACL creation request.
 */
func (a *ACLV3ApiService) CreateKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *CreateKafkaV3AclsOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/acls"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.CreateAclRequestData.IsSet() {
		localVarOptionalCreateAclRequestData, localVarOptionalCreateAclRequestDataok := localVarOptionals.CreateAclRequestData.Value().(CreateAclRequestData)
		if !localVarOptionalCreateAclRequestDataok {
			return nil, reportError("createAclRequestData should be CreateAclRequestData")
		}
		localVarPostBody = &localVarOptionalCreateAclRequestData
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// DeleteKafkaV3AclsOpts Optional parameters for the method 'DeleteKafkaV3Acls'
type DeleteKafkaV3AclsOpts struct {
	ResourceType optional.Interface
	ResourceName optional.String
	PatternType  optional.String
	Principal    optional.String
	Host         optional.String
	Operation    optional.String
	Permission   optional.String
}

/*
 * DeleteKafkaV3Acls Delete ACLs
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Deletes the list of ACLs that matches the search criteria.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param optional nil or *DeleteKafkaV3AclsOpts - Optional Parameters:
 * @param "ResourceType" (optional.Interface of AclResourceType) -  The ACL resource type.
 * @param "ResourceName" (optional.String) -  The ACL resource name.
 * @param "PatternType" (optional.String) -  The ACL pattern type.
 * @param "Principal" (optional.String) -  The ACL principal.
 * @param "Host" (optional.String) -  The ACL host.
 * @param "Operation" (optional.String) -  The ACL operation.
 * @param "Permission" (optional.String) -  The ACL permission.
 * @return InlineResponse200
 */
func (a *ACLV3ApiService) DeleteKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *DeleteKafkaV3AclsOpts) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/acls"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resource_type", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceName.IsSet() {
		localVarQueryParams.Add("resource_name", parameterToString(localVarOptionals.ResourceName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PatternType.IsSet() {
		localVarQueryParams.Add("pattern_type", parameterToString(localVarOptionals.PatternType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Principal.IsSet() {
		localVarQueryParams.Add("principal", parameterToString(localVarOptionals.Principal.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Host.IsSet() {
		localVarQueryParams.Add("host", parameterToString(localVarOptionals.Host.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Operation.IsSet() {
		localVarQueryParams.Add("operation", parameterToString(localVarOptionals.Operation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Permission.IsSet() {
		localVarQueryParams.Add("permission", parameterToString(localVarOptionals.Permission.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetKafkaV3AclsOpts Optional parameters for the method 'GetKafkaV3Acls'
type GetKafkaV3AclsOpts struct {
	ResourceType optional.Interface
	ResourceName optional.String
	PatternType  optional.String
	Principal    optional.String
	Host         optional.String
	Operation    optional.String
	Permission   optional.String
}

/*
 * GetKafkaV3Acls Search ACLs
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of ACLs that match the search criteria.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param optional nil or *GetKafkaV3AclsOpts - Optional Parameters:
 * @param "ResourceType" (optional.Interface of AclResourceType) -  The ACL resource type.
 * @param "ResourceName" (optional.String) -  The ACL resource name.
 * @param "PatternType" (optional.String) -  The ACL pattern type.
 * @param "Principal" (optional.String) -  The ACL principal.
 * @param "Host" (optional.String) -  The ACL host.
 * @param "Operation" (optional.String) -  The ACL operation.
 * @param "Permission" (optional.String) -  The ACL permission.
 * @return AclDataList
 */
func (a *ACLV3ApiService) GetKafkaV3Acls(ctx _context.Context, clusterId string, localVarOptionals *GetKafkaV3AclsOpts) (AclDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AclDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/acls"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resource_type", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceName.IsSet() {
		localVarQueryParams.Add("resource_name", parameterToString(localVarOptionals.ResourceName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PatternType.IsSet() {
		localVarQueryParams.Add("pattern_type", parameterToString(localVarOptionals.PatternType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Principal.IsSet() {
		localVarQueryParams.Add("principal", parameterToString(localVarOptionals.Principal.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Host.IsSet() {
		localVarQueryParams.Add("host", parameterToString(localVarOptionals.Host.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Operation.IsSet() {
		localVarQueryParams.Add("operation", parameterToString(localVarOptionals.Operation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Permission.IsSet() {
		localVarQueryParams.Add("permission", parameterToString(localVarOptionals.Permission.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
