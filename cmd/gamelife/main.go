package main

import (
	"game-life/univers"
	"time"
)

func main() {
	u := univers.NewUnivers()

	for i := 0; i < 600; i++ {
		u.Show()
		u.Computing()
		time.Sleep(time.Second / 5)
	}
}
