package main

func main() {
    commands := NewCommands()
	for {
        commands.Reset()
        commands.Runner()
	}
}
