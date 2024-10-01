package Reader

import (
	"encoding/csv"
	"os"

	"github.com/olekukonko/tablewriter"
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

func (c *CSVReader) Connect() (error, string) {
	records, err := c.Read()
	if err != nil {
		return err, err.Error()
	}

	c.Cache = NewCache(records)
	// fmt.Println(c.Cache.Data)
	return nil, "stored to the cache!"

}

func (c *CSVReader) Display() string {
	records := c.Cache.Data

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(records[0])

	for _, row := range records[1:] {
		table.Append(row)
	}

	table.Render()
	return "CSV data displayed successfully!"
}
