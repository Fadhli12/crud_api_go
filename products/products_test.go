package products

import (
	"context"
	"testing"

	"crud_api_go/products/entity"
	"crud_api_go/products/handler"
	"crud_api_go/products/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productRepo = &repository.ProductsRepositoryMock{
		Mock: mock.Mock{},
	}
	productHandler = handler.ProductsHandler{
		Repository: productRepo,
	}
)

func TestDeleteProduct(t *testing.T) {
	ctx := context.Background()
	product := entity.Products{
		ID: 1,
	}
	t.Run("Success", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&product)
		productRepo.Mock.On("Delete", &product).Return(&entity.Products{})
		err := productHandler.Delete(ctx, &product)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&entity.Products{
			ID: 2,
		})
		productRepo.Mock.On("Delete", &product).Return(&entity.Products{})
		err := productHandler.Delete(ctx, &product)
		assert.NotNil(t, err)
	})
}

func TestCreateProduct(t *testing.T) {
	ctx := context.Background()
	product := entity.Products{
		Name:    "This is Product",
		Picture: "#",
		Price:   1000,
		BrandId: 1,
	}
	t.Run("Success", func(t *testing.T) {
		productRepo.Mock.On("Create", &product).Return(&product)
		err := productHandler.Create(ctx, &product)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		product.Title = ""
		productRepo.Mock.On("Create", &product).Return(&product)
		err := productHandler.Create(ctx, &product)
		assert.NotNil(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	ctx := context.Background()
	product := entity.Products{
		ID:      1,
		Name:    "This is Product",
		Picture: "#",
		Price:   1000,
		BrandId: 1,
	}
	t.Run("Success", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&product)
		productRepo.Mock.On("Update", &product).Return(&product)
		err := productHandler.Update(ctx, &product)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&entity.Products{
			ID: 2,
		})
		productRepo.Mock.On("Update", &product).Return(&product)
		err := productHandler.Update(ctx, &product)
		assert.NotNil(t, err)
	})
}

func TestDetailProduct(t *testing.T) {
	ctx := context.Background()
	product := entity.Products{
		ID:      1,
		Name:    "This is Product",
		Picture: "#",
		Price:   1000,
		BrandId: 1,
	}
	t.Run("Success", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&product)
		result, err := productHandler.Detail(ctx, &product)
		assert.NotNil(t, result)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		productRepo.Mock.On("Detail", &product).Return(&entity.Products{
			ID: 2,
		})
		result, err := productHandler.Detail(ctx, &product)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
