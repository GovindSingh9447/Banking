package dto

type CustomerResponse struct {
	Id          string `db:"customer_id"`
	Name        string `db:"full_name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateofBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}
