package motor

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestMotorInsert(t *testing.T) {
	motorRepository := NewMotorRepository(go_database.GetConnection())

	ctx := context.Background()
	motor := entity.Motor{
		Merk:        "SIZUKA",
		Transmition: "Manual",
		Price:       16000000,
	}

	result, err := motorRepository.Insert(ctx, motor)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestMotorFindById(t *testing.T) {
	motorRepository := NewMotorRepository(go_database.GetConnection())

	motor, err := motorRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(motor)
}

func TestMotorFindAll(t *testing.T) {
	motorRepository := NewMotorRepository(go_database.GetConnection())

	motors, err := motorRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, motor := range motors {
		fmt.Println(motor)
	}
}

func TestMotorDelete(t *testing.T) {
	motorRepository := NewMotorRepository(go_database.GetConnection())

	ctx := context.Background()
	motor := entity.Motor{
		Id: 5,
	}

	result, err := motorRepository.Delete(ctx, motor)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCustomerUpdate(t *testing.T) {
	motorRepository := NewMotorRepository(go_database.GetConnection())
	ctx := context.Background()
	motor := entity.Motor{
		Merk:        "HONDAR",
		Transmition: "Automatic",
		Price:       18650000,
	}
	result, err := motorRepository.Update(ctx, 1, motor)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
