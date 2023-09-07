package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		tasks, err := GetTasks()
		if err != nil {
			return err
		}
		for _, task := range tasks {
			var mark string
			if task.IsComplete {
				mark = "[x] %s\n"
			} else {
				mark = "[x] %s\n"
			}
			fmt.Fprintf(v, mark, task.Name)
		}
		v.Wrap = true
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func taskIsComplete(task string) bool {
	return rune(task[1]) == 'x'
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, 1, false)
	return nil
}

func moveUp(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, -1, false)
	return nil
}

func toggleTask(g *gocui.Gui, v *gocui.View) error {
	x, y := v.Cursor()
	task, err := v.Line(y)
	if err != nil {
		return err
	}
	isComplete := taskIsComplete(task)
	v.SetCursor(1, y)
	v.EditDelete(false)
	if isComplete {
		v.EditWrite(' ')
	} else {
		v.EditWrite('x')
	}
	v.SetCursor(x, y)
	return nil
}

func GuiInit() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'k', gocui.ModNone, moveUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'j', gocui.ModNone, moveDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, toggleTask); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone,

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
