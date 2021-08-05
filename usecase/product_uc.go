package usecase

import (
	"errors"

	"github.com/yafiesetyo/merempah-api-clone/db/repository"
	"github.com/yafiesetyo/merempah-api-clone/db/repository/models"
	request "github.com/yafiesetyo/merempah-api-clone/server/requests"
	"github.com/yafiesetyo/merempah-api-clone/usecase/viewmodel"
)

// ProductUC...
type ProductUC struct {
	*ContractUC
}

// BuildBody...
func (uc ProductUC) BuildBody(data *models.Product, res *viewmodel.ProductVM) {
	res.ID = data.ID
	res.Name = data.Name
	res.Stock = data.Stock
	res.Store = data.Store
	res.Category = data.Category
}

// SelectAll...
func (uc ProductUC) SelectAll(search string) (res []viewmodel.ProductVM, err error) {
	repo := repository.NewProductRepository(uc.DB)
	data, err := repo.SelectAll(search)
	if err != nil {
		return res, err
	}

	for _, r := range data {
		tmp := viewmodel.ProductVM{}
		uc.BuildBody(&r, &tmp)
		res = append(res, tmp)
	}

	return res, err

}

// FindById...
func (uc ProductUC) FindById(id int64) (res viewmodel.ProductVM, err error) {
	repo := repository.NewProductRepository(uc.DB)
	data, err := repo.FindById(id)
	if err != nil {
		return res, errors.New("record doesnt exist")
	}
	uc.BuildBody(&data, &res)

	return res, err
}

// Add
func (uc ProductUC) Add(req *request.ProductRequest) (res viewmodel.ProductVM, err error) {
	repo := repository.NewProductRepository(uc.DB)
	res = viewmodel.ProductVM{
		ID:       req.ID,
		Name:     req.Name,
		Store:    req.Store,
		Stock:    req.Stock,
		Category: req.Category,
	}
	res.ID, err = repo.Add(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

// Edit...
func (uc ProductUC) Edit(id int64, req *request.ProductRequest) (res viewmodel.ProductVM, err error) {
	repo := repository.NewProductRepository(uc.DB)
	res = viewmodel.ProductVM{
		ID:       1,
		Name:     req.Name,
		Store:    req.Store,
		Stock:    req.Stock,
		Category: req.Category,
	}
	res.ID, err = repo.Edit(id, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// Delete
func (uc ProductUC) Delete(id int64) (res viewmodel.ProductVM, err error) {
	repo := repository.NewProductRepository(uc.DB)
	res.ID, err = repo.Delete(id)
	if err != nil {
		return res, err
	}

	return res, err
}
