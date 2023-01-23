package util

import "log"

func Boom(err error) {
	if err != nil {
		log.Fatalf("A grave mistake was made here: %v", err)
	}
}
