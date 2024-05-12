package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// create struct for customer

type Customer struct {
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

// create a map

var customers = map[string]Customer{
	"1": {
		Name:      "Jim Barns",
		Role:      "Customer",
		Email:     "jim@gmail.com",
		Phone:     111111111,
		Contacted: false,
	},
	"2": {
		Name:      "Carl Bet",
		Role:      "Customer",
		Email:     "jim@gmail.com",
		Phone:     111111111,
		Contacted: false,
	},
	"3": {
		Name:      "Jim Born",
		Role:      "Basic Customer",
		Email:     "jim@gmail.com",
		Phone:     111111111,
		Contacted: false,
	},
	"4": {
		Name:      "Carley Fish",
		Role:      "Customer",
		Email:     "jim@gmail.com",
		Phone:     111111111,
		Contacted: false,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry map[string]Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	for k, v := range newEntry {
		if _, ok := customers[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			customers[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		var updatedCustomer Customer

		reqBody, _ := ioutil.ReadAll(r.Body)

		json.Unmarshal(reqBody, &updatedCustomer)

		customers[id] = updatedCustomer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedCustomer)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
