package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"
	"fmt"
)

type CartRepository struct {
	db db.DB
}

func NewCartRepository(db db.DB) CartRepository {
	return CartRepository{db}
}

func (u *CartRepository) ReadCart() ([]model.Cart, error) {
	records, err := u.db.Load("carts")
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("Cart not found!")
	}

	var cart []model.Cart
	err = json.Unmarshal([]byte(records), &cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (u *CartRepository) UpdateCart(cart model.Cart) error {
	listCart, err := u.ReadCart()
	if err != nil {
		return err
	}

	for ind, userCart := range listCart {
		if userCart.Name == cart.Name {
			listCart[ind].Cart = cart.Cart
			listCart[ind].TotalPrice = cart.TotalPrice
		}
	}

	jsonData, err := json.Marshal(listCart)
	if err != nil {
		return err
	}

	err = u.db.Save("carts", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) AddCart(cart model.Cart) error {
	if cart.Name == "" {
		return errors.New("bad request")
	}

	listCart, err := u.ReadCart()
	if err != nil {
		return err
	}

	listCart = append(listCart, cart)
	jsonData, err := json.Marshal(listCart)
	if err != nil {
		return err
	}

	err = u.db.Save("carts", jsonData)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *CartRepository) ResetCarts() error {
	err := u.db.Reset("carts", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) CartUserExist(name string) (model.Cart, error) {
	listcCart, err := u.ReadCart()
	if err != nil {
		return model.Cart{}, err
	}
	for _, element := range listcCart {
		if element.Name == name {
			return element, nil
		}
	}
	return model.Cart{}, fmt.Errorf("Cart Empty!")
}
