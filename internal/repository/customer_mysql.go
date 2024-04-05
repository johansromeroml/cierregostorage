package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *CustomersMySQL) FindTopActive() (c []internal.TopActiveCustomer, err error) {
	// execute the query
	rows, err := r.db.Query("select c.first_name , c.last_name , sum(i.total) as total from invoices i join customers c on i.customer_id = c.id " +
		"where c.`condition` = 1 group by i.customer_id order by total desc limit 5;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.TopActiveCustomer
		// scan the row into the customer
		err := rows.Scan(&cs.FirstName, &cs.LastName, &cs.Total)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *CustomersMySQL) FindTotalByCondition() (c []internal.CustomerConditionTotals, err error) {
	// execute the query
	rows, err := r.db.Query("select c.`condition` , sum(i.total) from invoices i join customers c on i.customer_id = c.id group by c.`condition`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.CustomerConditionTotals
		// scan the row into the customer
		err := rows.Scan(&cs.Condition, &cs.Total)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) SaveWithId(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (?, ?, ?, ?)",
		(*c).Id, (*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}
