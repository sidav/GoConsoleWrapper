package main

// This file is for the lib test purposes only.

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	fmt.Print("I'M NOT DONE YET!")
	init_console()
	defer close_console()
	start_drawing()
}

func draw() {
	w, h := termbox.Size()
	clear_console()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			put_char('Ы', x, y)
		}
	}
	flush_console()
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
