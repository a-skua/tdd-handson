package main

import (
	"fmt"
	"github.com/a-skua/tdd-handson/lifegame"
	"time"
)

func show(table [][]lifegame.State) {
	for _, row := range table {
		for _, state := range row {
			if state == lifegame.State_Live {
				fmt.Print("■ ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	x, y := lifegame.X(6), lifegame.Y(6)
	lifegame := lifegame.New(lifegame.NewCell(x, y, []lifegame.State{
		false, false, false, false, false, false,
		false, true, true, false, false, false,
		false, true, true, false, false, false,
		false, false, false, true, true, false,
		false, false, false, true, true, false,
		false, false, false, false, false, false,
	}))

	for {
		show(lifegame.Table())
		fmt.Println()
		lifegame.Next()
		time.Sleep(200 * time.Millisecond)
	}
}
