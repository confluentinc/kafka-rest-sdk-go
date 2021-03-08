/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ListMirrorTopicsResponseDataAllOf1 struct for ListMirrorTopicsResponseDataAllOf1
type ListMirrorTopicsResponseDataAllOf1 struct {
	LinkName string `json:"link_name"`
	DestinationTopicName string `json:"destination_topic_name"`
	SourceTopicName string `json:"source_topic_name"`
	NumPartitions int32 `json:"num_partitions"`
	MirrorLags MirrorLags `json:"mirror_lags"`
	MirrorTopicStatus MirrorTopicStatus `json:"mirror_topic_status"`
	StateTimeMs int32 `json:"state_time_ms"`
}
