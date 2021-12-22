package brands

import (
	"context"
	"testing"

	"crud_api_go/brands/entity"
	"crud_api_go/brands/handler"
	"crud_api_go/brands/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	brandRepo = &repository.BrandsRepositoryMock{
		Mock: mock.Mock{},
	}
	brandHandler = handler.BrandsHandler{
		Repository: brandRepo,
	}
)

func TestDeleteBrand(t *testing.T) {
	ctx := context.Background()
	brand := entity.Brands{
		ID: 1,
	}
	t.Run("Success", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&brand)
		brandRepo.Mock.On("Delete", &brand).Return(&entity.Brands{})
		err := brandHandler.Delete(ctx, &brand)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&entity.Brands{
			ID: 2,
		})
		brandRepo.Mock.On("Delete", &brand).Return(&entity.Brands{})
		err := brandHandler.Delete(ctx, &brand)
		assert.NotNil(t, err)
	})
}

func TestCreateBrand(t *testing.T) {
	ctx := context.Background()
	brand := entity.Brands{
		Name:   "This is Brand",
		Logo:   "#",
		Banner: "#",
	}
	t.Run("Success", func(t *testing.T) {
		brandRepo.Mock.On("Create", &brand).Return(&brand)
		err := brandHandler.Create(ctx, &brand)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		brand.Title = ""
		brandRepo.Mock.On("Create", &brand).Return(&brand)
		err := brandHandler.Create(ctx, &brand)
		assert.NotNil(t, err)
	})
}

func TestUpdateBrand(t *testing.T) {
	ctx := context.Background()
	brand := entity.Brands{
		ID:     1,
		Name:   "This is Brand",
		Logo:   "#",
		Banner: "#",
	}
	t.Run("Success", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&brand)
		brandRepo.Mock.On("Update", &brand).Return(&brand)
		err := brandHandler.Update(ctx, &brand)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&entity.Brands{
			ID: 2,
		})
		brandRepo.Mock.On("Update", &brand).Return(&brand)
		err := brandHandler.Update(ctx, &brand)
		assert.NotNil(t, err)
	})
}

func TestDetailBrand(t *testing.T) {
	ctx := context.Background()
	brand := entity.Brands{
		ID:     1,
		Name:   "This is Brand",
		Logo:   "#",
		Banner: "#",
	}
	t.Run("Success", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&brand)
		result, err := brandHandler.Detail(ctx, &brand)
		assert.NotNil(t, result)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		brandRepo.Mock.On("Detail", &brand).Return(&entity.Brands{
			ID: 2,
		})
		result, err := brandHandler.Detail(ctx, &brand)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
