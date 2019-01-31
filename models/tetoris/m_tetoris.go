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
<<<<<<< HEAD
    Vector          Point
    Operation       Point
    Occupancy       []Point
    TargetType      int
    TargetOccupancy []Point
    Span            int
    Score           int
    HighScore       int
    End             bool
}
=======
	Vector          Point
	Operation       Point
	Occupancy       []Point
	TargetType      int
	TargetOccupancy []Point
	Wait            bool
	Score           int
	HighScore       int
	End             bool
}

type Time struct {
	Status bool
	Span   int
}
>>>>>>> fb9ad8766e0418ed71c4d4b81e2b84124f18a950
