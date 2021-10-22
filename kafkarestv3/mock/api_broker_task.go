// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_broker_task.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// BrokerTaskApi is a mock of BrokerTaskApi interface
type BrokerTaskApi struct {
	lockClustersClusterIdBrokersBrokerIdTasksGet sync.Mutex
	ClustersClusterIdBrokersBrokerIdTasksGetFunc func(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error)

	lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet sync.Mutex
	ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetFunc func(ctx context.Context, clusterId string, brokerId int32, taskType github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskData, *net_http.Response, error)

	lockClustersClusterIdBrokersTasksGet sync.Mutex
	ClustersClusterIdBrokersTasksGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error)

	lockClustersClusterIdBrokersTasksTaskTypeGet sync.Mutex
	ClustersClusterIdBrokersTasksTaskTypeGetFunc func(ctx context.Context, clusterId string, taskType github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdBrokersBrokerIdTasksGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
		}
		ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
			TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
		}
		ClustersClusterIdBrokersTasksGet []struct {
			Ctx       context.Context
			ClusterId string
		}
		ClustersClusterIdBrokersTasksTaskTypeGet []struct {
			Ctx       context.Context
			ClusterId string
			TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
		}
	}
}

// ClustersClusterIdBrokersBrokerIdTasksGet mocks base method by wrapping the associated func.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksGet(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersBrokerIdTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksGet.Unlock()

	if m.ClustersClusterIdBrokersBrokerIdTasksGetFunc == nil {
		panic("mocker: BrokerTaskApi.ClustersClusterIdBrokersBrokerIdTasksGetFunc is nil but BrokerTaskApi.ClustersClusterIdBrokersBrokerIdTasksGet was called.")
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

	m.calls.ClustersClusterIdBrokersBrokerIdTasksGet = append(m.calls.ClustersClusterIdBrokersBrokerIdTasksGet, call)

	return m.ClustersClusterIdBrokersBrokerIdTasksGetFunc(ctx, clusterId, brokerId)
}

// ClustersClusterIdBrokersBrokerIdTasksGetCalled returns true if ClustersClusterIdBrokersBrokerIdTasksGet was called at least once.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksGetCalled() bool {
	m.lockClustersClusterIdBrokersBrokerIdTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersBrokerIdTasksGet) > 0
}

// ClustersClusterIdBrokersBrokerIdTasksGetCalls returns the calls made to ClustersClusterIdBrokersBrokerIdTasksGet.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
} {
	m.lockClustersClusterIdBrokersBrokerIdTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksGet.Unlock()

	return m.calls.ClustersClusterIdBrokersBrokerIdTasksGet
}

// ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet mocks base method by wrapping the associated func.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet(ctx context.Context, clusterId string, brokerId int32, taskType github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskData, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Unlock()

	if m.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetFunc == nil {
		panic("mocker: BrokerTaskApi.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetFunc is nil but BrokerTaskApi.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		BrokerId  int32
		TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		BrokerId:  brokerId,
		TaskType:  taskType,
	}

	m.calls.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet = append(m.calls.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet, call)

	return m.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetFunc(ctx, clusterId, brokerId, taskType)
}

// ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetCalled returns true if ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet was called at least once.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetCalled() bool {
	m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet) > 0
}

// ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetCalls returns the calls made to ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.
func (m *BrokerTaskApi) ClustersClusterIdBrokersBrokerIdTasksTaskTypeGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
	TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
} {
	m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Unlock()

	return m.calls.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet
}

// ClustersClusterIdBrokersTasksGet mocks base method by wrapping the associated func.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksGet.Unlock()

	if m.ClustersClusterIdBrokersTasksGetFunc == nil {
		panic("mocker: BrokerTaskApi.ClustersClusterIdBrokersTasksGetFunc is nil but BrokerTaskApi.ClustersClusterIdBrokersTasksGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdBrokersTasksGet = append(m.calls.ClustersClusterIdBrokersTasksGet, call)

	return m.ClustersClusterIdBrokersTasksGetFunc(ctx, clusterId)
}

// ClustersClusterIdBrokersTasksGetCalled returns true if ClustersClusterIdBrokersTasksGet was called at least once.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksGetCalled() bool {
	m.lockClustersClusterIdBrokersTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersTasksGet) > 0
}

// ClustersClusterIdBrokersTasksGetCalls returns the calls made to ClustersClusterIdBrokersTasksGet.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdBrokersTasksGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksGet.Unlock()

	return m.calls.ClustersClusterIdBrokersTasksGet
}

// ClustersClusterIdBrokersTasksTaskTypeGet mocks base method by wrapping the associated func.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksTaskTypeGet(ctx context.Context, clusterId string, taskType github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskDataList, *net_http.Response, error) {
	m.lockClustersClusterIdBrokersTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksTaskTypeGet.Unlock()

	if m.ClustersClusterIdBrokersTasksTaskTypeGetFunc == nil {
		panic("mocker: BrokerTaskApi.ClustersClusterIdBrokersTasksTaskTypeGetFunc is nil but BrokerTaskApi.ClustersClusterIdBrokersTasksTaskTypeGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		TaskType:  taskType,
	}

	m.calls.ClustersClusterIdBrokersTasksTaskTypeGet = append(m.calls.ClustersClusterIdBrokersTasksTaskTypeGet, call)

	return m.ClustersClusterIdBrokersTasksTaskTypeGetFunc(ctx, clusterId, taskType)
}

// ClustersClusterIdBrokersTasksTaskTypeGetCalled returns true if ClustersClusterIdBrokersTasksTaskTypeGet was called at least once.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksTaskTypeGetCalled() bool {
	m.lockClustersClusterIdBrokersTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksTaskTypeGet.Unlock()

	return len(m.calls.ClustersClusterIdBrokersTasksTaskTypeGet) > 0
}

// ClustersClusterIdBrokersTasksTaskTypeGetCalls returns the calls made to ClustersClusterIdBrokersTasksTaskTypeGet.
func (m *BrokerTaskApi) ClustersClusterIdBrokersTasksTaskTypeGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TaskType  github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BrokerTaskType
} {
	m.lockClustersClusterIdBrokersTasksTaskTypeGet.Lock()
	defer m.lockClustersClusterIdBrokersTasksTaskTypeGet.Unlock()

	return m.calls.ClustersClusterIdBrokersTasksTaskTypeGet
}

// Reset resets the calls made to the mocked methods.
func (m *BrokerTaskApi) Reset() {
	m.lockClustersClusterIdBrokersBrokerIdTasksGet.Lock()
	m.calls.ClustersClusterIdBrokersBrokerIdTasksGet = nil
	m.lockClustersClusterIdBrokersBrokerIdTasksGet.Unlock()
	m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Lock()
	m.calls.ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet = nil
	m.lockClustersClusterIdBrokersBrokerIdTasksTaskTypeGet.Unlock()
	m.lockClustersClusterIdBrokersTasksGet.Lock()
	m.calls.ClustersClusterIdBrokersTasksGet = nil
	m.lockClustersClusterIdBrokersTasksGet.Unlock()
	m.lockClustersClusterIdBrokersTasksTaskTypeGet.Lock()
	m.calls.ClustersClusterIdBrokersTasksTaskTypeGet = nil
	m.lockClustersClusterIdBrokersTasksTaskTypeGet.Unlock()
}