/*
 * REST Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Contact: kafka-clients-proxy-team@confluent.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

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

type BrokerReplicaExclusionApi interface {

	/*
	 * ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet Get a Broker Replica Exclusions.
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a Broker Replica Exclusions in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60; and broker specified with &#x60;&#x60;broker_id&#x60;&#x60;.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param brokerId The Kafka broker ID.
	 * @return BrokerReplicaExclusionData
	 */
	ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet(ctx _context.Context, clusterId string, brokerId int32) (BrokerReplicaExclusionData, *_nethttp.Response, error)

	/*
	 * ClustersClusterIdBrokerReplicaExclusionsGet Get all Broker Replica Exclusions.
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns all Broker Replica Exclusions in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @return BrokerReplicaExclusionDataList
	 */
	ClustersClusterIdBrokerReplicaExclusionsGet(ctx _context.Context, clusterId string) (BrokerReplicaExclusionDataList, *_nethttp.Response, error)

	/*
	 * ClustersClusterIdBrokerReplicaExclusionscreatePost Create Broker Replica Exclusions
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create Broker Replica Exclusions for brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param optional nil or *ClustersClusterIdBrokerReplicaExclusionscreatePostOpts - Optional Parameters:
	 * @param "BrokerReplicaExclusionBatchRequestData" (optional.Interface of BrokerReplicaExclusionBatchRequestData) -  Alter Broker Replica Exclusions.
	 * @return AlterBrokerReplicaExclusionDataList
	 */
	ClustersClusterIdBrokerReplicaExclusionscreatePost(ctx _context.Context, clusterId string, localVarOptionals *ClustersClusterIdBrokerReplicaExclusionscreatePostOpts) (AlterBrokerReplicaExclusionDataList, *_nethttp.Response, error)

	/*
	 * ClustersClusterIdBrokerReplicaExclusionsdeletePost Delete Broker Replica Exclusions
	 *
	 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Delete Broker Replica Exclusions for brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
	 *
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param clusterId The Kafka cluster ID.
	 * @param optional nil or *ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts - Optional Parameters:
	 * @param "BrokerReplicaExclusionBatchRequestData" (optional.Interface of BrokerReplicaExclusionBatchRequestData) -  Alter Broker Replica Exclusions.
	 * @return AlterBrokerReplicaExclusionDataList
	 */
	ClustersClusterIdBrokerReplicaExclusionsdeletePost(ctx _context.Context, clusterId string, localVarOptionals *ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts) (AlterBrokerReplicaExclusionDataList, *_nethttp.Response, error)
}

// BrokerReplicaExclusionApiService BrokerReplicaExclusionApi service
type BrokerReplicaExclusionApiService service

/*
 * ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet Get a Broker Replica Exclusions.
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a Broker Replica Exclusions in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60; and broker specified with &#x60;&#x60;broker_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param brokerId The Kafka broker ID.
 * @return BrokerReplicaExclusionData
 */
func (a *BrokerReplicaExclusionApiService) ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet(ctx _context.Context, clusterId string, brokerId int32) (BrokerReplicaExclusionData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerReplicaExclusionData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/broker-replica-exclusions/{broker_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(clusterId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"broker_id"+"}", _neturl.PathEscape(parameterToString(brokerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

/*
 * ClustersClusterIdBrokerReplicaExclusionsGet Get all Broker Replica Exclusions.
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns all Broker Replica Exclusions in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @return BrokerReplicaExclusionDataList
 */
func (a *BrokerReplicaExclusionApiService) ClustersClusterIdBrokerReplicaExclusionsGet(ctx _context.Context, clusterId string) (BrokerReplicaExclusionDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerReplicaExclusionDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/broker-replica-exclusions"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

// ClustersClusterIdBrokerReplicaExclusionscreatePostOpts Optional parameters for the method 'ClustersClusterIdBrokerReplicaExclusionscreatePost'
type ClustersClusterIdBrokerReplicaExclusionscreatePostOpts struct {
	BrokerReplicaExclusionBatchRequestData optional.Interface
}

/*
 * ClustersClusterIdBrokerReplicaExclusionscreatePost Create Broker Replica Exclusions
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Create Broker Replica Exclusions for brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param optional nil or *ClustersClusterIdBrokerReplicaExclusionscreatePostOpts - Optional Parameters:
 * @param "BrokerReplicaExclusionBatchRequestData" (optional.Interface of BrokerReplicaExclusionBatchRequestData) -  Alter Broker Replica Exclusions.
 * @return AlterBrokerReplicaExclusionDataList
 */
func (a *BrokerReplicaExclusionApiService) ClustersClusterIdBrokerReplicaExclusionscreatePost(ctx _context.Context, clusterId string, localVarOptionals *ClustersClusterIdBrokerReplicaExclusionscreatePostOpts) (AlterBrokerReplicaExclusionDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterBrokerReplicaExclusionDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/broker-replica-exclusions:create"
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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.BrokerReplicaExclusionBatchRequestData.IsSet() {
		localVarOptionalBrokerReplicaExclusionBatchRequestData, localVarOptionalBrokerReplicaExclusionBatchRequestDataok := localVarOptionals.BrokerReplicaExclusionBatchRequestData.Value().(BrokerReplicaExclusionBatchRequestData)
		if !localVarOptionalBrokerReplicaExclusionBatchRequestDataok {
			return localVarReturnValue, nil, reportError("brokerReplicaExclusionBatchRequestData should be BrokerReplicaExclusionBatchRequestData")
		}
		localVarPostBody = &localVarOptionalBrokerReplicaExclusionBatchRequestData
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

// ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts Optional parameters for the method 'ClustersClusterIdBrokerReplicaExclusionsdeletePost'
type ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts struct {
	BrokerReplicaExclusionBatchRequestData optional.Interface
}

/*
 * ClustersClusterIdBrokerReplicaExclusionsdeletePost Delete Broker Replica Exclusions
 *
 * [![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Delete Broker Replica Exclusions for brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param optional nil or *ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts - Optional Parameters:
 * @param "BrokerReplicaExclusionBatchRequestData" (optional.Interface of BrokerReplicaExclusionBatchRequestData) -  Alter Broker Replica Exclusions.
 * @return AlterBrokerReplicaExclusionDataList
 */
func (a *BrokerReplicaExclusionApiService) ClustersClusterIdBrokerReplicaExclusionsdeletePost(ctx _context.Context, clusterId string, localVarOptionals *ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts) (AlterBrokerReplicaExclusionDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterBrokerReplicaExclusionDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/broker-replica-exclusions:delete"
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
	if localVarOptionals != nil && localVarOptionals.BrokerReplicaExclusionBatchRequestData.IsSet() {
		localVarOptionalBrokerReplicaExclusionBatchRequestData, localVarOptionalBrokerReplicaExclusionBatchRequestDataok := localVarOptionals.BrokerReplicaExclusionBatchRequestData.Value().(BrokerReplicaExclusionBatchRequestData)
		if !localVarOptionalBrokerReplicaExclusionBatchRequestDataok {
			return localVarReturnValue, nil, reportError("brokerReplicaExclusionBatchRequestData should be BrokerReplicaExclusionBatchRequestData")
		}
		localVarPostBody = &localVarOptionalBrokerReplicaExclusionBatchRequestData
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v AlterBrokerReplicaExclusionDataList
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
