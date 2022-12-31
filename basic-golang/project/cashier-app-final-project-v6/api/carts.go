package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	username := r.Context().Value("username").(string)
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Request Product Not Found"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	// TODO: answer here
	product := r.FormValue("product")
	if product == "" {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Request Product Not Found"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}

	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{} // TODO: replace this
	cart.Name = username
	cart.Cart = list
	totalPrice := 0.0
	for _, val := range cart.Cart {
		totalPrice += val.Total
	}
	cart.TotalPrice = totalPrice

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
		
	api.dashboardView(w, r)
}
