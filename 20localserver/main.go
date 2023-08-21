package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
)

func main() {
	Getrequest()
}
func Getrequest() {
	const mylocalurl = "http://localhost:8000/get"
	// const mylocalurl = "https://harishsirikoti.online"
	response, err := http.Get(mylocalurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	Content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content in host:\n", string(Content))
}
