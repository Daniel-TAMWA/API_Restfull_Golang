package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Create a Customer struct
type Customer struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Role      string `json:"Role"`
	Email     string `json:"Email"`
	Phone     int    `json:"Phone"`
	Contacted bool   `json:"Contacted"`
}

// create the Customers variable then create the variables, customer_1, customer_2, customer_3, customer_4 in Customers.
var Customers = make(map[string]*Customer)
var customer_1 = &Customer{Id: uuid.NewString(), Name: "Yoan", Role: "admin", Email: "yoan@email.com", Phone: 123654, Contacted: true}
var customer_2 = &Customer{Id: uuid.NewString(), Name: "Dan", Role: "contributor", Email: "dan@email.com", Phone: 852456, Contacted: false}
var customer_3 = &Customer{Id: uuid.NewString(), Name: "Lorris", Role: "contributor", Email: "lorris@email.com", Phone: 753951, Contacted: false}
var customer_4 = &Customer{Id: uuid.NewString(), Name: "Anti", Role: "contributor", Email: "anti@email.com", Phone: 951753, Contacted: true}

// homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Welcome in the CRM backend!</H1>")
	fmt.Println("Endpoint: homePage")
}

// create the function for to obtain all the list of customer
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint: getcustomers")
	json.NewEncoder(w).Encode(Customers)
}

// create the function for to select a customer
func getCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// read the request from the user, for which we want to fetch the data
	id := mux.Vars(r)["id"]

	if _, ok := Customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println("Endpoint: getcustomer")
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newCustomer)
	newCustomer.Id = uuid.NewString()

	if _, ok := Customers[newCustomer.Id]; ok {
		w.WriteHeader(http.StatusConflict)
	} else {
		Customers[newCustomer.Id] = &newCustomer
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(Customers)
	fmt.Println("Endpoint: addcustomer")
}

// create the function for to delete the customer

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := Customers[id]; ok {
		delete(Customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Customers)
	}
	fmt.Println("Endpoint: deletecustomer")
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	if _, ok := Customers[id]; ok {
		json.Unmarshal(reqBody, Customers[id])

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(Customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Customers)
	}
	fmt.Println("Endpoint: updatecustomer")
}

// The fonctions for all the handlerRequests.
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/customers", getCustomers).Methods("GET")
	myRouter.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	myRouter.HandleFunc("/customers", addCustomer).Methods("POST")
	myRouter.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	myRouter.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

// The main function
func main() {

	Customers[customer_1.Id] = customer_1
	Customers[customer_2.Id] = customer_2
	Customers[customer_3.Id] = customer_3
	Customers[customer_4.Id] = customer_4
	handleRequests()
}
