package test

import "os"

func CreateFile(file string) {
	myFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	defer myFile.Close()

}
