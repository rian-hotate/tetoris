package tetoris

import (
	"fmt"

	"github.com/rian-hotate/tetoris/game"
)

type TetorisApp struct {
	addr string // network addr.
	//network *RTCConnection
}

func (t *TetorisApp) Init() {
	// t.network := NewWebRTC(t.addr);
	fmt.Println(t.addr)
}

func (t *TetorisApp) Update() {
}
func (t *TetorisApp) Draw() {
}
func (t *TetorisApp) Close() {
	// t.network.Close()
}

// NewTetoris is to generate new game.
func NewTetoris(addr string) game.AppFactory {
	var factory game.AppFactory
	factory = &TetorisApp{addr: addr}
	return factory
}
