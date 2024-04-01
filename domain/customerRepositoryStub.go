package domain




type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error){
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
        // {"3001","Arvind","Balrampur","900001","30-09-2001", "Active"},
		// {"3002","Govind","Balrampur","700001","20-08-2001", "Active"},

	}

	return CustomerRepositoryStub{customers }
}