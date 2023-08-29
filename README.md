# Golangpractice
	go mod init [module-path] (for dependencies)
	go run [file name with extension] (to run the function)
	go help [] (get document for command help)
	go env GOPATH (to get environment)
	fp (to get print line)
	To get the type{
var username string = "Harish"
fmt.Println(username)
fmt.Printf("Variables is of type: %T \n", username)
} (%T)

	UserInputs{
	func main() {
	    fmt.Println("Welcom to user Input")
	    reader := bufio.NewReader(os.Stdin)
	    fmt.Println("Enter your rating:")
	    input, _ := reader.ReadString('\n')
	    fmt.Println("Your rating:", input)
	    fmt.Printf("Type of rating: %T", input)
	}
}

	Conversion string to int{
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	    if err != nil {
	        fmt.Println(err)
	    } else {
	        fmt.Println("Adding One", numRating+1)
	    }
}
	Print time and formate {
	presentTime := time.Now()
	    fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
	GOOS=”windows” go build (To build file for other operating systems linux,darwin)
	Memory Management
new()	make()
Allocate memory but no INIT	Allocate memory and INIT
You will get a memory address	You will get a memory address
Zeroed storage 	Non zeroed storage 
o	Garbage Collection automatically: The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated date to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. The runtime/debug package SetGCPercent function allows changing this percentage at run time.
 
	Pointer {
	    mynumber := 10
	    var ptr = &mynumber
	    fmt.Println("Pointer address:", ptr)
	    fmt.Println("Pointer Value: ", *ptr)
	    *ptr = *ptr + 5
	    fmt.Println("New mynumber value: ", mynumber)
}
	Dereferencing a pointer means accessing the value of the variable to which the pointer points. *a is the syntax to deference a. 
	Array declaration {
	var fruits [5]string
	    fruits[0] = "Apple"
	    fruits[1] = "Banana"
	    fruits[3] = "Mango"
	    fruits[4] = "grapes"
	    var vegetables = [4]string{"Carrot", "radish", "potato", "beans"}
	    fmt.Println("Array of Fruits :", fruits)
	    fmt.Println("Length of Fruits :", len(fruits)) //we have only 4 but it gives 5(output)
	    fmt.Println("Array of Vegetables :", vegetables)
	    fmt.Println("Length of Vegetables :", len(vegetables))
}

	Slicers: A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.
 {
	    var fruits = []string{"apple", "banana", "grapes", "Mango"}
	    fmt.Printf("Type of fruits: %T\n", fruits)
	    fruits = append(fruits, "apple2.0", "mango2.0")
	    fruits = append(fruits[2:5])
	    fmt.Println("Fruits after slice: ", fruits)
	    fmt.Println("Length of fruits: ", len(fruits)) // Length will change after slice
	    Scores := make([]int, 4) //another syntax
	    Scores[0] = 27697
	    Scores[1] = 67
	    Scores[2] = 8779
	    Scores[3] = 277
	    Scores = append(Scores, 34, 123)
	    sort.Ints(Scores) //Sorting of slicers
	    fmt.Println("To know slices are sorted: ", sort.IntsAreSorted(Scores))
	    fmt.Println(Scores)
}

	creating a slice using make
	func main() {  
	    i := make([]int, 5, 5)
	    fmt.Println(i)
	}
	The values are zeroed by default when a slice is created using make. The above program will output [0 0 0 0 0].
	Appending 2 slices to one:{
	veggies := []string{"potatoes","tomatoes","brinjal"}
	    fruits := []string{"oranges","apples"}
	    food := append(veggies, fruits...)
	    fmt.Println("food:",food)
}
	Remove 1 element form slices{
	courses := []string{"React", "Java", "Python", "Ruby", "C++", "Shift"}
	    index := 2
	    courses = append(courses[:index], courses[index+1:]...)//triple dots are imp
	    fmt.Println("Removed index value from slice: ", courses)
	}

	Maps {
	languages := make(map[string]string)
	    languages["GO"] = "GoLang"
	    languages["PY"] = "Python"
	    languages["JS"] = "JavaScript"
	    languages["TS"] = "TypeScript"
	    fmt.Println("single element: ", languages["PY"])
	    delete(languages, "JS")
}

Day 3
	Loops and goto{
	weekdays := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	    for i, day := range weekdays {
	        fmt.Printf("Index %v Day %v\n", i, day)
	    }
	    num := 1
	    for num < 10 {
	        if num == 5 {
	            goto somelink
	        }
	        fmt.Println(num)
	        num++
	    }
	somelink:
	    fmt.Println("this is goto")
}
	Passing unlimited values to function{
	func main() {
	    adderres, mess := adder(3, 4, 5, 6)
	    fmt.Println(mess, adderres)
	}
	func adder(values ...int) (int, string) {
	    total := 0
	    for _, val := range values {
	        total += val
	    }
	    return total, "Adder Results"
	}
}
Day 7
	Defer:- statements are executed in a last-in, first-out (LIFO) order, which means the most recently deferred function call will be executed first when the surrounding function completes.
	Files:{
	func main() {
	    content := "This file is created by harish using Go lang"
	
	    file, err := os.Create("./autofile.txt")
	    if err != nil {
	        panic(err)
	    }
	    length, err := io.WriteString(file, content)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println("The lenght of the file: ", length)
	    file.Close()
	    readfile("./autofile.txt")
	}
	func readfile(name string) {
	    readcontent, err := ioutil.ReadFile(name)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println("File content inside: \n", string(readcontent))
	}
}
	Nested Structs{
	type Person struct {
	    name    string
	    age     int
	    address Address
	} //nested structs
	type Address struct {
	    state string
	    city  string
	}
	p1 := Person{name: "Harish", age: 23, address: Address{state: "Telangana", city: "Hyderbad"}}						// in function
	
Day8
	Interfaces: When a type provides definition for all the methods in the interface, it is said to implement the interface.

o	An interface that has zero methods is called an empty interface. It is represented as interface{}.
	package main
	import ("fmt")
	func describe(i interface{}) {  
	    fmt.Printf("Type = %T, value = %v\n", i, i)
	}
	func main() {  
	    s := "Hello World"
	    describe(s)
	    i := 55
	    describe(i)
	    strt := struct {
	        name string
	    }{
	        name: "Naveen R",
	    }
	    describe(strt)
	}
o	Type swich: A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. 
	func findType(i interface{}) {  
	    switch i.(type) {
	    case string:
	        fmt.Printf("I am a string and my value is %s\n", i.(string))
	    case int:
	        fmt.Printf("I am an int and my value is %d\n", i.(int))
	    default:
	        fmt.Printf("Unknown type\n")
	    }
	}
	func main() {  
	    findType("Naveen")
	    findType(77)
	    findType(89.98)
	}

Day9
	Web request {
	responce, err := http.Get(url2)
	    if err != nil {
	        panic(err)
	    }
	    fmt.Printf("reponse type: %T\n", responce)
	    defer responce.Body.Close()
	    databytes, err := ioutil.ReadAll(responce.Body)
	    if err != nil {
	        panic(err)
	    }
	    Cotent := string(databytes)
	    fmt.Println("Content :\n", Cotent)
	}
	Handling of webrequests:{
	    responce, _ := url.Parse(Myurl)
	    fmt.Println(responce.Scheme)
	    fmt.Println(responce.Host)
	    fmt.Println(responce.Path)
	    fmt.Println(responce.Port())
	    fmt.Println(responce.RawQuery)
	    qprarams := responce.Query()
	    fmt.Println("value from qparams:", qprarams["coursename"])
	    for _, val := range qprarams {
	        fmt.Println(val)
	    }
	}
	Post in localhost:{
	const myposturl = "http://localhost:8000/post"
	    postbody := strings.NewReader(`
	    {
	        "name":"Harish",
	        "age":23,
	        "JobRole":"Developer"
	    }
	`)
	    reponse, err := http.Post(myposturl, "application/json", postbody)
	    if err != nil {
	        panic(err)
	    }
	    content, _ := ioutil.ReadAll(reponse.Body)
	    defer reponse.Body.Close()
	    fmt.Println("Content of Post:\n", string(content))
	}
Day10
	To Create JSON
	type emp struct {
	    Name    string `json:"Employe name"` //print to Employe name
	    Age     int
	    Salary  int      `json:"-"`              //Dont print
	    Hobbies []string `json:"tags,omitempty"` //if nil no print
	}
	
	func main() {
	    Employee := []emp{
	        {"Harish", 23, 20000, []string{"Kabbadi", "Dance"}},
	        {"Chitti", 25, 22000, []string{"Drawing", "Using phone"}},
	        {"Harish", 26, 30000, nil},
	    }
	    finalJson, err := json.MarshalIndent(Employee, "", "\t") //to print good formate
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println("Require JSON:", string(finalJson))
	}
	

Day 11
	Concurrency is an inherent part of the Go programming language. Concurrency is handled in Go using Goroutines and channels.
	Goroutines are functions or methods that run concurrently with other functions or methods. 
	Multiplexing Goroutines: Imagine you have many Goroutines, but there are only a limited number of OS threads available (like having only one chef in the kitchen). The Go runtime, which manages Goroutines, cleverly switches the attention of these OS threads between different Goroutines. This is similar to the chef rotating between different orders.
	Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines.
	func hello(done chan bool) {  
	    fmt.Println("Hello world goroutine")
	    done <- true}
	func main() {  
	    done := make(chan bool)
	    go hello(done)
	    <-done
	    fmt.Println("main function")}
	In the above program, we create a done bool channel in line no. 12 and pass it as a parameter to the hello Goroutine. In line no. 14 we are receiving data from the done channel. This line of code is blocking which means that until some Goroutine writes data to the done channel, the control will not move to the next line of code. Hence this eliminates the need for the time.Sleep which was present in the original program to prevent the main Goroutine from exiting.
	Deadlock: One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
	Buffered Channels and Worker Pools: It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.
	The length of the buffered channel is the number of elements currently queued in it.
	The capacity of a buffered channel is the number of values that the channel can hold. 
	Select: The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random. 
o	If a default case is present, this deadlock will not happen since the default case will be executed when no other case is ready. 
	The "critical section" is like a rule that says: "Hey, only one person (Goroutine) can change this thing at a time."
	A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.
o	Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock.
o	mutex.Lock()  
o	x = x + 1  
o	mutex.Unlock()
o	If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.

	Composition by embedding structs
o	Composition can be achieved in Go is by embedding one struct type into another.
