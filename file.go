package main

import "os"

type byteString []byte

func writeToFile(name string, data byteString) error {
	return os.WriteFile(name, data, 0666)
}

func getDataFromFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
