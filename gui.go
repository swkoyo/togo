package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, maxX, 10); err != nil {
	if err != gocui.ErrUnknownView {
			return err
		}
		logo, err := ioutil.ReadFile("logo.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(v, string(logo))
	}
	if v, err := g.SetView("main", -1, 10, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if jsonFile, err := os.Open("test.json"); err != nil {
			if err != nil {
				panic(err)
			}
			defer jsonFile.Close()

		}
		fmt.Fprintf(v, "Hello\nThere")
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func GuiInit() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone,

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
