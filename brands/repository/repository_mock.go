package repository

import (
	"context"
	"errors"

	"crud_api_go/brands/entity"
	"github.com/stretchr/testify/mock"
)

type BrandsRepositoryMock struct {
	Mock mock.Mock
}

// Create :
func (r *BrandsRepositoryMock) Create(ctx context.Context, brands *entity.Brands) error {
	args := r.Mock.Called(brands)
	if args.Get(0) == nil {
		return errors.New("Failed create a brand")
	}
	return nil
}

// List :
func (r *BrandsRepositoryMock) List(ctx context.Context) ([]*entity.Brands, error) {
	return nil, nil
}

// Detail :
func (r *BrandsRepositoryMock) Detail(ctx context.Context, brands *entity.Brands) (*entity.Brands, error) {
	args := r.Mock.Called(brands)
	if args.Get(0) == nil {
		return nil, errors.New("Failed create a brand")
	}

	fetchBrand := args.Get(0).(*entity.Brands)
	if fetchBrand.ID != brands.ID {
		return nil, errors.New("Brand not found")
	}
	return fetchBrand, nil
}

// Update :
func (r *BrandsRepositoryMock) Update(ctx context.Context, brands *entity.Brands) error {
	_, err := r.Detail(ctx, brands)
	if err != nil {
		return err
	}

	args := r.Mock.Called(brands)
	if args.Get(0) == nil {
		return errors.New("Failed update a brand")
	}
	return nil
}

// Delete :
func (r *BrandsRepositoryMock) Delete(ctx context.Context, brands *entity.Brands) error {
	_, err := r.Detail(ctx, brands)
	if err != nil {
		return err
	}

	args := r.Mock.Called(brands)
	if args.Get(0) == nil {
		return errors.New("Failed delete a brand")
	}

	return nil
}
