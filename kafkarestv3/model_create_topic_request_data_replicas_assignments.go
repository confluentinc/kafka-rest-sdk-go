/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// CreateTopicRequestDataReplicasAssignments struct for CreateTopicRequestDataReplicasAssignments
type CreateTopicRequestDataReplicasAssignments struct {
	PartitionId int32 `json:"partition_id"`
	BrokerIds []int32 `json:"broker_ids"`
}
