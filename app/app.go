package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"log"
	"net/http"
    "os"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
	   os.Getenv("SERVER_PORT")==""{
		log.Fatal("Environment variable not defined.....")
	   }
}


func Start() {
    // Create a new router
	sanityCheck()
	
    router := mux.NewRouter()


	//wring

	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
    //ch :=CustomerHandler{ service.NewCustomerService(domain.NewCustomerRepositoryStub() )}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
    // Define routes
    router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")


    // Starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	addr := fmt.Sprintf("%s:%s", address, port)
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))

   // log.Fatal(http.ListenAndServe("localhost:8888", router))

}

