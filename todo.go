package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	todoName    string
	isCompleted bool
	createdAt   time.Time
	completedAt *time.Time
}
type todos []Todo

func (todos *todos) add(title string) {
	todo := Todo{
		todoName:    title,
		isCompleted: false,
		completedAt: nil,
		createdAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *todos) validadeIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *todos) deleteTodo(index int) error {
	t := *todos

	if err := t.validadeIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

// method that takes an index as input and turns whatever
func (todos *todos) toggle(index int) error {
	t := *todos
	if err := t.validadeIndex(index); err != nil {
		return err
	}
	completed := t[index].isCompleted
	if !completed {
		completionTime := time.Now()
		t[index].completedAt = &completionTime
	}
	t[index].isCompleted = !completed
	return nil
}

func (todos *todos) edit(index int, title string) error {
	t := *todos
	if err := t.validadeIndex(index); err != nil {
		return err
	}
	t[index].todoName = title

	return nil
}

func (todos *todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "TodoName", "isCompleted", "CreatedAt", "CompletedAt")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.isCompleted {
			completed = "✅"
			if t.completedAt != nil {
				completedAt = t.completedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.todoName, completed, t.createdAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
