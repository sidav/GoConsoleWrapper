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

const (
	WHITE   = 0
	BLACK   = 1
	RED     = 2
	GREEN   = 3
	YELLOW  = 4
	BLUE    = 5
	MAGENTA = 6
	CYAN    = 7
	BEIGE   = 8
)

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

func SetColor(fg int, bg int) {
	fg_color = termbox.Attribute(fg)
	bg_color = termbox.Attribute(bg)
}

func SetFgColor(fg int) {
	fg_color = termbox.Attribute(fg)
}

func SetBgColor(bg int) {
	bg_color = termbox.Attribute(bg)
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
	termbox.Clear(WHITE, BLACK)
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

func ReadKey() string {
	ev := <-event_queue
	if ev.Key == termbox.KeyEsc {
		return "ESCAPE"
	}
	if ev.Key == termbox.KeyEnter {
		return "ENTER"
	}
	if ev.Key == termbox.KeyTab {
		return "TAB"
	}
	if ev.Type == termbox.EventKey {
		return string(ev.Ch)
	}
	return "KEY_EMPTY_WTF_HAPPENED"
}
