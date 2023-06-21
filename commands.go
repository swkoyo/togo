package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var COMMANDS = []Command{
	"Add task",
	"Remove task",
	"Toggle task",
	"Clear all",
	"Clear done",
	"Exit",
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Command string

type Commands struct {
	commands []Command
	todoList *TodoList
	scanner  *bufio.Scanner
}

func NewCommands() Commands {
	prevTodos := make([]Todo, 0)
	var file *os.File
	if _, err := os.Stat("todo.txt"); os.IsNotExist(err) {
		f, err := os.Create("todo.txt")
		if err != nil {
			panic(err)
		}
		file = f
	} else {
		f, err := os.Open("todo.txt")
		if err != nil {
			panic(err)
		}
		file = f
	}

    defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[1] == 'x' {
				prevTodos = append(prevTodos, Todo{line[4:], true})
			} else {
				prevTodos = append(prevTodos, Todo{line[3:], false})
			}
		}
	}

	todoList := NewTodoList(prevTodos...)

	commands := Commands{todoList: &todoList, scanner: bufio.NewScanner(os.Stdin)}
	commands.Reset()
	return commands
}

func (c *Commands) Reset() {
	data := make([]Command, 0)
	for _, command := range COMMANDS {
		switch command {
		case "Remove task":
			fallthrough
		case "Toggle task":
			fallthrough
		case "Clear all":
			if !c.todoList.IsEmpty() {
				data = append(data, command)
			}
		case "Clear done":
			if c.todoList.HasDone() {
				data = append(data, command)
			}
		default:
			data = append(data, command)
		}
	}
	c.commands = data
}

func (c *Commands) WriteToFile() {
    file, err := os.OpenFile("todo.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    str := ""
    for _, todo := range c.todoList.todos {
        fmt.Println(todo)
        if todo.done {
            str += fmt.Sprintf("[x] %s\n", todo.task)
        } else {
            str += fmt.Sprintf("[] %s\n", todo.task)
        }
    }
    file.WriteString(str)
}

func (c *Commands) PrintInterface() {
	ClearScreen()
	c.todoList.Print()
	c.PrintCommands()
}

func (c *Commands) PrintCommands() {
	fmt.Println("----- COMMANDS -----")
	for i, command := range c.commands {
		fmt.Printf("%d. %s\n", i, command)
	}
	fmt.Println("--------------------")
}

func (c *Commands) AddTask() {
	c.PrintInterface()
	fmt.Printf("Enter new task: ")
	var task string
	if c.scanner.Scan() {
		task = c.scanner.Text()
		strings.Trim(task, " ")
	}
	for len(task) == 0 {
		c.PrintInterface()
		fmt.Printf("Empty task given. Enter new task: ")
		if c.scanner.Scan() {
			task = c.scanner.Text()
			strings.Trim(task, " ")
		}
	}
	c.todoList.Add(task)
    c.WriteToFile()
}

func (c *Commands) RemoveTask() {
	c.PrintInterface()
	fmt.Printf("Enter task number: ")
	var input string
	if c.scanner.Scan() {
		input = c.scanner.Text()
	}
	taskNum, err := strconv.Atoi(input)
	for err != nil || taskNum < 1 || taskNum > c.todoList.Size() {
		c.PrintInterface()
		fmt.Printf("Enter valid task number: ")
		if c.scanner.Scan() {
			input = c.scanner.Text()
			taskNum, err = strconv.Atoi(input)
		}
	}
	c.todoList.Remove(taskNum - 1)
    c.WriteToFile()
}

func (c *Commands) ClearAllTasks() {
	c.todoList.ClearAll()
    c.WriteToFile()
}

func (c *Commands) ClearDoneTasks() {
	c.todoList.ClearDone()
    c.WriteToFile()
}

func (c *Commands) ToggleTask() {
	c.PrintInterface()
	fmt.Printf("Enter task number: ")
	var input string
	if c.scanner.Scan() {
		input = c.scanner.Text()
	}
	taskNum, err := strconv.Atoi(input)
	for err != nil || taskNum < 0 || taskNum > c.todoList.Size() {
		c.PrintInterface()
		fmt.Printf("Enter valid task number: ")
		if c.scanner.Scan() {
			input = c.scanner.Text()
			taskNum, err = strconv.Atoi(input)
		}
	}
	c.todoList.Toggle(taskNum - 1)
    c.WriteToFile()
}

func (c *Commands) Runner() {
	c.PrintInterface()
	fmt.Printf("Enter command: ")
	var input string
	if c.scanner.Scan() {
		input = c.scanner.Text()
	}
	val, err := strconv.Atoi(input)
	for err != nil || val < 0 || val >= len(c.commands) {
		c.PrintInterface()
		fmt.Printf("Invalid command given. Enter valid command: ")
		if c.scanner.Scan() {
			input = c.scanner.Text()
		}
		val, err = strconv.Atoi(input)
	}
	command := c.commands[val]

	switch command {
	case "Add task":
		c.AddTask()
	case "Remove task":
		c.RemoveTask()
	case "Toggle task":
		c.ToggleTask()
	case "Clear all":
		c.ClearAllTasks()
	case "Clear done":
		c.ClearDoneTasks()
	case "Exit":
		os.Exit(0)
	}
}
