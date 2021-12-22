package repository

import (
	"context"
	"errors"

	"crud_api_go/outlets/entity"
	"github.com/stretchr/testify/mock"
)

type OutletsRepositoryMock struct {
	Mock mock.Mock
}

// Create :
func (r *OutletsRepositoryMock) Create(ctx context.Context, outlets *entity.Outlets) error {
	args := r.Mock.Called(outlets)
	if args.Get(0) == nil {
		return errors.New("Failed create a outlet")
	}
	return nil
}

// List :
func (r *OutletsRepositoryMock) List(ctx context.Context) ([]*entity.Outlets, error) {
	return nil, nil
}

// Detail :
func (r *OutletsRepositoryMock) Detail(ctx context.Context, outlets *entity.Outlets) (*entity.Outlets, error) {
	args := r.Mock.Called(outlets)
	if args.Get(0) == nil {
		return nil, errors.New("Failed create a outlet")
	}

	fetchOutlet := args.Get(0).(*entity.Outlets)
	if fetchOutlet.ID != outlets.ID {
		return nil, errors.New("Outlet not found")
	}
	return fetchOutlet, nil
}

// Update :
func (r *OutletsRepositoryMock) Update(ctx context.Context, outlets *entity.Outlets) error {
	_, err := r.Detail(ctx, outlets)
	if err != nil {
		return err
	}

	args := r.Mock.Called(outlets)
	if args.Get(0) == nil {
		return errors.New("Failed update a outlet")
	}
	return nil
}

// Delete :
func (r *OutletsRepositoryMock) Delete(ctx context.Context, outlets *entity.Outlets) error {
	_, err := r.Detail(ctx, outlets)
	if err != nil {
		return err
	}

	args := r.Mock.Called(outlets)
	if args.Get(0) == nil {
		return errors.New("Failed delete a outlet")
	}

	return nil
}
