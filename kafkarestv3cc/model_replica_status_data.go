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

package kafkarestv3cc

// ReplicaStatusData struct for ReplicaStatusData
type ReplicaStatusData struct {
	Kind               string           `json:"kind"`
	Metadata           ResourceMetadata `json:"metadata"`
	ClusterId          string           `json:"clusterId"`
	TopicName          string           `json:"topicName"`
	BrokerId           int32            `json:"brokerId"`
	PartitionId        int32            `json:"partitionId"`
	Leader             bool             `json:"leader"`
	Observer           bool             `json:"observer"`
	IsrEligible        bool             `json:"isrEligible"`
	InIsr              bool             `json:"inIsr"`
	CaughtUp           bool             `json:"caughtUp"`
	LogStartOffset     int32            `json:"logStartOffset"`
	LogEndOffset       int32            `json:"logEndOffset"`
	LastCaughtUpTimeMs int32            `json:"lastCaughtUpTimeMs"`
	LastFetchTimeMs    int32            `json:"lastFetchTimeMs"`
	LinkName           string           `json:"linkName,omitempty"`
}
