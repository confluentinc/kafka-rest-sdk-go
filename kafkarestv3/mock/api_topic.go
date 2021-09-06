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
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_topic.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// TopicApi is a mock of TopicApi interface
type TopicApi struct {
	lockClustersClusterIdTopicsGet sync.Mutex
	ClustersClusterIdTopicsGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicDataList, *net_http.Response, error)

	lockClustersClusterIdTopicsPost sync.Mutex
	ClustersClusterIdTopicsPostFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsPostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicData, *net_http.Response, error)

	lockClustersClusterIdTopicsTopicNameDelete sync.Mutex
	ClustersClusterIdTopicsTopicNameDeleteFunc func(ctx context.Context, clusterId, topicName string) (*net_http.Response, error)

	lockClustersClusterIdTopicsTopicNameGet sync.Mutex
	ClustersClusterIdTopicsTopicNameGetFunc func(ctx context.Context, clusterId, topicName string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicData, *net_http.Response, error)

	calls struct {
		ClustersClusterIdTopicsGet []struct {
			Ctx       context.Context
			ClusterId string
		}
		ClustersClusterIdTopicsPost []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsPostOpts
		}
		ClustersClusterIdTopicsTopicNameDelete []struct {
			Ctx       context.Context
			ClusterId string
			TopicName string
		}
		ClustersClusterIdTopicsTopicNameGet []struct {
			Ctx       context.Context
			ClusterId string
			TopicName string
		}
	}
}

// ClustersClusterIdTopicsGet mocks base method by wrapping the associated func.
func (m *TopicApi) ClustersClusterIdTopicsGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicDataList, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsGet.Lock()
	defer m.lockClustersClusterIdTopicsGet.Unlock()

	if m.ClustersClusterIdTopicsGetFunc == nil {
		panic("mocker: TopicApi.ClustersClusterIdTopicsGetFunc is nil but TopicApi.ClustersClusterIdTopicsGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdTopicsGet = append(m.calls.ClustersClusterIdTopicsGet, call)

	return m.ClustersClusterIdTopicsGetFunc(ctx, clusterId)
}

// ClustersClusterIdTopicsGetCalled returns true if ClustersClusterIdTopicsGet was called at least once.
func (m *TopicApi) ClustersClusterIdTopicsGetCalled() bool {
	m.lockClustersClusterIdTopicsGet.Lock()
	defer m.lockClustersClusterIdTopicsGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsGet) > 0
}

// ClustersClusterIdTopicsGetCalls returns the calls made to ClustersClusterIdTopicsGet.
func (m *TopicApi) ClustersClusterIdTopicsGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdTopicsGet.Lock()
	defer m.lockClustersClusterIdTopicsGet.Unlock()

	return m.calls.ClustersClusterIdTopicsGet
}

// ClustersClusterIdTopicsPost mocks base method by wrapping the associated func.
func (m *TopicApi) ClustersClusterIdTopicsPost(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsPostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicData, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsPost.Lock()
	defer m.lockClustersClusterIdTopicsPost.Unlock()

	if m.ClustersClusterIdTopicsPostFunc == nil {
		panic("mocker: TopicApi.ClustersClusterIdTopicsPostFunc is nil but TopicApi.ClustersClusterIdTopicsPost was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsPostOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.ClustersClusterIdTopicsPost = append(m.calls.ClustersClusterIdTopicsPost, call)

	return m.ClustersClusterIdTopicsPostFunc(ctx, clusterId, localVarOptionals)
}

// ClustersClusterIdTopicsPostCalled returns true if ClustersClusterIdTopicsPost was called at least once.
func (m *TopicApi) ClustersClusterIdTopicsPostCalled() bool {
	m.lockClustersClusterIdTopicsPost.Lock()
	defer m.lockClustersClusterIdTopicsPost.Unlock()

	return len(m.calls.ClustersClusterIdTopicsPost) > 0
}

// ClustersClusterIdTopicsPostCalls returns the calls made to ClustersClusterIdTopicsPost.
func (m *TopicApi) ClustersClusterIdTopicsPostCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsPostOpts
} {
	m.lockClustersClusterIdTopicsPost.Lock()
	defer m.lockClustersClusterIdTopicsPost.Unlock()

	return m.calls.ClustersClusterIdTopicsPost
}

// ClustersClusterIdTopicsTopicNameDelete mocks base method by wrapping the associated func.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameDelete(ctx context.Context, clusterId, topicName string) (*net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNameDelete.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameDelete.Unlock()

	if m.ClustersClusterIdTopicsTopicNameDeleteFunc == nil {
		panic("mocker: TopicApi.ClustersClusterIdTopicsTopicNameDeleteFunc is nil but TopicApi.ClustersClusterIdTopicsTopicNameDelete was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		TopicName string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		TopicName: topicName,
	}

	m.calls.ClustersClusterIdTopicsTopicNameDelete = append(m.calls.ClustersClusterIdTopicsTopicNameDelete, call)

	return m.ClustersClusterIdTopicsTopicNameDeleteFunc(ctx, clusterId, topicName)
}

// ClustersClusterIdTopicsTopicNameDeleteCalled returns true if ClustersClusterIdTopicsTopicNameDelete was called at least once.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameDeleteCalled() bool {
	m.lockClustersClusterIdTopicsTopicNameDelete.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameDelete.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNameDelete) > 0
}

// ClustersClusterIdTopicsTopicNameDeleteCalls returns the calls made to ClustersClusterIdTopicsTopicNameDelete.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameDeleteCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TopicName string
} {
	m.lockClustersClusterIdTopicsTopicNameDelete.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameDelete.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNameDelete
}

// ClustersClusterIdTopicsTopicNameGet mocks base method by wrapping the associated func.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameGet(ctx context.Context, clusterId, topicName string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.TopicData, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNameGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameGet.Unlock()

	if m.ClustersClusterIdTopicsTopicNameGetFunc == nil {
		panic("mocker: TopicApi.ClustersClusterIdTopicsTopicNameGetFunc is nil but TopicApi.ClustersClusterIdTopicsTopicNameGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		TopicName string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		TopicName: topicName,
	}

	m.calls.ClustersClusterIdTopicsTopicNameGet = append(m.calls.ClustersClusterIdTopicsTopicNameGet, call)

	return m.ClustersClusterIdTopicsTopicNameGetFunc(ctx, clusterId, topicName)
}

// ClustersClusterIdTopicsTopicNameGetCalled returns true if ClustersClusterIdTopicsTopicNameGet was called at least once.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameGetCalled() bool {
	m.lockClustersClusterIdTopicsTopicNameGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNameGet) > 0
}

// ClustersClusterIdTopicsTopicNameGetCalls returns the calls made to ClustersClusterIdTopicsTopicNameGet.
func (m *TopicApi) ClustersClusterIdTopicsTopicNameGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TopicName string
} {
	m.lockClustersClusterIdTopicsTopicNameGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameGet.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNameGet
}

// Reset resets the calls made to the mocked methods.
func (m *TopicApi) Reset() {
	m.lockClustersClusterIdTopicsGet.Lock()
	m.calls.ClustersClusterIdTopicsGet = nil
	m.lockClustersClusterIdTopicsGet.Unlock()
	m.lockClustersClusterIdTopicsPost.Lock()
	m.calls.ClustersClusterIdTopicsPost = nil
	m.lockClustersClusterIdTopicsPost.Unlock()
	m.lockClustersClusterIdTopicsTopicNameDelete.Lock()
	m.calls.ClustersClusterIdTopicsTopicNameDelete = nil
	m.lockClustersClusterIdTopicsTopicNameDelete.Unlock()
	m.lockClustersClusterIdTopicsTopicNameGet.Lock()
	m.calls.ClustersClusterIdTopicsTopicNameGet = nil
	m.lockClustersClusterIdTopicsTopicNameGet.Unlock()
}
