package main

import "os"

type byteString []byte

func writeToFile(name string, data byteString) error {
	return os.WriteFile(name, data, os.ModeAppend)
}

func getDataFromFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
