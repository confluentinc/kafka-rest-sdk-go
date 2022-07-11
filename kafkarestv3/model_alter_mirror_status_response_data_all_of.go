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

// AlterMirrorStatusResponseDataAllOf struct for AlterMirrorStatusResponseDataAllOf
type AlterMirrorStatusResponseDataAllOf struct {
	MirrorTopicName string      `json:"mirror_topic_name"`
	ErrorMessage    *string     `json:"error_message"`
	ErrorCode       *int32      `json:"error_code"`
	MirrorLags      []MirrorLag `json:"mirror_lags"`
}
