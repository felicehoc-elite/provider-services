// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// EndpointSliceExpansion is an autogenerated mock type for the EndpointSliceExpansion type
type EndpointSliceExpansion struct {
	mock.Mock
}

// NewEndpointSliceExpansion creates a new instance of EndpointSliceExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointSliceExpansion(t testing.TB) *EndpointSliceExpansion {
	mock := &EndpointSliceExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
