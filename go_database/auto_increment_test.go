package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "coba@coba.com"
	comment := "Hai bang"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer with id:", insertedId)
}