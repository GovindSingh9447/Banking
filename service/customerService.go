package service

import "Banking/domain"
import "Banking/errors"


type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errors.AppError)
}


type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer,error){
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer,*errors.AppError){
	return s.repo.ById(id)
}


func NewCustomerService (repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}