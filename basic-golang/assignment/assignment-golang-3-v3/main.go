package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	var res []invoice.InvoiceData
	for _, val := range data {
		switch val.(type) {
			case invoice.FinanceInvoice:
				fi := val.(invoice.FinanceInvoice)
				tmp, err := fi.RecordInvoice()
				if err == nil {
					if fi.Status == invoice.PAID && fi.Approved {
						found := -1
						for ind, dt := range res {
							if dt.Date == tmp.Date && dt.Departemen == invoice.Finance {
								found = ind
							}
						}
						if found == -1 {
							res = append(res, tmp)
						} else {
							res[found].TotalInvoice += tmp.TotalInvoice
						}
					}
				} else {
					return []invoice.InvoiceData{}, err
				}
			case invoice.WarehouseInvoice:
				wi := val.(invoice.WarehouseInvoice)
				tmp, err := wi.RecordInvoice()
				if err == nil {
					if wi.InvoiceType == invoice.PURCHASE && wi.Approved {
						found := -1
						for ind, dt := range res {
							if dt.Date == wi.Date && dt.Departemen == invoice.Warehouse {
								found = ind
							}
						}
						if found == -1 {
							res = append(res, tmp)
						} else {
							res[found].TotalInvoice += tmp.TotalInvoice
						}
					}
				} else {
					return []invoice.InvoiceData{}, err
				}
			case invoice.MarketingInvoice:
				mi := val.(invoice.MarketingInvoice)
				tmp, err := mi.RecordInvoice()
				if err == nil {
					if mi.Approved {
						found := -1
						for ind, dt := range res {
							if dt.Date == mi.Date && dt.Departemen == invoice.Marketing {
								found = ind
							}
						}
						if found == -1 {
							res = append(res, tmp)
						} else {
							res[found].TotalInvoice += tmp.TotalInvoice
						}
					}
				} else {
					return []invoice.InvoiceData{}, err
				}
			default:
				return nil, nil
		}
	}
	return res, nil // TODO: replace this
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{"product A", 10, 10000, 0.1},
				{"product C", 5, 15000, 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
