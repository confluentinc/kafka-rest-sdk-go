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

// AlterMirrorsRequestData struct for AlterMirrorsRequestData
type AlterMirrorsRequestData struct {
	// The mirror topics specified as a list of topic names.
	MirrorTopicNames []string `json:"mirror_topic_names,omitempty"`
	// The mirror topics specified as a pattern.
	MirrorTopicNamePattern string `json:"mirror_topic_name_pattern,omitempty"`
}
