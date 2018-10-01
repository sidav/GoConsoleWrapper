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

	var (
		mapw, maph, splits, splitprob, splitratio, hprob int
	)

	fmt.Println("Enter map width (<= 80 for Windows!): ")
	fmt.Scanf("%d", &mapw)
	fmt.Println("Enter map height")
	fmt.Scanf("%d", &maph)
	if mapw == 0 {
		mapw = 80
	}
	if maph == 0 {
		maph = 20
	}
	fmt.Println("Enter splits amount (0 - over9000)")
	fmt.Scanf("%d", &splits)
	fmt.Println("Enter split probability (in percent)")
	fmt.Scanf("%d", &splitprob)
	fmt.Println("Enter minimum split ratio (in percent)")
	fmt.Scanf("%d", &splitratio)
	fmt.Println("Enter horizontal split probability (in percent)")
	fmt.Scanf("%d", &hprob)

	Randomize()
	fuck := GenerateDungeon(mapw, maph, splits, splitprob, splitratio, hprob)
	for x := 0; x < mapw; x++ {
		for y := 0; y < maph; y++ {
			cw.Put_char(fuck.getCell(x, y), x, y)
		}
	}
	cw.Flush_console()
	for key_pressed := cw.ReadKey(); ; {
		if key_pressed == "ENTER" {
			break
		}
	}
	// test_wrapper()
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
