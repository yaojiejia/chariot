package lib

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Cache is a struct that holds the data from the CSV file
type Cache struct {
	Client *redis.Client
}

// NewCache is a constructor for the Cache struct, return a new Cache object
func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Cache{
		Client: rdb,
	}
}

// Set is a method that sets the key-value pair in the cache
func (c *Cache) Set(key string, value string) error {
	err := c.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil

}

// Get is a method that gets the value from the cache
func (c *Cache) Get(key string) (string, error) {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// GetKeys is a method that gets all the keys from the cache
func (c *Cache) GetKeys() ([]string, error) {
	var (
		cursor uint64 = 0
		keys   []string
		err    error
	)
	for {
		var batch []string
		batch, cursor, err = c.Client.Scan(ctx, cursor, "*", int64(10000)).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, batch...)

		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

func (c *Cache) Flush() error {
	err := c.Client.FlushDB(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
