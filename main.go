package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"Country", "city", "state"},
		{"USA", "San Joe", "California"},
		{"USA", "Dallas", "Texas"},
		{"India", "Mumbai", "Maharashtra"},
	}

	csvfile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	cswriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = cswriter.Write(row)

	}
	cswriter.Flush()
	csvfile.Close()
}
