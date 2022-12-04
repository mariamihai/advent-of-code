package util

import (
	"log"
	"os"
)

func ReadFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func CloseFile() func(file *os.File) {
	return func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
