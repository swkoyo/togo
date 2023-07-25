package main

func main() {
	InitDB()
	defer DB.Close()
	GuiInit()
}
