package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)
import "math/rand"

func close_console() { //should be defered!
	termbox.Close()
}

func init_console() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer fmt.Print("Console successfully inited")
}

func put_char(c rune, x, y int) {
	termbox.SetCell(x, y, c, termbox.ColorDefault,
		termbox.Attribute(rand.Int()%8)+1)
}

func clear_console() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func flush_console() {
	termbox.Flush()
}
