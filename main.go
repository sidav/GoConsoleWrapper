package main

// This file is for the lib test purposes only.

import (
	"GoConsoleWrapper/console_wrapper" //TODO: check why this shit doesn't want to work.
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	fmt.Print("I'M NOT DONE YET!")
	console_wrapper.Init_console()
	defer console_wrapper.Close_console()
	start_drawing()
}

func draw() {
	w, h := termbox.Size()
	console_wrapper.Clear_console()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			console_wrapper.Put_char('Ы', x, y)
		}
	}
	console_wrapper.Flush_console()
}

func start_drawing() {
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	draw()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			draw()
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
