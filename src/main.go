package main

import (
    "models/tetoris"
    "github.com/nsf/termbox-go"
)

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }

    pieceCh := make(chan models.Piece)
    keyCh := make(chan termbox.Key)
    timerCh := make(chan models.Time)

    go drawLoop(pieceCh)
    go keyEventLoop(keyCh)
    go timerLoop(timerCh, pieceCh)

    controller(pieceCh, keyCh, timerCh)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    defer termbox.Close()
}
