package application

import (
	"app/internal"
	"app/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// -----------------------------------------
	fmt.Println("Initializing db recovery from JSON files")
	cfg := mysql.Config{
		User:   "root",
		Passwd: "masterkey",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fantasy_products",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
	}
	// - db: ping
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	// -------------------------------------------
	fmt.Println("products JSON")
	productsFile, err := os.Open("./docs/db/json/products.json")
	productsRepo := repository.NewProductsMySQL(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("Products file closed")
	}(productsFile)
	productsJSON, err := io.ReadAll(productsFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	products := make([]internal.ProductJSON, 100)
	json.Unmarshal(productsJSON, &products)
	for _, v := range products {
		product := internal.Product{
			Id: v.Id,
			ProductAttributes: internal.ProductAttributes{
				Description: v.Description,
				Price:       v.Price,
			},
		}
		productsRepo.SaveWithId(&product)
	}
	// -----------------------------------------
	fmt.Println("customers JSON")
	customersFile, err := os.Open("./docs/db/json/customers.json")
	customersRepo := repository.NewCustomersMySQL(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("Customers file closed")
	}(customersFile)
	customersJSON, err := io.ReadAll(customersFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	customers := make([]internal.CustomerJSON, 100)
	json.Unmarshal(customersJSON, &customers)
	for _, v := range customers {
		customer := internal.Customer{
			Id: v.Id,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			},
		}
		customersRepo.SaveWithId(&customer)
	}
	// ----------------------------------------
	fmt.Println("invoices JSON")
	invoicesFile, err := os.Open("./docs/db/json/invoices.json")
	invoicesRepo := repository.NewInvoicesMySQL(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("Invoices file closed")
	}(invoicesFile)
	invoicesJSON, err := io.ReadAll(invoicesFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	invoices := make([]internal.InvoiceJSON, 100)
	json.Unmarshal(invoicesJSON, &invoices)
	for _, v := range invoices {
		invoice := internal.Invoice{
			Id: v.Id,
			InvoiceAttributes: internal.InvoiceAttributes{
				Total:      v.Total,
				CustomerId: v.CustomerId,
				Datetime:   v.Datetime,
			},
		}
		err := invoicesRepo.SaveWithId(&invoice)
		if err != nil {
			fmt.Println(err)
		}
	}
	// ----------------------------------------
	fmt.Println("sales JSON")
	salesFile, err := os.Open("./docs/db/json/sales.json")
	salesRepo := repository.NewSalesMySQL(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(f *os.File) {
		f.Close()
		fmt.Println("Sales file closed")
	}(salesFile)
	salesJSON, err := io.ReadAll(salesFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	sales := make([]internal.SaleJSON, 100)
	json.Unmarshal(salesJSON, &sales)
	for _, v := range sales {
		sale := internal.Sale{
			Id: v.Id,
			SaleAttributes: internal.SaleAttributes{
				Quantity:  v.Quantity,
				ProductId: v.ProductId,
				InvoiceId: v.InvoiceId,
			},
		}
		err := salesRepo.SaveWithId(&sale)
		if err != nil {
			fmt.Println(err)
		}
	}
}
