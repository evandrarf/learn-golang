package repository

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
		Email:   "repository@test.com",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}