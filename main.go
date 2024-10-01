package main

import (
	"fmt"

	Reader "github.com/yaojiejia/chariot/reader"
)

func main() {
	r := Reader.NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")
	fmt.Println(r.Read())

}
