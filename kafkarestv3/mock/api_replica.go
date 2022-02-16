// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_replica.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// ReplicaApi is a mock of ReplicaApi interface
type ReplicaApi struct {
	lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet sync.Mutex
	ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetFunc func(ctx context.Context, clusterId, topicName string, partitionId, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaData, *net_http.Response, error)

	lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet sync.Mutex
	ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetFunc func(ctx context.Context, clusterId, topicName string, partitionId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet []struct {
			Ctx         context.Context
			ClusterId   string
			TopicName   string
			PartitionId int32
			BrokerId    int32
		}
		ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet []struct {
			Ctx         context.Context
			ClusterId   string
			TopicName   string
			PartitionId int32
		}
	}
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet mocks base method by wrapping the associated func.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet(ctx context.Context, clusterId, topicName string, partitionId, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaData, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Unlock()

	if m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetFunc == nil {
		panic("mocker: ReplicaApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetFunc is nil but ReplicaApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet was called.")
	}

	call := struct {
		Ctx         context.Context
		ClusterId   string
		TopicName   string
		PartitionId int32
		BrokerId    int32
	}{
		Ctx:         ctx,
		ClusterId:   clusterId,
		TopicName:   topicName,
		PartitionId: partitionId,
		BrokerId:    brokerId,
	}

	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet = append(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet, call)

	return m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetFunc(ctx, clusterId, topicName, partitionId, brokerId)
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetCalled returns true if ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet was called at least once.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetCalled() bool {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet) > 0
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetCalls returns the calls made to ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGetCalls() []struct {
	Ctx         context.Context
	ClusterId   string
	TopicName   string
	PartitionId int32
	BrokerId    int32
} {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet mocks base method by wrapping the associated func.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet(ctx context.Context, clusterId, topicName string, partitionId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaDataList, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Unlock()

	if m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetFunc == nil {
		panic("mocker: ReplicaApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetFunc is nil but ReplicaApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet was called.")
	}

	call := struct {
		Ctx         context.Context
		ClusterId   string
		TopicName   string
		PartitionId int32
	}{
		Ctx:         ctx,
		ClusterId:   clusterId,
		TopicName:   topicName,
		PartitionId: partitionId,
	}

	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet = append(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet, call)

	return m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetFunc(ctx, clusterId, topicName, partitionId)
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetCalled returns true if ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet was called at least once.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetCalled() bool {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet) > 0
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetCalls returns the calls made to ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.
func (m *ReplicaApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGetCalls() []struct {
	Ctx         context.Context
	ClusterId   string
	TopicName   string
	PartitionId int32
} {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet
}

// Reset resets the calls made to the mocked methods.
func (m *ReplicaApi) Reset() {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Lock()
	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet = nil
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasBrokerIdGet.Unlock()
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Lock()
	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet = nil
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicasGet.Unlock()
}
