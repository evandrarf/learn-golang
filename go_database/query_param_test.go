package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestQueryParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	ctx := context.Background()

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err = rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login: ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}