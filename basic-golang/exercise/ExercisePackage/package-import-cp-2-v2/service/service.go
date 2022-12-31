package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, errGet := s.database.GetProductByname(productName)
	if errGet != nil {
		return errGet
	}
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	tmp := entity.CartItem {
		ProductName: productName,
		Price: product.Price,
		Quantity: quantity,
	}
	carts, errLoad := s.database.Load()
	if errLoad != nil {
		return errLoad
	}
	carts = append(carts, tmp)

	errSave := s.database.Save(carts)
	if errSave != nil {
		return errSave
	}
	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	var tmp []entity.CartItem
	var found bool = false
	carts, errLoad := s.database.Load()
	if errLoad != nil {
		return errLoad
	}
	for _, val := range carts {
		if val.ProductName != productName {
			tmp = append(tmp, val)
		} else {
			found = true
		}
	}
	if found {
		errSave := s.database.Save(tmp)
		if errSave != nil {
			return errSave
		}
		return nil
	}
	return errors.New("product not found") // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (s *Service) ResetCart() error {
	errSave := s.database.Save([]entity.CartItem{})
	if errSave != nil {
		return errSave
	}
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	products := s.database.GetProductData()
	return products, nil // TODO: replace this
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {
	carts, errLoad := s.database.Load()
	if errLoad != nil {
		return entity.PaymentInformation{}, errLoad
	}
	totalPrice := 0
	for _, val := range carts {
		totalPrice += val.Price * val.Quantity
	}
	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}
	change := money - totalPrice
	errReset := s.ResetCart()
	if errReset != nil {
		return entity.PaymentInformation{}, errReset
	}
	return entity.PaymentInformation{ListProduct: carts, TotalPrice: totalPrice, MoneyPaid: money, Change: change}, nil // TODO: replace this
}
