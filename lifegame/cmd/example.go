package main

import (
	"fmt"
	"github.com/a-skua/tdd-handson/lifegame"
	"time"
)

const (
	D = lifegame.Die
	L = lifegame.Live
)

func show(table [][]lifegame.State) {
	for _, row := range table {
		for _, state := range row {
			if state.IsLive() {
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
		D, D, D, D, D, D,
		D, L, L, D, D, D,
		D, L, L, D, D, D,
		D, D, D, L, L, D,
		D, D, D, L, L, D,
		D, D, D, D, D, D,
	}))

	for {
		show(lifegame.Table())
		fmt.Println()
		lifegame.Next()
		time.Sleep(200 * time.Millisecond)
	}
}
