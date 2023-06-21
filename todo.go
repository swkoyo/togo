package main

import "fmt"

type Todo struct {
	task string
	done bool
}

func (t *Todo) Toggle() {
	t.done = !t.done
}

func (t *Todo) Print() {
	if t.done {
		fmt.Printf("[x] %s\n", t.task)
	} else {
		fmt.Printf("[ ] %s\n", t.task)
	}
}

type TodoList struct {
	todos []Todo
}

func (t *TodoList) Add(task string) {
	t.todos = append(t.todos, Todo{task, false})
}

func (t *TodoList) Remove(index int) {
	t.todos = append(t.todos[:index], t.todos[index+1:]...)
}

func (t *TodoList) Toggle(index int) {
	t.todos[index].Toggle()
}

func (t *TodoList) ClearDone() {
    todos := make([]Todo, 0)
    for _, todo := range t.todos {
        if !todo.done {
            todos = append(todos, todo)
        }
    }
    t.todos = todos
}

func (t *TodoList) ClearAll() {
    t.todos = []Todo{}
}

func (t *TodoList) Size() int {
    return len(t.todos)
}

func (t *TodoList) HasDone() bool {
    result := false
    for _, todo := range t.todos {
        if todo.done {
            result = true
            break
        }
    }
    return result
}

func (t *TodoList) IsEmpty() bool {
    return t.Size() == 0
}

func (t *TodoList) Print() {
    fmt.Println("---------- TODO LIST ----------")
	if len(t.todos) == 0 {
		fmt.Println("No tasks")
	} else {
		for i, todo := range t.todos {
			fmt.Printf("%d. ", i+1)
			todo.Print()
		}
	}
    fmt.Printf("-------------------------------\n\n")
}
