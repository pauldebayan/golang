package main // mandatory
import (
	"fmt"
	"net/http"
)

// Starting point - main

var taskOne = "Task list 1"

// Alternatively:
var taskTwo = "Task List 2"

// Arrays and Slices:
var taskItems = []string{taskOne, taskTwo}

// func main() {
// 	fmt.Print("Hello World!\n") //printn will give the new line by default

// 	fmt.Println(taskOne)
// 	fmt.Println(taskTwo)

// 	fmt.Println("Tasks: ", taskItems)

// 	// Loops:

// 	// for index, task := range taskItems { //  index -> _
// 	// 	//fmt.Println(index+1, ".", task) // Space around '.'
// 	// 	fmt.Printf("%d. %s\n", index+1, task)
// 	// }

// 	// Functions:
// 	var maxItems = 20
// 	printTasks(taskItems, maxItems)

// 	// Function Parameters:

// 	taskItems = addTasks("Go for a run")

// 	printTasks(taskItems, maxItems)

// 	// Compilation error if variables are not used

// }

func main() {
	// Local Server
	fmt.Println("Http Server")

	http.HandleFunc("/hello-go", func(w http.ResponseWriter, r *http.Request) {
		var greeting = "Hello user welcome to out ToDo list App!"
		fmt.Fprintf(w, greeting)
	})

	http.HandleFunc("/show-tasks", func(w http.ResponseWriter, r *http.Request) {

		for _, item := range taskItems {
			fmt.Fprintf(w, item)
		}
		//  curl localhost:8080/show-tasks will also work
	})

	// Similar to 'fmt' there is 'http', golang will automatically import it
	http.ListenAndServe(":8080", nil) // 2nd parameter is the handler - what if something goes wrong

}

func printTasks(taskItems []string, itemLimit int) {
	// taskItems is not global
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTasks(newTask string) []string { // we can return multiple values using ()
	// taskItems will not get updated
	// updatedTaskList will only work in this scope

	var updatedTaskList = append(taskItems, newTask)
	//printTasks(updatedTaskList, 0)

	// Return from function

	return updatedTaskList

}

// Building a web API

// go build <file>.go
// go run <file>.go
