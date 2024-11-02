package Reader

import (
	"testing"
)

func readerTest(t *testing.T) {
	r := NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")

	data, _ := r.Read()
	if data == "" {
		t.Errorf("data is empty")
	}
}
