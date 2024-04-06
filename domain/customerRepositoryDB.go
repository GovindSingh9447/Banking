package domain

import (
	"Banking/errors"
	"Banking/logger"
	"database/sql"
	

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errors.AppError) {

	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {

		
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	}else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)

	}
	
	

	
	if err != nil {
		logger.Error("Error while quering customer table " + err.Error())
		return nil, errors.NewUnexpectedError("Error while quering customer table")
	}

	// //customers := make([]Customer, 0)
	// err =sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while quering customer table " + err.Error())
	// 	return nil, errors.NewUnexpectedError("Error while quering customer table")
	// }


	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	// 	if err != nil {
	// 		logger.Error("Error while scaning customers " + err.Error())
	// 		return nil, errors.NewNotFoundError("Error while scaning customers")
	// 	}

	// 	customers = append(customers, c)
	// }

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppError) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	//row := d.client.QueryRow(customerSql,id)
	var c Customer
	err :=d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("Customer not found")
		}else{
			logger.Error("Error while scaning customer " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
          
	}
	return &c, nil

}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{dbClient}

}
