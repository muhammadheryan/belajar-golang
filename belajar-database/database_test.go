package belajar_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	defer db.Close()
}
