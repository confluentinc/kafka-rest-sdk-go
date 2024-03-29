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

// TopicConfigData struct for TopicConfigData
type TopicConfigData struct {
	Kind        string              `json:"kind"`
	Metadata    ResourceMetadata    `json:"metadata"`
	ClusterId   string              `json:"cluster_id"`
	Name        string              `json:"name"`
	Value       *string             `json:"value,omitempty"`
	IsDefault   bool                `json:"is_default"`
	IsReadOnly  bool                `json:"is_read_only"`
	IsSensitive bool                `json:"is_sensitive"`
	Source      string              `json:"source"`
	Synonyms    []ConfigSynonymData `json:"synonyms"`
	TopicName   string              `json:"topic_name"`
}
