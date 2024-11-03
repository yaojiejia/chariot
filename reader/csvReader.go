package Reader

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

type CSVReader struct {
	FileName string
	FilePath string
	Encoder  *csv.Reader
	Cache    *Cache
}

func NewCSVReader(fileName, filePath string) *CSVReader {
	return &CSVReader{
		FileName: fileName,
		FilePath: filePath,
		Cache:    NewCache(),
	}
}

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

func (c *CSVReader) ReadAndCache() error {
	// Read data from CSV
	records, err := c.Read()
	if err != nil {
		return err
	}

	// Ensure we have at least headers
	if len(records) < 1 {
		return nil
	}

	// Get headers from first row
	headers := records[0]

	// Create maps for each column
	columnData := make(map[string][]string)
	for _, header := range headers {
		columnData[header] = make([]string, 0)
	}

	// Process each data row
	for _, row := range records[1:] {
		// Add values to corresponding column arrays
		for i, value := range row {
			if i < len(headers) {
				columnData[headers[i]] = append(columnData[headers[i]], value)
			}
		}
	}

	// Store each column's data in cache
	for header, values := range columnData {
		// Convert values array to JSON string
		jsonData, err := json.Marshal(values)
		if err != nil {
			return err
		}

		// Store in cache with column header as key
		err = c.Cache.Set(header, string(jsonData))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CSVReader) Get(key string) (string, error) {
	return c.Cache.Get(key)
}
