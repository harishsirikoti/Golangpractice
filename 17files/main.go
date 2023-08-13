package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This file is created by harish using Go lang"

	file, err := os.Create("./autofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The lenght of the file: ", length)
	file.Close()
	readfile("./autofile.txt")
}
func readfile(name string) {
	readcontent, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println("File content inside: \n", string(readcontent))
}
