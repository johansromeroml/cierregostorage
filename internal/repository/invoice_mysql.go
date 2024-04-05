package repository

import (
	"database/sql"
	"fmt"

	"app/internal"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindUpdatedTotals() (i []internal.Invoice, err error) {
	// execute the update query
	_, err = r.db.Exec("UPDATE invoices as i1, " +
		"(select s.invoice_id as sid, sum(s.quantity*p.price) as new_total " +
		"from sales s join invoices i on s.invoice_id = i.id join products p on s.product_id = p.id " +
		"group by s.invoice_id ) as i2 " +
		"set i1.total = i2.new_total " +
		"where i1.id = i2.sid ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// execute the get query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*i).Id = int(id)

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) SaveWithId(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`id`, `datetime`, `total`, `customer_id`) VALUES (?, ?, ?, ?)",
		(*i).Id, (*i).Datetime, (*i).Total, (*i).CustomerId,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*i).Id = int(id)

	return
}
