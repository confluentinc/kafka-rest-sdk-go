/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// AlterMirrorsRequestData struct for AlterMirrorsRequestData
type AlterMirrorsRequestData struct {
	DestinationTopicNames []string `json:"destination_topic_names"`
}
