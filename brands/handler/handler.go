package handler

import (
	"context"
	"errors"
	"net/http"

	"crud_api_go/brands/domain"
	"crud_api_go/brands/entity"
	"crud_api_go/common"
)

type BrandsHandler struct {
	Repository domain.BrandsRepository
}

// NewBrandsHandler :
func NewBrandsHandler(brandRepository domain.BrandsRepository) domain.BrandsHandler {
	return &BrandsHandler{
		Repository: brandRepository,
	}
}

// Create :
func (h *BrandsHandler) Create(ctx context.Context, brand *entity.Brands) error {

	//Validation
	if brand.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if brand.Logo == "" {
		return common.ErrorRequest(errors.New(common.LogoRequired), http.StatusBadRequest)
	} else if brand.Banner == "" {
		return common.ErrorRequest(errors.New(common.BannerRequired), http.StatusBadRequest)
	}

	//Create brand
	if err := h.Repository.Create(ctx, brand); err != nil {
		return err
	}

	return nil
}

// List :
func (h *BrandsHandler) List(ctx context.Context) ([]*entity.Brands, error) {

	//Get list brands
	brands, err := h.Repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return brands, nil
}

// Detail :
func (h *BrandsHandler) Detail(ctx context.Context, brand *entity.Brands) (*entity.Brands, error) {

	//Get detail brand
	brand, err := h.Repository.Detail(ctx, brand)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

// Update :
func (h *BrandsHandler) Update(ctx context.Context, brand *entity.Brands) error {

	//Validation
	if brand.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if brand.Logo == "" {
		return common.ErrorRequest(errors.New(common.LogoRequired), http.StatusBadRequest)
	} else if brand.Banner == "" {
		return common.ErrorRequest(errors.New(common.BannerRequired), http.StatusBadRequest)
	}

	//Update brand
	if err := h.Repository.Update(ctx, brand); err != nil {
		return err
	}
	return nil
}

// Delete :
func (h *BrandsHandler) Delete(ctx context.Context, brand *entity.Brands) error {

	//Delete brand
	if err := h.Repository.Delete(ctx, brand); err != nil {
		return err
	}
	return nil
}
