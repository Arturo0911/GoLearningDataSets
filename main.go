package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

const (
	PATH_FILE      = "datasets/COVID.csv"
	TRAIN_FILE     = "covid_train.csv"
	TEST_FILE      = "covid_test.csv"
	PATH_IMAGES    = "images/"
	IRIS_PATH_FILE = "datasets/iris_headers.csv"
)

func checkMapSet(cases map[string]int, element string) bool {
	var result bool = false
	for key := range cases {
		if key == element {
			result = true
			break
		}
	}
	return result
}

func ReadingFile() {
	file, err := os.Open(PATH_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	covidDF := dataframe.ReadCSV(file)

	mapGeneral := make(map[string]map[string]int)

	for _, element := range covidDF.Names() {

		mapSet := make(map[string]int)
		if element != "Timestamp" {
			for _, value := range covidDF.Col(element).Records() {

				if checkMapSet(mapSet, value) {
					mapSet[value] += 1
				} else {
					mapSet[value] = 1
				}
			}
			mapGeneral[element] = mapSet
		}
	}

	for key, value := range mapGeneral {
		fmt.Println(key, ": ", value)
	}
}

// Reading and practicing about iris_file
func MakingFilesIris() {

	irisFile, err := os.Open(IRIS_PATH_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)

	trainingNum := (4 * irisDF.Nrow()) / 5
	testNum := (irisDF.Nrow()) / 5

	if trainingNum+testNum < irisDF.Nrow() {
		trainingNum++
	}

	// training sets
	trainSet := make([]int, trainingNum)
	testSet := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainSet[i] = i
	}

	for j := 0; j < testNum; j++ {
		testSet[j] = trainingNum + j
	}

}

func main() {
	//ReadingFile()
	MakingFilesIris()
}
