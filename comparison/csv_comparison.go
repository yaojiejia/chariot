package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	Reader "github.com/yaojiejia/chariot/reader"
)

func main() {
	parallelAccess()
	// sequentialAccess()

}

func parallelAccess() {
	start := time.Now()
	c := Reader.NewCSVReader("melb_data.csv", "melb_data.csv")

	c.ReadAndCache()
	fmt.Println(c.Get("Rooms"))
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func sequentialAccess() {

	start := time.Now()
	filePath := "../melb_data.csv"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV records: %v\n", err)
		return
	}

	if len(records) < 1 {
		fmt.Println("CSV file is empty.")
		return
	}

	headers := records[0]

	targetColumn := "Rooms"

	colIndex := -1
	for i, header := range headers {
		if header == targetColumn {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		fmt.Printf("Column '%s' not found in CSV headers.\n", targetColumn)
		return
	}

	fmt.Printf("Values in column '%s':\n", targetColumn)
	for _, row := range records[1:] { // Skip header row
		if colIndex < len(row) {
			fmt.Println(row[colIndex])
		} else {
			fmt.Println("N/A") // Handle missing columns
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
