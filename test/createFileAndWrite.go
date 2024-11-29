package test

import (
	"os"
)

func CreateFileAndWrite(file, contant string) {
	myFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer myFile.Close()

	WriteString(myFile, contant)

	err = myFile.Sync()

}
