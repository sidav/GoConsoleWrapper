package console_wrapper

import (
	"github.com/nsf/termbox-go"
)
import "math/rand"

func Close_console() { //should be defered!
	termbox.Close()
}

func Init_console() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func Put_char(c rune, x, y int) {
	termbox.SetCell(x, y, c, termbox.ColorDefault,
		termbox.Attribute(rand.Int()%8)+1)
}

func Clear_console() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Flush_console() {
	termbox.Flush()
}
