package main

import (
	"flag"

	"github.com/rian-hotate/tetoris/app/tetoris"
	"github.com/rian-hotate/tetoris/game"
)

func main() {
	wordPtr := flag.String("addr", "0.0.0.0:5000", "a string")
	flag.Parse()

	g := game.NewGame()
	err := g.Init()
	if err != nil {
		panic(err)
	}
	defer g.Close()

	// game app.
	app := tetoris.NewTetoris(*wordPtr)

	// run game app.
	g.Run(app)
}
