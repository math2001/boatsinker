package app

import (
	"fmt"
)

// Point represents a 2d vector
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d>", p.X, p.Y)
}

// Boat represents a boat no a board
type Boat struct {
	Size int
	Pos  Point
	Rot  int // 0, 1 -> horizontal vertical
}

// Board represents a board with the boats and hits
type Board struct {
	Hits  []Point
	Boats []Boat
}

// Player represents a client
type Player struct {
	Name  string `json:"name"`
	Board Board
}
