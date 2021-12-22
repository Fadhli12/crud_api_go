package outlets

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"crud_api_go/common"
	"crud_api_go/outlets/domain"
	"crud_api_go/outlets/entity"
	"crud_api_go/outlets/handler"
	"crud_api_go/outlets/repository"
	"github.com/gorilla/mux"
)

var handlerOutlets domain.OutletsHandler

// Routes :
func Routes(route *mux.Router, db *sql.DB) {
	handlerOutlets = handler.NewOutletsHandler(
		repository.NewOutletsRepository(db),
	)
	route.HandleFunc("/outlets", Create).Methods("POST")
	route.HandleFunc("/outlets", List).Methods("GET")
	route.HandleFunc("/outlets/{id:[0-9]+}", Detail).Methods("GET")
	route.HandleFunc("/outlets/{id:[0-9]+}", Update).Methods("PUT")
	route.HandleFunc("/outlets/{id:[0-9]+}", Delete).Methods("DELETE")
}

// CreateOutlet godoc
// @Summary Create a new outlet
// @Description Create a new outlet with the input payload
// @Tags outlet
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /outlets [post]
func Create(w http.ResponseWriter, r *http.Request) {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	outlet := entity.Outlets{}
	if err := json.Unmarshal(body, &outlet); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	if err := handlerOutlets.Create(r.Context(), &outlet); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// ListOutlet godoc
// @Summary List outlets
// @Description Returns outlet data
// @Tags outlet
// @Accept json
// @Produce json
// @Success 200 {object} common.Responses
// @Failure 500 {object} common.Responses
// @Router /outlets [get]
func List(w http.ResponseWriter, r *http.Request) {
	outlets, err := handlerOutlets.List(r.Context())
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: outlets,
	})
}

// DetailOutlet godoc
// @Summary Detail a outlet
// @Description Return detail outlet with param id
// @Tags outlet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /outlets/{id} [get]
func Detail(w http.ResponseWriter, r *http.Request) {

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	outletID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	outlet, err := handlerOutlets.Detail(r.Context(), &entity.Outlets{
		ID: int(outletID),
	})
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: outlet,
	})
}

// UpdateOutlet godoc
// @Summary Update a outlet
// @Description Update a outlet with the input payload and param id
// @Tags outlet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /outlets/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	outlet := entity.Outlets{}
	if err := json.Unmarshal(body, &outlet); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	outletID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
	}
	outlet.ID = int(outletID)

	if err := handlerOutlets.Update(r.Context(), &outlet); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// DeleteOutlet godoc
// @Summary Delete a outlet
// @Description Delete a outlet with param id
// @Tags outlet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /outlets/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	outletID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	if err := handlerOutlets.Delete(r.Context(), &entity.Outlets{
		ID: int(outletID),
	}); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}
