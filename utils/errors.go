package utils

import (
	"log"
	"os"
)

func CheckError(err error, message string) {
	if err != nil {
		log.Fatal(err, ": "+message)
		os.Exit(1)
	}

}
