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
// BrokerTaskData struct for BrokerTaskData
type BrokerTaskData struct {
	Kind string `json:"kind"`
	Metadata ResourceMetadata `json:"metadata"`
	ClusterId string `json:"cluster_id"`
	BrokerId int32 `json:"broker_id"`
	TaskType BrokerTaskType `json:"task_type"`
	TaskStatus BrokerTaskStatus `json:"task_status"`
	SubTaskStatuses map[string]string `json:"sub_task_statuses"`
	// The date and time at which this task was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time at which this task was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	ErrorCode *int32 `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Broker Relationship `json:"broker"`
}
