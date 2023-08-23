package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Getrequest()
	// PostjsonRequest()
	PostForm()
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

// Post to local host
func PostjsonRequest() {
	const myposturl = "http://localhost:8000/post"
	postbody := strings.NewReader(`
	{
		"name":"Harish",
		"age":23,
		"JobRole":"Developer"
	}
`)
	reponse, err := http.Post(myposturl, "application/json", postbody)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(reponse.Body)
	defer reponse.Body.Close()
	fmt.Println("Content of Post:\n", string(content))
}
func PostForm() {
	const myposturl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("Name", "Harish")
	data.Add("Age", "23")
	data.Add("Email", "maeh1030@gmail.com")
	response, err := http.PostForm(myposturl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of Postform:\n", string(content))

}
