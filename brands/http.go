package brands

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"crud_api_go/brands/domain"
	"crud_api_go/brands/entity"
	"crud_api_go/brands/handler"
	"crud_api_go/brands/repository"
	"crud_api_go/common"
	"github.com/gorilla/mux"
)

var handlerBrands domain.BrandsHandler

// Routes :
func Routes(route *mux.Router, db *sql.DB) {
	handlerBrands = handler.NewBrandsHandler(
		repository.NewBrandsRepository(db),
	)
	route.HandleFunc("/brands", Create).Methods("POST")
	route.HandleFunc("/brands", List).Methods("GET")
	route.HandleFunc("/brands/{id:[0-9]+}", Detail).Methods("GET")
	route.HandleFunc("/brands/{id:[0-9]+}", Update).Methods("PUT")
	route.HandleFunc("/brands/{id:[0-9]+}", Delete).Methods("DELETE")
}

// CreateBrand godoc
// @Summary Create a new brand
// @Description Create a new brand with the input payload
// @Tags brand
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /brands [post]
func Create(w http.ResponseWriter, r *http.Request) {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	brand := entity.Brands{}
	if err := json.Unmarshal(body, &brand); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	if err := handlerBrands.Create(r.Context(), &brand); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// ListBrand godoc
// @Summary List brands
// @Description Returns brand data
// @Tags brand
// @Accept json
// @Produce json
// @Success 200 {object} common.Responses
// @Failure 500 {object} common.Responses
// @Router /brands [get]
func List(w http.ResponseWriter, r *http.Request) {
	brands, err := handlerBrands.List(r.Context())
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: brands,
	})
}

// DetailBrand godoc
// @Summary Detail a brand
// @Description Return detail brand with param id
// @Tags brand
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /brands/{id} [get]
func Detail(w http.ResponseWriter, r *http.Request) {

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	brandID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	brand, err := handlerBrands.Detail(r.Context(), &entity.Brands{
		ID: int(brandID),
	})
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: brand,
	})
}

// UpdateBrand godoc
// @Summary Update a brand
// @Description Update a brand with the input payload and param id
// @Tags brand
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /brands/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	brand := entity.Brands{}
	if err := json.Unmarshal(body, &brand); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	brandID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
	}
	brand.ID = int(brandID)

	if err := handlerBrands.Update(r.Context(), &brand); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// DeleteBrand godoc
// @Summary Delete a brand
// @Description Delete a brand with param id
// @Tags brand
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /brands/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	brandID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	if err := handlerBrands.Delete(r.Context(), &entity.Brands{
		ID: int(brandID),
	}); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}
