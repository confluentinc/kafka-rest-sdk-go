/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// ListMirrorTopicsResponseData struct for ListMirrorTopicsResponseData
type ListMirrorTopicsResponseData struct {
	Kind            string            `json:"kind"`
	Metadata        ResourceMetadata  `json:"metadata"`
	LinkName        string            `json:"link_name"`
	MirrorTopicName string            `json:"mirror_topic_name"`
	SourceTopicName string            `json:"source_topic_name"`
	NumPartitions   int32             `json:"num_partitions"`
	MirrorLags      []MirrorLag       `json:"mirror_lags"`
	MirrorStatus    MirrorTopicStatus `json:"mirror_status"`
	StateTimeMs     int64             `json:"state_time_ms"`
}
