package outlets

import (
	"context"
	"testing"

	"crud_api_go/outlets/entity"
	"crud_api_go/outlets/handler"
	"crud_api_go/outlets/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	outletRepo = &repository.OutletsRepositoryMock{
		Mock: mock.Mock{},
	}
	outletHandler = handler.OutletsHandler{
		Repository: outletRepo,
	}
)

func TestDeleteOutlet(t *testing.T) {
	ctx := context.Background()
	outlet := entity.Outlets{
		ID: 1,
	}
	t.Run("Success", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&outlet)
		outletRepo.Mock.On("Delete", &outlet).Return(&entity.Outlets{})
		err := outletHandler.Delete(ctx, &outlet)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&entity.Outlets{
			ID: 2,
		})
		outletRepo.Mock.On("Delete", &outlet).Return(&entity.Outlets{})
		err := outletHandler.Delete(ctx, &outlet)
		assert.NotNil(t, err)
	})
}

func TestCreateOutlet(t *testing.T) {
	ctx := context.Background()
	outlet := entity.Outlets{
		Name:      "This is Outlet",
		Picture:   "#",
		Address:   "address outlet",
		Longitude: "1.000000",
		Latitude:  "1.000000",
		BrandId:   1,
	}
	t.Run("Success", func(t *testing.T) {
		outletRepo.Mock.On("Create", &outlet).Return(&outlet)
		err := outletHandler.Create(ctx, &outlet)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		outlet.Title = ""
		outletRepo.Mock.On("Create", &outlet).Return(&outlet)
		err := outletHandler.Create(ctx, &outlet)
		assert.NotNil(t, err)
	})
}

func TestUpdateOutlet(t *testing.T) {
	ctx := context.Background()
	outlet := entity.Outlets{
		ID:        1,
		Name:      "This is Outlet",
		Picture:   "#",
		Address:   "address outlet",
		Longitude: "1.000000",
		Latitude:  "1.000000",
		BrandId:   1,
	}
	t.Run("Success", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&outlet)
		outletRepo.Mock.On("Update", &outlet).Return(&outlet)
		err := outletHandler.Update(ctx, &outlet)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&entity.Outlets{
			ID: 2,
		})
		outletRepo.Mock.On("Update", &outlet).Return(&outlet)
		err := outletHandler.Update(ctx, &outlet)
		assert.NotNil(t, err)
	})
}

func TestDetailOutlet(t *testing.T) {
	ctx := context.Background()
	outlet := entity.Outlets{
		ID:        1,
		Name:      "This is Outlet",
		Picture:   "#",
		Address:   "address outlet",
		Longitude: "1.000000",
		Latitude:  "1.000000",
		BrandId:   1,
	}
	t.Run("Success", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&outlet)
		result, err := outletHandler.Detail(ctx, &outlet)
		assert.NotNil(t, result)
		assert.Nil(t, err)
	})
	t.Run("Failed", func(t *testing.T) {
		outletRepo.Mock.On("Detail", &outlet).Return(&entity.Outlets{
			ID: 2,
		})
		result, err := outletHandler.Detail(ctx, &outlet)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}
