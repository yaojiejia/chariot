package main

import (
	"fmt"

	// Reader "github.com/yaojiejia/chariot/reader"
	"github.com/yaojiejia/chariot/db"
)

func main() {
	// r := Reader.NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")
	// fmt.Println(r.Read())
	p := db.NewPSQL("localhost", "5432", "alex", "jiayaojie0715", "newdb")
	p.Connect()

	tables, err := p.GetTables()
	fmt.Printf("%v", tables)
	if err != nil {
		fmt.Println(err)
	}
	col, err := p.GetColumns("users_test", "public")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", col)

}
