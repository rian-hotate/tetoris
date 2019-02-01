package main

import ()

func main() {
    g := NewGame()
    err := g.init()
    if err != nil {
        panic(err)
    }
    defer g.close()

    g.run()
}
