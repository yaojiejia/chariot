package Reader

// Cache is a struct that holds the data from the CSV file
type Cache struct {
	Data [][]string
}

// NewCache is a constructor for the Cache struct, return a new Cache object
func NewCache(data [][]string) *Cache {
	return &Cache{
		Data: data,
	}
}
