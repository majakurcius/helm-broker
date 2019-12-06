// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import chart "k8s.io/helm/pkg/proto/hapi/chart"
import internal "github.com/kyma-project/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"
import semver "github.com/Masterminds/semver"

// chartGetter is an autogenerated mock type for the chartGetter type
type chartGetter struct {
	mock.Mock
}

// Get provides a mock function with given fields: namespace, name, ver
func (_m *chartGetter) Get(namespace internal.Namespace, name internal.ChartName, ver semver.Version) (*chart.Chart, error) {
	ret := _m.Called(namespace, name, ver)

	var r0 *chart.Chart
	if rf, ok := ret.Get(0).(func(internal.Namespace, internal.ChartName, semver.Version) *chart.Chart); ok {
		r0 = rf(namespace, name, ver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chart.Chart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Namespace, internal.ChartName, semver.Version) error); ok {
		r1 = rf(namespace, name, ver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
