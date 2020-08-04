/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ReassignmentDataAllOf struct for ReassignmentDataAllOf
type ReassignmentDataAllOf struct {
	ClusterId string `json:"cluster_id"`
	TopicName string `json:"topic_name"`
	PartitionId int32 `json:"partition_id"`
	AddingReplicas []int32 `json:"adding_replicas"`
	RemovingReplicas []int32 `json:"removing_replicas"`
	Replicas Relationship `json:"replicas"`
}
