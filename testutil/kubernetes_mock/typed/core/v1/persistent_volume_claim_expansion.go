// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// PersistentVolumeClaimExpansion is an autogenerated mock type for the PersistentVolumeClaimExpansion type
type PersistentVolumeClaimExpansion struct {
	mock.Mock
}

// NewPersistentVolumeClaimExpansion creates a new instance of PersistentVolumeClaimExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersistentVolumeClaimExpansion(t testing.TB) *PersistentVolumeClaimExpansion {
	mock := &PersistentVolumeClaimExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
