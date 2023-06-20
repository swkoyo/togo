package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

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

func (t *TodoList) Print() {
	fmt.Println("Todo List: ")
	if len(t.todos) == 0 {
		fmt.Println("No tasks")
	} else {
		for i, todo := range t.todos {
			fmt.Printf("%d. ", i+1)
			todo.Print()
		}
		fmt.Println()
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var COMMANDS = map[int]string{
	1: "Add task",
	2: "Remove task",
	3: "Toggle task",
	4: "Exit",
}

func PrintMenu() {
	fmt.Println("---------------------")
	fmt.Println("Please choose an option:")
	for i := 1; i <= len(COMMANDS); i++ {
		fmt.Printf("%d. %s\n", i, COMMANDS[i])
	}
}

func GetInput() (int, error) {
    scanner := bufio.NewScanner(os.Stdin)
	var input string
    if scanner.Scan() {
        input = scanner.Text()
    }
	val, err := strconv.Atoi(input)
	if err != nil {
        return 0, errors.New("Enter a valid input: ")
	}
	if val < 1 || val > len(COMMANDS) {
        return 0, errors.New("Enter a valid input: ")
	}
	return val, nil
}

func Runner(todoList *TodoList) {
	ClearScreen()
	todoList.Print()
	PrintMenu()
	input, err := GetInput()
	for err != nil {
		fmt.Print(err)
		input, err = GetInput()
	}

    scanner := bufio.NewScanner(os.Stdin)

	if input == 1 {
		fmt.Printf("Enter task: ")
        var task string
        if scanner.Scan() {
            task = scanner.Text()
        }
        fmt.Println(task)
		todoList.Add(task)
	} else if input == 2 {
		fmt.Printf("Enter task number: ")
		var taskNum int
		fmt.Scanln(&taskNum)
		todoList.Remove(taskNum - 1)
	} else if input == 3 {
		fmt.Printf("Enter task number: ")
		var taskNum int
		fmt.Scanln(&taskNum)
		todoList.Toggle(taskNum - 1)
	} else {
		os.Exit(0)
	}
}

func main() {
	todoList := TodoList{}
	for {
		Runner(&todoList)
	}
}
