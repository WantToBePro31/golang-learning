package invoice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	if mi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, errors.New("travel date is empty")
	}
	if mi.PricePerDay <= 0 {
		return InvoiceData{}, errors.New("price per day is not valid")
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
	dateFormat := strings.Split(mi.Date, "/")
	date := fmt.Sprintf("%s-%s-%s", dateFormat[0], month[dateFormat[1]], dateFormat[2])

	startDateFormat := strings.Split(mi.StartDate, "/")
	start, _ := strconv.Atoi(startDateFormat[0])
	endDateFormat := strings.Split(mi.EndDate, "/")
	end, _ := strconv.Atoi(endDateFormat[0])
	duration := end - start + 1

	var totalPrice float64 = float64(duration * mi.PricePerDay + mi.AnotherFee)
	return InvoiceData{Date: date, TotalInvoice: totalPrice, Departemen: Marketing}, nil // TODO: replace this
}
