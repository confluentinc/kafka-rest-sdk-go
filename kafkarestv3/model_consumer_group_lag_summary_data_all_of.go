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

// ConsumerGroupLagSummaryDataAllOf struct for ConsumerGroupLagSummaryDataAllOf
type ConsumerGroupLagSummaryDataAllOf struct {
	ClusterId         string       `json:"cluster_id"`
	ConsumerGroupId   string       `json:"consumer_group_id"`
	MaxLagConsumerId  string       `json:"max_lag_consumer_id"`
	MaxLagInstanceId  *string      `json:"max_lag_instance_id,omitempty"`
	MaxLagClientId    string       `json:"max_lag_client_id"`
	MaxLagTopicName   string       `json:"max_lag_topic_name"`
	MaxLagPartitionId int32        `json:"max_lag_partition_id"`
	MaxLag            int64        `json:"max_lag"`
	TotalLag          int64        `json:"total_lag"`
	MaxLagConsumer    Relationship `json:"max_lag_consumer"`
	MaxLagPartition   Relationship `json:"max_lag_partition"`
}
