package app

import (
	"Banking/domain"
	"Banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Start() {
    // Create a new router
    router := mux.NewRouter()


	//wring

	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
    //ch :=CustomerHandler{ service.NewCustomerService(domain.NewCustomerRepositoryStub() )}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
    // Define routes
    router.HandleFunc("/customer", ch.getAllCustomers).Methods("GET")

    // Starting server
    log.Fatal(http.ListenAndServe("localhost:8888", router))
}
