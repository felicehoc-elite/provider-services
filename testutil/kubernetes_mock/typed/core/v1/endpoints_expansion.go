// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// EndpointsExpansion is an autogenerated mock type for the EndpointsExpansion type
type EndpointsExpansion struct {
	mock.Mock
}

// NewEndpointsExpansion creates a new instance of EndpointsExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEndpointsExpansion(t testing.TB) *EndpointsExpansion {
	mock := &EndpointsExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
