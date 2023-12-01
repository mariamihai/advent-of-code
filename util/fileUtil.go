package util

import (
	"os"
)

func ReadFile(filename string) *os.File {
	file, err := os.Open(filename)
	Boom(err)

	return file
}

func CloseFile() func(file *os.File) {
	return func(file *os.File) {
		err := file.Close()
		Boom(err)
	}
}
