package invoice

import (
	"errors"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if wi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if wi.InvoiceType == "" && wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {
		return InvoiceData{}, errors.New("invoice type is not valid")
	}
	if len(wi.Products) == 0 {
		return InvoiceData{}, errors.New("invoice products is empty")
	}
	for _, val := range wi.Products {
		if val.Unit <= 0 {
			return InvoiceData{}, errors.New("unit product is not valid")
		}
		if val.Price <= 0.0 {
			return InvoiceData{}, errors.New("price product is not valid")
		}
	}

	var totalPrice float64
	for _, data := range wi.Products {
		totalPrice += (1.0 - data.Discount) * (float64(data.Unit) * data.Price)
	}
	return InvoiceData{Date: wi.Date, TotalInvoice: totalPrice, Departemen: Warehouse}, nil // TODO: replace this
}
