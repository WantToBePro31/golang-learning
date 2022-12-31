package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func (api *API) dashboardView(w http.ResponseWriter, r *http.Request) {

	var username string
	if r.FormValue("username") != "" {
		username = r.FormValue("username")
	} else {
		username = fmt.Sprintf("%s", r.Context().Value("username"))
	}

	var data model.Dashboard
	listProducts, err := api.products.ReadProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cart, err := api.cartsRepo.CartUserExist(username)
	if err != nil {
		data = model.Dashboard{
			Product: listProducts,
			Cart: model.Cart{
				Name: username,
			},
		}
	} else {
		data = model.Dashboard{
			Product: listProducts,
			Cart:    cart,
		}
	}

	filepath := path.Join("views", "dashboard.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
	}
}
