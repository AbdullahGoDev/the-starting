package test

import (
	"fmt"
	"os"
)

func ReadFile(file string) string {

	readFile, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	rere := fmt.Sprintf("%s", readFile)
	return rere
}
