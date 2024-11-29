package test

import "os"

func CheckFile(file string) bool {
	_, filec := os.Stat(file)

	if os.IsNotExist(filec) {
		return false
	}
	return filec == nil
}
