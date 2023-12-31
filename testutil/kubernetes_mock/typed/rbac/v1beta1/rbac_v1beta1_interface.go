// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
)

// RbacV1beta1Interface is an autogenerated mock type for the RbacV1beta1Interface type
type RbacV1beta1Interface struct {
	mock.Mock
}

// ClusterRoleBindings provides a mock function with given fields:
func (_m *RbacV1beta1Interface) ClusterRoleBindings() v1beta1.ClusterRoleBindingInterface {
	ret := _m.Called()

	var r0 v1beta1.ClusterRoleBindingInterface
	if rf, ok := ret.Get(0).(func() v1beta1.ClusterRoleBindingInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ClusterRoleBindingInterface)
		}
	}

	return r0
}

// ClusterRoles provides a mock function with given fields:
func (_m *RbacV1beta1Interface) ClusterRoles() v1beta1.ClusterRoleInterface {
	ret := _m.Called()

	var r0 v1beta1.ClusterRoleInterface
	if rf, ok := ret.Get(0).(func() v1beta1.ClusterRoleInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.ClusterRoleInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *RbacV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// RoleBindings provides a mock function with given fields: namespace
func (_m *RbacV1beta1Interface) RoleBindings(namespace string) v1beta1.RoleBindingInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.RoleBindingInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.RoleBindingInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.RoleBindingInterface)
		}
	}

	return r0
}

// Roles provides a mock function with given fields: namespace
func (_m *RbacV1beta1Interface) Roles(namespace string) v1beta1.RoleInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.RoleInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.RoleInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.RoleInterface)
		}
	}

	return r0
}

// NewRbacV1beta1Interface creates a new instance of RbacV1beta1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRbacV1beta1Interface(t testing.TB) *RbacV1beta1Interface {
	mock := &RbacV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
