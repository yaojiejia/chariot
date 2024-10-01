package reader

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	FileName string
	FilePath string
	Encoder  *csv.Reader
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

func (c *CSVReader) Connect() error {
	return nil
}

func (c *CSVReader) Display() string {
	return "Displaying the cache read from csv"
}
