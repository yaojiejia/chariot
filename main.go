package main

import (
	"fmt"

	"github.com/yaojiejia/chariot/reader"
)

func main() {
	r := reader.NewCSVReader("test.csv", "melb_data.csv")
	fmt.Println(r.Read())

}
