package domain

import (
	"context"

	"crud_api_go/brands/entity"
)

// BrandsHandler :
type BrandsHandler interface {
	Create(context.Context, *entity.Brands) error
	List(context.Context) ([]*entity.Brands, error)
	Detail(context.Context, *entity.Brands) (*entity.Brands, error)
	Update(context.Context, *entity.Brands) error
	Delete(context.Context, *entity.Brands) error
}

// BrandsRepository :
type BrandsRepository interface {
	Create(context.Context, *entity.Brands) error
	List(context.Context) ([]*entity.Brands, error)
	Detail(context.Context, *entity.Brands) (*entity.Brands, error)
	Update(context.Context, *entity.Brands) error
	Delete(context.Context, *entity.Brands) error
}
