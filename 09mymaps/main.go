package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["GO"] = "GoLang"
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	fmt.Println("single element: ", languages["PY"])
	delete(languages, "JS")
	fmt.Println("languages after removal: ", languages)

	//Sample loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
