// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// ReplicaSetExpansion is an autogenerated mock type for the ReplicaSetExpansion type
type ReplicaSetExpansion struct {
	mock.Mock
}

// NewReplicaSetExpansion creates a new instance of ReplicaSetExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewReplicaSetExpansion(t testing.TB) *ReplicaSetExpansion {
	mock := &ReplicaSetExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
