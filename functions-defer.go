package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("dealers.txt")
	defer closeFile(file)
	writeToFile(file, "A1 Auto")
}

func createFile(path string) *os.File {
	fmt.Println("Creating file...")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, dealerName string) {
	fmt.Println("Writing to file...")
	fmt.Fprintln(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file...")
	file.Close()
}
