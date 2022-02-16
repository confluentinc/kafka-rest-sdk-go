// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_aclv3.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// ACLV3Api is a mock of ACLV3Api interface
type ACLV3Api struct {
	lockCreateKafkaAcls sync.Mutex
	CreateKafkaAclsFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.CreateKafkaAclsOpts) (*net_http.Response, error)

	lockDeleteKafkaAcls sync.Mutex
	DeleteKafkaAclsFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.DeleteKafkaAclsOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.InlineResponse200, *net_http.Response, error)

	lockGetKafkaAcls sync.Mutex
	GetKafkaAclsFunc func(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.GetKafkaAclsOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AclDataList, *net_http.Response, error)

	calls struct {
		CreateKafkaAcls []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.CreateKafkaAclsOpts
		}
		DeleteKafkaAcls []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.DeleteKafkaAclsOpts
		}
		GetKafkaAcls []struct {
			Ctx               context.Context
			ClusterId         string
			LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.GetKafkaAclsOpts
		}
	}
}

// CreateKafkaAcls mocks base method by wrapping the associated func.
func (m *ACLV3Api) CreateKafkaAcls(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.CreateKafkaAclsOpts) (*net_http.Response, error) {
	m.lockCreateKafkaAcls.Lock()
	defer m.lockCreateKafkaAcls.Unlock()

	if m.CreateKafkaAclsFunc == nil {
		panic("mocker: ACLV3Api.CreateKafkaAclsFunc is nil but ACLV3Api.CreateKafkaAcls was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.CreateKafkaAclsOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.CreateKafkaAcls = append(m.calls.CreateKafkaAcls, call)

	return m.CreateKafkaAclsFunc(ctx, clusterId, localVarOptionals)
}

// CreateKafkaAclsCalled returns true if CreateKafkaAcls was called at least once.
func (m *ACLV3Api) CreateKafkaAclsCalled() bool {
	m.lockCreateKafkaAcls.Lock()
	defer m.lockCreateKafkaAcls.Unlock()

	return len(m.calls.CreateKafkaAcls) > 0
}

// CreateKafkaAclsCalls returns the calls made to CreateKafkaAcls.
func (m *ACLV3Api) CreateKafkaAclsCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.CreateKafkaAclsOpts
} {
	m.lockCreateKafkaAcls.Lock()
	defer m.lockCreateKafkaAcls.Unlock()

	return m.calls.CreateKafkaAcls
}

// DeleteKafkaAcls mocks base method by wrapping the associated func.
func (m *ACLV3Api) DeleteKafkaAcls(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.DeleteKafkaAclsOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.InlineResponse200, *net_http.Response, error) {
	m.lockDeleteKafkaAcls.Lock()
	defer m.lockDeleteKafkaAcls.Unlock()

	if m.DeleteKafkaAclsFunc == nil {
		panic("mocker: ACLV3Api.DeleteKafkaAclsFunc is nil but ACLV3Api.DeleteKafkaAcls was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.DeleteKafkaAclsOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.DeleteKafkaAcls = append(m.calls.DeleteKafkaAcls, call)

	return m.DeleteKafkaAclsFunc(ctx, clusterId, localVarOptionals)
}

// DeleteKafkaAclsCalled returns true if DeleteKafkaAcls was called at least once.
func (m *ACLV3Api) DeleteKafkaAclsCalled() bool {
	m.lockDeleteKafkaAcls.Lock()
	defer m.lockDeleteKafkaAcls.Unlock()

	return len(m.calls.DeleteKafkaAcls) > 0
}

// DeleteKafkaAclsCalls returns the calls made to DeleteKafkaAcls.
func (m *ACLV3Api) DeleteKafkaAclsCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.DeleteKafkaAclsOpts
} {
	m.lockDeleteKafkaAcls.Lock()
	defer m.lockDeleteKafkaAcls.Unlock()

	return m.calls.DeleteKafkaAcls
}

// GetKafkaAcls mocks base method by wrapping the associated func.
func (m *ACLV3Api) GetKafkaAcls(ctx context.Context, clusterId string, localVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.GetKafkaAclsOpts) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AclDataList, *net_http.Response, error) {
	m.lockGetKafkaAcls.Lock()
	defer m.lockGetKafkaAcls.Unlock()

	if m.GetKafkaAclsFunc == nil {
		panic("mocker: ACLV3Api.GetKafkaAclsFunc is nil but ACLV3Api.GetKafkaAcls was called.")
	}

	call := struct {
		Ctx               context.Context
		ClusterId         string
		LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.GetKafkaAclsOpts
	}{
		Ctx:               ctx,
		ClusterId:         clusterId,
		LocalVarOptionals: localVarOptionals,
	}

	m.calls.GetKafkaAcls = append(m.calls.GetKafkaAcls, call)

	return m.GetKafkaAclsFunc(ctx, clusterId, localVarOptionals)
}

// GetKafkaAclsCalled returns true if GetKafkaAcls was called at least once.
func (m *ACLV3Api) GetKafkaAclsCalled() bool {
	m.lockGetKafkaAcls.Lock()
	defer m.lockGetKafkaAcls.Unlock()

	return len(m.calls.GetKafkaAcls) > 0
}

// GetKafkaAclsCalls returns the calls made to GetKafkaAcls.
func (m *ACLV3Api) GetKafkaAclsCalls() []struct {
	Ctx               context.Context
	ClusterId         string
	LocalVarOptionals *github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.GetKafkaAclsOpts
} {
	m.lockGetKafkaAcls.Lock()
	defer m.lockGetKafkaAcls.Unlock()

	return m.calls.GetKafkaAcls
}

// Reset resets the calls made to the mocked methods.
func (m *ACLV3Api) Reset() {
	m.lockCreateKafkaAcls.Lock()
	m.calls.CreateKafkaAcls = nil
	m.lockCreateKafkaAcls.Unlock()
	m.lockDeleteKafkaAcls.Lock()
	m.calls.DeleteKafkaAcls = nil
	m.lockDeleteKafkaAcls.Unlock()
	m.lockGetKafkaAcls.Lock()
	m.calls.GetKafkaAcls = nil
	m.lockGetKafkaAcls.Unlock()
}
