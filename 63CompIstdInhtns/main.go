package main

import "fmt"

type author struct {
	fname string
	lname string
	bio   string
}
type blogpost struct {
	title   string
	content string
	author
}

func (e author) fulname() string {
	return e.fname + " " + e.lname
}
func (b blogpost) details() {
	fmt.Println("Name of author:", b.fulname())
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Bio :", b.bio)
}

type website struct {
	blogposts []blogpost
}

func (web website) contents() {
	fmt.Println("Contents of website")
	for _, v := range web.blogposts {
		v.details()
		fmt.Println("")
	}
}
func main() {
	Auth := author{"Harish", "Sirikoti", "Testing of bio print"}
	Detail := blogpost{"Random Title", "Random good content", Auth}
	blogPost2 := blogpost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		Auth,
	}
	blogPost3 := blogpost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		Auth,
	}
	// Detail.details()
	w := website{
		blogposts: []blogpost{Detail, blogPost2, blogPost3},
	}
	w.contents()
}
