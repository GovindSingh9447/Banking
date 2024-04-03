package app

import (
	"Banking/service"
	"encoding/json"
	//"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip" xml:"zip"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*customers := []Customer{
		{"Govind", "Patna", "800001"},
		{"Arvind", "Balrampur", "700001"},
		{"Vivek", "WB", "600001"},
	}*/



	// customers, _ := ch.service.GetAllCustomer()

	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil{
		writeResponse(w, err.Code, err.AsMessage())
	}else{
        writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		//fmt.Fprint(w,err.Message)
	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {

		panic(err)
	}
}
