package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url = "https://lco.dev"
const url2 = "https://harishsirikoti.online/"

func main() {
	responce, err := http.Get(url2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("reponse type: %T\n", responce)
	defer responce.Body.Close()
	databytes, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	Cotent := string(databytes)
	fmt.Println("Content :\n", Cotent)
}
