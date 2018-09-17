package console_wrapper

import (
	"github.com/nsf/termbox-go"
)

var (
	fg_color    = termbox.ColorWhite
	bg_color    = termbox.ColorBlack
	event_queue = make(chan termbox.Event)
)

/* PUBLIC SHIT BELOW */

func Close_console() { //should be defered!
	termbox.Close()
}

func Init_console() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	event_queue = make(chan termbox.Event)
}

func Run_event_listener() { // should be run as go func() {}()
	for {
		event_queue <- termbox.PollEvent()
	}
}

func Set_color(fg int, bg *int) {
	fg_color = termbox.Attribute(fg)
	if bg != nil {
		bg_color = termbox.Attribute(*bg)
	}

}

func Put_char(c rune, x, y int) {
	termbox.SetCell(x, y, c, fg_color, bg_color)
}

func Put_string(s string, x, y int) {
	length := len([]rune(s))
	for i := 0; i < length; i++ {
		Put_char([]rune(s)[i], x+i, y)
	}
}

func Clear_console() {
	termbox.Clear(fg_color, bg_color)
}

func Flush_console() {
	termbox.Flush()
}

func Await_keypress() rune {
	for {
		ev := <-event_queue
		if ev.Type == termbox.EventKey {
			return ev.Ch
		}
	}

}

func Read_key_char() rune {
	ev := <-event_queue
	return ev.Ch
}
