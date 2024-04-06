package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/errors"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	err :=req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResposeDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{repo}
}
