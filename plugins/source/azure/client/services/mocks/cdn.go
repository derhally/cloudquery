// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: ProfilesClient,EndpointsClient,CustomDomainsClient,OriginsClient,OriginGroupsClient,RoutesClient,RuleSetsClient,RulesClient,SecurityPoliciesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cdn "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	gomock "github.com/golang/mock/gomock"
)

// MockProfilesClient is a mock of ProfilesClient interface.
type MockProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfilesClientMockRecorder
}

// MockProfilesClientMockRecorder is the mock recorder for MockProfilesClient.
type MockProfilesClientMockRecorder struct {
	mock *MockProfilesClient
}

// NewMockProfilesClient creates a new mock instance.
func NewMockProfilesClient(ctrl *gomock.Controller) *MockProfilesClient {
	mock := &MockProfilesClient{ctrl: ctrl}
	mock.recorder = &MockProfilesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfilesClient) EXPECT() *MockProfilesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockProfilesClient) List(arg0 context.Context) (cdn.ProfileListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(cdn.ProfileListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProfilesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProfilesClient)(nil).List), arg0)
}

// MockEndpointsClient is a mock of EndpointsClient interface.
type MockEndpointsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointsClientMockRecorder
}

// MockEndpointsClientMockRecorder is the mock recorder for MockEndpointsClient.
type MockEndpointsClientMockRecorder struct {
	mock *MockEndpointsClient
}

// NewMockEndpointsClient creates a new mock instance.
func NewMockEndpointsClient(ctrl *gomock.Controller) *MockEndpointsClient {
	mock := &MockEndpointsClient{ctrl: ctrl}
	mock.recorder = &MockEndpointsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointsClient) EXPECT() *MockEndpointsClientMockRecorder {
	return m.recorder
}

// ListByProfile mocks base method.
func (m *MockEndpointsClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.EndpointListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.EndpointListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByProfile indicates an expected call of ListByProfile.
func (mr *MockEndpointsClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockEndpointsClient)(nil).ListByProfile), arg0, arg1, arg2)
}

// MockCustomDomainsClient is a mock of CustomDomainsClient interface.
type MockCustomDomainsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCustomDomainsClientMockRecorder
}

// MockCustomDomainsClientMockRecorder is the mock recorder for MockCustomDomainsClient.
type MockCustomDomainsClientMockRecorder struct {
	mock *MockCustomDomainsClient
}

// NewMockCustomDomainsClient creates a new mock instance.
func NewMockCustomDomainsClient(ctrl *gomock.Controller) *MockCustomDomainsClient {
	mock := &MockCustomDomainsClient{ctrl: ctrl}
	mock.recorder = &MockCustomDomainsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomDomainsClient) EXPECT() *MockCustomDomainsClientMockRecorder {
	return m.recorder
}

// ListByEndpoint mocks base method.
func (m *MockCustomDomainsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.CustomDomainListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.CustomDomainListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEndpoint indicates an expected call of ListByEndpoint.
func (mr *MockCustomDomainsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockCustomDomainsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

// MockOriginsClient is a mock of OriginsClient interface.
type MockOriginsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOriginsClientMockRecorder
}

// MockOriginsClientMockRecorder is the mock recorder for MockOriginsClient.
type MockOriginsClientMockRecorder struct {
	mock *MockOriginsClient
}

// NewMockOriginsClient creates a new mock instance.
func NewMockOriginsClient(ctrl *gomock.Controller) *MockOriginsClient {
	mock := &MockOriginsClient{ctrl: ctrl}
	mock.recorder = &MockOriginsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOriginsClient) EXPECT() *MockOriginsClientMockRecorder {
	return m.recorder
}

// ListByEndpoint mocks base method.
func (m *MockOriginsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.OriginListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.OriginListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEndpoint indicates an expected call of ListByEndpoint.
func (mr *MockOriginsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockOriginsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

// MockOriginGroupsClient is a mock of OriginGroupsClient interface.
type MockOriginGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOriginGroupsClientMockRecorder
}

// MockOriginGroupsClientMockRecorder is the mock recorder for MockOriginGroupsClient.
type MockOriginGroupsClientMockRecorder struct {
	mock *MockOriginGroupsClient
}

// NewMockOriginGroupsClient creates a new mock instance.
func NewMockOriginGroupsClient(ctrl *gomock.Controller) *MockOriginGroupsClient {
	mock := &MockOriginGroupsClient{ctrl: ctrl}
	mock.recorder = &MockOriginGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOriginGroupsClient) EXPECT() *MockOriginGroupsClientMockRecorder {
	return m.recorder
}

// ListByEndpoint mocks base method.
func (m *MockOriginGroupsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.OriginGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.OriginGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEndpoint indicates an expected call of ListByEndpoint.
func (mr *MockOriginGroupsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockOriginGroupsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

// MockRoutesClient is a mock of RoutesClient interface.
type MockRoutesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesClientMockRecorder
}

// MockRoutesClientMockRecorder is the mock recorder for MockRoutesClient.
type MockRoutesClientMockRecorder struct {
	mock *MockRoutesClient
}

// NewMockRoutesClient creates a new mock instance.
func NewMockRoutesClient(ctrl *gomock.Controller) *MockRoutesClient {
	mock := &MockRoutesClient{ctrl: ctrl}
	mock.recorder = &MockRoutesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesClient) EXPECT() *MockRoutesClientMockRecorder {
	return m.recorder
}

// ListByEndpoint mocks base method.
func (m *MockRoutesClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.RouteListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.RouteListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEndpoint indicates an expected call of ListByEndpoint.
func (mr *MockRoutesClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockRoutesClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

// MockRuleSetsClient is a mock of RuleSetsClient interface.
type MockRuleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRuleSetsClientMockRecorder
}

// MockRuleSetsClientMockRecorder is the mock recorder for MockRuleSetsClient.
type MockRuleSetsClientMockRecorder struct {
	mock *MockRuleSetsClient
}

// NewMockRuleSetsClient creates a new mock instance.
func NewMockRuleSetsClient(ctrl *gomock.Controller) *MockRuleSetsClient {
	mock := &MockRuleSetsClient{ctrl: ctrl}
	mock.recorder = &MockRuleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuleSetsClient) EXPECT() *MockRuleSetsClientMockRecorder {
	return m.recorder
}

// ListByProfile mocks base method.
func (m *MockRuleSetsClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.RuleSetListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.RuleSetListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByProfile indicates an expected call of ListByProfile.
func (mr *MockRuleSetsClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockRuleSetsClient)(nil).ListByProfile), arg0, arg1, arg2)
}

// MockRulesClient is a mock of RulesClient interface.
type MockRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRulesClientMockRecorder
}

// MockRulesClientMockRecorder is the mock recorder for MockRulesClient.
type MockRulesClientMockRecorder struct {
	mock *MockRulesClient
}

// NewMockRulesClient creates a new mock instance.
func NewMockRulesClient(ctrl *gomock.Controller) *MockRulesClient {
	mock := &MockRulesClient{ctrl: ctrl}
	mock.recorder = &MockRulesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRulesClient) EXPECT() *MockRulesClientMockRecorder {
	return m.recorder
}

// ListByRuleSet mocks base method.
func (m *MockRulesClient) ListByRuleSet(arg0 context.Context, arg1, arg2, arg3 string) (cdn.RuleListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByRuleSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.RuleListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByRuleSet indicates an expected call of ListByRuleSet.
func (mr *MockRulesClientMockRecorder) ListByRuleSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByRuleSet", reflect.TypeOf((*MockRulesClient)(nil).ListByRuleSet), arg0, arg1, arg2, arg3)
}

// MockSecurityPoliciesClient is a mock of SecurityPoliciesClient interface.
type MockSecurityPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityPoliciesClientMockRecorder
}

// MockSecurityPoliciesClientMockRecorder is the mock recorder for MockSecurityPoliciesClient.
type MockSecurityPoliciesClientMockRecorder struct {
	mock *MockSecurityPoliciesClient
}

// NewMockSecurityPoliciesClient creates a new mock instance.
func NewMockSecurityPoliciesClient(ctrl *gomock.Controller) *MockSecurityPoliciesClient {
	mock := &MockSecurityPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSecurityPoliciesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityPoliciesClient) EXPECT() *MockSecurityPoliciesClientMockRecorder {
	return m.recorder
}

// ListByProfile mocks base method.
func (m *MockSecurityPoliciesClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.SecurityPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.SecurityPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByProfile indicates an expected call of ListByProfile.
func (mr *MockSecurityPoliciesClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockSecurityPoliciesClient)(nil).ListByProfile), arg0, arg1, arg2)
}