package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func sanityCheck() {
	requiredEnv := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASS", "DB_ADDR", "DB_PORT", "DB_NAME"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			log.Fatalf("Environment variable %s not defined", env)
		}
	}
}

func Start() {
	// Perform sanity check for required environment variables
	sanityCheck()

	// Create a new router
	router := mux.NewRouter()

	// Initialize database client
	dbClient := getDbClient()

	// Initialize repositories
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// Initialize service handlers
	customerService := service.NewCustomerService(customerRepositoryDb)
	accountService := service.NewAccountService(accountRepositoryDb)

	// Initialize HTTP handlers
	ch := CustomerHandler{customerService}
	ah := AccountHandler{service: accountService}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods("POST")

	// Starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	addr := fmt.Sprintf("%s:%s", address, port)
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
