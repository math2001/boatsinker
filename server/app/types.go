package app

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d>", p.X, p.Y)
}

type Boat struct {
	Size int
	Pos  Point
	Rot  int // 0, 1 -> horizontal vertical
}

type Board struct {
	Hits  []Point
	Boats []Boat
}

type Player struct {
	Name  string `json:"name"`
	Board Board
}
