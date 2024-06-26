package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

func (s *CustomersDefault) FindTopActive() (c []internal.TopActiveCustomer, err error) {
	c, err = s.rp.FindTopActive()
	return
}

func (s *CustomersDefault) FindTotalByCondition() (c []internal.CustomerConditionTotals, err error) {
	c, err = s.rp.FindTotalByCondition()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}
