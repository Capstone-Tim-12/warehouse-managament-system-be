// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	mock "github.com/stretchr/testify/mock"
)

// RegionRepository is an autogenerated mock type for the RegionRepository type
type RegionRepository struct {
	mock.Mock
}

// FindAllProvince provides a mock function with given fields: ctx
func (_m *RegionRepository) FindAllProvince(ctx context.Context) ([]entity.Province, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Province
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Province, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Province); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Province)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDistrictByRegencyId provides a mock function with given fields: ctx, id
func (_m *RegionRepository) FindDistrictByRegencyId(ctx context.Context, id string) ([]entity.District, error) {
	ret := _m.Called(ctx, id)

	var r0 []entity.District
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.District, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.District); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.District)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRegencyByProvinceId provides a mock function with given fields: ctx, id
func (_m *RegionRepository) FindRegencyByProvinceId(ctx context.Context, id string) ([]entity.Regency, error) {
	ret := _m.Called(ctx, id)

	var r0 []entity.Regency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Regency, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Regency); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Regency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindVillageByDistrictId provides a mock function with given fields: ctx, id
func (_m *RegionRepository) FindVillageByDistrictId(ctx context.Context, id string) ([]entity.Village, error) {
	ret := _m.Called(ctx, id)

	var r0 []entity.Village
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Village, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Village); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Village)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDistrictById provides a mock function with given fields: ctx, id
func (_m *RegionRepository) GetDistrictById(ctx context.Context, id string) (*entity.District, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.District
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.District, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.District); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.District)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvinceById provides a mock function with given fields: ctx, id
func (_m *RegionRepository) GetProvinceById(ctx context.Context, id string) (*entity.Province, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Province
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Province, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Province); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Province)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegencyById provides a mock function with given fields: ctx, id
func (_m *RegionRepository) GetRegencyById(ctx context.Context, id string) (*entity.Regency, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Regency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Regency, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Regency); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Regency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVillageById provides a mock function with given fields: ctx, id
func (_m *RegionRepository) GetVillageById(ctx context.Context, id string) (*entity.Village, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Village
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Village, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Village); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Village)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegionRepository creates a new instance of RegionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegionRepository(t mockConstructorTestingTNewRegionRepository) *RegionRepository {
	mock := &RegionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}