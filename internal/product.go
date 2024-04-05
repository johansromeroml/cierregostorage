package internal

// ProductAttributes is the struct that represents the attributes of a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string
	// Price is the price of the product.
	Price float64
}

// Product is the struct that represents a product.
type Product struct {
	// Id is the unique identifier of the product.
	Id int
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

// ProductJSON is the struct that server for the Unmarshal off product JSONs
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type TopSeller struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)

	FindTopSellers() (p []TopSeller, err error)
	// Save saves a product.
	Save(p *Product) (err error)
}

// RepositoryProduct is the interface that wraps the basic methods that a product repository must have.
type RepositoryProduct interface {
	// FindAll returns all products saved in the database.
	FindAll() (p []Product, err error)

	FindTopSellers() (p []TopSeller, err error)
	// Save saves a product into the database.
	Save(p *Product) (err error)
}
