package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murali0706/customer/config"
	customerHttp "github.com/murali0706/customer/customer/http"
	repository "github.com/murali0706/customer/repository/mysql"
	"github.com/murali0706/customer/service"
)

func main() {
	dbConn, err := config.GetDB()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	customerRepo := repository.NewMysqlCustomerReposity(dbConn)

	cu := service.NewCustomerUsecase(customerRepo)
	customerHttp.NewCustomerHandler(r, cu)

	log.Fatal(http.ListenAndServe(":8080", r))
}
