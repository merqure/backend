package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/merqure/backend/models"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func SellersIndex(w http.ResponseWriter, r *http.Request) {
	sellers := models.Sellers{
		models.Seller{Name: "Apple"},
		models.Seller{Name: "Samsung"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(sellers); err != nil {
		panic(err)
	}
}

func SellerShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sellerId := vars["sellerId"]
	fmt.Fprintln(w, "Seller show:", sellerId)
}
