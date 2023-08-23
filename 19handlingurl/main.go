package main

import (
	"fmt"
	"net/url"
)

const Myurl string = "https://courses.learncodeonline.in/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {

	responce, _ := url.Parse(Myurl)
	fmt.Println(responce.Scheme)
	fmt.Println(responce.Host)
	fmt.Println(responce.Path)
	fmt.Println(responce.Port())
	fmt.Println(responce.RawQuery)
	qprarams := responce.Query()
	fmt.Println("value from qparams:", qprarams["coursename"])
	for _, val := range qprarams {
		fmt.Println(val)
	}
}
