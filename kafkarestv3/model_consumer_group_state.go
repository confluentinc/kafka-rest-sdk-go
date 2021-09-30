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

// ConsumerGroupState the model 'ConsumerGroupState'
type ConsumerGroupState string

// List of ConsumerGroupState
const (
	CONSUMERGROUPSTATE_UNKNOWN              ConsumerGroupState = "UNKNOWN"
	CONSUMERGROUPSTATE_PREPARING_REBALANCE  ConsumerGroupState = "PREPARING_REBALANCE"
	CONSUMERGROUPSTATE_COMPLETING_REBALANCE ConsumerGroupState = "COMPLETING_REBALANCE"
	CONSUMERGROUPSTATE_STABLE               ConsumerGroupState = "STABLE"
	CONSUMERGROUPSTATE_DEAD                 ConsumerGroupState = "DEAD"
	CONSUMERGROUPSTATE_EMPTY                ConsumerGroupState = "EMPTY"
)
