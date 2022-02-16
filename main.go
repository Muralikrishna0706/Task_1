package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murali0706/customer/routes"
)

func main() {
	routers:= mux.NewRouter()
	routers.HandleFunc("/customer/create", routes.CreateCustomer).Methods("POST")
	routers.HandleFunc("/customer/{id}", routes.GetCustomer).Methods("GET")
	routers.HandleFunc("/customer", routes.GetAllCustomers).Methods("GET")
	routers.HandleFunc("/customer/{id}", routes.UpdateCustomer).Methods("PATCH")
	routers.HandleFunc("/customer/delete/{id}", routes.DeleteCustomer).Methods("DELETE")
	http.ListenAndServe(":8080", routers)
}
