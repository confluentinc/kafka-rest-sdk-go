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

// AclData struct for AclData
type AclData struct {
	Kind         string           `json:"kind"`
	Metadata     ResourceMetadata `json:"metadata"`
	ClusterId    string           `json:"cluster_id"`
	ResourceType AclResourceType  `json:"resource_type"`
	ResourceName string           `json:"resource_name"`
	PatternType  string           `json:"pattern_type"`
	Principal    string           `json:"principal"`
	Host         string           `json:"host"`
	Operation    string           `json:"operation"`
	Permission   string           `json:"permission"`
}
