package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//ProductRepo Tank for object
type ProductRepo struct {
	DB *pg.DB
}

//GetProduct all
func (p *ProductRepo) GetProduct() ([]*model.Product, error) {

	var products []*model.Product
	err := p.DB.Model(&products).Select()
	if err != nil {
		return nil, err
	}

	return products, nil
}