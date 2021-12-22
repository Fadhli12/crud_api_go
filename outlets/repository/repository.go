package repository

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"crud_api_go/common"
	"crud_api_go/outlets/domain"
	"crud_api_go/outlets/entity"
)

type OutletsRepository struct {
	conn *sql.DB
}

// List Query
const (
	Create = `INSERT INTO "outlets" (name,picture,address,longitude,latitude,brand_id) VALUES ($1,$2,$3,$4,$5,$6)`
	List   = `SELECT id,name,picture,address,longitude,latitude,brand_id,created_at,updated_at FROM "outlets" ORDER BY created_at DESC`
	Detail = `SELECT id,name,picture,address,longitude,latitude,brand_id,created_at,updated_at FROM "outlets" WHERE id = $1 LIMIT 1`
	Update = `UPDATE "outlets" SET name = $1, picture = $2, address = $3, longitude = $4, latitude = $5, brand_id = $6, updated_at = $7 WHERE id = $8`
	Delete = `DELETE FROM "outlets" WHERE id = $1`
)

// NewOutletsRepository :
func NewOutletsRepository(db *sql.DB) domain.OutletsRepository {
	return &OutletsRepository{
		conn: db,
	}
}

// Create
func (r *OutletsRepository) Create(ctx context.Context, outlet *entity.Outlets) error {
	if err := r.conn.QueryRowContext(ctx, Create, outlet.Name, outlet.Picture, outlet.Address, outlet.Longitude, outlet.Latitude, outlet.BrandId).Err(); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// List
func (r *OutletsRepository) List(ctx context.Context) ([]*entity.Outlets, error) {
	rows, err := r.conn.QueryContext(ctx, List)
	if err != nil {
		return nil, common.ErrorRequest(err, http.StatusInternalServerError)
	}

	outlets := []*entity.Outlets{}
	for rows.Next() {

		outlet := entity.Outlets{}

		var CreatedAt sql.NullTime
		var UpdatedAt sql.NullTime

		if err := rows.Scan(
			&outlet.ID,
			&outlet.Name,
			&outlet.Picture,
			&outlet.Address,
			&outlet.Longitude,
			&outlet.Latitude,
			&outlet.BrandId,
			&CreatedAt,
			&UpdatedAt,
		); err != nil {
			return nil, common.ErrorRequest(err, http.StatusInternalServerError)
		}

		outlet.CreatedAt = CreatedAt.Time
		outlet.UpdateAt = UpdatedAt.Time

		outlets = append(outlets, &outlet)
	}
	return outlets, nil
}

// Detail
func (r *OutletsRepository) Detail(ctx context.Context, outlet *entity.Outlets) (*entity.Outlets, error) {
	result := entity.Outlets{}
	var CreatedAt sql.NullTime
	var UpdatedAt sql.NullTime

	if err := r.conn.QueryRowContext(ctx, Detail, outlet.ID).Scan(
		&result.ID,
		&result.Name,
		&result.Picture,
		&result.Address,
		&result.Longitude,
		&result.Latitude,
		&result.BrandId,
		&CreatedAt,
		&UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorRequest(err, http.StatusNotFound)
		}
		return nil, common.ErrorRequest(err, http.StatusInternalServerError)
	}

	result.CreatedAt = CreatedAt.Time
	result.UpdateAt = UpdatedAt.Time

	return &result, nil
}

// Update
func (r *OutletsRepository) Update(ctx context.Context, outlet *entity.Outlets) error {

	//Check if exist outlet
	if _, err := r.Detail(ctx, outlet); err != nil {
		return err
	}

	if _, err := r.conn.ExecContext(ctx, Update,
		&outlet.Name,
		&outlet.Picture,
		&outlet.Address,
		&outlet.Longitude,
		&outlet.Latitude,
		&outlet.BrandId,
		time.Now(),
		outlet.ID); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// Delete
func (r *OutletsRepository) Delete(ctx context.Context, outlet *entity.Outlets) error {

	//Check if exist outlet
	if _, err := r.Detail(ctx, outlet); err != nil {
		return err
	}

	if err := r.conn.QueryRowContext(ctx, Delete, outlet.ID).Err(); err != nil {
		if err == sql.ErrNoRows {
			return common.ErrorRequest(err, http.StatusNotFound)
		}
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}
