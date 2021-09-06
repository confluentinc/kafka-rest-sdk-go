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

// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_broker.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// BrokerApi is a mock of BrokerApi interface
type BrokerApi struct {
	lockClustersClusterIdBrokersBrokerIdGet sync.Mutex
	ClustersClusterIdBrokersBrokerIdGetFunc func(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerData, *net_http.Response, error)

	lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet sync.Mutex
	ClustersClusterIdBrokersBrokerIdPartitionReplicasGetFunc func(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaDataList, *net_http.Response, error)

	lockClustersClusterIdBrokersGet sync.Mutex
	ClustersClusterIdBrokersGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdBrokersBrokerIdGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
		}
		ClustersClusterIdBrokersBrokerIdPartitionReplicasGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
		}
		ClustersClusterIdBrokersGet []struct {
			Ctx       context.Context
			ClusterId string
		}
	}
}

// ClustersClusterIdBrokersBrokerIdGet mocks base method by wrapping the associated func.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdGet(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerData, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdGet.Unlock()

	if m.ClustersClusterIdBrokersBrokerIdGetFunc == nil {
		panic("mocker: BrokerApi.ClustersClusterIdBrokersBrokerIdGetFunc is nil but BrokerApi.ClustersClusterIdBrokersBrokerIdGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		BrokerId  int32
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		BrokerId:  brokerId,
	}

	m.calls.ClustersClusterIdBrokersBrokerIdGet = append(m.calls.ClustersClusterIdBrokersBrokerIdGet, call)

	return m.ClustersClusterIdBrokersBrokerIdGetFunc(ctx, clusterId, brokerId)
}

// ClustersClusterIdBrokersBrokerIdGetCalled returns true if ClustersClusterIdBrokersBrokerIdGet was called at least once.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdGetCalled() bool {
	m.lockClustersClusterIdBrokersBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersBrokerIdGet) > 0
}

// ClustersClusterIdBrokersBrokerIdGetCalls returns the calls made to ClustersClusterIdBrokersBrokerIdGet.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
} {
	m.lockClustersClusterIdBrokersBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdGet.Unlock()

	return m.calls.ClustersClusterIdBrokersBrokerIdGet
}

// ClustersClusterIdBrokersBrokerIdPartitionReplicasGet mocks base method by wrapping the associated func.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdPartitionReplicasGet(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Unlock()

	if m.ClustersClusterIdBrokersBrokerIdPartitionReplicasGetFunc == nil {
		panic("mocker: BrokerApi.ClustersClusterIdBrokersBrokerIdPartitionReplicasGetFunc is nil but BrokerApi.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		BrokerId  int32
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		BrokerId:  brokerId,
	}

	m.calls.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet = append(m.calls.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet, call)

	return m.ClustersClusterIdBrokersBrokerIdPartitionReplicasGetFunc(ctx, clusterId, brokerId)
}

// ClustersClusterIdBrokersBrokerIdPartitionReplicasGetCalled returns true if ClustersClusterIdBrokersBrokerIdPartitionReplicasGet was called at least once.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdPartitionReplicasGetCalled() bool {
	m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet) > 0
}

// ClustersClusterIdBrokersBrokerIdPartitionReplicasGetCalls returns the calls made to ClustersClusterIdBrokersBrokerIdPartitionReplicasGet.
func (m *BrokerApi) ClustersClusterIdBrokersBrokerIdPartitionReplicasGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
} {
	m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Unlock()

	return m.calls.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet
}

// ClustersClusterIdBrokersGet mocks base method by wrapping the associated func.
func (m *BrokerApi) ClustersClusterIdBrokersGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersGet.Lock()
	defer m.lockClustersClusterIdBrokersGet.Unlock()

	if m.ClustersClusterIdBrokersGetFunc == nil {
		panic("mocker: BrokerApi.ClustersClusterIdBrokersGetFunc is nil but BrokerApi.ClustersClusterIdBrokersGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdBrokersGet = append(m.calls.ClustersClusterIdBrokersGet, call)

	return m.ClustersClusterIdBrokersGetFunc(ctx, clusterId)
}

// ClustersClusterIdBrokersGetCalled returns true if ClustersClusterIdBrokersGet was called at least once.
func (m *BrokerApi) ClustersClusterIdBrokersGetCalled() bool {
	m.lockClustersClusterIdBrokersGet.Lock()
	defer m.lockClustersClusterIdBrokersGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersGet) > 0
}

// ClustersClusterIdBrokersGetCalls returns the calls made to ClustersClusterIdBrokersGet.
func (m *BrokerApi) ClustersClusterIdBrokersGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdBrokersGet.Lock()
	defer m.lockClustersClusterIdBrokersGet.Unlock()

	return m.calls.ClustersClusterIdBrokersGet
}

// Reset resets the calls made to the mocked methods.
func (m *BrokerApi) Reset() {
	m.lockClustersClusterIdBrokersBrokerIdGet.Lock()
	m.calls.ClustersClusterIdBrokersBrokerIdGet = nil
	m.lockClustersClusterIdBrokersBrokerIdGet.Unlock()
	m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Lock()
	m.calls.ClustersClusterIdBrokersBrokerIdPartitionReplicasGet = nil
	m.lockClustersClusterIdBrokersBrokerIdPartitionReplicasGet.Unlock()
	m.lockClustersClusterIdBrokersGet.Lock()
	m.calls.ClustersClusterIdBrokersGet = nil
	m.lockClustersClusterIdBrokersGet.Unlock()
}
