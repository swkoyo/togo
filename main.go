package main

func main() {
	todoList := TodoList{}
    commands := NewCommands(&todoList)
	for {
        commands.Reset()
        commands.Runner()
	}
}
