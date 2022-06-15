package main

import "net/http"

func Start() {

	http.HandleFunc("/green", green)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8090", nil)
}
