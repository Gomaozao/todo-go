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
	TodoName    string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		TodoName:    title,
		IsCompleted: false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validadeIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) deleteTodo(index int) error {
	t := *todos

	if err := t.validadeIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validadeIndex(index); err != nil {
		return err
	}
	completed := t[index].IsCompleted
	if !completed {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].IsCompleted = !completed
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validadeIndex(index); err != nil {
		return err
	}
	t[index].TodoName = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "TodoName", "isCompleted", "CreatedAt", "CompletedAt")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.IsCompleted {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.TodoName, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
