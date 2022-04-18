package calonmertua

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCalonmertuaInsert(t *testing.T) {
	calonmertuaRepository := NewCalonmertuaRepository(go_database.GetConnection())

	ctx := context.Background()
	calonmertua := entity.Calonmertua{
		Nama:   "EUIS",
		Alamat: "Ciawi Tali",
		Suku:   "Sunda",
	}

	result, err := calonmertuaRepository.Insert(ctx, calonmertua)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCalonmertuaFindById(t *testing.T) {
	calonmertuaRepository := NewCalonmertuaRepository(go_database.GetConnection())

	calonmertua, err := calonmertuaRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(calonmertua)
}

func TestCalonmertuaFindAll(t *testing.T) {
	calonmertuaRepository := NewCalonmertuaRepository(go_database.GetConnection())

	kandidat, err := calonmertuaRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, calonmertua := range kandidat {
		fmt.Println(calonmertua)
	}
}

func TestCalonmertuaDelete(t *testing.T) {
	calonmertuaRepository := NewCalonmertuaRepository(go_database.GetConnection())

	ctx := context.Background()
	calonmertua := entity.Calonmertua{
		Id: 5,
	}

	result, err := calonmertuaRepository.Delete(ctx, calonmertua)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCalonmertuaUpdate(t *testing.T) {
	calonmertuaRepository := NewCalonmertuaRepository(go_database.GetConnection())
	ctx := context.Background()
	calonmertua := entity.Calonmertua{
		Nama:   "Pringga",
		Alamat: "Andara",
		Suku:   "Badui",
	}
	result, err := calonmertuaRepository.Update(ctx, 6, calonmertua)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
