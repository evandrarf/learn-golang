package go_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('rudi', 'Rudi')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "epan" + strconv.Itoa(i) +"@coba.com"
		comment := "iya bang"

		res, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new customer: ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "radit" + strconv.Itoa(i) +"@coba.com"
		comment := "iya bang" + strconv.Itoa(i)

		_, err = tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}
	}
	
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}