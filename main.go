package main

import (
	"github.com/nsf/termbox-go"
	models "github.com/rian-hotate/tetoris/models/tetoris"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	pieceCh := make(chan models.Piece)
	keyCh := make(chan termbox.Key)

	go drawLoop(pieceCh)
	go keyEventLoop(keyCh)

	controller(pieceCh, keyCh)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Close()
}
