package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func FixingCSVFiles(pathFile string) {

	csvFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	readerCSV := csv.NewReader(csvFile)

	csvReader, err := readerCSV.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, elements := range csvReader {
		fmt.Println("Elements in a row: ", elements)
		fmt.Println("number of elements per row: ", len(elements))
	}

}
