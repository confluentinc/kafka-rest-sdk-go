/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// AbstractConfigData struct for AbstractConfigData
type AbstractConfigData struct {
	Kind string `json:"kind"`
	Metadata ResourceMetadata `json:"metadata"`
	ClusterId string `json:"cluster_id"`
	Name string `json:"name"`
	Value *string `json:"value,omitempty"`
	IsDefault bool `json:"is_default"`
	IsReadOnly bool `json:"is_read_only"`
	IsSensitive bool `json:"is_sensitive"`
	Source ConfigSource `json:"source"`
	Synonyms []ConfigSynonymData `json:"synonyms"`
}
