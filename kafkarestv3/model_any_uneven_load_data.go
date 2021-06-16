/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

import (
	"time"
)

// AnyUnevenLoadData struct for AnyUnevenLoadData
type AnyUnevenLoadData struct {
	Kind           string              `json:"kind"`
	Metadata       ResourceMetadata    `json:"metadata"`
	ClusterId      string              `json:"cluster_id"`
	Status         AnyUnevenLoadStatus `json:"status"`
	PreviousStatus AnyUnevenLoadStatus `json:"previous_status"`
	// The date and time at which this task was created.
	StatusUpdatedAt time.Time `json:"status_updated_at"`
	// The date and time at which this task was created.
	PreviousStatusUpdatedAt time.Time    `json:"previous_status_updated_at"`
	ErrorCode               *int32       `json:"error_code,omitempty"`
	ErrorMessage            *string      `json:"error_message,omitempty"`
	BrokerTasks             Relationship `json:"broker_tasks"`
}
