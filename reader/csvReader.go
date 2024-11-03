package Reader

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/yaojiejia/chariot/lib"
)

// CSVReader is a struct for reading CSV files and caching their data.
type CSVReader struct {
	FileName string
	FilePath string
	Cache    *lib.Cache
}

// NewCSVReader is a constructor for the CSVReader struct, returning a new CSVReader object.
func NewCSVReader(fileName, filePath string) *CSVReader {
	return &CSVReader{
		FileName: fileName,
		FilePath: filePath,
		Cache:    lib.NewCache(),
	}
}

// Read reads all records from the CSV file.
func (c *CSVReader) Read() ([][]string, error) {
	f, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

// ReadAndCache reads the CSV data and caches each column's data concurrently.
func (c *CSVReader) ReadAndCache() error {
	// Read data from CSV
	records, err := c.Read()
	if err != nil {
		return err
	}

	// Ensure we have at least headers
	if len(records) < 1 {
		return errors.New("no records found in CSV")
	}

	// Get headers from first row
	headers := records[0]

	// Initialize thread-safe columnData
	columnData := lib.NewSafeColumnData(headers)

	// Use a WaitGroup to wait for all row processing goroutines to finish
	var wg sync.WaitGroup

	// Define a concurrency limit to avoid spawning too many goroutines
	concurrencyLimit := 100 // Adjust based on your system's capabilities
	semaphore := make(chan struct{}, concurrencyLimit)

	// Process each data row concurrently
	for _, row := range records[1:] {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot

		// Launch a goroutine to process the row
		go func(r []string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the slot

			for i, value := range r {
				if i < len(headers) {
					columnData.Add(headers[i], value)
				}
			}
		}(row)
	}

	// Wait for all row processing goroutines to complete
	wg.Wait()

	// Now, marshal JSON and store in cache concurrently
	var cacheWg sync.WaitGroup
	cacheErrors := make(chan error, len(columnData.GetData()))
	cacheConcurrencyLimit := 50 // Adjust as needed
	cacheSemaphore := make(chan struct{}, cacheConcurrencyLimit)

	for header, values := range columnData.GetData() {
		cacheWg.Add(1)
		cacheSemaphore <- struct{}{} // Acquire a slot

		// Launch a goroutine to marshal and cache the column data
		go func(h string, v []string) {
			defer cacheWg.Done()
			defer func() { <-cacheSemaphore }() // Release the slot

			// Marshal to JSON
			jsonData, err := json.Marshal(v)
			if err != nil {
				cacheErrors <- fmt.Errorf("error marshaling column '%s': %w", h, err)
				return
			}

			// Store in cache
			err = c.Cache.Set(h, string(jsonData))
			if err != nil {
				cacheErrors <- fmt.Errorf("error setting cache for column '%s': %w", h, err)
				return
			}
		}(header, values)
	}

	// Wait for all cache operations to complete
	cacheWg.Wait()
	close(cacheErrors)

	// Check for errors
	for err := range cacheErrors {
		if err != nil {
			return err
		}
	}

	return nil
}

// Get retrieves the cached JSON string for the specified key.
func (c *CSVReader) Get(key string) (string, error) {
	return c.Cache.Get(key)
}
