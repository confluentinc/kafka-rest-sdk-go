/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// RemoveBrokerTaskDataAllOf struct for RemoveBrokerTaskDataAllOf
type RemoveBrokerTaskDataAllOf struct {
	ClusterId string `json:"cluster_id"`
	BrokerId int32 `json:"broker_id"`
	PartitionReassignmentStatus DeprecatedPartitionReassignmentStatus `json:"partition_reassignment_status"`
	BrokerShutdownStatus DeprecatedBrokerShutdownStatus `json:"broker_shutdown_status"`
	ErrorCode *int32 `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Broker Relationship `json:"broker"`
}