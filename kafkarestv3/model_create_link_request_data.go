/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// CreateLinkRequestData struct for CreateLinkRequestData
type CreateLinkRequestData struct {
	LinkName string `json:"link_name"`
	Configs []CreateTopicRequestDataConfigs `json:"configs,omitempty"`
}
