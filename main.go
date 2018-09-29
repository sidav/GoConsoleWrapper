package main

// This file is for the lib test purposes only.

import (
	cw "GoConsoleWrapper/console_wrapper"
	"fmt"
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
	for {
		fmt.Println(cw.ReadKey())
	}
	//for i := 0; i < 20; i++ {
	//	cw.Set_color(i, nil)
	//	cw.Put_string(fmt.Sprintf("This is %d", i), 0, i)
	//}
	//cw.Flush_console()
	//for i := cw.ReadKey(); i != 'e'; {
	//
	//}
}
