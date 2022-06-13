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

// ClusterDataAllOf struct for ClusterDataAllOf
type ClusterDataAllOf struct {
	ClusterId              string       `json:"cluster_id"`
	Controller             Relationship `json:"controller,omitempty"`
	Acls                   Relationship `json:"acls"`
	Brokers                Relationship `json:"brokers"`
	BrokerConfigs          Relationship `json:"broker_configs"`
	ConsumerGroups         Relationship `json:"consumer_groups"`
	Topics                 Relationship `json:"topics"`
	PartitionReassignments Relationship `json:"partition_reassignments"`
}
