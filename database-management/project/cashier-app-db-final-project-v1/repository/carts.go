package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	var joinedCart []model.JoinCart	
	if err := c.db.Table("carts").Select("carts.id as id, carts.product_id as product_id, products.name as name, carts.quantity as quantity, carts.total_price as total_price").Joins("JOIN products ON carts.product_id = products.id").Where("carts.deleted_at IS NULL").Find(&joinedCart).Error; err != nil {
		return []model.JoinCart{}, err 
	}
	return joinedCart, nil // TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		if product.Stock == 0 {
			return errors.New("Item is out of stock")
		}

		if err := tx.Model(&product).Where("id = ?", product.ID).Update("stock", product.Stock-1).Error; err != nil {
			return err
		}

		var cart model.Cart
		if err := tx.Model(&cart).Where("product_id = ?", product.ID).Find(&cart).Error; err != nil {
			return err
		}

		if cart.ID == 0 {
			newCart := model.Cart{
				ProductID: product.ID,
				Quantity: 1,
				TotalPrice: ((100.0-product.Discount)*product.Price)/100.0,
			}
			if err := tx.Create(&newCart).Error; err != nil {
				return err
			}
			return nil
		}

		cart.Quantity += 1;
		cart.TotalPrice = cart.Quantity * (((100.0-product.Discount)*product.Price)/100.0)

		if err := tx.Where("product_id = ?", product.ID).Updates(cart).Error; err != nil {
			return err
		}

		return nil
	})

	return nil // TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	c.db.Transaction(func(tx *gorm.DB) error {
		var cart model.Cart
		if err := tx.Model(&cart).Where("id = ?", id).Find(&cart).Error; err != nil {
			return err
		}

		if err := tx.Model(&cart).Where("id = ?", id).Delete(&cart).Error; err != nil {
			return err
		}

		var product model.Product
		if err := tx.Model(&product).Where("id = ?", productID).Find(&product).Error; err != nil {
			return err
		}
		if err := tx.Model(&product).Where("id = ?", productID).Update("stock", float64(product.Stock) + cart.Quantity).Error; err != nil {
			return err
		}

		return nil
	})

	return nil // TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	var ct model.Cart
	if err := c.db.Model(&ct).Where("id = ?", id).Updates(cart).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
