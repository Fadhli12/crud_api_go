package domain

import (
	"context"

	"crud_api_go/products/entity"
)

// ProductsHandler :
type ProductsHandler interface {
	Create(context.Context, *entity.Products) error
	List(context.Context) ([]*entity.Products, error)
	Detail(context.Context, *entity.Products) (*entity.Products, error)
	Update(context.Context, *entity.Products) error
	Delete(context.Context, *entity.Products) error
}

// ProductsRepository :
type ProductsRepository interface {
	Create(context.Context, *entity.Products) error
	List(context.Context) ([]*entity.Products, error)
	Detail(context.Context, *entity.Products) (*entity.Products, error)
	Update(context.Context, *entity.Products) error
	Delete(context.Context, *entity.Products) error
}
