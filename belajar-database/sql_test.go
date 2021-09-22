package belajar_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer (id, name) VALUES ('fahmi','fahmi')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Success")

}

func TestQuerySql(t *testing.T) {
	fmt.Println("=========================")
	fmt.Println("Result TestQuerySql")

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, " - ", name)

	}
	fmt.Println("=========================")
}

func TestQuerySqlComplex(t *testing.T) {
	fmt.Println("=========================")
	fmt.Println("Result TestQuerySqlComplex")

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT `id`, `name`, `email`, `balance`, `rating`, `created_at`, `birth_date`, `married` FROM `customer`"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		// var id, name, email string **Cannot handle null type**
		var id, name string
		var email sql.NullString //Can handle null type
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime //Can handle null type
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("-------")
		fmt.Println("ID : ", id)
		fmt.Println("name : ", name)
		fmt.Println("email : ", email)
		fmt.Println("balance : ", balance)
		fmt.Println("rating : ", rating)
		fmt.Println("created_at : ", created_at)
		fmt.Println("birth_date : ", birth_date.Time) //Get spesific value from struct sql.NullTime
		fmt.Println("married : ", married)

	}

	fmt.Println("=========================")
}

func TestQuerySqlinjection(t *testing.T) {
	fmt.Println("=========================")
	fmt.Println("Result TestQuerySqlinjection")

	db := GetConnection()
	defer db.Close()

	// Normal
	// username := "admin"
	// password := "admin"
	// Injection
	username := "admin'; #"
	password := "xxx"

	ctx := context.Background()
	script := "SELECT * FROM `user` where username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("LOGIN BERHASIL")
	} else {
		fmt.Println("LOGIN GAGAL")
	}

	fmt.Println("=========================")
}

func TestQuerySqlinjectionSafe(t *testing.T) {
	fmt.Println("=========================")
	fmt.Println("Result TestQuerySqlinjectionSafe")

	db := GetConnection()
	defer db.Close()

	// Normal
	// username := "admin"
	// password := "admin"
	// Injection
	username := "admin'; #"
	password := "xxx"

	ctx := context.Background()
	script := "SELECT * FROM `user` where username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("LOGIN BERHASIL")
	} else {
		fmt.Println("LOGIN GAGAL")
	}

	fmt.Println("=========================")
}

func TestPrepareStatment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statment, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statment.Close()

	for i := 0; i < 10; i++ {
		email := "user" + strconv.Itoa(i) + "@mail.com"
		comment := "comment ke-" + strconv.Itoa(i)

		result, err := statment.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Insert ID Ke - ", id, " Sukses")

	}

}
