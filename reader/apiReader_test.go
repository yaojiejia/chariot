package Reader

import (
	"testing"
)

func TestRead(t *testing.T) {
	r := NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")

	data, _ := r.Read()
	if data == "" {
		t.Errorf("data is empty")
	}
}

func TestCache(t *testing.T) {
	r := NewAPIReader("https://api.mockae.com/fakeapi/products/2", "")

	r.ReadAndCache()
	id, _ := r.Cache.Get("id")
	expectedID := "2"

	if id != expectedID {
		t.Errorf("id is different, got: %s, want: %s", id, expectedID)
	}
}
