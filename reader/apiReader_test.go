package Reader

import "fmt"

func readerTest() {
	r := NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")

	fmt.Println(r.Read())
}
