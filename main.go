package main

// This file is for the lib test purposes only.

import (
	cw "GoConsoleWrapper/console_wrapper" //TODO: check why this shit doesn't want to work.
	"time"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	go func() {
		cw.Run_event_listener()
	}()

	test_wrapper()
}

func test_wrapper() {
	for i := 0; i < 20; i++ {
		fuck := cw.Read_key_char()
		cw.Set_color(i, nil)
		cw.Put_char(fuck, 0, 0)
		cw.Flush_console()
		time.Sleep(200 * time.Millisecond)
	}
}
