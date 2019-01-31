package main

import (
    "github.com/nsf/termbox-go"
)

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }

    keyCh := make(chan termbox.Key)

    go keyEventLoop(keyCh)

    controller(keyCh)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    defer termbox.Close()
}
