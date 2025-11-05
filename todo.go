package main

import "flag"

// task is just a struct
// todo is also a struct

var hFlag = flag.Uint("h", 1, "") // THIRD argument should be all of the other flags explanation

func (t todo) createTodoList(listName string) {
	if listName == t.TodoName {
		//Output to the console the TodoList and it's  tasks
	}
	t.TodoName = listName
}

type todo struct {
	task
	tasks    []string // slice should be of type string or type task?
	TodoName string
}

type task struct {
	TaskName    string
	IsCompleted bool
}
