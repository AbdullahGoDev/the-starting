package main

import (
	"fmt"
	"testpro/test"
)

func main() {

	file := "myTestFile.txt"

	//contant := "hi bro \nhow are you today"

	test.CreateFileAndWrite(file, "just for test")
	fmt.Println(test.CheckFile(file))
	fmt.Println(test.ReadFile(file))

}
