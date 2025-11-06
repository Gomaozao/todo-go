package main

import (
	"errors"
	"fmt"
	"time"
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
