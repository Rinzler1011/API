package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// create struct for customer

type Customer struct {
	Id        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

// create a map

var customers = map[string]Customer{
	"1": {Id: uuid.New(), Name: "John Doe", Role: "Manager", Email: "johndoe@example.com", Phone: 1234567890, Contacted: false},
	"2": {Id: uuid.New(), Name: "Jane Smith", Role: "Developer", Email: "janesmith@example.com", Phone: 2345678901, Contacted: true},
	"3": {Id: uuid.New(), Name: "Emily Johnson", Role: "Designer", Email: "emilyjohnson@example.com", Phone: 3456789012, Contacted: false},
	"4": {Id: uuid.New(), Name: "Michael Brown", Role: "Analyst", Email: "michaelbrown@example.com", Phone: 4567890123, Contacted: true},
}

func updateCustomerId() {
	newMap := map[string]Customer{}
	for _, customer := range customers {
		newID := uuid.New()
		customer.Id = newID
		newMap[newID.String()] = customer
	}
	customers = newMap
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

	var newCustomer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, &newCustomer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCustomer.Id = uuid.New()
	customers[newCustomer.Id.String()] = newCustomer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if customer, ok := customers[id]; ok {
		var updatedCustomer Customer
		reqBody, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(reqBody, &updatedCustomer); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		updatedCustomer.Id = customer.Id
		customers[id] = updatedCustomer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedCustomer)
	} else {
		w.WriteHeader(http.StatusNotFound)
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

	updateCustomerId()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	index := http.FileServer(http.Dir("html"))
	router.Handle("/", index)

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
