// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_records.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// RecordsApi is a mock of RecordsApi interface
type RecordsApi struct {
	lockClustersClusterIdTopicsTopicNameRecordsPost sync.Mutex
	ClustersClusterIdTopicsTopicNameRecordsPostFunc func(ctx context.Context, clusterId, topicName string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsTopicNameRecordsPostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ProduceResponse, *net_http.Response, error)

	calls struct {
		ClustersClusterIdTopicsTopicNameRecordsPost []struct {
			Ctx               context.Context
			ClusterId         string
			TopicName         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsTopicNameRecordsPostOpts
		}
	}
}

// ClustersClusterIdTopicsTopicNameRecordsPost mocks base method by wrapping the associated func.
func (m *RecordsApi) ClustersClusterIdTopicsTopicNameRecordsPost(ctx context.Context, clusterId, topicName string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsTopicNameRecordsPostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ProduceResponse, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNameRecordsPost.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameRecordsPost.Unlock()

	if m.ClustersClusterIdTopicsTopicNameRecordsPostFunc == nil {
		panic("mocker: RecordsApi.ClustersClusterIdTopicsTopicNameRecordsPostFunc is nil but RecordsApi.ClustersClusterIdTopicsTopicNameRecordsPost was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		TopicName         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsTopicNameRecordsPostOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		TopicName:         topicName,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.ClustersClusterIdTopicsTopicNameRecordsPost = append(m.calls.ClustersClusterIdTopicsTopicNameRecordsPost, call)

	return m.ClustersClusterIdTopicsTopicNameRecordsPostFunc(ctx, clusterId, topicName, localVarOptionals)
}

// ClustersClusterIdTopicsTopicNameRecordsPostCalled returns true if ClustersClusterIdTopicsTopicNameRecordsPost was called at least once.
func (m *RecordsApi) ClustersClusterIdTopicsTopicNameRecordsPostCalled() bool {
	m.lockClustersClusterIdTopicsTopicNameRecordsPost.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameRecordsPost.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNameRecordsPost) > 0
}

// ClustersClusterIdTopicsTopicNameRecordsPostCalls returns the calls made to ClustersClusterIdTopicsTopicNameRecordsPost.
func (m *RecordsApi) ClustersClusterIdTopicsTopicNameRecordsPostCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	TopicName         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdTopicsTopicNameRecordsPostOpts
} {
	m.lockClustersClusterIdTopicsTopicNameRecordsPost.Lock()
	defer m.lockClustersClusterIdTopicsTopicNameRecordsPost.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNameRecordsPost
}

// Reset resets the calls made to the mocked methods.
func (m *RecordsApi) Reset() {
	m.lockClustersClusterIdTopicsTopicNameRecordsPost.Lock()
	m.calls.ClustersClusterIdTopicsTopicNameRecordsPost = nil
	m.lockClustersClusterIdTopicsTopicNameRecordsPost.Unlock()
}