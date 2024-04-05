package internal

// SaleAttributes is the struct that represents the attributes of a sale.
type SaleAttributes struct {
	// Quantity is the quantity of the sale.
	Quantity int
	// ProductId is the product id of the sale.
	ProductId int
	// InvoiceId is the invoice id of the sale.
	InvoiceId int
}

// Sale is the struct that represents a sale.
type Sale struct {
	// Id is the unique identifier of the sale.
	Id int
	// SaleAttributes is the attributes of the sale.
	SaleAttributes
}

type SaleJSON struct {
	Id        int `json:"id"`
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
}

// ServiceSale is the interface that wraps the basic ServiceSale methods.
type ServiceSale interface {
	// FindAll returns all sales.
	FindAll() (s []Sale, err error)
	// Save saves a sale.
	Save(s *Sale) (err error)
}

// RepositorySale is the interface that wraps the basic Sale methods.
type RepositorySale interface {
	// FindAll returns all sales.
	FindAll() (s []Sale, err error)
	// Save saves a sale.
	Save(s *Sale) (err error)
}
