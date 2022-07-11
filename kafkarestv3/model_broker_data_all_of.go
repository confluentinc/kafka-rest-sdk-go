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

// BrokerDataAllOf struct for BrokerDataAllOf
type BrokerDataAllOf struct {
	ClusterId         string       `json:"cluster_id"`
	BrokerId          int32        `json:"broker_id"`
	Host              *string      `json:"host,omitempty"`
	Port              *int32       `json:"port,omitempty"`
	Rack              *string      `json:"rack,omitempty"`
	Configs           Relationship `json:"configs"`
	PartitionReplicas Relationship `json:"partition_replicas"`
}
