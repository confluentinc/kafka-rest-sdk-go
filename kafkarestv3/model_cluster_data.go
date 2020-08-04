/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ClusterData struct for ClusterData
type ClusterData struct {
	Kind string `json:"kind"`
	Metadata ResourceMetadata `json:"metadata"`
	ClusterId string `json:"cluster_id"`
	Controller Relationship `json:"controller,omitempty"`
	Acls Relationship `json:"acls"`
	Brokers Relationship `json:"brokers"`
	BrokerConfigs Relationship `json:"broker_configs"`
	ConsumerGroups Relationship `json:"consumer_groups"`
	Topics Relationship `json:"topics"`
	PartitionReassignments Relationship `json:"partition_reassignments"`
}
