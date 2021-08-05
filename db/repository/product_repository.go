package repository

import (
	"database/sql"
	"strings"

	"github.com/yafiesetyo/merempah-api-clone/db/repository/models"
	"github.com/yafiesetyo/merempah-api-clone/helper"
	"github.com/yafiesetyo/merempah-api-clone/usecase/viewmodel"
)

// IProduct...
type IProduct interface {
	SelectAll(search string) ([]models.Product, error)
	FindById(id int64) (models.Product, error)
	Add(body *viewmodel.ProductVM) (int64, error)
	Edit(id int64, body *viewmodel.ProductVM) (int64, error)
	Delete(id int64) (int64, error)
}

// productRepository...
type productRepository struct {
	DB *sql.DB
}

// NewProductRepository...
func NewProductRepository(db *sql.DB) IProduct {
	return &productRepository{DB: db}
}

// SelectAll...
func (model productRepository) SelectAll(search string) (data []models.Product, err error) {
	query := models.ProductSelectStatement + ` WHERE name LIKE $1 or store LIKE $1`
	println(search)
	rows, err := model.DB.Query(query, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		d := models.Product{}
		err = rows.Scan(&d.ID, &d.Category, &d.Name, &d.Stock, &d.Store)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, err
}

// FindById...
func (model productRepository) FindById(id int64) (data models.Product, err error) {
	query := models.ProductSelectStatement + ` WHERE id = $1 LIMIT 1`
	err = model.DB.QueryRow(query, helper.EmptyId(id)).Scan(&data.ID, &data.Name, &data.Store, &data.Stock, &data.Category)
	return data, err
}

// Add...
func (model productRepository) Add(body *viewmodel.ProductVM) (res int64, err error) {
	query := `INSERT into product (name,store,stock,category) values ($1,$2,$3,$4) RETURNING id`
	err = model.DB.QueryRow(query,
		body.Name, body.Store, body.Stock, body.Category,
	).Scan(&res)

	return res, err

}

// Edit...
func (model productRepository) Edit(id int64, body *viewmodel.ProductVM) (res int64, err error) {
	query := `UPDATE product set name = $1, store = $2, stock = $3, category = $4 where id = $5 RETURNING id`
	err = model.DB.QueryRow(
		query, body.Name, body.Store, body.Stock, body.Category, id,
	).Scan(&res)

	return res, err
}

// Delete...
func (model productRepository) Delete(id int64) (res int64, err error) {
	query := `DELETE FROM product where id = $1 RETURNING id`
	err = model.DB.QueryRow(
		query, id,
	).Scan(&res)

	return res, err
}
