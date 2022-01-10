package main

import(
	"fmt"
	"io/ioutil"
	)
func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) //here we need to cast this b type to string, to read content of file instead of binary

	//write
	stringContent := string(b) + "\n more new stuff"
	fmt.Println("write to file", stringContent)
	err = ioutil.WriteFile("output.txt", []byte(stringContent), 0644)
	if err != nil {
		panic(err)
	}
}