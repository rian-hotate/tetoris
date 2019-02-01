package models

import (
	"sync"
)

const (
	HEIGHT = 20
	WIDTH  = 11
)

var Mu sync.Mutex

type Point struct {
	X int
	Y int
}
type Piece struct {
	Vector          Point
	Operation       Point
	Occupancy       []Point
	TargetType      int
	TargetOccupancy []Point
	Span            int
	Score           int
	End             bool
}
