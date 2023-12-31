// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

// NetworkPoliciesGetter is an autogenerated mock type for the NetworkPoliciesGetter type
type NetworkPoliciesGetter struct {
	mock.Mock
}

// NetworkPolicies provides a mock function with given fields: namespace
func (_m *NetworkPoliciesGetter) NetworkPolicies(namespace string) v1.NetworkPolicyInterface {
	ret := _m.Called(namespace)

	var r0 v1.NetworkPolicyInterface
	if rf, ok := ret.Get(0).(func(string) v1.NetworkPolicyInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.NetworkPolicyInterface)
		}
	}

	return r0
}

// NewNetworkPoliciesGetter creates a new instance of NetworkPoliciesGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkPoliciesGetter(t testing.TB) *NetworkPoliciesGetter {
	mock := &NetworkPoliciesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
