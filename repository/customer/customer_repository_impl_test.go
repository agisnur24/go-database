package customer

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCustomerInsert(t *testing.T) {
	customerRepository := NewCustomerRepository(go_database.GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Nama:   "Reza Cahya Fauzia",
		Status: "Perawan",
	}

	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCustomerFindById(t *testing.T) {
	customerRepository := NewCustomerRepository(go_database.GetConnection())

	customer, err := customerRepository.FindById(context.Background(), 5)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestCustomerFindAll(t *testing.T) {
	customerRepository := NewCustomerRepository(go_database.GetConnection())

	customers, err := customerRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func TestCustomerDelete(t *testing.T) {
	customerRepository := NewCustomerRepository(go_database.GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Id: 3,
	}

	result, err := customerRepository.Delete(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCustomerUpdate(t *testing.T) {
	customerRepository := NewCustomerRepository(go_database.GetConnection())
	ctx := context.Background()
	customer := entity.Customer{
		Nama:   "Reza Cawi Culun",
		Status: "Perawan",
	}
	result, err := customerRepository.Update(ctx, 1, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
