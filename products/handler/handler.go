package handler

import (
	"context"
	"errors"
	"net/http"

	"crud_api_go/common"
	"crud_api_go/products/domain"
	"crud_api_go/products/entity"
)

type ProductsHandler struct {
	Repository domain.ProductsRepository
}

// NewProductsHandler :
func NewProductsHandler(productRepository domain.ProductsRepository) domain.ProductsHandler {
	return &ProductsHandler{
		Repository: productRepository,
	}
}

// Create :
func (h *ProductsHandler) Create(ctx context.Context, product *entity.Products) error {

	//Validation
	if product.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if product.Picture == "" {
		return common.ErrorRequest(errors.New(common.PictureRequired), http.StatusBadRequest)
	} else if product.Price == 0 {
		return common.ErrorRequest(errors.New(common.PriceRequired), http.StatusBadRequest)
	} else if product.BrandId == 0 {
		return common.ErrorRequest(errors.New(common.BrandIdRequired), http.StatusBadRequest)
	}
	//Create product
	if err := h.Repository.Create(ctx, product); err != nil {
		return err
	}

	return nil
}

// List :
func (h *ProductsHandler) List(ctx context.Context) ([]*entity.Products, error) {

	//Get list products
	products, err := h.Repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Detail :
func (h *ProductsHandler) Detail(ctx context.Context, product *entity.Products) (*entity.Products, error) {

	//Get detail product
	product, err := h.Repository.Detail(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Update :
func (h *ProductsHandler) Update(ctx context.Context, product *entity.Products) error {

	//Validation
	if product.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if product.Picture == "" {
		return common.ErrorRequest(errors.New(common.PictureRequired), http.StatusBadRequest)
	} else if product.Price == 0 {
		return common.ErrorRequest(errors.New(common.PriceRequired), http.StatusBadRequest)
	} else if product.BrandId == 0 {
		return common.ErrorRequest(errors.New(common.BrandIdRequired), http.StatusBadRequest)
	}
	//Update product
	if err := h.Repository.Update(ctx, product); err != nil {
		return err
	}
	return nil
}

// Delete :
func (h *ProductsHandler) Delete(ctx context.Context, product *entity.Products) error {

	//Delete product
	if err := h.Repository.Delete(ctx, product); err != nil {
		return err
	}
	return nil
}
