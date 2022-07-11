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

// ConsumerLagDataAllOf struct for ConsumerLagDataAllOf
type ConsumerLagDataAllOf struct {
	ClusterId       string  `json:"cluster_id"`
	ConsumerGroupId string  `json:"consumer_group_id"`
	TopicName       string  `json:"topic_name"`
	PartitionId     int32   `json:"partition_id"`
	CurrentOffset   int64   `json:"current_offset"`
	LogEndOffset    int64   `json:"log_end_offset"`
	Lag             int64   `json:"lag"`
	ConsumerId      string  `json:"consumer_id"`
	InstanceId      *string `json:"instance_id,omitempty"`
	ClientId        string  `json:"client_id"`
}
