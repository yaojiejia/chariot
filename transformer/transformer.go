package Transformer

import (
	"github.com/yaojiejia/chariot/lib"
)

type Transformer struct {
	cache *lib.Cache
}

func NewTransformer(cache *lib.Cache) *Transformer {
	return &Transformer{
		cache: cache,
	}
}

func (t *Transformer) RemovingDuplicates() {
	// Get all keys from the cache
	keys, err := t.cache.GetKeys()
	if err != nil {
		panic(err)
	}

	// Remove duplicates
	for _, key := range keys {
		value, err := t.cache.Get(key)
		if err != nil {
			panic(err)
		}

		// Check if the value is unique
		if t.cache.IsUnique(value) {
			t.cache.Set(key, value)
		}
	}
}
