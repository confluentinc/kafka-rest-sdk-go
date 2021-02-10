/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ConsumerGroupLagDataAllOf struct for ConsumerGroupLagDataAllOf
type ConsumerGroupLagDataAllOf struct {
	ClusterId string `json:"cluster_id"`
	ConsumerGroupId string `json:"consumer_group_id"`
	MaxLagConsumerId string `json:"max_lag_consumer_id"`
	MaxLagInstanceId *string `json:"max_lag_instance_id,omitempty"`
	MaxLagClientId string `json:"max_lag_client_id"`
	MaxLagTopicName string `json:"max_lag_topic_name"`
	MaxLagPartitionId int32 `json:"max_lag_partition_id"`
	MaxLag int32 `json:"max_lag"`
	TotalLag int32 `json:"total_lag"`
	MaxLagConsumer Relationship `json:"max_lag_consumer"`
	MaxLagPartition Relationship `json:"max_lag_partition"`
}
