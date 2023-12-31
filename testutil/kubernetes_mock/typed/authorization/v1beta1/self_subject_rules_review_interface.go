// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/authorization/v1beta1"
)

// SelfSubjectRulesReviewInterface is an autogenerated mock type for the SelfSubjectRulesReviewInterface type
type SelfSubjectRulesReviewInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, selfSubjectRulesReview, opts
func (_m *SelfSubjectRulesReviewInterface) Create(ctx context.Context, selfSubjectRulesReview *v1beta1.SelfSubjectRulesReview, opts v1.CreateOptions) (*v1beta1.SelfSubjectRulesReview, error) {
	ret := _m.Called(ctx, selfSubjectRulesReview, opts)

	var r0 *v1beta1.SelfSubjectRulesReview
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.SelfSubjectRulesReview, v1.CreateOptions) *v1beta1.SelfSubjectRulesReview); ok {
		r0 = rf(ctx, selfSubjectRulesReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.SelfSubjectRulesReview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.SelfSubjectRulesReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, selfSubjectRulesReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSelfSubjectRulesReviewInterface creates a new instance of SelfSubjectRulesReviewInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewSelfSubjectRulesReviewInterface(t testing.TB) *SelfSubjectRulesReviewInterface {
	mock := &SelfSubjectRulesReviewInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
