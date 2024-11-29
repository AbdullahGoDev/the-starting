package test

import "os"

func WriteString(fileName *os.File, Contant string) {

	_, err := fileName.WriteString(Contant)

	if err != nil {
		panic(err)
	}
	defer fileName.Close()

}
