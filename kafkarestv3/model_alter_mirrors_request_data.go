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
	MirrorTopicNames []string `json:"mirror_topic_names"`
}
