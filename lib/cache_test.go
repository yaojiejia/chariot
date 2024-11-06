package lib

import "testing"

func Test_cache(t *testing.T) {
	r := NewCache()
	r.Set("hello", "world")
	val, _ := r.Get("hello")

	if val != "world" {
		t.Errorf("value is different, got: %s, want: %s", val, "world")
	}
}

func Test_cache_keys(t *testing.T) {
	r := NewCache()
	r.Flush()
	r.Set("hello", "world")
	r.Set("foo", "bar")
	keys, _ := r.GetKeys()

	if len(keys) != 2 {
		t.Errorf("keys length is different, got: %d, want: %d", len(keys), 2)
	}
}
