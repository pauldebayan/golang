package main // mandatory
import "fmt"

// Starting point - main

var taskOne = "Task list 1"

// Alternatively:
var taskTwo = "Task List 2"

// Arrays and Slices:
var taskItems = []string{taskOne, taskTwo}

func main() {
	fmt.Print("Hello World!\n") //printn will give the new line by default

	fmt.Println(taskOne)
	fmt.Println(taskTwo)

	fmt.Println("Tasks: ", taskItems)

	// Loops:

	// for index, task := range taskItems { //  index -> _
	// 	//fmt.Println(index+1, ".", task) // Space around '.'
	// 	fmt.Printf("%d. %s\n", index+1, task)
	// }

	// Functions:
	var maxItems = 20
	printTasks(taskItems, maxItems)

	// Function Parameters:

	taskItems = addTasks("Go for a run")

	printTasks(taskItems, maxItems)

	// Compilation error if variables are not used

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
