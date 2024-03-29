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

// CreateLinkRequestData struct for CreateLinkRequestData
type CreateLinkRequestData struct {
	SourceClusterId      string       `json:"source_cluster_id,omitempty"`
	DestinationClusterId string       `json:"destination_cluster_id,omitempty"`
	RemoteClusterId      string       `json:"remote_cluster_id,omitempty"`
	ClusterLinkId        string       `json:"cluster_link_id,omitempty"`
	Configs              []ConfigData `json:"configs,omitempty"`
}
