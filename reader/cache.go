package Reader

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Cache is a struct that holds the data from the CSV file
type Cache struct {
	client *redis.Client
}

// NewCache is a constructor for the Cache struct, return a new Cache object
func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Cache{
		client: rdb,
	}
}



// Set is a method that sets the key-value pair in the cache
func (c *Cache) Set(key string, value string) error {
	err := c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil

}

// Get is a method that gets the value from the cache
func (c *Cache) Get(key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
