package repository

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"crud_api_go/brands/domain"
	"crud_api_go/brands/entity"
	"crud_api_go/common"
)

type BrandsRepository struct {
	conn *sql.DB
}

// List Query
const (
	Create = `INSERT INTO "brands" (name,logo,banner) VALUES ($1,$2,$3)`
	List   = `SELECT id,name,logo,banner,created_at,updated_at FROM "brands" ORDER BY rating DESC`
	Detail = `SELECT id,name,logo,banner,created_at,updated_at FROM "brands" WHERE id = $1 LIMIT 1`
	Update = `UPDATE "brands" SET name = $1, logo = $2, banner = $3, updated_at = $4 WHERE id = $5`
	Delete = `DELETE FROM "brands" WHERE id = $1`
)

// NewBrandsRepository :
func NewBrandsRepository(db *sql.DB) domain.BrandsRepository {
	return &BrandsRepository{
		conn: db,
	}
}

// Create
func (r *BrandsRepository) Create(ctx context.Context, brand *entity.Brands) error {
	if err := r.conn.QueryRowContext(ctx, Create, brand.Name, brand.Logo, brand.Banner).Err(); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// List
func (r *BrandsRepository) List(ctx context.Context) ([]*entity.Brands, error) {
	rows, err := r.conn.QueryContext(ctx, List)
	if err != nil {
		return nil, common.ErrorRequest(err, http.StatusInternalServerError)
	}

	brands := []*entity.Brands{}
	for rows.Next() {

		brand := entity.Brands{}
		var CreatedAt sql.NullTime
		var UpdatedAt sql.NullTime

		if err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.Logo,
			&brand.Banner,
			&CreatedAt,
			&UpdatedAt,
		); err != nil {
			return nil, common.ErrorRequest(err, http.StatusInternalServerError)
		}

		brand.CreatedAt = CreatedAt.Time
		brand.UpdateAt = UpdatedAt.Time

		brands = append(brands, &brand)
	}
	return brands, nil
}

// Detail
func (r *BrandsRepository) Detail(ctx context.Context, brand *entity.Brands) (*entity.Brands, error) {
	result := entity.Brands{}
	var CreatedAt sql.NullTime
	var UpdatedAt sql.NullTime

	if err := r.conn.QueryRowContext(ctx, Detail, brand.ID).Scan(
		&result.ID,
		&result.Name,
		&result.Logo,
		&result.Banner,
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
func (r *BrandsRepository) Update(ctx context.Context, brand *entity.Brands) error {

	//Check if exist brand
	if _, err := r.Detail(ctx, brand); err != nil {
		return err
	}

	if _, err := r.conn.ExecContext(ctx, Update,
		brand.Name,
		brand.Logo,
		brand.Banner,
		time.Now(),
		brand.ID); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// Delete
func (r *BrandsRepository) Delete(ctx context.Context, brand *entity.Brands) error {

	//Check if exist brand
	if _, err := r.Detail(ctx, brand); err != nil {
		return err
	}

	if err := r.conn.QueryRowContext(ctx, Delete, brand.ID).Err(); err != nil {
		if err == sql.ErrNoRows {
			return common.ErrorRequest(err, http.StatusNotFound)
		}
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}
