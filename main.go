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

	//fmt.Print("Enter map width (<= 80 for Windows!) ")
	//fmt.Scanf("%d", &mapw)
	//fmt.Printf(": %d\n", mapw)
	//
	//fmt.Print("Enter map height")
	//fmt.Scanf("%d", &maph)
	//fmt.Printf(": %d\n", maph)
	//fmt.Print("Enter splits amount (0 - over9000)")
	//fmt.Scanf("%d", &splits)
	//fmt.Printf(": %d\n", splits)
	//
	//fmt.Print("Enter split probability (in percent)")
	//fmt.Scanf("%d", &splitprob)
	//fmt.Printf(": %d\n", splitprob)
	//
	//fmt.Print("Enter minimum split ratio (in percent)")
	//fmt.Scanf("%d", &splitratio)
	//fmt.Printf(": %d\n", splitratio)
	//
	//fmt.Print("Enter horizontal split probability (in percent)")
	//fmt.Scanf("%d", &hprob)
	//fmt.Printf(": %d\n", hprob)
	//fmt.Print("You can press ENTER to re-generate map. Press any key...")
	//cw.ReadKey()

	if mapw == 0 {
		mapw = 80
	}
	if maph == 0 {
		maph = 20
	}

	Randomize()

outerloop:
	for {
		cw.Clear_console()
		fuck := GenerateDungeon(mapw, maph, splits, splitprob, splitratio, hprob, 3)
		for x := 0; x < mapw; x++ {
			for y := 0; y < maph; y++ {
				chr := fuck.getCell(x, y)
				switch chr {
				case '+':
					cw.Set_color(cw.BLUE, nil)
				case '~':
					cw.Set_color(cw.RED, nil)
				default:
					cw.Set_color(cw.BEIGE, nil)
				}
				cw.Put_char(chr, x, y)
			}
		}
		cw.Flush_console()
		for key_pressed := cw.ReadKey(); key_pressed != "ENTER"; {
			if key_pressed == "ESCAPE" {
				break outerloop
			} else {
				key_pressed = cw.ReadKey()
			}
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
