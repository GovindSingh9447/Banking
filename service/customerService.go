package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/errors"
)

// primary port
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	c.ToDto()
	respone := c.ToDto()
	return &respone, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
