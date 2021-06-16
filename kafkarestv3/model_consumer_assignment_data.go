/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// ConsumerAssignmentData struct for ConsumerAssignmentData
type ConsumerAssignmentData struct {
	Kind            string           `json:"kind"`
	Metadata        ResourceMetadata `json:"metadata"`
	ClusterId       string           `json:"cluster_id"`
	ConsumerGroupId string           `json:"consumer_group_id"`
	ConsumerId      string           `json:"consumer_id"`
	TopicName       string           `json:"topic_name"`
	PartitionId     int32            `json:"partition_id"`
	Partition       Relationship     `json:"partition"`
	Lag             Relationship     `json:"lag"`
}
