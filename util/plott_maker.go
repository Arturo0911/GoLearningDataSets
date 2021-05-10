package util

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func MakerPlot(mainFile string) {

	file, err := os.Open(mainFile)
	if err == io.EOF {
		log.Fatal(err)
	}
	defer file.Close()

	irisDF := dataframe.ReadCSV(file)
	fmt.Println(irisDF)

}
