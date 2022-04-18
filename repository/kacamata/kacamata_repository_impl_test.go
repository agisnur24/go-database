package kacamata

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	kacamataRepository := NewKacamataRepository(go_database.GetConnection())

	ctx := context.Background()
	kacamata := entity.Kacamata{
		Merk:  "topengrazor@gmail.com",
		Price: 2500000,
	}

	result, err := kacamataRepository.Insert(ctx, kacamata)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	kacamataRepository := NewKacamataRepository(go_database.GetConnection())

	kacamata, err := kacamataRepository.FindById(context.Background(), 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(kacamata)
}

func TestCommentFindAll(t *testing.T) {
	kacamataRepository := NewKacamataRepository(go_database.GetConnection())

	glasses, err := kacamataRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, kacamata := range glasses {
		fmt.Println(kacamata)
	}
}

func TestCommentDelete(t *testing.T) {
	kacamataRepository := NewKacamataRepository(go_database.GetConnection())

	ctx := context.Background()
	kacamata := entity.Kacamata{
		Id: 5,
	}

	result, err := kacamataRepository.Delete(ctx, kacamata)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentUpdate(t *testing.T) {
	kacamataRepository := NewKacamataRepository(go_database.GetConnection())
	ctx := context.Background()
	kacamata := entity.Kacamata{
		Merk:  "helemdeka@gmail.com",
		Price: 5000000,
	}
	result, err := kacamataRepository.Update(ctx, 1, kacamata)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
