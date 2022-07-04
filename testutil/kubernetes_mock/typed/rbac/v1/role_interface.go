// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	rbacv1 "k8s.io/api/rbac/v1"

	testing "testing"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/rbac/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// RoleInterface is an autogenerated mock type for the RoleInterface type
type RoleInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Apply(ctx context.Context, role *v1.RoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RoleApplyConfiguration, metav1.ApplyOptions) *rbacv1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.RoleApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Create(ctx context.Context, role *rbacv1.Role, opts metav1.CreateOptions) (*rbacv1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1.Role, metav1.CreateOptions) *rbacv1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1.Role, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *RoleInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *RoleInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *RoleInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*rbacv1.Role, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *rbacv1.Role); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *RoleInterface) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *rbacv1.RoleList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *rbacv1.RoleList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.RoleList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *RoleInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1.Role, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *rbacv1.Role); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, role, opts
func (_m *RoleInterface) Update(ctx context.Context, role *rbacv1.Role, opts metav1.UpdateOptions) (*rbacv1.Role, error) {
	ret := _m.Called(ctx, role, opts)

	var r0 *rbacv1.Role
	if rf, ok := ret.Get(0).(func(context.Context, *rbacv1.Role, metav1.UpdateOptions) *rbacv1.Role); ok {
		r0 = rf(ctx, role, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rbacv1.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rbacv1.Role, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, role, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *RoleInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRoleInterface creates a new instance of RoleInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleInterface(t testing.TB) *RoleInterface {
	mock := &RoleInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
