package handler

import (
	"context"
	"errors"
	"net/http"

	"crud_api_go/common"
	"crud_api_go/outlets/domain"
	"crud_api_go/outlets/entity"
)

type OutletsHandler struct {
	Repository domain.OutletsRepository
}

// NewOutletsHandler :
func NewOutletsHandler(outletRepository domain.OutletsRepository) domain.OutletsHandler {
	return &OutletsHandler{
		Repository: outletRepository,
	}
}

// Create :
func (h *OutletsHandler) Create(ctx context.Context, outlet *entity.Outlets) error {

	//Validation
	if outlet.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if outlet.Picture == "" {
		return common.ErrorRequest(errors.New(common.PictureRequired), http.StatusBadRequest)
	} else if outlet.Address == "" {
		return common.ErrorRequest(errors.New(common.AddressRequired), http.StatusBadRequest)
	} else if outlet.Longitude == "" {
		return common.ErrorRequest(errors.New(common.LongitudeRequired), http.StatusBadRequest)
	} else if outlet.Latitude == "" {
		return common.ErrorRequest(errors.New(common.LatitudeRequired), http.StatusBadRequest)
	} else if outlet.BrandId == 0 {
		return common.ErrorRequest(errors.New(common.BrandIdRequired), http.StatusBadRequest)
	}

	//Create outlet
	if err := h.Repository.Create(ctx, outlet); err != nil {
		return err
	}

	return nil
}

// List :
func (h *OutletsHandler) List(ctx context.Context) ([]*entity.Outlets, error) {

	//Get list outlets
	outlets, err := h.Repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return outlets, nil
}

// Detail :
func (h *OutletsHandler) Detail(ctx context.Context, outlet *entity.Outlets) (*entity.Outlets, error) {

	//Get detail outlet
	outlet, err := h.Repository.Detail(ctx, outlet)
	if err != nil {
		return nil, err
	}
	return outlet, nil
}

// Update :
func (h *OutletsHandler) Update(ctx context.Context, outlet *entity.Outlets) error {

	//Validation
	if outlet.Name == "" {
		return common.ErrorRequest(errors.New(common.NameRequired), http.StatusBadRequest)
	} else if outlet.Picture == "" {
		return common.ErrorRequest(errors.New(common.PictureRequired), http.StatusBadRequest)
	} else if outlet.Address == "" {
		return common.ErrorRequest(errors.New(common.AddressRequired), http.StatusBadRequest)
	} else if outlet.Longitude == "" {
		return common.ErrorRequest(errors.New(common.LongitudeRequired), http.StatusBadRequest)
	} else if outlet.Latitude == "" {
		return common.ErrorRequest(errors.New(common.LatitudeRequired), http.StatusBadRequest)
	} else if outlet.BrandId == 0 {
		return common.ErrorRequest(errors.New(common.BrandIdRequired), http.StatusBadRequest)
	}
	//Update outlet
	if err := h.Repository.Update(ctx, outlet); err != nil {
		return err
	}
	return nil
}

// Delete :
func (h *OutletsHandler) Delete(ctx context.Context, outlet *entity.Outlets) error {

	//Delete outlet
	if err := h.Repository.Delete(ctx, outlet); err != nil {
		return err
	}
	return nil
}
