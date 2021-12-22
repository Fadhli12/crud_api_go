package domain

import (
	"context"

	"crud_api_go/outlets/entity"
)

// OutletsHandler :
type OutletsHandler interface {
	Create(context.Context, *entity.Outlets) error
	List(context.Context) ([]*entity.Outlets, error)
	Detail(context.Context, *entity.Outlets) (*entity.Outlets, error)
	Update(context.Context, *entity.Outlets) error
	Delete(context.Context, *entity.Outlets) error
}

// OutletsRepository :
type OutletsRepository interface {
	Create(context.Context, *entity.Outlets) error
	List(context.Context) ([]*entity.Outlets, error)
	Detail(context.Context, *entity.Outlets) (*entity.Outlets, error)
	Update(context.Context, *entity.Outlets) error
	Delete(context.Context, *entity.Outlets) error
}
