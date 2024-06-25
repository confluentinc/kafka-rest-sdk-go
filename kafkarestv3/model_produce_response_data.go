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

// ProduceResponseData struct for ProduceResponseData
type ProduceResponseData struct {
	Size          int64   `json:"size"`
	Type          *string `json:"type"`
	Subject       *string `json:"subject,omitempty"`
	SchemaId      *int32  `json:"schema_id,omitempty"`
	SchemaVersion *int32  `json:"schema_version,omitempty"`
}
