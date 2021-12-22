package repository

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"crud_api_go/common"
	"crud_api_go/products/domain"
	"crud_api_go/products/entity"
)

type ProductsRepository struct {
	conn *sql.DB
}

// List Query
const (
	Create = `INSERT INTO "products" (name,picture,price,brand_id) VALUES ($1,$2,$3,$4)`
	List   = `SELECT id,name,picture,price,brand_id,created_at,updated_at FROM "products" ORDER BY created_at DESC`
	Detail = `SELECT id,name,picture,price,brand_id,created_at,updated_at FROM "products" WHERE id = $1 LIMIT 1`
	Update = `UPDATE "products" SET name = $1, picture = $2, price = $3, brand_id = $4, updated_at = $5 WHERE id = $6`
	Delete = `DELETE FROM "products" WHERE id = $1`
)

// NewProductsRepository :
func NewProductsRepository(db *sql.DB) domain.ProductsRepository {
	return &ProductsRepository{
		conn: db,
	}
}

// Create
func (r *ProductsRepository) Create(ctx context.Context, product *entity.Products) error {
	if err := r.conn.QueryRowContext(ctx, Create, product.Name, product.Picture, product.Price, product.BrandId).Err(); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// List
func (r *ProductsRepository) List(ctx context.Context) ([]*entity.Products, error) {
	rows, err := r.conn.QueryContext(ctx, List)
	if err != nil {
		return nil, common.ErrorRequest(err, http.StatusInternalServerError)
	}

	products := []*entity.Products{}
	for rows.Next() {

		product := entity.Products{}

		var CreatedAt sql.NullTime
		var UpdatedAt sql.NullTime

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Picture,
			&product.Price,
			&product.BrandId,
			&CreatedAt,
			&UpdatedAt,
		); err != nil {
			return nil, common.ErrorRequest(err, http.StatusInternalServerError)
		}

		product.CreatedAt = CreatedAt.Time
		product.UpdateAt = UpdatedAt.Time

		products = append(products, &product)
	}
	return products, nil
}

// Detail
func (r *ProductsRepository) Detail(ctx context.Context, product *entity.Products) (*entity.Products, error) {
	result := entity.Products{}
	var CreatedAt sql.NullTime
	var UpdatedAt sql.NullTime

	if err := r.conn.QueryRowContext(ctx, Detail, product.ID).Scan(
		&result.ID,
		&result.Name,
		&result.Picture,
		&result.Price,
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
func (r *ProductsRepository) Update(ctx context.Context, product *entity.Products) error {

	//Check if exist product
	if _, err := r.Detail(ctx, product); err != nil {
		return err
	}

	if _, err := r.conn.ExecContext(ctx, Update,
		&product.Name,
		&product.Picture,
		&product.Price,
		&product.BrandId,
		time.Now(),
		product.ID); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}

// Delete
func (r *ProductsRepository) Delete(ctx context.Context, product *entity.Products) error {

	//Check if exist product
	if _, err := r.Detail(ctx, product); err != nil {
		return err
	}

	if err := r.conn.QueryRowContext(ctx, Delete, product.ID).Err(); err != nil {
		if err == sql.ErrNoRows {
			return common.ErrorRequest(err, http.StatusNotFound)
		}
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}
