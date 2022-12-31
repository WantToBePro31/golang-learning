package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return []model.Product{}, err
	}
	return products, nil // TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	var product model.Product
	if err := p.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	var prod model.Product
	if err := p.db.Model(&prod).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
