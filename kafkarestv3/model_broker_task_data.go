// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

import (
	"time"
)

// BrokerTaskData struct for BrokerTaskData
type BrokerTaskData struct {
	Kind              string            `json:"kind"`
	Metadata          ResourceMetadata  `json:"metadata"`
	ClusterId         string            `json:"cluster_id"`
	BrokerId          int32             `json:"broker_id"`
	TaskType          BrokerTaskType    `json:"task_type"`
	TaskStatus        BrokerTaskStatus  `json:"task_status"`
	ShutdownScheduled *bool             `json:"shutdown_scheduled,omitempty"`
	SubTaskStatuses   map[string]string `json:"sub_task_statuses"`
	// The date and time at which this task was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time at which this task was last updated.
	UpdatedAt    time.Time    `json:"updated_at"`
	ErrorCode    *int32       `json:"error_code,omitempty"`
	ErrorMessage *string      `json:"error_message,omitempty"`
	Broker       Relationship `json:"broker"`
}
