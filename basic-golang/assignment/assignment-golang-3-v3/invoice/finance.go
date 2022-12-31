package invoice

import (
	"errors"
	"fmt"
	"strings"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if fi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if len(fi.Details) == 0 {
		return InvoiceData{}, errors.New("invoice details is empty")
	}
	if fi.Status == "" && fi.Status != PAID && fi.Status != UNPAID {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}
	for _, val := range fi.Details {
		if val.Total <= 0 {
			return InvoiceData{}, errors.New("total price is not valid")
		}
	}
	month := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}
	dateFormat := strings.Split(fi.Date, "/")
	date := fmt.Sprintf("%s-%s-%s", dateFormat[0], month[dateFormat[1]], dateFormat[2])

	var totalPrice float64
	for _, data := range fi.Details {
		totalPrice += float64(data.Total)
	}
	return InvoiceData{Date: date, TotalInvoice: totalPrice, Departemen: Finance}, nil // TODO: replace this
}
