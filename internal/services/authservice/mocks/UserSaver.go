// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "mzhn/auth/internal/domain/dto"
	entity "mzhn/auth/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserSaver is an autogenerated mock type for the UserSaver type
type UserSaver struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, user
func (_m *UserSaver) Save(ctx context.Context, user *dto.CreateUser) (*entity.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateUser) (*entity.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateUser) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateUser) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserSaver creates a new instance of UserSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserSaver(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserSaver {
	mock := &UserSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}