// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
)

// CertificatesV1beta1Interface is an autogenerated mock type for the CertificatesV1beta1Interface type
type CertificatesV1beta1Interface struct {
	mock.Mock
}

// CertificateSigningRequests provides a mock function with given fields:
func (_m *CertificatesV1beta1Interface) CertificateSigningRequests() v1beta1.CertificateSigningRequestInterface {
	ret := _m.Called()

	var r0 v1beta1.CertificateSigningRequestInterface
	if rf, ok := ret.Get(0).(func() v1beta1.CertificateSigningRequestInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.CertificateSigningRequestInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *CertificatesV1beta1Interface) RESTClient() rest.Interface {
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

// NewCertificatesV1beta1Interface creates a new instance of CertificatesV1beta1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCertificatesV1beta1Interface(t testing.TB) *CertificatesV1beta1Interface {
	mock := &CertificatesV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
