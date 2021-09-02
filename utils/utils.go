package utils

import "log"

func ExitIfError(message string, err error) {
	if err != nil {
		log.Fatalf("%s, %v", message, err)
	}
}
