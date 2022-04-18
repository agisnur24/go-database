package comment

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "topengrazor@gmail.com",
		Comment: "Test Repositiry",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestCommentDelete(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Id: 5,
	}

	result, err := commentRepository.Delete(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "helemdeka@gmail.com",
		Comment: "Test Repositiry",
	}
	result, err := commentRepository.Update(ctx, 1, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
