package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Menu is struct
type Menu struct {
	IDMenu   int
	Nama     string
	Kategori string
	Harga    int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/restaurant")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func tampilMenu() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT idMenu, nama, kategori, harga FROM menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []Menu

	for rows.Next() {
		var each = Menu{}

		err = rows.Scan(&each.IDMenu, &each.Nama, &each.Kategori, &each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println("Menu: ", each.Nama, each.Harga)
	}

}

func main() {
	tampilMenu()
}
