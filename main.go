package main

import "fmt"
import "time"

// TODO
// Persist data


type task struct {
	name         string
	completed    bool
	creationTime time.Time
}

func createTask(name string, taskList *[]task) {
	newTask := task{
		name: name,
		completed:    false,
		creationTime: time.Now()}
	*taskList = append(*taskList, newTask)
}

func main() {
	taskList := []task{}
	fmt.Println("Hi from main!")
	createTask("test", &taskList)
	fmt.Println(taskList)
}
