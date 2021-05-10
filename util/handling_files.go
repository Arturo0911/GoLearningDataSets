package util

import (
	"log"
	"os"
)

func MakingPredictionsFromFile(trainFile, testFile string) {

	trainer, err := os.Open(trainFile)
	if err != nil {
		log.Fatal(err)
	}
	defer trainer.Close()

}
