// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_broker_replica_exclusion.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// BrokerReplicaExclusionApi is a mock of BrokerReplicaExclusionApi interface
type BrokerReplicaExclusionApi struct {
	lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet sync.Mutex
	ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetFunc func(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerReplicaExclusionData, *net_http.Response, error)

	lockClustersClusterIdBrokerReplicaExclusionsGet sync.Mutex
	ClustersClusterIdBrokerReplicaExclusionsGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerReplicaExclusionDataList, *net_http.Response, error)

	lockClustersClusterIdBrokerReplicaExclusionscreatePost sync.Mutex
	ClustersClusterIdBrokerReplicaExclusionscreatePostFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionscreatePostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AlterBrokerReplicaExclusionDataList, *net_http.Response, error)

	lockClustersClusterIdBrokerReplicaExclusionsdeletePost sync.Mutex
	ClustersClusterIdBrokerReplicaExclusionsdeletePostFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AlterBrokerReplicaExclusionDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
		}
		ClustersClusterIdBrokerReplicaExclusionsGet []struct {
			Ctx       context.Context
			ClusterId string
		}
		ClustersClusterIdBrokerReplicaExclusionscreatePost []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionscreatePostOpts
		}
		ClustersClusterIdBrokerReplicaExclusionsdeletePost []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts
		}
	}
}

// ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet mocks base method by wrapping the associated func.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerReplicaExclusionData, *net_http.Response, error) {
	m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Unlock()

	if m.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetFunc == nil {
		panic("mocker: BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetFunc is nil but BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet was called.")
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

	m.calls.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet = append(m.calls.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet, call)

	return m.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetFunc(ctx, clusterId, brokerId)
}

// ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetCalled returns true if ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet was called at least once.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetCalled() bool {
	m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet) > 0
}

// ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetCalls returns the calls made to ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsBrokerIdGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
} {
	m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Unlock()

	return m.calls.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet
}

// ClustersClusterIdBrokerReplicaExclusionsGet mocks base method by wrapping the associated func.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerReplicaExclusionDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokerReplicaExclusionsGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsGet.Unlock()

	if m.ClustersClusterIdBrokerReplicaExclusionsGetFunc == nil {
		panic("mocker: BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsGetFunc is nil but BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdBrokerReplicaExclusionsGet = append(m.calls.ClustersClusterIdBrokerReplicaExclusionsGet, call)

	return m.ClustersClusterIdBrokerReplicaExclusionsGetFunc(ctx, clusterId)
}

// ClustersClusterIdBrokerReplicaExclusionsGetCalled returns true if ClustersClusterIdBrokerReplicaExclusionsGet was called at least once.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsGetCalled() bool {
	m.lockClustersClusterIdBrokerReplicaExclusionsGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokerReplicaExclusionsGet) > 0
}

// ClustersClusterIdBrokerReplicaExclusionsGetCalls returns the calls made to ClustersClusterIdBrokerReplicaExclusionsGet.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdBrokerReplicaExclusionsGet.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsGet.Unlock()

	return m.calls.ClustersClusterIdBrokerReplicaExclusionsGet
}

// ClustersClusterIdBrokerReplicaExclusionscreatePost mocks base method by wrapping the associated func.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionscreatePost(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionscreatePostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AlterBrokerReplicaExclusionDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Unlock()

	if m.ClustersClusterIdBrokerReplicaExclusionscreatePostFunc == nil {
		panic("mocker: BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionscreatePostFunc is nil but BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionscreatePost was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionscreatePostOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.ClustersClusterIdBrokerReplicaExclusionscreatePost = append(m.calls.ClustersClusterIdBrokerReplicaExclusionscreatePost, call)

	return m.ClustersClusterIdBrokerReplicaExclusionscreatePostFunc(ctx, clusterId, localVarOptionals)
}

// ClustersClusterIdBrokerReplicaExclusionscreatePostCalled returns true if ClustersClusterIdBrokerReplicaExclusionscreatePost was called at least once.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionscreatePostCalled() bool {
	m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Unlock()

	return len(m.calls.ClustersClusterIdBrokerReplicaExclusionscreatePost) > 0
}

// ClustersClusterIdBrokerReplicaExclusionscreatePostCalls returns the calls made to ClustersClusterIdBrokerReplicaExclusionscreatePost.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionscreatePostCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionscreatePostOpts
} {
	m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Unlock()

	return m.calls.ClustersClusterIdBrokerReplicaExclusionscreatePost
}

// ClustersClusterIdBrokerReplicaExclusionsdeletePost mocks base method by wrapping the associated func.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsdeletePost(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AlterBrokerReplicaExclusionDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Unlock()

	if m.ClustersClusterIdBrokerReplicaExclusionsdeletePostFunc == nil {
		panic("mocker: BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsdeletePostFunc is nil but BrokerReplicaExclusionApi.ClustersClusterIdBrokerReplicaExclusionsdeletePost was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.ClustersClusterIdBrokerReplicaExclusionsdeletePost = append(m.calls.ClustersClusterIdBrokerReplicaExclusionsdeletePost, call)

	return m.ClustersClusterIdBrokerReplicaExclusionsdeletePostFunc(ctx, clusterId, localVarOptionals)
}

// ClustersClusterIdBrokerReplicaExclusionsdeletePostCalled returns true if ClustersClusterIdBrokerReplicaExclusionsdeletePost was called at least once.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsdeletePostCalled() bool {
	m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Unlock()

	return len(m.calls.ClustersClusterIdBrokerReplicaExclusionsdeletePost) > 0
}

// ClustersClusterIdBrokerReplicaExclusionsdeletePostCalls returns the calls made to ClustersClusterIdBrokerReplicaExclusionsdeletePost.
func (m *BrokerReplicaExclusionApi) ClustersClusterIdBrokerReplicaExclusionsdeletePostCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ClustersClusterIdBrokerReplicaExclusionsdeletePostOpts
} {
	m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Lock()
	defer m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Unlock()

	return m.calls.ClustersClusterIdBrokerReplicaExclusionsdeletePost
}

// Reset resets the calls made to the mocked methods.
func (m *BrokerReplicaExclusionApi) Reset() {
	m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Lock()
	m.calls.ClustersClusterIdBrokerReplicaExclusionsBrokerIdGet = nil
	m.lockClustersClusterIdBrokerReplicaExclusionsBrokerIdGet.Unlock()
	m.lockClustersClusterIdBrokerReplicaExclusionsGet.Lock()
	m.calls.ClustersClusterIdBrokerReplicaExclusionsGet = nil
	m.lockClustersClusterIdBrokerReplicaExclusionsGet.Unlock()
	m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Lock()
	m.calls.ClustersClusterIdBrokerReplicaExclusionscreatePost = nil
	m.lockClustersClusterIdBrokerReplicaExclusionscreatePost.Unlock()
	m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Lock()
	m.calls.ClustersClusterIdBrokerReplicaExclusionsdeletePost = nil
	m.lockClustersClusterIdBrokerReplicaExclusionsdeletePost.Unlock()
}
