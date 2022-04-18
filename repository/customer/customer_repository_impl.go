package customer

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type CustomerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db}
}

func (repository *CustomerRepositoryImpl) Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := "INSERT INTO customer(nama, status) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, customer.Nama, customer.Status)
	if err != nil {
		return customer, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return customer, err
	}
	customer.Id = int32(id)
	return customer, nil
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Customer, error) {
	script := "SELECT id, nama FROM customer WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	customer := entity.Customer{}
	if err != nil {
		return customer, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&customer.Id, &customer.Nama, &customer.Status)
		return customer, nil
	} else {
		//tidak ada
		return customer, errors.New("Id" + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]entity.Customer, error) {
	script := "SELECT id, nama, status FROM customer"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		rows.Scan(&customer.Id, &customer.Nama, &customer.Status)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := "DELETE FROM customer WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, customer.Id)
	if err != nil {
		return customer, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return customer, err
	}
	if rows == 0 {
		return customer, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, id int32, customer entity.Customer) (entity.Customer, error) {
	script := "SELECT id, nama, status FROM customer WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return customer, err
	}

	if rows.Next() {
		script := "UPDATE customer SET nama = ?, status = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, customer.Nama, customer.Status, id)
		if err != nil {
			return customer, err
		}
		customer.Id = id
		return customer, nil
	} else {
		return customer, errors.New(("Id" + strconv.Itoa(int(id)) + " Update Failed"))
	}
}
