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

// ProduceRequestData struct for ProduceRequestData
type ProduceRequestData struct {
	Type                string       `json:"type,omitempty"`
	Subject             *string      `json:"subject,omitempty"`
	SubjectNameStrategy *string      `json:"subject_name_strategy,omitempty"`
	SchemaId            *int32       `json:"schema_id,omitempty"`
	SchemaVersion       *int32       `json:"schema_version,omitempty"`
	Schema              *string      `json:"schema,omitempty"`
	Data                *interface{} `json:"data,omitempty"`
}
