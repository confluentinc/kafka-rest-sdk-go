/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type ConsumerGroupApi interface {

    /*
     * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet List Consumer Assignments
     *
     * Returns a list of partition assignments for the specified consumer.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param consumerGroupId The consumer group ID.
     * @param consumerId The consumer ID.
     * @return ConsumerAssignmentDataList
     */
    ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string) (ConsumerAssignmentDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet Get Consumer Assignment
     *
     * Returns information about the assignment for the specified consumer to the specified partition.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param consumerGroupId The consumer group ID.
     * @param consumerId The consumer ID.
     * @param topicName The topic name.
     * @param partitionId The partition ID.
     * @return ConsumerAssignmentData
     */
    ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string, topicName string, partitionId int32) (ConsumerAssignmentData, *_nethttp.Response, error)

    /*
     * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet Get Consumer
     *
     * Returns the consumer specified by the &#x60;&#x60;consumer_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param consumerGroupId The consumer group ID.
     * @param consumerId The consumer ID.
     * @return ConsumerData
     */
    ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string) (ConsumerData, *_nethttp.Response, error)

    /*
     * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet List Consumers
     *
     * Returns a list of consumers that belong to the specified consumer group.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param consumerGroupId The consumer group ID.
     * @return ConsumerDataList
     */
    ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet(ctx _context.Context, clusterId string, consumerGroupId string) (ConsumerDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdConsumerGroupsConsumerGroupIdGet Get Consumer Group
     *
     * Returns the consumer group specified by the &#x60;&#x60;consumer_group_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param consumerGroupId The consumer group ID.
     * @return ConsumerGroupData
     */
    ClustersClusterIdConsumerGroupsConsumerGroupIdGet(ctx _context.Context, clusterId string, consumerGroupId string) (ConsumerGroupData, *_nethttp.Response, error)

    /*
     * ClustersClusterIdConsumerGroupsGet List Consumer Groups
     *
     * Returns the list of consumer groups that belong to the specified Kafka cluster.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @return ConsumerGroupDataList
     */
    ClustersClusterIdConsumerGroupsGet(ctx _context.Context, clusterId string) (ConsumerGroupDataList, *_nethttp.Response, error)
}

// ConsumerGroupApiService ConsumerGroupApi service
type ConsumerGroupApiService service

/*
 * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet List Consumer Assignments
 *
 * Returns a list of partition assignments for the specified consumer.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param consumerGroupId The consumer group ID.
 * @param consumerId The consumer ID.
 * @return ConsumerAssignmentDataList
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string) (ConsumerAssignmentDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerAssignmentDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_group_id"+"}", _neturl.QueryEscape(parameterToString(consumerGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_id"+"}", _neturl.QueryEscape(parameterToString(consumerId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
 * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet Get Consumer Assignment
 *
 * Returns information about the assignment for the specified consumer to the specified partition.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param consumerGroupId The consumer group ID.
 * @param consumerId The consumer ID.
 * @param topicName The topic name.
 * @param partitionId The partition ID.
 * @return ConsumerAssignmentData
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdAssignmentsTopicNamePartitionsPartitionIdGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string, topicName string, partitionId int32) (ConsumerAssignmentData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerAssignmentData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}/assignments/{topic_name}/partitions/{partition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_group_id"+"}", _neturl.QueryEscape(parameterToString(consumerGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_id"+"}", _neturl.QueryEscape(parameterToString(consumerId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"topic_name"+"}", _neturl.QueryEscape(parameterToString(topicName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"partition_id"+"}", _neturl.QueryEscape(parameterToString(partitionId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
 * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet Get Consumer
 *
 * Returns the consumer specified by the &#x60;&#x60;consumer_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param consumerGroupId The consumer group ID.
 * @param consumerId The consumer ID.
 * @return ConsumerData
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersConsumerIdGet(ctx _context.Context, clusterId string, consumerGroupId string, consumerId string) (ConsumerData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers/{consumer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_group_id"+"}", _neturl.QueryEscape(parameterToString(consumerGroupId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_id"+"}", _neturl.QueryEscape(parameterToString(consumerId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
 * ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet List Consumers
 *
 * Returns a list of consumers that belong to the specified consumer group.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param consumerGroupId The consumer group ID.
 * @return ConsumerDataList
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsConsumerGroupIdConsumersGet(ctx _context.Context, clusterId string, consumerGroupId string) (ConsumerDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups/{consumer_group_id}/consumers"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_group_id"+"}", _neturl.QueryEscape(parameterToString(consumerGroupId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
 * ClustersClusterIdConsumerGroupsConsumerGroupIdGet Get Consumer Group
 *
 * Returns the consumer group specified by the &#x60;&#x60;consumer_group_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param consumerGroupId The consumer group ID.
 * @return ConsumerGroupData
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsConsumerGroupIdGet(ctx _context.Context, clusterId string, consumerGroupId string) (ConsumerGroupData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerGroupData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups/{consumer_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"consumer_group_id"+"}", _neturl.QueryEscape(parameterToString(consumerGroupId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
 * ClustersClusterIdConsumerGroupsGet List Consumer Groups
 *
 * Returns the list of consumer groups that belong to the specified Kafka cluster.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @return ConsumerGroupDataList
 */
func (a *ConsumerGroupApiService) ClustersClusterIdConsumerGroupsGet(ctx _context.Context, clusterId string) (ConsumerGroupDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConsumerGroupDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/consumer-groups"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
