package app

import (
	"Banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip" xml:"zip"`
}


type CustomerHandler struct{
	service service.CustomerService
}


func (ch *CustomerHandler)getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*customers := []Customer{
		{"Govind", "Patna", "800001"},
		{"Arvind", "Balrampur", "700001"},
		{"Vivek", "WB", "600001"},
	}*/

	customers, _ := ch.service.GetAllCustomer()

	

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
