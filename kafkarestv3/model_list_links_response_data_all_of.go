/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ListLinksResponseDataAllOf struct for ListLinksResponseDataAllOf
type ListLinksResponseDataAllOf struct {
	SourceClusterId string `json:"source_cluster_id"`
	LinkName string `json:"link_name"`
	LinkId string `json:"link_id"`
	TopicsNames []string `json:"topics_names,omitempty"`
}
