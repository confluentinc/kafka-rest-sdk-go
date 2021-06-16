/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// AlterMirrorStatusResponseData struct for AlterMirrorStatusResponseData
type AlterMirrorStatusResponseData struct {
	Kind            string           `json:"kind"`
	Metadata        ResourceMetadata `json:"metadata"`
	MirrorTopicName string           `json:"mirror_topic_name"`
	ErrorMessage    *string          `json:"error_message"`
	ErrorCode       *int32           `json:"error_code"`
	MirrorLags      []MirrorLag      `json:"mirror_lags,omitempty"`
}
