package main

import (

	// Reader "github.com/yaojiejia/chariot/reader"

	"fmt"

	Reader "github.com/yaojiejia/chariot/reader"
)

func main() {
	r := Reader.NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")

	r.ReadAndCache()
	id, _ := r.Cache.Get("id")
	fmt.Println(id)
	// var c lib.Config
	// c.GetConfig()
	// p := db.NewPSQL(c.Host, c.Port, c.User, c.Password, c.Database)
	// p.Connect()

	// tables, err := p.GetTables()
	// fmt.Printf("%v", tables)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// col, err := p.GetColumns("users_test", "public")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%v", col)

}
