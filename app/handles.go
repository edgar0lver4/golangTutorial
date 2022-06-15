package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string
	Zipcode string
}

func green(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Edgar", "Edo mex", "55290"},
		{"Andrei", "Hidalgo", "55879"},
	}

	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
}
