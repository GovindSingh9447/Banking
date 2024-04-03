package domain

import "Banking/errors"

type Customer struct{


	Id   string
	Name string
	City  string
	Zipcode string
	DateofBirth string
	Status   string 
}


type CustomerRepository interface{
	FindAll(status string) ([]Customer, *errors.AppError)

	ById(string) (*Customer,*errors.AppError)
}

