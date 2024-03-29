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

// ConsumerGroupDataAllOf struct for ConsumerGroupDataAllOf
type ConsumerGroupDataAllOf struct {
	ClusterId         string       `json:"cluster_id"`
	ConsumerGroupId   string       `json:"consumer_group_id"`
	IsSimple          bool         `json:"is_simple"`
	PartitionAssignor string       `json:"partition_assignor"`
	State             string       `json:"state"`
	Coordinator       Relationship `json:"coordinator"`
	Consumer          Relationship `json:"consumer,omitempty"`
	LagSummary        Relationship `json:"lag_summary"`
}
