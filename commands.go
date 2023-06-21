package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "os/exec"
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
	commands     []Command
	todoList *TodoList
	scanner  *bufio.Scanner
}

func NewCommands(t *TodoList) Commands {
	commands := Commands{todoList: t, scanner: bufio.NewScanner(os.Stdin)}
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
}

func (c *Commands) ClearAllTasks() {
	c.todoList.ClearAll()
}

func (c *Commands) ClearDoneTasks() {
    c.todoList.ClearDone()
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
